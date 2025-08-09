package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry // 引导类路径
	extClasspath  Entry // 扩展类路径
	userClasspath Entry // 用户类路径
}

// Parse 解析 JRE 和类路径选项，返回一个 Classpath 实例
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

// ReadClass 读取指定类的字节码
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class" // 添加 .class 后缀
	// 尝试从引导类路径读取类
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	// 尝试从扩展类路径读取类
	if data, entry, err := self.userClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	// 尝试从用户类路径读取类
	return self.userClasspath.readClass(className)
}

// String 返回 Classpath 的字符串表示形式
func (self *Classpath) String() string {
	return self.userClasspath.String()
}

// Parse 解析 JRE 和类路径选项，返回一个 Classpath 实例
func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	// 获取 JRE 目录
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

// parseUserClasspath 解析用户指定的类路径选项
func (self *Classpath) parseUserClasspath(cpOption string) {
	// 如果用户没有指定类路径选项，则默认为当前目录
	if cpOption == "" {
		cpOption = "."
	}
	// 如果类路径选项包含路径分隔符，则创建一个复合条目
	self.userClasspath = newEntry(cpOption)
}

// getJreDir 获取 JRE 的目录路径
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

// exists 检查指定路径是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}
