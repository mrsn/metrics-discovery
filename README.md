# metrics-discovery
Can be used in a monitoring systems like nagios or zabbix to discover items on aws 

#### Installing
	go get github.com/weltn24/metrics-discovery

#### Discover ELBs
	
	metrics-discovery -aws.region eu-central-1 -type ELB

#### Discover ALBs
	
	metrics-discovery -aws.region eu-central-1 -type ALB

#### Discover EC2 instances

	metrics-discovery -aws.region eu-central-1 -type EC2

#### Discover RDS instances

	metrics-discovery -aws.region eu-central-1 -type RDS

#### Discover Queues

	metrics-discovery -aws.region eu-central-1 -type SQS

#### Discover CloudFront distributions

	metrics-discovery -aws.region us-east-1 -type CloudFront

#### Discover Lambda functions

	metrics-discovery -aws.region eu-central-1 -type Lambda

#### Discover ECS clusters

    metrics-discovery -aws.region eu-central-1 -type ECSClusters

#### Discover Services running on ECS
	
	metrics-discovery -aws.region eu-central-1 -type ECSServices

#### Discover Checks in Trusted Advisor
	
	metrics-discovery -aws.region us-east-1 -type AWS/TrustedAdvisor
