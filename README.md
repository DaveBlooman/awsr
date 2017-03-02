# AWSR

A Developer AWS Read only CLI tool.  

### Why?

The AWS CLI is powerful, but doesn't offer the nicest interface for a simple workflow.  AWSR also focusses only on reading from the AWS API, instances, buckets, IAM etc.  AWSR supports using environment variables, but also from your `.aws/credentials` file.  This allows for fast switching between predefined profiles.

## Usage

```sh
NAME:
   awsr - Developer driven command line tool for AWS, but only for read.

USAGE:
   awsr [global options] command [command options] [arguments...]

VERSION:
   0.1.0

AUTHOR:
   DaveBlooman

COMMANDS:
     ec2
     iam
     s3
     vpcs
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```


#### Name
There are generic switches which have been implemented on subcommands, for example `-n` for a regex on name.

```sh
awsr ec2 -n consul

+---------------------+--------------+---------+-------------------------------+--------+
| InstanceID          | IP Address   | State   | Launch Time                   | Name   |
+---------------------+--------------+---------+-------------------------------+--------+
| i-00000000000000001 | 10.10.12.139 | running | 2016-11-14 18:34:00 +0000 UTC | Consul |
| i-00000000000000002 | 10.10.11.39  | running | 2016-11-26 19:31:26 +0000 UTC | Consul |
| i-00000000000000003 | 10.10.10.130 | running | 2016-12-26 04:50:28 +0000 UTC | Consul |
| i-00000000000000004 | 10.10.10.147 | running | 2016-11-14 18:31:57 +0000 UTC | Consul |
| i-00000000000000005 | 10.10.12.214 | stopped | 2016-11-14 18:36:05 +0000 UTC | Consul |
+---------------------+--------------+---------+-------------------------------+--------+
```

#### Status
In our next scenario, we only want running instances.

```sh
awsr ec2 -n consul -s running

+---------------------+--------------+---------+-------------------------------+--------+
| InstanceID          | IP Address   | State   | Launch Time                   | Name   |
+---------------------+--------------+---------+-------------------------------+--------+
| i-00000000000000001 | 10.10.12.139 | running | 2016-11-14 18:34:00 +0000 UTC | Consul |
| i-00000000000000002 | 10.10.11.39  | running | 2016-11-26 19:31:26 +0000 UTC | Consul |
| i-00000000000000003 | 10.10.10.130 | running | 2016-12-26 04:50:28 +0000 UTC | Consul |
| i-00000000000000004 | 10.10.10.147 | running | 2016-11-14 18:31:57 +0000 UTC | Consul |
+---------------------+--------------+---------+-------------------------------+--------+
```

#### Limit

Responses for some services are limited as they are likely to include large responses.  When combined with a name match, you may see an empty response as the limit wasn't enough to include your query.  To increase the limit, use the `-l` flag

```sh
awsr iam -n consul -l 40

+---------------------------------------------------+-------------------------------+--------------------+-----------------------+
| ARN                                               | Create Date                   | Role Name          | Role ID               |
+---------------------------------------------------+-------------------------------+--------------------+-----------------------+
| arn:aws:iam::000000000000:role/consul-server-role | 2015-02-27 17:30:16 +0000 UTC | consul-server-role | AR000000000000000000K |
+---------------------------------------------------+-------------------------------+--------------------+-----------------------+
```

#### Environment

Environment is the term for different AWS credentials.  You may have different AWS accounts within you credentials file, so each of these credentials is treated as a separate AWS environment.

```sh
awsr iam -n consul -e dev

+---------------------------------------------------+-------------------------------+--------------------+-----------------------+
| ARN                                               | Create Date                   | Role Name          | Role ID               |
+---------------------------------------------------+-------------------------------+--------------------+-----------------------+
| arn:aws:iam::000000000001:role/consul-server-role | 2015-02-25 12:33:11 +0000 UTC | consul-server-role | AR000000000000000001K |
+---------------------------------------------------+-------------------------------+--------------------+-----------------------+
```


## Install

To install, use `go get`:

```bash
$ go get -d github.com/DaveBlooman/awsr
```

## Contribution

1. Fork ([https://github.com/DaveBlooman/awsr/fork](https://github.com/DaveBlooman/awsr/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[DaveBlooman](https://github.com/DaveBlooman)
