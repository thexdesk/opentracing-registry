package opentracingmiddleware

import (
	"context"

	"github.com/docker/distribution"
	"github.com/docker/distribution/reference"
	middleware "github.com/docker/distribution/registry/middleware/registry"
)

func init() {
	middleware.Register("opentracing", NewRepository)
}

type repository struct {
	distribution.Repository

	manifestService distribution.ManifestService
	blobStore       distribution.BlobStore
	tagService      distribution.TagService
}

// NewRepository creates a distribution.Repository wrapper that instruments
// distributed tracing using `github.com/opentracing/opentracing-go`.
func NewRepository(ctx context.Context, repository distribution.Repository, options map[string]interface{}) (distribution.Repository, error) {
	return &repository{
		Repository: repository,
		manifestService: manifestService{
			repository: repository,
		},
	}, nil
}

// Named returns the name of the repository.
func (r *repository) Named() reference.Named {
	return r.Repository.Named()
}

// Manifests returns a reference to this repository's manifest service.
// with the supplied options applied.
func (r *repository) Manifests(ctx context.Context, options ...distribution.ManifestServiceOption) (distribution.ManifestService, error) {
	return r.manifestService, nil

	manifests, err := r.Repository.Manifests(ctx, options)
	if err != nil {
		return manifests, err
	}

	return manifestService{
		ManifestService: manifests,
	}, nil
}

// Blobs returns a reference to this repository's blob service.
func (r *repository) Blobs(ctx context.Context) distribution.BlobStore {
	return blobStore{
		BlobStore: r.Repository.Blobs(ctx),
	}
}

// Tags returns a reference to this repositories tag service
func (r *repository) Tags(ctx context.Context) distribution.TagService {
	return tagService{
		TagService: r.Repository.Tags(ctx),
	}
}
