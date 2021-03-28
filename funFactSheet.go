package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Fund struct {
	Proj_id         string `json:"proj_id"`
	ProjectAbbrName string `json:"proj_abbr_name"`
	ProjectNameEn   string `json:"proj_name_en"`
	ProjectNameTh   string `json:"proj_name_th"`
	UniqueId        string `json:"unique_id"`
	FundStatus      string `json:"fund_status"`
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/getFundFactSheet", getFundFactSheet)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func getFundFactSheet(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:8000/fundFactSheet/fundByStatus?status=RG"
	fmt.Fprintf(w, "Welcome to the getFundFactSheet!")
	fmt.Println("Endpoint Hit: getFundFactSheet")
	spaceClient := http.Client{
		// Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")
	// req.Header.Set("Ocp-Apim-Subscription-Key", "3a5189136390449dade9d8cadca3d0a5")
	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Println(res.Status)
	// fmt.Println(string(body))
	// fmt.Fprintf(w, "%+v", string(body))
	var responseObject []Fund
	err = json.Unmarshal((body), &responseObject)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(body)
	// fmt.Fprintf(w, responseObject[0].projectId)
	// fmt.Println(responseObject)
}

func main() {
	handleRequest()
	fmt.Print("Hi Go Lang")
}
