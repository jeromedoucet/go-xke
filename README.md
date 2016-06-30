# Implementation

## Module d'enregistrement

La premiere etape consiste a s'enregistrer aupres des clients. Pour cela, il va falloir completer l'implementation de 
registration.go.

### cas nominal

Lancez le premier test a l'aide de la commande suivante:
go test github.com/vil-coyote-acme/go-xke/registration -run ^Test_register_should_return_no_error_on_200_response$

Il s'agit du cas nominal d'un enregistrement. Vous devrez pour cela completer la creation de l'instance de registration
en precisant les attributs PlayerId et Ip.

    AIDE : lorsque l'on utilise une instanciation literral de la forme t := myType{},
    vous pouvez affecter les attributs du type que vous souhaitez comme ceci : 
    t := myType{myStringField: "a string value", myIntValue: 42}

Une fois ceci fait, vous devrez utiliser le package json du sdk pour serialiser registration.
    
    AIDE : lorsque l'on doit dependre d'un autre package, il faut alors 
    l'ajouter en tant qu'import. Si vous utilisez intellij avec le plugin go,
    l'ide vous proposera automatiquement d'ajouter l'import, sinon 
    il vous faudra ajouter a la main "encoding/json" dans les imports
    
    AIDE : vous pouvez remarquer que la fonction de serialisation du 
    package json vous renvois deux valeurs. La premiere est votre structure
    de donnee serialise presente sous forme d'un tableau d'octet, la seconde
    est valeur de type 'error'. Si cette derniere n'est pas nil, 
    cela signifie que la serialisation a rencontre un probleme. Si vous n'exploitez
    pas cette valeur (ce qui est ici le cas), vous pouvez choisir de l'ignorer
    en la remplacant par un '_'. Exemple :
    valeur, error := appelFonction()
    valeur, _ := appelFonction()
    
Enfin, il vous faudra faire un appel http en POST en utilisant clientRegistrationURL comme url. La fonction Post se trouve
dans le package "net/http" (usage http.Post(....)) et necessitera de creer un 'Reader' a partir de la structure 
registration serialise. Pour ce faire utiliser NewBuffer(myByteArray) du package "bytes".

### cas de rejet de la requete de la part du client

Si le nom de joueur que vous avez choisi est deja pris, ou si vous essayez de vous enregistrer plusieurs fois 
(et donc de truander !) le client peut rejeter votre demande d'enregistrement.

Le test correspondant est le suivant :
go test github.com/vil-coyote-acme/go-xke/registration -run ^Test_register_should_return_no_error_on_non_200_response$

Il vous faudra alors exploiter la premiere valeur retourne par http.Post(...) qui possede un attribut 'StatusCode'
contenant le code retour http de la requete. 

Si le code retour n'est pas 200, il vous faudra alors retourner une error, a l'aide de la fonction New("string message")
du package "errors"

    AIDE: syntaxe 'if' : if value.intAttribut != 42 {return ...}

### cas d'echec de la requete 

Le test correspondant est le suivant :
go test github.com/vil-coyote-acme/go-xke/registration -run ^Test_register_should_return_error_when_connection_issue$

Il ne vous reste ici qu'a considerer le cas ou la valeur de retour de type 'error' du Post est non nil. Si tel est le cas,
 il vous faudra retourner cette error !
 

## Module Serveur

Il s'agit du module principal de votre programme. Il recevra les commandes envoyees par les clients, les enverra au
barman puis se fera payer en retour aupres des clients.

### Requete vers le barman

Les premiers tests sont les suivants :
go test github.com/vil-coyote-acme/go-xke/server -run ^Test_postOrder_should_fail$
go test github.com/vil-coyote-acme/go-xke/server -run ^Test_postOrder_should_do_without_error$

Il porte sur l'envoi de la commande au barman et consiste a implementer la fonction postOrder et la requete Post que
vous devrez faire est tres similaire a ce que vous avez deja fait dans le module registration

### Payment

Les premiers tests sont les suivants :
go test github.com/vil-coyote-acme/go-xke/server -run ^Test_getDataFromCallback_should_fail_with_error_in_url$
go test github.com/vil-coyote-acme/go-xke/server -run ^Test_getDataFromCallback_should_not_fail$

Il s'agit d'implementer la fonction getDataFromCallback. Il vous faudra faire un Get sur l'url de callback contenu dans **l'order**, afin de 'recuperer'
votre du. Un composant tiers se chargera de compter votre score

# lancement de votre programme :

Se postionner sous $GOPATH/src/github.com/vil-coyote-acme/go-xke 

puis lancer go run xke-app.go -clientIp={{clientIp}} -ourIp={{ourIp}} -bartenderIp={{bartenderIp}} -playerId={{playerId}}

Avec clientIp et bartenderIp qui seront une valeur ip:port fourni lors de l'exercice, ourIp qui sera votre Ip sur le reseau 
au moment de l'exercice et playerId qui sera votre nom de joueur.

Bonne chance !
 

