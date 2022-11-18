package serializer

import "electronic-album/internal/model"

type User struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Nickname    string `json:"nickname"`
	Sex         string `json:"sex"`
	Age         uint   `json:"age"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

func BuildUser(it model.User) User {
	return User{
		ID:          it.ID,
		Username:    it.Username,
		Nickname:    it.Nickname,
		Sex:         it.Sex,
		Age:         it.Age,
		Status:      it.Status,
		Description: it.Description,
	}
}
