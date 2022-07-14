package main

import (
	"errors"
	"fmt"
	"net/http"
)

func LogOutput(message string) {
	fmt.Println(message)
}

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForId(userId string) (string, bool) {
	val, ok := sds.userData[userId]

	return val, ok
}

func NewSimpleDataStore() *SimpleDataStore {
	return &SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Mary",
			"3": "Pat",
		},
	}
}

type Datastore interface {
	UserNameForId(string) (string, bool)
}

type Logger interface {
	Log(string)
}

type LoggerAdapter func(string)

func (la LoggerAdapter) Log(message string) {
	la(message)
}

type SimpleLogic struct {
	l  Logger
	ds Datastore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	un, ok := sl.ds.UserNameForId(userID)

	sl.l.Log("in SayHello for " + userID)

	fmt.Println(un, ok)

	if !ok {
		return "", errors.New("unknown User")
	}

	return "Hello " + un, nil

}

func (sl SimpleLogic) SayGoodBye(userID string) (string, error) {
	un, ok := sl.ds.UserNameForId(userID)

	sl.l.Log("in SayGoodBye for " + userID)

	if !ok {
		return "", errors.New("unknown User")
	}

	return "GoodBye, " + un, nil

}

func NewSimpleLogic(l Logger, ds Datastore) *SimpleLogic {
	return &SimpleLogic{
		l, ds,
	}
}

type logic interface {
	SayHello(userId string) (string, error)
}

type Controller struct {
	l  Logger
	lg logic
}

func (c Controller) HandleGreeting(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In SayHello")

	userId := r.URL.Query().Get("user_id")
	message, err := c.lg.SayHello(userId)
	fmt.Println(userId, err)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(message))
}

func NewController(l Logger, lg logic) *Controller {
	return &Controller{l: l, lg: lg}
}

func main() {
	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)

	c := NewController(l, logic)

	http.HandleFunc("/", c.HandleGreeting)
	http.ListenAndServe(":8080", nil)
}
