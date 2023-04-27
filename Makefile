gcloud.set:
	gcloud config set project natural-chinese-384616
	gcloud config set compute/region us-central1
	gcloud config set compute/zone us-central1-c

source.zips:
	rm infra/sources/feedreader.zip
	zip -j -r infra/sources/feedreader.zip extract/feedreader/function.go extract/feedreader/go.mod