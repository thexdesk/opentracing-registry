package opentracingmiddleware

import (
	"context"

	"github.com/docker/distribution"
	"github.com/opencontainers/go-digest"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type manifestService struct {
	distribution.ManifestService
}

// Exists returns true if the manifest exists.
func (s *manifestService) Exists(ctx context.Context, dgst digest.Digest) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ManifestService.Exists")
	defer span.Finish()
	span.LogFields(
		log.String("digest", dgst.String()),
	)

	return s.ManifestService.Exists(ctx, dgst)
}

// Get retrieves the manifest specified by the given digest
func (s *manifestService) Get(ctx context.Context, dgst digest.Digest, options ...distribution.ManifestServiceOption) (distribution.Manifest, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ManifestService.Get")
	defer span.Finish()
	span.LogFields(
		log.String("digest", dgst.String()),
	)

	return s.ManifestService.Get(ctx, dgst, options...)
}

// Put creates or updates the given manifest returning the manifest digest
func (s *manifestService) Put(ctx context.Context, manifest distribution.Manifest, options ...distribution.ManifestServiceOption) (digest.Digest, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ManifestService.Put")
	defer span.Finish()

	return s.ManifestService.Put(ctx, manifest, options...)
}

// Delete removes the manifest specified by the given digest. Deleting
// a manifest that doesn't exist will return ErrManifestNotFound
func (s *manifestService) Delete(ctx context.Context, dgst digest.Digest) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ManifestService.Delete")
	defer span.Finish()
	span.LogFields(
		log.String("digest", dgst.String()),
	)

	return s.ManifestService.Delete(ctx, dgst)
}
