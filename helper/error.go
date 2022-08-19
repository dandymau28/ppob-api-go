package helper

import "go.mongodb.org/mongo-driver/mongo"

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicErrorMongo(err error) {
	if err != nil {
		if err == mongo.ErrNoDocuments {

			return
		}
		panic(err)
	}
}
