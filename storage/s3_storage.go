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
	"strings"
	"bufio"
	"bytes"
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
	info,err := s.Stat(dir)
	if err != nil{
		return false,err
	}

	return info.IsDir,nil
}

func (s *S3Storage)ListFile(dir string)([]*FTPFileInfo,error){
	out,err := s.cli.ListFile(s.Bucket,dir)
	if err != nil{
		return nil,err
	}

	var ret []*FTPFileInfo
	for _,v := range out.CommonPrefixes{
		i := newFTPFileInfo()
		i.IsDir = true
		i.Name = filepath.Base(*v.Prefix)
		i.Mode = os.ModeDir
		i.ModTime = time.Now()

		ret = append(ret,i)
	}

	for _,v := range out.Contents{
		i := newFTPFileInfo()
		i.IsDir = false
		i.Name = filepath.Base(*v.Key)
		i.Mode = os.ModePerm
		i.Size = *v.Size
		i.ModTime = *v.LastModified

		ret = append(ret,i)
	}

	return ret,nil
}

func (s *S3Storage)ReName(oldname string, newname string)error{
	out,err := s.cli.ListAllFile(s.Bucket,oldname)
	if err != nil{
		return err
	}

	for _,v := range out.Contents{
		newkey := strings.Replace(*v.Key,oldname,newname,1)
		_,err = s.cli.CopyObject(s.Bucket,newkey,s.Bucket,*v.Key)
		if err != nil{
			return err
		}

		_,err = s.cli.DeleteObject(s.Bucket,*out.Name)
		if err != nil{
			return err
		}
	}

	return nil
}

func (s *S3Storage)MKDir(dir string)error{
	return nil
}

func (s *S3Storage)DeleteFile(filename string)error{
	_,err := s.cli.DeleteObject(s.Bucket,filename)

	return err
}

func (s *S3Storage)DeleteDir(dir string)error{
	out,err := s.cli.ListAllFile(s.Bucket,dir)
	if err != nil{
		return err
	}

	for _,v := range out.Contents{
		_,err = s.cli.DeleteObject(s.Bucket,*v.Key)
		if err != nil{
			return err
		}
	}

	return nil
}

func (s *S3Storage)GetFile(filename string, wr io.Writer)(int64,error){
	out,err := s.cli.DownloadObject(s.Bucket,filename)
	if err != nil{
		return 0,err
	}

	return io.Copy(wr,out.Body)
}

func (s *S3Storage)StoreFile(filename string, rd io.Reader)(int64,error){
	return 0,nil
}
