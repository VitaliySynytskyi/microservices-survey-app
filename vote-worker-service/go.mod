module github.com/VitaliySynytskyi/microservices-survey-app/vote-worker-service

go 1.15

require (
	github.com/jackc/pgx v3.6.2+incompatible
	github.com/jackc/pgx/v4 v4.8.1
	github.com/joeshaw/envdecode v0.0.0-20200121155833-099f1fc765bd
	github.com/VitaliySynytskyi/microservices-survey-app/vote-service
	github.com/rs/zerolog v1.20.0
	github.com/streadway/amqp v1.0.0
)
