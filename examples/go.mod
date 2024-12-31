module github.com/calvinit/jiguang-sdk-go/examples

go 1.16

require (
	github.com/calvinit/jiguang-sdk-go v0.0.0
	github.com/go-resty/resty/v2 v2.16.2
	github.com/rs/zerolog v1.33.0
	github.com/sirupsen/logrus v1.9.3
	go.uber.org/zap v1.24.0 // It's the latest version that supports go 1.16.
)

replace (
	github.com/calvinit/jiguang-sdk-go => ../
	golang.org/x/mod => golang.org/x/mod v0.4.2
	// It's the latest version that supports go 1.16, as later versions removed `// +build ignore`,
	// and `//go:build ignore` is not supported in go 1.16.
	golang.org/x/net => golang.org/x/net v0.17.0
	// The `unsafe.Slice` function was added in go 1.17: https://pkg.go.dev/unsafe#Slice.
	golang.org/x/sys => golang.org/x/sys v0.0.0-20201204225414-ed752295db88
	golang.org/x/tools => golang.org/x/tools v0.1.0
)