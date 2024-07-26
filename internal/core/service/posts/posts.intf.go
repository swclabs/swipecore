package posts

import (
	"context"
	"mime/multipart"
	"swclabs/swipecore/internal/core/domain/dto"
)

// IPostsService : Module for managing posts.
// Actor: Admin & Customer
type IPostsService interface {
	// UploadCollections uploads a new collection.
	// ctx is the context to manage the request's lifecycle.
	// banner contains the collection details to be uploaded.
	// Returns id of collection was uploaded and error if any issues occur during the upload process.
	UploadCollections(ctx context.Context, banner dto.Collection) (int64, error)

	// UploadCollectionsImage uploads a new image of collection.
	// ctx is the context to manage the request's lifecycle.
	// cardBannerID contains the id of collection to be uploaded.
	// fileHeader is  the header of the file to be uploaded
	// Returns an error if any issues occur during the upload process.
	UploadCollectionsImage(ctx context.Context, cardBannerID string, fileHeader *multipart.FileHeader) error

	// SlicesOfCollections return a slices of collection.
	// ctx is the context to manage the request's lifecycle.
	// cardBannerID contains the id of collection to be returns.
	// limit is the maximum number of Collection to retrieve.
	// Returns an error if any issues occur during the upload process.
	SlicesOfCollections(ctx context.Context, position string, limit int) (*dto.Collections, error)

	UploadHeadlineBanner(ctx context.Context, banner dto.HeadlineBanner) error
	SliceOfHeadlineBanner(ctx context.Context, position string, limit int) (*dto.HeadlineBanners, error)
}
