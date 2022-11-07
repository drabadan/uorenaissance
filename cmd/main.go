package main

import (
	"log"

	"github.com/drabadan/gostealthclient"
	"github.com/drabadan/uorenaissance/pkg/crafter"
)

func main() {
	log.Println("[INFO] Starting script")
	// sc.Bootstrap(AnimalLoreScript)
	// sc.Bootstrap(OccloCottonScavengerScript)
	gostealthclient.Bootstrap(crafter.CraftCloth)
	// gostealthclient.Bootstrap(combat.AutoHealing)
	// sc.Bootstrap(test)
	// sc.Bootstrap(AutoKillEnemy)
	// gostealthclient.Bootstrap(harvester.OccloBankLumberjacking)
	// sc.Bootstrap(Snooping)
	// sc.Bootstrap(combat.AutoHealingSelectTargets)
	// sc.Bootstrap(combat.TrainMagery)
	// sc.Bootstrap(harvester.OccloBankMining)
	// sc.Bootstrap(harvester.CheckTileFromClientTarget)
	// gostealthclient.Bootstrap(lores.ItemIdentification)
	// gostealthclient.Bootstrap(harvester.MinocMining)
	// gostealthclient.Bootstrap(harvester.CoveHouseLumberjacking)
	// gostealthclient.Bootstrap(combat.Archery)

	log.Println("[INFO] Script finished")
}
