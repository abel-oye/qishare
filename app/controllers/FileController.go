package controllers

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"github.com/robfig/revel"
	"io"
	"os"
	"path"
	"strings"
)

type FileController struct {
	Application
}

//下载
func (this *FileController) Download(code string) revel.Result {
	//判断是否需要权限下载

	//通过code查询出文件路径

	return this.Redirect(path.Join("/attach", code))
}

//上传
func (this *FileController) Upload(code string) revel.Result {
	//判断是否需要权限下载

	//通过code查询出文件路径

	return this.Redirect(path.Join("/attach", code))
}

//上传文件
/**
*@return 保存的文件名称（不含后缀）
*@return 上载的文件名称
*@return 上载的文件后缀
 */
func saveFile(req *revel.Request, formFileName string, savepath string) (string, string, string) {
	saveFileName := strings.Replace(uuid.NewUUID().String(), "-", "", -1)

	file, fh, err := req.FormFile(formFileName)

	if err != nil {
		fmt.Println(err)
	}
	//获得文件后缀
	suffix := path.Ext(fh.Filename)

	defer file.Close()
	//创建保存路径
	os.MkdirAll(savepath, os.ModePerm)
	f, err1 := os.OpenFile(savepath+"\\"+saveFileName+suffix, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err1 != nil {
		fmt.Println(err1)
	}
	defer f.Close()
	io.Copy(f, file)
	return saveFileName, fh.Filename, suffix
}

//删除文件
func removeFile() {

}
