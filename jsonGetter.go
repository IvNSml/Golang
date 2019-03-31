package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

type Person struct {
	Name         string
	Age          int
	Nicknames    []string
	Achievements map[int]string
}
type Persons []Person
type SayMyName interface {
	SayName()
}

func (array Persons) SayName() {
	for _,person:=range array{
		fmt.Println("My name is ", person.Name)
		rand.Seed(time.Now().Unix())
		fmt.Println(person.Nicknames[rand.Int()%len(person.Nicknames)])
	}
}
func main() {
	people := CreatePerson()
	people.SayName()
}

func CreatePerson() Persons {
	var text, err = ioutil.ReadFile("/home/ivan/Документы/homework.json")
	if err != nil {
		fmt.Println(err)
	}
	var persons = make([]Person,0)
	var buff Person
	var isBreak string

	for true {
		err= json.Unmarshal(text, &buff)
		if err != nil {
			log.Fatal(err)
		}
		persons = append(persons, buff)
		_=json.Unmarshal(text,&isBreak)
		break//there is a incompleteness
	}
	return persons
}
