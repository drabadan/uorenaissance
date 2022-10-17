package main

import (

	// stealth client alias
	"fmt"
	"log"
	"time"

	sc "github.com/drabadan/gostealthclient"
	m "github.com/drabadan/gostealthclient/pkg/model"
	"github.com/drabadan/uorenaissance/pkg/harvester"
)

/**
Strength:90
Dexterity:45
Intelligence:90

100.0 AnimalLore
100.0 EvalInt
100.0 Magery
100.0 MagicResist
100.0 AnimalTaming
100.0 Veterinary
100.0 Meditation
*/

/**
Ossytank Axer
GM Swords
90 Tactics
GM Anatomy
GM Veterinary
GM Animal Lore
90 Healing
90 Lumber Jacking
30 Magery
*/

/** Animals

* Info * : ID: $0011BD9D Name: a cat Type: $00C9 Color: $0904
* Info * : Quantity: 0 X: 3652 Y: 2535 Z: 0
* Info * : Tooltip:

* Info * : ID: $0011DB9A Name: a Type: $0122 Color: $0000
* Info * : Quantity: 0 X: 3655 Y: 2552 Z: 0
* Info * : Tooltip:
 */

/**
recsu ->
recdu <-
*/

func defaultCancellationConditions() bool {
	return <-sc.Connected() || <-sc.Dead()
}

// Animal Lore script, looping through takes animalId as skill target
// exits if disconnected
func AnimalLoreScript() interface{} {
	sc.SetFindDistance(20)
	animalTypes := []uint16{0xc9, 0x122}

	/*sc.ClientRequestObjectTarget()
	if <-sc.WaitForClientTargetResponse(time.Duration(30 * time.Second)) {
		p := <-sc.ClientTargetResponse()*/

	for {
		if !<-sc.Connected() || <-sc.Dead() {
			break
		}

		sc.FindTypesArrayEx(animalTypes, []uint16{0xFFFF}, []uint32{0x0}, false)
		/*list := <-sc.GetFoundList()
		if len(list) == 0 {
			log.Println("[WARNING] No Animals found - waiting")
			time.Sleep(5 * time.Second)
			continue
		}*/

		animalId := uint32(0x16718D) // p.ID

		log.Printf("[INFO] Selected animal: %x\n", animalId)

		/*if !<-sc.MoveXY(<-sc.GetX(animalId), <-sc.GetY(animalId), true, 2, true) {
			continue
		}*/

		sc.WaitTargetObject(animalId)
		sc.UseSkill("Animal Lore")
		time.Sleep(10 * time.Second)
	}
	//}

	return nil
}

/**
 COTTON BUSH
 * Info * : ID: $4235C9C0 Name:  Type: $0C54 Color: $0000
* Info * : Quantity: 1 X: 3734 Y: 2602 Z: 40
* Info * : Tooltip:

COTTON RAW
* Info * : ID: $4235CFCE Name:  Type: $0DF9 Color: $0000
* Info * : Quantity: 1 X: 3734 Y: 2602 Z: 40
* Info * : Tooltip:
*/

// Occlo cotton bushes gathering script
func OccloCottonScavengerScript() interface{} {
	sc.SetFindDistance(20)
	waypoints := []m.Point2D{
		{3732, 2586},
		{3740, 2606},
	}

	for _, point := range waypoints {
		<-sc.MoveXY(point.X, point.Y, true, 1, true)

		for {

			if !<-sc.Connected() || <-sc.Dead() {
				break
			}

			bushId := <-sc.FindTypesArrayEx([]uint16{0xc54, 0xc53, 0xc52, 0xc51}, []uint16{0xffff}, []uint32{sc.Ground()}, false)
			if bushId == 0 {
				log.Println("[INFO] Bush not found")
				break
			}

			<-sc.MoveXY(<-sc.GetX(bushId), <-sc.GetY(bushId), true, 1, true)
			time.Sleep(1 * time.Second)
			sc.UseObject(bushId)
			time.Sleep(1 * time.Second)
			for {
				cottonGround := <-sc.FindType(0xdf9, sc.Ground())
				if cottonGround == 0 {
					log.Println("[INFO] Cotton on ground not found")
					break
				}

				<-sc.MoveXY(<-sc.GetX(cottonGround), <-sc.GetY(cottonGround), true, 1, true)
				<-sc.MoveItem(cottonGround, 0, <-sc.Backpack(), 0, 0, 0)
				time.Sleep(1 * time.Second)
			}

		}

	}

	return nil
}

/**
TOOL 1
* Info * : ID: $40063994 Name:  Type: $1019 Color: $0000
* Info * : Quantity: 1 X: 3667 Y: 2597 Z: 0
* Info * : Tooltip:

TOOL 2
* Info * : ID: $40017F80 Name:  Type: $1062 Color: $0000
* Info * : Quantity: 1 X: 3669 Y: 2594 Z: 0
* Info * : Tooltip:

SPOOL OF THREAD
* Info * : ID: $42387E9C Name: spool of thread : 30 Type: $0FA0 Color: $0000
* Info * : Quantity: 30 X: 53 Y: 143 Z: 0
* Info * : Tooltip:

*/
func CraftCloth() interface{} {
	<-sc.MoveXY(3669, 2596, true, 0, true)
	sc.SetFindDistance(0x14)
	sc.OpenDoor()
	for {
		c := <-sc.FindType(0xDF9, <-sc.Backpack())
		if c == 0 {
			break
		}

		sc.WaitTargetObject(0x40063994)
		sc.UseObject(c)
		time.Sleep(500 * time.Millisecond)
	}

	for {
		c := <-sc.FindType(0xFA0, <-sc.Backpack())
		if c == 0 {
			break
		}

		sc.WaitTargetObject(0x40017F80)
		sc.UseObject(c)
		time.Sleep(500 * time.Millisecond)
	}

	return nil
}

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
		if <-sc.InJournalBetweenTimes("too far away|barely help|finish applying|fail to resurrect|cured the target|heal what little damage|You are able to resurrect", tb, time.Now()) > -1 {
			break
		}
	}
}

func autoHealing(targets []uint32, approach bool) {
	backpack := <-sc.Backpack()
	for {
		if !defaultCancellationConditions() {
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

const BONDED_FOREST_OSTARD = 0xF6686

func AutoHealing() interface{} {
	self := <-sc.Self()
	targets := []uint32{BONDED_FOREST_OSTARD, self, 0x171354, 0x172033}
	autoHealing(targets, false)
	return nil
}

func selectTarget() m.TargetInfo {
	sc.ClientPrint("Select target...")
	sc.ClientRequestObjectTarget()
	<-sc.WaitForClientTargetResponse(time.Duration(30 * time.Second))
	t := <-sc.ClientTargetResponse()
	sc.ClientPrint(fmt.Sprintf("Target selected: %x", t.ID))
	sc.AddToSystemJournal(fmt.Sprintf("Target selected: %x", t.ID))
	return t
}

func AutoHealingSelectTargets() interface{} {
	self := <-sc.Self()
	targets := []uint32{self}

	for i := 0; i < 2; i++ {
		targets = append(targets, selectTarget().ID)
	}

	log.Printf("Starting autohealing with targets: % x", targets)

	autoHealing(targets, false)
	return nil
}

func AutoKillEnemy() interface{} {
	// select pet
	sc.ClientRequestObjectTarget()
	if <-sc.WaitForClientTargetResponse(time.Duration(30 * time.Second)) {
		p := <-sc.ClientTargetResponse()

		// select enemy
		sc.ClientRequestObjectTarget()
		if <-sc.WaitForClientTargetResponse(time.Duration(30 * time.Second)) {
			e := <-sc.ClientTargetResponse()

			sc.WaitTargetObject(e.ID)
			sc.UOSay("all kill")

			log.Printf("Selected pet %x, Selected enemy: %x", p.ID, e.ID)

			// autoHealing([]uint32{<-sc.Self(), p.ID})
		}
	}
	return nil
}

func Snooping() interface{} {
	for {
		sc.UseObject(0x42C4F10F)
		time.Sleep(1500 * time.Millisecond)
	}
}

func test() interface{} {
	// sc.UOSay("hello")
	return nil
}

func main() {
	log.Println("[INFO] Starting script")
	// sc.Bootstrap(AnimalLoreScript)
	// sc.Bootstrap(OccloCottonScavengerScript)
	// sc.Bootstrap(CraftCloth)
	// sc.Bootstrap(AutoHealing)
	// sc.Bootstrap(test)
	// sc.Bootstrap(AutoKillEnemy)
	sc.Bootstrap(harvester.OccloBankLumberjacking)
	// sc.Bootstrap(Snooping)
	// sc.Bootstrap(AutoHealingSelectTargets)
	log.Println("[INFO] Script finished")
}

//despise
