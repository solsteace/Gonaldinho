package soccer

import (
	"sync"
	"time"
)

type Keeper struct {
	Player
	saves int
}

// Keeper has 80% chance of getting the ball from opponents and 100% from friendlies
func (k *Keeper) Catch(wg *sync.WaitGroup, ballChan chan Ball, stats *Stats) {
	for ball := range ballChan {
		if ball.lastKick != nil {
			var catchRate int64 = 80
			if ball.lastKick.Team == k.Player.Team {
				catchRate += 20
			}

			succeed := (time.Now().UnixNano() % 100) < catchRate
			if succeed {
				stats.makeEntry("save", ball.lastKick, &k.Player)
				k.saves++
			} else {
				stats.makeEntry("goal", ball.lastKick, &k.Player)
				stats.goals++
			}
		}

		ball.lastKick = &k.Player
		if stats.goals < 5 {
			ballChan <- ball
		} else {
			close(ballChan)
			wg.Done()
		}
	}
}
