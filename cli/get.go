package cli

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/WanderningMaster/kv/config"
	"github.com/WanderningMaster/kv/internal/assert"
	"github.com/WanderningMaster/kv/internal/encryption"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get value by key",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		cfg := ctx.Value("cfg").(config.Config)

		data, err := os.ReadFile(path.Join(cfg.Path, args[0]))
		if os.IsNotExist(err) {
			fmt.Println("Invalid key")
			return nil
		}
		assert.Assert(err)

		fmt.Println("Data: ", string(data))
		fmt.Println("Key: ", cfg.Key)
		decrypted := strings.TrimSpace(
			string(encryption.Dec(string(data), cfg.Key)),
		)
		clipboard.WriteAll(decrypted)
		fmt.Println("Copied!")

		return nil
	},
}
