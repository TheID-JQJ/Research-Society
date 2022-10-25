package utils

import (
	"os"
	"path/filepath"
	"time"
)

type FileUtil struct {
}

func (FileUtil) Mkdir(basePath string) string {
	//	1.获取当前时间,并且格式化时间
	folderName := time.Now().Format("2006/01/02")
	folderPath := filepath.Join(basePath, folderName)
	os.MkdirAll(folderPath, os.ModePerm)
	return folderPath
}
