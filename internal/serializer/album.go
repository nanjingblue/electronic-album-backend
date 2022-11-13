package serializer

import (
	"electronic-album/internal/model"
)

type Album struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	UserName  string `json:"username"`
	AlbumName string `json:"album_name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func BuildAlbum(it model.Album, username string) Album {
	return Album{
		ID:        it.ID,
		UserID:    it.UserID,
		UserName:  username,
		AlbumName: it.AlbumName,
		CreatedAt: it.CreatedAt.Unix(),
		UpdatedAt: it.UpdatedAt.Unix(),
	}
}
