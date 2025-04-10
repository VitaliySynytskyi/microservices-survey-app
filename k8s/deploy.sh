#!/bin/bash

# Встановлюємо змінну для URL Docker registry
export REGISTRY_URL="your-registry-url"

# Створюємо namespace
kubectl apply -f namespace.yaml

# Створюємо ConfigMap і Secret
kubectl apply -f vote-storage/postgres-init-configmap.yaml
kubectl apply -f vote-storage/postgres-secret.yaml

# Створюємо PersistentVolumeClaims
kubectl apply -f vote-storage/postgres-pvc.yaml
kubectl apply -f survey-storage/mongo-pvc.yaml

# Розгортаємо StatefulSets для баз даних і черги повідомлень
kubectl apply -f survey-storage/mongo-statefulset.yaml
kubectl apply -f survey-storage/mongo-service.yaml
kubectl apply -f vote-storage/postgres-statefulset.yaml
kubectl apply -f vote-storage/postgres-service.yaml
kubectl apply -f vote-queue/rabbitmq-statefulset.yaml
kubectl apply -f vote-queue/rabbitmq-service.yaml

# Чекаємо поки бази даних будуть готові
echo "Чекаємо поки бази даних стануть доступними..."
kubectl wait --for=condition=Ready pod/survey-storage-0 -n survey-app --timeout=120s
kubectl wait --for=condition=Ready pod/vote-storage-0 -n survey-app --timeout=120s
kubectl wait --for=condition=Ready pod/vote-queue-0 -n survey-app --timeout=120s

# Розгортаємо сервіси
kubectl apply -f surveys/surveys-deployment.yaml
kubectl apply -f surveys/surveys-service.yaml
kubectl apply -f votes/votes-deployment.yaml
kubectl apply -f votes/votes-service.yaml
kubectl apply -f vote-worker/vote-worker-deployment.yaml
kubectl apply -f frontend/frontend-deployment.yaml
kubectl apply -f frontend/frontend-service.yaml

# Налаштовуємо Горизонтальне автомасштабування
kubectl apply -f hpa.yaml

# Налаштовуємо мережеві політики
kubectl apply -f network-policies.yaml

# Налаштовуємо Ingress
kubectl apply -f ingress.yaml

echo "Розгортання завершено!"
echo "Щоб перевірити статус подів, виконайте: kubectl get pods -n survey-app"
echo "Щоб перевірити статус сервісів, виконайте: kubectl get svc -n survey-app"
echo "Додайте запис в /etc/hosts: 127.0.0.1 survey-app.local" 