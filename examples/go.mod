module github.com/calvinit/jiguang-sdk-go/examples

go 1.16

retract [v0.0.0-0, v0.0.0-99991231235959-zzzzzzzzzzzz] // Wiping out this module including all pseudo-versions.

require (
	github.com/calvinit/jiguang-sdk-go v0.3.0
	github.com/go-resty/resty/v2 v2.16.5
	github.com/rs/zerolog v1.33.0
	github.com/sirupsen/logrus v1.9.3
	go.uber.org/zap v1.24.0 // It's the latest version that supports go 1.16.
)

replace (
	github.com/calvinit/jiguang-sdk-go => ../
	golang.org/x/crypto => golang.org/x/crypto v0.33.0
	golang.org/x/mod => golang.org/x/mod v0.4.2
	golang.org/x/net => golang.org/x/net v0.17.0
	golang.org/x/sys => golang.org/x/sys v0.0.0-20201204225414-ed752295db88
	golang.org/x/tools => golang.org/x/tools v0.1.0
)