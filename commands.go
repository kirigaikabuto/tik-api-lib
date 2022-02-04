package tik_api_lib

type LoginCommand struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func (cmd *LoginCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(Service).Login(cmd)
}
