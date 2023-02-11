IMAGE=${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/orc-api:latest

# Local
run:
	go run orc-api
seed:
	go run orc-api seed
test:
	go test

# Container (c)
build:
	docker build --tag ${IMAGE} .
login:
	aws ecr get-login-password | docker login --username AWS --password-stdin ${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com
push:
	docker push ${IMAGE}
run-c:
	docker run -p 8081:8081 --env-file ./.env.container ${IMAGE}
ssh:
	ssh -i "${KEY_PATH}" ec2-user@${EC2_HOST}