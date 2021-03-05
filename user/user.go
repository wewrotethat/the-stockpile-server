package user

//User is the model that holds information about a given person on the system
type User struct {
	ID       uint   `bson:"id"`
	UserName string `bson:"userName"`
	FullName string `bson:"fullName"`
	Email    string `bson:"email"`
	Phone    string `bson:"phone"`
	Password string `bson:"password"`
}
