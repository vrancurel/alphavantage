package alphavantage

import (
	"fmt"
	"log"
	"time"
)

// ListingType represents the type of a listing (ETF or Stock).
type ListingType int

const (
	// Etf represents an Exchange Traded Fund listing.
	Etf ListingType = iota
	// Stock represents a Stock listing.
	Stock
)

// ListingState represents the state of a listing (Delisted or Active).
type ListingState int

const (
	// Delisted represents a delisted listing.
	Delisted ListingState = iota
	// Active represents an active listing.
	Active
)

// SymbolStatus represents the status of a symbol.
type SymbolStatus struct {
	Symbol       string
	Name         string
	Exchange     string
	Type         ListingType
	ActiveSince  time.Time
	DelistedFrom time.Time
	Status       ListingState
}

// ListingStatus represents the listing status containing multiple symbol statuses.
type ListingStatus struct {
	SymbolStatuses []SymbolStatus
}

func toListingStatus(data [][]string, delisted bool) (*ListingStatus, error) {
	var ls ListingStatus
	ls.SymbolStatuses = make([]SymbolStatus, 0)
	for idx, row := range data {
		// skip header
		if idx == 0 {
			continue
		}
		var s SymbolStatus
		if len(row) != 7 {
			log.Printf("Invalid number of fields for %v", len(row))
			continue
		}
		s.Symbol = row[0]
		s.Name = row[1]
		s.Exchange = row[2]
		switch row[3] {
		case "ETF":
			s.Type = Etf
		case "Stock":
			s.Type = Stock
		default:
			log.Printf("Unknown listing type %v", row[3])
			continue
		}
		var err error
		s.ActiveSince, err = time.Parse("2006-01-02", row[4])
		if err != nil {
			log.Printf("Error parsing time %v", row[4])
			continue
		}
		if delisted {
			s.DelistedFrom, err = time.Parse("2006-01-02", row[5])
			if err != nil {
				log.Printf("Error parsing time %v", row[5])
				continue
			}
		} else {
			s.DelistedFrom = time.Time{}
		}
		switch row[6] {
		case "Active":
			s.Status = Active
		case "Delisted":
			s.Status = Delisted
		default:
			log.Printf("Unknown status %v", row[6])
			continue
		}
		ls.SymbolStatuses = append(ls.SymbolStatuses, s)
	}
	return &ls, nil
}

// ListingStatus fetches and returns the listing status data (either active or delisted).
// if not active will list delisted tickers
func (c *Client) ListingStatus(delisted bool) (*ListingStatus, error) {
	const function = "LISTING_STATUS"
	var state string
	if delisted {
		state = "delisted"
	} else {
		state = "active"
	}
	url := fmt.Sprintf("%s/query?function=%s&state=%s&apikey=%s", baseURL, function, state, c.apiKey)
	data, err := c.makeHTTPRequestForCsv(url)
	if err != nil {
		return nil, err
	}

	return toListingStatus(data, delisted)
}
