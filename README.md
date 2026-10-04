# Puissance 4 Web

**Projet scolaire individuel**
Ce projet a été réalisé de manière individuelle dans le cadre de mes études en informatique à Ynov Campus.

## Description
Il s'agit d'une implémentation web légère du jeu de stratégie classique **Puissance 4** (Connect 4), développée avec **Go** (Golang) pour le backend et **HTML/CSS** pour le frontend.

Le projet met à disposition une grille jouable où deux joueurs peuvent s'affronter localement. Le serveur gère toute la logique du jeu : validation des coups, détection de victoire (horizontale, verticale, diagonale) et égalité.

## Fonctionnalités
- **Gameplay Classique** : Grille 6x7 respectant les règles standards.
- **Mode Deux Joueurs** : Multijoueur local avec alternance de tours (🔴 Rouge vs 🟡 Jaune).
- **Logique de Jeu (Serveur)** :
  - Détection automatique de la victoire (4 jetons alignés).
  - Détection d'égalité (plateau rempli).
  - Prévention de coups invalides (colonnes pleines).
- **Interface Utilisateur Interactive** :
  - Effets de survol (hover) sur les boutons de sélection.
  - Indicateurs visuels du tour actuel.
  - **Animation de feux d'artifice** en cas de victoire.
  - Arrière-plan animé.
- **Fonction de réinitialisation** : Possibilité de recommencer la partie immédiatement à la fin d'un match.

## Stack Technique
- **Backend** : Go (Golang)
- **Frontend** : HTML5, CSS3
- **Templating** : Package Go `html/template`

## Arborescence du Projet
```text
projet_puissance4-web/
├── main.go            # Point d'entrée : serveur HTTP et logique de jeu
├── go.mod             # Définition du module Go
├── templates/         # Templates HTML (home.html, game.html)
├── style/             # Fichiers CSS (style.css)
└── static/            # Ressources statiques (images, gifs)
```

## Installation et Lancement
1. Assurez-vous d'avoir installé Go.
2. Clonez ce dépôt.
3. Exécutez le serveur depuis le répertoire racine :
   ```bash
   go run main.go
   ```
4. Ouvrez votre navigateur et accédez à `http://localhost:8080`.
