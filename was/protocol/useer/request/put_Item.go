package request

import (
	"github.com/codestates/WBA-BC-Project-02/was/model/entity"
	"github.com/codestates/WBA-BC-Project-02/was/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	UserID    string `json:"user_id"`
	BlackIron int    `json:"black_iron"`
}

func (i *Item) ToPutItemUser(ID string) (*entity.User, error) {
	OBJID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}
	return &entity.User{
		ID: OBJID,
		Item: &dom.Item{
			BlackIron: i.BlackIron,
		},
	}, nil
}
