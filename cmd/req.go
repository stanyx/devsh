package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

func headerKeysSorted(h http.Header) []string {
	var keys []string
	for k := range h {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func SendGETRequest(r *bufio.Reader) {

	line, _ := r.ReadString('\n')
	url := strings.Trim(line, "\r\n")

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %+v\n", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error: %+v\n", err)
		return
	}
	fmt.Println("Status: ", resp.Status)
	fmt.Println()
	fmt.Println("Headers: ")

	for _, k := range headerKeysSorted(resp.Header) {
		fmt.Printf("%s := %v\n", k, resp.Header[k])
	}
	fmt.Println()
	fmt.Println("Response: ")
	fmt.Println(string(body))
}

// func sendPOSTRequest(url string, body map[string]interface{}) {

// }
