package response

import "Your_Words_Are_Right/model/database"

type UserInfoResponse struct {
	Response
	User database.User `json:"user"`
}
type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}
