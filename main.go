package main

import (
	"fmt"

	"th.truecorp.it.dsm.batch/batch-comparing-ppr/elasticclient"
)

func main() {
	client, err := elasticclient.NewClient()
	if err != nil {
		panic(err)
	}
	indexName := "k8stdev-intcom-*"
	resultSearch, err := elasticclient.SearchBodyStr(client, elasticclient.QUERY_LEGACY_GET_ALL_PREPAID_PROFILE_LIST, &indexName)

	if err != nil {
		panic(err)
	}

	fmt.Println(resultSearch)
}
