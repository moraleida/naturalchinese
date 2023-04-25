package feedReader

import (
	"cloud.google.com/go/storage"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/mmcdole/gofeed"
	"log"
)

func init() {
	functions.CloudEvent("IngestFeed", ingestFeed)
}

func getBucket(ctx context.Context) (*storage.BucketHandle, string) {
	storageclient, storageerr := storage.NewClient(ctx)

	if storageerr != nil {
		log.Fatal(storageerr)
	}

	bucketName := "natural-chinese-raw-files"

	return storageclient.Bucket(bucketName), bucketName
}

func getObjectHashString(el *gofeed.Item) string {
	jsonObj, _ := json.Marshal(el)
	objHash := md5.Sum(jsonObj)
	return hex.EncodeToString(objHash[:])
}

func writeObject(hash string, ctx context.Context, bucket *storage.BucketHandle) {
	newObj := bucket.Object(hash)
	w := newObj.NewWriter(ctx)

	decoded, decodeerr := hex.DecodeString(hash)

	if decodeerr != nil {
		log.Fatal(decodeerr)
	}

	if _, writerr := fmt.Fprintf(w, "%q", string(decoded)); writerr != nil {
		log.Fatal(writerr)
	}

	if closeerr := w.Close(); closeerr != nil {
		log.Fatal(closeerr)
	}
}

// Entry point for Cloud Function
func ingestFeed(ctx context.Context, e event.Event) error {
	//	feedURL := ingest(e.Data())
	fp := gofeed.NewParser()
	//	feed, _ := fp.ParseURL(feedURL)
	feed, _ := fp.ParseURL("http://www.xinhuanet.com/politics/news_politics.xml")
	bucket, bucketName := getBucket(ctx)

	for _, element := range feed.Items {
		hashString := getObjectHashString(element)
		_, attrserr := bucket.Object(hashString).Attrs(ctx)

		if attrserr != storage.ErrObjectNotExist {
			continue
		}

		writeObject(hashString, ctx, bucket)
		log.Println("INFO: stored " + element.Link + "|" + hashString + "|" + bucketName)
	}

	return nil
}
