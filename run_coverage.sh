go clean -testcache

go test ./... -coverprofile cover.out

go tool cover -func cover.out

rm cover.out