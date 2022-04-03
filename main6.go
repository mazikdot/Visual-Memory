package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "strings"
)

var (
	no_of_frames int
	no_of_pages  int
	frames       []int
	pages        []int
	counter      int
	time         []int
	flag1        int
	flag2        int
	i            int
	j            int
	pos          int
	faults       int
)

func initialized() {
	frames = make([]int, 3)
	pages = make([]int, 10)
	time = make([]int, 10)
	counter = 0
	faults = 0
	no_of_frames = 3
	no_of_pages = 10
}

func findLRU(time []int, n int) int {
	var minimum int
	pos = 0
	minimum = time[0]
	for i := 1; i < n; i++ {
		if time[i] < minimum {
			minimum = time[i]
			pos = i
		}
	}
	return pos
}
func main() {
	initialized()
	fmt.Printf("-----------------------------\n")
	fmt.Printf("Page Request -> ")
	for i := 0; i < no_of_pages; i++ {
		fmt.Scanf("%d", &pages[i])
	}
	for i := 0; i < no_of_frames; i++ {
		frames[i] = -1
	}
	for i := 0; i < no_of_pages; i++ {
		flag1 = 0
		flag2 = 0
		for j := 0; j < no_of_frames; j++ {
			if frames[j] == pages[i] {
				counter++
				time[j] = counter
				flag1 = 1
				flag2 = 1
				break
			}
		}
		if flag1 == 0 {
			for j := 0; j < no_of_frames; j++ {
				if frames[j] == -1 {
					counter++
					faults++
					frames[j] = pages[i]
					time[j] = counter
					flag2 = 1
					break
				}
			}
		}
		if flag2 == 0 {
			pos = findLRU(time, no_of_frames)
			counter++
			faults++
			frames[pos] = pages[i]
			time[pos] = counter
		}
		fmt.Printf("\n")
		for j := 0; j < no_of_frames; j++ {
			if frames[j] == -1 {
				fmt.Printf("-\t")
			} else {
				fmt.Printf("%d\t", frames[j])
			}

		}

	}
	fmt.Printf("\nTotal Page Faults = %d\t", faults)
	fmt.Printf("\n-----------------------------")
}
