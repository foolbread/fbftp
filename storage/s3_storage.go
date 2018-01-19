/*
author: foolbread
file: storage/s3_storage.go
date: 2017/9/27
*/
package storage

import (
	"io"
	"github.com/foolbread/fbftp/util"
	"github.com/foolbread/fbcommon/golog"
	"path/filepath"
	"time"
	"os"
	"strings"
	"bytes"
	"sync"
)

const(
	common_piece_size = 1024*1024*4 //4MB
	max_piece_size = 1024*1024*32 //256MB
)

func s3CleanFilePath(p string)string{
	return strings.TrimRight(strings.TrimLeft(p,"/"),"/")
}

func s3CleanDirPath(p string)string{
	if p == "/"{
		return strings.TrimLeft(p,"/")
	}

	return  strings.TrimLeft(p,"/")+"/"
}


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
	file = s3CleanFilePath(file)

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
	dir = s3CleanDirPath(dir)

	info,err := s.Stat(dir)
	if err != nil{
		return false,err
	}

	return info.IsDir,nil
}

func (s *S3Storage)ListFile(dir string)([]*FTPFileInfo,error){
	dir = s3CleanDirPath(dir)
	golog.Info("s3 list file:",dir)
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
		if dir == *v.Key{
			continue
		}
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
	oldname = s3CleanFilePath(oldname)
	newname = s3CleanFilePath(newname)

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
	dir = s3CleanDirPath(dir)
	_,err := s.cli.PutObject(s.Bucket,dir,nil)

	return err
}

func (s *S3Storage)DeleteFile(filename string)error{
	filename = s3CleanFilePath(filename)

	_,err := s.cli.DeleteObject(s.Bucket,filename)

	return err
}

func (s *S3Storage)DeleteDir(dir string)error{
	dir = s3CleanDirPath(dir)

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
	filename = s3CleanFilePath(filename)

	out,err := s.cli.DownloadObject(s.Bucket,filename)
	if err != nil{
		return 0,err
	}

	return io.Copy(wr,out.Body)
}

func (s *S3Storage)StoreFile(filename string, rd io.Reader)(int64,error){
	filename = s3CleanFilePath(filename)

	var cnt int = 0
	var cur_piece_size int64 = common_piece_size
	var uid string
	var wg sync.WaitGroup
	var ch chan *util.FBS3PieceInfo
	var sz int64
	var rbuf [4096]byte
	var ubuf[]byte = make([]byte,0,cur_piece_size)
	var bover bool = false

	for {
		if bover{
			break
		}

		for {
			n, err := rd.Read(rbuf[:])
			if err != nil &&err != io.EOF{
				return sz, err
			}else if err == io.EOF{
				bover = true
				break
			}

			sz += int64(n)

			ubuf = append(ubuf, rbuf[:n]...)
			if len(ubuf) >= int(cur_piece_size) {
				break
			}
		}

		if cnt == 0{
			if int64(len(ubuf)) < cur_piece_size{
				_,err := s.cli.PutObject(s.Bucket,filename,bytes.NewReader(ubuf[:]))
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

		if len(ubuf) > 0{
			wg.Add(1)

			go func(data []byte,part int){
				defer wg.Done()
				out,err := s.cli.UploadPartObject(s.Bucket,filename,part,uid,bytes.NewReader(data))
				if err != nil{
					return
				}

				ch <- &util.FBS3PieceInfo{int64(part),*out.ETag}
			}(ubuf[:],cnt+1)
		}

		cnt++
		cur_piece_size = cur_piece_size*2
		if cur_piece_size > max_piece_size{
			cur_piece_size = max_piece_size
		}

		ubuf = make([]byte,0,cur_piece_size)
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

	return sz,nil
}
