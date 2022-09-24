// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var drainOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--chunk-size", Description: "Return large lists in chunks rather than all at once. Pass 0 to disable. This flag is beta and may change in the future."},
	prompt.Suggest{Text: "--delete-emptydir-data", Description: "Continue even if there are pods using emptyDir (local data that will be deleted when the node is drained)."},
	prompt.Suggest{Text: "--disable-eviction", Description: "Force drain to use delete, even if eviction is supported. This will bypass checking PodDisruptionBudgets, use with caution."},
	prompt.Suggest{Text: "--dry-run", Description: "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource."},
	prompt.Suggest{Text: "--force", Description: "Continue even if there are pods not managed by a ReplicationController, ReplicaSet, Job, DaemonSet or StatefulSet."},
	prompt.Suggest{Text: "--grace-period", Description: "Period of time in seconds given to each pod to terminate gracefully. If negative, the default value specified in the pod will be used."},
	prompt.Suggest{Text: "--ignore-daemonsets", Description: "Ignore DaemonSet-managed pods."},
	prompt.Suggest{Text: "--ignore-errors", Description: "Ignore errors occurred between drain nodes in group."},
	prompt.Suggest{Text: "--pod-selector", Description: "Label selector to filter pods on the node"},
	prompt.Suggest{Text: "-l", Description: "Selector (label query) to filter on"},
	prompt.Suggest{Text: "--selector", Description: "Selector (label query) to filter on"},
	prompt.Suggest{Text: "--skip-wait-for-delete-timeout", Description: "If pod DeletionTimestamp older than N seconds, skip waiting for the pod.  Seconds must be greater than 0 to skip."},
	prompt.Suggest{Text: "--timeout", Description: "The length of time to wait before giving up, zero means infinite"},
}
