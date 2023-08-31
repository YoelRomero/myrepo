FROM debian:latest

# Установка зависимостей

RUN apt-get update && apt-get install -y \
    build-essential \
    wget            \
    ca-certificates \
&& rm -rf /var/lib/apt/lists/*

# Установка Go

ENV GOLANG_VERSION 1.21.0
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz

RUN wget -q $GOLANG_DOWNLOAD_URL -O golang.tar.gz \
&& tar -C /usr/local -xzf golang.tar.gz \
&& rm golang.tar.gz

ENV PATH $PATH:/usr/local/go/bin

# Установка зависимостей Go
COPY /source_dir /app
WORKDIR /app

# COPY redis_user.crt /usr/local/etc/redis/redis_user.crt
# COPY redis_user_private.key /usr/local/etc/redis/redis_user_private.key
# COPY redis_ca.pem /usr/local/etc/redis/redis_ca.pem

RUN go get /app
RUN go install /app

# Настройка порта и запуск приложения
# RUN go build -o main .
# RUN chmod +x main
EXPOSE 8000

#
CMD ["go", "run", "/app/main.go"]  
