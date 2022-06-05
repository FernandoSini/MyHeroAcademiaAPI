package models

type Password struct {
	NewPassword     string `bson:"newPassword" json:"newPassword"`
	CurrentPassword string `bson:"currentPassword" json:"currentPassword"`
}
