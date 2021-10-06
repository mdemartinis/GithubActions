resource "aws_instance" "test1instance" {
  ami           = var.AMI_USEAST1
  instance_type = var.INSTANCE_TYPE
  tags          = var.TAGS

  metadata_options {
    http_tokens = "required"
  }
}