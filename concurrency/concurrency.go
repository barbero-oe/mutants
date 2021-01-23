package concurrency

import "sync"

// Merge combines output channels to a single one
func Merge(channels []<-chan struct{}) <-chan struct{} {
	var wg sync.WaitGroup
	wg.Add(len(channels))
	out := make(chan struct{})
	consume := func(ch <-chan struct{}) {
		defer wg.Done()
		for range ch {
			out <- struct{}{}
		}
	}
	for _, ch := range channels {
		go consume(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
