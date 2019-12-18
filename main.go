package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/sqs"
	"ms-test-sqs-go/worker"
	"os"
)

func main() {
	awsConfig := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), "AWS_SECRET_ACCESS_KEY", ""),
		Region:           aws.String("AWS_DEFAULT_REGION"),
		Endpoint:         aws.String("AWS_SQS_URL"),
		DisableSSL:       aws.Bool(true),
	}

	sqsClient := worker.CreateSqsClient(awsConfig)
	workerConfig := &worker.Config{
		QueueName:          "screenshot",
		MaxNumberOfMessage: 15,
		WaitTimeSecond:     5,
	}
	eventWorker := worker.New(sqsClient, workerConfig)
	ctx := context.Background()

	// start the worker
	eventWorker.Start(ctx, worker.HandlerFunc(func(msg *sqs.Message) error {
		fmt.Println(aws.StringValue(msg.Body))
		return nil
	}))
}
