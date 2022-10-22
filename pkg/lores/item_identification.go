package lores

import (
	"time"

	sc "github.com/drabadan/gostealthclient"
	"github.com/drabadan/uorenaissance/pkg/misc"
)

func ItemIdentification() interface{} {
	i := misc.SelectTarget()
	for {
		sc.WaitTargetObject(i.ID)
		sc.UseSkill("Item Identification")
		time.Sleep(time.Second * 12)
	}
}
