package main

import (
	"fmt"
	"os"

	"github.com/DaveBlooman/awsr/command"
	"github.com/urfave/cli"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "ec2",
		Usage:  "",
		Action: cli.ActionFunc(command.CmdEc2),
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "env, e",
				Usage: "Environment, e.g: dev",
			},
			cli.StringFlag{
				Name:  "name, n",
				Usage: "Name via instance tag, e.g: Ruby",
			},
			cli.StringFlag{
				Name:  "region, r",
				Usage: "AWS region, e.g: eu-west-1",
			},
			cli.StringFlag{
				Name:  "status, s",
				Usage: "Current server status, e.g: running",
			},
		},
	},
	{
		Name:   "iam",
		Usage:  "",
		Action: cli.ActionFunc(command.CmdIam),
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "env, e",
				Usage: "Environment, e.g: dev",
			},
			cli.StringFlag{
				Name:  "name, n",
				Usage: "Name via instance tag, e.g: Ruby",
			},
			cli.StringFlag{
				Name:  "region, r",
				Usage: "AWS region, e.g: eu-west-1",
			},
			cli.StringFlag{
				Name:  "limit, l",
				Usage: "Amount of IAM roles, e.g: 50",
			},
		},
	},
	{
		Name:  "s3",
		Usage: "s3 buckets",
		Subcommands: []cli.Command{
			{
				Name:   "buckets",
				Usage:  "list buckets",
				Action: cli.ActionFunc(command.CmdS3Buckets),
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "env, e",
						Usage: "Environment, e.g: dev",
					},
				},
			},
			{
				Name:   "objects",
				Usage:  "remove an existing template",
				Action: cli.ActionFunc(command.CmdS3Objects),
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "bucket, b",
						Usage: "Bucket, e.g: bucketname",
					},
					cli.StringFlag{
						Name:  "prefix, p",
						Usage: "Prefix of objects, e.g: img_",
					},
				},
			},
		},
	},
	{
		Name:   "vpcs",
		Usage:  "",
		Action: cli.ActionFunc(command.CmdVpcs),
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
