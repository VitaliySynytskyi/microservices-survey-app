# microservices-survey-app

## Розгортання в Kubernetes

Цей проект можна розгорнути в Kubernetes кластері, дотримуючись наступних кроків:

### Передумови

1. Діючий Kubernetes кластер (minikube, kind, або інший)
2. Встановлений kubectl
3. Docker registry (Docker Hub, Google Container Registry, або інший)

### Кроки розгортання

1. Клонуйте репозиторій:
   ```bash
   git clone https://github.com/yourusername/microservices-survey-app.git
   cd microservices-survey-app
   ```

2. Оновіть URL Docker registry в скриптах:
   ```bash
   # Відкрийте файли k8s/build-images.sh та k8s/deploy.sh
   # Замініть "your-registry-url" на URL вашого Docker registry
   ```

3. Зберіть та відправте образи в Docker registry:
   ```bash
   cd k8s
   chmod +x build-images.sh
   ./build-images.sh
   ```

4. Розгорніть додаток в Kubernetes:
   ```bash
   chmod +x deploy.sh
   ./deploy.sh
   ```

5. Перевірте стан розгортання:
   ```bash
   kubectl get pods -n survey-app
   kubectl get svc -n survey-app
   ```

6. Для доступу до додатку ззовні, додайте запис в файл /etc/hosts:
   ```
   127.0.0.1 survey-app.local
   ```

7. Відкрийте у браузері: http://survey-app.local

### Структура K8s маніфестів

```
k8s/
├── namespace.yaml                      # Namespace для додатку
├── frontend/                           # Frontend сервіс
│   ├── frontend-deployment.yaml
│   └── frontend-service.yaml
├── surveys/                            # Survey сервіс
│   ├── surveys-deployment.yaml
│   └── surveys-service.yaml
├── survey-storage/                     # MongoDB для Survey сервісу
│   ├── mongo-pvc.yaml
│   ├── mongo-service.yaml
│   └── mongo-statefulset.yaml
├── votes/                              # Vote сервіс
│   ├── votes-deployment.yaml
│   └── votes-service.yaml
├── vote-worker/                        # Vote Worker сервіс
│   └── vote-worker-deployment.yaml
├── vote-queue/                         # RabbitMQ для Vote сервісу
│   ├── rabbitmq-service.yaml
│   └── rabbitmq-statefulset.yaml
├── vote-storage/                       # PostgreSQL для Vote сервісу
│   ├── postgres-init-configmap.yaml
│   ├── postgres-pvc.yaml
│   ├── postgres-secret.yaml
│   ├── postgres-service.yaml
│   └── postgres-statefulset.yaml
├── hpa.yaml                            # Horizontal Pod Autoscalers
├── network-policies.yaml               # Мережеві політики
├── ingress.yaml                        # Ingress для зовнішнього доступу
├── build-images.sh                     # Скрипт для збірки образів
└── deploy.sh                           # Скрипт для розгортання
```

### Моніторинг та логування

Для моніторингу додатку рекомендуємо встановити Prometheus та Grafana:

```bash
# Додаємо репозиторій Helm для Prometheus
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

# Встановлюємо Prometheus Stack (включає Prometheus, Grafana, AlertManager)
helm install prometheus prometheus-community/kube-prometheus-stack --namespace monitoring --create-namespace
```

Для централізованого логування рекомендуємо використовувати ELK Stack або Loki.