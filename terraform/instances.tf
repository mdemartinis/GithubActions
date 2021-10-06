resource "aws_instance" "Test1Instance" {
  ami           = var.AMI_USEAST1
  instance_type = var.INSTANCE_TYPE
  tags          = var.TAGS
}