# Liatrio Server

The Liatrio Server is a simple Go application, deployed on an AWS EKS Cluster using Terraform, and it is orchestrated with Kubernetes. It includes rate-limited API endpoints and is served in a Docker container. The project also has unit tests to ensure the functionality of the code and uses GitHub Actions for Continuous Integration/Continuous Deployment (CI/CD).

## Contents

The repository includes the following files:

- `main.go`: The main Go file of the application that serves an API endpoint on a rate-limited basis.
- `main.tf`: The Terraform configuration file to provision necessary resources on AWS.
- `main_test.go`: The Go test file to ensure the functionality of the code.
- `output.tf`: The Terraform configuration file for output variables.
- `terraform.tf`: The Terraform configuration file for required providers and versions.
- `variables.tf`: The Terraform configuration file for defining variables.
- `deployment.yaml`: The Kubernetes configuration file for deploying the application.
- `service.yaml`: The Kubernetes configuration file for creating a service for the application.
- `.github/workflows/go.yml`: The GitHub Actions workflow file for Go code operations.
- `.github/workflows/terraform.yml`: The GitHub Actions workflow file for Terraform operations.
- `.github/workflows/docker.yml`: The GitHub Actions workflow file for Docker operations.

## Prerequisites

- [Go](https://golang.org/doc/install) installed (version 1.20 or later)
- [Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli) installed (version 1.3 or later)
- Docker installed and configured on your local machine
- An account on [Docker Hub](https://hub.docker.com/)
- [AWS CLI](https://aws.amazon.com/cli/) installed and configured
- [kubectl](https://kubernetes.io/docs/tasks/tools/) installed

## How to Run

### Local Development

1. **Running the Go server locally:**

```bash
go run main.go
```
2. **Running tests locally:**

```bash
go test
```
## GitHub Actions

GitHub Actions are used for CI/CD. Three workflows are included:

- go.yml: This workflow is triggered on every push to the main branch. It checks out the code, sets up Go, installs dependencies, verifies go.sum, and runs linting and tests.
- docker.yml: This workflow is triggered when the Go Code workflow is completed. It logs into Docker Hub, builds the Docker image of the application, and pushes it to Docker Hub.
- terraform.yml: This workflow is triggered when the Docker Build/Publish workflow is completed. It sets up Terraform, formats the Terraform files, initializes Terraform, validates the Terraform files, plans and applies the Terraform configurations.

To make use of the GitHub Actions, you need to fork this repository and set up the following secrets in your GitHub repository:

- TF_API_TOKEN: Terraform Cloud API token
- AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY: AWS credentials
- DOCKERHUB_USERNAME and DOCKERHUB_TOKEN: Docker Hub credentials

## Docker

The Docker image of the application is available on Docker Hub. It can be pulled and run using the following commands:

```bash
docker pull raviteja0628/liatrioserver:1.0.0
docker run -p 8080:8080 raviteja0628/liatrioserver:1.0.0
```
This will start the server inside a Docker container and bind it to port 8080.

## Kubernetes

After the Terraform workflow successfully completes on GitHub Actions, the Kubernetes workflow will start. It will deploy the application on the EKS cluster, and create a LoadBalancer service for the application.

## Deployment to Cloud

- User will need a setup AWS account with enabled Access keys for the variables
- Terraform Cloud account to manage state files and more
- DockerHub account to hande docker images
- Setup Github Account 
