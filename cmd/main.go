package main

import (

	// stealth client alias

	"log"

	sc "github.com/drabadan/gostealthclient"
	"github.com/drabadan/uorenaissance/pkg/harvester"
)

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
