package main

import (
	"fmt"
	"time"
)

func main() {
	// filename := "log.txt"

	var lastMod time.Time

	fmt.Println(lastMod)
	go watch()

	// watcher(filename, lastMod)

	// content, modLast, err := modFile(filename, time.Time{})
	// if err != nil {
	// 	fmt.Printf("%v", err)
	// }

	// fmt.Println(modLast)
	// fmt.Println(string(content))

	// fmt.Println(time.Second)

	// ============================================

	// tick := time.Tick(100 * time.Millisecond)
	// boom := time.After(500 * time.Millisecond)
	// for {
	// 	select {
	// 	case <-tick:
	// 		fmt.Println("tick.")
	// 	case <-boom:
	// 		fmt.Println("BOOM!")
	// 		return
	// 	default:
	// 		fmt.Println("    .")
	// 		time.Sleep(50 * time.Millisecond)
	// 	}
	// }
}
