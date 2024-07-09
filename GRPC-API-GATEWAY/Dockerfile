FROM golang:1.21.2-alpine3.18 AS build-stage
WORKDIR /grpc_api_gateway
COPY ./ /grpc_api_gateway
RUN mkdir -p /grpc_api_gateway/build
RUN go mod download
RUN go build -v -o /grpc_api_gateway/build/api ./cmd


FROM gcr.io/distroless/static-debian11
COPY --from=build-stage /grpc_api_gateway/build/api /
COPY --from=build-stage /grpc_api_gateway/.env /
EXPOSE 3000
CMD ["/api"]

