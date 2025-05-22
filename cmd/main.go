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

var (
	itemSepStr    = "|"
	mappingSepStr = ":"
	gcloudStr     = "g"
	kubectlStr    = "k"
	namespaceStr  = "n"
)

func main() {

	noNamespace := flag.Bool("no-namespace", false, "Do not show the namespace")
	noCompact := flag.Bool("no-compact", false, "Do not show the compacted version")
	flag.Parse()

	gProject := getGcloudProjectStr()
	kubeCtx := getKubectlContextStr()

	itemSep := getItemSeparator(*noCompact)

	if *noNamespace {
		fmt.Printf("%s%s%s\n", gProject, itemSep, kubeCtx)
	} else {
		namespace := getKubectlNamespaceContextStr()
		fmt.Printf("%s%s%s%s%s\n", gProject, itemSep, kubeCtx, itemSep, namespace)
	}
}

func getItemSeparator(isCompact bool) string {
	space := ""
	if !isCompact {
		space = " "
	}
	return fmt.Sprintf("%s%s%s", space, itemSepStr, space)
}

func getGcloudProjectStr() string {
	project := getGcloudProject()
	return fmt.Sprintf("%s%s%s", gcloudStr, mappingSepStr, project)
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

func getKubectlContextStr() string {
	kubectlCtx := getKubectlContext()
	return fmt.Sprintf("%s%s%s", kubectlStr, mappingSepStr, kubectlCtx)
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

func getKubectlNamespaceContextStr() string {
	kubectlNamespace := getKubectlNamespaceContext()
	return fmt.Sprintf("%s%s%s", namespaceStr, mappingSepStr, kubectlNamespace)
}

func getKubectlNamespaceContext() string {
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
