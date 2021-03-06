module github.com/nodece/casbin-hraft-dispatcher

go 1.15

require (
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible
	github.com/casbin/casbin/v2 v2.23.4
	github.com/cenkalti/backoff/v4 v4.1.0
	github.com/go-chi/chi v1.5.1
	github.com/golang/mock v1.4.4
	github.com/hashicorp/go-multierror v1.1.0
	github.com/hashicorp/raft v1.2.0
	github.com/hashicorp/raft-boltdb v0.0.0-20171010151810-6e5ba93211ea
	github.com/json-iterator/go v1.1.10
	github.com/pkg/errors v0.8.1
	github.com/smartystreets/goconvey v1.6.4
	github.com/stretchr/testify v1.6.1
	go.etcd.io/bbolt v1.3.5
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20201202161906-c7110b5ffcbb
	google.golang.org/protobuf v1.25.0
)
