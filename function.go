package covid

import (
	"encoding/json"
	"fmt"
	"github.com/afritzler/covid-skill/pkg/types"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	covidAPIAll     = "https://corona.lmao.ninja/all"
	covidAPICountry = "https://corona.lmao.ninja/countries"
)

// Cases returns the COVID-19 cases for a given country name.
func Cases(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["country"]
	url := getenv("COVID_API_ALL", covidAPIAll)
	var replies []interface{}
	country := false

	if ok && len(keys[0]) >= 1 {
		url = fmt.Sprintf("%s/%s", covidAPICountry, keys[0])
		country = true
	}
	log.Printf("setting url to: %s", url)


	res, err := http.Get(url)
	if err != nil {
		log.Printf("failed to get url %s: %v", url, err)
		replies = append(replies, generateTextMessage(types.RequestErrorMessage, 0))
	}

	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Printf("failed to read body: %v", err)
		replies = append(replies, generateTextMessage(types.RequestErrorMessage, 0))
	}

	response := types.Response{}
	countryResponse := types.CountryResponse{}

	summary := ""
	if country {
		err = json.Unmarshal([]byte(body), &countryResponse)
		summary = fmt.Sprintf("COVID-19 cases for %s: cases %d, deaths: %d, recovered: %d, critival %s, cases today %d, deaths today %d", countryResponse.Country, countryResponse.Cases, countryResponse.Deaths, countryResponse.Recovered, countryResponse.Critical, countryResponse.TodayCases, countryResponse.TodayDeaths)
	} else {
		err = json.Unmarshal([]byte(body), &response)
		summary = fmt.Sprintf("World wide COVID-19 cases %d, deaths: %d, recovered: %d", response.Cases, response.Deaths, response.Recovered)
	}
	if err != nil {
		log.Printf("failed to unmarshal body: %v", err)
		replies = append(replies, generateTextMessage(types.RequestErrorMessage, 0))
	}

	replies = append(replies, generateTextMessage(summary, 0))

	output, err := json.Marshal(types.Replies{Replies: replies})
	if err != nil {
		log.Printf("failed to marshal replies: %+v\n", err)
		panic("something went wrong here with marshalling")
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func generateTextMessage(text string, delay int) types.TextMessage {
	return types.TextMessage{
		Type:    types.ButtonsType,
		Content: text,
		Delay:   delay,
	}
}
