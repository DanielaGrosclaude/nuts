package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Usuario struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name,omitempty"`
	Adress    string             `bson:"adress" json:"adress,omitempty"`
	Birth     time.Time          `bson:"birth" json:"birth,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Biography string             `bson:"biography" json:"biography,omitempty"`
	Location  string             `bson:"location" json:"location,omitempty"`
	WebSite   string             `bson:"web_site" json:"web_site,omitempty"`
}
