package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// DirEntry 代表一个目录条目,表示目录形式的类路径
type DirEntry struct {
	absDir string // 存放目录的绝对路径
}

// newDirEntry 创建一个新的 DirEntry 实例
func newDirEntry(path string) *DirEntry {
	// 获取给定路径的绝对路径
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err) // 如果获取绝对路径失败，抛出异常
	}
	return &DirEntry{absDir}
}

// readClass 读取指定类的字节码
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	// 把目录和 class 文件名拼成一个完整的路径
	fileName := filepath.Join(self.absDir, className)
	// 替换类名中的点为路径分隔符，并添加 .class 后缀
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

// String 返回 DirEntry 的字符串表示形式
func (self *DirEntry) String() string {
	return self.absDir
}
