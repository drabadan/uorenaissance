package combat

import (
	"log"
	"time"

	sc "github.com/drabadan/gostealthclient"
)

type Spell struct {
	Name      string
	ManaCount int
	Regs      []uint16
}

func TrainMagery() interface{} {
	self := <-sc.Self()
	backpack := <-sc.Backpack()
	/*lightning := &Spell{
		Name:      "Lightning",
		ManaCount: 11,
		Regs: []uint16{
			sc.SA(),
			sc.MR(),
		},
	}*/

	energyBolt := &Spell{
		Name:      "Energy Bolt",
		ManaCount: 20,
		Regs: []uint16{
			sc.NS(),
			sc.BP(),
		},
	}

	for {
		if !<-sc.Connected() || <-sc.Dead() {
			log.Fatal("Dead or not connected")
		}

		for _, r := range energyBolt.Regs {
			<-sc.FindType(r, backpack)
			fq := <-sc.FindFullQuantity()
			if fq == 0 {
				log.Fatal("Not enough regs")
			}

			log.Printf("Regs %v count: %v", r, fq)
		}

		if <-sc.Mana() < uint32(energyBolt.ManaCount) {
			for {
				if <-sc.Mana() == <-sc.MaxMana() {
					break
				}
				sc.UseSkill("Meditation")
				time.Sleep(time.Second * 9)
			}
		}

		sc.CastToObj(energyBolt.Name, self)
		time.Sleep(time.Second * 3)

		if <-sc.GetSkillValue("Magery") >= 85 {
			break
		}
	}

	return nil
}
