package load

import (
	"crypto/md5"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func Init(minsec int, maxsec int, duration int, errLog *Logger) {
	for {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		//We sleep for a range between min and max.
		val := time.Duration(r.Intn(maxsec-minsec)+minsec) * time.Second
		exhaustForNseconds(val)
		errLog.Printf("CPU stress cycle for %d seconds finished.\n", val)

	}
}
func exhaustForNseconds(duration time.Duration) {
	s := []byte("load")
	runtime.GOMAXPROCS(runtime.NumCPU() * 20)
	quit := make(chan struct{})
	time.AfterFunc(duration*time.Second, func() { close(quit) })
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
