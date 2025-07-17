package alphavantage

import (
	"testing"

	"github.com/AMekss/assert"
)

func TestToNewsSentiment(t *testing.T) {
	var buf = `
{
    "items": "6",
    "sentiment_score_definition": "x <= -0.35: Bearish; -0.35 < x <= -0.15: Somewhat-Bearish; -0.15 < x < 0.15: Neutral; 0.15 <= x < 0.35: Somewhat_Bullish; x >= 0.35: Bullish",
    "relevance_score_definition": "0 < x <= 1, with a higher score indicating higher relevance.",
    "feed": [
        {
            "title": "TechCorp's Longest-Serving Designer John Smith To Retire, Marking End Of Design Leader Era - TechCorp  ( NASDAQ:TECH ) ",
            "url": "https://www.example1.com/news/24/02/37114172/techcorps-longest-serving-designer-john-smith-to-retire-marking-end-of-design-leader-era",
            "time_published": "20240214T013334",
            "authors": [
                "Jane Reporter"
            ],
            "summary": "TechCorp Inc.'s TECH longest-serving senior industrial designer, John Smith, has reportedly decided to retire, marking the near-complete turnover of a team formerly guided by Design Leader.",
            "banner_image": "https://cdn.example1.com/files/images/story/2024/TechCorp-City-Country-media-preview-store_1.jpeg?width=1200&height=800&fit=crop",
            "source": "Example1",
            "category_within_source": "News",
            "source_domain": "www.example1.com",
            "topics": [
                {
                    "topic": "Earnings",
                    "relevance_score": "0.360215"
                },
                {
                    "topic": "Technology",
                    "relevance_score": "1.0"
                },
                {
                    "topic": "Financial Markets",
                    "relevance_score": "0.108179"
                }
            ],
            "overall_sentiment_score": 0.173373,
            "overall_sentiment_label": "Somewhat-Bullish",
            "ticker_sentiment": [
                {
                    "ticker": "TECH",
                    "relevance_score": "0.634626",
                    "ticker_sentiment_score": "0.344843",
                    "ticker_sentiment_label": "Somewhat-Bullish"
                }
            ]
        },
        {
            "title": "Bob Analyst Attributes Tuesday Stocks Sell-Off To 'Bad Judgment' By Investors: 'Market Won't Bottom All At Once'",
            "url": "https://www.example1.com/analyst-ratings/analyst-color/24/02/37114017/bob-analyst-attributes-tuesday-stocks-sell-off-to-bad-judgment-by-investors-market-w",
            "time_published": "20240214T012429",
            "authors": [
                "Example1 AI"
            ],
            "summary": "Bob Analyst, the host of EXBC's \"Money Show,\" has attributed the recent market sell-off to poor decision-making by investors.",
            "banner_image": "https://cdn.example1.com/files/images/story/2024/bob-analyst-shutterstock_12.png?width=1200&height=800&fit=crop",
            "source": "Example1",
            "category_within_source": "Markets",
            "source_domain": "www.example1.com",
            "topics": [
                {
                    "topic": "Economy - Monetary",
                    "relevance_score": "0.576289"
                },
                {
                    "topic": "Financial Markets",
                    "relevance_score": "0.108179"
                },
                {
                    "topic": "Manufacturing",
                    "relevance_score": "0.5"
                },
                {
                    "topic": "Earnings",
                    "relevance_score": "0.310843"
                },
                {
                    "topic": "Technology",
                    "relevance_score": "0.5"
                }
            ],
            "overall_sentiment_score": -0.160165,
            "overall_sentiment_label": "Somewhat-Bearish",
            "ticker_sentiment": [
                {
                    "ticker": "SOFT",
                    "relevance_score": "0.102888",
                    "ticker_sentiment_score": "0.034995",
                    "ticker_sentiment_label": "Neutral"
                },
                {
                    "ticker": "TECH",
                    "relevance_score": "0.102888",
                    "ticker_sentiment_score": "0.034995",
                    "ticker_sentiment_label": "Neutral"
                },
                {
                    "ticker": "ELEC",
                    "relevance_score": "0.102888",
                    "ticker_sentiment_score": "0.034995",
                    "ticker_sentiment_label": "Neutral"
                }
            ]
        },
        {
            "title": "Big 7 Lose More Than $250 Billion As Blazing Inflation Data Triggers Market Sell-Off - TechCorp  ( NASDAQ:TECH ) , ShopCorp  ( NASDAQ:SHOP ) ",
            "url": "https://www.example1.com/markets/equities/24/02/37105400/big-7-lose-more-than-250-billion-as-blazing-inflation-data-triggers-market-sell-off",
            "time_published": "20240213T205849",
            "authors": [
                "Mike Writer"
            ],
            "summary": "The Big Seven stocks experienced a sharp plunge in market value Tuesday, shedding over $250 billion in value a single session as hotter-than-expected inflation data sparked a widespread sell-off in U.S. stocks. The group, which includes heavyweights SoftCorp Inc. SOFT, TechCorp Inc.",
            "banner_image": "https://cdn.example1.com/files/images/story/2024/Bull-and-bear-illustration_0.jpeg?width=1200&height=800&fit=crop",
            "source": "Example1",
            "category_within_source": "Markets",
            "source_domain": "www.example1.com",
            "topics": [
                {
                    "topic": "Economy - Monetary",
                    "relevance_score": "0.990678"
                },
                {
                    "topic": "Retail & Wholesale",
                    "relevance_score": "0.25"
                },
                {
                    "topic": "Financial Markets",
                    "relevance_score": "0.796627"
                },
                {
                    "topic": "Manufacturing",
                    "relevance_score": "0.25"
                },
                {
                    "topic": "Technology",
                    "relevance_score": "0.25"
                },
                {
                    "topic": "Finance",
                    "relevance_score": "0.25"
                }
            ],
            "overall_sentiment_score": -0.184637,
            "overall_sentiment_label": "Somewhat-Bearish",
            "ticker_sentiment": [
                {
                    "ticker": "SOFT",
                    "relevance_score": "0.279015",
                    "ticker_sentiment_score": "-0.201297",
                    "ticker_sentiment_label": "Somewhat-Bearish"
                },
                {
                    "ticker": "SRCH",
                    "relevance_score": "0.188193",
                    "ticker_sentiment_score": "0.0",
                    "ticker_sentiment_label": "Neutral"
                },
                {
                    "ticker": "SOCL",
                    "relevance_score": "0.188193",
                    "ticker_sentiment_score": "0.0",
                    "ticker_sentiment_label": "Neutral"
                },
                {
                    "ticker": "CHIP",
                    "relevance_score": "0.188193",
                    "ticker_sentiment_score": "0.0",
                    "ticker_sentiment_label": "Neutral"
                },
                {
                    "ticker": "TECH",
                    "relevance_score": "0.188193",
                    "ticker_sentiment_score": "0.0",
                    "ticker_sentiment_label": "Neutral"
                },
                {
                    "ticker": "ELEC",
                    "relevance_score": "0.279015",
                    "ticker_sentiment_score": "-0.201297",
                    "ticker_sentiment_label": "Somewhat-Bearish"
                },
                {
                    "ticker": "INVS",
                    "relevance_score": "0.094762",
                    "ticker_sentiment_score": "-0.092722",
                    "ticker_sentiment_label": "Neutral"
                },
                {
                    "ticker": "SHOP",
                    "relevance_score": "0.188193",
                    "ticker_sentiment_score": "0.0",
                    "ticker_sentiment_label": "Neutral"
                }
            ]
        },
        {
            "title": "ChipCorp's Soaring Success with Phone Modem Deal - Analyst Eyes Higher Valuation - ChipCorp  ( NASDAQ:CHIP ) ",
            "url": "https://www.example1.com/analyst-ratings/analyst-color/24/02/37105391/chipcorps-soaring-success-with-phone-modem-deal-analyst-eyes-higher-valuation",
            "time_published": "20240213T205754",
            "authors": [
                "Sarah Analyst"
            ],
            "summary": "Example Finance analyst Tom Kumar reiterates an Overweight rating on ChipCorp Inc CHIP with a price target of $165.",
            "banner_image": "https://cdn.example1.com/files/images/story/2024/chipcorp.jpeg?width=1200&height=800&fit=crop",
            "source": "Example1",
            "category_within_source": "Trading",
            "source_domain": "www.example1.com",
            "topics": [
                {
                    "topic": "Financial Markets",
                    "relevance_score": "0.266143"
                },
                {
                    "topic": "Manufacturing",
                    "relevance_score": "0.5"
                },
                {
                    "topic": "Earnings",
                    "relevance_score": "0.108179"
                },
                {
                    "topic": "Technology",
                    "relevance_score": "0.5"
                }
            ],
            "overall_sentiment_score": 0.109896,
            "overall_sentiment_label": "Neutral",
            "ticker_sentiment": [
                {
                    "ticker": "CHIP",
                    "relevance_score": "0.506613",
                    "ticker_sentiment_score": "0.198157",
                    "ticker_sentiment_label": "Somewhat-Bullish"
                },
                {
                    "ticker": "TECH",
                    "relevance_score": "0.318897",
                    "ticker_sentiment_score": "0.068598",
                    "ticker_sentiment_label": "Neutral"
                }
            ]
        },
        {
            "title": "These 3 Companies Are Shattering Quarterly Records",
            "url": "https://www.example2.com/commentary/2225782/these-3-companies-are-shattering-quarterly-records",
            "time_published": "20240213T203400",
            "authors": [
                "David Reporter"
            ],
            "summary": "The Q4 cycle continues to chug along, with many positive quarterly reports already delivered. Notably, these three companies posted results that reflected quarterly records.",
            "banner_image": "https://staticx-tuner.example2.com/images/articles/main/0a/1429.jpg",
            "source": "Example2 Commentary",
            "category_within_source": "n/a",
            "source_domain": "www.example2.com",
            "topics": [
                {
                    "topic": "Earnings",
                    "relevance_score": "0.999999"
                },
                {
                    "topic": "Technology",
                    "relevance_score": "0.5"
                },
                {
                    "topic": "Manufacturing",
                    "relevance_score": "0.5"
                }
            ],
            "overall_sentiment_score": 0.292724,
            "overall_sentiment_label": "Somewhat-Bullish",
            "ticker_sentiment": [
                {
                    "ticker": "SRCH",
                    "relevance_score": "0.123844",
                    "ticker_sentiment_score": "0.046335",
                    "ticker_sentiment_label": "Neutral"
                },
                {
                    "ticker": "TECH",
                    "relevance_score": "0.414559",
                    "ticker_sentiment_score": "0.320203",
                    "ticker_sentiment_label": "Somewhat-Bullish"
                },
                {
                    "ticker": "SHOE",
                    "relevance_score": "0.359881",
                    "ticker_sentiment_score": "0.282165",
                    "ticker_sentiment_label": "Somewhat-Bullish"
                }
            ]
        },
        {
            "title": "The Big Seven are so big, they are worth as much as all the stocks in Country1, Country2 and Country3 put together",
            "url": "https://www.example3.com/story/the-big-seven-are-so-big-they-are-worth-as-much-as-all-the-stocks-in-country1-country2-and-country3-put-together-8495b644",
            "time_published": "20240213T094800",
            "authors": [
                "Alex Writer"
            ],
            "summary": "The epic scale of the so-called Big Seven was put under the microscope by analysts at Example Bank, which finds the grouping of U.S. tech giants equal in size to multiple major stock markets put together.",
            "banner_image": "https://images.example3.com/im-76258514?width=700&height=427",
            "source": "Example3",
            "category_within_source": "Top Stories",
            "source_domain": "www.example3.com",
            "topics": [
                {
                    "topic": "Retail & Wholesale",
                    "relevance_score": "0.333333"
                },
                {
                    "topic": "Financial Markets",
                    "relevance_score": "0.538269"
                },
                {
                    "topic": "Manufacturing",
                    "relevance_score": "0.333333"
                },
                {
                    "topic": "Technology",
                    "relevance_score": "0.333333"
                }
            ],
            "overall_sentiment_score": 0.211976,
            "overall_sentiment_label": "Somewhat-Bullish",
            "ticker_sentiment": [
                {
                    "ticker": "SOFT",
                    "relevance_score": "0.329551",
                    "ticker_sentiment_score": "0.169969",
                    "ticker_sentiment_label": "Somewhat-Bullish"
                },
                {
                    "ticker": "SOCL",
                    "relevance_score": "0.329551",
                    "ticker_sentiment_score": "0.169969",
                    "ticker_sentiment_label": "Somewhat-Bullish"
                },
                {
                    "ticker": "CHIP",
                    "relevance_score": "0.329551",
                    "ticker_sentiment_score": "0.169969",
                    "ticker_sentiment_label": "Somewhat-Bullish"
                },
                {
                    "ticker": "TECH",
                    "relevance_score": "0.329551",
                    "ticker_sentiment_score": "0.169969",
                    "ticker_sentiment_label": "Somewhat-Bullish"
                },
                {
                    "ticker": "ELEC",
                    "relevance_score": "0.329551",
                    "ticker_sentiment_score": "0.169969",
                    "ticker_sentiment_label": "Somewhat-Bullish"
                },
                {
                    "ticker": "SHOP",
                    "relevance_score": "0.329551",
                    "ticker_sentiment_score": "0.169969",
                    "ticker_sentiment_label": "Somewhat-Bullish"
                }
            ]
        }
    ]
}
`
	newsSentiment, err := toNewsSentiment([]byte(buf))
	assert.NoError(t.Fatalf, err)

	assert.EqualInt(t, 6, len(newsSentiment.Feed))
}
