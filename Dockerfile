FROM golang:1.14.3-alpine AS build
WORKDIR /src
COPY . .
RUN rm .env
RUN go build
EXPOSE 6443
CMD ["./LocalGuideContentService"]