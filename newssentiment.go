package alphavantage

import (
	"encoding/json"
	"fmt"
)

// Topic represents a topic related to a news item.
type Topic struct {
	Topic          string  `json:"topic"`
	RelevanceScore float64 `json:"relevance_score,string"`
}

// TickerSentiment represents the sentiment of a ticker in a news item.
type TickerSentiment struct {
	Ticker               string  `json:"ticker"`
	RelevanceScore       float64 `json:"relevance_score,string"`
	TickerSentimentScore float64 `json:"ticker_sentiment_score,string"`
	TickerSentimentLabel string  `json:"ticker_sentiment_label"`
}

// FeedItem represents a news feed item with various attributes.
type FeedItem struct {
	Title                 string            `json:"title"`
	URL                   string            `json:"url"`
	TimePublished         string            `json:"time_published"`
	Authors               []string          `json:"authors"`
	Summary               string            `json:"summary"`
	BannerImage           string            `json:"banner_image"`
	Source                string            `json:"source"`
	CategoryWithinSource  string            `json:"category_within_source"`
	SourceDomain          string            `json:"source_domain"`
	Topics                []Topic           `json:"topics"`
	OverallSentimentScore AVFloat64         `json:"overall_sentiment_score,string"`
	OverallSentimentLabel string            `json:"overall_sentiment_label"`
	TickerSentiment       []TickerSentiment `json:"ticker_sentiment"`
}

// NewsSentiment represents the sentiment analysis of news items.
type NewsSentiment struct {
	Items                    string     `json:"items"`
	SentimentScoreDefinition string     `json:"sentiment_score_definition"`
	RelevanceScoreDefinition string     `json:"relevance_score_definition"`
	Feed                     []FeedItem `json:"feed"`
}

func toNewsSentiment(buf []byte) (*NewsSentiment, error) {
	newsSentiment := &NewsSentiment{}
	if err := json.Unmarshal(buf, newsSentiment); err != nil {
		return nil, err
	}
	return newsSentiment, nil
}

// NewsSentiment fetches and returns the news sentiment data for the specified tickers.
func (c *Client) NewsSentiment(tickers string, sortType SortType, limit int, timeFrom string, timeTo string) (*NewsSentiment, error) {
	const function = "NEWS_SENTIMENT"
	url := fmt.Sprintf("%s/query?function=%s&tickers=%s&sort=%s&limit=%d&apikey=%s", baseURL, function, tickers, sortType, limit, c.apiKey)

	if timeFrom != "" {
		url += fmt.Sprintf("&time_from=%s", timeFrom)
	}
	if timeTo != "" {
		url += fmt.Sprintf("&time_to=%s", timeTo)
	}

	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	return toNewsSentiment(body)
}
