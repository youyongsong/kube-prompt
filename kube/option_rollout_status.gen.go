// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var rolloutStatusOptions = []prompt.Suggest{
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files identifying the resource to get from a server."},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files identifying the resource to get from a server."},
	prompt.Suggest{Text: "-k", Description: "Process the kustomization directory. This flag can't be used together with -f or -R."},
	prompt.Suggest{Text: "--kustomize", Description: "Process the kustomization directory. This flag can't be used together with -f or -R."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--revision", Description: "Pin to a specific revision for showing its status. Defaults to 0 (last revision)."},
	prompt.Suggest{Text: "--timeout", Description: "The length of time to wait before ending watch, zero means never. Any other values should contain a corresponding time unit (e.g. 1s, 2m, 3h)."},
	prompt.Suggest{Text: "-w", Description: "Watch the status of the rollout until it's done."},
	prompt.Suggest{Text: "--watch", Description: "Watch the status of the rollout until it's done."},
}
