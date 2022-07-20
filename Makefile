CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})

APP_CMD_DIR=${CURRENT_DIR}/cmd

IMG_NAME=${APP}
REGISTRY=${REGISTRY}
TAG=latest
ENV_TAG=latest
NETWORK_NAME=ur_default
PROJECT_NAME=urecruit

# Including
include .build_info

go-run:
	go run cmd/main.go

run:
	docker-compose -f docker-compose.yml up --force-recreate

test:
	docker-compose -f docker-compose.test.yml up --force-recreate

config:
	docker-compose -f docker-compose.yml config

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

proto-gen:
	./scripts/gen_proto.sh ${CURRENT_DIR}
	./scripts/remove_omitempty.sh ${CURRENT_DIR}

clone-protos:
	rm -rf protos/* && cp -R ur_protos/* protos

pull-proto-module:
	git submodule update --init --recursive

update-proto-module:
	git submodule update --remote --merge

clear:
	rm -rf ${CURRENT_DIR}/bin/*

network:
	docker network create --driver=bridge ${NETWORK_NAME}

migrate-up:
	docker run --mount type=bind,source="${CURRENT_DIR}/migrations,target=/migrations" --network ${NETWORK_NAME} migrate/migrate \
		-path=/migrations/ -database=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable up

migrate-down:
	docker run --mount type=bind,source="${CURRENT_DIR}/migrations,target=/migrations" --network ${NETWORK_NAME} migrate/migrate \
		-path=/migrations/ -database=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable down

mark-as-production-image:
	docker tag ${REGISTRY}/${PROJECT_NAME}/${APP}/${IMG_NAME}:${TAG} ${REGISTRY}/${PROJECT_NAME}/${APP}/${IMG_NAME}:production
	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}/${IMG_NAME}:production

build-image:
	@docker build --rm -t ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} .
	@docker tag ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

push-image:
	@docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG}
	@docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

swag-init:
	swag init -g api/main.go -o api/docs


