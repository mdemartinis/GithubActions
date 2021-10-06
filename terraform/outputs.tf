output "instance-tags" {
  value = aws_instance.Test1Instance.tags
}

output "bucket-tags" {
  value = aws_s3_bucket.Test1Bucket.tags
}