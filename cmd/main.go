package main

import (
	"log"

	"github.com/drabadan/gostealthclient"
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
	// sc.Bootstrap(harvester.OccloBankLumberjacking)
	// sc.Bootstrap(Snooping)
	// sc.Bootstrap(combat.AutoHealingSelectTargets)
	// sc.Bootstrap(combat.TrainMagery)
	// sc.Bootstrap(harvester.OccloBankMining)
	// sc.Bootstrap(harvester.CheckTileFromClientTarget)
	// sc.Bootstrap(lores.ItemIdentification)
	// gostealthclient.Bootstrap(harvester.MinocMining)
	gostealthclient.Bootstrap(harvester.CoveHouseLumberjacking)

	log.Println("[INFO] Script finished")
}
