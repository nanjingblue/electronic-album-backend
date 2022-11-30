package serializer

import "electronic-gallery/internal/model"

type User struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Nickname    string `json:"nickname"`
	Gender      string `json:"gender"`
	Age         uint   `json:"age"`
	Status      string `json:"status"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
}

func BuildUser(it model.User) User {
	return User{
		ID:          it.ID,
		Username:    it.Username,
		Nickname:    it.Nickname,
		Gender:      it.Gender,
		Age:         it.Age,
		Status:      it.Status,
		Description: it.Description,
		Avatar:      it.AvatarURl(),
	}
}
