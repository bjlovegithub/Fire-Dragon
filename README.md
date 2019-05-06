# Fire-Dragon
Backend for Best Wishes iOS App.

## Revel HTTP Framework for Go

## SSL to Protect Data
Using an self-signed key in Revel.

# Development
## Docker

### Run locally
```docker build --build-arg mysql_user="" --build-arg mysql_pass="" --build-arg mysql_host_port="" --build-arg mysql_db="" --build-arg jwt_sec_key="" -t fire-dragon .```

```docker run -p 127.0.0.1:443:443 b3b66e45a5bf```

# Deployment
## Prepare MySQL DB
* Create an RDS DB.
* Create the theree tables from model folder.

## Prepare docker image
* Build the docker image on EC2 instance.
* Push the docker image to ECR.

## Create an ALB
* Create an self-signed SSL key for AWS LB: http://www.lalitgolani.com/2019/04/deploying-self-signed-ssl-certificate.html

## Start ECS Service
* Create an ECS cluster
* Build a new service for the docker image created above.
* Attach the ALB to the ECS service.
* Properly configure the security groups for ALB and ECS service.

