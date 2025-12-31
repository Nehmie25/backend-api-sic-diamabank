# Multi-stage build
FROM golang:1.24.10-bookworm AS builder

WORKDIR /app

# Installer les dépendances nécessaires pour godror
RUN apt-get update && apt-get install -y --no-install-recommends \
    pkg-config \
    build-essential \
    ca-certificates \
    libaio1 \
    wget \
    unzip \
    && rm -rf /var/lib/apt/lists/*

# Télécharger et installer Oracle Instant Client (version 21)
RUN mkdir -p /opt/oracle && cd /opt/oracle && \
    wget --no-verbose https://download.oracle.com/otn_software/linux/instantclient/219000/instantclient-basic-linux.x64-21.9.0.0.0dbru.zip && \
    unzip -q instantclient-basic-linux.x64-21.9.0.0.0dbru.zip && \
    rm instantclient-basic-linux.x64-21.9.0.0.0dbru.zip && \
    ls -la

# Copier les fichiers de dépendances
COPY go.mod go.sum* ./

# Télécharger les dépendances
RUN go mod download

# Copier le code source
COPY . .

# Compiler l'application avec Oracle Instant Client
RUN LD_LIBRARY_PATH=/opt/oracle/instantclient_21_9:$LD_LIBRARY_PATH \
    PKG_CONFIG_PATH=/opt/oracle/instantclient_21_9/lib/pkgconfig \
    CGO_ENABLED=1 \
    GOOS=linux \
    go build -a -installsuffix cgo -ldflags="-s -w" -o main .

# Stage final - image minimaliste
FROM debian:bookworm-slim

WORKDIR /root/

# Installer uniquement les dépendances runtime nécessaires
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    libc6 \
    libaio1 \
    libstdc++6 \
    && rm -rf /var/lib/apt/lists/*

# Créer le répertoire Oracle
RUN mkdir -p /opt/oracle

# Copier Oracle Instant Client depuis le builder
COPY --from=builder /opt/oracle/instantclient_21_9 /opt/oracle/instantclient_21_9

# Configurer les variables d'environnement pour Oracle
ENV LD_LIBRARY_PATH=/opt/oracle/instantclient_21_9:$LD_LIBRARY_PATH
ENV ORACLE_HOME=/opt/oracle/instantclient_21_9

# Copier le binaire depuis le builder
COPY --from=builder /app/main .

# Exposer le port
EXPOSE 8080

# Lancer l'application
CMD ["./main"]
