---
layout: default
title: Programation concurrente GO
---

# Optimisation

Votre code fonctionne, félicitations ! Mais il faut maintenant l'améliorer.

Comme vous avez pu vous en rendre compte, vous êtes évalué sur votre efficacité à traiter les commandes, et pour le moment vous les traitez de manière séquentielle ! Ce n'est pas très efficace.

En effet, le barman ne traite pas immédiatement votre requête, et plus celle-ci vaut chère, plus il prend son temps! Néanmoins, 
nous avons un excellent barman qui est capable de traiter 5 commande maximum par serveur en même temps.
Mais si jamais un serveur essayait de dépasser cette limite, il se verrait immédiatement et lourdement sanctionné !

Vous allez ici devoir vous servir des [goroutines](https://www.golang-book.com/books/intro/10#section1) et des [channels](https://www.golang-book.com/books/intro/10#section2). Le serveur fonctionne encore, mais des indications sous la forme de *TODO* ont été laissés dans le code pour vous guider.

Nous vous proposons d'utiliser le channel *orderChan* du *server* pour gérer les commandes du client.
<blockquote class = 'help' markdown="1">

```go
// pousser une valeur dans un channel
c <- Poney{Name: "mon petit poney"}

```
</blockquote>

Assurez-vous que le channel a été bien declaré avant de l'utiliser.

<blockquote class = 'help' markdown="1">

```go
// initialisation d'un channel bloquant de Poney
c:= make(chan Poney)

```

</blockquote>

Vous pouvez traiter plusieurs commandes de manière concurrente, mais rappelez-vous de vous limiter à 5.

<blockquote class = 'help' markdown="1">

Aidez-vous des go routines pour creer les travails qui consomment des *orders* du channel de manière concurrente.

```go

for i := 0; i < 5; i ++ {
	go consumeOrder()
}

```

</blockquote>

Maintenant, chaque consommateur puisse prendre l'ordre envoyé au channel et l'envoyer au barman.

<blockquote class = 'help' markdown="1">
```go
// consommer une valeur dans un channel
v := <- c

```
</blockquote>

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

#### Plus d’infos :
- [https://golang.org/pkg/net/http/](https://golang.org/pkg/net/http/)
- [https://golang.org/pkg/io/](https://golang.org/pkg/io/)
- [https://golang.org/pkg/encoding/json/](https://golang.org/pkg/encoding/json/)



