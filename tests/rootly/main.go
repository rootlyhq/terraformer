package main

import (
	"log"
	"os"
	"os/exec"
	"sort"

	"github.com/GoogleCloudPlatform/terraformer/cmd"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"

	rootly_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/rootly"
)

const command = "terraform init && terraform plan"

func main() {
	organization := ""
	token := os.Getenv("ROOTLY_API_TOKEN")
	services := []string{}
	provider := &rootly_terraforming.RootlyProvider{}
	for service := range provider.GetSupportedService() {
		services = append(services, service)
	}
	sort.Strings(services)
	provider = &rootly_terraforming.RootlyProvider{
		Provider: terraformutils.Provider{},
	}
	err := cmd.Import(provider, cmd.ImportOptions{
		Resources:   services,
		PathPattern: cmd.DefaultPathPattern,
		PathOutput:  cmd.DefaultPathOutput,
		State:       "local",
		Connect:     true,
	}, []string{organization, token})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	rootPath, _ := os.Getwd()
	for _, serviceName := range services {
		currentPath := cmd.Path(cmd.DefaultPathPattern, provider.GetName(), serviceName, cmd.DefaultPathOutput)
		if err := os.Chdir(currentPath); err != nil {
			log.Println(err)
			os.Exit(1)
		}
		cmd := exec.Command("sh", "-c", command)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		err := os.Chdir(rootPath)
		if err != nil {
			log.Println(err)
		}
	}
}
