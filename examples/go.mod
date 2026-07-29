module github.com/cavlabs/jiguang-sdk-go/examples

go 1.26

retract [v0.0.0-0, v0.0.0-99991231235959-zzzzzzzzzzzz] // Wiping out this module including all pseudo-versions.

require (
	github.com/cavlabs/jiguang-sdk-go v1.0.8
	github.com/go-resty/resty/v2 v2.17.2
	github.com/hashicorp/go-retryablehttp v0.7.8
	github.com/rs/zerolog v1.35.1
	github.com/sirupsen/logrus v1.9.4
	go.uber.org/zap v1.28.0
)

require (
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/mattn/go-colorable v0.1.15 // indirect
	github.com/mattn/go-isatty v0.0.24 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.57.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
)

replace github.com/cavlabs/jiguang-sdk-go => ../
