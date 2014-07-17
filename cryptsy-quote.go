package mucupa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CryptsyGetMarket(n string) {
	url := "http://pubapi.cryptsy.com/api.php?method=singlemarketdata&marketid=" + n
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("1 ERror in CryptsyGetMarket")
		fmt.Println(err)
	} else {
		fmt.Println("our request to Cryptsy:")
		fmt.Printf("%v", req)
		fmt.Println()
	}
	cli := &http.Client{}
	resp, err := cli.Do(req)

	if err != nil {
		fmt.Println("ERror in CryptsyGetMarket")
		fmt.Println(err)
	} else {
		fmt.Println("Our response from Cryptsy:")
		fmt.Printf("%v", resp)
		fmt.Println()
	}

	body, err := ioutil.ReadAll(resp.Body)

	var j interface{}
	fmt.Println("Body from Cryptsy:")
	fmt.Println(body)

	fmt.Println("Body from Cryptsy (cast as string):")
	fmt.Println(string(body[:]))

	err2 := json.Unmarshal(body, &j)

	if err2 == nil {
		fmt.Println("Unmarshaled JSON:")
		fmt.Printf("%v", j)
	}

	// testing output

}
