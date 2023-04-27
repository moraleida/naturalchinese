INTEGRATION_TEST_PATH?=./it

build:
	pack build --builder gcr.io/buildpacks/builder:v1 \
		 --env GOOGLE_APPLICATION_CREDENTIALS=/workspace/serverless_function_source_code/key.json \
		 --env GOOGLE_FUNCTION_SIGNATURE_TYPE=cloudevent \
		 --env GOOGLE_FUNCTION_TARGET=ingestFeed naturalchinese.com/ingest

run:
	docker run --rm -p 8080:8080 naturalchinese.com/ingest

test:
	./test.sh

test.start:
	docker-compose up -d --remove-orphans

test.stop:
	docker-compose down

test.integration:
	$(ENV_LOCAL_TEST) \
	go test -tags=integration $(INTEGRATION_TEST_PATH) -count=1 -run=$(INTEGRATION_TEST_SUITE_PATH)

test.integration.debug:
	$(ENV_LOCAL_TEST) \
	go test -tags=integration $(INTEGRATION_TEST_PATH) -count=1 -v -run=$(INTEGRATION_TEST_SUITE_PATH)

emulate.pubsub:
	gcloud beta emulators pubsub start --project=feedreader --host-port='localhost:8043'

emulate.cloudstorage:
	docker run -d -e PORT=9023 -p 9023:9023 -v "cloudstorage":/storage oittaa/gcp-storage-emulator

emulate.topic:
	curl -s -X PUT 'http://localhost:8043/v1/projects/feedreader/topics/mytopic'
	curl -s -X PUT 'http://localhost:8043/v1/projects/feedreader/subscriptions/mysub' \
					  -H 'Content-Type: application/json' \
					  --data '{"topic":"projects/feedreader/topics/mytopic","pushConfig": \
					  {"pushEndpoint":"http://localhost:8080/projects/feedreader/topics/mytopic"}}'