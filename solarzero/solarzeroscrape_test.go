package solarzero_test

import (
	// "testing"

	"testing"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	mocks "github.com/mwinters-stuff/solar-zero-scrape-golang/internal/mocks/app"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/solarzero"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/solarzero/jsontypes"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	// "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

	mockInfluxDBWriter *mocks.InfluxDBWriter
	mocksAWSInterface  *mocks.AWSInterface
}

func (*SolarZeroScrapeTestSuite) makeConfig() solarzero.AllSolarZeroOptions {
	config := solarzero.AllSolarZeroOptions{}
	config.InfluxDBOptions.HostURL = "https://influxdb.url/"
	config.InfluxDBOptions.Token = "ANTOKENTHATSBIG"
	config.InfluxDBOptions.Bucket = "solarzero/autogen"
	config.InfluxDBOptions.Org = "example.org"
	return config
}

func (suite *SolarZeroScrapeTestSuite) SetupTest() {
	suite.loghook = scrapeLogHook{}
	solarzero.Logger = log.Hook(&suite.loghook)
	suite.mockInfluxDBWriter = mocks.NewInfluxDBWriter(suite.T())
	suite.mocksAWSInterface = mocks.NewAWSInterface(suite.T())

	solarzero.NewInfluxDBWriter = func(config *jsontypes.Configuration) solarzero.InfluxDBWriter {
		return suite.mockInfluxDBWriter
	}

	solarzero.NewAWSInterface = func(config *jsontypes.Configuration) solarzero.AWSInterface {
		return suite.mocksAWSInterface
	}

}

func (suite *SolarZeroScrapeTestSuite) TestNew() {
	cfg := suite.makeConfig()

	suite.mockInfluxDBWriter.EXPECT().Connect(mock.Anything).Run(func(client influxdb2.Client) {
		assert.Equal(suite.T(), "https://influxdb.url/", client.ServerURL())
	}).Return(nil)

	szs := solarzero.NewSolarZeroScrape(&cfg)
	assert.NotNil(suite.T(), szs)

}

func (suite *SolarZeroScrapeTestSuite) TestAuthenticateFullyFailsAws() {
	cfg := suite.makeConfig()

	suite.mockInfluxDBWriter.EXPECT().Connect(mock.Anything).Return(nil)

	suite.mocksAWSInterface.EXPECT().Authenticate().Return(false)

	szs := solarzero.NewSolarZeroScrape(&cfg)
	assert.NotNil(suite.T(), szs)

	assert.False(suite.T(), szs.AuthenticateFully())
}

func TestSolarZeroScrapeSuite(t *testing.T) {
	suite.Run(t, new(SolarZeroScrapeTestSuite))
}
