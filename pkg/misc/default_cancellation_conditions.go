package misc

import sc "github.com/drabadan/gostealthclient"

func DefaultCancellationConditions() bool {
	return <-sc.Connected() || <-sc.Dead()
}
