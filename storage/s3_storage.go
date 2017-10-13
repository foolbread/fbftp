/*
author: foolbread
file: storage/s3_storage.go
date: 2017/9/27
*/
package storage

import (
	"io"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type S3Storage struct {
	Bucket string
	AccKey string
	SecKey string
	EndPoint string
	Token string

	cli *s3.S3
}

func NewS3Storage(acckey string,seckey string,endpoint string,token string,bucket string)*S3Storage{
	r := new(S3Storage)
	r.AccKey = acckey
	r.SecKey = seckey
	r.EndPoint = endpoint
	r.Token = token
	r.Bucket = bucket

	cre := credentials.NewStaticCredentials(acckey,seckey,token)

	config := aws.NewConfig().WithRegion("us-east-1").
		WithEndpoint(endpoint).
		WithCredentials(cre).WithS3ForcePathStyle(true)

	sess,_ := session.NewSession(config)
	r.cli = s3.New(sess)

	return r
}

func (s *S3Storage)Stat(file string)(*FTPFileInfo,error){
	return nil,nil
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
