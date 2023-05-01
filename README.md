# GO Event Publisher

This is an application to send events (files) to S3, this project is built on top of a cloudformation stack, which creates the resources needed and runs the application on a docker container inside of an ec2 instance.

![arquitetura](https://user-images.githubusercontent.com/27701706/235545271-8b11ecb2-b40e-4583-805b-b61ffb90adac.png)

To build and deploy this application, you need to have an AWS account and deploy the cloudformation stack (cloudformation.yaml).
It will create the following resources:

* S3InstanceProfile - Creates a Profile needed for EC2 instance
* S3Policy - Creates a policy to allow access to S3 resources
* S3InstanceProfileRole - Creates a Role to attach the policies to
* S3Bucket - Creates an S3 bucket
* EC2Instance - Creates an EC2 instance
* InstanceSecurityGroup - Creates an security group for the EC2 instance

After deploying the cloudformation stack, the bash commands defined on the cloudformation file will automatically build and run the application, and the S3-bucket selected will start to be populated with the files from the event folder.

