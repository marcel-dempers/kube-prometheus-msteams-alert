package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"io/ioutil"
	"app/models"
	"bytes"
)


func Alert(writer http.ResponseWriter, request *http.Request, p httprouter.Params, config models.Configuration) {
	
	jsonBytes, err := ioutil.ReadAll(request.Body)
	
	if err != nil {
		panic(err)
	}

	jsonString := string(jsonBytes)
	
	var alertData models.Alerts

	err = json.Unmarshal([]byte(jsonString), &alertData)
	
	if err != nil {
		panic(err)
	}
	
	for _, alert := range alertData.Alerts {

		var alertBody models.MSTeamsAlert
		alertBody.Title = alert.Labels.Cluster + " : " + alert.Annotations.Summary
		alertBody.Text = alert.Annotations.Description
		alertBody.ThemeColor = "EA4300"
		
		outputJSON, _ := json.Marshal(alertBody)
		var jsonBytes = []byte(outputJSON)

		req, reqErr := http.NewRequest("POST", config.MicrosoftTeamsURL, bytes.NewBuffer(jsonBytes))
		client := &http.Client{}
		resp, reqErr := client.Do(req)
		if reqErr != nil {
			panic(reqErr)
		}
		defer resp.Body.Close()
	
		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))
	}
}