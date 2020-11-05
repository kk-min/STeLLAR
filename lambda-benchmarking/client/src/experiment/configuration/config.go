package configuration

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

//SubExperiment is the schema for all sub-experiment configurations.
type SubExperiment struct {
	Title                   string  `json:"Title"`
	Bursts                  int     `json:"Bursts"`
	BurstSizes              []int   `json:"BurstSizes"`
	PayloadLengthBytes      int     `json:"PayloadLengthBytes"`
	CooldownSeconds         float64 `json:"CooldownSeconds"`
	FunctionIncrementLimits []int64 `json:"FunctionIncrementLimits"`
	IATType                 string  `json:"IATType"`
	Provider                string  `json:"Provider"`
	GatewaysNumber          int     `json:"GatewaysNumber"`
	Visualization           string  `json:"Visualization"`
	GatewayEndpoints        []string
	ID                      int
}

//Extract will read the given JSON configuration file and load it as an array of sub-experiment configurations.
func Extract(configFile *os.File) []SubExperiment {
	configByteValue, _ := ioutil.ReadAll(configFile)

	var parsedConfigs []SubExperiment
	if err := json.Unmarshal(configByteValue, &parsedConfigs); err != nil {
		log.Error(err)
	}

	return parsedConfigs
}
