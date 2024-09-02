GOOS=linux
GOARCH=amd64
CGO_ENABLED=0
BUILDVARS=GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=${CGO_ENABLED}
BUILD_CMD=${BUILDVARS} go build ${LDFLAGS}
BUILD_DIR=build

COMPOSE_TEST_CMD=docker-compose -p dev-${PROJECTNAME} -f docker-compose.yaml
COMPOSE_DEBUG_CMD=docker-compose -p dev-${PROJECTNAME} -f docker-compose-debug.yaml

PROJECTNAME=$(shell basename "$(PWD)")

GREEN=\033[0;32m
YELLOW=\033[0;33m
ORANGE=\033[0;31m
NC=\033[0m


build-tiny-url:
	@echo ">${YELLOW} building tiny-url...${NC}"
	${BUILD_CMD} -o ${BUILD_DIR}/tiny cmd/main.go
	@echo ">${GREEN} tiny-url is built${NC}"

compose-debug-up: build-tiny-url
	@echo ">${YELLOW} docker-compose debug env up...${NC}"
	${COMPOSE_DEBUG_CMD} up --build -d
	@echo ">${GREEN} up${NC}"

compose-debug-down:
	@echo ">${YELLOW} docker compose down...${NC}"
	${COMPOSE_DEBUG_CMD} down
	@echo ">${GREEN} down${NC}"

compose-test-up: build-tiny-url
	@echo ">${YELLOW} docker-compose up...${NC}"
	${COMPOSE_TEST_CMD} up --build -d
	@echo ">${GREEN} up${NC}"

compose-test-down:
	@echo ">${YELLOW} docker compose down...${NC}"
	${COMPOSE_TEST_CMD} down
	@echo ">${GREEN} down${NC}"
