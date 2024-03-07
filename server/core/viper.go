package core

import (
	"flag"
	"fmt"
	"github.com/WaynerEP/restaurant-app/server/core/internal"
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func Viper(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		// Check if command-line argument is empty
		if config == "" {
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" { // Check if the environment variable stored in internal.ConfigEnv is empty
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
					fmt.Printf("You are using the %s environment name in gin mode, and the config path is %s\n", gin.EnvGinMode, internal.ConfigDefaultFile)
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
					fmt.Printf("You are using the %s environment name in gin mode, and the config path is %s\n", gin.EnvGinMode, internal.ConfigReleaseFile)
				case gin.TestMode:
					config = internal.ConfigTestFile
					fmt.Printf("You are using the %s environment name in gin mode, and the config path is %s\n", gin.EnvGinMode, internal.ConfigTestFile)
				}
			} else { // The environment variable stored in internal.ConfigEnv is not empty, assign its value to config
				config = configEnv
				fmt.Printf("You are using the %s environment variable, and the config path is %s\n", internal.ConfigEnv, config)
			}
		} else { // Command-line argument is not empty, assign its value to config
			fmt.Printf("You are using the value passed by the -c parameter in the command line, and the config path is %s\n", config)
		}
	} else { // Assign the first value from the variadic parameter passed to the function to config
		config = path[0]
		fmt.Printf("You are using the value passed by the func Viper() function, and the config path is %s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}

	// Adapt root location to ensure the corresponding migration position is found, ensuring root path validity
	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return v
}
