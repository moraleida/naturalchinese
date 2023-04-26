data "local_file" "source" {
  filename = "${path.cwd}/sources/feedreader.zip"
}
/*

resource "local_file" "feedreader_source" {
  filename = "${path.module}/service_feedreader_function.zip"
  source = "../../../extract/feedReader/feedreader.zip"
}
*/

// Publish the function code to the appropriate storage
resource "google_storage_bucket_object" "service_feedreader_object" {
  bucket = google_storage_bucket.service_feedreader_storage_bucket.name
  name   = "${var.project_name}-service_feedreader_source.zip"
  source = data.local_file.source.filename
  depends_on = [ google_storage_bucket.service_feedreader_storage_bucket ]
}

// Define the Cloud Function behavior
resource "google_cloudfunctions2_function" "service_feedreader_function" {
  name        = "service_feedreader_function"
  description = "Event-triggered function responding to PubSub Events to check on the RSS feeds for new texts."
  project     = var.project_id
  location    = var.region

  event_trigger {
    pubsub_topic = google_pubsub_topic.service_feedreader_topic.id
    event_type   = "google.cloud.pubsub.topic.v1.messagePublished"
  }

  build_config {
    runtime     = "go120"
    entry_point = "ingestFeed"
    source {
      storage_source {
        bucket = google_storage_bucket.service_feedreader_storage_bucket.name
        object = google_storage_bucket_object.service_feedreader_object.name
      }
    }
  }

  service_config {
    max_instance_count = 1
    available_memory   = "256M"
    timeout_seconds    = 60
  }
}