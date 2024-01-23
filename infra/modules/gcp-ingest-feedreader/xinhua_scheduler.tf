resource "google_cloud_scheduler_job" "trigger_scheduler" {
  name        = "${var.project_name}-service_feedreader-trigger-scheduler-xinhua"
  description = "Triggers the feedReader service to read XinHua feeds"
  schedule    = "*/20 * * * *" // run every 20 minutes
  project     = var.project_id
  region      = var.region

  pubsub_target {
    topic_name = google_pubsub_topic.service_feedreader_topic.id
    data       = base64encode("${var.xinhua_feed}")
  }
}