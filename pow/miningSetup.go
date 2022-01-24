package pow

import (
	"fmt"

	"github.com/zapproject/pythia/config"
)

func SetupMiningGroup(cfg *config.Config) (*MiningGroup, error) {
	var hashers []Hasher

	if len(hashers) == 0 {
		fmt.Printf("Using %d threads\n", cfg.NumProcessors)
		for i := 0; i < cfg.NumProcessors; i++ {
			hashers = append(hashers, NewCpuMiner(int64(i)))
		}

	}
	return NewMiningGroup(hashers), nil
}
