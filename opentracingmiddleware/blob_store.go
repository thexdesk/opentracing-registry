package opentracingmiddleware

import (
	"context"
	"net/http"

	"github.com/docker/distribution"
	"github.com/opencontainers/go-digest"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type blobStore struct {
	distribution.BlobStore
}

// Stat provides metadata about a blob identified by the digest. If the
// blob is unknown to the describer, ErrBlobUnknown will be returned.
func (s *blobStore) Stat(ctx context.Context, dgst digest.Digest) (distribution.Descriptor, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "BlobStore.Stat")
	defer span.Finish()
	span.LogFields(
		log.String("digest", dgst.String()),
	)

	return s.BlobStore.Stat(ctx, dgst)
}

// Get returns the entire blob identified by digest along with the descriptor.
func (s *blobStore) Get(ctx context.Context, dgst digest.Digest) ([]byte, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "BlobStore.Get")
	defer span.Finish()
	span.LogFields(
		log.String("digest", dgst.String()),
	)

	return s.BlobStore.Get(ctx, dgst)
}

// Open provides a ReadSeekCloser to the blob identified by the provided
// descriptor. If the blob is not known to the service, an error will be
// returned.
func (s *blobStore) Open(ctx context.Context, dgst digest.Digest) (distribution.ReadSeekCloser, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "BlobStore.Open")
	defer span.Finish()
	span.LogFields(
		log.String("digest", dgst.String()),
	)

	return s.BlobStore.Open(ctx, dgst)
}

// Put inserts the content p into the blob service, returning a descriptor
// or an error.
func (s *blobStore) Put(ctx context.Context, mediaType string, p []byte) (distribution.Descriptor, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "BlobStore.Put")
	defer span.Finish()
	span.LogFields(
		log.String("mediaType", mediaType),
	)

	return s.BlobStore.Put(ctx, mediaType, p)
}

// Create allocates a new blob writer to add a blob to this service. The
// returned handle can be written to and later resumed using an opaque
// identifier. With this approach, one can Close and Resume a BlobWriter
// multiple times until the BlobWriter is committed or cancelled.
func (s *blobStore) Create(ctx context.Context, options ...distribution.BlobCreateOption) (distribution.BlobWriter, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "BlobStore.Create")
	defer span.Finish()

	return s.BlobStore.Create(ctx, options...)
}

// Resume attempts to resume a write to a blob, identified by an id.
func (s *blobStore) Resume(ctx context.Context, id string) (distribution.BlobWriter, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "BlobStore.Resume")
	defer span.Finish()
	span.LogFields(
		log.String("id", id),
	)

	return s.BlobStore.Resume(ctx, id)
}

// ServeBlob attempts to serve the blob, identifed by dgst, via http. The
// service may decide to redirect the client elsewhere or serve the data
// directly.
//
// This handler only issues successful responses, such as 2xx or 3xx,
// meaning it serves data or issues a redirect. If the blob is not
// available, an error will be returned and the caller may still issue a
// response.
//
// The implementation may serve the same blob from a different digest
// domain. The appropriate headers will be set for the blob, unless they
// have already been set by the caller.
func (s *blobStore) ServeBlob(ctx context.Context, w http.ResponseWriter, r *http.Request, dgst digest.Digest) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "BlobStore.ServeBlob")
	defer span.Finish()
	span.LogFields(
		log.String("digest", dgst.String()),
	)

	return s.BlobStore.ServeBlob(ctx, w, r, dgst)
}

func (s *blobStore) Delete(ctx context.Context, dgst digest.Digest) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "BlobStore.Delete")
	defer span.Finish()
	span.LogFields(
		log.String("digest", dgst.String()),
	)

	return s.BlobStore.Delete(ctx, dgst)
}
