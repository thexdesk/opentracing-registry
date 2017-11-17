package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/docker/cli/cli/command"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	lightstep "github.com/lightstep/lightstep-tracer-go"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/palantir/stacktrace"
)

const (
	pushImage = "ubuntu"
)

func main() {
	// span := opentracing.SpanFromContext(ctx)
	// if span != nil {
	// 	carrier := xmetaheaders.XMetaHeadersCarrier{
	// 		TextMapWriter: opentracing.HTTPHeadersCarrier(headers),
	// 	}
	// 	opentracing.GlobalTracer().Inject(span.Context(), opentracing.HTTPHeaders, carrier)
	// }

	if len(os.Args) != 2 {
		log.Fatal("requires exactly one arg")
	}

	ctx := context.Background()
	lightstepTracer := lightstep.NewTracer(lightstep.Options{
		AccessToken: os.Getenv("LIGHTSTEP_ACCESS_TOKEN"),
	})
	defer lightstepTracer.Close(ctx)

	opentracing.SetGlobalTracer(lightstepTracer)

	err := push(ctx, os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
}

func push(ctx context.Context, host string) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return stacktrace.Propagate(err, "failed to create docker client from env")
	}

	pullStream, err := cli.ImagePull(ctx, pushImage, types.ImagePullOptions{})
	if err != nil {
		return stacktrace.Propagate(err, "failed to initiate image pull '%s'", pushImage)
	}

	// Complete `docker pull pushImage`.
	_, err = ioutil.ReadAll(pullStream)
	if err != nil {
		return stacktrace.Propagate(err, "failed to read image pull stream")
	}

	pushTag := fmt.Sprintf("%s/opentracing/%s", host, pushImage)
	err = cli.ImageTag(ctx, pushImage, pushTag)
	if err != nil {
		return stacktrace.Propagate(err, "failed to tag '%s'", pushImage)
	}

	encodedAuth, err := command.EncodeAuthToBase64(types.AuthConfig{})
	if err != nil {
		return stacktrace.Propagate(err, "failed to encode auth config")
	}

	span := opentracing.StartSpan("APIClient.ImagePush")
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)
	pushStream, err := cli.ImagePush(ctx, pushTag, types.ImagePushOptions{
		RegistryAuth:  encodedAuth,
		PrivilegeFunc: registryAuthenticationPrivilegedFunc,
	})
	if err != nil {
		return stacktrace.Propagate(err, "failed to initiate image push '%s'", pushTag)
	}

	// Complete `docker push host/opentracing/pushImage`.
	_, err = ioutil.ReadAll(pushStream)
	if err != nil {
		return stacktrace.Propagate(err, "failed to read image push stream")
	}

	return nil
}

func registryAuthenticationPrivilegedFunc() (string, error) {
	return command.EncodeAuthToBase64(types.AuthConfig{})
}
