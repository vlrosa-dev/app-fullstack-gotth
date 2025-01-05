package structs

type User struct {
	FirstName string `form:"first_name"`
	LastName  string `form:"last_name"`
	Email     string `form:"email"`
	Password  string `form:"password"`
}
