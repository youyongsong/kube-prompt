// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var deleteOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--all", Description: "Delete all resources, including uninitialized ones, in the namespace of the specified resource types."},
	prompt.Suggest{Text: "-A", Description: "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace."},
	prompt.Suggest{Text: "--all-namespaces", Description: "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace."},
	prompt.Suggest{Text: "--cascade", Description: "Must be \"background\", \"orphan\", or \"foreground\". Selects the deletion cascading strategy for the dependents (e.g. Pods created by a ReplicationController). Defaults to background."},
	prompt.Suggest{Text: "--dry-run", Description: "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource."},
	prompt.Suggest{Text: "--field-selector", Description: "Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1,key2=value2). The server only supports a limited number of field queries per type."},
	prompt.Suggest{Text: "-f", Description: "containing the resource to delete."},
	prompt.Suggest{Text: "--filename", Description: "containing the resource to delete."},
	prompt.Suggest{Text: "--force", Description: "If true, immediately remove resources from API and bypass graceful deletion. Note that immediate deletion of some resources may result in inconsistency or data loss and requires confirmation."},
	prompt.Suggest{Text: "--grace-period", Description: "Period of time in seconds given to the resource to terminate gracefully. Ignored if negative. Set to 1 for immediate shutdown. Can only be set to 0 when --force is true (force deletion)."},
	prompt.Suggest{Text: "--ignore-not-found", Description: "Treat \"resource not found\" as a successful delete. Defaults to \"true\" when --all is specified."},
	prompt.Suggest{Text: "-k", Description: "Process a kustomization directory. This flag can't be used together with -f or -R."},
	prompt.Suggest{Text: "--kustomize", Description: "Process a kustomization directory. This flag can't be used together with -f or -R."},
	prompt.Suggest{Text: "--now", Description: "If true, resources are signaled for immediate shutdown (same as --grace-period=1)."},
	prompt.Suggest{Text: "-o", Description: "Output mode. Use \"-o name\" for shorter output (resource/name)."},
	prompt.Suggest{Text: "--output", Description: "Output mode. Use \"-o name\" for shorter output (resource/name)."},
	prompt.Suggest{Text: "--raw", Description: "Raw URI to DELETE to the server.  Uses the transport specified by the kubeconfig file."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "-l", Description: "Selector (label query) to filter on, not including uninitialized ones."},
	prompt.Suggest{Text: "--selector", Description: "Selector (label query) to filter on, not including uninitialized ones."},
	prompt.Suggest{Text: "--timeout", Description: "The length of time to wait before giving up on a delete, zero means determine a timeout from the size of the object"},
	prompt.Suggest{Text: "--wait", Description: "If true, wait for resources to be gone before returning. This waits for finalizers."},
}
