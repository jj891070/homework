package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	// week2_1()
	// week2_2()
	// week3()
	week4()

}

func week2_1() {
	var people [5]int64
	var total int64
	total = 0
	for key := range people {
		people[key] = rand.Int63n(100)
		fmt.Printf("people %d : %d \n", key, people[key])
		total = total + people[key]
	}

	total = total / int64(len(people))
	fmt.Printf("平均成績 : %d \n", total)
}

func week2_2() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	small := x[0]
	for _, value := range x {
		if small > value {
			small = value
		}
	}
	fmt.Printf("The small : %d \n", small)

}

func week3() {
	Q1()
	Q2()

}

func week4() {
	ch := make(chan int)
	mx := new(sync.RWMutex)
	store := 0
	go func() {

		for {
			// fmt.Print("5000")
			mx.Lock()
			if store < 60000 {
				store = store + 5000
			} else {
				ch <- 5
			}
			mx.Unlock()
		}
	}()

	go func() {
		for {
			// fmt.Print("3000")
			mx.Lock()
			if store < 60000 {
				store = store + 3000
			} else {
				ch <- 3
			}
			mx.Unlock()
		}
	}()

	go func() {
		for {
			// fmt.Print("6000")
			mx.Lock()
			if store < 60000 {
				store = store + 6000
			} else {
				ch <- 6
			}
			mx.Unlock()
		}
	}()

	select {
	case msg1 := <-ch:
		fmt.Println("received done ", msg1)
		fmt.Println("total ", store)

	}

}
