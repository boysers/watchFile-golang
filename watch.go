package main

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

func watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
			case err, ok := <-watcher.Errors:
				fmt.Println(ok)
				if !ok {
					return
				}
				log.Println("error", err)
			}
		}
	}()

	err = watcher.Add("./log.txt")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

// func modFile(filename string, lastFileMod time.Time) ([]byte, time.Time, error) {
// 	info, err := os.Stat(filename)
// 	if err != nil {
// 		return nil, lastFileMod, err
// 	}
// 	if !info.ModTime().After(lastFileMod) {
// 		return nil, lastFileMod, nil
// 	}

// 	content, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return nil, info.ModTime(), err
// 	}
// 	return content, info.ModTime(), nil
// }

// func watcher(filename string, lastMod time.Time) {
// 	lastError := ""
// 	fileTicker := time.NewTicker(1 * time.Second)
// 	defer func() {
// 		fmt.Println("stop")
// 		fileTicker.Stop()
// 	}()
// 	for {
// 		select {
// 		case <-fileTicker.C:
// 			var content []byte
// 			var err error

// 			content, lastMod, err := modFile(filename, lastMod)
// 			if err != nil {
// 				if e := err.Error(); e != lastError {
// 					lastError = e
// 					content = []byte(lastError)
// 				}
// 			} else {
// 				lastError = ""
// 			}

// 			if content != nil {
// 				// fmt.Println(string(content))
// 				fmt.Println(lastMod)
// 			}
// 		}
// 	}
// }

// // ---------------

// var filename = "log.txt"

// func readFileIfModified(lastMod time.Time) ([]byte, time.Time, error) {
// 	fi, err := os.Stat(filename)
// 	if err != nil {
// 		return nil, lastMod, err
// 	}
// 	if !fi.ModTime().After(lastMod) {
// 		return nil, lastMod, nil
// 	}
// 	p, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return nil, fi.ModTime(), err
// 	}
// 	return p, fi.ModTime(), nil
// }
