package main

import (
	"slices"
)

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func NewPlayer(name string, goals, misses, assists int) Player {
	return Player{
		Name:    name,
		Goals:   goals,
		Misses:  misses,
		Assists: assists,
		Rating:  calculateRating(goals, misses, assists),
	}
}

func calculateRating(goals, misses, assists int) float64 {
	return (float64(goals) + (float64(assists) / 2.0)) / float64(max(misses, 1))
}

func compName(a, b Player) int {
	switch {
	case a.Name > b.Name:
		return 1
	case a.Name < b.Name:
		return -1
	default:
		return 0
	}
}

func compGolals(a, b Player) int {
	switch {
	case a.Goals < b.Goals:
		return 1
	case a.Goals > b.Goals:
		return -1
	default:
		return compName(a, b)
	}
}

func compRating(a, b Player) int {
	switch {
	case a.Rating < b.Rating:
		return 1
	case a.Rating > b.Rating:
		return -1
	default:
		return compName(a, b)
	}
}

func compGm(a, b Player) int {
	switch {
	case a.Rating < b.Rating:
		return 1
	case a.Rating > b.Rating:
		return -1
	default:
		return compName(a, b)
	}
}

func goalsSort(players []Player) []Player {
	res := make([]Player, len(players))
	copy(res, players)
	slices.SortFunc(res, compGolals)
	return res
}

func ratingSort(players []Player) []Player {
	res := make([]Player, len(players))
	copy(res, players)
	slices.SortFunc(res, compRating)
	return res
}

func gmSort(players []Player) []Player {
	res := make([]Player, len(players))
	copy(res, players)
	slices.SortFunc(res, compGm)
	return res
}
