package meta

import "fileStore/db"

type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas =make(map[string]FileMeta)
}

//新增/更新文件元信息
func UpdateFileMeta(filemeta FileMeta) {
	fileMetas[filemeta.FileSha1]=filemeta

}

func GetFileMeta(filesha1 string) FileMeta {
	return fileMetas[filesha1]

}

func RemoveMeta(filesha1 string)  {
	delete(fileMetas,filesha1)
}

func UpdateFileMetaDB(filemeta FileMeta) bool{
	return db.SaveFile(filemeta.FileName,filemeta.FileSha1,filemeta.FileSize,filemeta.Location)
}