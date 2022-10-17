package harvester

import (
	"log"
	"math"
	"time"

	sc "github.com/drabadan/gostealthclient"
	m "github.com/drabadan/gostealthclient/pkg/model"
)

func FindUniqueCaptchaElem(pics []m.TilePic) m.TilePic {
	var sum int
	for _, p := range pics {
		sum += int(p.Id)
	}

	average := sum / len(pics)

	difs := make([]uint32, len(pics))
	for i, p := range pics {
		difs[i] = uint32((math.Abs(float64(average) - float64(p.Id))))
	}

	var max uint32
	var idx int
	for i, e := range difs {
		if i == 0 || e > max {
			max = e
			idx = i
		}
	}

	return pics[idx]
}

func ResolveCaptchaIfPresent() {
	for {
		if <-sc.Connected() {
			c := int(<-sc.GetGumpsCount())

			if c > 0 {
				log.Println("Gumps present!")
				for i := 0; i < c; i++ {
					gi := <-sc.GetGumpInfo(uint16(i))

					if gi.GumpId != CAPTCHA_GUMP_ID {
						sc.CloseSimpleGump(uint16(i))
						continue
					}

					log.Println("Antimacro captcha present...")
					log.Printf("Gump tilepics: %v", gi.ExtInfo.TilePic)

					p := FindUniqueCaptchaElem(gi.ExtInfo.TilePic)
					for _, gb := range gi.ExtInfo.GumpButtons {
						if gb.ElemNum == p.ElemNum+1 {
							<-sc.NumGumpButton(uint16(i), gb.ReturnValue)
							log.Printf("Gumps has been resolved with gumpButton: %v", gb)
							break
						}
					}
				}
			} else {
				break
			}
		}
		time.Sleep(time.Second * 2)
	}

	log.Println("No gumps present, proceeding...")
}
