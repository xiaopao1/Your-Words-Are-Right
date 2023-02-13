package database

import "Your_Words_Are_Right/Init"

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

func (u *User) CheckUsernameIsExist(username string) bool {
	sql := " SELECT 1  FROM tb_users WHERE  user_name=?"
	result := Init.DB.Exec(sql, username)
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

func (u *User) CreateUser(username string, password string) (token string, userId int64) {
	//Init.DB.Create(&User{Name: username})
	return username + password, userId
}
