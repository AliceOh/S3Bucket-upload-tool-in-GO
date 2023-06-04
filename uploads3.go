package main

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/mattetti/filebuffer"
)

var uploadenvFlags = flag.NewFlagSet(appName, flag.ContinueOnError)

var uploadenvArgs = `arguments:

  file        file name with path to push to s3 bucket
  s3uri       s3 uri to push zip to (s3://<bucketname>/<bucketpath>)
`

var uploadS3Help = fmt.Sprintf(
	"usage: %s [global options...] uploadse [command options...] <file-to-upload> <s3uri> \n\n%s\n\n%s",
	appName,
	uploadenvArgs,
	helpGlobalOptions,
)

func uploads3(argv []string) int {
	// Parse arguments.
	if err := uploadenvFlags.Parse(argv); err != nil {
		fmt.Println(err.Error())
		fmt.Println(helpRoot)
		return 1
	}

	// Print help string and return if required.
	if flagHelp != nil && *flagHelp {
		fmt.Println(uploadS3Help)
		return 0
	}

	// Create the context for the following requests. We allow 60 seconds for
	// pushing the environment zip.
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Get arguments.
	filePath := flags.Arg(1)
	if filePath == "" {
		fmt.Println("expected path argument")
		fmt.Println(uploadS3Help)
		return 1
	}

	s3URI := flags.Arg(2)
	if s3URI == "" {
		fmt.Println("expected s3URI argument")
		fmt.Println(uploadS3Help)
		return 1
	}

	// Check filePath argument.
	if fs, err := os.Stat(filePath); err != nil {
		fmt.Printf("error stating '%s': %v\n", filePath, err)
		return 1
	} else if fs.IsDir() {
		fmt.Printf("expected '%s' to be a file\n", filePath)
		return 1
	}

	// Check filePath is a zip file.
	// zf, err := zip.OpenReader(filePath)
	// if err != nil {
	// 	fmt.Printf("error opening '%s' as a zip file: %v\n", filePath, err)
	// 	return 1
	// }
	// zf.Close()

	// Check s3URI argument.
	uri, err := url.ParseRequestURI(s3URI)
	if err != nil {
		fmt.Printf("invalid URI '%s' for S3 content: %v", s3URI, err)
		return 1
	}
	if uri.Scheme != "s3" {
		fmt.Printf("invalid URI '%s' for S3 content: expected scheme to be s3", s3URI)
		return 1
	}
	if uri.Host == "" {
		fmt.Printf("invalid URI '%s' for S3 content: expected host to be bucket name", s3URI)
		return 1
	}
	if uri.Path == "" {
		fmt.Printf("invalid URI '%s' for S3 content: path is empty", s3URI)
		return 1
	}

	// Read and encrypt the to-be-uploaded file.
	inp, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("error reading '%s': %v\n", filePath, err)
		return 1
	}

	// Upload to S3.
	config, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		fmt.Printf("error loading AWS config: %v\n", err)
		return 1
	}

	client := s3.NewFromConfig(config)
	s3Req := &s3.PutObjectInput{
		Bucket: aws.String(uri.Host),
		Key:    aws.String(strings.TrimPrefix(uri.Path, "/")),
		Body:   filebuffer.New(inp),
	}

	_, err = client.PutObject(ctx, s3Req)
	if err != nil {
		fmt.Printf("failed to put '%s': %v", s3URI, err)
		return 1
	}

	return 0
}
