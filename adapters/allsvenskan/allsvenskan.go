package allsvenskan

import (
	"go-al/match"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Adapter struct{}

func (a Adapter) Fetch() <-chan []match.Match {
	r := make(chan []match.Match)
	go func() {
		defer close(r)
		var matches []match.Match
		res, err := http.Get("https://tabellen.se/fotboll/allsvenskan")
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		doc.Find(".match-inner").Each(func(i int, s *goquery.Selection) {
			result := strings.Split(s.Find(".live_result").Text(), "-")
			date := "date"
			homeTeamName := s.Find(".home").Text()
			awayTeamName := s.Find(".away").Text()
			liveInfo := s.Find(".time_info").Text()
			homeTeamScore := ""
			awayTeamScore := ""
			if len(result) == 2 {
				homeTeamScore = result[0]
				awayTeamScore = result[1]
			}
			match := match.New(strings.TrimSpace(homeTeamName), homeTeamScore, strings.TrimSpace(awayTeamName), awayTeamScore, date, liveInfo)
			matches = append(matches, match)
		})
		r <- matches
	}()
	return r
}
