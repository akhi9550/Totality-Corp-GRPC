FROM golang:1.21.2-alpine3.18 AS build-stage
WORKDIR /grpc_user_svc
COPY ./ /grpc_user_svc
RUN mkdir -p /grpc_user_svc/build
RUN go mod download
RUN go build -v -o /grpc_user_svc/build/api ./cmd


FROM gcr.io/distroless/static-debian11
COPY --from=build-stage /grpc_user_svc/build/api /
COPY --from=build-stage /grpc_user_svc/.env /
EXPOSE 50051
CMD ["/api"]