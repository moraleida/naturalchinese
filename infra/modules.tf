module "gcp-ingest-feedreader" {
  source = "./modules/gcp-ingest-feedreader"

  module_name = "${var.project_name}-ingest-feedreader"

  depends_on  = [
    google_project_service.cloud_run_api,
    google_project_service.cloud_scheduler_api,
    google_project_service.artifact_registry,
    google_project_service.eventarc_api
  ]
}

