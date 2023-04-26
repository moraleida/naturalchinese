resource "google_storage_bucket" "raw_files" {
  name                        = "${var.module_name}-raw-files"
  location                    = var.region
  project                     = var.project_id
  uniform_bucket_level_access = true
}

resource "google_storage_bucket" "service_feedreader_storage_bucket" {
  name                        = "${var.module_name}-source"
  uniform_bucket_level_access = true
  location                    = var.region
  project                     = var.project_id
}
