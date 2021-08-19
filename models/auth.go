package models

import "github.com/unknwon/com"

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	Db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	return auth.ID > 0
}

func GetUserId(username string) string {
	var auth Auth
	Db.Select("id").Where(Auth{Username: username}).First((&auth))

	return com.StrTo(auth.ID).String()

}
