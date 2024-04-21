package es

import (
	"context"
	"encoding/json"
	"fmt"
	log "goconf/core/logger"
	"goconf/core/sdk"

	"github.com/elastic/go-elasticsearch/v8"
)

func Setup() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Println("ElasticSearch connect error :", err)
	}

	// 发送健康检查请求
	res, err := client.Cluster.Health(
		client.Cluster.Health.WithContext(context.Background()),
	)
	if err != nil {
		fmt.Println("ElasticSearch connect error :", err)
	} else {
		defer res.Body.Close()
		var healthInfo map[string]interface{}
		fmt.Println("ElasticSearch connect success !")
		if err := json.NewDecoder(res.Body).Decode(&healthInfo); err != nil {
			log.Error("Error parsing the response body:", err)
		} else {
			// 获取集群健康状态
			status, found := healthInfo["status"]
			if !found {
				log.Error("Status not found in response")
			} else {
				// 打印集群健康状态
				fmt.Println("Elasticsearch cluster status:", status)
			}
		}
	}

	sdk.Runtime.SetEsDb("*", client)

}
