package sample

import (
	"fmt"
	"github.com/journeymidnight/Yig-S3-SDK-Go/s3lib"
	"io/ioutil"
	"os"
	"strings"
)

func AppendObjectSample() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()
	sc := s3lib.NewS3(endpoint, accessKey, secretKey)
	// Create a bucket
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}
	var nextPos int64

	// 1. Append strings to an object
	strs := []string{"yig1", "yig2", "yig3"}
	for _, s := range strs {
		fmt.Println("Append String:", s)
		nextPos, err = sc.AppendObject(bucketName, objectKey, strings.NewReader(s), nextPos)
		if err != nil {
			HandleError(err)
		}
	}
	out, err := sc.GetObject(bucketName, objectKey)
	b, _ := ioutil.ReadAll(out)
	fmt.Println("Get appended string:", string(b))
	out.Close()

	// Append files to an object
	strs = []string{"sample/L.jpeg", "sample/L1.jpeg", "sample/L2.jpeg"}
	for _, s := range strs {
		fmt.Println("Append file:", s)
		f, err := os.Open(s)
		defer f.Close()
		if err != nil {
			HandleError(err)
		}
		nextPos, err = sc.AppendObject(bucketName, objectKey, f, nextPos)
		if err != nil {
			HandleError(err)
		}

	}
	out, err = sc.GetObject(bucketName, objectKey)
	if err != nil {
		HandleError(err)
	}

	out.Close()

	//  Get Next Append Position
	nextPos, err = sc.GetObjectNextAppendPosition(bucketName, objectKey, nextPos)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("Next Pos is: ", nextPos)

	// Append With ACL And Meta
	strs = []string{"yig1", "yig2", "yig3"}
	c := make(map[string]string)
	c["a"] = "b"
	for _, s := range strs {
		fmt.Println("Append String:", s)
		nextPos, err = sc.AppendObjectWithAclAndMeta(bucketName, objectKey, strings.NewReader(s), nextPos, "public-read", c)
		if err != nil {
			HandleError(err)
		}
	}
	out, err = sc.GetObject(bucketName, objectKey)
	out.Close()

	fmt.Printf("AppendObjectSample Run Success !\n\n")
}
