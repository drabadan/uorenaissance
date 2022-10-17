package harvester

import (
	"log"
	"time"

	m "github.com/drabadan/gostealthclient/pkg/model"
)

type ActionPriority int

const (
	Low ActionPriority = iota
	High
	Immediate
)

type scriptAction struct {
	Command  func()
	Priority ActionPriority
}

type ScriptAction interface {
	Execute()
}

func (a *scriptAction) Execute() {
	a.Command()
}

type scriptActionsQueue struct {
	lowPrioActions       []*scriptAction
	highPrioActions      []*scriptAction
	immediatePrioActions []*scriptAction
}

func (q *scriptActionsQueue) Next() {
	log.Println("Next action")

	var a *scriptAction

	if len(q.immediatePrioActions) > 0 {
		a = q.immediatePrioActions[len(q.immediatePrioActions)-1]
		q.immediatePrioActions = q.immediatePrioActions[:len(q.immediatePrioActions)-1]
	} else if len(q.highPrioActions) > 0 {
		a = q.highPrioActions[len(q.highPrioActions)-1]
		q.highPrioActions = q.highPrioActions[:len(q.highPrioActions)-1]
	} else if len(q.lowPrioActions) > 0 {
		a = q.lowPrioActions[len(q.lowPrioActions)-1]
		q.lowPrioActions = q.lowPrioActions[:len(q.lowPrioActions)-1]
	}

	a.Execute()
	q.Next()
}

type LumberjackingScriptConfig struct {
	Self           uint32
	SearchDistance uint16
	UnloadPoint    m.Point2D
	MaxWeight      uint16
}

type HarvestingSpot struct {
	LastHarvested    time.Time
	SinceLastHarvest time.Duration
	Point2D          m.Point2D
}

type LumberjackingScript struct {
	Config LumberjackingScriptConfig
	Spots  []HarvestingSpot
}

func QueueLumberjack() interface{} {
	queue := &scriptActionsQueue{
		lowPrioActions:       []*scriptAction{},
		highPrioActions:      []*scriptAction{},
		immediatePrioActions: []*scriptAction{},
	}
	/*ls := &LumberjackingScript{
		Config: LumberjackingScriptConfig{},
		Spots:  []HarvestingSpot{},
	}*/

	// LowPrio commands
	// TravelToSpot -> HarvestSpot -> loop
	// HighPrio commands
	// Unload
	// Immideate commands
	// ResolveCaptcha

	queue.Next()
	return nil
}

/*
type Collection interface {
	createIterator() Iterator
}

type UserCollection struct {
	users []*User
}

func (u *UserCollection) createIterator() Iterator {
	return &UserIterator{
		users: u.users,
	}
}

type Iterator interface {
	hasNext() bool
	getNext() *User
}

type UserIterator struct {
	index int
	users []*User
}

func (u *UserIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false

}
func (u *UserIterator) getNext() *User {
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}

type User struct {
	name string
	age  int
}

func main() {

	user1 := &User{
		name: "a",
		age:  30,
	}
	user2 := &User{
		name: "b",
		age:  20,
	}

	userCollection := &UserCollection{
		users: []*User{user1, user2},
	}

	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
*/
