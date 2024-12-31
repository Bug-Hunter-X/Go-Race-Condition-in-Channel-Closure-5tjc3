# Go Race Condition in Channel Closure

This repository demonstrates a potential race condition in Go programs that involve channels and goroutines.  The `bug.go` file showcases a scenario where the channel closure might not be synchronized with the loop that reads from the channel, leading to unexpected behavior or program crashes.

The solution (`bugSolution.go`) fixes the race condition by employing a WaitGroup to properly synchronize the goroutine's completion with the main loop's consumption of channel values. This ensures the channel is correctly closed and the loop terminates gracefully.  

The key takeaway is to always carefully synchronize goroutine termination with the consumption of data from channels to avoid such race conditions.