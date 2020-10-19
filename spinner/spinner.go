package spinner

import (
	"fmt"
	"time"
)
//Movespinner Экспортируемая функция в main.
func Movespinner() {
	go spinner(50 * time.Millisecond)
	tick := 1
	for range time.Tick(1 * time.Second) {
		if tick > 5 {
			break
		}
		fmt.Println(tick, "seconds passed")
		tick++
	}
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}