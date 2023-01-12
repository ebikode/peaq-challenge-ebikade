package main

// import (
// 	"fmt"
// 	"math/big"
// )

import (
	"fmt"
	"os"

	"github.com/irediaes/peaq-challenge/challenge-3/exchange/pkg/server"
)

func main() {
	err := server.InitServer()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

}
