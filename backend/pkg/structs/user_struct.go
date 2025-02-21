package structs

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
}
