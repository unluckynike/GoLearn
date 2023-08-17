package models

type User struct {
	//将json标签正确地包围在反引号内，并且添加了访问修饰符，以便可以在其他包中访问该结构体。
	Username string `json:"username"`
	Password string `json:"password"`
}
