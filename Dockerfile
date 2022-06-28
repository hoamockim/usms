ARG GO_VERSION=1.16

FROM golang:${GO_VERSION}-alpine as builder

ARG SERVICE_TYPE

ENV SERVICE_NAME="usms_$SERVICE_TYPE"
RUN echo $SERVICE_NAME

ENV MIGRATION_RESOURCES="./app/cmd/migration/resources/"

WORKDIR ./src/usms/
COPY . .

RUN go make update && \
    go build -o $SERVICE_NAME ./app/cmd/$SERVICE_TYPE/ && \
    echo $PWD

RUN mkdir -p /app
RUN cp -p ./$SERVICE_NAME /app && \
    cp -p ./entrypoint.sh /app
RUN if [ "$SERVICE_TYPE" = "migration" ]; then cp -a $MIGRATION_RESOURCES /app/resources; fi

WORKDIR /app
RUN ls -a
RUN chmod +x ./entrypoint.sh
RUN chmod +x ./$SERVICE_NAME
CMD ["./entrypoint.sh"]
