package function

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Handle a serverless request
func Handle(req []byte) string {
	url := os.Getenv("URL")

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return fmt.Sprintf("Hello, Go. You said: %s", string(byteArray))
}
