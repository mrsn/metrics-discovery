package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/rds"
)

type Result struct {
	Data interface{} `json:"data"`
}

func main() {
	var (
		discoveryType = flag.String("type", "", "type of discovery. Only ELB, RDS and CloudFront supported right now")
		awsRegion     = flag.String("aws.region", "eu-central-1", "AWS region")
		list          interface{}
		err           error
	)
	flag.Parse()
	awsSession := session.New(aws.NewConfig().WithRegion(*awsRegion))

	switch *discoveryType {
	case "ELB":
		list, err = getAllElasticLoadBalancers(awsSession)
		if err != nil {
			log.Fatalf("Could not descibe load balancers: %v", err)
		}
	case "RDS":
		list, err = getAllDBInstances(awsSession)
		if err != nil {
			log.Fatalf("Could not describe db instances: %v", err)
		}
	case "CloudFront":
		list, err = getAllCloudFrontDistributions(awsSession)
		if err != nil {
			log.Fatalf("Could not list distributions")
		}
	default:
		log.Fatalf("discovery type %s not supported", *discoveryType)
	}
	err = json.NewEncoder(os.Stdout).Encode(Result{Data: list})
	if err != nil {
		log.Fatal(err)
	}
}

func getAllDBInstances(awsSession *session.Session) ([]map[string]string, error) {
	resp, err := rds.New(awsSession).DescribeDBInstances(&rds.DescribeDBInstancesInput{})
	if err != nil {
		return nil, fmt.Errorf("getting RDS instances:%v", err)
	}

	rdsIdentifiers := make([]map[string]string, len(resp.DBInstances))
	for _, rds := range resp.DBInstances {
		rdsIdentifiers = append(rdsIdentifiers, map[string]string{
			"{#RDSIDENTIFIER}": *rds.DBInstanceIdentifier,
		})
	}
	return rdsIdentifiers, nil
}

func getAllCloudFrontDistributions(awsSession *session.Session) ([]map[string]string, error) {
	resp, err := cloudfront.New(awsSession).ListDistributions(&cloudfront.ListDistributionsInput{})
	if err != nil {
		return nil, fmt.Errorf("listing CloudFront distributions %v", err)
	}

	dists := make([]map[string]string, len(resp.DistributionList.Items))
	for _, dist := range resp.DistributionList.Items {
		dists = append(dists, map[string]string{
			"{#DISTID}": *dist.Id,
		})
	}
	return dists, nil
}

func getAllElasticLoadBalancers(awsSession *session.Session) ([]map[string]string, error) {
	resp, err := elb.New(awsSession).DescribeLoadBalancers(&elb.DescribeLoadBalancersInput{})
	if err != nil {
		return nil, fmt.Errorf("reading ELBs:%v", err)
	}

	elbs := make([]map[string]string, len(resp.LoadBalancerDescriptions))
	for _, elb := range resp.LoadBalancerDescriptions {
		elbs = append(elbs, map[string]string{
			"{#LOADBALANCERNAME}": *elb.LoadBalancerName,
		})
	}
	return elbs, nil
}
