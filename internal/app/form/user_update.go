package form

type UserUpdateForm struct {
	PasswordOld string `body:"passwordOld"`
	Password    string `body:"password"`
}
