-include .env
-include .env.example

inject-env:
# use .env file if it exist
ifneq ("$(wildcard .env)","")
	$(eval export $(shell sed 's/=.*//' .env))
else
	echo ".env file not found, it will run on .env.example instead."
	$(eval export $(shell sed 's/=.*//' .env.example))
endif

serve-rest: inject-env
	go run ./cmd/... -c=./config serve-rest

