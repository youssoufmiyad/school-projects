# Projet Chip-8 en Golang

![Alt text](screenshots/test1.png?raw=true "Logo")

## Description

Ce projet a été réalisé dans le cadre de mon cursus en B2 Informatique à Ynov Campus Paris.

L'objectif était d'émulater un système Chip-8 en utilisant le langage de programmation Golang, nous avons opter pour la bibliothèque graphique Ebiten.

## Prérequis

Avant de lancer le projet, assurez-vous d'avoir installé les éléments suivants sur votre système :

- Golang
- Ebiten : github.com/hajimehoshi/ebiten/v2

## Compilez et exécutez le programme :

Une fois dans le répertoire :
```
go run main.go
```

***Chargement d'un jeu :***

Une fois l'émulateur en cours d'exécution, il vous sera demandé de choisir parmis 7 ROMS dans la console, l'émulateur lancera ensuite la ROM de votre choix.

![Alt text](screenshots/launch.png?raw=true "Launch") ![Alt text](screenshots/secretLaunch.png?raw=true "Code launch")


Contrôles du Chip-8 :

Les touches du clavier sont émulées par les touches hexadécimales du Chip-8 (0-F).

1 2 3 4&emsp;1 2 3 C<br />
A Z E R&emsp;4 5 6 D<br />
Q S D F&emsp;7 8 9 E<br />
W X C V&emsp;A 0 B F<br />

![Alt text](screenshots/breakout.png?raw=true "Game")

Pour quitter appuyer : [échap]

## Fonctionnalités

Émulation de la machine virtuelle Chip-8 : Nôtre émulateur est capable d'interpréter et d'exécuter des programmes écrits en langage de la machine Chip-8.

Interface graphique : Utilisation de la bibliothèque Ebiten pour afficher l'écran de la machine Chip-8 et rendre les graphiques du jeu.

Clavier virtuel : Mise en place d'une interface utilisateur pour émuler le clavier de la machine Chip-8. Les touches du clavier d'origine étaient hexadécimales (0-F).

Chargement des ROMs : Capacité à charger des fichiers ROM Chip-8 qui contiennent les programmes à exécuter.

Gestion des instructions : L'émulateur est capable de déchiffrer et d'exécuter les différentes instructions du jeu.


