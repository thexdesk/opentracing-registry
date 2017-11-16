package opentracingmiddleware

import (
	"context"

	"github.com/docker/distribution"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type tagService struct {
	distribution.TagService
}

// Get retrieves the descriptor identified by the tag. Some
// implementations may differentiate between "trusted" tags and
// "untrusted" tags. If a tag is "untrusted", the mapping will be returned
// as an ErrTagUntrusted error, with the target descriptor.
func (s *tagService) Get(ctx context.Context, tag string) (distribution.Descriptor, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "TagService.Get")
	defer span.Finish()
	span.LogFields(
		log.String("tag", tag),
	)

	return s.TagService.Get(ctx, tag)
}

// Tag associates the tag with the provided descriptor, updating the
// current association, if needed.
func (s *tagService) Tag(ctx context.Context, tag string, desc distribution.Descriptor) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "TagService.Tag")
	defer span.Finish()
	span.LogFields(
		log.String("tag", tag),
		log.String("mediaType", desc.MediaType),
		log.String("digest", desc.Digest.String()),
		log.Int64("size", desc.Size),
	)

	return s.TagService.Tag(ctx, tag, desc)
}

// Untag removes the given tag association
func (s *tagService) Untag(ctx context.Context, tag string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "TagService.Untag")
	defer span.Finish()
	span.LogFields(
		log.String("tag", tag),
	)

	return s.TagService.Untag(ctx, tag)
}

// All returns the set of tags managed by this tag service
func (s *tagService) All(ctx context.Context) ([]string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "TagService.All")
	defer span.Finish()

	return s.TagService.All(ctx)
}

// Lookup returns the set of tags referencing the given digest.
func (s *tagService) Lookup(ctx context.Context, desc distribution.Descriptor) ([]string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "TagService.Lookup")
	defer span.Finish()
	span.LogFields(
		log.String("mediaType", desc.MediaType),
		log.String("digest", desc.Digest.String()),
		log.Int64("size", desc.Size),
	)

	return s.TagService.Lookup(ctx, desc)
}
