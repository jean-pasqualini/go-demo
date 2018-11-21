package container

import (
	"gopkg.in/mgo.v2"
)
import "go.uber.org/dig"

func Get() *dig.Container {
	return build()
}

func build() *dig.Container {
	container := dig.New()
	container.Provide(func() *mgo.Database {
		session, _ := mgo.Dial("127.0.0.1")
		session.SetMode(mgo.Monotonic, true)

		return session.DB("test")
	})

	return container;
}