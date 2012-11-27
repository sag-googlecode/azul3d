package main

import "fmt"

var ch = make(chan [][]interface{}, 1)

func main() {
    frame := [][]interface{}{}

    frame = append(frame, []interface{}{"call1", 1, 2})
    frame = append(frame, []interface{}{"call2", 1, 2})
    frame = append(frame, []interface{}{"zomg", 1, 2})

    ch <- frame
    //ch <- []interface{}{1}
    fmt.Println(<-ch)
}
