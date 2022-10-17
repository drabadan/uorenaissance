package harvester

import (
	"log"
	"time"

	sc "github.com/drabadan/gostealthclient"
	m "github.com/drabadan/gostealthclient/pkg/model"
)

func defaultCancellationConditions() bool {
	return <-sc.Connected() || <-sc.Dead()
}

func unload(t uint16, c uint32, p m.Point2D) interface{} {
	<-sc.MoveXY(p.X, p.Y, true, 1, true)
	sc.UOSay("bank")
	time.Sleep(1 * time.Second)

	for i := 0; i < 10; i++ {
		i := <-sc.FindTypeEx(t, 0xFFFF, <-sc.Backpack(), true)

		if i == 0 {
			break
		}

		<-sc.MoveItem(i, 0, c, 0, 0, 0)
		time.Sleep(2 * time.Second)
	}

	<-sc.FindTypeEx(t, 0xFFFF, c, true)
	log.Printf("Unload completed, quantity in target container: %v", <-sc.FindFullQuantity())

	return nil
}

func equipt(t uint16) {
	if <-sc.ObjAtLayer(sc.LhandLayer()) == 0 {
		tool := <-sc.FindType(t, <-sc.Backpack())
		if tool == 0 {
			sc.AddToSystemJournal("No equip item found")
			log.Fatal("No equip item found")
		}
		<-sc.DragItem(tool, 1)
		sc.WearItem(sc.LhandLayer(), tool)
	}

}

func afterTree(t m.FoundTile) {
	log.Printf("Tree has finished, % v", t)
}

func Lumberjacking(waypoints []m.Point2D, afterTree func(m.FoundTile)) interface{} {
	if len(waypoints) == 0 {
		log.Fatal("No waypoints, halting...")
	}

	self := <-sc.Self()
	searchDist := uint16(SEARCH_TREES_RANGE)
	unloadPoint := UNLOAD_POINT
	maxWeight := <-sc.MaxWeight() - 20
	sc.SetMoveOpenDoor(true)

	for {
		for _, w := range waypoints {
			if !defaultCancellationConditions() {
				log.Println("[CRITICAL] Dead or not Connected")
				break
			}

			<-sc.MoveXY(w.X, w.Y, true, 1, true)

			x := <-sc.GetX(self)
			y := <-sc.GetY(self)

			ft := <-sc.GetStaticTilesArray(x-searchDist, y-searchDist, x+searchDist, y+searchDist, <-sc.WorldNum(), TREE_TILES)

			if len(ft) == 0 {
				log.Println("Tree tiles not found, next waypoint")
				time.Sleep(3 * time.Second)
				continue
			}

			for _, t := range ft {
			nextTree:
				for i := 0; i < 10; i++ {
					if <-sc.Weight() >= maxWeight-20 {
						unload(LOGS_TYPE, UNLOAD_BAG_BANK, unloadPoint)
					}

					ResolveCaptchaIfPresent()
					equipt(AXE_TYPE)
					<-sc.MoveXY(uint16(t.X), uint16(t.Y), true, 1, true)

					h := <-sc.ObjAtLayer(sc.LhandLayer())
					sc.WaitTargetTile(t.Tile, uint16(t.X), uint16(t.Y), byte(t.Z))
					sc.UseObject(h)
					tb := time.Now()
				retryTree:
					for i := 0; i < 100; i++ {
						time.Sleep(100 * time.Millisecond)
						if <-sc.InJournalBetweenTimes(LUMBERJACKING_MESSAGES["RETRY"], tb, time.Now()) > -1 {
							log.Printf("Next tree due to: %v", <-sc.LastJournalMessage())
							break nextTree
						}
						if <-sc.InJournalBetweenTimes(LUMBERJACKING_MESSAGES["BREAK"], tb, time.Now()) > -1 {
							log.Printf("Retry tree due to: %v", <-sc.LastJournalMessage())
							break retryTree
						}
					}
				}

				afterTree(t)
			}

		}

	}
}

func OccloBankLumberjacking() interface{} {
	w := []m.Point2D{
		{X: 3558, Y: 2564},
		{X: 3623, Y: 2433},
		{X: 3693, Y: 2406},
		{X: 3741, Y: 2522},
	}

	return Lumberjacking(w, afterTree)
}
