resource "aws_s3_bucket" "test1bucket" {
  bucket = var.BUCKET_NAME
  acl    = "private"
  tags   = var.TAGS

  versioning {
    enabled = true
  }
}