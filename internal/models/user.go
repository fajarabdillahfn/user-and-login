package models

type User struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"-"`
	Role     string `bson:"role" json:"role"`
}
