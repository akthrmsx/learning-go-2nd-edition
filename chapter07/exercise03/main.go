package main

import (
	"io"
	"os"
	"sort"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name  string
	Teams map[string]Team
	Wins  map[string]int
}

func (l *League) MatchResult(team1 string, score1 int, team2 string, score2 int) {
	if _, ok := l.Teams[team1]; !ok {
		return
	}

	if _, ok := l.Teams[team2]; !ok {
		return
	}

	if score1 == score2 {
		return
	}

	if score1 > score2 {
		l.Wins[team1] += 1
	} else {
		l.Wins[team2] += 1
	}
}

func (l League) Ranking() []string {
	names := make([]string, 0, len(l.Teams))

	for k := range l.Teams {
		names = append(names, k)
	}

	sort.Slice(names, func(i, j int) bool {
		return l.Wins[names[i]] > l.Wins[names[j]]
	})
	return names
}

type Ranker interface {
	Ranking() []string
}

func PrintRanking(ranker Ranker, w io.Writer) {
	result := ranker.Ranking()

	for _, v := range result {
		io.WriteString(w, v)
		w.Write([]byte("\n"))
	}
}

func main() {
	league := League{
		Name: "Big League",
		Teams: map[string]Team{
			"Italy": {
				Name:    "Italy",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"France": {
				Name:    "France",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"India": {
				Name:    "India",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"Nigeria": {
				Name:    "Nigeria",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
		},
		Wins: map[string]int{},
	}

	league.MatchResult("Italy", 50, "France", 70)
	league.MatchResult("India", 85, "Nigeria", 80)
	league.MatchResult("Italy", 60, "India", 55)
	league.MatchResult("France", 100, "Nigeria", 110)
	league.MatchResult("Italy", 65, "Nigeria", 70)
	league.MatchResult("France", 95, "India", 80)

	PrintRanking(league, os.Stdout)
}
