package serializer

import "electronic-gallery/internal/model"

type Picture struct {
	ID          uint   `json:"id"`
	PictureName string `json:"picture_name"`
	URL         string `json:"url"`
	UserID      uint   `json:"user_id"`
	GalleryID   uint   `json:"gallery_id"`
	CreatedAt   int64  `json:"created_at"`
}

func BuildPicture(it model.Picture) Picture {
	return Picture{
		ID:          it.ID,
		PictureName: it.PictureName,
		URL:         it.CoverURl(),
		UserID:      it.UserID,
		GalleryID:   it.GalleryID,
		CreatedAt:   it.CreatedAt.Unix(),
	}
}

func BuildPictures(it []model.Picture) []Picture {
	var pictures []Picture
	for _, item := range it {
		pictures = append(pictures, BuildPicture(item))
	}
	return pictures
}
