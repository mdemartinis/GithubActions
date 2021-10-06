# These variables are used in the provider, in order to connect to AWS properly.
# ACCESS and SECRET keys will be declared but not populated.
# They will be fulfilled later with .tfvars file to keep the secrets. 

variable "AWS_ACCESS_KEY" {

}
variable "AWS_SECRET_KEY" {

}
variable "AWS_REGION" {
  default = "us-east-1"
}

# Next, AMI and size for EC2 instances
# This AMI corresponds to Ubuntu 20.04 LTS
variable "AMI_USEAST1" {
  default = "ami-036490d46656c4818"
}

variable "INSTANCE_TYPE" {
  default = "t2.micro"
}

# Bucket name. Random from pipeline

variable "BUCKET_NAME" {
  
}

# General Tags for all resources

variable "TAGS" {
  type    = map(string)
  default = {
    name  = "Flugel"
    owner = "InfraTeam"
  }
}