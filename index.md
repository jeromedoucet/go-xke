---
layout: default
title: Programation concurrente GO
---


# Le serveur X

## Objectif 

Cet atelier a pour objectif de vous faire coder en go et de découvrir le langage.

Le principe est simple. Nous vous proposons un jeu où vous êtes un serveur dans un bar.  Vous devez passer les commandes des clients et les faire payer ses boisons. Nous avons codé pour vous le  *bartender*, le *ordermaker* et le *client*. On vous donne le serveur avec des méthodes à implémenter et des tests unitaires.

## Les pre-requis

 - Avoir git installé
 - Installer Go : [https://golang.org/dl/](https://golang.org/dl/)
 - Creation du workspace Go : un repertoire contenant src/ , pkg/  et  bin/
 - Export de la variable $GOPATH pointant vers le workspace Go
 - Creation d'un repertoire $GOPATH/src/github.com/vil-coyote-acme/ 
 - À l'interieur du repertoire $GOPATH/src/github.com/vil-coyote-acme/ cloner le repository [https://github.com/vil-coyote-acme/go-xke.git](https://github.com/vil-coyote-acme/go-xke.git) et executer les commandes:

```sh
$ git clone https://github.com/vil-coyote-acme/go-xke.git
$ cd go-xke
$ ./init.sh
```

### 1. Le server, la communication http et parsing json
Basculez sur la branche step-1 : 

```sh
$ git checkout step-1
```

Allez sur la console et executez

```sh
$ go get github.com/vil-coyote-acme/go-concurrency/commons
$ go get github.com/stretchr/testify/assert
```

Modifiez les méthodes de **server.go** et **registration.go**. Utilisez les **TODO** dans le code pour faire passer les tests

La suite par ici:
[step-1](step-1)

#### Plus d’infos :
- [https://golang.org/pkg/net/http/](https://golang.org/pkg/net/http/)
- [https://golang.org/pkg/io/](https://golang.org/pkg/io/)
- [https://golang.org/pkg/encoding/json/](https://golang.org/pkg/encoding/json/)

### 2. Les goroutine, les channels et la communication concurrente

Basculez sur la branche step-2 : 

```sh
$ git checkout step-2
```

La suite par ici:
[step-2](step-2)


