# Les titres de JUL – Exploitation de l'API Spotify

Site web écrit en Go qui affiche l'ensemble des titres de l'artiste JUL à partir de l'API Web de Spotify, avec des filtres et une page de détail pour chaque titre.

Le serveur récupère lui-même les données auprès de Spotify, puis les affiche avec des templates `html/template` (fichiers `.gohtml`). Le site n'utilise aucun JavaScript.

## Fonctionnalités

- **Liste des titres de JUL** : nom, pochette, date de sortie, artistes et durée (mm:ss), avec un bouton vers le détail.
- **Filtres combinables** (formulaire en GET, filtrage écrit à la main en Go) :
  - **Année** de sortie, déduite des données ;
  - **Type** : single ou album ;
  - **Feat** : oui (au moins un autre artiste que JUL) ou non (JUL seul).
- **Détail d'un titre** : nom, pochette, album, artistes, date de sortie, durée et lien vers Spotify.
- **Gestion des erreurs** : page 404 personnalisée (route inconnue ou identifiant inexistant) et page d'erreur propre si Spotify ne répond pas (5xx, timeout, réponse illisible).
- **Gestion des appels à Spotify** : timeout sur le client HTTP, renouvellement automatique du token, respect de l'en-tête `Retry-After` en cas de 429, espacement des requêtes et cache des titres.

## Prérequis

- [Go](https://go.dev/dl/) 1.22 ou plus récent (les routes utilisent les motifs `GET /tracks/{id}` et `/{$}`)
- Un compte Spotify **Premium** : depuis février 2026, c'est obligatoire pour qu'une application en Development Mode fonctionne
- Une application créée sur le [Dashboard Spotify for Developers](https://developer.spotify.com/dashboard), avec l'API **Web API** cochée

## Configuration des variables d'environnement

1. Créer un fichier `.env` à la racine du projet, en copiant `.env.example`.
2. Renseigner les identifiants de l'application Spotify (Dashboard → l'application → **Settings**) :
   ```
   SPOTIFY_CLIENT_ID=votre_client_id
   SPOTIFY_CLIENT_SECRET=votre_client_secret
   ```

| Variable | Rôle |
|---|---|
| `SPOTIFY_CLIENT_ID` | Identifiant de l'application Spotify |
| `SPOTIFY_CLIENT_SECRET` | Secret de l'application Spotify |

Le fichier `.env` est ignoré par Git : il ne doit jamais être commité.

## Lancement

Depuis la racine du projet (les dossiers `templates/`, `static/` et le fichier `.env` sont cherchés à partir de là) :

```bash
go run .
```

Le site est ensuite accessible sur **http://localhost:8080**.

**Premier chargement :** pour afficher la liste, le serveur doit récupérer tous les albums de JUL, puis les titres de chaque album (une requête par album). Les requêtes sont espacées pour respecter la limite de Spotify : le premier chargement peut prendre quelques minutes. Il ne faut pas recharger la page pendant ce temps. Le résultat est ensuite enregistré dans `cache/jul_tracks.json` et réutilisé pendant 24 heures, y compris après un redémarrage. Pour forcer une nouvelle récupération, il suffit de supprimer le dossier `cache/`.

## Routes

| Méthode | Route | Description |
|---|---|---|
| GET | `/` | Liste de tous les titres de JUL |
| GET | `/?year=2025&type=album&feat=oui` | Liste filtrée (paramètres facultatifs et combinables) |
| GET | `/tracks/{id}` | Détail du titre dont l'identifiant Spotify est `{id}` |
| GET | `/static/...` | Fichiers statiques (CSS) |
| — | toute autre route | Page 404 personnalisée |

Paramètres des filtres :

| Paramètre | Valeurs | Effet |
|---|---|---|
| `year` | `2026`, `2025`… | Titres sortis cette année-là |
| `type` | `single`, `album` | Titres issus d'un single ou d'un album |
| `feat` | `oui`, `non` | Titres avec ou sans autre artiste que JUL |

## Structure du projet

```
├── main.go                  Point d'entrée : chargement du .env, routes, serveur sur localhost:8080
├── spotify/
│   ├── token.go             Récupération du token et renouvellement automatique
│   ├── client.go            Requêtes GET vers l'API : timeout, 429, 404, 5xx, espacement
│   ├── jul.go               Récupération des titres de JUL, cache mémoire et disque, détail d'un titre
│   └── models.go            Structures Go des réponses JSON (Track, Album, Artist, Image)
├── handlers/
│   ├── home.go              Page liste
│   ├── filters.go           Filtrage écrit à la main (année, type, feat)
│   ├── tracks.go            Page détail
│   └── errors.go            Pages 404 et erreur Spotify
├── templates/               home.gohtml, tracks.gohtml, error.gohtml
├── static/css/              home.css, tracks.css
├── postman/                 Collection Postman des requêtes Q0 à Q4
├── NOTES.md                 Difficultés rencontrées et solutions apportées
└── .env.example             Variables d'environnement nécessaires (sans valeur)
```

## Collection Postman

Le fichier `postman/TP-Spotify-JUL.postman_collection.json` contient les requêtes Q0 à Q4. Après l'import dans Postman, renseigner `client_id` et `client_secret` dans les variables de la collection, puis exécuter Q0 : le token est enregistré automatiquement pour les requêtes suivantes.