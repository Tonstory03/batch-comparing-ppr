package elasticclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"th.truecorp.it.dsm.batch/batch-comparing-ppr/utils"
)

func search(es *elasticsearch.Client, params []func(*esapi.SearchRequest)) (*ResultSearch, error) {

	var result ResultSearch

	// params := []func(*esapi.SearchRequest){es.Search.WithBody(&buffer), es.Search.WithIndex(*indexName)}

	// json.NewEncoder(&buffer).Encode(body)

	response, err := es.Search(params...)

	if err != nil {
		return nil, err
	}

	json.NewDecoder(response.Body).Decode(&result)

	return &result, nil
}

func SearchBodyMap(es *elasticsearch.Client, body map[string]interface{}, indexName *string) (*ResultSearch, error) {

	var buffer bytes.Buffer

	json.NewEncoder(&buffer).Encode(body)

	params := []func(*esapi.SearchRequest){es.Search.WithBody(&buffer)}

	if indexName != nil {
		params = append(params, es.Search.WithIndex(*indexName))
	}

	return search(es, params)
}

func SearchBodyStr(es *elasticsearch.Client, body string, indexName *string) (*ResultSearch, error) {

	params := []func(*esapi.SearchRequest){es.Search.WithBody(bytes.NewBuffer([]byte(body)))}

	if indexName != nil {
		params = append(params, es.Search.WithIndex(*indexName))
	}

	return search(es, params)
}

// func getSearchBodyProcessFailure(env, startTime, endTime string) (map[string]interface{}, error) {

// 	return getSearchBodyRetryProcess(env, startTime, endTime, nil)
// }

// func getSearchBodyProcessFailurePaging(env, startTime, endTime string, from, size int) (map[string]interface{}, error) {

// 	return getSearchBodyRetryProcess(env, startTime, endTime, &SearchPaging{From: &from, Size: &size})
// }

// func GetSearchBodyProcessFailure(env, startTime, endTime string, from, size int) (map[string]interface{}, error) {

// 	return getSearchBodyRetryProcess(env, startTime, endTime, SearchRequest{From: &from, Size: &size})
// }

// func getSearchBodyProcessFailure(env, startTime, endTime string, result SearchRequest) (map[string]interface{}, error) {
// 	var query interface{} = fmt.Sprintf(QUERY_BODY_PROCESS_FAILURE, getTopic(env), startTime, endTime)
// 	var sort interface{} = `{"kafkaTimestamp": "asc"}`
// 	var from interface{} = 0
// 	var size interface{} = 9999
// 	qs := getQuery(&query, &sort, &from, &size)

// }

func getQuery(query, sort, from, size *interface{}) string {

	listKV := []KV{
		KV{Key: "query", Value: utils.ToJsonText(query)},
		KV{Key: "sort", Value: utils.ToJsonText(sort)},
		KV{Key: "from", Value: utils.ToJsonText(from)},
		KV{Key: "size", Value: utils.ToJsonText(size)},
	}

	return fmt.Sprintf(`{%s}`, convertKV2String(listKV))
}

func convertKV2String(keyValues []KV) string {

	var arr []string = make([]string, 0)

	for _, keyValue := range keyValues {
		if v := keyValue.Convert2String(); v != nil {
			arr = append(arr, *v)
		}
	}

	return strings.Join(arr, ",")
}
