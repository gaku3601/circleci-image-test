package function

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Handle a serverless request
func Handle(req []byte) string {
	values := url.Values{}
	values.Add("myname2", "gakuo")

	reqq, err := http.NewRequest(
		"POST",
		"http://link:8080/",
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		return err.Error()
	}

	// Content-Type 設定
	reqq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(reqq)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)

	return fmt.Sprintf("Hello, Go. You said: %s", string(byteArray))
}
