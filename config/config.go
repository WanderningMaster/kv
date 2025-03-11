package config

import (
	"encoding/hex"
	"encoding/json"
	"io/fs"
	"os"
	stdpath "path"

	"github.com/WanderningMaster/kv/internal/assert"
	"github.com/WanderningMaster/kv/internal/encryption"
)

type Config struct {
	Path    string `json:"path"`
	Key     string `json:"-"`
	KeyPath string `json:"keyPath"`
}

var rwAll fs.FileMode = 0777
var rwRoot fs.FileMode = 0600

func openOrCreate(path string, readBuffer bool, perm fs.FileMode) (string, bool) {
	exists := true

	file, err := os.Open(path)
	defer file.Close()

	if os.IsNotExist(err) {
		exists = false

		file, err = os.Create(path)
		assert.Assert(err)

		err = os.Chmod(path, perm)
		assert.Assert(err)
	}

	if !readBuffer {
		return "", exists
	}

	stat, err := file.Stat()
	assert.Assert(err)

	if stat.IsDir() {
		return "", exists
	}

	filesize := stat.Size()

	buffer := make([]byte, filesize)
	_, err = file.Read(buffer)
	assert.Assert(err)

	return string(buffer), exists
}

func mkdirIfNotExists(path string, perm fs.FileMode) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, perm)
		assert.Assert(err)
	}
}

func LoadConfig() Config {
	configDir, err := os.UserConfigDir()
	assert.Assert(err)

	homeDir, err := os.UserHomeDir()
	assert.Assert(err)

	configDirPath := stdpath.Join(configDir, "kv")
	mkdirIfNotExists(configDirPath, rwAll)

	storagePath := stdpath.Join(homeDir, "kv")
	mkdirIfNotExists(storagePath, rwAll)

	var cfg Config
	configPath := stdpath.Join(configDirPath, "config.json")
	jsonstr, exists := openOrCreate(configPath, true, rwAll)

	if exists {
		err = json.Unmarshal([]byte(jsonstr), &cfg)
		assert.Assert(err)

		keyBytes, _ := openOrCreate(cfg.KeyPath, true, rwAll)
		assert.Assert(err)

		cfg.Key = keyBytes
	} else {
		cfg.Path = storagePath
		cfg.KeyPath = stdpath.Join(configDirPath, "key")

		keyBytes := encryption.GenKey()
		key := hex.EncodeToString(keyBytes)
		cfg.Key = key

		_, _ = openOrCreate(cfg.KeyPath, false, rwAll)
		err = os.WriteFile(cfg.KeyPath, []byte(key), rwAll)
		assert.Assert(err)

		jsonstr, err := json.Marshal(cfg)
		assert.Assert(err)

		err = os.WriteFile(configPath, jsonstr, rwAll)
	}

	return cfg
}
