package asciiartweb

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type gconfigs struct {
	HostURL string `json:"HostURL"`
}

func (this *gconfigs) IsValid() bool {
	if this.HostURL == "" {
		return false
	}
	return true
}

// PathGConfigs - path to GlobalConfigs json
const PathGConfigs = "app/configs.json"

// SetGlobalConfigs - Set config depending on the startup environment (JSON | ENV)
func SetGlobalConfigs() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		SetJsonConfigs()
	} else {
		log.Printf("Set Port from ENV to gConfigs")
		gConfigs.HostURL = fmt.Sprintf("localhost:%v", port)
	}
}

// SetJsonConfigs -----------------------------
func SetJsonConfigs() {
	if gConfigs == nil {
		gConfigs = &gconfigs{}
	}
	file, err := os.Open(PathGConfigs)
	if err != nil {
		if err == os.ErrNotExist {
			CreateJsonConfigs()
			return
		}
	}
	defer file.Close()
	tconfigs := &gconfigs{}
	dec := json.NewDecoder(file)
	if err := dec.Decode(tconfigs); err != nil {
		log.Printf("Incorrect json file")
		CreateJsonConfigs()
		return
	}
	if !tconfigs.IsValid() {
		log.Printf("Incorrect json param names in file")
		CreateJsonConfigs()
		return
	}
	gConfigs = tconfigs
	log.Print("Setted Json Configs for GConfigs")
}

// CreateJsonConfigs - Just Create json file with PathGConfigs
func CreateJsonConfigs() {
	f, err := os.Create(PathGConfigs)
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
	defer f.Close()

	jsonData, err := json.Marshal(gConfigs)
	f.Write(jsonData)
	log.Printf("Created '%v' for GlobalConfigs", PathGConfigs)
}
