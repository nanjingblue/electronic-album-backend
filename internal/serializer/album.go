package serializer

import (
	"electronic-album/internal/model"
)

type Album struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	UserName  string `json:"username"`
	AlbumName string `json:"album_name"`
	Image     string `json:"image"`
	Status    int    `json:"status"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func BuildAlbum(it model.Album, username string) Album {
	return Album{
		ID:        it.ID,
		UserID:    it.UserID,
		UserName:  username,
		AlbumName: it.AlbumName,
		Image:     it.Image,
		Status:    it.Status,
		CreatedAt: it.CreatedAt.Unix(),
		UpdatedAt: it.UpdatedAt.Unix(),
	}
}

func BuildAlbums(it []model.Album, username string) (albums []Album) {
	for _, item := range it {
		albums = append(albums, BuildAlbum(item, username))
	}
	return
}
