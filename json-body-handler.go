package mucupa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func jsonBodyUnmarshal(r *http.Request) interface{} {

	bodyBytes, err := ioutil.ReadAll(r.Body)
	var j interface{}

	if err != nil {
		fmt.Println("body error")
		fmt.Println(err)
	} else {
		fmt.Println("bodyBytes: ", bodyBytes)
		bodyString := string(bodyBytes[:])
		fmt.Println("bodyString: ", bodyString)
		err := json.Unmarshal(bodyBytes, &j)
		if err == nil {
			// do stuff
		}

	}
	return j

}
