# go-xke

Bienvenu dans ce hands'on de programmation concurrente avec go !

## rappels

Afin de bien commencer ce hands'on, il est imperatif de s'assurer que votre environnement go soit bien initialisé.

-- todo : a finir !

## but de ce hand'on

Son but est tout d'abord la découverte de go et de son écosysteme en douceur. Il est donc découpé en deux partie.

- finir l'implémentation de waiter.go
- optimiser waiter.go à l'aide des goroutines et des channels

## 1) finir l'implémentation

Ce hands'on est fourni avec des tests unitaires et des indications qui vont vous guider.

Commençez par lancer la commande "go get ./..." dans le repertoire racine des sources du hands'on. Ceci aura pour effet
de télécharger toutes les dépendances nécessaires.

Puis lancez la commande "go test go-xke", 3 tests seront alors en erreur. À vous de les faire passer !

Si tout se passe bien, lorsque les tests passent, votre waiter est pret !

Il vous reste alors à modifier la variable "host" avec l'ip qui vous sera fourni et à lancer la commande
"go run waiter.go -player=nomDeJoueur"

## 2) optimiser waiter.go

Vous pouvez remarquer maintenant que le Bar peut mettre beaucoup de temps à répondre. Ce temps de réponse est lié au nombre
de point qu'un type de commande donne peut rapporter.

le but du jeu étant d'être plus efficace que les autres participants, il va falloir utiliser intelligement les goroutine et les channels.

 
Bonne chance !




