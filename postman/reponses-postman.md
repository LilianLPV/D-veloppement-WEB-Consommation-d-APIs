# Réponses obtenues dans Postman

Résultats des requêtes Q0 à Q4 de la collection `TP-Spotify-JUL.postman_collection.json`.
Le token d'accès a été remplacé par `<access_token>` : aucun secret n'apparaît dans ce fichier.

---

## Q0 – Obtenir un token

```
POST https://accounts.spotify.com/api/token
```

Body en `x-www-form-urlencoded` : `grant_type=client_credentials`, `client_id`, `client_secret`.

**Réponse : 200 OK**

```json
{
  "access_token": "<access_token>",
  "token_type": "Bearer",
  "expires_in": 3600
}
```

Le token est valable 1 heure (`expires_in` = 3600 secondes). Il s'envoie ensuite dans l'en-tête `Authorization: Bearer <access_token>`.

---

## Q1 – Rechercher un artiste

```
GET https://api.spotify.com/v1/search?q=JUL&type=artist&limit=3
```

**Réponse : 200 OK**

**ID de JUL : `3IW7ScrzXmPvZhB27hmfgy`** (premier résultat). La recherche est approximative : elle renvoie aussi Julien Doré et Juliette Armanet.

```json
{
  "artists": {
    "href": "https://api.spotify.com/v1/search?offset=0&limit=3&query=JUL&type=artist",
    "limit": 3,
    "next": "https://api.spotify.com/v1/search?offset=3&limit=3&query=JUL&type=artist",
    "offset": 0,
    "previous": null,
    "total": 4,
    "items": [
      {
        "external_urls": {
          "spotify": "https://open.spotify.com/artist/3IW7ScrzXmPvZhB27hmfgy"
        },
        "href": "https://api.spotify.com/v1/artists/3IW7ScrzXmPvZhB27hmfgy",
        "id": "3IW7ScrzXmPvZhB27hmfgy",
        "images": [
          {
            "url": "https://i.scdn.co/image/ab6761610000e5ebaa70253f73793a0ff061069e",
            "height": 640,
            "width": 640
          },
          {
            "url": "https://i.scdn.co/image/ab67616100005174aa70253f73793a0ff061069e",
            "height": 320,
            "width": 320
          },
          {
            "url": "https://i.scdn.co/image/ab6761610000f178aa70253f73793a0ff061069e",
            "height": 160,
            "width": 160
          }
        ],
        "name": "Jul",
        "type": "artist",
        "uri": "spotify:artist:3IW7ScrzXmPvZhB27hmfgy"
      },
      {
        "external_urls": {
          "spotify": "https://open.spotify.com/artist/6jStlKLflTMUN6BajxrNlj"
        },
        "href": "https://api.spotify.com/v1/artists/6jStlKLflTMUN6BajxrNlj",
        "id": "6jStlKLflTMUN6BajxrNlj",
        "images": [
          {
            "url": "https://i.scdn.co/image/ab6761610000e5eba7475b5895e9a1d3800d60e4",
            "height": 640,
            "width": 640
          },
          {
            "url": "https://i.scdn.co/image/ab67616100005174a7475b5895e9a1d3800d60e4",
            "height": 320,
            "width": 320
          },
          {
            "url": "https://i.scdn.co/image/ab6761610000f178a7475b5895e9a1d3800d60e4",
            "height": 160,
            "width": 160
          }
        ],
        "name": "Julien Doré",
        "type": "artist",
        "uri": "spotify:artist:6jStlKLflTMUN6BajxrNlj"
      },
      {
        "external_urls": {
          "spotify": "https://open.spotify.com/artist/61CPKXT0bcKj8MKTNTMOXa"
        },
        "href": "https://api.spotify.com/v1/artists/61CPKXT0bcKj8MKTNTMOXa",
        "id": "61CPKXT0bcKj8MKTNTMOXa",
        "images": [
          {
            "url": "https://i.scdn.co/image/ab6761610000e5eb3869679fca4ed2840ceecb41",
            "height": 640,
            "width": 640
          },
          {
            "url": "https://i.scdn.co/image/ab676161000051743869679fca4ed2840ceecb41",
            "height": 320,
            "width": 320
          },
          {
            "url": "https://i.scdn.co/image/ab6761610000f1783869679fca4ed2840ceecb41",
            "height": 160,
            "width": 160
          }
        ],
        "name": "Juliette Armanet",
        "type": "artist",
        "uri": "spotify:artist:61CPKXT0bcKj8MKTNTMOXa"
      }
    ]
  }
}
```

---

## Q2 – Rechercher les titres de JUL

```
GET https://api.spotify.com/v1/search?q=artist:JUL&type=track&limit=10
```

**Réponse : 200 OK**

Les résultats sont paginés : pour voir les titres suivants, il faut « changer de page » en ajoutant `&offset=10`, ou suivre l'URL du champ `next`. Depuis février 2026, `limit` est plafonné à 10.

> Remarque : la réponse ci-dessous a été obtenue avec la limite par défaut (`limit=5` dans `href`), d'où `"limit": 5`, `"total": 3` et `"next": null`.

```json
{
  "tracks": {
    "href": "https://api.spotify.com/v1/search?offset=0&limit=5&query=artist%3AJUL&type=track",
    "limit": 5,
    "next": null,
    "offset": 0,
    "previous": null,
    "total": 3,
    "items": [
      {
        "album": {
          "album_type": "single",
          "artists": [
            {
              "external_urls": {
                "spotify": "https://open.spotify.com/artist/3IW7ScrzXmPvZhB27hmfgy"
              },
              "href": "https://api.spotify.com/v1/artists/3IW7ScrzXmPvZhB27hmfgy",
              "id": "3IW7ScrzXmPvZhB27hmfgy",
              "name": "Jul",
              "type": "artist",
              "uri": "spotify:artist:3IW7ScrzXmPvZhB27hmfgy"
            }
          ],
          "external_urls": {
            "spotify": "https://open.spotify.com/album/1BG4MKsWTJvsYB0cyVF6I7"
          },
          "href": "https://api.spotify.com/v1/albums/1BG4MKsWTJvsYB0cyVF6I7",
          "id": "1BG4MKsWTJvsYB0cyVF6I7",
          "images": [
            {
              "height": 640,
              "width": 640,
              "url": "https://i.scdn.co/image/ab67616d0000b27328bc7d7e8b7290dd785ca2cc"
            },
            {
              "height": 300,
              "width": 300,
              "url": "https://i.scdn.co/image/ab67616d00001e0228bc7d7e8b7290dd785ca2cc"
            },
            {
              "height": 64,
              "width": 64,
              "url": "https://i.scdn.co/image/ab67616d0000485128bc7d7e8b7290dd785ca2cc"
            }
          ],
          "is_playable": true,
          "name": "Ragnar 2",
          "release_date": "2026-10-23",
          "release_date_precision": "day",
          "total_tracks": 1,
          "type": "album",
          "uri": "spotify:album:1BG4MKsWTJvsYB0cyVF6I7"
        },
        "artists": [
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/3IW7ScrzXmPvZhB27hmfgy"
            },
            "href": "https://api.spotify.com/v1/artists/3IW7ScrzXmPvZhB27hmfgy",
            "id": "3IW7ScrzXmPvZhB27hmfgy",
            "name": "Jul",
            "type": "artist",
            "uri": "spotify:artist:3IW7ScrzXmPvZhB27hmfgy"
          }
        ],
        "disc_number": 1,
        "duration_ms": 414666,
        "explicit": true,
        "external_ids": {
          "isrc": "FRX282695997"
        },
        "external_urls": {
          "spotify": "https://open.spotify.com/track/4hEum3F9ASQVMqRm7VgcbW"
        },
        "href": "https://api.spotify.com/v1/tracks/4hEum3F9ASQVMqRm7VgcbW",
        "id": "4hEum3F9ASQVMqRm7VgcbW",
        "is_local": false,
        "is_playable": true,
        "name": "Ragnar 2",
        "track_number": 1,
        "type": "track",
        "uri": "spotify:track:4hEum3F9ASQVMqRm7VgcbW"
      },
      {
        "album": {
          "album_type": "single",
          "artists": [
            {
              "external_urls": {
                "spotify": "https://open.spotify.com/artist/3IW7ScrzXmPvZhB27hmfgy"
              },
              "href": "https://api.spotify.com/v1/artists/3IW7ScrzXmPvZhB27hmfgy",
              "id": "3IW7ScrzXmPvZhB27hmfgy",
              "name": "Jul",
              "type": "artist",
              "uri": "spotify:artist:3IW7ScrzXmPvZhB27hmfgy"
            }
          ],
          "external_urls": {
            "spotify": "https://open.spotify.com/album/3sjKkgJkvnRwYPPnFcaoED"
          },
          "href": "https://api.spotify.com/v1/albums/3sjKkgJkvnRwYPPnFcaoED",
          "id": "3sjKkgJkvnRwYPPnFcaoED",
          "images": [
            {
              "height": 640,
              "width": 640,
              "url": "https://i.scdn.co/image/ab67616d0000b27343c8b89de39bc66cf21c4674"
            },
            {
              "height": 300,
              "width": 300,
              "url": "https://i.scdn.co/image/ab67616d00001e0243c8b89de39bc66cf21c4674"
            },
            {
              "height": 64,
              "width": 64,
              "url": "https://i.scdn.co/image/ab67616d0000485143c8b89de39bc66cf21c4674"
            }
          ],
          "is_playable": true,
          "name": "Parfum quartier",
          "release_date": "2017-09-24",
          "release_date_precision": "day",
          "total_tracks": 1,
          "type": "album",
          "uri": "spotify:album:3sjKkgJkvnRwYPPnFcaoED"
        },
        "artists": [
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/3IW7ScrzXmPvZhB27hmfgy"
            },
            "href": "https://api.spotify.com/v1/artists/3IW7ScrzXmPvZhB27hmfgy",
            "id": "3IW7ScrzXmPvZhB27hmfgy",
            "name": "Jul",
            "type": "artist",
            "uri": "spotify:artist:3IW7ScrzXmPvZhB27hmfgy"
          }
        ],
        "disc_number": 1,
        "duration_ms": 226210,
        "explicit": true,
        "external_ids": {
          "isrc": "FR0U91701011"
        },
        "external_urls": {
          "spotify": "https://open.spotify.com/track/1fMOAIge1jLxq4ONi9bIFL"
        },
        "href": "https://api.spotify.com/v1/tracks/1fMOAIge1jLxq4ONi9bIFL",
        "id": "1fMOAIge1jLxq4ONi9bIFL",
        "is_local": false,
        "is_playable": true,
        "name": "Parfum quartier",
        "track_number": 1,
        "type": "track",
        "uri": "spotify:track:1fMOAIge1jLxq4ONi9bIFL"
      },
      {
        "album": {
          "album_type": "album",
          "artists": [
            {
              "external_urls": {
                "spotify": "https://open.spotify.com/artist/3IW7ScrzXmPvZhB27hmfgy"
              },
              "href": "https://api.spotify.com/v1/artists/3IW7ScrzXmPvZhB27hmfgy",
              "id": "3IW7ScrzXmPvZhB27hmfgy",
              "name": "Jul",
              "type": "artist",
              "uri": "spotify:artist:3IW7ScrzXmPvZhB27hmfgy"
            }
          ],
          "external_urls": {
            "spotify": "https://open.spotify.com/album/3yMDPzPjM2VqLHENlVgiEF"
          },
          "href": "https://api.spotify.com/v1/albums/3yMDPzPjM2VqLHENlVgiEF",
          "id": "3yMDPzPjM2VqLHENlVgiEF",
          "images": [
            {
              "height": 640,
              "width": 640,
              "url": "https://i.scdn.co/image/ab67616d0000b27385d9fb2e7833cb239769c02d"
            },
            {
              "height": 300,
              "width": 300,
              "url": "https://i.scdn.co/image/ab67616d00001e0285d9fb2e7833cb239769c02d"
            },
            {
              "height": 64,
              "width": 64,
              "url": "https://i.scdn.co/image/ab67616d0000485185d9fb2e7833cb239769c02d"
            }
          ],
          "is_playable": true,
          "name": "Dans ma paranoïa",
          "release_date": "2014-02-24",
          "release_date_precision": "day",
          "total_tracks": 17,
          "type": "album",
          "uri": "spotify:album:3yMDPzPjM2VqLHENlVgiEF"
        },
        "artists": [
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/3IW7ScrzXmPvZhB27hmfgy"
            },
            "href": "https://api.spotify.com/v1/artists/3IW7ScrzXmPvZhB27hmfgy",
            "id": "3IW7ScrzXmPvZhB27hmfgy",
            "name": "Jul",
            "type": "artist",
            "uri": "spotify:artist:3IW7ScrzXmPvZhB27hmfgy"
          }
        ],
        "disc_number": 1,
        "duration_ms": 316777,
        "explicit": false,
        "external_ids": {
          "isrc": "FR2MX1300150"
        },
        "external_urls": {
          "spotify": "https://open.spotify.com/track/0SKNkEMV313mE6oYsgVCgm"
        },
        "href": "https://api.spotify.com/v1/tracks/0SKNkEMV313mE6oYsgVCgm",
        "id": "0SKNkEMV313mE6oYsgVCgm",
        "is_local": false,
        "is_playable": true,
        "name": "J'oublie tout",
        "track_number": 5,
        "type": "track",
        "uri": "spotify:track:0SKNkEMV313mE6oYsgVCgm"
      }
    ]
  }
}
```

---

## Q3 – Récupérer un titre par son ID

```
GET https://api.spotify.com/v1/tracks/0goer5e8B6SJaTV2B51k6f
```

**Réponse : 200 OK** (« La bandite », album *Rien 100 Rien*)

```json
{
  "album": {
    "album_type": "album",
    "artists": [
      {
        "external_urls": {
          "spotify": "https://open.spotify.com/artist/3IW7ScrzXmPvZhB27hmfgy"
        },
        "href": "https://api.spotify.com/v1/artists/3IW7ScrzXmPvZhB27hmfgy",
        "id": "3IW7ScrzXmPvZhB27hmfgy",
        "name": "Jul",
        "type": "artist",
        "uri": "spotify:artist:3IW7ScrzXmPvZhB27hmfgy"
      }
    ],
    "external_urls": {
      "spotify": "https://open.spotify.com/album/2GLuHlc49dJKY8yzxUZb8p"
    },
    "href": "https://api.spotify.com/v1/albums/2GLuHlc49dJKY8yzxUZb8p",
    "id": "2GLuHlc49dJKY8yzxUZb8p",
    "images": [
      {
        "url": "https://i.scdn.co/image/ab67616d0000b2734a7624705f2fd897adb437a2",
        "width": 640,
        "height": 640
      },
      {
        "url": "https://i.scdn.co/image/ab67616d00001e024a7624705f2fd897adb437a2",
        "width": 300,
        "height": 300
      },
      {
        "url": "https://i.scdn.co/image/ab67616d000048514a7624705f2fd897adb437a2",
        "width": 64,
        "height": 64
      }
    ],
    "is_playable": true,
    "name": "Rien 100 Rien",
    "release_date": "2019-06-28",
    "release_date_precision": "day",
    "total_tracks": 25,
    "type": "album",
    "uri": "spotify:album:2GLuHlc49dJKY8yzxUZb8p"
  },
  "artists": [
    {
      "external_urls": {
        "spotify": "https://open.spotify.com/artist/3IW7ScrzXmPvZhB27hmfgy"
      },
      "href": "https://api.spotify.com/v1/artists/3IW7ScrzXmPvZhB27hmfgy",
      "id": "3IW7ScrzXmPvZhB27hmfgy",
      "name": "Jul",
      "type": "artist",
      "uri": "spotify:artist:3IW7ScrzXmPvZhB27hmfgy"
    }
  ],
  "disc_number": 1,
  "duration_ms": 160893,
  "explicit": true,
  "external_ids": {
    "isrc": "FR0U91900069"
  },
  "external_urls": {
    "spotify": "https://open.spotify.com/track/0goer5e8B6SJaTV2B51k6f"
  },
  "href": "https://api.spotify.com/v1/tracks/0goer5e8B6SJaTV2B51k6f",
  "id": "0goer5e8B6SJaTV2B51k6f",
  "is_local": false,
  "is_playable": true,
  "name": "La bandite",
  "track_number": 1,
  "type": "track",
  "uri": "spotify:track:0goer5e8B6SJaTV2B51k6f"
}
```

---

## Q4 – Réponses d'erreur

Les trois erreurs ont été reproduites sur `GET /v1/tracks/{id}`.

### Requête de référence (avec un token valide)

```
GET https://api.spotify.com/v1/tracks/1fMOAIge1jLxq4ONi9bIFL
```

**Réponse : 200 OK** (« Parfum quartier »). C'est la même requête qui est ensuite modifiée pour provoquer chaque erreur.

```json
{
  "album": {
    "album_type": "single",
    "artists": [
      {
        "external_urls": {
          "spotify": "https://open.spotify.com/artist/3IW7ScrzXmPvZhB27hmfgy"
        },
        "href": "https://api.spotify.com/v1/artists/3IW7ScrzXmPvZhB27hmfgy",
        "id": "3IW7ScrzXmPvZhB27hmfgy",
        "name": "Jul",
        "type": "artist",
        "uri": "spotify:artist:3IW7ScrzXmPvZhB27hmfgy"
      }
    ],
    "external_urls": {
      "spotify": "https://open.spotify.com/album/3sjKkgJkvnRwYPPnFcaoED"
    },
    "href": "https://api.spotify.com/v1/albums/3sjKkgJkvnRwYPPnFcaoED",
    "id": "3sjKkgJkvnRwYPPnFcaoED",
    "images": [
      {
        "url": "https://i.scdn.co/image/ab67616d0000b27343c8b89de39bc66cf21c4674",
        "width": 640,
        "height": 640
      },
      {
        "url": "https://i.scdn.co/image/ab67616d00001e0243c8b89de39bc66cf21c4674",
        "width": 300,
        "height": 300
      },
      {
        "url": "https://i.scdn.co/image/ab67616d0000485143c8b89de39bc66cf21c4674",
        "width": 64,
        "height": 64
      }
    ],
    "is_playable": true,
    "name": "Parfum quartier",
    "release_date": "2017-09-24",
    "release_date_precision": "day",
    "total_tracks": 1,
    "type": "album",
    "uri": "spotify:album:3sjKkgJkvnRwYPPnFcaoED"
  },
  "artists": [
    {
      "external_urls": {
        "spotify": "https://open.spotify.com/artist/3IW7ScrzXmPvZhB27hmfgy"
      },
      "href": "https://api.spotify.com/v1/artists/3IW7ScrzXmPvZhB27hmfgy",
      "id": "3IW7ScrzXmPvZhB27hmfgy",
      "name": "Jul",
      "type": "artist",
      "uri": "spotify:artist:3IW7ScrzXmPvZhB27hmfgy"
    }
  ],
  "disc_number": 1,
  "duration_ms": 226210,
  "explicit": true,
  "external_ids": {
    "isrc": "FR0U91701011"
  },
  "external_urls": {
    "spotify": "https://open.spotify.com/track/1fMOAIge1jLxq4ONi9bIFL"
  },
  "href": "https://api.spotify.com/v1/tracks/1fMOAIge1jLxq4ONi9bIFL",
  "id": "1fMOAIge1jLxq4ONi9bIFL",
  "is_local": false,
  "is_playable": true,
  "name": "Parfum quartier",
  "track_number": 1,
  "type": "track",
  "uri": "spotify:track:1fMOAIge1jLxq4ONi9bIFL"
}
```

### Q4a – Sans token

Même requête, Authorization réglé sur **No Auth**.

**Réponse : 401 Unauthorized**

```json
{
  "error": {
    "status": 401,
    "message": "Missing/invalid/expired access token"
  }
}
```

### Q4b – Token invalide

Même requête, avec un token modifié.

**Réponse : 401 Unauthorized**

```json
{
  "error": {
    "status": 401,
    "message": "Missing/invalid/expired access token"
  }
}
```

### Q4c – Identifiant de titre inexistant

```
GET https://api.spotify.com/v1/tracks/0000000000000000000000
```

Identifiant au bon format (22 caractères) qui ne correspond à aucun titre.

**Réponse : 404 Not Found**

```json
{
  "error": {
    "status": 404,
    "message": "Resource not found"
  }
}
```

### Récapitulatif

| Cas | Code | Message |
|---|---|---|
| Sans token | 401 Unauthorized | Missing/invalid/expired access token |
| Token invalide | 401 Unauthorized | Missing/invalid/expired access token |
| ID inexistant | 404 Not Found | Resource not found |

Spotify renvoie le même message pour un token absent, invalide ou expiré : seul le code 401 indique un problème d'authentification, sans en préciser la cause.