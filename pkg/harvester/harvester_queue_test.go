package harvester_test

import (
	"testing"

	sc "github.com/drabadan/gostealthclient"
	"github.com/drabadan/uorenaissance/pkg/harvester"
)

func TestHarvesterQueue(t *testing.T) {
	t.Run("Harvester queue should execute next action", func(t *testing.T) {

		ans := sc.Bootstrap(harvester.QueueLumberjack)
		if ans != nil {
			t.Fatalf("Failed to find unique element, return value")
		}
	})
}
