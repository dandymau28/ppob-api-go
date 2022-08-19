package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"ppob-api.go/ppob-api/entity"
	"ppob-api.go/ppob-api/helper"
)

type UserRepositoryImpl struct {
	client *mongo.Client
}

func NewUserRepository(client *mongo.Client) UserRepository {
	return &UserRepositoryImpl{
		client: client,
	}
}

func (repository *UserRepositoryImpl) FindByPhone(phone string) (entity.User, error) {
	userCollection := repository.client.Database("test").Collection("users")

	var user entity.User

	ctx := context.TODO()
	err := userCollection.FindOne(ctx, bson.M{"noHandphone": phone}).Decode(&user)
	if err != nil {
		user = entity.User{}
		if err == mongo.ErrNoDocuments {
			return user, errors.New("no documents found")
		}
		return user, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) Save(user *entity.User) error {
	userCollection := helper.GetCollection(repository.client, "users")
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	res, err := userCollection.InsertOne(context.TODO(), user)

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		user.Id = oid
	}

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (repository *UserRepositoryImpl) Update(user *entity.User) error {
	userCollection := helper.GetCollection(repository.client, "users")
	user.OtpExpire = time.Now().Add(2 * time.Minute)
	user.UpdatedAt = time.Now()
	fmt.Println(user.Otp)

	filter := bson.D{{Key: "id", Value: user.Id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "otp", Value: user.Otp}, {Key: "otpExpire", Value: user.OtpExpire}, {Key: "updateAt", Value: user.UpdatedAt}}}}

	result, err := userCollection.UpdateOne(context.TODO(), filter, update)

	fmt.Println(result)

	return err
}
