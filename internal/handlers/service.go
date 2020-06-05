package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	db "github.com/ssubedir/go-shoot-dashboard/internal/db"

	"go.mongodb.org/mongo-driver/bson"
)

type Services struct {
}

func NewServices() *Services {

	return &Services{}

}

func (s *Services) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	type status struct {
		Service string
		Status  bool
	}
	ToJSON(&status{"go-shoot", true}, w)
}

func (s *Services) Enqueued(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	Connections := db.DBConnection()
	mongo := Connections["mongodb"].(db.Mongo)
	mongoConn := mongo.Connect()
	defer mongoConn.Disconnect()
	//Mongo Setup
	client := mongoConn.MONGO
	goshoot := client.Database("go-shoot")
	enqueued := goshoot.Collection("enqueued")

	var eTask []EnqueuedTask

	serv, err := enqueued.Find(context.TODO(), bson.M{})

	type Error struct {
		Error string
	}

	if err != nil {

		ToJSON(Error{"Error fetching enqueued tasks"}, w)
	}
	if err = serv.All(context.TODO(), &eTask); err != nil {
		ToJSON(Error{"Error seralizing enqueued tasks"}, w)
	}

	ToJSON(eTask, w)

}

func (s *Services) Scheduled(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	Connections := db.DBConnection()
	mongo := Connections["mongodb"].(db.Mongo)
	mongoConn := mongo.Connect()
	defer mongoConn.Disconnect()
	//Mongo Setup
	client := mongoConn.MONGO
	goshoot := client.Database("go-shoot")
	scheduled := goshoot.Collection("scheduled")

	var sTask []ScheduledTask

	serv, err := scheduled.Find(context.TODO(), bson.M{})

	type Error struct {
		Error string
	}

	if err != nil {

		ToJSON(Error{"Error fetching scheduled tasks"}, w)
	}
	if err = serv.All(context.TODO(), &sTask); err != nil {
		ToJSON(Error{"Error seralizing scheduled tasks"}, w)
	}

	ToJSON(sTask, w)

}

func (s *Services) Successful(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	Connections := db.DBConnection()
	mongo := Connections["mongodb"].(db.Mongo)
	mongoConn := mongo.Connect()
	defer mongoConn.Disconnect()
	//Mongo Setup
	client := mongoConn.MONGO
	goshoot := client.Database("go-shoot")
	succeeded := goshoot.Collection("succeeded")

	var sTask []SucceededTask

	serv, err := succeeded.Find(context.TODO(), bson.M{})

	type Error struct {
		Error string
	}

	if err != nil {

		ToJSON(Error{"Error fetching succeeded tasks"}, w)
	}
	if err = serv.All(context.TODO(), &sTask); err != nil {
		ToJSON(Error{"Error seralizing succeeded tasks"}, w)
	}

	ToJSON(sTask, w)

}

func (s *Services) TaskSummary(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	Connections := db.DBConnection()
	mongo := Connections["mongodb"].(db.Mongo)
	mongoConn := mongo.Connect()
	defer mongoConn.Disconnect()
	//Mongo Setup
	client := mongoConn.MONGO
	goshoot := client.Database("go-shoot")
	succeeded := goshoot.Collection("succeeded")
	scheduled := goshoot.Collection("scheduled")
	enqueued := goshoot.Collection("enqueued")
	var suTask []SucceededTask
	var scTask []ScheduledTask
	var eTask []EnqueuedTask

	type Error struct {
		Error string
	}

	serv0, err := succeeded.Find(context.TODO(), bson.M{})
	if err != nil {

		ToJSON(Error{"Error fetching succeeded tasks"}, w)
	}
	if err = serv0.All(context.TODO(), &suTask); err != nil {
		ToJSON(Error{"Error seralizing succeeded tasks"}, w)
	}
	serv1, err1 := scheduled.Find(context.TODO(), bson.M{})
	if err1 != nil {

		ToJSON(Error{"Error fetching scheduled tasks"}, w)
	}
	if err = serv1.All(context.TODO(), &scTask); err != nil {
		ToJSON(Error{"Error seralizing scheduled tasks"}, w)
	}
	serv2, _ := enqueued.Find(context.TODO(), bson.M{})
	if err1 != nil {

		ToJSON(Error{"Error fetching enqueued tasks"}, w)
	}
	if err = serv2.All(context.TODO(), &eTask); err != nil {
		ToJSON(Error{"Error seralizing enqueued tasks"}, w)
	}

	type Resp struct {
		Enqueued  int
		Scheduled int
		Succeeded int
	}

	ToJSON(Resp{Enqueued: len(eTask), Scheduled: len(scTask), Succeeded: len(suTask)}, w)

}

// ToJSON serializes the given interface into a string based JSON format
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}
