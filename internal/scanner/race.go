package scanner

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/Dyrandy/bigmom/internal/menus"
)

func fetchURL(url string, wg *sync.WaitGroup, method string, headers map[string]string, body string) {
	defer wg.Done()

	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	for key, value := range headers {
		req.Header.Set(key, value[1:len(value)-1])
		// fmt.Println(value)
		// fmt.Println(strings.ReplaceAll(value, "\t", ""))
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)
}

func DoRaceAttack(base64Packet string, method string, host string, path string, query string, isTLS bool) {
	var wg sync.WaitGroup
	var url string

	base64Plain, err := base64.StdEncoding.DecodeString(base64Packet)
	if err != nil {
		panic(err)
	}
	menus.Logo()
	fmt.Printf("%s", base64Plain)
	fmt.Printf("\n\n > Attack This Packet? (Y/N) : ")
	var choose string
	fmt.Scanln(&choose)
	if strings.ToUpper(choose) == "Y" {
		str := bytes.Split(base64Plain, []byte{10})
		headers := make(map[string]string)
		isBody := false
		bodyData := ""
		for index, part := range str {
			if index != 0 {
				keyAndValue := string(part)
				splitedList := strings.SplitN(keyAndValue, ":", 2)
				// fmt.Println("asd: ", splitedList[0], len(splitedList))
				if len(splitedList) == 1 && isBody == false {
					isBody = true
					continue
				} else if isBody == true {
					bodyData = bodyData + keyAndValue
					// fmt.Println("qweqwe")
					continue
				}
				if len(splitedList) > 1 {
					if strings.ToLower(splitedList[0]) != "host" && strings.ToLower(splitedList[0]) != "connection" && strings.ToLower(splitedList[0]) != "cache-control" && isBody == false && strings.ToLower(splitedList[0]) != "content-length" && strings.ToLower(splitedList[0]) != "upgrade-insecure-requests" {
						headers[splitedList[0]] = splitedList[1]
						// fmt.Println(splitedList[0])
					}
				}

			}
			// fmt.Println("part :", keyAndValue, asd)

		}
		// fmt.Println(bodyData)
		// for key, value := range headers {
		// 	fmt.Println(key, value)
		// }
		// fmt.Println(headers)

		if isTLS == true {
			url = "https://" + host + path + query
		} else {
			url = "http://" + host + path + query
		}

		// fmt.Println(url)
		for i := 0; i < 5; i++ {
			wg.Add(1)
			fetchURL(url, &wg, method, headers, bodyData)
		}
		wg.Wait()
	} else {
		return
	}

}
