package apis

import (
	"encoding/base64"
	"errors"
	"fmt"
	"goconf/core/sdk/api"
	"goconf/core/sdk/pkg"
	"goconf/core/sdk/pkg/utils"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FileResponse struct {
	Size     int64  `json:"size"`
	Path     string `json:"path"`
	FullPath string `json:"full_path"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}

const path = "static/uploadfile/"

type File struct {
	api.Api
}

func (e File) UploadFile(c *gin.Context) {
	e.MakeContext(c)
	tag, _ := c.GetPostForm("type")
	urlPrefix := fmt.Sprintf("%s://%s/", "http", c.Request.Host)
	var fileResponse FileResponse

	switch tag {
	case "1":
		var done bool
		fileResponse, done = e.singleFile(c, fileResponse, urlPrefix)
		if done {
			return
		}
		e.OK(fileResponse, "上传成功")
		return
	case "2":
		multipartFile := e.multipleFile(c, urlPrefix)
		e.OK(multipartFile, "上传成功")
		return
	case "3":
		fileResponse = e.baseImg(c, fileResponse, urlPrefix)
		e.OK(fileResponse, "上传成功")
	default:
		var done bool
		fileResponse, done = e.singleFile(c, fileResponse, urlPrefix)
		if done {
			return
		}
		e.OK(fileResponse, "上传成功")
		return
	}

}

func (e File) baseImg(c *gin.Context, fileResponse FileResponse, urlPerfix string) FileResponse {
	files, _ := c.GetPostForm("file")
	file2list := strings.Split(files, ",")
	ddd, _ := base64.StdEncoding.DecodeString(file2list[1])
	guid := uuid.New().String()
	fileName := guid + ".jpg"
	err := utils.IsNotExitsMKDir(path)
	if err != nil {
		e.Error(500, errors.New(""), "初始化文件路径失败")
	}
	base64File := path + fileName
	_ = os.WriteFile(base64File, ddd, 0666)
	typeStr := strings.Replace(strings.Replace(file2list[0],"data:","", -1), ";base64", "", -1)
	fileResponse = FileResponse{
		Size: pkg.GetFileSize(base64File),
		Path: base64File,
		FullPath: urlPerfix + base64File,
		Name: "",
		Type: typeStr,
	}
	source, _ := c.GetPostForm("source")
	err = thirdUpload(source, fileName, base64File)
	if err != nil {
		e.Error(200, errors.New(""), "上传第三方失败")
		return fileResponse
	}
	if source != "1" {
		fileResponse.Path = "/static/uploadfile/" + fileName
		fileResponse.FullPath = "/static/uploadfile/" + fileName
	}
	return fileResponse
}

func (e File) singleFile(c *gin.Context, fileResponse FileResponse, urlPrefix string) (FileResponse, bool) {
	files, err := c.FormFile("file")

	if err != nil {
		e.Error(200, errors.New(""), "图片不能为空")
		return FileResponse{}, true
	}
	guid := uuid.New().String()

	fileName := guid + utils.GetExt(files.Filename)

	err = utils.IsNotExitsMKDir(path)
	if err != nil {
		e.Error(500, errors.New(""), "初始化文件路径失败")
	}
	singleFile := path + fileName
	_ = c.SaveUploadedFile(files, singleFile)
	fileType, _ := utils.GetType(singleFile)
	fileResponse = FileResponse{
		Size:     pkg.GetFileSize(singleFile),
		Path:     singleFile,
		FullPath: urlPrefix + singleFile,
		Name:     files.Filename,
		Type:     fileType,
	}
	fileResponse.Path = "/static/uploadfile/" + fileName
	fileResponse.FullPath = "/static/uploadfile/" + fileName
	return fileResponse, false

}

func (e File) multipleFile(c *gin.Context, urlPerfix string) []FileResponse {
	files := c.Request.MultipartForm.File["file"]
	source, _ := c.GetPostForm("source")
	var multipartFile []FileResponse
	for _, f := range files {
		guid := uuid.New().String()
		fileName := guid + utils.GetExt(f.Filename)
		err := utils.IsNotExitsMKDir(path)
		if err != nil {
			e.Error(500, errors.New(""), "初始化文件路径失败")
		}
		multipartFileName := path + fileName
		err1 := c.SaveUploadedFile(f, multipartFileName)
		fileType, _ := utils.GetType(multipartFileName)
		if err1 == nil {
			err := thirdUpload(source, fileName, multipartFileName)
			if err !=nil {
				e.Error(500, errors.New(""), "上传第三方失败")
			} else {
				fileResponse := FileResponse {
					Size: pkg.GetFileSize(multipartFileName),
					Path: multipartFileName,
					FullPath: urlPerfix + multipartFileName,
					Name: f.Filename,
					Type: fileType,
				}
				if source != "1" {
					fileResponse.Path = "/static/uploadfile/" + fileName
					fileResponse.FullPath = "/static/uploadfile/" + fileName
				}
				multipartFile = append(multipartFile, fileResponse)
			}
		}
	}
	return multipartFile
}


func thirdUpload(source string, name string, path string) error {
	switch source {
	case "2":
		return ossUpload("img/"+name, path)
	case "3":
		return qiniuUpload("img/"+name, path)
	}
	return nil
}

func ossUpload(name string, path string) error {
	// oss := file_store.ALiYunOSS{}
	fmt.Println(name,path)
	// return oss.UpLoad(name, path)
	return nil 
}

func qiniuUpload(name string, path string) error {
	// oss := file_store.ALiYunOSS{}
	// return oss.UpLoad(name, path)
	fmt.Println(name,path)
	return nil 
}