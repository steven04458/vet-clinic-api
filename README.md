# vet-clinic-api

Projet Final : API pour une Clinique Vétérinaire

## Description

Cette API permet de gérer les informations d'une clinique vétérinaire, y compris les chats, les visites et les traitements. Elle fournit des endpoints pour créer, lire, mettre à jour et supprimer des enregistrements pour chaque entité.

## Prérequis

- Go 1.16 ou supérieur

## Installation

1. Clonez le dépôt :
    ```sh
    git clone https://github.com/votre-utilisateur/vet-clinic-api.git
    cd vet-clinic-api
    ```

2. Installez les dépendances :
    ```sh
    go mod download
    ```

3. Configurez les variables d'environnement nécessaires (voir `config/config.go` pour les détails).

## Démarrage

Démarrez le serveur :
    ```sh
    go run main.go
    ```

Le serveur sera démarré sur le port `8080` par défaut.

## Endpoints

### Chats

- `POST /api/v1/cats` : Créer un nouveau chat
- `GET /api/v1/cats` : Récupérer tous les chats
- `GET /api/v1/cats/{id}` : Récupérer un chat par ID
- `GET /api/v1/cats/{id}/history` : Récupérer un chat par ID avec les visits et traitements
- `PUT /api/v1/cats/{id}` : Mettre à jour un chat par ID
- `DELETE /api/v1/cats/{id}` : Supprimer un chat par ID

### Visites

- `POST /api/v1/visits` : Créer une nouvelle visite
- `GET /api/v1/visits` : Récupérer toutes les visites
- `GET /api/v1/visits/{id}` : Récupérer une visite par ID
- `GET /api/v1/visits/filter` : Récupérer une visite par filtrage
- `PUT /api/v1/visits/{id}` : Mettre à jour une visite par ID
- `DELETE /api/v1/visits/{id}` : Supprimer une visite par ID

### Traitements

- `POST /api/v1/treatments` : Créer un nouveau traitement
- `GET /api/v1/treatments` : Récupérer tous les traitements
- `GET /api/v1/treatments/{id}` : Récupérer un traitement par ID
- `PUT /api/v1/treatments/{id}` : Mettre à jour un traitement par ID
- `DELETE /api/v1/treatments/{id}` : Supprimer un traitement par ID

## Configuration

La configuration de l'application se fait via le fichier `config/config.go`. Vous pouvez y définir les paramètres de connexion à la base de données et d'autres paramètres de configuration.
