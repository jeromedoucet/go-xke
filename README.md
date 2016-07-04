# Le serveur'X
## Objectif 
Cet atelier a pour objectif de vous faire coder en go et de découvrir le langage.

Le principe est simple. Nous vous proposons un jeu où vous êtes un serveur dans un bar.  Vous devez passer les commandes des clients et les faire payer ses boisons. Nous avons codé pour vous le  bartender, le ordermaker et le client. On vous donne le serveur avec des méthodes à implémenter et des tests unitaires.

## Les pre-requis 
 - installation de Go : https://golang.org/dl/
 - creation du workspace Go : un repertoire contenant src/ , pkg/  et  bin/
 - export de la variable $GOPATH pointant vers le workspace Go
 - create un repertoire $GOPATH/src/github.com/vil-coyote-acme/ 
 - executer les commande:

```sh
$ git clone https://github.com/vil-coyote-acme/go-xke.git
$ cd go-xke
$ ./init.sh
```

### 1. Le server et la communication http et parsing json
Basculer sur la branche step-1 : 

```sh
$ git checkout step-1
```

Allez sur la console et executez

```sh
$ go get github.com/vil-coyote-acme/go-concurrency/commons
$ go get github.com/stretchr/testify/assert
```


Modifier les méthodes de **server.go** et **registration.go**. Utiliser les **TODO** dans le code pour faire passer les tests

La suite par ici:
https://github.com/vil-coyote-acme/go-xke/tree/step-1

#### Plus d’infos :
- https://golang.org/pkg/net/http/
- https://golang.org/pkg/io/
- https://golang.org/pkg/encoding/json/

### 2. Les goroutine, les channels et la communication concurrente

Basculez sur la branche step-2 : 

```sh
$ git checkout step-2
```

La suite par ici:
https://github.com/vil-coyote-acme/go-xke/tree/step-2
