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
	"path"
	"net/url"
	"io"
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

func (u *FBS3Client)DeleteObject(bucket string, key string)(*s3.DeleteObjectOutput,error){
	var in s3.DeleteObjectInput
	in.SetBucket(bucket)
	in.SetKey(key)

	return u.cli.DeleteObject(&in)
}

func (u *FBS3Client)CopyObject(newbucket string, newkey string, oldbucket string, oldkey string)(*s3.CopyObjectOutput,error){
	var in s3.CopyObjectInput
	in.SetBucket(newbucket)
	in.SetKey(newkey)
	in.SetCopySource(path.Join("/",oldbucket,url.QueryEscape(oldkey)))

	return u.cli.CopyObject(&in)
}

func (u *FBS3Client)PutObject(bucket string, key string, wr io.ReadSeeker)(*s3.PutObjectOutput,error){
	var in s3.PutObjectInput
	in.SetBucket(bucket)
	in.SetKey(key)
	in.SetBody(wr)

	return u.cli.PutObject(&in)
}

func (u *FBS3Client)ListFile(bucket string, key string)(*s3.ListObjectsOutput,error){
	var ret *s3.ListObjectsOutput = new(s3.ListObjectsOutput)
	var nextkey string
	for {
		var in s3.ListObjectsInput
		in.SetBucket(bucket)
		in.SetPrefix(key)
		in.SetDelimiter("/")
		in.SetMarker(nextkey)

		out,err := u.cli.ListObjects(&in)
		if err != nil{
			return nil,err
		}

		for _,v := range out.CommonPrefixes{
			ret.CommonPrefixes = append(ret.CommonPrefixes,v)
		}

		for _,v := range out.Contents{
			ret.Contents = append(ret.Contents,v)
		}

		if !*out.IsTruncated{
			break
		}

		nextkey = *out.NextMarker
	}


	return ret,nil
}

func (u *FBS3Client)ListAllFile(bucket string, key string)(*s3.ListObjectsOutput,error){
	var ret *s3.ListObjectsOutput = new(s3.ListObjectsOutput)
	var nextkey string

	for{
		var in s3.ListObjectsInput
		in.SetBucket(bucket)
		in.SetPrefix(key)
		in.SetMarker(nextkey)

		out,err := u.cli.ListObjects(&in)
		if err != nil{
			return nil,err
		}

		for _,v := range out.CommonPrefixes{
			ret.CommonPrefixes = append(ret.CommonPrefixes,v)
		}

		for _,v := range out.Contents{
			ret.Contents = append(ret.Contents,v)
		}

		if !*out.IsTruncated{
			break
		}

		nextkey = *out.NextMarker
	}

	return ret,nil
}