package channels

import (
	"fmt"
	"math/rand"
	"time"
)

/* Notes on Go Concurrency */
// - In Go, concurrency is achieved with the `go` keyword.
//		ex. go someFunc()
// - The rest of the code will proceed. The `go` keyword spawns a goroutine

/* Notes on Go Channels */
// - created using make
//		ex. ch := make(chan int)

func sendToChannel(ch chan string) {
	ch <- "go boot.dev!"
}

func sendToChannelWithDelay(secs int, ch chan string) {
	time.Sleep(time.Duration(secs) * time.Second)
	ch <- "go boot.dev!"
}

func recieveFromChannel(ch chan string) {
	// Reads and removes a value from the channel (first in first out)
	// Will block until a value from the channel is received
	channelReceiver := <-ch
	fmt.Println("channelReceiver: ", channelReceiver)
}

func basicChannelUsage() {
	// Declare a channel
	someChannel := make(chan string)

	// I think I have to use `go` when working with channels.
	//		Not doing so appears to crash my program.
	// Correction! I would only need to use `go` if the channel is unbuffered.
	// Unbuffered channels will block until the value has been retrieved.

	go sendToChannel(someChannel)
	recieveFromChannel(someChannel)
}

func usingChannelsAsCueToStopStart() {
	someChannel := make(chan string)

	go sendToChannelWithDelay(1, someChannel)

	<-someChannel
	fmt.Println("this part only triggers once info is retrieved from channel")
}

func bufferedChannel(startingPokemon []string) {
	// Buffered channels have a predefined number of slots (for values)
	// Buffered channels won't block unless it fills the set buffer amount
	//		So, no need for `go` as long as it hasn't yet hit the cap

	// Usefull for when we want to act on the info in batches
	pokemonCh := make(chan string, len(startingPokemon))
	for _, pkmn := range startingPokemon {
		pokemonCh <- pkmn
	}

	// Offloading from a buffered channel has to happen one step at a time.
	pkmStartList := []string{}
	for len(pokemonCh) > 0 {
		pkmn := <-pokemonCh
		pkmStartList = append(pkmStartList, pkmn)
	}

	randomIdx := rand.Intn(len(pkmStartList))

	fmt.Println("Chosen start pokemon: ", pkmStartList[randomIdx])
}

func closingChannels() {
	// I'm sensing that it would be good practice to close channels once done
	// This is even though the garbage collector will handle it eventually.

	chToClose := make(chan string)
	close(chToClose)

	// It's good practice to check
	if _, chOk := <-chToClose; !chOk {
		fmt.Println("The channel is already closed")
	}

	// Sending to a closed channel causes a panic
}

/* Use Select to listing to multiple channels */
// I didn't come up with a personal example, but it looks like the following
/*
	select {
	case i, ok := <-chInts:
		if ok {
			fmt.Println(i)
		}
	case s, ok := <-chStrings:
		if ok {
			fmt.Println(s)
		}
	case <-chIgnore:
		For when the value is not used, but we may want to trigger something
	default:
		Execute immediately when available channel to read from
	}
*/

/* Read and Write only */
// - There's also a way to specify channels as READ-only or WRITE-only.
// 		Didn't include this here

func ChannelsMain() {
	fmt.Println("===== 13_channels =====")

	basicChannelUsage()
	usingChannelsAsCueToStopStart()

	startingPokemon := []string{"Bulbasaur", "Charmander", "Squirtle"}
	bufferedChannel(startingPokemon)

	closingChannels()
}
