package harvester

import (
	"log"
	"time"

	sc "github.com/drabadan/gostealthclient"
	m "github.com/drabadan/gostealthclient/pkg/model"
)

type Unloader interface {
	unload()
}

type unloader struct {
	p m.Point2D
	c uint32
	t uint16
}

type bankUnloader struct {
	unloader
}

type chestUnloader struct {
	unloader
}

type homeLockedSpotUnloader struct {
	unloader
}

func (u chestUnloader) unload() {
	log.Println("CHEST UNLOADING STARTED")

	// we will halt if the waight after unload has not changed
	w := <-sc.Weight()
	<-sc.MoveXY(u.p.X, u.p.Y, true, 1, true)
	sc.UseObject(u.c)
	time.Sleep(1 * time.Second)

	for i := 0; i < 10; i++ {
		item := <-sc.FindTypeEx(u.t, 0xFFFF, <-sc.Backpack(), true)

		if item == 0 {
			break
		}

		<-sc.MoveItem(item, 0, u.c, 0, 0, 0)
		time.Sleep(2 * time.Second)
		if <-sc.Weight() >= w {
			sc.AddToSystemJournal("[CRITICAL] Weight has not changed after moving item, last journal message:")
			sc.AddToSystemJournal(<-sc.LastJournalMessage())
			log.Fatal("[CRITICAL] Weight has not changed after moving item")
		}
	}

	<-sc.FindTypeEx(u.t, 0xFFFF, u.c, true)
	log.Printf("Unload completed, quantity in target container: %v", <-sc.FindFullQuantity())
}

func (u bankUnloader) unload() {
	log.Println("BANK UNLOADING STARTED")
	<-sc.MoveXY(u.p.X, u.p.Y, true, 1, true)
	sc.UOSay("bank")
	time.Sleep(1 * time.Second)

	for i := 0; i < 10; i++ {
		item := <-sc.FindTypeEx(u.t, 0xFFFF, <-sc.Backpack(), true)

		if item == 0 {
			break
		}

		<-sc.MoveItem(item, 0, u.c, 0, 0, 0)
		time.Sleep(2 * time.Second)
	}

	<-sc.FindTypeEx(u.t, 0xFFFF, u.c, true)
	log.Printf("Unload completed, quantity in target container: %v", <-sc.FindFullQuantity())
}

func (u homeLockedSpotUnloader) unload() {
	log.Println("HOME SPOT UNLOADING STARTED")

	// we will halt if the waight after unload has not changed
	w := <-sc.Weight()
	<-sc.MoveXY(u.p.X, u.p.Y, true, 1, true)
	time.Sleep(1 * time.Second)

	for i := 0; i < 10; i++ {
		item := <-sc.FindTypeEx(u.t, 0xFFFF, <-sc.Backpack(), true)

		if item == 0 {
			break
		}

		color := <-sc.GetColor(item)
		container := <-sc.FindTypeEx(u.t, color, sc.Ground(), false)

		<-sc.MoveItem(item, 0, container, 0, 0, 0)
		time.Sleep(2 * time.Second)
		if <-sc.Weight() >= w {
			sc.AddToSystemJournal("[CRITICAL] Weight has not changed after moving item, last journal message:")
			sc.AddToSystemJournal(<-sc.LastJournalMessage())
			log.Fatal("[CRITICAL] Weight has not changed after moving item")
		}
	}

	<-sc.FindTypeEx(u.t, 0xFFFF, u.c, true)
	log.Printf("Unload completed, quantity in target container: %v", <-sc.FindFullQuantity())
}
