package contract

import (
	"github.com/codestates/WBA-BC-Project-02/common/model/entity"
	wasCommon "github.com/codestates/WBA-BC-Project-02/was/common"
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
	"github.com/codestates/WBA-BC-Project-02/was/model/query"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dexInstance *dexContractModel

type dexContractModel struct {
	collection *mongo.Collection
}

func NewDexContractModel(col *mongo.Collection) *dexContractModel {
	if dexInstance != nil {
		return dexInstance
	}
	dexInstance = &dexContractModel{
		collection: col,
	}
	return dexInstance
}

func (d *dexContractModel) FindDexContractNonTxHashes(dexAddr string) (*entity.DexContract, error) {
	filter := query.GetContractAddressFilter(dexAddr)

	opt := options.FindOne().SetProjection(bson.M{enum.TransactionHashes: 0})

	dex := &entity.DexContract{}

	if err := query.NewFindAction(dex, d.collection).
		InjectFilter(filter).
		FindOne(opt); err != nil {
		return nil, err
	}

	return dex, nil
}

func (d *dexContractModel) InsertOne(dexContract *entity.DexContract) error {
	ctx, cancel := wasCommon.NewContext(wasCommon.ModelContextTimeOut)
	defer cancel()

	if _, err := d.collection.InsertOne(ctx, dexContract); err != nil {
		return err
	}
	return nil
}
