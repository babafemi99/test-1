package videoRepo

import (
	"context"
	"test1/internal/entities/cloudinary"
)

type VideoRepository interface {
	PersistVideo(ctx context.Context, path string) (cloudinary.Response, error)
}
