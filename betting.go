// Copyright 2013 Alessandro De Donno

// "Betfair API-NG Golang Library" is dual-licensed: for free software projects
// please refer to GPLv3 (see declaration above), for commercial software
// please contact the author.
// If you are a contributor and need any clarification, please contact the
// author.

// For free software projects:

// This file is part of "Betfair API-NG Golang Library".

// "Betfair API-NG Golang Library" is free software: you can redistribute it
// and/or modify it under the terms of the GNU General Public License as
// published by the Free Software Foundation, either version 3 of the License,
// or (at your option) any later version.

// "Betfair API-NG Golang Library" is distributed in the hope that it will be
// useful, but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with "Betfair API-NG Golang Library".  If not, see
// <http://www.gnu.org/licenses/>.

package betfair

import (
	"encoding/json"
	"strings"
	"time"
	// "log"
)

type baseEnumVal string

// OrderProjVal Enum for order projections
type OrderProjVal baseEnumVal

// MarketProjVal Enum for market projections
type MarketProjVal baseEnumVal

// MatchProjVal Enum for match projections
type MatchProjVal baseEnumVal

// PriceDataVal Enum for price data values
type PriceDataVal baseEnumVal

// RunnerStatusVal Enum for runner status values
type RunnerStatusVal baseEnumVal

// OrderTypeVal Enum of Order types
type OrderTypeVal baseEnumVal

// OrderStatusVal Enum of current order status (executable or complete), same as order projections
type OrderStatusVal baseEnumVal

// PersistenceTypeVal Enum of what to do with order when market turns Inplay
type PersistenceTypeVal baseEnumVal

// SideVal Enum of side, back or lay
type SideVal baseEnumVal

// Constants for side, back or lay
const (
	SideBack SideVal = "BACK"
	SideLay          = "LAY"
)

// Constant values for use in order projections
const (
	OrderProjectionAll               OrderProjVal = "ALL"
	OrderProjectionExecutable                     = "EXECUTABLE"
	OrderProjectionExecutionComplete              = "EXECUTION_COMPLETE"
)

// Constant values for order status
const (
	OrderStatusExecutionComplete OrderStatusVal = OrderProjectionExecutionComplete
	OrderStatusExecutable                       = OrderProjectionExecutable
)

// Constant values for use in match projections
const (
	MatchProjectionNoRollup           MatchProjVal = "NO_ROLLUP"
	MatchProjectionRolledUpByPrice                 = "ROLLED_UP_BY_PRICE"
	MatchProjectionRolledUpByAvgPrice              = "ROLLED_UP_BY_AVG_PRICE"
)

// Constant values for use in market projections
const (
	MarketProjectionCompetition       MarketProjVal = "COMPETITION"
	MarketProjectionEvent                           = "EVENT"
	MarketProjectionEventType                       = "EVENT_TYPE"
	MarketProjectionMarketStartTime                 = "MARKET_START_TIME"
	MarketProjectionMarketDescription               = "MARKET_DESCRIPTION"
	MarketProjectionRunnerDescription               = "RUNNER_DESCRIPTION"
	MarketProjectionRunnerMetadata                  = "RUNNER_METADATA"
)

// Constant values for price data projections
const (
	PriceDataSPAvailable  PriceDataVal = "SP_AVAILABLE"
	PriceDataSPTraded                  = "SP_TRADED"
	PriceDataEXBestOffers              = "EX_BEST_OFFERS"
	PriceDataEXAllOffers               = "EX_ALL_OFFERS"
	PriceDataEXTraded                  = "EX_TRADED"
)

// Constant values representing runner status
const (
	RunnerStatusActive        RunnerStatusVal = "ACTIVE"
	RunnerStatusWinner                        = "WINNER"
	RunnerStatusLoser                         = "LOSER"
	RunnerStatusRemovedVacant                 = "REMOVED_VACANT"
	RunnerStatusRemoved                       = "REMOVED"
	RunnerStatusHidden                        = "HIDDEN"
)

// Constant values for what to do with order when it turns in play
const (
	PersistenceTypeLapse         PersistenceTypeVal = "LAPSE"
	PersistenceTypePersist                          = "PERSIST"
	PersistenceTypeMarketOnClose                    = "MARKET_ON_CLOSE"
)

// Constant values for types of order
const (
	OrderTypeLimit         OrderTypeVal = "LIMIT"
	OrderTypeLimitOnClose               = "LIMIT_ON_CLOSE"
	OrderTypeMarketOnClose              = "MARKET_ON_CLOSE"
)

// ProjectionParams contains the various projections
// for assigning to requests
type ProjectionParams struct {
	MarketProjection []MarketProjVal
	PriceProjection  *PriceProjection
	OrderProjection  OrderProjVal
	MatchProjection  MatchProjVal
}

// MarketFilter allows various filtering of market types
type MarketFilter struct {
	TextQuery       string   `json:"textQuery,omitempty"`
	ExchangeIds     []string `json:"exchangeIds,omitempty"`
	EventTypeIds    []string `json:"eventTypeIds,omitempty"`
	EventIds        []string `json:"eventIds,omitempty"`
	CompetitionIds  []string `json:"competitionIds,omitempty"`
	MarketCountries []string `json:"marketCountries,omitempty"`
	MarketIds       []string `json:"marketIds,omitempty"`
	MarketTypeCodes []string `json:"marketTypeCodes,omitempty"`
}

// PriceProjection sets data returned from price queries
type PriceProjection struct {
	PriceData []PriceDataVal `json:"priceData,omitempty"`
}

// Params sets up the required parameters for betfair requests
type Params struct {
	MarketFilter     *MarketFilter    `json:"filter,omitempty"`
	MarketIds        []string         `json:"marketIds,omitempty"`
	PriceProjection  *PriceProjection `json:"priceProjection,omitempty"`
	MarketProjection []MarketProjVal  `json:"marketProjection,omitempty"`
	OrderProjection  OrderProjVal     `json:"orderProjection,omitempty"`
	MatchProjection  MatchProjVal     `json:"matchProjection,omitempty"`
	MaxResults       int              `json:"maxResults,omitempty"`
	Locale           string           `json:"locale,omitempty"`
}

// SetProjections applies the projections from a param object to the general
// params object
func (p *Params) SetProjections(params *ProjectionParams) {
	p.PriceProjection = params.PriceProjection
	p.MarketProjection = params.MarketProjection
	p.OrderProjection = params.OrderProjection
	p.MatchProjection = params.MatchProjection
}

// EventType represents the name and betfair ID of events
type EventType struct {
	ID   string
	Name string
}

// EventTypeResult Response for query on event types
type EventTypeResult struct {
	EventType   *EventType
	MarketCount int
}

type Competition struct {
	Id   string
	Name string
}

type CompetitionResult struct {
	Competition       *Competition
	MarketCount       int
	CompetitionRegion string
}

type CountryCodeResult struct {
	CountryCode string
	MarketCount int
}

type Event struct {
	Id          string
	Name        string
	CountryCode string
	Timezone    string
	Venue       string
	OpenDate    time.Time
}

type EventResult struct {
	Event       *Event
	MarketCount int
}

// PriceSize gives price and size of stake
type PriceSize struct {
	Price float32 `json:"price,omitempty"`
	Size  float32 `json:"size,omitempty"`
}

// StartingPrices Information about the Betfair Starting Price. Only available in BSP markets
type StartingPrices struct {
	NearPrice         float32     `json:"nearPrice,omitempty"`
	FarPrice          float32     `json:"farPrice,omitempty"`
	BackStakeTaken    []PriceSize `json:"backStakeTaken,omitempty"`
	LayLiabilityTaken []PriceSize `json:"layLiabilityTaken,omitempty"`
	ActualSP          float32     `json:"actualSP,omitempty"`
}

// ExchangePrices Prices available to back and lay with volume
type ExchangePrices struct {
	AvailableToBack []PriceSize `json:"availableToBack,omitempty"`
	AvailableToLay  []PriceSize `json:"availableToLay,omitempty"`
	TradedVolume    []PriceSize `json:"TradedVolume,omitempty"`
}

// Order Struct representing a market order
type Order struct {
	BetId           string
	OrderType       OrderTypeVal
	Status          OrderStatusVal
	PersistenceType PersistenceTypeVal
	Side            SideVal
	Price           float32
	Size            float32
	BspLiability    float32
	PlacedDate      time.Time
	AvgPriceMatched float32
	SizeMatched     float32
	SizeRemaining   float32
	SizeLapsed      float32
	SizeCancelled   float32
	SizeVoided      float32
}

//Match An individual bet Match, or rollup by price or avg price. Rollup depends on the requested MatchProjection
type Match struct {
	BetID     string    `json:"betId,omitempty"`
	MatchDate time.Time `json:"matchDate,omitempty"`
	MatchID   string    `json:"matchId,omitempty"`
	Price     string    `json:"price"`
	Side      SideVal   `json:"side"`
	Size      float32   `json:"size"`
}

// MarketBook showing data on a specific market
type MarketBook struct {
	MarketId              string
	IsMarketDataDelayed   bool
	Status                string
	BetDelay              int
	BspReconciled         bool
	Complete              bool
	Inplay                bool
	NumberOfWinners       int
	NumberOfRunners       int
	NumberOfActiveRunners int
	LastMatchTime         time.Time
	TotalMatched          float32
	TotalAvailable        float32
	CrossMatching         bool
	RunnersVoidable       bool
	Version               int
	Runners               []Runner
}

// Runner Details for each runner, containing current bets
type Runner struct {
	SelectionID      uint32          `json:"selectionId,omitempty"`
	Handicap         float32         `json:"handicap,omitempty"`
	Status           RunnerStatusVal `json:"status,omitempty"`
	AdjustmentFactor float32         `json:"adjustmentFactor,omitempty"`
	LastPriceTraded  float32         `json:"lastPriceTraded,omitempty"`
	TotalMatched     float32         `json:"totalMatched,omitempty"`
	RemovalDate      time.Time       `json:"removalDate,omitempty"`
	StartingPrices   StartingPrices  `json:"sp,omitempty"`
	ExchangePrices   ExchangePrices  `json:"ex,omitempty"`
	Orders           []Order         `json:"orders,omitempty"`
	Matches          []Match         `json:"matches,omitempty"`
}

// RunnerCatalog Information about the Runners (selections) in a market.
type RunnerCatalog struct {
	SelectionId  uint32
	RunnerName   string
	Handicap     float32
	SortPriority int
	Metadata     map[string]string
}

// MarketCatalogue Information about a market.
type MarketCatalogue struct {
	MarketId        string
	MarketName      string
	MarketStartTime time.Time
	Description     *MarketDescription
	TotalMatched    float32
	Runners         []RunnerCatalog
	EventType       *EventType
	Competition     *Competition
	Event           *Event
}

// Market definition.
type MarketDescription struct {
	PersistenceEnabled bool
	BspMarket          bool
	MarketTime         time.Time
	SuspendTime        time.Time
	SettleTime         time.Time
	BettingType        string
	TurnInPlayEnabled  bool
	MarketType         string
	Regulator          string
	MarketBaseRate     float32
	DiscountAllowed    bool
	Wallet             string
	Rules              string
	RulesHasDate       bool
	Clarifications     string
}

// MarketType Result.
type MarketTypeResult struct {
	MarketType  string
	MarketCount int
}

// Returns a list of Competitions (i.e., World Cup 2013) associated with the
// markets selected by the MarketFilter.
func (s *Session) ListCompetitions(filter *MarketFilter) ([]CompetitionResult, error) {
	var results []CompetitionResult
	params := new(Params)
	params.MarketFilter = filter
	err := doBettingRequest(s, "listCompetitions", params, &results)
	return results, err
}

// Returns a list of Countries associated with the markets selected by the
// MarketFilter.
func (s *Session) ListCountries(filter *MarketFilter) ([]CountryCodeResult, error) {
	var results []CountryCodeResult
	params := new(Params)
	params.MarketFilter = filter
	err := doBettingRequest(s, "listCountries", params, &results)
	return results, err
}

// Returns a list of Events (i.e, Reading vs. Man United) associated with the
// markets selected by the MarketFilter.
func (s *Session) ListEvents(filter *MarketFilter) ([]EventResult, error) {
	var results []EventResult
	params := new(Params)
	params.MarketFilter = filter
	err := doBettingRequest(s, "listEvents", params, &results)
	return results, err
}

// Returns a list of Event Types (i.e. Sports) associated with the markets
// selected by the MarketFilter.
func (s *Session) ListEventTypes(filter *MarketFilter) ([]EventTypeResult, error) {
	var results []EventTypeResult
	params := new(Params)
	params.MarketFilter = filter
	err := doBettingRequest(s, "listEventTypes", params, &results)
	return results, err
}

// ListMarketBook Returns a list of dynamic data about markets. Dynamic data includes prices,
// the status of the market, the status of selections, the traded volume, and
// the status of any orders you have placed in the market.
func (s *Session) ListMarketBook(marketIds []string, projectionsParam *ProjectionParams) ([]MarketBook, error) {
	var results []MarketBook
	params := new(Params)
	params.MarketIds = marketIds
	params.SetProjections(projectionsParam)
	err := doBettingRequest(s, "listMarketBook", params, &results)
	return results, err
}

// ListMarketCatalogue Returns a list of information about markets that does not change (or
// changes very rarely). You use listMarketCatalogue to retrieve the name
// of the market, the names of selections and other information about markets.
// Market Data Request Limits apply to requests made to listMarketCatalogue.
func (s *Session) ListMarketCatalogue(filter *MarketFilter, maxResults int, projectionsParam *ProjectionParams) ([]MarketCatalogue, error) {
	var results []MarketCatalogue
	params := new(Params)
	params.MarketFilter = filter
	params.MaxResults = maxResults
	params.SetProjections(projectionsParam)
	err := doBettingRequest(s, "listMarketCatalogue", params, &results)
	return results, err
}

// Returns a list of market types (i.e. MATCH_ODDS, NEXT_GOAL) associated
// with the markets selected by the MarketFilter. The market types are always
// the same, regardless of locale.
func (s *Session) ListMarketTypes(filter *MarketFilter) ([]MarketTypeResult, error) {
	var results []MarketTypeResult
	params := new(Params)
	params.MarketFilter = filter
	err := doBettingRequest(s, "listMarketTypes", params, &results)
	return results, err
}

func doBettingRequest(s *Session, method string, params *Params, v interface{}) error {

	params.Locale = s.config.Locale

	bytes, err := json.Marshal(params)
	if err != nil {
		return err
	}
	body := strings.NewReader(string(bytes))

	data, err := doRequest(s, "betting", method+"/", body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, v); err != nil {
		return err
	}

	return nil
}
