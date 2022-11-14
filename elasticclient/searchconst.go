package elasticclient

const ACTION_VERSION_EXP = "VersionExpireHistory"

const ACTION_UPSERT = "upsert"

const ACTION_DELETE = "delete"

const ACTION_FETCHALL = "fetchAll"

const QUERY_LEGACY_GET_ALL_PREPAID_PROFILE_LIST = `{
    "from": 0,
    "size": 100,
    "sort": {
        "@timestamp": "asc"
    },
    "_source": [
        "response"
    ],
    "query": {
        "bool": {
            "must": [
                {
                    "bool": {
                        "filter": [
                            {
                                "range": {
                                    "@timestamp": {
                                        "gte": "2022-11-14T14:00:00.000+07:00",
                                        "lt": "2022-11-14T14:15:00.000+07:00"
                                    }
                                }
                            },
                            {
                                "term": {
                                    "logType": 2
                                }
                            },
                            {
                                "match": {
                                    "response": "getAllPrepaidProfileListResponse"
                                }
                            },
                            {
                                "match": {
                                    "response": "OSBbllngA00001"
                                }
                            }
                        ]
                    }
                }
            ]
        }
    }
}`

const QUERY_BODY_PROCESS_FAILURE = `{
	"bool": {
	  "must": [
		{
		  "bool": {
			"filter": [
			  {
				"term": {
				  "topic.keyword": "%s"
				}
			  },
			  {
				"range": {
				  "timestamp": {
					"gte": "%s",
					"lt": "%s"
				  }
				}
			  },
			  {
				"term": {
					"isRetryMessage": {
						"value": true
					}
				}
			  },
			  {
				"exists": {
				  "field": "action"
				}
			  }
			]
		  }
		}
	  ],
	  "must_not": [
		{
		  "bool": {
			"filter": [
			  {
				"term": {
				  "isRetryMessage": {
					"value": false
				  }
				}
			  },
			  {
				"term": {
				  "resultStatus.keyword": "F"
				}
			  }
			]
		  }
		}
	  ]
	}
}`

const QUERY_BODY_BY_CODENAME = `{
	"bool": {
	  "must": [
		{
		  "bool": {
			"filter": [
			  {
				"term": {
				  "topic.keyword": "%s"
				}
			  },
			  {
				"range": {
					"kafkaTimestamp": {
						"gt": "%s"
					}
				}
			  },
			  {
				"term": {
					"offerCode.keyword": "%s"
				}
			  },
			  {
				"term": {
					"offerName.keyword": "%s"
				}
			  },
			  {
				"exists": {
				  "field": "action"
				}
			  }
			]
		  }
		}
	  ],
	  "must_not": [
		{
		  "bool": {
			"filter": [
			  {
				"term": {
				  "isRetryMessage": {
					"value": false
				  }
				}
			  },
			  {
				"term": {
				  "resultStatus.keyword": "F"
				}
			  }
			]
		  }
		}
	  ]
	}
}`
