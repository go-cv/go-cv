# Stage 1 · go builder
ARG GO_BUILDER=golang
ARG GO_VERSION=latest
FROM ${GO_BUILDER}:${GO_VERSION} AS build

ARG GO_OS
ARG GO_ARCH
ARG GIT_HOST
ARG REPO_ORG
ARG REPO_NAME
ARG APP_VERSION

# Copy the project inside the builder container
WORKDIR $GOPATH/src/${GIT_HOST}/$REPO_ORG/$REPO_NAME/
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=${GO_OS} GOARCH=${GO_ARCH} \
    go build \
    -installsuffix cgo \
    -ldflags="-w -s -X 'main.APP_VERSION=${APP_VERSION}' -X 'main.COMMIT_ID=$(git log HEAD --oneline | awk '{print $1}' | head -n1)'" \
    -o /gocv

# Stage 2 · scratch image
FROM scratch

# Copy the necessary stuff from the build stage
COPY --from=build /gocv /gocv
# Copy the certificates - in case of fetches
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/cert.pem
# Copy themes for templates
COPY --from=build $GOPATH/src/${GIT_HOST}/${REPO_ORG}/${REPO_NAME}/themes /themes
# Copy default config
COPY --from=build $GOPATH/src/${GIT_HOST}/${REPO_ORG}/${REPO_NAME}/config.yaml /config.yaml

# Create content and output directories
# Note: In scratch image, we need to mount these at runtime

# Execute the binary
ENTRYPOINT ["/gocv", "serve"]
