package main

import (
	"errors"
	"fmt"
)

func main() {
	msg := sayHello("Aran")
	fmt.Println(msg)
}

func sayHello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

type Item struct {
	id    int
	name  string
	price float64
}

func removeFromListById(items []Item, id int) ([]Item, error) {
	for i, item := range items {
		if item.id == id {
			return append(items[:i], items[i+1:]...), nil
		}
	}
	return items, errors.New("item not found")
}
