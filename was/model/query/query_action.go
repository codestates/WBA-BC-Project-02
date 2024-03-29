package query

import (
	wasCommon "github.com/codestates/WBA-BC-Project-02/was/common"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type queryAction[T any] struct {
	Result     T
	Filter     interface{}
	Update     interface{}
	collection *mongo.Collection
}

func NewFindAction[T any](t T, collection *mongo.Collection) *queryAction[T] {
	return &queryAction[T]{
		Result:     t,
		collection: collection,
	}
}

func (q *queryAction[T]) InjectFilter(filter interface{}) *queryAction[T] {
	q.Filter = filter
	return q
}

func (q *queryAction[T]) InjectUpdate(updateFilter interface{}) *queryAction[T] {
	q.Update = updateFilter
	return q
}

func (q *queryAction[T]) FindOne(opt *options.FindOneOptions) error {
	ctx, cancel := wasCommon.NewContext(wasCommon.ModelContextTimeOut)
	defer cancel()
	if err := q.collection.FindOne(ctx, q.Filter, opt).Decode(q.Result); err != nil {
		return err
	}
	return nil
}

func (q *queryAction[T]) FindOneAndUpdate(opt *options.FindOneAndUpdateOptions) error {
	ctx, cancel := wasCommon.NewContext(wasCommon.ModelContextTimeOut)
	defer cancel()

	if err := q.collection.FindOneAndUpdate(ctx, q.Filter, q.Update, opt).Decode(q.Result); err != nil {
		return err
	}
	return nil
}

func (q *queryAction[T]) Find(opt *options.FindOptions) error {
	ctx, cancel := wasCommon.NewContext(wasCommon.ModelContextTimeOut)
	defer cancel()

	cursor, err := q.collection.Find(ctx, q.Filter, opt)
	if err != nil {
		return err
	}

	if err = cursor.All(ctx, &q.Result); err != nil {
		return err
	}

	return nil
}
