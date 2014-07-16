package mucupa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func QuoteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("quote success")
	// read body
	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println("body error")
		fmt.Println(err)
	} else {
		fmt.Println("bodyBytes: ", bodyBytes)
		bodyString := string(bodyBytes[:])
		fmt.Println("bodyString: ", bodyString)
		var j interface{}
		err := json.Unmarshal(bodyBytes, &j)
		if err == nil {
			// do stuff
		}

		fmt.Printf("%v", j)
	}

	// do cryptsy stuff
}
