package repository

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/rs/xid"
)

type ImageUploader interface {
	UploadProfileImage(context context.Context, r io.Reader) (string, error)
	GetProfileURL(name string) string
}

type imageUploader struct {
	s          *session.Session
	bucketName string
	projetID   string
}

// NewUserImageRepository will create an object that represent the problem.Repository interface
func NewUserImageRepository(s *session.Session, bucketName, projectID string) ImageUploader {
	return &imageUploader{s, bucketName, projectID}
}

func (u *imageUploader) UploadProfileImage(context context.Context, r io.Reader) (string, error) {
	guid := xid.New()

	_, _, err := upload(context, r, u.projetID, u.bucketName, guid.String(), true)

	if err != nil {
		return "", err
	}

	return guid.String(), nil
}

func (u *imageUploader) GetProfileURL(name string) string {
	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", u.bucketName, name)
}

func upload(ctx context.Context, r io.Reader, projectID, bucket, name string, public bool) (*storage.ObjectHandle, *storage.ObjectAttrs, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, nil, err
	}

	bh := client.Bucket(bucket)
	if _, err = bh.Attrs(ctx); err != nil {
		return nil, nil, err
	}

	obj := bh.Object(name)
	w := obj.NewWriter(ctx)
	if _, err := io.Copy(w, r); err != nil {
		return nil, nil, err
	}
	if err := w.Close(); err != nil {
		return nil, nil, err
	}

	if public {
		if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
			return nil, nil, err
		}
	}

	attrs, err := obj.Attrs(ctx)

	return obj, attrs, err
}
