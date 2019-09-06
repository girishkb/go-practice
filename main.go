package main

import (
	f "fmt"
	g "github.com/golang/protobuf/proto"
	l "log"
)

func main()  {
	elliot := &Person{
		Name: "Girish",
		Age:  30,
	}
	data,err := g.Marshal(elliot)
	if err != nil {
		l.Fatal("Marshmalling error")
	}
	f.Println(data)
	newElliot := &Person{}
	err =  g.Unmarshal(data,newElliot)
	if err != nil {
		l.Fatal("UnMarshmalling error")
	}
	f.Println(newElliot)

	finalElliot := &Person{
		Name:            "",
		Age:             0,
		SocialFollowers: &SocialFollowers{
			Youtube: 2500,
			Twitter: 1400,
		},
	}

	f.Println(finalElliot)
}
