/*
resource "google_cloud_run_service" "scheduler_service" {
  location = var.region
  name     = "${var.project_name}-service_feedreader-service"

  template {
    spec {
      containers {
        image = ""
      }
    }
  }
}
*/
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
    /*
    cloud_run_service {
      service = google_cloud_run_service.scheduler_service.name
      region  = google_cloud_run_service.scheduler_service.location
    }
    */
    cloud_function = google_cloudfunctions2_function.service_feedreader_function.name
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