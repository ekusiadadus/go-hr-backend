package internal

import (
	"os"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"
)

func ConnectElasticsearch() (*elasticsearch.Client, error) {
	// 環境変数 ES_ADDRESS がある場合は記述されているアドレスに接続
	// ない場合は、 http://localhost:9200 に接続
	var addr string
	if os.Getenv("ES_ADDRESS") != "" {
		addr = os.Getenv("ES_ADDRESS")
	} else {
		addr = "http://localhost:9200"
	}
	cfg := elasticsearch.Config{
		Addresses: []string{
			addr,
		},
	}
	es, err := elasticsearch.NewClient(cfg)

	return es, err
}