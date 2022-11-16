package main

import (
	"fmt"
	"os"

	// "github.com/qor/oss/s3"
	// "github.com/oss/filesystem"
	awss3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/padiazg/qor-oss-test/s3"
)

var (
	fileName = "test.json"

	cfg1 = &s3.Config{
		AccessID:  "",
		AccessKey: "",
		Region:    "sa-east-1",
		Bucket:    "portal-test-pato",
		Endpoint:  "https://s3.sa-east-1.amazonaws.com",
		ACL:       awss3.BucketCannedACLPrivate,
		// ACL:              awss3.BucketCannedACLPublicRead,
		S3ForcePathStyle: true,
	}
)

func main() {
	storage := s3.New(cfg1)

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("failed to open file %q, %+v\n", fileName, err)
	}

	// Save a reader interface into storage
	if _, e := storage.Put(fileName, f); e != nil {
		fmt.Printf("storage.Put: %+v\n", e)
	}

	// Get file with path
	// if _, e := storage.Get(fileName); e != nil {
	// 	fmt.Printf("storage.Get %s\n", e.Error())
	// }

	// // Get object as io.ReadCloser
	// if o, e := storage.GetStream(fileName); e != nil {
	// 	fmt.Printf("storage.GetStream %s\n", e.Error())
	// } else {
	// 	localFile, err := os.Create("test-local.json")
	// 	if err != nil {
	// 		fmt.Printf("storage.GetStream. Creating local file %s\n", e.Error())
	// 		return
	// 	}

	// 	if _, err = io.Copy(localFile, o); err != nil {
	// 		fmt.Printf("storage.GetStream. Saving to local file %s\n", e.Error())
	// 		return
	// 	}
	// }

	// // Delete file with path
	// storage.Delete(fileName)

	// List all objects under path
	// if l, e := storage.List(""); e != nil {
	// 	fmt.Printf("storage.List %s\n", e.Error())
	// } else {
	// 	for _, o := range l {
	// 		fmt.Printf("> %s %s\n", o.Name, o.Path)
	// 	}
	// }

	// Get Public Accessible URL (useful if current file saved privately)
	// if u, e := storage.GetURL(fileName); e != nil {
	// 	fmt.Printf("storage.GetURL %s\n", e.Error())
	// } else {
	// 	fmt.Printf("%s: %s\n", fileName, u)
	// }
}
