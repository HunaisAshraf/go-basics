package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// apiCall()
	// verifyUrl()
	postApi()
}

func postApi() {
	link := "https://jsonplaceholder.typicode.com/posts"

	body := strings.NewReader(`
		{
			"title":"new world",
			"body": "the world full of dangers",
			"userId":56	
		}
	`)

	res, err := http.Post(link, "application/json", body)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	responseString, err := io.ReadAll(res.Body)

	fmt.Println(string(responseString))

}

func apiCall() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		panic(err)
	}
	fmt.Printf("type of response %T\n", res)

	var response strings.Builder
	data, err := io.ReadAll(res.Body)

	byteLen, _ := response.Write(data)
	fmt.Println(byteLen)
	fmt.Println(response.String())

	finalData := string(data)

	fmt.Println(finalData)
	fmt.Printf("type of data: %T\n", finalData)
	res.Body.Close()
}

func verifyUrl() {
	link := "https://jsonplaceholder.typicode.com/todos?data=hello&data=world&num=10"

	data, _ := url.Parse(link)
	fmt.Println(data.Scheme)
	fmt.Println(data.Host)
	fmt.Println(data.Path)
	fmt.Println(data.RawQuery)
	fmt.Println(data.Hostname())

	val := data.Query()

	fmt.Println(val)

	for key, val := range val {
		fmt.Printf("key: %v, value: %v, type of value: %T \n", key, val, val)
		if len(val) > 1 {
			for _, k := range val {
				fmt.Println(k)
			}
		}
	}

}
