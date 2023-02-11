#!/bin/bash
aws ecr get-login-password | docker login --username AWS --password-stdin ${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com
docker pull ${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/orc-api:latest
docker stop $(docker ps -a -q --filter="name=orc-api")
docker rm $(docker ps -a -q --filter="name=orc-api")
docker run -d -p 8081:8081 --name orc-api --env-file ./orc-api/.env.prod ${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/orc-api:latest

