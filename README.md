## Threecolts Golang test

### Getting ready

Welcome to the Threecolts hiring test. Please make sure your development environment is ready:

- Golang 1.18.X https://go.dev/doc/install

Then type in a command line:

```
go run main.go
```

Optional: you can run tests with jest:

```
go test ./...
```

##output
```
#go test ./... -v
=== RUN   TestCountUniqueUrls
    main_test.go:52: TestCountUniqueUrls:
    main_test.go:58: Input: [https://example.com https://example.com/]  Output: 1
    main_test.go:58: Input: [https://example.com http://example.com]  Output: 2
    main_test.go:58: Input: [https://example.com? https://example.com]  Output: 1
    main_test.go:58: Input: [https://example.com?a=1&b=2 https://example.com?b=2&a=1]  Output: 1
--- PASS: TestCountUniqueUrls (0.00s)
=== RUN   TestCountUniqueUrlsPerTopLevelDomain
    main_test.go:66: TestingCountUniqueUrlsPerTopLevelDomain:
    main_test.go:72: Input: [https://example.com]  Output: map[example.com:1]
    main_test.go:72: Input: [https://example.com https://subdomain.example.com]  Output: map[example.com:2]
--- PASS: TestCountUniqueUrlsPerTopLevelDomain (0.00s)
PASS
ok      setup   0.460s
```
