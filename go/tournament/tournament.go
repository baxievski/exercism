// Package tournament implements a solution for the exercism tournament challenge
package tournament

import (
	"encoding/csv"
	"fmt"
	"io"
	"sort"
)

// Team represents a team in a footbal competition
type Team struct {
	name          string
	matchesPlayed int
	wins          int
	draws         int
	losses        int
}

// Table represents a table of results of matches played between teams
type Table struct {
	teams map[string]*Team
}

func (team *Team) points() int {
	return 3*team.wins + team.draws
}

func (team *Team) win() {
	team.matchesPlayed++
	team.wins++
}

func (team *Team) loose() {
	team.matchesPlayed++
	team.losses++
}

func (team *Team) draw() {
	team.matchesPlayed++
	team.draws++
}

func (t *Table) addMatch(team1Name, team2Name, result string) error {
	var team1, team2 *Team
	if team, ok := t.teams[team1Name]; ok {
		team1 = team
	} else {
		team1 = &Team{name: team1Name}
	}
	if team, ok := t.teams[team2Name]; ok {
		team2 = team
	} else {
		team2 = &Team{name: team2Name}
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
	t.teams[team1Name] = team1
	t.teams[team2Name] = team2
	return nil
}

func (t *Table) getSortedTeams() []*Team {
	teams := make([]*Team, 0, len(t.teams))
	for _, v := range t.teams {
		teams = append(teams, v)
	}
	sort.Slice(teams, func(p, q int) bool {
		if teams[p].points() != teams[q].points() {
			return teams[p].points() > teams[q].points()
		}
		return teams[p].name < teams[q].name
	})
	return teams
}

// Tally reads a stream of matches and writes a stream representing the table
func Tally(i io.Reader, o io.Writer) error {
	table := &Table{teams: map[string]*Team{}}
	r := csv.NewReader(i)
	r.Comma = ';'
	r.Comment = '#'
	r.FieldsPerRecord = 3
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		err = table.addMatch(record[0], record[1], record[2])
		if err != nil {
			return err
		}
	}
	l := "%-31v|%3v |%3v |%3v |%3v |%3v\n"
	fmt.Fprintf(o, l, "Team", "MP", "W", "D", "L", "P")
	for _, t := range table.getSortedTeams() {
		fmt.Fprintf(o, l, t.name, t.matchesPlayed, t.wins, t.draws, t.losses, t.points())
	}
	return nil
}
