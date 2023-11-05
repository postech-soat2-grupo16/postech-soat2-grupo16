provider "aws" {
  region = var.aws_region
}

#Configuração do Terraform State
terraform {
  backend "s3" {
    bucket = "terraform-state-soat"
    key    = "ecs-fastfood-api/terraform.tfstate"
    region = "us-east-1"

    dynamodb_table = "terraform-state-soat-locking"
    encrypt        = true
  }
}

### Task Config ###

resource "aws_ecs_task_definition" "task_definition_ecs" {
  family                   = "task-definition-fast-food-app"
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  execution_role_arn       = var.execution_role_ecs
  task_role_arn            = var.execution_role_ecs

  cpu    = 512
  memory = 1024

  container_definitions = jsonencode([
    {
      name      = "container-1"
      image     = var.ecr_image
      cpu       = 512,
      memory    = 1024,
      essential = true,
      portMappings = [
        {
          containerPort = 8000
          hostPort      = 8000
          protocol      = "tcp"
          appProtocol   = "http"
        }
      ]

      environment = [
        { "name" : "DATABASE_URL", "value" : var.db_url }
      ]

      logConfiguration = {
        logDriver = "awslogs",
        options = {
          awslogs-create-group  = "true",
          awslogs-group         = "my-ecs-logs",
          awslogs-region        = "us-east-1",
          awslogs-stream-prefix = "awslogs-container"
        }
      },
    }
  ])

  runtime_platform {
    operating_system_family = "LINUX"
    cpu_architecture        = "X86_64"
  }

  tags = {
    infra = "task-definition-ecs"
  }
}

output "task_definition_ecs_arn" {
  value = aws_ecs_task_definition.task_definition_ecs.arn
}

resource "aws_ecs_service" "ecs_service_api_soat" {
  name                              = "ecs-service-api-soat"
  cluster                           = var.ecs_cluster
  task_definition                   = aws_ecs_task_definition.task_definition_ecs.id
  launch_type                       = "FARGATE"
  platform_version                  = "1.4.0"
  desired_count                     = var.desired_tasks
  health_check_grace_period_seconds = 30

  network_configuration {
    subnets = [
      var.subnet_a,
      var.subnet_b
    ]
    security_groups  = [var.sg_cluster_ecs]
    assign_public_ip = false
  }

  load_balancer {
    target_group_arn = var.target_group_arn
    container_name   = "container-1"
    container_port   = 8000
  }

  tags = {
    infra = "ecs-service-api"
  }
}
