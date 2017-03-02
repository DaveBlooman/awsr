package command

import (
	"fmt"
	"strconv"

	"github.com/DaveBlooman/awsr/output"
	"github.com/apcera/termtables"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/urfave/cli"
)

func CmdIam(c *cli.Context) error {

	flags := fetchConfigGetFlags(c)

	sess := setupAWS(flags)
	svc := iam.New(sess)

	var iamLimit int64

	if flags["limit"] == "" {
		iamLimit = 10
	} else {
		convert, err := strconv.ParseInt(flags["limit"], 10, 64)
		if err != nil {
			output.Error(err.Error())
		}
		iamLimit = convert
	}

	params := &iam.ListRolesInput{
		MaxItems: aws.Int64(iamLimit),
	}
	resp, err := svc.ListRoles(params)

	if err != nil {
		fmt.Println(err.Error())
	}

	table := termtables.CreateTable()
	table.AddHeaders(SetTitle("ARN"), SetTitle("Create Date"), SetTitle("Role Name"), SetTitle("Role ID"))

	for _, roles := range resp.Roles {
		if useName(flags, *roles.RoleName) {
			table.AddRow(*roles.Arn, *roles.CreateDate, *roles.RoleName, *roles.RoleId)
		}
	}

	fmt.Println(table.Render())

	return nil

}
