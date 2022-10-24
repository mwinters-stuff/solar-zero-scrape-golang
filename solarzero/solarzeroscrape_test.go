package solarzero_test

import (
	// "testing"

	"testing"

	"github.com/mwinters-stuff/solar-zero-scrape-golang/solarzero"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/solarzero/jsontypes"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	// "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type scrapeLogHook struct {
	LastEvent *zerolog.Event
	LastLevel zerolog.Level
	LastMsg   string
}

func (h *scrapeLogHook) Run(e *zerolog.Event, l zerolog.Level, m string) {
	h.LastEvent = e
	h.LastLevel = l
	h.LastMsg = m
}

type SolarZeroScrapeTestSuite struct {
	suite.Suite
	loghook scrapeLogHook
}

func (*SolarZeroScrapeTestSuite) makeConfig() jsontypes.Configuration {
	config := jsontypes.Configuration{}
	config.InfluxDB.HostURL = "https://influxdb.url/"
	config.InfluxDB.Token = "ANTOKENTHATSBIG"
	config.InfluxDB.Bucket = "solarzero/autogen"
	config.InfluxDB.Org = "example.org"
	return config
}

func (suite *SolarZeroScrapeTestSuite) SetupTest() {
	suite.loghook = scrapeLogHook{}
	solarzero.Logger = log.Hook(&suite.loghook)

}

func (suite *SolarZeroScrapeTestSuite) TestNew() {
	cfg := suite.makeConfig()
	szs := solarzero.NewSolarZeroScrape(cfg)
	assert.NotNil(suite.T(), szs)

}

func (suite *SolarZeroScrapeTestSuite) TestAuthenticateFully() {
	cfg := suite.makeConfig()
	szs := solarzero.NewSolarZeroScrape(cfg)
	assert.NotNil(suite.T(), szs)

	assert.False(suite.T(), szs.AuthenticateFully())
}

func TestSolarZeroScrapeSuite(t *testing.T) {
	suite.Run(t, new(SolarZeroScrapeTestSuite))
}
