# The Bartender
## Objectif 
Cet atelier a pour objectif de vous faire coder en go et de découvrir le langage.

Le principe est simple. Nous vous propossons un jeu où vous êtes un client qui entre dans un bar.  Vous devez commander et payer vos boisons. Nous avons codé pour vous le  bartender, le order maker et le server. On vous donne des méthodes à implementer et des test unitaires.

## Les pre-requis 
 - installation de Go : https://golang.org/dl/
 - creation du workspace Go : un repertoire contenant src/ , pkg/  et  bin/
 - export de la variable $GOPATH pointant vers le workspace Go
 - create un repertoire $GOPATH/src/github.com/vil-coyote-acme/ 
 - executer la commande:

```sh
$ git clone https://github.com/vil-coyote-acme/go-xke.git
```

### 1. Le server et la communication http et parsing json
Basculer sur la branche step1 : 

```sh
$ git checkout step-1
```

Modifier les méthodes de **server.go**. Utiliser les **TODO** dans le code pour faire passer les tests

#### Plus d’infos :
- https://golang.org/pkg/net/http/
- https://golang.org/pkg/io/
- https://golang.org/pkg/encoding/json/
