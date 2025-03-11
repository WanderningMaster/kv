package cli

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/WanderningMaster/kv/config"
	"github.com/WanderningMaster/kv/internal/assert"
	"github.com/WanderningMaster/kv/internal/encryption"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setCmd)

}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set key-value",
	Args:  cobra.MinimumNArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		cfg := ctx.Value("cfg").(config.Config)

		key, value := args[0], strings.Join(args[1:], " ")

		fmt.Printf("'%s'\n", value)

		cipherText := encryption.Enc(value, cfg.Key)
		f, err := os.Create(path.Join(cfg.Path, key))
		assert.Assert(err)

		defer f.Close()

		w := bufio.NewWriter(f)
		w.WriteString(string(cipherText))
		w.Flush()

		fmt.Println("Saved!")

		return nil
	},
}
