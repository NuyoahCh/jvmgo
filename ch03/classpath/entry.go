package classpath

import (
	"os"
	"strings"
)

// pathListSeparator 存放路径分隔符
const pathListSeparator = string(os.PathListSeparator)

// Entry 接口定义了一个类路径条目的行为
type Entry interface {
	// ReadClass 读取指定类的字节码
	readClass(className string) ([]byte, Entry, error)

	// String 返回条目的字符串表示形式
	String() string
}

// newEntry 根据给定的路径创建一个新的 Entry 实例
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
