# Stage 1: Build React frontend
FROM node:14 as frontend-builder
WORKDIR /frontend
COPY frontend/package*.json ./
RUN npm install
COPY frontend/ ./
ENV REACT_APP_BASE_URL=/api
RUN npm run build

# Stage 2: Build Go backend
FROM golang:1.23 as backend-builder
WORKDIR /backend
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/ .
COPY --from=frontend-builder /frontend/build /root/public
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

# Stage 3: Final stage
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=backend-builder /backend/main .
COPY --from=backend-builder /root/public ./public
EXPOSE 8080
CMD ["./main"]