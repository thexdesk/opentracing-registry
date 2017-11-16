package opentracingmiddleware

import (
	"context"

	"github.com/docker/distribution"
	"github.com/docker/distribution/reference"
	middleware "github.com/docker/distribution/registry/middleware/repository"
)

func init() {
	middleware.Register("opentracing", NewRepository)
}

type opentracingRepository struct {
	distribution.Repository
}

// NewRepository creates a distribution.Repository wrapper that instruments
// distributed tracing using `github.com/opentracing/opentracing-go`.
func NewRepository(ctx context.Context, repository distribution.Repository, options map[string]interface{}) (distribution.Repository, error) {
	return &opentracingRepository{
		Repository: repository,
	}, nil
}

// Named returns the name of the repository.
func (r *opentracingRepository) Named() reference.Named {
	return r.Repository.Named()
}

// Manifests returns a reference to this repository's manifest service.
// with the supplied options applied.
func (r *opentracingRepository) Manifests(ctx context.Context, options ...distribution.ManifestServiceOption) (distribution.ManifestService, error) {
	manifests, err := r.Repository.Manifests(ctx, options...)
	if err != nil {
		return manifests, err
	}

	return &manifestService{
		ManifestService: manifests,
	}, nil
}

// Blobs returns a reference to this repository's blob service.
func (r *opentracingRepository) Blobs(ctx context.Context) distribution.BlobStore {
	return &blobStore{
		BlobStore: r.Repository.Blobs(ctx),
	}
}

// Tags returns a reference to this repositories tag service
func (r *opentracingRepository) Tags(ctx context.Context) distribution.TagService {
	return &tagService{
		TagService: r.Repository.Tags(ctx),
	}
}
