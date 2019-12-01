package repository

import (
	"bytes"
	"context"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/rs/xid"
)

type UserS3Image interface {
	PutImageFile(context context.Context, data []byte) (string, error)
}

type userS3Image struct {
	s          *session.Session
	bucketName string
	fileDir    string
}

// NewUserImageRepository will create an object that represent the problem.Repository interface
func NewUserImageRepository(s *session.Session, bucketName, fileDir string) UserS3Image {
	return &userS3Image{s, bucketName, fileDir}
}

func (u *userS3Image) PutImageFile(context context.Context, data []byte) (string, error) {
	guid := xid.New()

	tempFileName := "profiles/" + guid.String()

	_, err := s3.New(u.s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(u.bucketName),
		Key:                  aws.String(tempFileName),
		ACL:                  aws.String("public-read"), // could be private if you want it to be access by only authorized users
		Body:                 bytes.NewReader(data),
		ContentType:          aws.String(http.DetectContentType(data)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})
	if err != nil {
		return "", err
	}

	return tempFileName, err
}
