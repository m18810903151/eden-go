package router

type User struct {
	Name     string
	Age      string
	Gender   int
	HeadUrl  string
	Nickname string
}

func GetUserInfo() User {

	return User{"张三", "男", 12, "", "张三"}

}
