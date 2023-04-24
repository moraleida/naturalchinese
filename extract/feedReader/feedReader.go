package feedReader

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/mmcdole/gofeed"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.CloudEvent("IngestFeed", ingestFeed)
}

type PubSubEvent struct {
	Message struct {
		Data       string
		Attributes string
	}
	Subscription string
}

type RSSPost struct {
	Channel struct {
		Posts []struct {
			Title    string   `xml:"title"`
			Link     string   `xml:"link"`
			Category []string `xml:"category"`
			PubDate  string   `xml:"pubDate"`
			Updated  string   `xml:"updated"`
			Encoded  string   `xml:"encoded"`
			Content  string   `xml:"content"`
		} `xml:"item"`
	} `xml:"channel"`
}

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

// Decode the json object
func getJSON(data []byte) map[string]interface{} {
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(data, &jsonMap)

	if err != nil {
		panic("malformed input")
	}

	return jsonMap
}

// Parse the json object and return the expected feed URL
func ingest(data []byte) string {
	jsonBody := getJSON(data) // map[string]interface{}

	// message
	messageKeyValue := jsonBody["message"].(map[string]interface{}) // map[string]interface{}
	// data key
	dataKeyValue := messageKeyValue["data"].(string)                  // string
	decodedData, err := base64.StdEncoding.DecodeString(dataKeyValue) // []byte

	if err != nil {
		panic("unable to decode data")
	}

	return fmt.Sprintf("%q", decodedData)
}

// Entry point for Cloud Function
func ingestFeed(ctx context.Context, e event.Event) error {
	//	feedURL := ingest(e.Data())
	fp := gofeed.NewParser()
	//	feed, _ := fp.ParseURL(feedURL)
	feed, _ := fp.ParseURL("http://www.xinhuanet.com/politics/news_politics.xml")
	var count = 0

	for _, element := range feed.Items {
		jsonObj, _ := json.Marshal(element)

		fmt.Println(count)
		fmt.Printf("%s\n\n", string(jsonObj))

		count++
	}

	return nil
}
