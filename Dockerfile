# Gunakan base image golang:latest
FROM golang:latest

# Set timezone to Asia/Jakarta
ENV TZ=Asia/Jakarta

# Set WORKDIR ke direktori proyek
WORKDIR /go/src/bank-ina-test

# Install MySQL client
RUN apt-get update && \
    apt-get install -y default-mysql-client

# Tambahkan kode proyek Anda ke dalam container
COPY . .

# Install dependencies Go (jika menggunakan Go module)
RUN go mod tidy

# Install driver database MySQL untuk Go
RUN go get -u github.com/go-sql-driver/mysql

# Expose port yang digunakan oleh aplikasi Go (sesuaikan dengan kebutuhan)
EXPOSE 8080

# Jalankan aplikasi saat container berjalan
CMD ["go", "run", "main.go"]
