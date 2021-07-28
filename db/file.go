package db

import (
	"file/db/mysql"
	"fmt"
)

func SaveFile(filename string,filehash string,filesize int64,fileaddr string) bool {
	stmt, err := mysql.DBConn().Prepare("insert into tbl_file(file_sha1,file_name,file_size,file_addr,status) values (?,?,?,?,1)")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer stmt.Close()
	result, err := stmt.Exec(filehash, filename, filesize, fileaddr)
	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if affected<=0{
		fmt.Println("uploaded before")
		return true
	}
	return false

}
