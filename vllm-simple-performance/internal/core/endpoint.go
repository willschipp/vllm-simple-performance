package core

import (
	"log"
	"net/http"
	"sync"
	"encoding/json"
	"io"
	"time"
	"bytes"
	"fmt"
	"os"
)


// target structure
// {
//     "model": [model],
//     "prompt": [prompt],
// }
type Payload struct {
	Model string `json:"model"`
	Prompt string `json:"prompt"`
}

func SendPrompt(url string, model string, prompt string,wg *sync.WaitGroup) {
	defer wg.Done()

	payload := Payload{
		Model: model,
		Prompt: prompt,
	}

	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("error unmarshalling prompt %v",err)
	}
	//timer
	start := time.Now()

	response, err := http.Post(url,"application/json",bytes.NewReader(jsonBytes))

	// response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error invoking: %v",err)
	} //end if
	defer response.Body.Close()

	//timer
	elapsed := time.Since(start)

	log.Println("Status Code:", response.StatusCode)
	log.Printf("Response time: %s\n",elapsed)
	//get the response
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("error reading body %v",err)
	} //end if

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(body,&jsonResponse); err != nil {
		// log.Fatalf("error parsing body %v",err)
		log.Printf("error parsing body %v",err)
	} //end if
	log.Printf("Response %+v\n",jsonResponse) //dump the map
}

func GetMetrics(url string,location string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Can't get metrics %v",err)
	}
	defer response.Body.Close()
	// get the timestamp
	timestamp := time.Now().UTC().Format("2006-01-01-15:01:01")
	filename := fmt.Sprintf("metrics_%s.txt",timestamp)
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("couldn't write the filename %v",err)
	}
	defer file.Close()

	//write out the response
	_, err = io.Copy(file,response.Body)
	if err != nil {
		log.Fatalf("couldn't write to the file %v",err)
	}
	//done	
}