package model

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Id bson.ObjectId `bson:"_id"`
	FirstName string `bson:"firstName"`
	LastName string `bson:"lastName"`
	EMail string `bson:"email"`
	PwHash []byte `bson:"pwHash"`
	Roles []string `bson:"roles"`
}

type SessionUser struct {
	Id string `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	EMail string `json:"email"`
	Roles []string `json:"roles"`
}

func CreateUser(user *User) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	usrObjWithoutId := struct {
		FirstName string `bson:"firstName"`
		LastName string `bson:"lastName"`
		EMail string `bson:"email"`
		PwHash []byte `bson:"pwHash"`
		Roles []string `bson:"roles"`
	}{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		EMail:     user.EMail,
		PwHash:    user.PwHash,
		Roles:     user.Roles,
	}

	res, err := userCollection.InsertOne(ctx, usrObjWithoutId)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func LoadUserByEmail(email string) (*User, error) {
	usr := &User{}

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(usr)

	return usr, err
}

func LoadUserById(idHex string) (*User, error) {
	if !bson.IsObjectIdHex(idHex) {
		return nil, errors.New("not an ObjectId")
	}

	usr := &User{}

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	err := userCollection.FindOne(ctx, bson.ObjectIdHex(idHex)).Decode(usr)

	return usr, err
}

func (usr *User) ToSessionUser() *SessionUser {
	return &SessionUser{
		Id: usr.Id.Hex(),
		FirstName: usr.FirstName,
		LastName: usr.LastName,
		EMail: usr.EMail,
		Roles: usr.Roles,
	}
}
