package cronjob

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"th.truecorp.it.dsm.batch/batch-comparing-ppr/elasticclient"

	"th.truecorp.it.dsm.batch/batch-comparing-ppr/utils"
)

var isRetryProcessKafkaRunning = false

func retryProcessKafka() {

	if isRetryProcessKafkaRunning {
		return
	}

	now := time.Now()

	minute := int(math.Floor(float64(now.Minute()/30)) * 30)

	lastTimeCommitProcess := utils.Time2StrFormatISO8601Timezone(time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), minute, 0, 0, time.Local))

	fmt.Println(lastTimeCommitProcess)

	// uuidVal := uuid.New().String()

	isRetryProcessKafkaRunning = true

	fmt.Println("Start retry process kafka ", time.Now())

	// es, err := elasticclient.NewClient()

	// if err != nil {
	// 	fmt.Println("Error NewClient:", err.Error())
	// }

	// resultQueryFailure, err := searchProcessFailure(es, profile, uuidVal)

	// // value not found
	// if resultQueryFailure.Hits.Total.Value == 0 {
	// 	return
	// }

	// countRetry := 0

	// for _, dataHits := range resultQueryFailure.Hits.Hits {

	// }

	// fmt.Println("Finish retry process count retry process:", countRetry, ", no retry :", (len(resultQueryFailure.Hits.Hits) - countRetry), " at ", time.Now())
	isRetryProcessKafkaRunning = false
}

func searchProcessByCodeName(es *elasticsearch.Client, env, offerCode, offerName string, kafkaTimestamp int64) (*elasticclient.ResultSearch, error) {

	var (
		from = utils.NewIntPointer(0)
		size = utils.NewIntPointer(9999)
	)

	query, err := getSearchBodyProcessByCodeName(env, offerCode, offerName, kafkaTimestamp, &elasticclient.SearchPaging{From: from, Size: size})

	if err != nil {
		return nil, err
	}

	return elasticclient.Search(es, query)
}

func getTopic(env string) string {
	return fmt.Sprintf("%s-cat-offer", env)
}

func getSearchBodyProcessFailure(env string, searchRangeTimestamp elasticclient.SearchRangeTimestamp, searchPaging *elasticclient.SearchPaging) (map[string]interface{}, error) {

	var result elasticclient.SearchRequest = elasticclient.SearchRequest{}

	// Check and add from size.
	if searchPaging != nil && searchPaging.From != nil && searchPaging.Size != nil {
		result.From = searchPaging.From
		result.Size = searchPaging.Size
	}

	// Sort by kafka timestamp
	result.Sort = map[string]string{
		"kafkaTimestamp": "asc",
	}

	// Add query.
	var query string = fmt.Sprintf(elasticclient.QUERY_BODY_PROCESS_FAILURE, getTopic(env), searchRangeTimestamp.StartTime, searchRangeTimestamp.EndTime)
	var mapQuery map[string]interface{}

	err := json.Unmarshal([]byte(query), &mapQuery)

	if err != nil {
		return nil, err
	}

	result.Query = mapQuery

	return result.Convert2Map()
}

func getSearchBodyProcessByCodeName(env, offerCode, offerName string, kafkaTimestamp int64, searchPaging *elasticclient.SearchPaging) (map[string]interface{}, error) {

	var result elasticclient.SearchRequest = elasticclient.SearchRequest{}

	// Check and add from size.
	if searchPaging != nil && searchPaging.From != nil && searchPaging.Size != nil {
		result.From = searchPaging.From
		result.Size = searchPaging.Size
	}

	// Sort by kafka timestamp
	result.Sort = map[string]string{
		"kafkaTimestamp": "asc",
	}

	// Add query.
	var query string = fmt.Sprintf(elasticclient.QUERY_BODY_BY_CODENAME, getTopic(env), strconv.FormatInt(kafkaTimestamp, 10), offerCode, offerName)
	var mapQuery map[string]interface{}

	err := json.Unmarshal([]byte(query), &mapQuery)

	if err != nil {
		return nil, err
	}

	result.Query = mapQuery

	return result.Convert2Map()
}
