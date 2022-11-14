package cronjob

import (
	"fmt"
	"math"
	"time"

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
