package main

import (
	"io"
	"os"
	"sort"
)

type Team struct {
	Name string
	Players []string
}

type League struct {
	Name string
	Teams map[string]Team
	Wins map[string]int
}

func (l *League) MatchResult(firstTeamName string, firstTeamScore int, secondTeamName string, secondTeamScore int) {
	if _, ok := l.Teams[firstTeamName]; !ok {
		return
	}
	if _, ok := l.Teams[secondTeamName]; !ok {
		return
	}
	if firstTeamScore == secondTeamScore {
		return
	}
	if firstTeamScore > secondTeamScore {
		l.Wins[firstTeamName]++
	} else {
		l.Wins[secondTeamName]++
	}
}

type kv struct {
	Key string
	Value int
}

func (l League) Ranking() []string {
	names := make([]string, 0, len(l.Teams))
	for name := range l.Teams {
		names = append(names, name)
	}
	sort.Slice(names, func(i, j int) bool {
		return l.Wins[names[i]] > l.Wins[names[j]]
	})
	return names
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	for _, val := range r.Ranking() {
		io.WriteString(w, val)
		io.WriteString(w, "\n")
	}
}

func main() {
	l := League{
		Name: "The Major League",
		Teams: map[string]Team{
			"Austria": {
				Name: "Austria",
				Players: []string{"Player 1", "Player 2", "Player 3", "Player 4", "Player 5"},
			},
			"Nigeria": {
				Name: "Nigeria",
				Players: []string{"Player 1", "Player 2", "Player 3", "Player 4", "Player 5"},
			},
			"London": {
				Name: "London",
				Players: []string{"Player 1", "Player 2", "Player 3", "Player 4", "Player 5"},
			},
			"Egypt": {
				Name: "Egypt",
				Players: []string{"Player 1", "Player 2", "Player 3", "Player 4", "Player 5"},
			},
		},
		Wins: map[string]int{},
	}
	
	l.MatchResult("Nigeria", 50, "London", 30)
	l.MatchResult("Nigeria", 70, "Egypt", 20)
	l.MatchResult("London", 50, "Austria", 70)
	l.MatchResult("Egypt", 50, "London", 50)
	l.MatchResult("Nigeria", 50, "Austria", 40)

	RankPrinter(l, os.Stdout)
}