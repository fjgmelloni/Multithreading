package worker

import (
	"fmt"
	"net/http"
)

type Job struct {
	Cep      string
	Response http.ResponseWriter
	Request  *http.Request
}

var JobQueue chan<- Job
var handlerFunc func(string, http.ResponseWriter, *http.Request)

func SetHandler(f func(string, http.ResponseWriter, *http.Request)) {
	handlerFunc = f
}

func StartPool(numWorkers int) {
	jobChan := make(chan Job, 100)
	JobQueue = jobChan

	for i := 0; i < numWorkers; i++ {
		go worker(i, jobChan)
	}
}

func worker(id int, jobs <-chan Job) {
	for job := range jobs {
		fmt.Printf("[Worker %d] Processando CEP %s\n", id, job.Cep)
		handlerFunc(job.Cep, job.Response, job.Request)
	}
}
