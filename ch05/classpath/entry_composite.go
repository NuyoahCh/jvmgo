package classpath

import (
	"errors"
	"strings"
)

// CompositeEntry 代表一个复合条目,表示多个类路径条目的组合
type CompositeEntry []Entry

// newCompositeEntry 创建一个新的 CompositeEntry 实例
func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	// 使用路径分隔符将路径列表分割成多个路径
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

// readClass 遍历所有条目，读取指定类的字节码
func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	// 遍历 CompositeEntry 中的每个条目
	for _, entry := range self {
		// 调用条目的 readClass 方法读取类的字节码
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, err
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

// String 返回 CompositeEntry 的字符串表示形式
func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
