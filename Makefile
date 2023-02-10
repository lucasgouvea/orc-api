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
push:
	docker push ${IMAGE}
run-c:
	docker run -p 8081:8081 --env-file ./.env.container ${IMAGE}