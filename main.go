package main

import (
	"fmt"
	"io"
	"os"

	awss3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/qor/oss/s3"
	// "github.com/oss/filesystem"
	// "github.com/padiazg/qor-oss-test/s3"
)

var (
	fileName  = "test.json"
	accessID  = "81QKqhvmrOEEuuUs"
	accessKey = "TbUMQkkVRLUZayAkMVMIyPWc6Z4ru08o"
	region    = "us-east-1"
	bucket    = "bucket"
	endPoint  = "http://localhost:9000"
)

func main() {

	storage := s3.New(&s3.Config{
		AccessID:         accessID,
		AccessKey:        accessKey,
		Region:           region,
		Bucket:           bucket,
		S3Endpoint:       endPoint,
		ACL:              awss3.BucketCannedACLPublicRead,
		S3ForcePathStyle: true,
	})
	// storage := filesystem.New("/tmp")

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("failed to open file %q, %v\n", fileName, err)
	}

	// Save a reader interface into storage
	if _, e := storage.Put(fileName, f); e != nil {
		fmt.Printf("storage.Put %s\n", e.Error())
	}

	// Get file with path
	// if _, e := storage.Get(fileName); e != nil {
	// 	fmt.Printf("storage.Get %s\n", e.Error())
	// }

	// Get object as io.ReadCloser
	if o, e := storage.GetStream(fileName); e != nil {
		fmt.Printf("storage.GetStream %s\n", e.Error())
	} else {
		localFile, err := os.Create("test-local.json")
		if err != nil {
			fmt.Printf("storage.GetStream. Creating local file %s\n", e.Error())
			return
		}

		if _, err = io.Copy(localFile, o); err != nil {
			fmt.Printf("storage.GetStream. Saving to local file %s\n", e.Error())
			return
		}
	}

	// Delete file with path
	// storage.Delete(fileName)

	// List all objects under path
	if l, e := storage.List("/"); e != nil {
		fmt.Printf("storage.List %s\n", e.Error())
	} else {
		for _, o := range l {
			fmt.Printf("> %s %s\n", o.Name, o.Path)
		}
	}

	// // Get Public Accessible URL (useful if current file saved privately)
	if u, e := storage.GetURL(fileName); e != nil {
		fmt.Printf("storage.GetURL %s\n", e.Error())
	} else {
		fmt.Printf("%s: %s\n", fileName, u)
	}
}
