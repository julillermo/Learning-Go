package mutexes

import (
	"fmt"
)

/* Note Mutex */
// - Mutex comes from the "sync" package
// - I think this is for prevent race conditions amongst goroutines

// I didn't cover mutexes much, but the generall idea is that only 1
// 	gorouine at a time can use run a locked mutex until it gets unlocked

func MutexMain() {
	fmt.Println("===== 14_mutexes =====")

}
