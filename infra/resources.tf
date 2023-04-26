// Resources that need to be activated
resource "google_project_service" "artifact_registry" {
  disable_on_destroy = false
  service            = "artifactregistry.googleapis.com"
  project            = var.project_id
}

resource "google_project_service" "cloud_build_api" {
  disable_on_destroy = false
  service            = "cloudbuild.googleapis.com"
  project            = var.project_id
}

resource "google_project_service" "cloud_functions_api" {
  disable_on_destroy = false
  service            = "cloudfunctions.googleapis.com"
  project            = var.project_id
}

resource "google_project_service" "cloud_run_api" {
  disable_on_destroy = false
  service            = "run.googleapis.com"
  project            = var.project_id
}

resource "google_project_service" "cloud_scheduler_api" {
  disable_on_destroy = false
  service            = "cloudscheduler.googleapis.com"
  project            = var.project_id
}

resource "google_project_service" "eventarc_api" {
  disable_on_destroy = false
  service            = "eventarc.googleapis.com"
  project            = var.project_id
}