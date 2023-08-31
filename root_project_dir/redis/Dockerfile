FROM redis:6.2-bookworm

# Копирование конфигурационного файла Redis
COPY redis.conf /usr/local/etc/redis/redis.conf

# COPY redis_user.crt /usr/local/etc/redis/redis_user.crt
# COPY redis_user_private.key /usr/local/etc/redis/redis_user_private.key
# COPY redis_ca.pem /usr/local/etc/redis/redis_ca.pem

# Открытие порта Redis
EXPOSE 6379

# Запуск Redis с использованием кастомного конфигурационного файла
CMD ["redis-server", "/usr/local/etc/redis/redis.conf"]

