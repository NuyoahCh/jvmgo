package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

// newWildcardEntry 创建一个新的通配符条目,表示目录下的所有 JAR 文件
func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] // 去掉末尾的星号 remove *
	compositeEntry := []Entry{}
	// 根据后缀名选出JAR文件，并且返回SkipDir跳过 子目录(通配符类路径不能递归匹配子目录下的JAR文件)
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // 如果有错误，直接返回
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	// 遍历 baseDir 目录下的所有文件和子目录
	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}
