FROM node:alpine as frontend-builder
WORKDIR /app
COPY frontend/ .
RUN npm install && npm run build

FROM golang:alpine as backend-builder
WORKDIR /app
COPY backend/ .
COPY --from=frontend-builder /app/dist ./frontend/dist
RUN go build -o todo-list

FROM alpine:latest
COPY --from=backend-builder /app/todo-list .
EXPOSE 8080
CMD ["./todo-list"]
