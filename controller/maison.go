package MaisonController

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-demo/model"
	"gopkg.in/mgo.v2"
	"io/ioutil"
)

func GetAction(c *gin.Context) {

	maison := model.Maison{
		Name: "chocapic",
	}
	content, _ := json.Marshal(maison)

	c.String(200, string(content))
};

func PostAction(context *gin.Context) {
	var maison model.Maison
	session, _ := mgo.Dial("127.0.0.1")
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("maison")

	requestBody, _ := ioutil.ReadAll(context.Request.Body)

	json.Unmarshal([]byte(requestBody), &maison)

	content, _ := json.Marshal(maison)

	_ = c.Insert(maison)

	context.String(201, string(content))
};

