package command

import (
	"fmt"
	"strings"

	"github.com/DaveBlooman/awsr/configuration"
	"github.com/DaveBlooman/awsr/output"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func checkOptionFlags(flags map[string]string) {
	for k, v := range flags {
		if len(v) == 0 {
			msg := fmt.Sprintf("Error: '%s' is a required flag", k)
			output.Error(msg)
		}
	}
}

func setupAWS(flags map[string]string) *session.Session {

	var config aws.Config

	if flags["region"] == "" {
		flags["region"] = "eu-west-1"
	}

	file, err := configuration.Load()
	if err != nil {
		output.Error(err.Error())
	}

	if flags["env"] == "" {
		config = aws.Config{
			Region: aws.String(flags["region"]),
		}
	} else {
		env, err := file.GetSection(flags["env"])
		if err != nil {
			output.Error(err.Error())
		}

		access, err := env.GetKey("aws_access_key_id")
		if err != nil {
			output.Error("Unable to read .aws/credentials file")
		}
		secret, err := env.GetKey("aws_secret_access_key")
		if err != nil {
			output.Error("Unable to read .aws/credentials file")
		}
		config = aws.Config{
			Region:      aws.String(flags["region"]),
			Credentials: credentials.NewStaticCredentials(access.Value(), secret.Value(), ""),
		}
	}

	sess, err := session.NewSession(&config)
	if err != nil {
		output.Error("failed to create AWS session")
	}
	return sess
}

func fetchConfigGetFlags(c *cli.Context) map[string]string {
	return map[string]string{
		"env":    c.String("env"),
		"name":   c.String("name"),
		"region": c.String("region"),
		"limit":  c.String("limit"),
		"status": c.String("status"),
		"bucket": c.String("bucket"),
		"prefix": c.String("prefix"),
	}
}

func useName(flags map[string]string, instance string) bool {
	if flags["name"] != "" {
		if strings.Contains(strings.ToLower(instance), strings.ToLower(flags["name"])) {
			return true
		}
		return false
	}
	return true
}

func SetTitle(text string) string {
	return output.ChangeColor(text, color.FgBlue)
}
