package opentracingmiddleware

import (
	"context"
	"io"
	"time"

	"github.com/docker/distribution"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type blobWriter struct {
	distribution.BlobWriter
}

// Write writes len(p) bytes from p to the underlying data stream.
// It returns the number of bytes written from p (0 <= n <= len(p))
// and any error encountered that caused the write to stop early.
// Write must return a non-nil error if it returns n < len(p).
// Write must not modify the slice data, even temporarily.
//
// Implementations must not retain p.
func (w *blobWriter) Write(p []byte) (n int, err error) {
	return w.BlobWriter.Write(p)
}

// The behavior of Close after the first call is undefined.
// Specific implementations may document their own behavior.
func (w *blobWriter) Close() error {
	return w.BlobWriter.Close()
}

// ReadFrom reads data from r until EOF or error.
// The return value n is the number of bytes read.
// Any error except io.EOF encountered during the read is also returned.
//
// The Copy function uses ReaderFrom if available.
func (w *blobWriter) ReadFrom(r io.Reader) (n int64, err error) {
	return w.BlobWriter.ReadFrom(r)
}

// Size returns the number of bytes written to this blob.
func (w *blobWriter) Size() int64 {
	return w.BlobWriter.Size()
}

// ID returns the identifier for this writer. The ID can be used with the
// Blob service to later resume the write.
func (w *blobWriter) ID() string {
	return w.BlobWriter.ID()
}

// StartedAt returns the time this blob write was started.
func (w *blobWriter) StartedAt() time.Time {
	return w.BlobWriter.StartedAt()
}

// Commit completes the blob writer process. The content is verified
// against the provided provisional descriptor, which may result in an
// error. Depending on the implementation, written data may be validated
// against the provisional descriptor fields. If MediaType is not present,
// the implementation may reject the commit or assign "application/octet-
// stream" to the blob. The returned descriptor may have a different
// digest depending on the blob store, referred to as the canonical
// descriptor.
func (w *blobWriter) Commit(ctx context.Context, provisional distribution.Descriptor) (canonical distribution.Descriptor, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "BlobWriter.Commit")
	defer span.Finish()
	span.LogFields(
		log.String("digest", provisional.Digest.String()),
	)

	canonical, err = w.BlobWriter.Commit(ctx, provisional)
	span.LogFields(
		log.String("mediaType", canonical.MediaType),
		log.Int64("size", canonical.Size),
	)
	return
}

// Cancel ends the blob write without storing any data and frees any
// associated resources. Any data written thus far will be lost. Cancel
// implementations should allow multiple calls even after a commit that
// result in a no-op. This allows use of Cancel in a defer statement,
// increasing the assurance that it is correctly called.
func (w *blobWriter) Cancel(ctx context.Context) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "BlobWriter.Cancel")
	defer span.Finish()

	return w.BlobWriter.Cancel(ctx)
}
