package leaky_bucket

import (
	"errors"
	"log"
	"strconv"
	"time"

	u "rate-limiter/util"
)

// configuring 5 req/s as the processing rate
const MAX_REQ uint32 = 5

func Init() chan *u.PostRequest {
	ch := make(chan *u.PostRequest, 2*MAX_REQ)

	go func() {
		ticker := time.NewTicker(time.Duration(1e9 / float64(MAX_REQ)))
		log.Println("Initialized listener thread")

		for range ticker.C {
			pr := <-ch
			log.Println("[LIMITER] Processing request:", pr.Id, pr.Origin)
		}
	}()

	return ch
}

func ProcessWithLimit(pr *u.PostRequest, ch chan *u.PostRequest) error {
	// using twice the configured req/s for handling surge in requests
	if uint32(len(ch)) >= 2*MAX_REQ {
		return errors.New("Rate limiting " + strconv.Itoa(int(pr.Id)) + " " + pr.Origin)
	}

	ch <- pr
	return nil
}
