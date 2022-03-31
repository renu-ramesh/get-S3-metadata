package contenttype

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"strings"
// )

// func contenttype_test() {

// 	url := "https://3axjp2rzq5.execute-api.us-east-1.amazonaws.com/Dev/v1"
// 	method := "POST"

// 	payload := strings.NewReader(`{
//   "payload": "ewogICAgInNvdXJjZSI6ewogICAgICAgICJwYXRoIjoidGVzdC1yZXBvL2FnYWluLXRlc3QvbG9nbyB0ZXN0IGhvcml6IGJpcy5wbmciLAogICAgICAgICJidWNrZXQiOiJldmVhcmF0ZXN0YnVja2V0IiwKICAgICAgICAicmVnaW9uIjoidXMtZWFzdC0xIgogICAgfSwKICAgICJjcmVkZW50aWFscyI6ewogICAgfQp9"
// }`)

// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, url, payload)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	req.Header.Add("Content-Type", "application/json")

// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(string(body))
// }
