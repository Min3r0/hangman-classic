Voici notre projet Hangman.

Ce programme effectue le jeu du pendu sur le terminal, vous avez donc 10 essai pour essayer de trouver le bon mot.
Il est possible d'essayer des lettres, une mauvaise lettre vous fera perdre 1 essai.

Hangman s'affichera donc comme cela:
```console
cas 1:








         
         

         
         
=========
```
```console
cas 2:
         
      |  
      |  
      |  
      |  
      |  
=========
```
```console
cas 3:

  +---+  
      |  
      |  
      |  
      |  
      |  
=========
```
```console
cas 4:

  +---+  
  |   |  
      |  
      |  
      |  
      |  
=========
```
```console
cas 5:

  +---+  
  |   |  
  O   |  
      |  
      |  
      |  
=========
```
```console
cas 6:

  +---+  
  |   |  
  O   |  
  |   |  
      |  
      |  
=========
```
```console
cas 7:

  +---+  
  |   |  
  O   |  
 /|   |  
      |  
      |  
=========
```
```console
cas 8:

  +---+  
  |   |  
  O   |  
 /|\  |  
      |  
      |  
=========
```
```console
cas 9:

  +---+  
  |   |  
  O   |  
 /|\  |  
 /    |  
      |  
=========
```
```console
cas 10:

  +---+  
  |   |  
  O   |  
 /|\  |  
 / \  |  
      |  
=========
```

A cela nous avons ajouter des features qui sont les suivantes:
- Le joueur à l'opportunité de mettre des mots si il le souhaite, chaque mauvais mot lui fera perdre 2 essais.
- Si le joueurs trouve le mots le jeu s'arretera avec un message comme quoi il a gagné.
- Le joueurs n'aura pas l'opportunité de rentrer le meme mot ou lettre une seconde fois.

Nous avons ajouté aussi L'ASCII-ART, le joueur aura le choix au début de la partie entre trois different type d'ASCII-ART:
- Le Standard:
```console
__     __                        _                                           
\ \   / /                       | |                                    ___   
 \ \_/ /    ___    _   _        | |__     __ _  __   __   ___         ( _ )  
  \   /    / _ \  | | | |       |  _ \   / _` | \ \ / /  / _ \        / _ \  
   | |    | (_) | | |_| |       | | | | | (_| |  \ V /  |  __/       | (_) | 
   |_|     \___/   \__,_|       |_| |_|  \__,_|   \_/    \___|        \___/   
                                                                             
```

- Le Shadow:
```console

_|      _|                         _|                                            _|_|   
  _|  _|     _|_|   _|    _|       _|_|_|     _|_|_| _|      _|   _|_|         _|    _| 
    _|     _|    _| _|    _|       _|    _| _|    _| _|      _| _|_|_|_|         _|_|   
    _|     _|    _| _|    _|       _|    _| _|    _|   _|  _|   _|             _|    _| 
    _|       _|_|     _|_|_|       _|    _|   _|_|_|     _|       _|_|_|         _|_|   
```

- Le Thinkertoy:
```console

o   o                o                          o-o  
 \ /                 |                         |   | 
  O   o-o o  o       O--o  oo  o   o o-o        o-o  
  |   | | |  |       |  | | |   \ /  |-'       |   | 
  o   o-o o--o       o  o o-o-   o   o-o        o-o  

```