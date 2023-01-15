package user

import (
	"github.com/codestates/WBA-BC-Project-02/common/model/entity"
	wasCommon "github.com/codestates/WBA-BC-Project-02/was/common"
	"go.mongodb.org/mongo-driver/mongo"
)

var instance *userModel

type userModel struct {
	collection *mongo.Collection
}

func NewUserModel(col *mongo.Collection) *userModel {
	if instance != nil {
		return instance
	}
	instance = &userModel{
		collection: col,
	}
	return instance
}

func (u *userModel) InsertUser(user *entity.User) error {
	ctx, cancel := wasCommon.NewContext(wasCommon.ModelContextTimeOut)
	defer cancel()

	if _, err := u.collection.InsertOne(ctx, user); err != nil {
		return err
	}
	return nil
}
