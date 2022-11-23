# note: call scripts from /scripts

test-cover:
	go test .\pkg\util\ -coverprofile=coverage.txt -covermode count
	go tool cover -html=coverage
doc:
	godoc -http=:6060 -play