package digitalocean

import (
	"fmt"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	KEY      string = "AHF4Y7QPRPHBRNLST7IM"
	SECRET   string = "k/oJdajH1+FwvDOyuReWwOtz6bZr9ngEv5e1ttXPGMY"
	ENDPOINT string = "sgp1.digitaloceanspaces.com"
	REGION   string = "us-east-1"
	BUCKET   string = "files-distrodakwah"
)

func UploadFiles(files []multipart.File) ([]string, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(KEY, SECRET, ""),
		Endpoint:    &ENDPOINT,
		Region:      &REGION,
	}))
	uploader := s3manager.NewUploader(sess)
	results := make([]string, len(files))
	for i, file := range files {
		result, err := uploader.Upload(&s3manager.UploadInput{
			ACL:         aws.String("public-read"),
			Bucket:      aws.String(BUCKET),
			Key:         aws.String(fmt.Sprintf("dd_product_images/test%v", time.Now().Unix())),
			Body:        file,
			ContentType: aws.String("image/jpeg"),
		})
		if err != nil {
			return []string{}, err
		}
		results[i] = aws.StringValue(&result.Location)
		fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))

	}
	return results, nil
}
