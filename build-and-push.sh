#!/bin/bash

set -e

# Check if a git tag is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <git-tag>"
  exit 1
fi

GIT_TAG=$1
AWS_REGION="ap-southeast-2"
ECR_REGISTRY="293875060805.dkr.ecr.ap-southeast-2.amazonaws.com"
IMAGE_NAME="any-api-backend"

# Log in to ECR
echo "Logging in to ECR..."
aws ecr get-login-password --region ${AWS_REGION} | docker login --username AWS --password-stdin ${ECR_REGISTRY}

# Build the Docker image
echo "Building Docker image..."
docker build -t ${IMAGE_NAME}:${GIT_TAG} .

# Tag the image
echo "Tagging image with ${GIT_TAG} and latest"
docker tag ${IMAGE_NAME}:${GIT_TAG} ${ECR_REGISTRY}/${IMAGE_NAME}:${GIT_TAG}
docker tag ${IMAGE_NAME}:${GIT_TAG} ${ECR_REGISTRY}/${IMAGE_NAME}:latest

# Push the image to ECR
echo "Pushing image to ECR..."
docker push ${ECR_REGISTRY}/${IMAGE_NAME}:${GIT_TAG}
docker push ${ECR_REGISTRY}/${IMAGE_NAME}:latest

echo "Successfully built and pushed ${IMAGE_NAME}:${GIT_TAG} and ${IMAGE_NAME}:latest to ECR."
