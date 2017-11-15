package opentracingmiddleware

import (
	"context"

	"github.com/docker/distribution"
	"github.com/docker/distribution/digest"
)

type manifestService struct {
	distribution.ManifestService
}

// Exists returns true if the manifest exists.
func (s *manifestService) Exists(ctx context.Context, dgst digest.Digest) (bool, error) {
	return s.ManifestService.Exists(ctx, dgst)
}

// Get retrieves the manifest specified by the given digest
func (s *manifestService) Get(ctx context.Context, dgst digest.Digest, options ...distribution.ManifestServiceOption) (distribution.Manifest, error) {
	return s.ManifestService.Get(ctx, dgst, options)
}

// Put creates or updates the given manifest returning the manifest digest
func (s *manifestService) Put(ctx context.Context, manifest Manifest, options ...distribution.ManifestServiceOption) (digest.Digest, error) {
	return s.ManifestService.Put(ctx, manifest, options)
}

// Delete removes the manifest specified by the given digest. Deleting
// a manifest that doesn't exist will return ErrManifestNotFound
func (s *manifestService) Delete(ctx context.Context, dgst digest.Digest) error {
	return s.ManifestService.Delete(ctx, dgst)
}
