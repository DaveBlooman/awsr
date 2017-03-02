package command

import (
	"fmt"

	"github.com/DaveBlooman/awsr/output"
	"github.com/apcera/termtables"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/urfave/cli"
)

func CmdEc2(c *cli.Context) error {

	flags := fetchConfigGetFlags(c)

	sess := setupAWS(flags)
	client := ec2.New(sess)

	params := &ec2.DescribeInstancesInput{}
	resp, err := client.DescribeInstances(params)
	if err != nil {
		output.Error(err.Error())
	}

	reservations := resp.Reservations

	table := termtables.CreateTable()
	table.AddHeaders(SetTitle("InstanceID"), SetTitle("IP Address"), SetTitle("State"), SetTitle("Launch Time"), SetTitle("Name"))

	for _, reservation := range reservations {
		for _, instance := range reservation.Instances {
			id := *instance.InstanceId
			state := *instance.State.Name
			launchtime := instance.LaunchTime.String()
			tempIP := instance.PrivateIpAddress
			tags := instance.Tags

			var ipAddress string
			if tempIP == nil {
				ipAddress = "n/a"
			} else {
				ipAddress = *tempIP
			}

			var name string
			for _, tag := range tags {
				if *tag.Key == "Name" {
					name = *tag.Value
				}
			}

			if flags["status"] != "" && *instance.State.Name != flags["status"] {
				continue
			}

			if name != "" {
				if useName(flags, name) {
					table.AddRow(id, ipAddress, state, launchtime, name)
				}
			} else {
				table.AddRow(id, ipAddress, state, launchtime, name)
			}
		}
	}
	fmt.Println(table.Render())

	return nil
}
