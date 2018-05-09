package function

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// Handle a serverless request
func Handle(req []byte) string {
	myurl := os.Getenv("URL")

	values := url.Values{}
	values.Add("myname", "gaku")
	values.Add("myname2", "gakuo")
	resp, _ := http.PostForm(
		myurl,
		values,
	)

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	println(string(body))
	return fmt.Sprintf("Hello, Go. You said: %s", string(body))
}
