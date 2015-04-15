// MultiQuickSortDemo project main.go
package main

import (
	"fmt"
	"sync"
	//	"time"
)

var mutex sync.Mutex

func main() {
	fmt.Println("Hello World!")
	count := 1
	end := make(chan bool, 0)
	testArray := []int{10, 20, 30, 5, 9, 65, 87, 49, 100, 33, 100, 1000, 1, 4, 3, 100, 1, 4, 65, 89, 999, 10, 1001}
	divide(testArray, 0, len(testArray)-1, &count, end, len(testArray))
	//	spread(testArray, lowflag, highflag, 0, 9, pival)
	<-end
	for i := 0; i < len(testArray); i++ {
		fmt.Printf("%d ", testArray[i])
	}
	//	fmt.Print(pival)
	fmt.Println(count)
}

func divide(segment []int, low int, high int, count *int, end chan bool, sum int) (flagl bool, flagr bool, pival int) {
	i := low
	j := high
	temp := segment[i]
	for i < j {
		for i < j && segment[j] > temp {
			j--
		}
		if i < j {
			segment[i] = segment[j]
		}
		for i < j && segment[i] <= temp {
			i++
		}
		if i < j {
			segment[j] = segment[i]
		}
	}
	segment[i] = temp
	//	add(count)
	pival = i
	if (pival - low) < 1 {
		flagl = true
		add(count, end, sum)
	}
	if (high - pival) < 1 {
		flagr = true
		add(count, end, sum)
	}
	spread(segment, flagl, flagr, low, high, pival, count, end, sum)
	return
}
func spread(segment []int, lowflag bool, highflag bool, low int, high int, pival int, count *int, end chan bool, sum int) {
	if !highflag {
		go divide(segment, pival+1, high, count, end, sum)
	}
	if !lowflag {
		go divide(segment, low, pival-1, count, end, sum)
	}
}

func add(count *int, end chan bool, sum int) {
	mutex.Lock()
	defer mutex.Unlock()
	*count = *count + 1
	if *count >= sum {
		end <- true
	}

}
