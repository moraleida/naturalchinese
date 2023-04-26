INTEGRATION_TEST_PATH?=./it

build:
	pack build --builder gcr.io/buildpacks/builder:v1 --env GOOGLE_FUNCTION_SIGNATURE_TYPE=cloudevent --env GOOGLE_FUNCTION_TARGET=ingestFeed naturalchinese.com/ingest

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
