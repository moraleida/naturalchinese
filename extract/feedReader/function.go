// Package feedreader is meant to be implemented as an event-triggered Google Cloud Function.
//
// Entrypoint: ingestFeed
// Expects: PubSub Event with an RSS feed URL in the message.data field
package feedreader

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/mmcdole/gofeed"
)

// init is the default function for implementing GCP Cloud Functions.
// This should not be changed unless the ingestFeed func signature changes.
func init() {
	functions.CloudEvent("ingestFeed", ingestFeed)
}

// getBucket returns the Bucket client for storing objects
func getBucket(ctx context.Context) (*storage.BucketHandle, string) {
	storageclient, storageerr := storage.NewClient(ctx)

	if storageerr != nil {
		log.Fatal(storageerr)
	}

	bucketName := "natural-chinese-ingest-feedreader-raw-files"

	return storageclient.Bucket(bucketName), bucketName
}

// getObjectHashString returns an md5 hash of the RSS Item content meant to be used as the Object name
func getObjectHashString(el *gofeed.Item) string {
	jsonObj, _ := json.Marshal(el)
	objHash := md5.Sum(jsonObj)
	return hex.EncodeToString(objHash[:])
}

// getObjectContent returns the string representation of the json object for a specific RSS Item
func getObjectContent(el *gofeed.Item) string {
	jsonObj, _ := json.Marshal(el)
	return string(jsonObj)
}

// writeObjectToBucket creates a new Object with the content hash as its name and the text representation
// of the json object for the RSS Item as the content
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

// ingestFeed is the entry point for this Cloud Function
// This func expects a PubSub Event containing a Feed URL in the message.data field.
// It will read the contents of the feed and create any objects in the Raw Objects
// bucket with their contents.
func ingestFeed(ctx context.Context, e event.Event) error {
	// Uncomment for local testing
	// _ := os.Setenv("STORAGE_EMULATOR_HOST", "localhost:9023")
	// log.Println("LOG: EMULATOR: " + os.Getenv("STORAGE_EMULATOR_HOST"))

	//	feedURL := ingest(e.Data())
	fp := gofeed.NewParser()
	feedurl, _ := base64.StdEncoding.DecodeString(string(e.Data()))
	feed, _ := fp.ParseURL(string(feedurl))
	bucket, bucketName := getBucket(ctx)

	for _, element := range feed.Items {
		hashString := getObjectHashString(element)
		contentString := getObjectContent(element)

		_, attrserr := bucket.Object(hashString).Attrs(ctx)
		if attrserr != storage.ErrObjectNotExist {
			// Do not recreate objects if their hashes didn't change
			continue
		}

		writeObjectToBucket(hashString, contentString, ctx, bucket)
		// @TODO writeReferenceToDatabase
		log.Println("INFO: stored " + element.Link + "|" + hashString + "|" + bucketName)
	}

	return nil
}
