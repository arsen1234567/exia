package models

import "time"

type GoogleNews struct {
	Link          string    `json:"link"`
	Snippet       string    `json:"snippet"`
	Source        string    `json:"source"`
	Imageurl      string    `json:"imageurl"`
	Position      int       `json:"position"`
	Date_news     string    `json:"date_news"`
	Keyword       string    `json:"keyword"`
	Title         string    `json:"title"`
	Sentences     string    `json:"sentences"`
	Dict_sent     float64   `json:"dict_sent"`
	Dict_sent_wrd string    `json:"dict_sent_wrd"`
	Article_text  string    `json:"article_text"`
	Id            int       `json:"id"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
}

type GoogleNewsSummary struct {
	CompanyName   string    `json:"company_name"`
	CompanyCount  int       `json:"company_count"`
	PositiveCount int       `json:"positive_count"`
	NewsSource    string    `json:"source"`
	NegativeCount int       `json:"negative_count"`
	NeutralCount  int       `json:"neutral_count"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
}

type NewsSummary struct {
	Keyword   string `json:"keyword"`
	Snippet   string `json:"snippet"`
	Source    string `json:"source"`
	Sentiment string `json:"dict_sent_wrd"`
}
