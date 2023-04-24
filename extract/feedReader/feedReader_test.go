package feedReader

import (
	"fmt"
	"testing"
)

func TestFeedReader(t *testing.T) {
	result := ingest()

	fmt.Printf("%q\n", result)
}
