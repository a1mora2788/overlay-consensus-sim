package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type EpochResult struct {
	Epoch             int     `json:"epoch"`
	NumNodes          int     `json:"num_nodes"`
	NumActive         int     `json:"num_active"`
	OverlayEnergy     float64 `json:"overlay_energy"`
	PowEnergy         float64 `json:"pow_energy"`
	BlocksProduced    int     `json:"blocks_produced"`
	CumOverlayEnergy  float64 `json:"cum_overlay_energy"`
	CumPowEnergy      float64 `json:"cum_pow_energy"`
}

type Summary struct {
	NodesStart         int     `json:"nodes_start"`
	Epochs             int     `json:"epochs"`
	BlocksProduced     int     `json:"blocks_produced"`
	CumOverlayEnergy   float64 `json:"cum_overlay_energy"`
	CumPowEnergy       float64 `json:"cum_pow_energy"`
	OverlayVsPowRatio  float64 `json:"overlay_vs_pow_ratio"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	nodes := 100
	epochs := 200
	powPerBlock := 1000.0        // baseline PoW energy cost
	overlayPerBlock := 200.0     // overlay energy baseline

	results := []EpochResult{}
	totalOverlay := 0.0
	totalPow := 0.0
	totalBlocks := 0

	// prepare CSV writer
	csvFile, err := os.Create("per_epoch_results.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// write CSV header
	writer.Write([]string{"epoch", "num_nodes", "num_active", "overlay_energy", "pow_energy", "blocks_produced", "cum_overlay_energy", "cum_pow_energy"})

	for e := 1; e <= epochs; e++ {
		numActive := rand.Intn(nodes) + 1 // at least 1 active node
		blocks := rand.Intn(2)            // 0 or 1 block produced this epoch

		overlayEnergy := float64(numActive) * overlayPerBlock * float64(blocks)
		powEnergy := powPerBlock * float64(blocks)

		totalOverlay += overlayEnergy
		totalPow += powEnergy
		totalBlocks += blocks

		result := EpochResult{
			Epoch:            e,
			NumNodes:         nodes,
			NumActive:        numActive,
			OverlayEnergy:    overlayEnergy,
			PowEnergy:        powEnergy,
			BlocksProduced:   blocks,
			CumOverlayEnergy: totalOverlay,
			CumPowEnergy:     totalPow,
		}
		results = append(results, result)

		// write row to CSV
		writer.Write([]string{
			fmt.Sprintf("%d", result.Epoch),
			fmt.Sprintf("%d", result.NumNodes),
			fmt.Sprintf("%d", result.NumActive),
			fmt.Sprintf("%.3f", result.OverlayEnergy),
			fmt.Sprintf("%.3f", result.PowEnergy),
			fmt.Sprintf("%d", result.BlocksProduced),
			fmt.Sprintf("%.3f", result.CumOverlayEnergy),
			fmt.Sprintf("%.3f", result.CumPowEnergy),
		})
	}
	writer.Flush()

	summary := Summary{
		NodesStart:        nodes,
		Epochs:            epochs,
		BlocksProduced:    totalBlocks,
		CumOverlayEnergy:  totalOverlay,
		CumPowEnergy:      totalPow,
		OverlayVsPowRatio: totalOverlay / totalPow,
	}

	jsonData, _ := json.MarshalIndent(summary, "", "  ")
	fmt.Println(string(jsonData))
}
