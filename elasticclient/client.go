package elasticclient

import (
	"net/http"

	"github.com/elastic/go-elasticsearch/v7"
	"th.truecorp.it.dsm.batch/batch-comparing-ppr/config"
	"th.truecorp.it.dsm.batch/batch-comparing-ppr/utils"
)

func NewClient() (*elasticsearch.Client, error) {
	elasticConfig := config.GetElasticConfig()
	cfg := elasticsearch.Config{
		Addresses: []string{
			elasticConfig.Endpoint,
		},
		Header: http.Header{},
	}

	if elasticConfig.EnableAuth {
		cfg.Header["Authorization"] = []string{utils.GetBasicAuth(elasticConfig.Username, elasticConfig.Password)}
	}

	return elasticsearch.NewClient(cfg)
}
