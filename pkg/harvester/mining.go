package harvester

import (
	"fmt"
	"log"
	"time"

	sc "github.com/drabadan/gostealthclient"
	m "github.com/drabadan/gostealthclient/pkg/model"
)

func CheckTileFromClientTarget() interface{} {
	sc.ClientRequestTileTarget()
	<-sc.WaitForClientTargetResponse(time.Second * 30)
	r := <-sc.ClientTargetResponse()

	mt := <-sc.ReadStaticsXY(uint16(r.X), uint16(r.Y), <-sc.WorldNum())
	sta := <-sc.GetStaticTilesArray(uint16(r.X)+4, uint16(r.Y)+4, uint16(r.X)-4, uint16(r.Y)-4, <-sc.WorldNum(), []uint16{0xFFFF, 0x0, 1342})

	log.Println(r)
	log.Println(mt)
	log.Println(sta)
	sc.AddToSystemJournal(fmt.Sprint(mt))
	return nil
}

func smelt(smeltPoint m.Point2D, forgeId uint32) {
	b := <-sc.Backpack()
	if <-sc.MoveXY(smeltPoint.X, smeltPoint.Y, true, 1, true) {
		for _, t := range ORE_TYPES {
			for {
				<-sc.FindType(t, b)
				if <-sc.FindItem() == 0 || <-sc.GetQuantity(<-sc.FindItem()) < 2 {
					break
				}

				sc.WaitTargetObject(forgeId)
				sc.UseObject(<-sc.FindItem())
				time.Sleep(time.Second * 1)
			}
		}
	} else {
		log.Fatal("Could not reach smelt point")
	}
}

var minable_points = make([]m.Point2D, 0)

func Mining(waypoints []m.Point2D, ul Unloader, afterTree func(m.FoundTile)) interface{} {
	if len(waypoints) == 0 {
		log.Fatal("No waypoints, halting...")
	}

	self := <-sc.Self()
	searchDist := uint16(SEARCH_TREES_RANGE)
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

			ft := <-sc.GetStaticTilesArray(x-searchDist, y-searchDist, x+searchDist, y+searchDist, <-sc.WorldNum(), MINING_TILES)

			for _, t := range ft {
			nextTree:
				for r := 0; r < 10; r++ {
					ResolveCaptchaIfPresent()
					if <-sc.Weight() >= maxWeight-20 {
						smelt(SMELT_POINT_MINOC, FORGE_MINOC)

						dw := float64(<-sc.MaxWeight()) * 0.5

						log.Printf("Current dw: %v and my weight is: %v", dw, <-sc.Weight())

						if <-sc.Weight() >= uint16(dw) {
							// unloadToBank(INGOT_TYPE, unloadBag, unloadPoint)
							ul.unload()
						}
						<-sc.MoveXY(w.X, w.Y, true, 1, true)
					}

					h := <-sc.ObjAtLayer(sc.RhandLayer())

					if h == 0 {
						h = equipt(PICAXE_TYPE, sc.RhandLayer())
					}

					<-sc.MoveXY(uint16(t.X), uint16(t.Y), true, 1, true)

					sc.WaitTargetTile(t.Tile, uint16(t.X), uint16(t.Y), byte(t.Z))
					sc.UseObject(h)
					tb := time.Now()
				retryTree:
					for i := 0; i < 100; i++ {
						time.Sleep(100 * time.Millisecond)
						if <-sc.InJournalBetweenTimes(MINING_MESSAGES["BREAK"], tb, time.Now()) > -1 {
							if DEBUG {
								log.Printf("Next spot x:%v, y:%v due to: %v", t.X, t.Y, <-sc.LastJournalMessage())
							}
							time.Sleep(time.Second)
							break nextTree
						}
						if <-sc.InJournalBetweenTimes(MINING_MESSAGES["RETRY"], tb, time.Now()) > -1 {
							// minable_points = append(minable_points, m.Point2D{X: uint16(t.X), Y: uint16(t.Y)})
							if DEBUG {
								log.Printf("Retry spot due to: %v", <-sc.LastJournalMessage())
							}

							break retryTree
						}
					}
				}
			}
		}
	}
}

/*func OccloBankMining() interface{} {
	w := []m.Point2D{
		{X: 3724, Y: 2463},
		{X: 3678, Y: 2458},
		{X: 3673, Y: 2454},
		{X: 3668, Y: 2455},
		{X: 3724, Y: 2468},
		{X: 3725, Y: 2472},
	}

	return Mining(w, 0x439A1161, afterTree)
}*/

func MinocMining() interface{} {
	w := []m.Point2D{
		{X: 2566, Y: 485},
	}

	u := &bankUnloader{
		unloader{
			p: UNLOAD_POINT_MINOC_BANK,
			c: 0x439A1161,
			t: INGOT_TYPE,
		},
	}

	var ul Unloader = u
	return Mining(w, ul, afterTree)
}
