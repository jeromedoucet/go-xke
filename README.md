# Optimisation

Votre code fontionne, félicitations ! Mais il s'agit maintenant de l'améliorer. Comme vous avez pu vous en rendre compte,
vous êtes evalué sur votre efficacité a traiter les commandes, et pour le moment tout est séquentiel ! Ce n'est pas très 
efficace.

En effet, le barman ne traite pas immédiatement votre requête, et plus celle-ci vaut chère, plus il prend son temps! Il
s'agit néanmoins d'un excellent barman qui est capable de traiter 5 commande maximum par serveur en meme temps.
Mais si jamais un serveur essayait de dépasser cette limite, il se verrait immédiatement et lourdement sanctionné !

Vous aller ici devoir vous servir des goroutines et des channels. Le serveur fonctionne encore, mais des indications
sous la forme de TODO ont ete laissé dans le code pour vous guider.

Rappel :
```go
// initialisation d'un channel bloquant de Poney
c:= make(chan Poney)

// pousser une valeur dans un channel
c <- Poney{Name: "mon petit poney"}

// consommer une valeur dans un channel
v := <- c

// faire executer une fonction en concurrence :
go maFonction()

```

# Pour aller plus loin...

```go
// initialisation d'un channel de Poney non bloquant de taille n :
c := make(chan Poney, n)

// utilisation du mot-cle 'select' pour 'choisir' un channel sur lequel consommer
select {
    case v1 := <- c1:
        // faire un truc avec v1 venant du channel c1
    case v2 := <- c2:
        // faire un autre truc avec v2 venant du channel c2
    default:
        // faire un truc par defaut si on ne peut ni consommer du c1 ni du c2
}

```

