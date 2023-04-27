// Package feedreader is meant to be implemented as an event-triggered Google Cloud Function.
//
// Entrypoint: ingestFeed
// Expects: PubSub Event with an RSS feed URL in the message.data field
package feedreader

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/mmcdole/gofeed"
)

// Default function for implementing GCP Cloud Functions.
// This should not be changed unless the ingestFeed func signature changes.
func init() {
	functions.CloudEvent("ingestFeed", ingestFeed)
}

func getBucket(ctx context.Context) (*storage.BucketHandle, string) {
	storageclient, storageerr := storage.NewClient(ctx)

	if storageerr != nil {
		log.Fatal(storageerr)
	}

	bucketName := "natural-chinese-ingest-feedreader-raw-files"

	return storageclient.Bucket(bucketName), bucketName
}

func getObjectHashString(el *gofeed.Item) string {
	jsonObj, _ := json.Marshal(el)
	objHash := md5.Sum(jsonObj)
	return hex.EncodeToString(objHash[:])
}

func getObjectContent(el *gofeed.Item) string {
	jsonObj, _ := json.Marshal(el)
	return string(jsonObj)
}

func writeObjectToBucket(hash string, content string, ctx context.Context, bucket *storage.BucketHandle) {
	newObj := bucket.Object(hash)
	w := newObj.NewWriter(ctx)

	if _, writerr := fmt.Fprintf(w, "%s", content); writerr != nil {
		log.Fatal(writerr)
	}

	if closeerr := w.Close(); closeerr != nil {
		log.Fatal(closeerr)
	}
}

/**
 *	Entry point for Cloud Function
 *  This function expects a PubSub Event containing a Feed URL in the message.data field. It will read the contents of the feed and create/update any objects in the Raw Objects bucket with their contents.
 */
func ingestFeed(ctx context.Context, e event.Event) error {
	// Uncomment for local testing
	// _ := os.Setenv("STORAGE_EMULATOR_HOST", "localhost:9023")
	// log.Println("LOG: EMULATOR: " + os.Getenv("STORAGE_EMULATOR_HOST"))

	//	feedURL := ingest(e.Data())
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("http://www.xinhuanet.com/politics/news_politics.xml")
	bucket, bucketName := getBucket(ctx)

	for _, element := range feed.Items {
		hashString := getObjectHashString(element)
		contentString := getObjectContent(element)

		_, attrserr := bucket.Object(hashString).Attrs(ctx)
		if attrserr != storage.ErrObjectNotExist {
			log.Println("LOG: STORAGE ERROR: " + attrserr.Error())
			continue
		}

		writeObjectToBucket(hashString, contentString, ctx, bucket)
		// @TODO writeReferenceToDatabase
		log.Println("INFO: stored " + element.Link + "|" + hashString + "|" + bucketName)
	}

	return nil
}
