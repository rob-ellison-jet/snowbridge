# Extended configuration of SQS as a failure target (all options)

failure_target {
  use "sqs" {
    # SQS queue name
    queue_name = "mySqsQueue"

    # AWS region of SQS queue
    region     = "us-west-1"

    # Role ARN to use on SQS queue
    role_arn   = "arn:aws:iam::123456789012:role/myrole"
  }
}
