package mucupa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	bodyString := string(body[:])
	fmt.Println(bodyString)

	err2 := json.Unmarshal(body, &j)

	if err2 == nil {
		fmt.Println("Unmarshaled JSON:")
		fmt.Printf("%v", j)
	}

	// testing output
	CryptsyTestLog(bodyString)

}

func CryptsyTestLog(textToLog string) {
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(textToLog)
}
