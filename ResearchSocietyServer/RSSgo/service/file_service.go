package service

import (
	"log"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"hkc.ink/rss/database"
	"hkc.ink/rss/model"
	"hkc.ink/rss/utils"
)

var fileUtil utils.FileUtil

type FileService struct {
}

func (FileService) Upload(ctx *gin.Context) (string, error) {
	file, err := ctx.FormFile("file")
	if err != nil {
		log.Println("获取文件失败")
		return "", err
	} else {
		log.Println("接收的数据:", file.Filename)
		// //获取文件名称
		// log.Println(file.Filename)
		// //文件大小
		// log.Println(file.Size)
		//获取文件的后缀名
		ext := path.Ext(file.Filename)
		//根据当前时间戳生成一个新的文件名
		fileNameInt := time.Now().Unix()
		fileNameStr := strconv.FormatInt(fileNameInt, 10)
		fileName := fileNameStr + ext
		//保存上传文件
		filePath := filepath.Join(fileUtil.Mkdir("upload"), "/", fileName)
		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
			log.Println("文件保存失败")
			return "", err
		}
		//保存文件路径到数据库
		db := database.GetDB()
		file := &model.File{
			FileName: fileName,
			FilePath: filePath,
		}
		result := db.Create(file)
		if result.Error != nil {
			return "", result.Error
		}

		return fileName, nil
	}
}

func (FileService) GetFilePath(fileName string) string {
	db := database.GetDB()
	file := &model.File{}
	db.Where("file_name=?", fileName).First(file)
	if file.FilePath != "" {
		return file.FilePath
	} else {
		return ""
	}
}
