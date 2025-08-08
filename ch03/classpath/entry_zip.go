package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ZipEntry 代表一个 ZIP 文件条目,表示 JAR 或 ZIP 格式的类路径
type ZipEntry struct {
	absPath string // 存放 zip 文件的绝对路径
}

// newZipEntry 创建一个新的 ZipEntry 实例
func newZipEntry(path string) *ZipEntry {
	// 获取给定路径的绝对路径
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

// readClass 读取指定类的字节码
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	// 确保 className 以 .class 结尾，首先打开 ZIP 文件
	r, err := zip.OpenReader(self.absPath)
	// 如果打开 ZIP 文件失败，返回错误
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	// todo 此方法存在性能问题，entry_zip2 进行完善升级
	for _, f := range r.File {
		// 找到 class 文件
		if f.Name == className {
			// 打开 class 文件
			rc, err := f.Open()
			// 如果打开 class 文件失败，返回错误
			if err != nil {
				return nil, nil, err
			}
			// 每次读取 class 文件时都会重新打开一次 zip 文件，频繁的文件打开和关闭会导致性能下降
			defer rc.Close()
			// 读取 class 文件的内容
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

// String 返回 ZipEntry 的字符串表示形式
func (self *ZipEntry) String() string {
	return self.absPath
}
