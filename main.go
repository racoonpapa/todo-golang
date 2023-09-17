package main

import (
	"fmt"
	"os"
	"time"
	"todo-golang/todo"
)

func main() {
	var bytes []byte
	var err error
	var list *todo.List
	var started time.Time
	var elapsedFrom, elapsedTo float64

	if bytes, err = os.ReadFile("todo.json"); err != nil {
		panic(err)
	}

	started = time.Now()
	if list, err = todo.NewListFromJson(bytes); err != nil {
		panic(err)
	}
	elapsedFrom = time.Now().Sub(started).Seconds()

	started = time.Now()
	if _, err = list.ToJson(); err != nil {
		panic(err)
	}
	elapsedTo = time.Now().Sub(started).Seconds()

	fmt.Printf("fromJson - elapsed: %f seconds, toJson - elapsed: %f seconds\n", elapsedFrom, elapsedTo)
}
