package main

import (
	"context"
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
	// defer r.Body.Close() // ensure that request body is always closed.

	// carrier := xmetaheaders.XMetaHeadersCarrier{
	// 	TextMapReader: opentracing.HTTPHeadersCarrier(r.Header),
	// }
	// wireContext, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, carrier)
	// if err != nil {
	// 	logrus.Warnf("failed to extract opentracing headers: %s", err)
	// }

	// span := opentracing.StartSpan("App.ServeHTTP", ext.RPCServerOption(wireContext))
	// defer span.Finish()
	// span.LogFields(
	// 	log.String("headers", fmt.Sprintf("%s", r.Header)),
	// 	log.String("method", r.Method),
	// 	log.String("url", r.URL.String()),
	// )

	// // Prepare the context with our own little decorations.
	// ctx := opentracing.ContextWithSpan(r.Context(), span)

	ctx := context.Background()
	lightstepTracer := lightstep.NewTracer(lightstep.Options{
		AccessToken: os.Getenv("LIGHTSTEP_ACCESS_TOKEN"),
	})
	defer lightstepTracer.Close(ctx)

	opentracing.SetGlobalTracer(lightstepTracer)

	registry.RootCmd.Execute()
}
