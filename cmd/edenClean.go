package cmd

import (
	"fmt"
	"github.com/lf-edge/eden/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "clean harness",
	Long:  `Clean harness.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		viperLoaded, err := utils.LoadConfigFile(config)
		if err != nil {
			return fmt.Errorf("error reading config: %s", err.Error())
		}
		if viperLoaded {
			eserverPidFile = viper.GetString("eserver-pid")
			evePidFile = viper.GetString("eve-pid")
			eveDist = viper.GetString("eve-dist")
			eveBaseDist = viper.GetString("eve-base-dist")
			adamDist = viper.GetString("adam-dist")
			certsDir = viper.GetString("certs-dist")
			eserverImageDist = viper.GetString("image-dist")
			qemuFileToSave = viper.GetString("eve-config")
			binDir = viper.GetString("bin-dist")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		command, err := os.Executable()
		if err != nil {
			log.Fatalf("cannot obtain executable path: %s", err)
		}
		if err := utils.CleanEden(command, eveDist, eveBaseDist, adamDist, certsDir, eserverImageDist,
			binDir, qemuFileToSave, eserverPidFile, evePidFile); err != nil {
			log.Fatalf("cannot CleanEden: %s", err)
		}
		log.Infof("CleanEden done")
	},
}

func cleanInit() {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	cleanCmd.Flags().StringVarP(&eserverPidFile, "eserver-pid", "", filepath.Join(currentPath, "dist", "eserver.pid"), "file with eserver pid")
	cleanCmd.Flags().StringVarP(&evePidFile, "eve-pid", "", filepath.Join(currentPath, "dist", "eve.pid"), "file with EVE pid")
	cleanCmd.Flags().StringVarP(&eveDist, "eve-dist", "", filepath.Join(currentPath, "dist", "eve"), "directory to save EVE")
	cleanCmd.Flags().StringVarP(&eveBaseDist, "eve-base-dist", "", filepath.Join(currentPath, "dist", "evebaseos"), "directory to save Base image of EVE")
	cleanCmd.Flags().StringVarP(&qemuFileToSave, "eve-config", "", filepath.Join(currentPath, "dist", defaultQemuFileToSave), "file to save qemu config")

	cleanCmd.Flags().StringVarP(&adamDist, "adam-dist", "", filepath.Join(currentPath, "dist", "adam"), "adam dist to start (required)")
	cleanCmd.Flags().StringVarP(&eserverImageDist, "image-dist", "", filepath.Join(currentPath, "dist", "images"), "image dist for eserver")

	cleanCmd.Flags().StringVarP(&certsDir, "certs-dist", "o", filepath.Join(currentPath, "dist", "certs"), "directory with certs")
	cleanCmd.Flags().StringVarP(&binDir, "bin-dist", "", filepath.Join(currentPath, "dist", "bin"), "directory for binaries")
	if err := viper.BindPFlags(stopCmd.Flags()); err != nil {
		log.Fatal(err)
	}
	cleanCmd.Flags().StringVar(&config, "config", "", "path to config file")
}