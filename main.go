package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	jsonpointer "github.com/mattn/go-jsonpointer"
)

type request struct {
	Ackey  string
	Params []string
}

type response struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

// toJSON formats search of result
func (r *request) toJSON(b []byte, keyword string) ([]byte, error) {
	var respJSON interface{}

	err := json.Unmarshal(b, &respJSON)
	if err != nil {
		return nil, err
	}
	count, err := jsonpointer.Get(respJSON, "/response/result/numFound")
	if err != nil {
		return nil, err
	}
	count, err = strconv.Atoi(count.(string))
	if err != nil {
		return nil, err
	}
	result, err := json.Marshal(response{keyword, count.(int)})
	if err != nil {
		return nil, err
	}

	return result, nil
}

// search articles from API
func (r *request) search(keyword string) (*http.Response, error) {
	values := url.Values{}
	values.Add("ackey", r.Ackey)
	values.Add("q", "Body:"+keyword)
	values.Add("wt", "json")
	const endpoint = "http://54.92.123.84/search?"

	resp, err := http.Get(endpoint + values.Encode())
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetPopularity returns most popular article
func (r *request) GetPopularity() (string, error) {
	var result []byte
	var results [][]byte
	for _, v := range r.Params {
		resp, err := r.search(v)
		defer resp.Body.Close()

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		result, err = r.toJSON(b, v)
		if err != nil {
			return "", err
		}
		results = append(results, result)
	}

	return string(results[0]), nil
}

func main() {
	// Get ACKEY from environmental variable
	godotenv.Load()
	ackey := os.Getenv("ACKEY")

	req := request{ackey, os.Args[1:]}
	result, err := req.GetPopularity()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)
}
