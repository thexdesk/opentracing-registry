package main

import (
	_ "net/http/pprof"
	"os"

	"github.com/docker/distribution/registry"
	_ "github.com/docker/distribution/registry/auth/htpasswd"
	_ "github.com/docker/distribution/registry/auth/silly"
	_ "github.com/docker/distribution/registry/auth/token"
	_ "github.com/docker/distribution/registry/proxy"
	_ "github.com/docker/distribution/registry/storage/driver/azure"
	_ "github.com/docker/distribution/registry/storage/driver/filesystem"
	_ "github.com/docker/distribution/registry/storage/driver/gcs"
	_ "github.com/docker/distribution/registry/storage/driver/inmemory"
	_ "github.com/docker/distribution/registry/storage/driver/middleware/cloudfront"
	_ "github.com/docker/distribution/registry/storage/driver/middleware/redirect"
	_ "github.com/docker/distribution/registry/storage/driver/oss"
	_ "github.com/docker/distribution/registry/storage/driver/s3-aws"
	_ "github.com/docker/distribution/registry/storage/driver/s3-goamz"
	_ "github.com/docker/distribution/registry/storage/driver/swift"
	_ "github.com/hinshun/opentracing-registry/opentracingmiddleware"
	"github.com/lightstep/lightstep-tracer-go"
	"github.com/opentracing/opentracing-go"
)

func main() {
	lightstepTracer := lightstep.NewTracer(lightstep.Options{
		AccessToken: os.Getenv("LIGHTSTEP_ACCESS_TOKEN"),
	})
	opentracing.SetGlobalTracer(lightstepTracer)
	registry.RootCmd.Execute()
}
