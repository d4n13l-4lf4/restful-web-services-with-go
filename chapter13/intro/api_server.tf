provider "aws" {
  profile = "default"
  region = "some-region"
}

resource "aws_instance" "api_server" {
  ami = "ami-03818140b4ac9ae2b"
  instance_type = "t2.micro"
  key_name = aws_key_pair.api_server_key.key_name
}

resource "aws_key_pair" "api_server_key" {
  key_name = "api-server-key"
  public_key = "some key"
}