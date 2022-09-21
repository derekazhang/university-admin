package request

type ManageAdminLoginParam struct {
	UserName    string `json:"username"`
	PasswordMd5 string `json:"passwordMd5"`
}

type ManageAdminParam struct {
	LoginUserName string `json:"loginUserName"`
	LoginPassword string `json:"loginPassword"`
	NickName      string `json:"nickName"`
}
