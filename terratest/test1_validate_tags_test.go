package test

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
}
func TestTerraformTags(t *testing.T) {
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "./",
		Vars: map[string]interface{}{
			"BUCKET_NAME": "test1bucketmdemartinis",
		},
	})

	// Clean up resources with "terraform destroy" at the end of the test.
	defer terraform.Destroy(t, terraformOptions)

	// Run "terraform init" and "terraform apply". Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// connect to AWS and retrieve instance and bucket

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		exitErrorf("Unable to open session to AWS, %v", err)
	}

	// Create S3 service client
	svc := s3.New(sess)

	// Read S3 buckets
	bucketsResult, err := svc.ListBuckets(nil)
	if err != nil {
		exitErrorf("Unable to list buckets, %v", err)
	}

	fmt.Println("Buckets:")

	for _, b := range bucketsResult.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}

	getInput := &s3.GetBucketTaggingInput{
		Bucket: aws.String("test1bucketmdemartinis"),
	}

	bucketTags, err := svc.GetBucketTagging(getInput)
	if err != nil {
		exitErrorf("Unable to get Bucket Tags, %v", err)
	} else {
		for _, b := range bucketTags.TagSet {
			fmt.Printf(" * Tag Name: %s - Tag Value: %s\n",
				aws.StringValue(b.Key), aws.StringValue(b.Value))
		}
	}

	// Create EC2 service client
	svcec2 := ec2.New(sess)

	// Read EC2 instances

	input := &ec2.DescribeTagsInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("resource-type"),
				Values: []*string{
					aws.String("instance"),
				},
			},
		},
	}

	ec2result, err := svcec2.DescribeTags(input)
	if err != nil {
		exitErrorf("Unable to get EC2 Instance Tags, %v", err)
	}

	fmt.Println(ec2result)

	// Run `terraform output` to get the values of output variables and check they have the expected values.
	output_instance := terraform.Output(t, terraformOptions, "instance-tags")
	output_bucket := terraform.Output(t, terraformOptions, "bucket-tags")

	expected_tags := "map[name:Flugel owner:InfraTeam]"

	assert.Equal(t, expected_tags, output_instance)
	assert.Equal(t, expected_tags, output_bucket)

}
