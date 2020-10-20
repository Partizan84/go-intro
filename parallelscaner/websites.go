package parallelscaner

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"sync"
)
//ParallelScaner параллельный сканер.
func ParallelScaner() {
	wg := &sync.WaitGroup{}

	start := time.Now()
	for _, url := range os.Args[1:] {
		wg.Add(1)
		go fetch(url, wg)
	}

	wg.Wait()

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	//bytes, err := ioutil.ReadAll(resp.Body)
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // to devNull
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2fs %7d %s\n", time.Since(start).Seconds(), nbytes, url)
}