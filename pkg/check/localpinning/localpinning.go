package localpinning

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/ethersphere/bee/pkg/swarm"
	"github.com/ethersphere/beekeeper/pkg/bee"
	"github.com/ethersphere/beekeeper/pkg/beeclient/api"
	"github.com/ethersphere/swarm/chunk"
	"github.com/prometheus/client_golang/prometheus/push"
)

// Options represents localpinning check options
type Options struct {
	FileName       string
	LargeFileCount int
	LargeFileSize  int64
	Seed           int64
	SmallFileSize  int64
}

// randomIndexes finds n random indexes <max and but excludes skipped
func randomIndexes(rnd *rand.Rand, n, max int, skipped []int) (indexes []int, err error) {
	if n > max-len(skipped) {
		return []int{}, fmt.Errorf("not enough nodes")
	}

	found := false
	for !found {
		i := rnd.Intn(max)
		if !contains(indexes, i) && !contains(skipped, i) {
			indexes = append(indexes, i)
		}
		if len(indexes) == n {
			found = true
		}
	}

	return
}

// contains checks if a given set of int contains given int
func contains(s []int, v int) bool {
	for _, a := range s {
		if a == v {
			return true
		}
	}
	return false
}

// metricsHandler wraps pushing metrics condition
func metricsHandler(pusher *push.Pusher, yes bool) {
	if yes {
		if err := pusher.Push(); err != nil {
			fmt.Printf("push metrics: %s\n", err)
		}
	}
}

func nodeHasChunk(n bee.Node, addr swarm.Address) (bool, error) {
	var counter = 0
	for i := 0; i < retries; i++ {
		time.Sleep(1 * time.Second)
		has, err = c.Nodes[pivot].HasChunk(ctx, chunk.Address())
		if err != nil {
			if counter > 5 {
				metricsHandler(pusher, pushMetrics)
				return false, err
			}
			if errors.Is(err, api.ErrServiceUnavailable) {
				continue
			}
		}
		if has {
			return true, nil
		}
	}
	return false, nil
}
