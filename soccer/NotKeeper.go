package soccer

import (
	"sync"
	"time"
)

type NotKeeper struct {
	Player
	nPass int
}

// I dunno the soccer term for player that is not a keeper so... yea
// NotKeeper has 50% chance of getting the ball
func (nk *NotKeeper) Catch(wg *sync.WaitGroup, ballChan chan Ball, stats *Stats) {
	for ball := range ballChan {
		succeed := (time.Now().UnixNano() % 100) < 50
		if ball.lastKick != nil {
			if succeed {
				stats.makeEntry("catch", ball.lastKick, &nk.Player)
				ball.lastKick = &nk.Player
				nk.nPass++
			} else {
				stats.makeEntry("miss", ball.lastKick, &nk.Player)
			}
		}

		if stats.goals < 5 {
			ballChan <- ball
		} else {
			close(ballChan)
			wg.Done()
		}
	}

}
