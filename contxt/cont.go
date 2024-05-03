package contxt

import (
	"context"
	"fmt"
	"log"
	"time"
)

func Exec() {
	fmt.Println("Context Example")
	start := time.Now()
	ctx := context.Background()
	val, err := callAPI(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(val)
	fmt.Println("Time taken: ", time.Since(start))
}

type Response struct {
	value int
	err   error
}

func callAPI(ctx context.Context) (int, error) {
	// Get context with timeout. Ensures that the API call does not take more than 3 seconds
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	respch := make(chan Response)

	go func() {

		// Schedule in a goroutine
		val, err := someThirdPartyAPI()
		respch <- Response{val, err}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("API call timed out")
		case resp := <-respch:
			return resp.value, resp.err
		}

	}
}

func someThirdPartyAPI() (int, error) {
	time.Sleep(5 * time.Second)
	return 2828, nil
}
