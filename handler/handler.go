package handler

import (
	"encoding/json"
	"file/meta"
	"file/util"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func UploadHandler(w http.ResponseWriter,r *http.Request)  {
	if r.Method=="GET"{
		//返回上传html
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w,"internel server error")
			return
		}
		io.WriteString(w,string(data))

	}else if r.Method=="POST" {
		//接收文件流以及存储到本地
		file, header, err := r.FormFile("file")
		if err != nil {
			fmt.Println("failed to get data,err:%s",err.Error())
			return
		}
		defer file.Close()
		fileMeta := meta.FileMeta{
			FileName: header.Filename,
			Location: "/tmp/" + header.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		create, err := os.Create(fileMeta.FileName)
		if err != nil {
			fmt.Println("failed to create file,err:%s",err.Error())
			return
		}
		defer create.Close()
		fileMeta.FileSize, err = io.Copy(create, file)
		if err != nil {
			fmt.Println("failed to save data into file,err:%s",err.Error())
			return
		}
		create.Seek(0,0)
		fileMeta.FileSha1 = util.FileSha1(create)
		meta.UpdateFileMeta(fileMeta)

		http.Redirect(w,r,"/filestore/upload/succeed",http.StatusFound)

	}
	
}
func UploadSucceedHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,"Upload finished!")
}

func GetFileMetaHandler(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	filemeta := r.Form["filehash"][0]
	data := meta.GetFileMeta(filemeta)
	bytes, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}
func DownloadFileHandler(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	get := r.Form.Get("filehash")
	fileMeta := meta.GetFileMeta(get)
	open, err := os.Open(fileMeta.Location)
	if err != nil {
		fmt.Println(1)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer open.Close()
	bytes, err := ioutil.ReadAll(open)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	// attachment表示文件将会提示下载到本地，而不是直接在浏览器中打开
	w.Header().Set("content-disposition", "attachment; filename=\""+fileMeta.FileName+"\"")


	w.Write(bytes)

}

func FileMetaUpdateHandler(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	operation := r.Form.Get("OP")
	filehash := r.Form.Get("filehash")
	newFileName := r.Form.Get("filename")
	if operation!="0"{
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if r.Method!="POST"{
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}


	fileMeta := meta.GetFileMeta(filehash)
	fileMeta.FileName=newFileName
	//meta.UpdateFileMeta(fileMeta)
	_ = meta.UpdateFileMetaDB(fileMeta)
	bytes, err := json.Marshal(fileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)

}

func DeleteHandler(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	filesha1 := r.Form.Get("filehash")

	getFileMeta := meta.GetFileMeta(filesha1)
	os.Remove(getFileMeta.Location)
	meta.RemoveMeta(filesha1)
	w.WriteHeader(http.StatusOK)
}