package harvester

import (
	"log"
	"time"

	sc "github.com/drabadan/gostealthclient"
	m "github.com/drabadan/gostealthclient/pkg/model"
)

func afterTree(t m.FoundTile) {
	log.Printf("Tree has finished, % v", t)
}

func Lumberjacking(waypoints []m.Point2D, unloader Unloader, afterTree func(m.FoundTile)) interface{} {
	if len(waypoints) == 0 {
		log.Fatal("No waypoints, halting...")
	}

	self := <-sc.Self()
	searchDist := uint16(SEARCH_TREES_RANGE)
	maxWeight := <-sc.MaxWeight() - 20
	sc.SetMoveOpenDoor(true)

	// we start the main loop with unloading
	unloader.unload()

	// the main loop
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
				if DEBUG {
					log.Println("Tree tiles not found, next waypoint")
				}
				time.Sleep(3 * time.Second)
				continue
			}

			for _, t := range ft {
			nextTree:
				for i := 0; i < 10; i++ {
					if <-sc.Weight() >= maxWeight-20 {
						unloader.unload()
					}

					ResolveCaptchaIfPresent()
					equipt(AXE_TYPE, sc.LhandLayer())
					<-sc.MoveXY(uint16(t.X), uint16(t.Y), true, 1, true)

					h := <-sc.ObjAtLayer(sc.LhandLayer())
					sc.WaitTargetTile(t.Tile, uint16(t.X), uint16(t.Y), byte(t.Z))
					sc.UseObject(h)
					tb := time.Now()
				retryTree:
					for i := 0; i < 100; i++ {
						time.Sleep(100 * time.Millisecond)
						if <-sc.InJournalBetweenTimes(LUMBERJACKING_MESSAGES["BREAK"], tb, time.Now()) > -1 {
							if DEBUG {
								log.Printf("Next tree due to: %v", <-sc.LastJournalMessage())
							}
							break nextTree
						}
						if <-sc.InJournalBetweenTimes(LUMBERJACKING_MESSAGES["RETRY"], tb, time.Now()) > -1 {
							if DEBUG {
								log.Printf("Retry tree due to: %v", <-sc.LastJournalMessage())
							}
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

	u := &bankUnloader{
		unloader{
			p: UNLOAD_POINT,
			c: UNLOAD_BAG_BANK,
			t: LOGS_TYPE,
		},
	}

	var ul Unloader = u
	return Lumberjacking(w, ul, afterTree)
}

func CoveHouseLumberjacking() interface{} {
	w := []m.Point2D{
		{X: 2534, Y: 1160},
		{X: 2598, Y: 1137},
	}

	u := &homeLockedSpotUnloader{
		unloader{
			p: m.Point2D{X: 2555, Y: 1181},
			// c: 0x430C3DE1,
			c: 0x43B15EEA,
			t: LOGS_TYPE,
		},
	}

	var ul Unloader = u

	return Lumberjacking(w, ul, afterTree)
}
