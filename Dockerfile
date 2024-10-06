# Gunakan image golang resmi sebagai base image
FROM golang:alpine AS builder

# Atur direktori kerja di dalam container
WORKDIR /app

# Salin file go.mod dan go.sum
COPY go.mod go.sum ./
# Download dependensi
RUN go mod download

# Salin seluruh kode sumber
COPY . .
# Build aplikasi
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
# Gunakan image alpine yang lebih ringan untuk runtime
FROM alpine:latest  
# Instal ca-certificates untuk HTTPS
RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Salin binary dari stage builder
COPY --from=builder /app/main .

# Expose port yang digunakan oleh aplikasi
EXPOSE 3000

# Jalankan aplikasi
CMD ["./main"]