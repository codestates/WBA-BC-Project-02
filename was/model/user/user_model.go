package user

import (
	"github.com/codestates/WBA-BC-Project-02/common/model/entity"
	wasCommon "github.com/codestates/WBA-BC-Project-02/was/common"
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
	"github.com/codestates/WBA-BC-Project-02/was/model/query"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (u *userModel) FindUser(address string) (*entity.User, error) {
	filter := query.GetAddressFilter(address)

	user := &entity.User{}

	if err := query.NewFindAction(user, u.collection).
		InjectFilter(filter).
		FindOne(nil); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userModel) FindUserAndPWDUpdate(address, password string) (*entity.User, error) {
	f := query.GetAddressFilter(address)

	upf := query.GetUpdatePWDFilter(password)

	prj := options.FindOneAndUpdate().SetProjection(bson.M{enum.Transactions: 0})

	user := &entity.User{}

	if err := query.NewFindAction(user, u.collection).
		InjectFilter(f).
		InjectUpdate(upf).
		FindOneAndUpdate(prj); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userModel) FindUserAndIncreaseIron(address string) (*entity.User, error) {
	f := query.GetAddressFilter(address)

	upf := query.GetBlackIronIncreaseFilter()

	prj := options.FindOneAndUpdate().SetReturnDocument(options.After)

	user := &entity.User{}

	if err := query.NewFindAction(user, u.collection).
		InjectFilter(f).
		InjectUpdate(upf).
		FindOneAndUpdate(prj); err != nil {
		return nil, err
	}

	return user, nil
}
