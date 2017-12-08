package main

import (
	"fmt"
	"net/http"
	"os"
	log "github.com/sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
	"app/models"
	"io/ioutil"
	
	"encoding/json"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	if os.Getenv("ENVIRONMENT") == "Production" {
		log.SetLevel(log.WarnLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

func Option(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func GetConfiguration() models.Configuration{

	file, e := ioutil.ReadFile("/app/Configs/config.json")
    if e != nil {
        fmt.Printf("Read config file error: %v\n", e)
        os.Exit(1)
	}

	var config models.Configuration
	err := json.Unmarshal(file, &config)
	
	if err != nil {
        fmt.Printf("Problem in config file: %v\n", err.Error())
        os.Exit(1)
	}

	return config
}
func main() {
	
	config := GetConfiguration()
	
	fmt.Println(config)
	router := httprouter.New()

	fmt.Println("enabling alert api route...")
	
	router.POST("/alert", func(w http.ResponseWriter, r *http.Request, p httprouter.Params){
		Alert(w,r,p, config)
	})
	

	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":10010", router))

}
