# Hangman Classic

Ce projet est une implémentation du jeu classique du **Pendu** en Go. Il a été réalisé durant ma **première année d'études** dans le cadre d'un projet pédagogique. L'objectif de ce projet était d'apprendre les bases de la programmation en Go, telles que la gestion des fichiers, la structuration d'un programme, et l'organisation en packages.

## Description

Le jeu propose au joueur de découvrir un mot en devinant des lettres une à une, tout en représentant visuellement le bonhomme "pendu" dans le terminal. Le projet met en œuvre les concepts suivants :

- Manipulation des fichiers texte pour charger des mots et différentes polices d'ASCII art.
- Organisation structurée en packages pour une gestion claire des fonctionnalités.
- Interactions dynamiques avec l'utilisateur via un terminal.
- Fonctionnalité permettant de sauvegarder l'état de la partie.

## Fonctionnalités

- **Gestion de la difficulté** : trois niveaux différents (facile, moyen, difficile) avec des listes de mots spécifiques.
- **Personnalisation visuelle** : trois styles d'ASCII art disponibles (`Standard`, `Shadow`, `Thinkertoy`).
- **Affichage d'un pendu dynamique** : le pendu évolue en fonction des erreurs dans les tentatives.
- **Sauvegarde et restauration** : possibilité de sauvegarder une partie en cours pour la reprendre ultérieurement.
- **Validations des entrées utilisateur** : empêche la saisie de caractères déjà utilisés et vérifie la validité des lettres saisies.

## Comment jouer ?

1. Clonez le projet :
   ```bash
   git clone https://github.com/Min3r0/hangman-classic.git
   cd hangman-classic
   ```

2. Installez les dépendances avec Go Modules :
   ```bash
   go mod tidy
   ```

3. Exécutez le jeu via la commande suivante :
   ```bash
   go run main.go
   ```

4. Suivez les instructions affichées dans le terminal pour jouer. Bonne chance !



## Organisation du projet

Le projet est organisé en plusieurs packages pour maximiser la clarté et la modularité du code :

- **`Game`** : Gère la logique du jeu et les interactions utilisateur.
- **`Print`** : Contient tout l'affichage, notamment le mot masqué, le pendu ASCII et les lettres utilisées.
- **`Verify`** : Contient les fonctions de vérification (lettres valides, conditions de victoire, etc.).
- **`RequestUsr`** : Permet de gérer les entrées utilisateur (choix du niveau, lettres, etc.).
- **`CreateList`** : Manipule les listes utilisées dans le jeu (génération des mots à deviner, etc.).
- **`StartAndStop`** : Implémente les fonctionnalités de sauvegarde et d'arrêt du jeu.
- **`Structure`** : Déclare les structures de données spécifiques au jeu.

## Notes pédagogiques

Ce projet a été réalisé dans le cadre d’un projet de cours pour apprendre :

- L'organisation d’un projet en packages.
- La manipulation des entrées/sorties en Go.
- L'ASCII art pour des affichages dynamiques et interactifs.
- Les principes de base de la gestion d'état d'un programme.

## Améliorations possibles

- Intégrer une interface graphique pour une expérience plus immersive.
- Ajouter un mode multijoueur avec suivi des scores.
- Permettre la gestion de plusieurs langues pour les mots.
- Implémenter un système d’aide permettant de révéler une lettre.


---
Créé par Romain AUGÉ
