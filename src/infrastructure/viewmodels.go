package infrastructure

import "github.com/allen-utec/vota-api-users/src/domain"

type UserVM struct {
	ID       uint   `json:"id"`
	Nickname string `json:"nickname"`
}

func formatUser(user domain.User) UserVM {
	userVM := UserVM{
		ID:       user.ID,
		Nickname: user.Nickname,
	}
	return userVM
}
