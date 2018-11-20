package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"io"
	"io/ioutil"
	"net/http"
	_ "net/http"
	"strings"
)

type Arbre struct {
	name string
}

type Maison struct {
	Name string `json:"title"`
}


type response1 struct {
	Page   int
}

func keepLines(s string, n int) string {
	result := strings.Join(strings.Split(s, "\n")[:n], "\n")
	return strings.Replace(result, "\r", "", -1)
}

func(m Maison) hello() string  {
	return "ma maison est " + m.Name;
}


func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		maison := Maison{
			Name: "chocapic",
		}
		content, error := json.Marshal(maison)

		if (error != nil) {
			panic(error);
		}

		c.String(200, string(content))
	})

	r.POST("/maison", func(context *gin.Context) {
		var maison Maison
		session, _ := mgo.Dial("127.0.0.1")
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB("test").C("maison")

		requestBody, _ := ioutil.ReadAll(context.Request.Body)

		json.Unmarshal([]byte(requestBody), &maison)

		content, _ := json.Marshal(maison)

		_ = c.Insert(maison)

		context.String(201, string(content))
	});
	r.Run("0.0.0.0:8181") // listen and serve on 0.0.0.0:8080
}

func aa() {
	reader := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")

	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		if (n < 4) {
			fmt.Print(string(p[:3]));
		} else {
			fmt.Print(string(p[0]), string(p[1]), string(p[2]), string(p[3]));
		}
		if (n > 5) {
			panic("err");
		}
		//fmt.Println(string(p[:n]))
	}

	b, err := ioutil.ReadAll(reader)



	return;

	if (err != nil) {
		panic(err)
	}

	println(string(b));
}

func other()  {
	panic("yolo")
	client := http.Client{};
	resp, err := client.Get("http://www.google.fr");

	if err != nil {
		fmt.Println("aaaaaaaaah")
		panic(err)
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	secoya := Maison{Name: "sequoya"}
	fmt.Println(secoya.hello());
}
