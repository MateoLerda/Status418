package responses

type UserInfo struct {
	Code   string `json:code`
	Email    string `json:email`
	Username string `json:username`
	Role      string `json:role`
}