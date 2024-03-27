package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

/*
Context is like a bag or container that holds information that is shared between different parts of the program,
especially when it comes to handling a request. This information can include things like timeouts, cancellation signals,
and other data that is specific to that request.

For example, imagine you are building a web server that handles a lot of incoming requests.
Each request has its own specific needs and requirements, such as a deadline for how long it should take to complete.
The context allows you to keep track of these individual requirements for each request,
and make sure that they are handled properly.
*/

func main() {
	start := time.Now()
	ctx := context.WithValue(context.Background(), "uuid", "1234")
	userId := 1
	val, err := fetchUserData(ctx, userId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Results : %v\n", val)
	fmt.Println("Took", time.Since(start))

}

type Response struct {
	Value int
	Err   error
}

func fetchUserData(ctx context.Context, userId int) (int, error) {
	uuid := ctx.Value("uuid")
	fmt.Println("uuid : ", uuid)
	resch := make(chan Response)

	// set timeout for third party API call to get response
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()

	go func() {
		val, err := fetchThirdPartyStuff()
		resch <- Response{
			Value: val,
			Err:   err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("Timeout")
		case resp := <-resch:
			return resp.Value, resp.Err
		}
	}
}

func fetchThirdPartyStuff() (int, error) {
	// we are not sure about this time, as its 3rd party
	// We can't control this, or Do we ? Yes, By using Context.
	time.Sleep(time.Millisecond * 100)
	return 101, nil
}
