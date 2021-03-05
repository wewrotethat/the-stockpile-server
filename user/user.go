package user

//User is the model that holds information about a given person on the system
type User struct {
	ID       uint   `bson:"id"        json:"id"`
	UserName string `bson:"userName"  json:"userName"`
	FullName string `bson:"fullName"  json:"fullName"`
	Email    string `bson:"email"     json:"email"`
	Phone    string `bson:"phone"     json:"phone"`
	Password string `bson:"password"  json:"-"`
}
