variable "aws_region" {
  type    = string
  default = "us-east-1"
}

variable "aws_az_a" {
  type    = string
  default = "us-east-1a"
}

variable "aws_az_b" {
  type    = string
  default = "us-east-1b"
}

variable "vpc_id" {
  type    = string
  default = "value"
}

variable "subnet_a" {
  type    = string
  default = "value"
}

variable "subnet_b" {
  type    = string
  default = "value"
}

variable "sg_load_balancer" {
  type    = string
  default = "value"
}

variable "ecr_image" {
  description = "ECR Image"
  type        = string
  sensitive   = true
}

variable "execution_role_ecs" {
  description = "Execution role ECS"
  type        = string
  sensitive   = true
}

variable "desired_tasks" {
  description = "Mininum executing tasks"
  type    = number
  default = 1
}

variable "ecs_cluster" {
  description = "Cluster ECS ARN"
  type = string
  sensitive = true
}

variable "sg_cluster_ecs" {
  description = "Cluster ECS Security group"
  type    = string
  default = "value"
}

variable "target_group_arn" {
  description = "Target Group ARN"
  type      = string
  sensitive = true
}

variable "db_url" {
  description = "database url"
  type = string
  sensitive = true
}