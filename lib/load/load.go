package load

import (
	"crypto/md5"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func Init(minsec int, maxsec int, duration int, errLog *log.Logger) {
	errLog.Println("load starting")
	for {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		//We sleep for a range between min and max.
		time.Sleep(time.Duration(r.Intn(maxsec-minsec)+minsec) * time.Second)
		exhaustForNseconds(time.Duration(duration) * time.Second)
		errLog.Printf("CPU stress cycle for %d seconds finished.\n", duration)
	}
}

func exhaustForNseconds(duration time.Duration) {
	s := []byte("load")
	runtime.GOMAXPROCS(runtime.NumCPU() * 20)
	quit := make(chan struct{})
	time.AfterFunc(duration, func() { close(quit) })
	var wg sync.WaitGroup
	for i := 0; i <= runtime.NumCPU()*20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-quit:
					return
				default:
					md5.Sum(s)
				}
			}
		}()
	}
	wg.Wait()
}
