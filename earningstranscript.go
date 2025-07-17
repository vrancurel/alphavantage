package alphavantage

import (
	"encoding/json"
	"fmt"
)

// EarningsTranscript represents the earnings call transcript data.
type EarningsTranscript struct {
	Symbol     string              `json:"symbol"`
	Quarter    string              `json:"quarter"`
	Transcript []TranscriptSegment `json:"transcript"`
}

// TranscriptSegment represents an individual segment of the earnings call transcript.
type TranscriptSegment struct {
	Speaker   string `json:"speaker"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Sentiment string `json:"sentiment"`
}

// toEarningsTranscript parses the JSON response into the EarningsTranscript struct.
func toEarningsTranscript(buf []byte) (*EarningsTranscript, error) {
	transcript := &EarningsTranscript{}
	if err := json.Unmarshal(buf, transcript); err != nil {
		return nil, err
	}
	return transcript, nil
}

// EarningsCallTranscript fetches and returns the earnings call transcript data for the specified symbol and quarter.
func (c *Client) EarningsCallTranscript(symbol string, quarter string) (*EarningsTranscript, error) {
	const function = "EARNINGS_CALL_TRANSCRIPT"
	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&quarter=%s&apikey=%s", baseURL, function, symbol, quarter, c.apiKey)

	body, err := c.makeHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	return toEarningsTranscript(body)
}
