/*
author: foolbread
file: storage/s3_storage.go
date: 2017/9/27
*/
package storage

import (
	"io"
	"github.com/foolbread/fbftp/util"
	"path/filepath"
	"time"
	"os"
)

type S3Storage struct {
	Bucket string

	cli *util.FBS3Client
}

func NewS3Storage(acckey string,seckey string,endpoint string,token string,bucket string)*S3Storage{
	r := new(S3Storage)
	r.Bucket = bucket
	r.cli = util.NewFBS3Client(acckey,seckey,endpoint,token)

	return r
}

func (s *S3Storage)Stat(file string)(*FTPFileInfo,error){
	var info *FTPFileInfo = newFTPFileInfo()
	res,err := s.cli.HeadObject(s.Bucket,file)
	if err != nil{
		//may be dir
		res,err := s.cli.ListFile(s.Bucket,file)
		if err != nil || len(res.CommonPrefixes) == 0{
			return nil,os.ErrNotExist
		}

		info.IsDir = true
		info.ModTime = time.Now()
		info.Mode = os.ModeDir
	}else{
		info.IsDir = false
		info.Name = filepath.Base(file)
		info.Size = *res.ContentLength
		info.ModTime = *res.LastModified
		info.Mode = os.ModePerm
	}
	return info,nil
}

func (s *S3Storage)ChangeDir(dir string)(bool,error){
	return true,nil
}

func (s *S3Storage)ListFile(dir string)([]*FTPFileInfo,error){
	return nil,nil
}

func (s *S3Storage)ReName(oldname string, newname string)error{
	return nil
}

func (s *S3Storage)MKDir(dir string)error{
	return nil
}

func (s *S3Storage)DeleteFile(filename string)error{
	return nil
}

func (s *S3Storage)DeleteDir(dir string)error{
	return nil
}

func (s *S3Storage)GetFile(filename string, wr io.Writer)(int64,error){
	return 0,nil
}

func (s *S3Storage)StoreFile(filename string, rd io.Reader)(int64,error){
	return 0,nil
}
