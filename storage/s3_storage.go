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
	"bytes"
	"sync"
	"fmt"
)

const(
	common_piece_size = 1024*1024*4 //4MB
	max_piece_size = 1024*1024*256 //256MB
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
		res,err := s.cli.ListObjects(s.Bucket,file)
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
	out,err := s.cli.ListObjects(s.Bucket,dir)
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
	out,err := s.cli.ListAllObjects(s.Bucket,oldname)
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
	out,err := s.cli.ListAllObjects(s.Bucket,dir)
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
	var cnt int = 0
	var cur_piece_size int64 = common_piece_size
	var uid string
	var wg sync.WaitGroup
	var ch chan *util.FBS3PieceInfo
	var sz int64

	for{
		rd := io.LimitReader(rd,cur_piece_size)

		var buf []byte = make([]byte,cur_piece_size,cur_piece_size)
		n,err := rd.Read(buf[:])
		if err != nil && err != io.EOF{
			return 0,err
		}

		if cnt != 0 && n == 0{
			fmt.Println("mutil file finish...")
			break
		}

		sz += int64(n)

		if cnt == 0{
			if int64(n) < cur_piece_size{
				fmt.Println("single file put...")
				_,err := s.cli.PutObject(s.Bucket,filename,bytes.NewReader(buf[:n]))
				return sz,err
			}else{
				out,err := s.cli.CreateMutilPartUpload(s.Bucket,filename)
				if err != nil{
					return 0,err
				}

				uid = *out.UploadId
				ch = make(chan *util.FBS3PieceInfo,1024)
			}
		}

		wg.Add(1)

		go func(data []byte,part int){
			defer wg.Done()

			out,err := s.cli.UploadPartObject(s.Bucket,filename,part,uid,bytes.NewReader(data))
			if err != nil{
				return
			}

			ch <- &util.FBS3PieceInfo{int64(part),*out.ETag}
		}(buf[:n],cnt+1)

		cnt++
		cur_piece_size = cur_piece_size*2
		if cur_piece_size > max_piece_size{
			cur_piece_size = max_piece_size
		}
	}

	wg.Wait()
	close(ch)
	var ps []*util.FBS3PieceInfo
	for v := range ch{
		ps = append(ps,v)
	}

	_,err := s.cli.CompleteMultiPartUpload(s.Bucket,filename,ps,uid)
	if err != nil{
		return 0,err
	}

	return sz,err
}
