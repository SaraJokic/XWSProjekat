package model

type LoginRequest struct {
	Username string `bson:"username" json:"username" form:"username" binding:"required"`
	Password string `bson:"password" json:"password" form:"password" binding:"required"`
}
