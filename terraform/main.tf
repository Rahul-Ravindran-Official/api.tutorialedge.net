provider "aws" {
    region = "eu-west-1"
}

resource "aws_vpc" "private_lambda_exec" {
    cidr_block = "10.0.0.0/16"

    tags = {
        Name = "private_lambda_exec"
    }
}

resource "aws_subnet" "main" {
  vpc_id     = "${aws_vpc.private_lambda_exec.id}"
  cidr_block = "10.0.1.0/24"

  tags = {
    Name = "private_lambda_exec"
  }
}

resource "aws_security_group" "allow_tls" {
    name = "private_lambda_exec"
    description = "A locked down ASG for lambda execution"
    vpc_id = "${aws_vpc.private_lambda_exec.id}"

    tags = {
        Name = "No-Access"
    }
}