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

type FBS3PieceInfo struct {
	Idx  int64
	Etag string
}

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

func (u *FBS3Client)PutObject(bucket string, key string, rd io.ReadSeeker)(*s3.PutObjectOutput,error){
	var in s3.PutObjectInput
	in.SetBucket(bucket)
	in.SetKey(key)
	in.SetBody(rd)

	return u.cli.PutObject(&in)
}

func (u *FBS3Client)DownloadObject(bucket string, key string)(*s3.GetObjectOutput,error){
	var in s3.GetObjectInput
	in.SetBucket(bucket)
	in.SetKey(key)

	return u.cli.GetObject(&in)
}

func (u *FBS3Client)WriteObject(bucket string, key string)error{
	return nil
}

func (u *FBS3Client) CreateMutilPartUpload(bucket string, key string)(*s3.CreateMultipartUploadOutput,error){
	var in s3.CreateMultipartUploadInput
	in.SetBucket(bucket)
	in.SetKey(key)

	return u.cli.CreateMultipartUpload(&in)
}

func (u *FBS3Client)UploadPartObject(bucket string,key string,part int,uid string,rs io.ReadSeeker)(*s3.UploadPartOutput,error){
	var in s3.UploadPartInput
	in.SetBucket(bucket)
	in.SetKey(key)
	in.SetBody(rs)
	in.SetPartNumber(int64(part))
	in.SetUploadId(uid)

	return u.cli.UploadPart(&in)
}

func (u *FBS3Client)CompleteMultiPartUpload(bucket string,key string,ps []*FBS3PieceInfo,uid string)(*s3.CompleteMultipartUploadOutput,error){
	var cps []*s3.CompletedPart
	for _,v := range ps{
		var p s3.CompletedPart
		p.SetPartNumber(v.Idx)
		p.SetETag(v.Etag)

		cps = append(cps, &p)
	}

	var cmu s3.CompletedMultipartUpload
	cmu.SetParts(cps)

	var in s3.CompleteMultipartUploadInput
	in.SetBucket(bucket)
	in.SetKey(key)
	in.SetUploadId(uid)
	in.SetMultipartUpload(&cmu)

	return u.cli.CompleteMultipartUpload(&in)
}

func (u *FBS3Client) ListObjects(bucket string, key string)(*s3.ListObjectsOutput,error){
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

func (u *FBS3Client) ListAllObjects(bucket string, key string)(*s3.ListObjectsOutput,error){
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