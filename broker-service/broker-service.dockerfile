FROM alpine:latest 

RUN mkdir /app

COPY brokerApp /app

CMD [ "/app/brokerApp" ]

# # base go image
# FROM golang:1.21.7-alpine as builder

# RUN mkdir /app

# COPY . /app

# WORKDIR /app

# RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

# # Build tiny docker image
# FROM alpine:latest 

# RUN mkdir /app

# COPY --from=builder /app/brokerApp /app

# CMD [ "/app/brokerApp" ]
