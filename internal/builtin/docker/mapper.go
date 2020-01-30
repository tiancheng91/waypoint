package docker

import (
	"github.com/mitchellh/devflow/internal/builtin/pack"
)

// NewRegistry is the factory function to create a Registry.
func NewRegistry() *Registry {
	return &Registry{}
}

// PackImageMapper maps a pack.DockerImage to our Image structure.
//
// NOTE(mitchellh): the pack builder can probably just reuse the image
// from here but at the time of writing I was still building all the
// mapper subsystems so I wanted to test it out.
func PackImageMapper(src *pack.DockerImage) *Image {
	return &Image{
		Image: src.Image,
		Tag:   src.Tag,
	}
}
