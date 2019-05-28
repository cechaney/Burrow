package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"time"

	"github.com/cechaney/burrow/core"
)

type animalsPageData struct {
	Dog string
	Cat string
	Fox string
}

var httpClient = &http.Client{Timeout: 10 * time.Second}

func animalsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	dogChannel := make(chan string)
	catChannel := make(chan string)
	foxChannel := make(chan string)

	go getDogImage("https://dog.ceo/api/breeds/image/random", dogChannel)
	go getCatImage("https://aws.random.cat/meow", catChannel)
	go getFoxImage("https://randomfox.ca/floof/", foxChannel)

	data := animalsPageData{
		Dog: <-dogChannel,
		Cat: <-catChannel,
		Fox: <-foxChannel,
	}

	content, err := core.Templates.FindString("animals.html")

	if err != nil {

		core.Logger.Error(err)

		http.Error(w, "", http.StatusInternalServerError)

	} else {

		t, err := template.New("html").Parse(content)

		if err != nil {

			core.Logger.Error(err)

			http.Error(w, "", http.StatusInternalServerError)

		} else {
			t.Execute(w, data)
		}

	}

}

func getDogImage(url string, dogChannel chan string) {

	defer close(dogChannel)

	type dogResponseType struct {
		URL string `json:"message"`
	}

	dogResponse := dogResponseType{}

	err := getJSON(url, &dogResponse)

	if err != nil {

		core.Logger.Error(err)

		dogChannel <- ""

	} else {
		dogChannel <- dogResponse.URL
	}

}

func getCatImage(url string, catChannel chan string) {

	defer close(catChannel)

	type catResponseType struct {
		URL string `json:"file"`
	}

	catResponse := catResponseType{}

	err := getJSON(url, &catResponse)

	if err != nil {

		core.Logger.Error(err)

		catChannel <- ""

	} else {
		catChannel <- catResponse.URL
	}

}

func getFoxImage(url string, foxChannel chan string) {

	defer close(foxChannel)

	type foxResponseType struct {
		URL string `json:"image"`
	}

	foxResponse := foxResponseType{}

	err := getJSON(url, &foxResponse)

	if err != nil {

		core.Logger.Error(err)

		foxChannel <- ""

	} else {
		foxChannel <- foxResponse.URL
	}
}

func getJSON(url string, target interface{}) error {

	r, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)

}

//GetAnimalController builds and returns the animalHandler
func GetAnimalController() core.Controller {

	controller := core.Controller{
		Name:    "animals",
		Path:    "/animals",
		Handler: animalsHandler,
	}

	return controller

}
