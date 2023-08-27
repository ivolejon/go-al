package match

import (
	"strings"
)

type Status int

const (
	Pending Status = iota
	Finished
	NotPlayed
)

func (s Status) String() string {
	return []string{"Igång", "Klar", "Inte spelad"}[s]
}

type Match struct {
	HomeTeam  string
	HomeScore string
	AwayTeam  string
	AwayScore string
	Date      string
	liveInfo  string
}

// New returns a new match with the given data.
func New(homeTeam string, homeScore string, awayTeam string, awayScore string, date string, liveInfo string) Match {
	m := Match{HomeTeam: homeTeam, HomeScore: homeScore, AwayTeam: awayTeam, AwayScore: awayScore, Date: date, liveInfo: liveInfo}
	return m
}

// Status returns the status of the match.
func (m Match) Status() string {
	if m.HomeScore == "" && m.AwayScore == "" {
		return NotPlayed.String()
	}
	return Finished.String()
}

// IsFinished returns whether the match is finished.
func (m Match) IsFinished() bool {
	return m.HomeScore != "" && m.AwayScore != ""
}

// IsPending returns whether the match is pending.
func (m Match) IsPending() bool {
	return m.liveInfo != ""
}

// IsNotPlayed returns whether the match is not played.
func (m Match) IsNotPlayed() bool {
	return m.HomeScore == "" && m.AwayScore == ""
}

func replaceAt(s string, i int, c rune) string {
	r := []rune(s)
	r[i] = c
	return string(r)
}

func (m Match) String(width int) string {
	homeTeamStringLength := len([]rune(m.HomeTeam))
	awayTeamStringLength := len([]rune(m.AwayTeam))
	padding := strings.Repeat(" ", (width-4)-(homeTeamStringLength+awayTeamStringLength))
	row := m.HomeTeam + padding + m.AwayTeam
	center := (width - 4) / 2
	if m.IsFinished() == true || m.IsPending() == true {
		row = replaceAt(row, center, '-')
		row = replaceAt(row, center-2, rune(m.HomeScore[0]))
		row = replaceAt(row, center+2, rune(m.AwayScore[0]))
	} else {
		row = replaceAt(row, center, '-')
	}
	return row
}
