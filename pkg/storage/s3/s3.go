package s3

// import (
// 	"mime/multipart"
//
// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/credentials"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/s3"
// 	"pstpmn.com/internal/core/port"
// )
//
// type storageS3 struct {
// 	Key      string
// 	Secret   string
// 	Endpoint string
// 	Region   string
// 	Bucket   string
// }
//
// func (h *storageS3) createAWSConfig() *aws.Config {
// 	return &aws.Config{
// 		Credentials:      credentials.NewStaticCredentials(h.Key, h.Secret, ""),
// 		Endpoint:         aws.String(h.Endpoint),
// 		S3ForcePathStyle: aws.Bool(false),
// 		Region:           aws.String(h.Region),
// 	}
// }
//
// func (h *storageS3) createObjectInput(path string, f *multipart.FileHeader) *s3.PutObjectInput {
// 	objectKey := path
// 	objectBody, _ := f.Open()
// 	return &s3.PutObjectInput{
// 		Bucket:      aws.String(h.Bucket),
// 		Key:         aws.String(objectKey),
// 		Body:        objectBody,
// 		ACL:         aws.String("public-read"),
// 		ContentType: aws.String("image/jpeg"),
// 	}
// }
//
// func (h *storageS3) UploadImage(image *multipart.FileHeader, path string) error {
// 	awsConfig := h.createAWSConfig()
// 	newSession := session.New(awsConfig)
// 	s3Client := s3.New(newSession)
// 	objectInput := h.createObjectInput(path, image)
// 	_, err := s3Client.PutObject(objectInput)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
//
// func NewStorageS3(bucket, region, endpoint, key, secret string) port.IStorageService {
// 	return &storageS3{
// 		Key:      key,
// 		Secret:   secret,
// 		Endpoint: endpoint,
// 		Region:   region,
// 		Bucket:   bucket,
// 	}
// }
