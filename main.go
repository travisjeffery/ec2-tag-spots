package main

import (
	"fmt"
	"log"

	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/travisjeffery/ec2-tag-spots/tagger"
)

var (
	region = kingpin.Flag("region", "Name of AWS region to use").Default("us-east-1").String()
	requestId = kingpin.Flag("spot-fleet-request-id", "ID of spot fleet request to get instances").Required().String()
	tags = kingpin.Flag("tags", "Key/value tags to add to instances").Required().StringMap()
)

func main() {
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version("1.0.0").Author("Travis Jeffery")
	kingpin.CommandLine.Help = "Create tags for spot fleets."
	kingpin.Parse()

	t := tagger.New(region, requestId, tags)
	instances, err := t.GetInstances()
	if err != nil {
		log.Fatalf("error getting instances: %s", err.Error())
		return
	}
	err = t.CreateTags(instances)
	if err != nil {
		log.Fatalf("error creating tags: %s", err.Error())
		return
	}
	fmt.Println("success creating tags")
}