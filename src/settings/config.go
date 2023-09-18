package settings

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

const fname = "cfg.json"

var config map[string]string

func init() {
	fp, err := os.Open(fname)
	defer fp.Close()
	if err != nil {
		log.Fatalf("Error reading '%s' file: %v", fname, err)
	}
	jbytes, _ := io.ReadAll(fp)
	json.Unmarshal(jbytes, &config)
}

func Get(key string) string {
	if val, ok := config[key]; ok {
		return val
	}
	return ""
}

func IsDebug() bool {
	return Get("mode") == "debug"
}
