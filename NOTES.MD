# Notes du TP – Difficultés rencontrées et solutions apportées

Ces notes suivent l'ordre du TP : la découverte de l'API dans Postman, puis le développement du serveur Go.

## Partie 2 – Développement du serveur Go

### 1. Demande du token
**Problèmes successifs :**
- `string.NewReader` n'existe pas : le bon package est `strings`.
- Paramètres concaténés à la main, sans encodage : remplacés par `url.Values` et `Encode()`.
- En-tête `Content-Type: application/x-www-form-urlencoded` oublié : Spotify ne peut pas lire le body.
- Code de statut non vérifié : avec un faux secret, la fonction renvoyait un token vide **sans erreur**, car le JSON d'erreur se décode sans problème.
- `io.ReadAll` appelé avant le décodage JSON : le body d'une réponse ne se lit qu'une seule fois.
- Vérification du statut placée avant `defer resp.Body.Close()` : en cas d'erreur, le body n'était jamais fermé.

**Solution :** body construit avec `url.Values`, en-tête ajouté, puis dans cet ordre : `Do`, `defer Close`, `ReadAll`, vérification du statut (avec le body dans le message d'erreur), et enfin `json.Unmarshal`.

### 2. Renouvellement automatique du token
**Problème :** en écrivant `currentToken := …` dans la fonction, les variables du package n'étaient jamais modifiées.
**Cause :** `:=` déclare une **nouvelle** variable locale qui masque celle du package ; il faut `=` pour modifier une variable existante.
**Solution :** `GetToken()` renvoie le token en mémoire s'il expire dans plus d'une minute, et en demande un nouveau sinon. La marge d'une minute évite qu'un token expire pendant une série de requêtes.

### 3. Blocage de 23 heures par Spotify (429)
**Problème :** après plusieurs redémarrages du serveur, Spotify a répondu `429 Too Many Requests` avec un `Retry-After` d'environ 23 heures. Créer une nouvelle application n'a pas suffi, car le blocage portait sur le compte.
**Cause :** afficher la liste demande une requête par album, soit plus d'une centaine. Ces requêtes partaient en rafale, et à **chaque** redémarrage, puisque le cache n'était qu'en mémoire.
**Solutions :**
- **Espacement des requêtes** : au plus une requête toutes les 500 ms.
- **Cache sur disque** (`cache/jul_tracks.json`, valable 24 heures) : un redémarrage ne relance plus aucune requête.
- **`limit=50`** sur les titres d'album : une seule requête par album au lieu de deux pour les albums de plus de 20 titres.
- **Mémorisation du blocage** : pendant un long `Retry-After`, le serveur ne contacte plus Spotify et affiche directement la page d'erreur, au lieu d'insister.
- Pour une attente courte (10 secondes au plus), la requête est relancée automatiquement après le délai `Retry-After`.

### 4. Premier chargement lent
**Constat :** le premier affichage de la liste prend plusieurs minutes, à cause de l'espacement des requêtes.
**Explication :** le sujet impose d'afficher **tous** les titres et de filtrer **côté serveur**. Il faut donc récupérer toute la discographie, alors que l'API n'a plus d'endpoint groupé en Development Mode. Le cache rend les chargements suivants instantanés.

