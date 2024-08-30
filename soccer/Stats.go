package soccer

import (
	"fmt"
	"strings"
	"sync"
)

type (
	logEntry struct {
		action   string
		passer   *Player
		receiver *Player
	}

	Stats struct {
		goals   int
		goalLog []*Player
		log     []logEntry
		mu      sync.Mutex
	}
)

func (s *Stats) makeEntry(action string, passer, receiver *Player) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if action == "goal" {
		s.goalLog = append(s.goalLog, passer)
	}

	s.log = append(s.log, logEntry{action, passer, receiver})
}

func (s *Stats) ShowResult() {
	fmt.Println("==== SUMMARY ====")
	fmt.Println("-> Goals")

	team1Score, team2Score := 0, 0
	for _, goal := range s.goalLog {
		fmt.Println(goal)
		if goal.Team == "1" {
			team1Score++
		} else {
			team2Score++
		}
	}

	fmt.Println("==== Final score ====")
	fmt.Printf("(Team `1`) %d : %d (Team `2`)\n", team1Score, team2Score)
}

func (s *Stats) DisplayLog() {
	var out []string
	for _, entry := range s.log {
		var newLine string
		switch entry.action {
		case "catch":
			newLine = fmt.Sprintf(
				"%s received the ball coming from %s!",
				entry.receiver.Name,
				entry.passer.Name)
		case "miss":
			newLine = fmt.Sprintf(
				"%s misses the ball coming from %s!",
				entry.receiver.Name,
				entry.passer.Name)
		case "save":
			newLine = fmt.Sprintf(
				"%s saved the ball coming from %s!",
				entry.receiver.Name,
				entry.passer.Name)
		case "goal":
			newLine = fmt.Sprintf(
				"%s failed to save the ball! A Goal has been scored by %s, that's a score for %s!",
				entry.receiver.Name,
				entry.passer.Name,
				entry.passer.Team)
		}

		out = append(out, newLine)
	}

	fmt.Println(strings.Join(out, "\n"))
}
