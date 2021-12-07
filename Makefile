GCP_PROJECT := $(shell gcloud config get-value project)
SERVICE_NAME := grpc-calculator
REGION := us-west1
ENDPOINT := $(shell gcloud run services list \
            	--project=${GCP_PROJECT} \
            	--region=${REGION} \
            	--platform=managed \
            	--format="value(status.address.url)" \
            	--filter="metadata.name=grpc-calculator")

#ENDPOINT=$(\
#  gcloud run services list \
#  --project=${GCP_PROJECT} \
#  --region=${GCP_REGION} \
#  --platform=managed \
#  --format="value(status.address.url)" \
#  --filter="metadata.name=grpc-calculator")
#ENDPOINT=${ENDPOINT#https://} && echo ${ENDPOINT}

install:
	brew install grpcurl

pb_gen:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	protos/*.proto

build:
	docker build -t ${SERVICE_NAME} .
	docker tag ${SERVICE_NAME} gcr.io/${GCP_PROJECT}/${SERVICE_NAME}:latest # github.shaとかコミットハッシュつけたほうがよい
	docker push gcr.io/${GCP_PROJECT}/${SERVICE_NAME}:latest # github.shaとかコミットハッシュつけたほうがよい

deploy:
	gcloud run deploy ${SERVICE_NAME} --image gcr.io/${GCP_PROJECT}/${SERVICE_NAME}:latest \
		--project ${GCP_PROJECT} \
		--platform managed \
		--region ${REGION} \
		--allow-unauthenticated
#		--service-account go-boiler-api@${GCP_PROJECT}.iam.gserviceaccount.com \

show-grpc-endpoint:
	gcloud run services list \
	--project=${GCP_PROJECT} \
	--region=${REGION} \
	--platform=managed \
	--format="value(status.address.url)" \
	--filter="metadata.name=grpc-calculator"

ping:
	grpcurl \
		-proto protos/calculator.proto \
		-d '{"first_operand": 2.0, "second_operand": 3.0, "operation": "ADD"}' \
		localhost:443 \
		Calculator/Calculate
#		grpc-calculator-6qulfhgl6a-uw.a.run.app:443 \

client:
	go run client/*.go --grpc_endpoint grpc-calculator-6qulfhgl6a-uw.a.run.app:443
