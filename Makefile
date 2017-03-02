ec2:
	go install && cli ec2 -n consul

iam:
	go install && cli iam -l 50 -n bot

install:
	glide install & go install
