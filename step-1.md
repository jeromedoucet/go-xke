---
layout: default
title: Step-1
---

# Implementation

Avant de commencer avec le Step-1 assurez-vous d'avoir executé sur le code de la branche:

```sh
$ go get github.com/vil-coyote-acme/go-concurrency/commons
$ go get github.com/stretchr/testify/assert
```

## Module d'enregistrement

La première étape consiste à s'enregistrer auprès des clients. Pour cela, il va falloir compléter l'implémentation de *registration.go*.

### Cas nominal

Lancez le premier test à l'aide de la commande suivante:

```
$ go test github.com/vil-coyote-acme/go-xke/registration -run ^Test_register_should_return_no_error_on_200_response$
```

Modifiez le fichier *registration.go*

Il s'agit du cas nominal d'un enregistrement. Vous devrez pour cela, compléter la création de l'instance de *registration* en précisant les attributs *PlayerId* et *Ip*.

<blockquote class = 'help' markdown="1">
    lorsque l'on utilise une instanciation literral de la forme t := myType{},
    vous pouvez affecter les attributs du type que vous souhaitez comme ceci : 
    t := myType{myStringField: "a string value", myIntValue: 42}
</blockquote>

Une fois ceci fait, vous devrez utiliser le package [*json*](https://golang.org/pkg/encoding/json/) du sdk pour sérialiser *registration*.

<blockquote class = 'help' markdown="1">    
    1. lorsque l'on doit dépendre d'un autre package, il faut alors 
    l'ajouter en tant qu'import. Si vous utilisez intellij avec le plugin go,
    l'ide vous proposera automatiquement d'ajouter l'import, sinon 
    il vous faudra ajouter à la main "encoding/json" dans les imports
   
    2. vous pouvez remarquer que la fonction de serialisation du 
    package *json* vous renvois deux valeurs. La première est votre structure
    de données sérialisé présente sous forme d'un tableau d'octet, la seconde
    est valeur de type '*error*'. Si cette dernière n'est pas nil, 
    cela signifie que la serialisation a rencontré un problème. Si vous n'exploitez
    pas cette valeur (ce qui est ici le cas), vous pouvez choisir de l'ignorer
    en la remplaçant par un '_'. Exemple :
    valeur, error := appelFonction()
    valeur, _ := appelFonction()
</blockquote>
    
Enfin, il vous faudra faire un appel http en POST en utilisant la variable *clientRegistrationURL* comme url. La fonction Post se trouve
dans le package [*net/http*](https://golang.org/pkg/net/http/).  Vous aurez besoin de créer un *'Reader'* à partir de la structure 
registration sérialisé. Pour ce faire, utiliser *NewBuffer(myByteArray)* du package [*bytes*](https://golang.org/pkg/bytes/).

<blockquote class = 'help' markdown="1">
    http.Post(clientRegistrationURL, "application/json", bytes.NewBuffer(body))
</blockquote>

### Cas de rejet de la requete de la part du client

Si le nom de joueur que vous avez choisi est déjà pris, ou si vous essayez de vous enregistrer plusieurs fois 
(et donc de truander !) le client peut rejeter votre demande d'enregistrement.

Le test correspondant est les suivant :

```
$ go test github.com/vil-coyote-acme/go-xke/registration -run ^Test_register_should_return_no_error_on_non_200_response$
``` 

Il vous faudra alors exploiter la première valeur retourne par http.Post(...) qui possède un attribut *StatusCode*
contenant le code retour *http* de la requête. 

Si le code retour n'est pas 200, il vous faudra alors retourner un error, a l'aide de la fonction New("string message")
du package *errors*

<blockquote class = 'help' markdown="1">
     syntaxe 'if' : if value.intAttribut != 42 {return ...}
</blockquote>

### cas d'échec de la requête 

Le test correspondant est le suivant :

```
$ go test github.com/vil-coyote-acme/go-xke/registration -run ^Test_register_should_return_error_when_connection_issue$
```

Il ne vous reste ici qu'à considérer le cas ou la valeur de retour de type *'error'* du Post est non nil. Si tel est le cas, il vous faudra retourner cet *error* !

<blockquote class = 'help' markdown="1">

if httpError != nil {
        return httpError
    }
</blockquote>
 

## Module Serveur

Il s'agit du module principal de votre programme. Il recevra les commandes envoyées par les clients, les transformera de JSON à *Order*, les enverra au
barman, puis se fera payer en retour auprès des clients. 

<blockquote class = 'help' markdown="1">
Modifiez le fichier go-xke/server/server.go
Comencez pour les  *TODO* de la methode *handleOrder*
</blockquote>

### Transformer les commandes envoyées par les clients

Utilisez le package [*io*](https://golang.org/pkg/io/) pour lire la *order* du *body* du *Request*
et utilisez le package [*json*](https://golang.org/pkg/encoding/json/) pour transformer ce message vers un element de type *commons.Order*

<blockquote class = 'help' markdown="1">

Vous pouvez déclarer une variable de type *commons.Order* et vous aider de *&* pour passer la réference (son position de memoire) de la variable
à la méthode *Unmarshal*. 
Utilisez json.Unmarshal(buf, &order)

</blockquote>

### Requête vers le barman

Vous pouvez implementer l'envoi de la commande au barman dans la fonction *postOrder*.

Pour vous aider, nous vous avons fourni des tests:

```
$ go test github.com/vil-coyote-acme/go-xke/server -run ^Test_postOrder_should_fail$
$ go test github.com/vil-coyote-acme/go-xke/server -run ^Test_postOrder_should_do_without_error$
````

La requête *Post* que vous devrez faire est similaire à cela que vous avez déjà fait dans le module *registration*.

Utilisez le package [*json*](https://golang.org/pkg/encoding/json/) du sdk pour sérialiser l'ordre avant de
faire un appel http en POST, en utilisant la variable *bartenderUrl* comme url. Rappelez-vous que la fonction Post se trouve
dans le package [*net/http*](https://golang.org/pkg/net/http/).

Retourner la réponse de l'appel POST.

<blockquote class = 'help' markdown="1">

return http.Post(bartenderUrl, "application/json", bytes.NewBuffer(marshalledMessage))

</blockquote>

### Payment

Implementez la fonction *getDataFromCallback*.

Vous pouvez tester le fonctionnement de cette fonction travers les tests:

```
$ go test github.com/vil-coyote-acme/go-xke/server -run ^Test_getDataFromCallback_should_fail_with_error_in_url$
$ go test github.com/vil-coyote-acme/go-xke/server -run ^Test_getDataFromCallback_should_not_fail$
```

 Il vous faudra faire un *Get* sur l'url de callback (*order.CallBackUrl*) contenu dans **l'order**, afin de 'recuperer'
votre dû.

<blockquote class = 'help' markdown="1">

paymentRes, paymentErr := http.Get(order.CallBackUrl)

</blockquote>

 Un composant tiers se chargera de compter votre score

# Lancement de votre programme :

Bravo! 
vous avez une première version du serveur. Pour le tester, postionnez-vous sous *$GOPATH/src/github.com/vil-coyote-acme/go-xke*, puis lancez: 

```
$ go run xke-app.go -clientIp={{clientIp}} -ourIp={{ourIp}} -bartenderIp={{bartenderIp}} -playerId={{playerId}}
```

Avec *clientIp* et *bartenderIp* qui seront une valeur ip:port fourni lors de l'exercice, *ourIp* qui sera votre *Ip* sur le réseau au moment de l'exercice et *playerId* qui sera votre nom de joueur.

Bonne chance !

#### Plus d’infos :
- [https://golang.org/pkg/net/http/](https://golang.org/pkg/net/http/)
- [https://golang.org/pkg/io/](https://golang.org/pkg/io/)
- [https://golang.org/pkg/encoding/json/](https://golang.org/pkg/encoding/json/)
 
#### La suite

Basculez sur la branche step-2 :

```
$ git checkout step-2
```

la suite par ici: [step-2](step-2) 

