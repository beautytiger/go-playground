package main

import (
	"encoding/json"
	"fmt"
)

type City struct {
	Name string `json:"name"`
	Area int	`json:"area"`
}

type Street struct {
	Name string	`json:"name"`
	Length int `json:"length"`
	City *City `json:"city"`
}

type Building struct {
	Name string `json:"building_name"`
	Height int	`json:"building_height"`
	Floor int	`json:"building_f"`
	Street *Street `json:"street"`
}

var city City = City{
	Name: "xian",
	Area: 199,
}

var street Street = Street{
	Name: "nanjing road",
	Length: 15,
	City: &city,
}

var building Building = Building{
	Name:   "lujiazui",
	Height: 388,
	Street: &street,
}

func main() {
	r, err := json.Marshal(building)
	if err != nil {
		fmt.Println(err)
	}
	s := string(r)
	fmt.Println(s)
	b := new(Building)
	j := `{"building_name":"lujiazui","building_height":388,"building_f":0,"street":{"name":"nanjing road","length":15,"city":{"name":"xian","area":199}}}`
	err = json.Unmarshal([]byte(j), &b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
}