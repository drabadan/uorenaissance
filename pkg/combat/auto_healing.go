package combat

import (
	"fmt"
	"log"
	"time"

	sc "github.com/drabadan/gostealthclient"
	"github.com/drabadan/uorenaissance/pkg/misc"
)

/**
BANDAGES
* Info * : ID: $424859C3 Name: clean bandage : 220 Type: $0E21 Color: $0000
* Info * : Quantity: 197 X: 50 Y: 74 Z: 0
* Info * : Tooltip:
*/
func healTarget(t uint32, b uint32) {
	sc.WaitTargetObject(t)
	sc.UseObject(b)

	tb := time.Now()
	for i := 0; i < 100; i++ {
		time.Sleep(200 * time.Millisecond)
		if <-sc.InJournalBetweenTimes(BANDAGE_HEALING_JOURNAL_MESSAGES, tb, time.Now()) > -1 {
			break
		}
	}
}

func autoHealing(targets []uint32, approach bool) {
	backpack := <-sc.Backpack()
	for {
		if !misc.DefaultCancellationConditions() {
			log.Println("[CRITICAL] Dead or not Connected")
			break
		}

		b := <-sc.FindTypeEx(0xe21, 0x0, backpack, true)
		q := <-sc.FindFullQuantity()
		if q == 0 {
			sc.ClientPrint("[CRITICAL] Bandages not found, halting")
			log.Println("[CRITICAL] Bandages not found, halting")
			break
		}

		if q < 25 {
			sc.ClientPrint(fmt.Sprintf("[WARNING] Low on bandages: %v", q))
		} else {
			sc.ClientPrint(fmt.Sprintf("[INFO] Bandages: %v", q))
		}

		for _, t := range targets {
			if approach {
				if !<-sc.MoveXY(<-sc.GetX(t), <-sc.GetY(t), true, 1, true) {
					log.Printf("[CRITICAL] Failed to reach % x", t)
					break
				}
			}

			hp := <-sc.GetHP(t)

			if hp == <-sc.GetMaxHP(t) {
				time.Sleep(1 * time.Second)
				continue
			}

			healTarget(t, b)
			sc.AddToSystemJournal(fmt.Sprintf("Health of t: %x is %v hp; Bandages: %v", t, hp, q))
		}
	}
}

func AutoHealing() interface{} {
	self := <-sc.Self()
	targets := []uint32{BONDED_FOREST_OSTARD, self}
	autoHealing(targets, false)
	return nil
}

func AutoHealingSelectTargets() interface{} {
	self := <-sc.Self()
	targets := []uint32{self}

	for i := 0; i < 2; i++ {
		targets = append(targets, misc.SelectTarget().ID)
	}

	log.Printf("Starting autohealing with targets: % x", targets)

	autoHealing(targets, false)
	return nil
}
