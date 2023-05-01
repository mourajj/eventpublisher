#!/bin/bash
# Install Docker and start the service
sudo yum update -y
sudo amazon-linux-extras install -y docker
sudo service docker start
sudo yum install -y git
cd
git clone https://github.com/mourajj/eventpublisher.git
cd eventpublisher
sudo make build
sudo make run