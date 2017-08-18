package main

import (
	"fmt"

	queue "tools/lib/queue"
)

func main() {
	q := queue.New()
	fmt.Printf("len:%v", q.Len())
}
