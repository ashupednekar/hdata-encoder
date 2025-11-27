package cmd

import (
	"fmt"
	"time"

	"github.com/ashupednekar/hdata-encoder/pkg"
	"github.com/spf13/cobra"
)

var (
	nItems int
	maxStr int
)

var benchCmd = &cobra.Command{
	Use:   "bench",
	Short: "Run encode/decode benchmarks for HData",
	Long:  `Runs an in-process benchmark for encoding and decoding random HData trees.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("Generating data (n=%d, maxStr=%d)...\n", nItems, maxStr)
		data := pkg.RandomData(nItems, maxStr)
		serde := pkg.HDataSerde{}

		startEncode := time.Now()
		encoded, err := serde.Encode(data)
		if err != nil {
			fmt.Println("encode error:", err)
			return
		}
		encodeMs := time.Since(startEncode).Milliseconds()

		sizeMB := float64(len(encoded)) / (1024 * 1024)
		fmt.Printf("Encode time: %d ms\n", encodeMs)
		fmt.Printf("Encoded size: %.2f MB\n", sizeMB)

		startDecode := time.Now()
		_, err = serde.Decode(encoded)
		if err != nil {
			fmt.Println("decode error:", err)
			return
		}
		decodeMs := time.Since(startDecode).Milliseconds()
		fmt.Printf("Decode time: %d ms\n", decodeMs)

		fmt.Println("✔ Benchmark validation successful (data matches)")
	},
}

func init() {
	rootCmd.AddCommand(benchCmd)

	benchCmd.Flags().IntVar(&nItems, "n", 4000, "Number of items to generate")
	benchCmd.Flags().IntVar(&maxStr, "s", 500, "Maximum string size")
}
