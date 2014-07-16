package mucupa

import (
	"fmt"
	"net/http"
)

func QuoteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("quote success")
	// read body
	// j := jsonBodyUnmarshal(r)
	// do cryptsy stuff
	CryptsyGetMarket("61")
}
