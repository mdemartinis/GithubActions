resource "aws_s3_bucket" "Test1Bucket" {
  bucket = var.BUCKET_NAME
  acl    = "private"
  tags   = var.TAGS
}