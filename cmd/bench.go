package cmd

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/ashupednekar/hdata-encoder/pkg"
	"github.com/spf13/cobra"
)

var (
	nItems int
	maxStr int
	iter   int
)

type BenchResult struct {
	Index    int
	EncodeMs int64
	SizeMB   float64
	DecodeMs int64
	Err      error
}

var benchCmd = &cobra.Command{
	Use:   "bench",
	Short: "Run encode/decode benchmarks for HData",
	Run: func(cmd *cobra.Command, args []string) {
		serde := pkg.HDataSerde{}

		results := make([]BenchResult, iter)
		var wg sync.WaitGroup
		wg.Add(iter)

		for i := 0; i < iter; i++ {
			i := i
			go func() {
				defer wg.Done()
				fmt.Printf("Generating data (n=%d, maxStr=%d)...\n", nItems, maxStr)
				data := pkg.RandomData(nItems, maxStr)
				for len(data) == 0 {
					data = pkg.RandomData(nItems, maxStr)
				}
				results[i] = RunCmdConcurrent(i, serde, data)
			}()
		}

		wg.Wait()

		// sort by index (optional but safe)
		sort.Slice(results, func(a, b int) bool {
			return results[a].Index < results[b].Index
		})

		for _, r := range results {
			if r.Err != nil {
				fmt.Printf("[%d] ERROR: %v\n", r.Index, r.Err)
				continue
			}
			fmt.Printf("[%d] encode=%dms size=%.2fMB decode=%dms\n",
				r.Index, r.EncodeMs, r.SizeMB, r.DecodeMs)
		}
	},
}

func RunCmdConcurrent(idx int, serde pkg.HDataSerde, data pkg.DataInput) BenchResult {
	startEncode := time.Now()
	encoded, err := serde.Encode(data)
	if err != nil {
		return BenchResult{Index: idx, Err: err}
	}
	encodeMs := time.Since(startEncode).Milliseconds()

	sizeMB := float64(len(encoded)) / (1024 * 1024)

	startDecode := time.Now()
	_, err = serde.Decode(encoded)
	if err != nil {
		return BenchResult{Index: idx, Err: err}
	}
	decodeMs := time.Since(startDecode).Milliseconds()

	return BenchResult{
		Index:    idx,
		EncodeMs: encodeMs,
		SizeMB:   sizeMB,
		DecodeMs: decodeMs,
	}
}

func init() {
	rootCmd.AddCommand(benchCmd)
	benchCmd.Flags().IntVarP(&nItems,
		"numItems", "n",
		4000,
		"Number of items to generate",
	)

	benchCmd.Flags().IntVarP(&maxStr,
		"maxStr", "s",
		500,
		"Maximum string size",
	)

	benchCmd.Flags().IntVarP(&iter,
		"numIters", "i",
		5,
		"Number of iterations",
	)
}
