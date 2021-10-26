-inject-env:
# use .env file if it exist
ifneq ("$(wildcard .env)","")
	$(eval -include .env)
	$(eval export $(shell sed 's/=.*//' .env))
else
 	$(eval include -include .env.example)
	$(eval export $(shell sed 's/=.*//' .env.example))
endif

serve-rest: inject-env
	go run ./cmd/... -c=./config serve-rest