package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var ch = make(chan int)
var wg = sync.WaitGroup{}

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		ch <- 20
	}()
	wg.Wait()
	calling(ctx)

}

func calling(ctx context.Context) {

	select {
	case <-ch:
		fmt.Println("Recieved : ", ch)

	case <-ctx.Done():
		fmt.Println("time out")
		return

	}

}
