# Optimisation

Votre code fontionne, felicitations ! Mais il s'agit maintenant de l'ameliorer. Comme vous avez pu vous en rendre compte,
vous etes evalue sur votre efficacite a traiter les commandes, et pour le moment tout est sequentiel ! Ce n'est pas tres 
efficace.

En effet, le barman ne traite pas immediatement votre requete, et plus celle-ci vaut chere, plus il prend son temps! Il
s'agit neanmoins d'un excellent barman qui est capable de traiter 5 commande maximum par serveur en meme temps.
Mais si jamais un serveur essayait de depasser cette limite, il se verrait immediatement et lourdement sanctionne !

Vous aller ici devoir vous servir des goroutines et des channels. Le serveur fonctionne encore, mais des indications
sous la forme de TODO ont ete laisse dans le code pour vous guider.

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

