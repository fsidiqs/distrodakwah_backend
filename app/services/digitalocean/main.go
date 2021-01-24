package digitalocean

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	productClass "distrodakwah_backend/app/services/library/productlibrary"
)

var (
	KEY      string = "AHF4Y7QPRPHBRNLST7IM"
	SECRET   string = "k/oJdajH1+FwvDOyuReWwOtz6bZr9ngEv5e1ttXPGMY"
	ENDPOINT string = "sgp1.digitaloceanspaces.com"
	REGION   string = "us-east-1"
	BUCKET   string = "files-distrodakwah"
)

func UploadFiles(files []productClass.ProductImage) ([]string, error) {

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
			Key:         aws.String(fmt.Sprintf("dd_product_images/%v", file.FileName)),
			Body:        file.Content,
			ContentType: aws.String("image/jpeg"),
		})
		if err != nil {

			return []string{}, err
		}

		// results[i] = aws.StringValue(&result.Location)
		_ = result
		results[i] = "https://files-distrodakwah.sgp1.digitaloceanspaces.com/testing/resumes/fajar_sidiq_salviro1-1608879556.pdf"

	}
	return results, nil
}
