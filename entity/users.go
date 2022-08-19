package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id                 primitive.ObjectID `bson:"id,omitempty" json:"id,omitempty"`
	Username           string             `bson:"username,omitempty" json:"username,omitempty"`
	Nohandphone        string             `bson:"noHandphone,omitempty" json:"noHandphone,omitempty" validate:"required,max:15"`
	Macaddress         string             `bson:"macAddress,omitempty" json:"macAddress,omitempty" validate:"required"`
	Email              string             `bson:"email,omitempty" json:"email,omitempty" validate:"email"`
	Password           string             `bson:"password,omitempty" json:"password,omitempty" validate:"min:8"`
	Token              string             `bson:"token,omitempty" json:"token,omitempty"`
	TokenExpire        time.Time          `bson:"tokenExpire,omitempty" json:"tokenExpire,omitempty"`
	RefreshToken       string             `bson:"refreshToken,omitempty" json:"refreshToken,omitempty"`
	RefreshTokenExpire time.Time          `bson:"refreshTokenExpire,omitempty" json:"refreshTokenExpire,omitempty"`
	Otp                string             `bson:"otp,omitempty" json:"otp,omitempty"`
	OtpExpire          time.Time          `bson:"otpExpire,omitempty" json:"otpExpire,omitempty"`
	LoginTime          time.Time          `bson:"loginTime,omitempty" json:"loginTime,omitempty"`
	LoginClient        primitive.ObjectID `bson:"loginClient,omitempty" json:"loginClient,omitempty"`
	DeletedAt          time.Time          `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
	CreatedAt          time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt          time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
