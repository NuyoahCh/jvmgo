package classpath

import "archive/zip"
import "errors"
import "io/ioutil"
import "path/filepath"

// ZipEntry2 代表一个 ZIP 文件条目,表示 JAR 或 ZIP 格式的类路径
type ZipEntry2 struct {
	absPath string          // 存放 zip 文件的绝对路径
	zipRC   *zip.ReadCloser // zip.ReadCloser 用于复用已打开的 ZIP 文件，避免重复打开
}

// newZipEntry2 创建一个新的 ZipEntry2 实例
func newZipEntry2(path string) *ZipEntry2 {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &ZipEntry2{absPath, nil}
}

// readClass 读取指定类的字节码
func (self *ZipEntry2) readClass(className string) ([]byte, Entry, error) {
	if self.zipRC == nil {
		err := self.openJar()
		if err != nil {
			return nil, nil, err
		}
	}

	classFile := self.findClass(className)
	if classFile == nil {
		return nil, nil, errors.New("class not found: " + className)
	}

	data, err := readClass(classFile)
	return data, self, err
}

// todo: close zip
// openJar 打开文件
func (self *ZipEntry2) openJar() error {
	// 之后复用已打开的 zip.ReadCloser，避免重复打开，提升了读取效率。
	r, err := zip.OpenReader(self.absPath)
	if err == nil {
		self.zipRC = r
	}
	return err
}

// findClass 在 ZIP 文件中查找指定的类文件
func (self *ZipEntry2) findClass(className string) *zip.File {
	for _, f := range self.zipRC.File {
		if f.Name == className {
			return f
		}
	}
	return nil
}

// readClass 读取指定类文件的内容
func readClass(classFile *zip.File) ([]byte, error) {
	rc, err := classFile.Open()
	if err != nil {
		return nil, err
	}
	// read class data
	data, err := ioutil.ReadAll(rc)
	rc.Close()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (self *ZipEntry2) String() string {
	return self.absPath
}
