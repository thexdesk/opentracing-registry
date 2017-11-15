package opentracingmiddleware

import (
	"context"

	"github.com/docker/distribution"
)

type tagService struct {
	distribution.TagService
}

// Get retrieves the descriptor identified by the tag. Some
// implementations may differentiate between "trusted" tags and
// "untrusted" tags. If a tag is "untrusted", the mapping will be returned
// as an ErrTagUntrusted error, with the target descriptor.
func (s *tagService) Get(ctx context.Context, tag string) (distribution.Descriptor, error) {
	return s.TagService.Get(ctx, tag)
}

// Tag associates the tag with the provided descriptor, updating the
// current association, if needed.
func (s *tagService) Tag(ctx context.Context, tag string, desc distribution.Descriptor) error {
	return s.TagService.Tag(ctx, tag, desc)
}

// Untag removes the given tag association
func (s *tagService) Untag(ctx context.Context, tag string) error {
	return s.TagService.Untag(ctx, tag)
}

// All returns the set of tags managed by this tag service
func (s *tagService) All(ctx context.Context) ([]string, error) {
	return s.TagService.All(ctx)
}

// Lookup returns the set of tags referencing the given digest.
func (s *tagService) Lookup(ctx context.Context, digest distribution.Descriptor) ([]string, error) {
	return s.TagService.Lookup(ctx, digest)
}
