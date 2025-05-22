module github.com/calvinit/jiguang-sdk-go/examples

go 1.24

retract [v0.0.0-0, v0.0.0-99991231235959-zzzzzzzzzzzz] // Wiping out this module including all pseudo-versions.

require (
	github.com/calvinit/jiguang-sdk-go v0.4.6
	github.com/go-resty/resty/v2 v2.16.5
	github.com/hashicorp/go-retryablehttp v0.7.7
	github.com/rs/zerolog v1.34.0
	github.com/sirupsen/logrus v1.9.3
	go.uber.org/zap v1.27.0
)

require (
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
)

replace github.com/calvinit/jiguang-sdk-go => ../
