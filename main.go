package main

import (
	"file/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/filestore/upload",handler.UploadHandler)
	http.HandleFunc("/filestore/upload/succeed",handler.UploadSucceedHandler)
	http.HandleFunc("/filestore/meta",handler.GetFileMetaHandler)
	http.HandleFunc("/filestore/download",handler.DownloadFileHandler)
	http.HandleFunc("/filestore/update",handler.FileMetaUpdateHandler)
	http.HandleFunc("/filestore/delete",handler.DeleteHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("failed to start server,err:%s",err.Error())
	}

}
