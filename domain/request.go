package domain

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

//type RequestID string

type Request struct {
	Id          bson.ObjectId `valid:"-" json:"id" db:"_id" bson:"_id"`
	Name        string        `valid:"required" json:"name" db:"name"`
	Url         string        `valid:"required" json:"url" db:"url"`
	Status      string        `json:"status" db:"status"`
	Environment string        `valid:"-" json:"environment" db:"environment"`
	CreatedAt   time.Time     `valid:"-" json:"created_at" db:"created_at"`
	UpdatedAt   time.Time     `valid:"-" json:"updated_at" db:"updated_at"`
}

/*func NewRequest(id RequestID) *Request {
	return &Request{
		RequestID: id,
		Environment: "DEFAULT",
	}
}*/
