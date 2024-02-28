SERVICE_NAME = "std-server"
DIST_DIR = dist
MAIN_FILE = cmd/main.go

run:
	@go run ${MAIN_FILE}

clean:
	@rm -rf ${DIST_DIR}/

build: clean
	@echo "Building ${SERVICE_NAME}..."
	@go build -o "${DIST_DIR}/${SERVICE_NAME}" -a -tags netgo -ldflags '-w -extldflags "-static"' ${MAIN_FILE}
	@echo "Dist: ${DIST_DIR}/${SERVICE_NAME}"

