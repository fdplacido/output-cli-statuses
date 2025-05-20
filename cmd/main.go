package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"os/exec"
	"strings"
)

//go:embed gcloud.json
var gcloudMapping []byte

//go:embed kubectl.json
var kubectlMapping []byte

func main() {

	noNamespace := flag.Bool("no-namespace", false, "Do not show the namespace")
	flag.Parse()

	projectShort := getGcloudProject()
	contextShort := getKubectlContext()

	if *noNamespace {
		fmt.Printf("g:%s|k:%s\n", projectShort, contextShort)
	} else {
		namespace := getNamespace()
		fmt.Printf("g:%s|k:%s|n:%s\n", projectShort, contextShort, namespace)
	}
}

func getGcloudProject() string {
	project := run("gcloud config get-value core/project")

	projectMap := loadMapping(gcloudMapping)

	projectShort, ok := projectMap[project]
	if !ok {
		projectShort = "??"
	}

	return projectShort
}

func getKubectlContext() string {
	context := run("kubectl config current-context")

	contextMap := loadMapping(kubectlMapping)

	contextShort, ok := contextMap[context]
	if !ok {
		contextShort = "??"
	}

	return contextShort
}

func getNamespace() string {
	namespace := run("kubectl config view --minify --output=jsonpath={..namespace}")
	if namespace == "" {
		namespace = "default"
	}

	return namespace
}

func loadMapping(configData []byte) map[string]string {
	var m map[string]string
	if err := json.Unmarshal(configData, &m); err != nil {
		return map[string]string{}
	}
	return m
}

func run(cmd string) string {
	parts := strings.Fields(cmd)
	out, err := exec.Command(parts[0], parts[1:]...).Output()
	if err != nil {
		return "?"
	}
	return strings.TrimSpace(string(out))
}
