package aws_s3

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"image/png"
	domainImage "pixstall-artist/domain/image"
	"pixstall-artist/domain/image/model"
)

type awsS3ImageRepository struct {
	s3 *s3.S3
}

const (
	BucketName = "pixstall-store-dev"
)

func NewAWSS3ImageRepository(s3 *s3.S3) domainImage.Repo {
	return &awsS3ImageRepository{
		s3: s3,
	}
}

func (a awsS3ImageRepository) SaveImage(ctx context.Context, pathImage model.PathImage) error {
	// create buffer
	buff := new(bytes.Buffer)

	// encode image to buffer
	err := png.Encode(buff, pathImage.Image)
	if err != nil {
		fmt.Println("failed to create buffer", err)
	}

	// convert buffer to reader
	reader := bytes.NewReader(buff.Bytes())

	// use it in `PutObjectInput`
	_, err = a.s3.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(BucketName),
		Key:    aws.String(pathImage.Path + pathImage.Name),
		Body:   reader,
		ContentType: aws.String("image"),
		ACL: aws.String("public-read"),  //profile should be public accessible
	})

	if err != nil {
		return err
	}
	return nil
}

func (a awsS3ImageRepository) SaveImages(ctx context.Context, pathImages []model.PathImage) ([]string, error) {

	var resultPaths []string
	for _, pathImage := range pathImages {
		// create buffer
		buff := new(bytes.Buffer)

		// encode image to buffer
		err := png.Encode(buff, pathImage.Image)
		if err != nil {
			fmt.Println("failed to create buffer", err)
		}
		uploadPath := pathImage.Path + pathImage.Name
		// convert buffer to reader
		reader := bytes.NewReader(buff.Bytes())

		// use it in `PutObjectInput`
		_, err = a.s3.PutObjectWithContext(ctx, &s3.PutObjectInput{
			Bucket: aws.String(BucketName),
			Key:    aws.String(uploadPath),
			Body:   reader,
			ContentType: aws.String("image"),
			ACL: aws.String("public-read"),  //profile should be public accessible
		})
		if err == nil {
			resultPaths = append(resultPaths, uploadPath)
		}
	}
	return resultPaths, nil
}