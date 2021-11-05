SHELL := /bin/zsh

# ==============================================================================
# Testing running system
#
# To generate a private/public key PEM file.
# openssl genpkey -algorithm RSA -out private.pem -pkeyopt rsa_keygen_bits:2048
# openssl rsa -pubout -in private.pem -out public.pem
#
#
#
# curl --user "admin@example.com:gophers" http://localhost:3000/v1/users/token/54bb2165-71e1-41a6-af3e-7da4a0e1e2c1
# export TOKEN="COPY TOKEN STRING FROM LAST CALL"
# curl -H "Authorization: Bearer ${TOKEN}" http://localhost:3000/v1/users/1/2
#
# expvarmon -ports=":4000" -vars="build,requests,goroutines,errors,mem:memstats.Alloc"
# hey -m GET -c 100 -n 1000000 "http://localhost:3000/readiness"



# ==============================================================================
# Running tests within the local computer

test:
	go test -v ./... -count=1
#	staticcheck ./...

run:
	go run app/sales-api/main.go

runa:
	go run app/admin/main.go
tidy:
	go mod tidy
	go mod vendor