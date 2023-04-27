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

resource "google_cloud_scheduler_job" "trigger_scheduler" {
  name        = "${var.project_name}-service_feedreader-trigger-scheduler"
  description = "Triggers the feedReader service"
  schedule    = "*/20 * * * *" // run every 20 minutes
  project     = var.project_id
  region      = var.region

  pubsub_target {
    topic_name = google_pubsub_topic.service_feedreader_topic.id
    data       = base64encode(formatdate("DD MMM YYYY hh:mm ZZZ", timestamp()))
  }
}