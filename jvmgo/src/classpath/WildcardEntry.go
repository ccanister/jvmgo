package classpath

import (
	"os"
	"path/filepath"
	"strings"
)


func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1]	//去掉*
	compositeEntry := []Entry{}	//切片

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		//跳过子目录
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	//遍历当前文件夹下所有文件
	filepath.Walk(baseDir, walkFn)

	return compositeEntry
}
