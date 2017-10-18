/*
author: foolbread
file: util/
date: 2017/10/14
*/
package util

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type FBS3Client struct {
	cli *s3.S3
}

func NewFBS3Client(acckey string,seckey string,endpoint string,token string)*FBS3Client{
	r := new(FBS3Client)

	cre := credentials.NewStaticCredentials(acckey,seckey,token)

	config := aws.NewConfig().WithRegion("us-east-1").
		WithEndpoint(endpoint).
		WithCredentials(cre).WithS3ForcePathStyle(true)

	sess,_ := session.NewSession(config)

	r.cli = s3.New(sess)

	return r
}

func (u *FBS3Client)HeadObject(bucket string, key string)(*s3.HeadObjectOutput,error){
	var in s3.HeadObjectInput
	in.SetBucket(bucket)
	in.SetKey(key)

	return u.cli.HeadObject(&in)
}

func (u *FBS3Client)ListFile(bucket string, key string)(*s3.ListObjectsOutput,error){
	var in s3.ListObjectsInput
	in.SetBucket(bucket)
	in.SetPrefix(key)
	in.SetDelimiter("/")

	return u.cli.ListObjects(&in)
}

func (u *FBS3Client)ListAllFile(bucket string, key string)(*s3.ListObjectsOutput,error){
	var in s3.ListObjectsInput
	in.SetBucket(bucket)
	in.SetPrefix(key)

	return u.cli.ListObjects(&in)
}