#!/bin/bash

# Встановлюємо змінну для URL Docker registry
export REGISTRY_URL="your-registry-url"

# Збираємо та відправляємо образи в реєстр
echo "Збираємо образи..."

# Frontend Service
cd ../frontend-service
docker build -t ${REGISTRY_URL}/frontend-service:latest .
docker push ${REGISTRY_URL}/frontend-service:latest

# Survey Service
cd ../survey-service
docker build -t ${REGISTRY_URL}/survey-service:latest .
docker push ${REGISTRY_URL}/survey-service:latest

# Vote Service
cd ../vote-service
docker build -t ${REGISTRY_URL}/vote-service:latest .
docker push ${REGISTRY_URL}/vote-service:latest

# Vote Worker Service
cd ../vote-worker-service
docker build -t ${REGISTRY_URL}/vote-worker-service:latest .
docker push ${REGISTRY_URL}/vote-worker-service:latest

echo "Всі образи успішно зібрано та відправлено у реєстр!" 