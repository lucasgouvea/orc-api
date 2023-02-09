run:
	go run orc-api
test:
	go test
# Migrations
mig-up:
	go run orc-api migrations:up
mig-down:
	go run orc-api migrations:down