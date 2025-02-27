// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var setEnvOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--all", Description: "If true, select all resources in the namespace of the specified resource types"},
	prompt.Suggest{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	prompt.Suggest{Text: "-c", Description: "The names of containers in the selected pod templates to change - may use wildcards"},
	prompt.Suggest{Text: "--containers", Description: "The names of containers in the selected pod templates to change - may use wildcards"},
	prompt.Suggest{Text: "--dry-run", Description: "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource."},
	prompt.Suggest{Text: "-e", Description: "Specify a key-value pair for an environment variable to set into each container."},
	prompt.Suggest{Text: "--env", Description: "Specify a key-value pair for an environment variable to set into each container."},
	prompt.Suggest{Text: "--field-manager", Description: "Name of the manager used to track field ownership."},
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files the resource to update the env"},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files the resource to update the env"},
	prompt.Suggest{Text: "--from", Description: "The name of a resource from which to inject environment variables"},
	prompt.Suggest{Text: "--keys", Description: "Comma-separated list of keys to import from specified resource"},
	prompt.Suggest{Text: "-k", Description: "Process the kustomization directory. This flag can't be used together with -f or -R."},
	prompt.Suggest{Text: "--kustomize", Description: "Process the kustomization directory. This flag can't be used together with -f or -R."},
	prompt.Suggest{Text: "--list", Description: "If true, display the environment and any changes in the standard format. this flag will removed when we have kubectl view env."},
	prompt.Suggest{Text: "--local", Description: "If true, set env will NOT contact api-server but run locally."},
	prompt.Suggest{Text: "-o", Description: "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-as-json|jsonpath-file."},
	prompt.Suggest{Text: "--output", Description: "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-as-json|jsonpath-file."},
	prompt.Suggest{Text: "--overwrite", Description: "If true, allow environment to be overwritten, otherwise reject updates that overwrite existing environment."},
	prompt.Suggest{Text: "--prefix", Description: "Prefix to append to variable names"},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--resolve", Description: "If true, show secret or configmap references when listing variables"},
	prompt.Suggest{Text: "-l", Description: "Selector (label query) to filter on"},
	prompt.Suggest{Text: "--selector", Description: "Selector (label query) to filter on"},
	prompt.Suggest{Text: "--show-managed-fields", Description: "If true, keep the managedFields when printing objects in JSON or YAML format."},
	prompt.Suggest{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
}
