package controllers

import (
	"encoding/json"
	"fmt"
	"jaabo/models"
	"net/http"
)

func JobsIndexHandler(rw http.ResponseWriter, r *http.Request) {
	//job := models.Job{"Awesome PHP Developer"}
	dataStore, err := models.NewDataStore()
	if err != nil {
		// TODO: ideal case log error and return a 404
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dataStore.Close()

	switch r.Method {
	case "GET":
		jobs := dataStore.FindJobs()
		js, err := json.Marshal(jobs)

		if err != nil {
			// TODO: ideal case log error and return a 404
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.Write(js)

	case "POST":
		decoder := json.NewDecoder(r.Body)
		fmt.Printf("%+v\n", r.Body)
		var jobRequest models.Job
		err := decoder.Decode(&jobRequest)
		if err != nil {
			panic(err)
		}

		err = dataStore.UpdateJob(&jobRequest)
		if err != nil {
			http.Error(rw, "Job Couldn't be created", 500)
		}

	case "PUT":
		decoder := json.NewDecoder(r.Body)
		fmt.Printf("%+v\n", r.Body)
		var jobRequest models.Job
		err := decoder.Decode(&jobRequest)
		if err != nil {
			panic(err)
		}

		err = dataStore.CreateJob(&jobRequest)
		if err != nil {
			http.Error(rw, "Job Couldn't be created", 500)
		}

	default:
		http.Error(rw, "Method not allowed", 405)
	}

}

func JobShowHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Single Job")
}
