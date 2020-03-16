package url

import (
	"context"
	"github.com/error2215/url_shortener/server/db/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindShortenedUrl(ctx context.Context, url string) (string, error) {
	res := make(map[string]interface{})
	col := mongo.GetClient().Database("db").Collection("url")
	filter := bson.D{{"url", url}}
	projection := options.FindOne().SetProjection(bson.D{{"_id", 1}})
	err := col.FindOne(ctx, filter, projection).Decode(&res)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return "", nil
		}
		return "", err
	}
	return res["_id"].(primitive.ObjectID).Hex(), nil
}

func CreateShortenedUrl(ctx context.Context, url string) (string, error) {
	collection := mongo.GetClient().Database("db").Collection("url")
	data := map[string]string{"url": url}
	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func FindRealUrl(ctx context.Context, shortened string) (string, error) {
	res := make(map[string]interface{})
	col := mongo.GetClient().Database("db").Collection("url")
	id, err := primitive.ObjectIDFromHex(shortened)
	if err != nil {
		return "", err
	}
	filter := bson.D{{"_id", id}}
	projection := options.FindOne().SetProjection(bson.D{{"url", 1}})
	err = col.FindOne(ctx, filter, projection).Decode(&res)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return "", nil
		}
		return "", err
	}
	return res["url"].(string), nil
}
