package PROJETRED

import (
	"fmt"
	"time"
)

func (p *Perso) TrainingFight() {
	c := Monstre{"Carrie", 40, 40, 5}
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("			███████╗███╗   ██╗████████╗██████╗  █████╗ ██╗███╗   ██╗███████╗███╗   ███╗███████╗███╗   ██╗████████╗ ")
	fmt.Println("			██╔════╝████╗  ██║╚══██╔══╝██╔══██╗██╔══██╗██║████╗  ██║██╔════╝████╗ ████║██╔════╝████╗  ██║╚══██╔══╝ ")
	fmt.Println("			█████╗  ██╔██╗ ██║   ██║   ██████╔╝███████║██║██╔██╗ ██║█████╗  ██╔████╔██║█████╗  ██╔██╗ ██║   ██║  ")
	fmt.Println("			██╔══╝  ██║╚██╗██║   ██║   ██╔══██╗██╔══██║██║██║╚██╗██║██╔══╝  ██║╚██╔╝██║██╔══╝  ██║╚██╗██║   ██║    ")
	fmt.Println("			███████╗██║ ╚████║   ██║   ██║  ██║██║  ██║██║██║ ╚████║███████╗██║ ╚═╝ ██║███████╗██║ ╚████║   ██║    ")
	fmt.Println("			╚══════╝╚═╝  ╚═══╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝╚══════╝╚═╝     ╚═╝╚══════╝╚═╝  ╚═══╝   ╚═╝    ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("Salut à toi ", p.name, "le", p.class)
	fmt.Println("Nous avons capturés plusieurs Carries pour pouvoir t'entraîner avant d'aller t'aventurer dans le forêt Noir.")
	fmt.Println(" ")
	fmt.Println("1 : Pour combattre la Carries d'entrainement ! ")
	fmt.Println("2 : Pour retourner à PATISSIA")
	fmt.Println(" ")
	fmt.Printf("Entre ton choix :  ")
	fmt.Scan(&choice)
	fmt.Println(" ")

	switch choice {
	case "1":
		fmt.Println("\033[H\033[2J")
		c.Fight(p)
	case "2":
		fmt.Println("\033[H\033[2J")
		p.patissia()
	default:
		fmt.Println("\033[H\033[2J")
		fmt.Println("Recommence mon donuts sucré au sucre ! ")
		p.TrainingFight()
	}
}

func (c *Monstre) Fight(p *Perso) {
	fmt.Println("			███████╗██╗ ██████╗ ██╗  ██╗████████╗    ██╗██╗██╗ ")
	fmt.Println("			██╔════╝██║██╔════╝ ██║  ██║╚══██╔══╝    ██║██║██║ ")
	fmt.Println("			█████╗  ██║██║  ███╗███████║   ██║       ██║██║██║ ")
	fmt.Println("			██╔══╝  ██║██║   ██║██╔══██║   ██║       ╚═╝╚═╝╚═╝ ")
	fmt.Println("			██║     ██║╚██████╔╝██║  ██║   ██║       ██╗██╗██╗ ")
	fmt.Println("			╚═╝     ╚═╝ ╚═════╝ ╚═╝  ╚═╝   ╚═╝       ╚═╝╚═╝╚═╝ ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	c.CharCreationMonstre()
	fmt.Println("Attention prepare toi, tu vas rentrer dans la cage en chocolat. ")
	tourDeCombat := 1
	for !(c.pvnow <= 0 || p.pvnow <= 0) {
		fmt.Println("vous êtes au tour:", tourDeCombat)

		if tourDeCombat%3 == 0 {
			fmt.Println(" ")
			fmt.Println(" ")
			time.Sleep(2 * time.Second)
			fmt.Print(" Attention, carrie s'enerve ! Elle inflige le double de ses dégats ce tour !")
			fmt.Println(" ")
			degats := 5 * 2
			time.Sleep(2 * time.Second)
			fmt.Println("- 10 HP")
			p.pvnow -= degats
			time.Sleep(3 * time.Second)
			p.Dead()
		} else {
			fmt.Println(" ")
			fmt.Println(" ")
			time.Sleep(2 * time.Second)
			fmt.Println("Le monstre attaque !")
			fmt.Println(" ")
			time.Sleep(2 * time.Second)
			fmt.Println("- 5 HP")
			p.pvnow = p.pvnow - 5
			time.Sleep(3 * time.Second)
		}

		p.Dead()
		fmt.Println("Il te reste ", p.pvnow, "HP /", p.pvmax, "HP")
		fmt.Println(" ")
		fmt.Println("Il reste a cette Carrie", c.pvnow, "HP /", c.pvmax, "HP")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println("A ton tour ! ")
		fmt.Println(" ")
		fmt.Println("1 : Frappe Chocolatée (-7 HP à votre adversaire)")
		fmt.Println("2 : Épée en sucre glace (-10 HP à votre adversaire)")
		fmt.Println(" ")
		fmt.Printf("Choisie ton attaque : ")
		fmt.Scan(&choice)
		fmt.Println(" ")

		switch choice {
		case "1":
			fmt.Println("\033[H\033[2J")
			if p.LookAttaque("Frappe Chocolatée") {
				fmt.Println("Tu lui à décroché un pain au chocolat !")
				fmt.Println("Carrie : - 7 HP")
				c.pvnow -= 7
				fmt.Println(c.pvnow, "/", c.pvmax)
				c.DeadMonstre()
			}
		case "2":
			fmt.Println("\033[H\033[2J")
			if p.LookAttaque("Épée en sucre glace") {
				fmt.Println("Tu lui à envoyé un coup d'épée sucrement parfait !!")
				fmt.Println("Carrie : - 10 HP")
				c.pvnow -= 10
				fmt.Println(c.pvnow, "/", c.pvmax)
				c.DeadMonstre()
			}
		default:
			fmt.Println("\033[H\033[2J")
			fmt.Println("Dommage pour toi tu à passé ton tour...")

		}
		tourDeCombat++
	}
	if c.pvnow <= 0 {
		fmt.Println("\033[H\033[2J")
		fmt.Println("			██╗    ██╗██╗███╗   ██╗    ██╗██╗██╗ ")
		fmt.Println("			██║    ██║██║████╗  ██║    ██║██║██║ ")
		fmt.Println("			██║ █╗ ██║██║██╔██╗ ██║    ██║██║██║ ")
		fmt.Println("			██║███╗██║██║██║╚██╗██║    ╚═╝╚═╝╚═╝ ")
		fmt.Println("			╚███╔███╔╝██║██║ ╚████║    ██╗██╗██╗ ")
		fmt.Println("			 ╚══╝╚══╝ ╚═╝╚═╝  ╚═══╝    ╚═╝╚═╝╚═╝ ")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println("Bravo !! Tu a réussi à tuer cette Carrie.")
		fmt.Println(" ")
		fmt.Println("Il te reste : ", p.pvnow, "/", p.pvmax)
		fmt.Println(" ")
		fmt.Println("1 : Pour recombattre une Carrie")
		fmt.Println("2 : Retourner dans la fôret Noir")
		fmt.Scan(&choice)
		fmt.Println(" ")

		switch choice {
		case "1":
			fmt.Println("\033[H\033[2J")
			c.Fight(p)
		case "2":
			fmt.Println("\033[H\033[2J")
			p.ForetNoir()
		default:
			fmt.Println("\033[H\033[2J")
			p.ForetNoir()

		}
	}
	if p.pvnow <= 0 {
		fmt.Println("\033[H\033[2J")
		fmt.Println("			██╗      ██████╗ ███████╗███████╗ ")
		fmt.Println("			██║     ██╔═══██╗██╔════╝██╔════╝  ")
		fmt.Println("			██║     ██║   ██║███████╗█████╗ ")
		fmt.Println("			██║     ██║   ██║╚════██║██╔══╝   ")
		fmt.Println("			███████╗╚██████╔╝███████║███████╗██╗██╗██╗")
		fmt.Println("			╚══════╝ ╚═════╝ ╚══════╝╚══════╝╚═╝╚═╝╚═╝")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println("Dommage !! Tu n'a pas réussi à tuer cette Carrie.")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println("1 : Pour recombattre la Carrie")
		fmt.Println("2 : Retourner dans la fôret Noir")
		fmt.Scan(&choice)
		fmt.Println(" ")

		switch choice {
		case "1":
			fmt.Println("\033[H\033[2J")
			c.Fight(p)
		case "2":
			fmt.Println("\033[H\033[2J")
			p.ForetNoir()
		default:
			fmt.Println("\033[H\033[2J")
			p.ForetNoir()

		}

	}
}
