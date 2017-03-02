package command

import (
	"fmt"

	"github.com/apcera/termtables"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/urfave/cli"
)

func CmdS3Buckets(c *cli.Context) error {

	flags := fetchConfigGetFlags(c)

	sess := setupAWS(flags)
	client := s3.New(sess)

	var params *s3.ListBucketsInput
	resp, err := client.ListBuckets(params)

	if err != nil {
		return err
	}

	table := termtables.CreateTable()
	table.AddHeaders(SetTitle("Name"), SetTitle("Created On"))

	for _, bucket := range resp.Buckets {
		name := *bucket.Name
		created := bucket.CreationDate.String()

		table.AddRow(name, created)
	}

	fmt.Println(table.Render())

	return nil

}

func CmdS3Objects(c *cli.Context) error {
	flags := fetchConfigGetFlags(c)

	sess := setupAWS(flags)
	client := s3.New(sess)

	params := &s3.ListObjectsInput{
		Bucket: aws.String(flags["bucket"]),
		Prefix: aws.String(flags["prefix"]),
	}
	resp, err := client.ListObjects(params)
	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}
