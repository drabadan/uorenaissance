package harvester

import (
	"log"
	"time"

	sc "github.com/drabadan/gostealthclient"
	m "github.com/drabadan/gostealthclient/pkg/model"
)

func OccloCottonScavengerScript() interface{} {
	sc.SetFindDistance(20)
	waypoints := []m.Point2D{
		{X: 3732, Y: 2586},
		{X: 3740, Y: 2606},
	}

	for _, point := range waypoints {
		<-sc.MoveXY(point.X, point.Y, true, 1, true)

		for {
			if !<-sc.Connected() || <-sc.Dead() {
				break
			}

			bushId := <-sc.FindTypesArrayEx(COTTON_BUSH_TYPES, []uint16{0xffff}, []uint32{sc.Ground()}, false)
			if bushId == 0 {
				log.Println("[INFO] Bush not found")
				break
			}

			<-sc.MoveXY(<-sc.GetX(bushId), <-sc.GetY(bushId), true, 1, true)
			time.Sleep(1 * time.Second)
			sc.UseObject(bushId)
			time.Sleep(1 * time.Second)
			ResolveCaptchaIfPresent()
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
