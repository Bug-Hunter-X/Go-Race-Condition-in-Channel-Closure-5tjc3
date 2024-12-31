```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        ch := make(chan int)
        done := make(chan struct{})

        wg.Add(1)
        go func() {
                defer wg.Done()
                for i := 0; i < 5; i++ {
                        ch <- i
                }
                close(ch)
                close(done)
        }()

        wg.Wait()
        <-done

        for v := range ch {
                fmt.Println(v)
        }
}
```