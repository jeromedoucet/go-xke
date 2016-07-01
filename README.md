# Implementation

## Module d'enregistrement

La première étape consiste à s'enregistrer auprès des clients. Pour cela, il va falloir compléter l'implémentation de *registration.go*.

### cas nominal

Lancez le premier test à l'aide de la commande suivante:

```
$ go test github.com/vil-coyote-acme/go-xke/registration -run ^Test_register_should_return_no_error_on_200_response$
```

Il s'agit du cas nominal d'un enregistrement. Vous devrez pour cela, compléter la création de l'instance de *registration* en précisant les attributs *PlayerId* et *Ip*.

    AIDE : lorsque l'on utilise une instanciation literral de la forme t := myType{},
    vous pouvez affecter les attributs du type que vous souhaitez comme ceci : 
    t := myType{myStringField: "a string value", myIntValue: 42}

Une fois ceci fait, vous devrez utiliser le package json du sdk pour serialiser *registration*.
    
    AIDE : lorsque l'on doit dépendre d'un autre package, il faut alors 
    l'ajouter en tant qu'import. Si vous utilisez intellij avec le plugin go,
    l'ide vous proposera automatiquement d'ajouter l'import, sinon 
    il vous faudra ajouter à la main "encoding/json" dans les imports
    
    AIDE : vous pouvez remarquer que la fonction de serialisation du 
    package *json* vous renvois deux valeurs. La première est votre structure
    de données sérialisé présente sous forme d'un tableau d'octet, la seconde
    est valeur de type '*error*'. Si cette dernière n'est pas nil, 
    cela signifie que la serialisation a rencontré un problème. Si vous n'exploitez
    pas cette valeur (ce qui est ici le cas), vous pouvez choisir de l'ignorer
    en la remplaçant par un '_'. Exemple :
    valeur, error := appelFonction()
    valeur, _ := appelFonction()
    
Enfin, il vous faudra faire un appel http en POST en utilisant clientRegistrationURL comme url. La fonction Post se trouve
dans le package "net/http" (usage http.Post(....)) et nécessitera de créer un *'Reader'* à partir de la structure 
registration sérialisé. Pour ce faire, utiliser *NewBuffer(myByteArray)* du package *'bytes'*.

### cas de rejet de la requete de la part du client

Si le nom de joueur que vous avez choisi est déjà pris, ou si vous essayez de vous enregistrer plusieurs fois 
(et donc de truander !) le client peut rejeter votre demande d'enregistrement.

Le test correspondant sont les suivants :

```
$ go test github.com/vil-coyote-acme/go-xke/registration -run ^Test_register_should_return_no_error_on_non_200_response$
```

Il vous faudra alors exploiter la première valeur retourne par http.Post(...) qui possède un attribut *'StatusCode'*
contenant le code retour *http* de la requête. 

Si le code retour n'est pas 200, il vous faudra alors retourner un error, a l'aide de la fonction New("string message")
du package *'errors'*

    AIDE: syntaxe 'if' : if value.intAttribut != 42 {return ...}

### cas d'échec de la requête 

Le test correspondant est le suivant :

```
$ go test github.com/vil-coyote-acme/go-xke/registration -run ^Test_register_should_return_error_when_connection_issue$
```

Il ne vous reste ici qu'à considérer le cas ou la valeur de retour de type *'error'* du Post est non nil. Si tel est le cas, il vous faudra retourner cet *error* !
 

## Module Serveur

Il s'agit du module principal de votre programme. Il recevra les commandes envoyées par les clients, les enverra au
barman puis se fera payer en retour auprès des clients.

### Requête vers le barman

Les premiers tests sont les suivants :
```
$ go test github.com/vil-coyote-acme/go-xke/server -run ^Test_postOrder_should_fail$
$ go test github.com/vil-coyote-acme/go-xke/server -run ^Test_postOrder_should_do_without_error$
````

Il porte sur l'envoi de la commande au barman et consiste à implementer la fonction *postOrder* et la requête *Post* que
vous devrez faire est très similaire à ce que vous avez déjà fait dans le module *registration*

### Payment

Les premiers tests sont les suivants :

```
$ go test github.com/vil-coyote-acme/go-xke/server -run ^Test_getDataFromCallback_should_fail_with_error_in_url$
$ go test github.com/vil-coyote-acme/go-xke/server -run ^Test_getDataFromCallback_should_not_fail$
```

Il s'agit d'implémenter la fonction *getDataFromCallback*. Il vous faudra faire un *Get* sur l'url de callback contenu dans **l'order**, afin de 'recuperer'
votre dû. Un composant tiers se chargera de compter votre score

# lancement de votre programme :

Se postionner sous **$GOPATH/src/github.com/vil-coyote-acme/go-xke **

puis lancer 

```
$ go run xke-app.go -clientIp={{clientIp}} -ourIp={{ourIp}} -bartenderIp={{bartenderIp}} -playerId={{playerId}}
```

Avec *clientIp* et *bartenderIp* qui seront une valeur ip:port fourni lors de l'exercice, *ourIp* qui sera votre *Ip* sur le réseau  au moment de l'exercice et *playerId* qui sera votre nom de joueur.

Bonne chance !
 

