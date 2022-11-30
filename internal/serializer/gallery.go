package serializer

import (
	"electronic-gallery/internal/model"
)

type Gallery struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	UserName    string `json:"username"`
	AlbumName   string `json:"gallery_name"`
	Cover       string `json:"cover"`
	Status      string `json:"status"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

func BuildGallery(it model.Gallery, username string) Gallery {
	return Gallery{
		ID:          it.ID,
		UserID:      it.UserID,
		UserName:    username,
		AlbumName:   it.GalleryName,
		Cover:       it.CoverURl(),
		Status:      it.Status,
		Description: it.Description,
		CreatedAt:   it.CreatedAt.Unix(),
		UpdatedAt:   it.UpdatedAt.Unix(),
	}
}

func BuildGallerys(it []model.Gallery, username string) (albums []Gallery) {
	for _, item := range it {
		albums = append(albums, BuildGallery(item, username))
	}
	return
}
