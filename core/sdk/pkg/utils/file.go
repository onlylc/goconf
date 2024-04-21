package utils

import (
	"log"
	"net/http"
	"os"
	"path"
)

// GetExt 获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

func CheckExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

func IsNotExitsMKDir(src string) error {
	if exist := !CheckExist(src); exist == false {
		if err := MKDir(src); err != nil {
			return err
		}
	}
	return nil
}

func MKDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// GetType 获取文件类型
func GetType(p string) (string, error) {
	file, err := os.Open(p)
	defer file.Close()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	buff := make([]byte, 512)

	_, err = file.Read(buff)

	if err != nil {
		log.Println(err)
	}

	filetype := http.DetectContentType(buff)

	//ext := GetExt(p)
	//var list = strings.Split(filetype, "/")
	//filetype = list[0] + "/" + ext
	return filetype, nil
}
