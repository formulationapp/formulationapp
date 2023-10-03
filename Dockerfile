FROM node:18 as clientBuilder
WORKDIR /app
COPY client/ .
RUN yarn
RUN yarn build-only
COPY landing/ /app/dist/landing
RUN sed -i -e 's/css\//\/landing\/css\//g' /app/dist/landing/index.html
RUN sed -i -e 's/images\//\/landing\/images\//g' /app/dist/landing/index.html
RUN sed -i -e 's/js\//\/landing\/js\//g' /app/dist/landing/index.html

FROM golang AS builder
WORKDIR /app
COPY --from=clientBuilder /app/dist /app/cmd/public/
RUN ls /app/cmd/public
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /main github.com/formulationapp/formulationapp/cmd

FROM alpine
WORKDIR /
COPY --from=builder /main /main
EXPOSE 3000
ENTRYPOINT ["/main"]