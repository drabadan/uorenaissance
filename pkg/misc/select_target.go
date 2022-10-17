package misc

import (
	"fmt"
	"time"

	sc "github.com/drabadan/gostealthclient"
	m "github.com/drabadan/gostealthclient/pkg/model"
)

func SelectTarget() m.TargetInfo {
	sc.ClientPrint("Select target...")
	sc.ClientRequestObjectTarget()
	<-sc.WaitForClientTargetResponse(time.Duration(30 * time.Second))
	t := <-sc.ClientTargetResponse()
	sc.ClientPrint(fmt.Sprintf("Target selected: %x", t.ID))
	sc.AddToSystemJournal(fmt.Sprintf("Target selected: %x", t.ID))
	return t
}
