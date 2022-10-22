package harvester

import (
	"log"
	"time"

	sc "github.com/drabadan/gostealthclient"
)

func defaultCancellationConditions() bool {
	return <-sc.Connected() || <-sc.Dead()
}

func equipt(t uint16, l byte) uint32 {
	if <-sc.ObjAtLayer(l) == 0 {
		tool := <-sc.FindType(t, <-sc.Backpack())
		if tool == 0 {
			sc.AddToSystemJournal("No equip item found")
			log.Fatal("No equip item found")
		}
		<-sc.DragItem(tool, 1)
		sc.WearItem(l, tool)
		time.Sleep(time.Second * 2)
		return tool
	}

	return 0
}
