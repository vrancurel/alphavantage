package alphavantage

import (
	"testing"

	"github.com/AMekss/assert"
)

func TestToEarningsTranscript(t *testing.T) {
	var buf = `
{
    "symbol": "STOCK1",
    "quarter": "2024Q1",
    "transcript": [
        {
            "speaker": "Example Speaker1",
            "title": "Global Head of Investor Relations",
            "content": "Thank you. I'd like to welcome you to STOCK1's First Quarter 2024 Earnings Presentation.",
            "sentiment": "0.6"
        },
        {
            "speaker": "Example Speaker2",
            "title": "CEO",
            "content": "Thank you for joining us. In the first quarter, we had solid performance across revenue and cash flow.",
            "sentiment": "0.7"
        }
    ]
}
`
	transcript, err := toEarningsTranscript([]byte(buf))
	assert.NoError(t.Fatalf, err)

	assert.EqualStrings(t, "STOCK1", transcript.Symbol)
	assert.EqualStrings(t, "2024Q1", transcript.Quarter)
	assert.EqualStrings(t, "Example Speaker1", transcript.Transcript[0].Speaker)
	assert.EqualStrings(t, "Global Head of Investor Relations", transcript.Transcript[0].Title)
	assert.EqualStrings(t, "0.6", transcript.Transcript[0].Sentiment)
	assert.EqualStrings(t, "Example Speaker2", transcript.Transcript[1].Speaker)
	assert.EqualStrings(t, "CEO", transcript.Transcript[1].Title)
	assert.EqualStrings(t, "0.7", transcript.Transcript[1].Sentiment)
}
