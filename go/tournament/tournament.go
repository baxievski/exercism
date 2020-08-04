// Package tournament implements a solution for the exercism tournament challenge
package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Team represents a team
type Team struct {
	Name          string
	MatchesPlayed int
	Wins          int
	Draws         int
	Losses        int
}

// Table represents a table of teams
type Table struct {
	Teams map[string]*Team
}

func (team *Team) points() int {
	resultsToPoints := map[string]int{
		"win":  3,
		"draw": 1,
		"loss": 0,
	}
	points := resultsToPoints["win"] * team.Wins
	points += resultsToPoints["draw"] * team.Draws
	points += resultsToPoints["loss"] * team.Losses
	return points
}

func (team *Team) win() {
	team.MatchesPlayed++
	team.Wins++
}

func (team *Team) loose() {
	team.MatchesPlayed++
	team.Losses++
}

func (team *Team) draw() {
	team.MatchesPlayed++
	team.Draws++
}

func (t *Table) addMatch(team1Name, team2Name, result string) error {
	var team1, team2 *Team
	if team, ok := t.Teams[team1Name]; ok {
		team1 = team
	} else {
		team1 = &Team{Name: team1Name}
	}
	if team, ok := t.Teams[team2Name]; ok {
		team2 = team
	} else {
		team2 = &Team{Name: team2Name}
	}

	if result == "draw" {
		team1.draw()
		team2.draw()
	} else if result == "win" {
		team1.win()
		team2.loose()
	} else if result == "loss" {
		team1.loose()
		team2.win()
	} else {
		return fmt.Errorf("invalid result '%v'", result)
	}
	t.Teams[team1Name] = team1
	t.Teams[team2Name] = team2
	return nil
}

func (t *Table) getSortedTeams() []*Team {
	teams := make([]*Team, 0, len(t.Teams))
	for _, v := range t.Teams {
		teams = append(teams, v)
	}
	sort.Slice(teams, func(p, q int) bool {
		if teams[p].points() != teams[q].points() {
			return teams[p].points() > teams[q].points()
		}
		if teams[p].Name != teams[q].Name {
			return teams[p].Name < teams[q].Name
		}
		return teams[p].MatchesPlayed < teams[q].MatchesPlayed
	})
	return teams
}

// Tally reads a stream of matches and writes a stream representing the table
func Tally(input io.Reader, output io.Writer) error {
	scanner := bufio.NewScanner(input)
	table := &Table{Teams: map[string]*Team{}}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		if line == "" {
			continue
		}
		match := strings.Split(string(line), ";")
		if len(match) != 3 {
			return fmt.Errorf("invalid input %v", line)
		}

		if table.addMatch(match[0], match[1], match[2]) != nil {
			return fmt.Errorf("invalid input %v", line)
		}
	}

	template := "%-31v|%3v |%3v |%3v |%3v |%3v\n"
	fmt.Fprintf(output, template, "Team", "MP", "W", "D", "L", "P")
	for _, team := range table.getSortedTeams() {
		fmt.Fprintf(
			output,
			template,
			team.Name,
			team.MatchesPlayed,
			team.Wins,
			team.Draws,
			team.Losses,
			team.points(),
		)
	}

	return nil
}
