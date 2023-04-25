package feedReader

/**
 * sample data for testing purposes
 */
func getBody() []byte {
	return []byte(`{
			"message": {
			  "data": "d29ybGQ=",
			  "attributes": {
				 "attr1":"attr1-value"
			  }
			},
			"subscription": "projects/MY-PROJECT/subscriptions/MY-SUB"
		  }`)
}
