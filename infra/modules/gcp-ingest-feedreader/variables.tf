variable "project_name" {
  type = string
  default = "natural-chinese"
}

variable "project_id" {
  type = string
  default = "natural-chinese-384616"
}

variable "region" {
  type = string
  default  = "us-central1"
}

variable "zone" {
  type = string
  default    = "us-central1-c"
}

variable "module_name" {
  type = string
  default = "default"
}

variable "xinhua_feed" {
  type = string
  default = "http://www.xinhuanet.com/politics/news_politics.xml"
}