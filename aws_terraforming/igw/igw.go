package igw

import (
	"waze/terraform/terraform_utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var ignoreKey = map[string]bool{
	"id": true,
}

func createResources(igws *ec2.DescribeInternetGatewaysOutput) []terraform_utils.TerraformResource {
	resoures := []terraform_utils.TerraformResource{}
	for _, internetGateway := range igws.InternetGateways {
		resourceName := ""
		if len(internetGateway.Tags) > 0 {
			for _, tag := range internetGateway.Tags {
				if aws.StringValue(tag.Key) == "Name" {
					resourceName = aws.StringValue(tag.Value)
					break
				}
			}
		}
		resoures = append(resoures, terraform_utils.TerraformResource{
			ResourceType: "aws_internet_gateway",
			ResourceName: resourceName,
			Item:         nil,
			ID:           aws.StringValue(internetGateway.InternetGatewayId),
			Provider:     "aws",
		})
	}
	return resoures
}

func Generate(region string) error {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(region)})
	svc := ec2.New(sess)
	igws, err := svc.DescribeInternetGateways(&ec2.DescribeInternetGatewaysInput{})
	if err != nil {
		return err
	}
	resources := createResources(igws)
	err = terraform_utils.GenerateTfState(resources)
	if err != nil {
		return err
	}
	resources, err = terraform_utils.TfstateToTfConverter("terraform.tfstate", "aws", ignoreKey)
	if err != nil {
		return err
	}
	err = terraform_utils.GenerateTf(resources, "internet_gateway", region)
	if err != nil {
		return err
	}
	return nil

}