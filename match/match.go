package match

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
}

// New returns a new match with the given data.
func New(homeTeam string, homeScore string, awayTeam string, awayScore string, date string) Match {
	m := Match{HomeTeam: homeTeam, HomeScore: homeScore, AwayTeam: awayTeam, AwayScore: awayScore, Date: date}
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
	return m.HomeScore == "" && m.AwayScore == ""
}

// IsNotPlayed returns whether the match is not played.
func (m Match) IsNotPlayed() bool {
	return m.HomeScore == "" && m.AwayScore == ""
}

func (m Match) String() string {
	return m.HomeTeam + " " + m.HomeScore + " - " + m.AwayScore + " " + m.AwayTeam + " " + m.Date
}
