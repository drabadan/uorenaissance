package crafter

import (
	"time"

	sc "github.com/drabadan/gostealthclient"
)

// TODO: refactor this for clientRequestObjectTarget
func CraftCloth() interface{} {
	for {
		c := <-sc.FindType(0xDF9, <-sc.Backpack())
		if c == 0 {
			break
		}

		sc.WaitTargetObject(0x40063994)
		sc.UseObject(c)
		time.Sleep(500 * time.Millisecond)
	}

	for {
		c := <-sc.FindType(0xFA0, <-sc.Backpack())
		if c == 0 {
			break
		}

		sc.WaitTargetObject(0x40017F80)
		sc.UseObject(c)
		time.Sleep(500 * time.Millisecond)
	}

	return nil
}
