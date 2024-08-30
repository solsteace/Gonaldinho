package main

import (
	"sync"

	"github.com/solsteace/gonaldinho/soccer"
)

func main() {
	ballChan := make(chan soccer.Ball)
	stats := soccer.Stats{}

	// Setup players (4 + 1 each team)
	keepers := []soccer.Keeper{
		{Player: soccer.Player{Name: "Jono", Team: "1"}},
		{Player: soccer.Player{Name: "Jeremy", Team: "2"}},
	}
	notKeepers := []soccer.NotKeeper{
		{Player: soccer.Player{Name: "Jajang", Team: "1"}},
		{Player: soccer.Player{Name: "Jamal", Team: "1"}},
		{Player: soccer.Player{Name: "Januar", Team: "1"}},
		{Player: soccer.Player{Name: "Jojon", Team: "1"}},
		{Player: soccer.Player{Name: "Jason", Team: "2"}},
		{Player: soccer.Player{Name: "Jonathan", Team: "2"}},
		{Player: soccer.Player{Name: "Jeremiah", Team: "2"}},
		{Player: soccer.Player{Name: "Jeremy", Team: "2"}},
	}

	var wg sync.WaitGroup
	wg.Add(1)

	// Get ready players!
	for _, notKeeper := range notKeepers {
		go notKeeper.Catch(&wg, ballChan, &stats)
	}
	for _, keeper := range keepers {
		go keeper.Catch(&wg, ballChan, &stats)
	}

	// Kick off! The match is on!
	ballChan <- soccer.Ball{}

	wg.Wait()
	stats.DisplayLog()
	stats.ShowResult()
}
