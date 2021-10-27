package config

import "os"

func GetS3BucketName() string {
	return os.Getenv("S3_BUCKET")
}
