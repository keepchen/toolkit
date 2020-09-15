package file

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

//File 结构体
type File struct{}

//NewFile 实例化
func NewFile() *File {
	return new(File)
}

//Unzip 解压zip文件
//
//@param archive 压缩包文件地址
//
//@param target 解压地址
//返回文件列表和错误信息
func (*File) Unzip(archive, target string) ([]string, error) {
	var fileList []string
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return fileList, err
	}
	err = os.MkdirAll(target, 0755)
	if err != nil {
		return fileList, err
	}
	for _, file := range reader.File {
		//这里的file.Name是带有父级文件夹名称的，因此需要剔除
		split := strings.Split(file.Name, string(filepath.Separator))
		filename := split[len(split)-1]
		path := filepath.Join(target, filename)
		if file.FileInfo().IsDir() {
			err = os.MkdirAll(path, file.Mode())
			if err != nil {
				return fileList, err
			}
			continue
		}
		fileList = append(fileList, filename)
		fileReader, err := file.Open()
		if err != nil {
			return fileList, err
		}
		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			_ = fileReader.Close()
			return fileList, err
		}
		_, err = io.Copy(targetFile, fileReader)
		if err != nil {
			_ = fileReader.Close()
			return fileList, err
		}
	}

	return fileList, err
}
