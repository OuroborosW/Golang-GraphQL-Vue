// resolvers.go
package main

import (
	"context"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
)

func SearchResolver(params graphql.ResolveParams) (interface{}, error) {
	input, isOK := params.Args["input"].(string)
	if isOK {
		// 使用正则表达式查询MongoDB以实现不区分大小写的前缀搜索
		filter := bson.D{
			{"name", bson.D{
				{"$regex", "^" + input},
				{"$options", "i"}, // 不区分大小写
			}},
		}
		cur, err := statesCollection.Find(context.Background(), filter)
		if err != nil {
			return nil, err
		}
		var results []map[string]interface{}
		for cur.Next(context.Background()) {
			var elem map[string]interface{}
			err := cur.Decode(&elem)
			if err != nil {
				return nil, err
			}
			results = append(results, elem)
		}
		return results, nil
	}
	return nil, nil
}
