package combat

import (
	"log"
	"time"

	sc "github.com/drabadan/gostealthclient"
	"github.com/drabadan/uorenaissance/pkg/misc"
)

var enemy uint32

func grabArrowsIfFound(b uint32) {
	for {
		a := <-sc.FindType(ARROW_TYPE, sc.Ground())
		if a == 0 {
			break
		}

		<-sc.MoveItem(a, 0, b, 0, 0, 0)
		time.Sleep(2 * time.Second)
	}
}

func checkEngage() {
	// log.Print("Check engage")
	if <-sc.GetHP(enemy) < 18 {
		// log.Printf("Disengage because enemy hp too low: % v", <-sc.GetHP(enemy))
		sc.SetWarMode(false)
	} else {
		sc.Attack(enemy)
	}
}

func archeryTraining(targets []uint32) interface{} {
	backpack := <-sc.Backpack()
	for {
		if !<-sc.Connected() || <-sc.Dead() {
			log.Fatal("Dead or not connected!")
		}

		<-sc.FindType(ARROW_TYPE, backpack)
		if <-sc.FindFullQuantity() == 0 {
			log.Fatal("No arrows")
		}

		checkEngage()

		if !healOnce(backpack, targets, false, checkEngage) {
			log.Fatal("Failed to healonce")
		}

		time.Sleep(1 * time.Second)
		grabArrowsIfFound(backpack)
	}
}

func Archery() interface{} {
	// init
	sc.SetFindDistance(2)
	targets := []uint32{<-sc.Self()}

	for i := 0; i < 1; i++ {
		targets = append(targets, misc.SelectTarget().ID)
		enemy = misc.SelectTarget().ID
	}

	return archeryTraining(targets)
}
