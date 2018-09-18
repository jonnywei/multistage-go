package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	fmt.Println("Starting server")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		urls, found := request.URL.Query()["url"]
		if !found || len(urls) < 1 {
			urls = []string{"http://sohu.com"}
		}

		response, err := http.Get(urls[0])
		if err != nil {
			fmt.Println(fmt.Errorf("An error occured:%s", err))
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("500-error happened"))
			return
		}

		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		writer.WriteHeader(http.StatusOK)
		writer.Write(body)

	})
	http.ListenAndServe(":8080", nil)
}
