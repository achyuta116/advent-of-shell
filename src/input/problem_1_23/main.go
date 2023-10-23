package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Animal struct {
	Species     string `json:"species"`
	Quantity    uint `json:"quantity"`
	LastVaxDate string `json:"last_vaccination_date"`
	NextVaxDate string `json:"due_vaccination_date"`
}

type Farmer struct {
	FarmSize  string `json:"size"`
	LastName  string `json:"last_name,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	Name      string `json:"name,omitempty"`
	Age       uint `json:"age"`
	Animals   []Animal `json:"animals"`
}

var (
	firstNames          = []string{"Alice", "Bob", "Charlie", "David", "Eva", "Frank", "Grace", "Henry", "Ivy", "Jack"}
	lastNames           = []string{"Smith", "Johnson", "Brown", "Davis", "Wilson", "Evans", "Harris", "Lee", "Robinson", "Wong"}
	domesticAnimalNames = []string{
		"Dog", "Cat", "Cow", "Horse", "Sheep", "Chicken", "Pig", "Goat", "Donkey", "Rabbit", "Turkey",
	}
)

func main() {
	seed, _ := strconv.Atoi(os.Getenv("SEED_VALUE"))

	rand.Seed(int64(seed))

    farmers := []Farmer{}
	for i := 0; i < 100; i++ {
		farmer := Farmer{}
		firstName := firstNames[rand.Intn(len(firstNames))]
		lastName := lastNames[rand.Intn(len(lastNames))]

		if rand.Intn(2) == 0 {
			farmer.Name = firstName + " " + lastName
		} else {
			farmer.FirstName = firstName
			farmer.LastName = lastName
		}

		farmer.Age = uint(rand.Intn(35)) + 18

		numAnimals := rand.Intn(5) + 5
		if rand.Intn(2) == 0 {
			if numAnimals > 5 {
				farmer.FarmSize = "Large"
			} else {
				farmer.FarmSize = "Small"
			}
		} else {
			if numAnimals > 5 {
				farmer.FarmSize = "Small"
			} else {
				farmer.FarmSize = "Large"
			}
		}

		for j := 0; j < numAnimals; j++ {
			animal := Animal{
				Species:  domesticAnimalNames[rand.Intn(len(domesticAnimalNames))],
				Quantity: uint(rand.Intn(5)) + 5,
			}

			startDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
			endDate := time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC)
            
            duration := endDate.Sub(startDate)
            randomDuration := time.Duration(rand.Int63n(int64(duration)))

            firstDate := startDate.Add(randomDuration)
            secondDate := firstDate.AddDate(0, rand.Intn(5) + 1, 0)

            if rand.Intn(2) == 0 {
                firstDate, secondDate = secondDate, firstDate
            }
            layout := "2006-01-02T15:04:05Z"
            animal.LastVaxDate = firstDate.Format(layout)
            animal.NextVaxDate = secondDate.Format(layout)

			farmer.Animals = append(farmer.Animals, animal)
		}

        farmers = append(farmers, farmer)
	}

    jsonBytes, err := json.Marshal(farmers)
    if err != nil {
        log.Fatalln("Marshaling failed")
    }

    fmt.Print(string(jsonBytes))
}
