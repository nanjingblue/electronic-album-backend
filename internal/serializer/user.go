package serializer

import "electronic-album/internal/model"

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Sex      string `json:"sex"`
	Age      uint   `json:"age"`
	Status   string `json:"status"`
}

func BuildUser(it model.User) User {
	return User{
		ID:       it.ID,
		Username: it.Username,
		Sex:      it.Sex,
		Age:      it.Age,
		Status:   it.Status,
	}
}
