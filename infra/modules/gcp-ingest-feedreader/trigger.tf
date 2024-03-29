// Define the PubSub topic that will be used to trigger the Cloud Function
resource "google_pubsub_topic" "service_feedreader_topic" {
  project = var.project_id
  name    = "feedreader-topic"
}

resource "google_eventarc_trigger" "feedreader_trigger" {
  location   = var.region
  name       = "${var.project_name}-service_feedreader-trigger"

  matching_criteria {
    attribute = "type"
    value     = "google.cloud.pubsub.topic.v1.messagePublished"
  }

  destination {
    cloud_function = google_cloudfunctions2_function.service_feedreader_function.name
  }

  lifecycle {
    create_before_destroy = true
  }
}