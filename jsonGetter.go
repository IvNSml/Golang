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
type SayMyName interface {
	SayName()
}

func(p Person) SayName()  {
	fmt.Println("My name is ", p.Name)
}
func main() {
	Vasya := CreatePerson()
	rand.Seed(time.Now().Unix())
	fmt.Println(Vasya.Nicknames[rand.Int()%len(Vasya.Nicknames)])
	Vasya.SayName()
}

func CreatePerson() Person{
	var text, err = ioutil.ReadFile("/home/ivan/Документы/homework.json")
	if err != nil {
		fmt.Println(err)
	}
	var man Person
	err =json.Unmarshal(text, &man)

	if err != nil {
		log.Fatal(err)
	}
	return man
}

