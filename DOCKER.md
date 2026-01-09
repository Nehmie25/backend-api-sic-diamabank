# Documentation Docker - SIC DiamaBankAPI

## Prérequis

- Docker (v20+)
- Docker Compose (v1.29+)
- Accès à la base de données Oracle

## Configuration

### 1. Créer le fichier `.env`

Copiez le fichier `.env.example` en `.env` et configurez les variables d'environnement :

```bash
cp .env.example .env
```

Éditez `.env` avec vos identifiants Oracle :

```
DB_USER=SICPROD
DB_PASSWORD=votre_mot_de_passe
DB_HOST=10.0.16.3
DB_PORT=1521
DB_NAME=ORCLPDB
API_PORT=8080
```

## Utilisation

### Build l'image Docker

```bash
docker build -t sic-diamabank-api .
```

### Lancer avec Docker Compose

```bash
docker-compose up -d
```

L'API sera accessible sur `http://localhost:8080`

### Lancer avec Docker uniquement

```bash
docker run -d \
  --name sic-diamabank-api \
  -p 8080:8080 \
  -e DB_USER=SICPROD \
  -e DB_PASSWORD=your_password \
  -e DB_HOST=10.0.16.3 \
  -e DB_PORT=1521 \
  -e DB_NAME=ORCLPDB \
  sic-diamabank-api
```

## Commandes utiles

### Voir les logs

```bash
# Avec Docker Compose
docker-compose logs -f api

# Avec Docker seul
docker logs -f sic-diamabank-api
```

### Arrêter le conteneur

```bash
# Avec Docker Compose
docker-compose down

# Avec Docker seul
docker stop sic-diamabank-api
docker rm sic-diamabank-api
```

### Rebuild après modifications

```bash
docker-compose up -d --build
```

## Endpoints de l'API

- `POST /declaration/personnephysique?date=YYYY-MM-DD`
- `POST /declaration/personnemorale?date=YYYY-MM-DD`
- `POST /declaration/engagements?date=YYYY-MM-DD`
- `POST /declaration/encours?date=YYYY-MM-DD`
- `POST /declaration/comptedebiteurs?date=YYYY-MM-DD`

## Dépannage

### L'API ne se connecte pas à Oracle

- Vérifiez les identifiants dans `.env`
- Vérifiez que la base de données est accessible depuis le conteneur
- Consultez les logs : `docker-compose logs api`

### Port 8080 déjà en utilisation

Changez le port dans `.env` ou utilisez :

```bash
docker run -p 8081:8080 ... sic-diamabank-api
```

## Architecture

```
├── Dockerfile           # Multi-stage build (builder + alpine)
├── docker-compose.yml   # Configuration des services
├── .dockerignore        # Fichiers exclus du build
└── main.go             # Application Go configurée pour Docker
```

**Dockerfile utilise un build multi-stage** pour réduire la taille de l'image finale (Alpine Linux).
