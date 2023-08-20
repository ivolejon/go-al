package allsvenskan

import (
	"go-al/match"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Adapter struct{}

func (a Adapter) Fetch() <-chan []match.Match {
	r := make(chan []match.Match)
	go func() {
		defer close(r)
		var matches []match.Match
		res, err := http.Get("https://www.svenskfotboll.se/serier-cuper/tabell-och-resultat/allsvenskan-2023/101999/")
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		doc.Find(".match-list__match").Each(func(i int, s *goquery.Selection) {
			date := s.Find(".match-list__info .match-list__place .match-list__date").Text()
			homeTeamName := s.Find(".match-list__home .match-list__team-name").Text()
			homeTeamScore := s.Find(".match-list__home .match-list__score").Text()
			awayTeamName := s.Find(".match-list__away .match-list__team-name").Text()
			awayTeamScore := s.Find(".match-list__away .match-list__score").Text()
			match := match.New(homeTeamName, homeTeamScore, awayTeamName, awayTeamScore, date)
			matches = append(matches, match)
		})
		r <- matches
	}()
	return r
}
