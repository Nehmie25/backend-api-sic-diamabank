# Backend API SIC DiamaBanK

## Description du Projet

L'API SIC DiamaBanK est une application backend développée en **Go** qui expose des endpoints REST pour récupérer et gérer les données déclaratives du système d'information de crédit (SIC) de DiamaBanK.

Cette API permet d'accéder aux données de:
- **Personnes Physiques** - Informations sur les individus emprunteurs
- **Personnes Morales** - Informations sur les entreprises emprunteurs
- **Engagements** - Données des engagements de crédit
- **Encours** - Informations sur les encours de crédit
- **Comptes Débiteurs** - Données des comptes débiteurs

L'application se connecte à une **base de données Oracle** et fournit une interface REST simple et efficace pour consulter ces données déclaratives.

## Architecture

```
internal/
├── controllers/     # Gestion des requêtes HTTP et réponses
├── services/        # Logique métier
├── repositories/    # Accès et manipulation des données
├── database/        # Configuration et connexion à la base de données
├── models/          # Définition des structures de données
└── router/          # Configuration des routes HTTP
```

## Routes de l'API

### Accueil
- **GET** `/` - Message de bienvenue et liste des endpoints disponibles

### Déclarations
- **GET** `/declaration/personnephysique?date=DD/MM/YY` - Récupérer les données des personnes physiques par date
- **GET** `/declaration/personnemorale?date=DD/MM/YY` - Récupérer les données des personnes morales par date
- **GET** `/declaration/engagements?date=DD/MM/YY` - Récupérer les données des engagements par date
- **GET** `/declaration/encours?date=DD/MM/YY` - Récupérer les données des encours par date
- **GET** `/declaration/comptedebiteurs?date=DD/MM/YY` - Récupérer les données des comptes débiteurs par date

### Paramètre date
Le paramètre `date` est optionnel et doit être au format **DD/MM/YY**
- Exemple: `?date=15/01/26`

## Configuration

### Variables d'Environnement

Créez un fichier `configs/.env` à la racine du projet avec les variables suivantes:

```env
# Configuration de la base de données Oracle
DB_USER=user
DB_PASSWORD=password
DB_HOST=ip
DB_PORT=port
DB_NAME=nom_base_de_donnée


# Port de l'API (optionnel, par défaut 8080)
API_PORT=8080
```

**Note**: Les valeurs par défaut sont déjà définies dans le code si le fichier `.env` n'existe pas.

## Installation et Lancement

### Méthode 1: Avec Docker Compose (Recommandé)

#### Prérequis
- Docker et Docker Compose installés

#### Étapes
1. Clonez le projet et naviguez dans le répertoire:
   ```bash
   git clone git@github.com:Nehmie25/backend-api-sic-diamabank.git
   ```

2. Créez le fichier de configuration .env:

3. Lancez l'application avec Docker Compose:
   ```bash
   docker compose up
   ```

4. L'API sera accessible sur `http://localhost:8080`

#### Arrêter l'application
```bash
docker compose down
```

### Méthode 2: Lancer localement avec Go

#### Prérequis
- Go 1.20 ou supérieur
- Oracle Instant Client (pour la connexion à Oracle)
- Accès réseau à la base de données Oracle

#### Étapes
1. Naviguez dans le répertoire du projet:
   ```bash
   cd /backend-api-sic-diamabank
   ```

2. Téléchargez les dépendances:
   ```bash
   go mod download
   ```

3. Lancez l'application:
   ```bash
   go run main.go
   ```

4. L'API sera accessible sur `http://localhost:8080`

## Utilisation

### Exemple de requête - Personnes Physiques

```bash
curl "http://localhost:8080/declaration/personnephysique?date=15/01/26"
```

### Exemple de requête - Engagements

```bash
curl "http://localhost:8080/declaration/engagements?date=15/01/26"
```

### Réponse de bienvenue

```bash
curl http://localhost:8080/
```

Réponse attendue:
```json
{
  "message": "Bienvenue sur l'API SIC DiamBank",
  "status": "En ligne",
  "version": "1.0.0",
  "endpoints": [
    "GET /declaration/personnephysique?date=DD/MM/YY",
    "GET /declaration/personnemorale?date=DD/MM/YY",
    "GET /declaration/engagements?date=DD/MM/YY",
    "GET /declaration/encours?date=DD/MM/YY",
    "GET /declaration/comptedebiteurs?date=DD/MM/YY"
  ]
}
```