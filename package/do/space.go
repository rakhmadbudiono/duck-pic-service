package do

import (
	"log"
	"net/http"

	"github.com/rakhmadbudiono/duck-pic-service/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var s3Config *aws.Config
var cfg *config.Config

func init() {
	cfg = config.New()

	s3Config = &aws.Config{
		Credentials: credentials.NewStaticCredentials(cfg.DO.SpaceKey, cfg.DO.SpaceSecret, ""),
		Endpoint:    aws.String(cfg.DO.SpaceHost),
		Region:      aws.String("us-east-1"),
	}
}

// GetObject fetches object from remote storage
func GetObject(path string) ([]byte, string) {
	sess, err := session.NewSession(s3Config)
	if err != nil {
		log.Printf("Unable to create new session %v", err)
	}

	downloader := s3manager.NewDownloader(sess)

	buff := &aws.WriteAtBuffer{}
	_, err = downloader.Download(buff,
		&s3.GetObjectInput{
			Bucket: aws.String(cfg.DO.SpaceBucket),
			Key:    aws.String(path),
		},
	)
	if err != nil {
		log.Fatalf("Unable to download item %q, %v", path, err)
	}

	bytes := buff.Bytes()

	return buff.Bytes(), http.DetectContentType(bytes)
}
