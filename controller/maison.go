package MaisonController

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-demo/container"
	"go-demo/model"
	"go-demo/request"
	"gopkg.in/mgo.v2"
)

var (
	dig = container.Get();
)

func GetAction(c *gin.Context) {
	maison := model.Maison{
		Name: "chocapic",
	}
	content, _ := json.Marshal(maison)

	c.String(200, string(content))
};

func PostAction(context *gin.Context) {

	// Invoke dependencies
	var mongoHouseCollection *mgo.Collection
	dig.Invoke(func(mongo *mgo.Database) {
		mongoHouseCollection = mongo.C("maison");
	})

	var maison model.Maison;
	var request = http.Request{Request: context.Request};

	json.Unmarshal([]byte(request.GetBodyString()), &maison)

	mongoHouseCollection.Insert(maison)

	var content, _ = json.Marshal(maison)
	context.String(201, string(content))
};

