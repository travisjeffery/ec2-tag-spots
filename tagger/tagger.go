package tagger

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Client struct {
	requestId *string
	tags      *map[string]string
	ec2       *ec2.EC2
}

func New(region, requestId *string, tags *map[string]string) *Client {
	sess := session.New(&aws.Config{Region: region})
	return &Client{
		requestId: requestId,
		tags:      tags,
		ec2:       ec2.New(sess),
	}
}

func (c *Client) CreateTags(instances []*string) error {
	var tags []*ec2.Tag
	for k, v := range *c.tags {
		tags = append(tags, &ec2.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		})
	}
	input := &ec2.CreateTagsInput{
		Resources: instances,
		Tags:      tags,
	}
	_, err := c.ec2.CreateTags(input)
	return err
}

func (c *Client) GetInstances() ([]*string, error) {
	input := &ec2.DescribeSpotFleetInstancesInput{
		SpotFleetRequestId: c.requestId,
	}
	output, err := c.ec2.DescribeSpotFleetInstances(input)
	if err != nil {
		return nil, err
	}
	var instances []*string
	for _, i := range output.ActiveInstances {
		instances = append(instances, i.InstanceId)
	}
	return instances, nil
}
