package main

import (
	"fmt"
)


func fibonacci(position uint, memoized *map[uint]uint) uint {
  if position <= 1 {
    return position
  }

  valueForKey, hasValueForKey := (*memoized)[position]
  fmt.Println((*memoized))
  
  if hasValueForKey {
    return valueForKey
  }

  (*memoized)[position] = fibonacci(position - 2, memoized) + fibonacci(position - 1, memoized)
    return (*memoized)[position]
}

func main() {
  memoized := make(map[uint]uint)
	response := fibonacci(uint(10), &memoized)
  fmt.Println(response)

}
