package s3

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/globalsign/mgo/bson"
)

func Upload(userID string, s *session.Session, file multipart.File, fileHeader *multipart.FileHeader, wg *sync.WaitGroup) error {
	defer wg.Done()

	size := fileHeader.Size
	buffer := make([]byte, size)
	_, _ = file.Read(buffer)

	fileName := fmt.Sprintf("%v/%v-%v", userID, bson.NewObjectId().Hex(), fileHeader.Filename)
	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(os.Getenv("AWS_S3_UPLOAD_BUCKET")),
		Key:                  aws.String(fileName),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})

	if err != nil {
		return err
	}

	return nil
}
