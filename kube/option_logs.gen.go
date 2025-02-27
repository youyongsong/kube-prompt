// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var logsOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--all-containers", Description: "Get all containers' logs in the pod(s)."},
	prompt.Suggest{Text: "-c", Description: "Print the logs of this container"},
	prompt.Suggest{Text: "--container", Description: "Print the logs of this container"},
	prompt.Suggest{Text: "-f", Description: "Specify if the logs should be streamed."},
	prompt.Suggest{Text: "--follow", Description: "Specify if the logs should be streamed."},
	prompt.Suggest{Text: "--ignore-errors", Description: "If watching / following pod logs, allow for any errors that occur to be non-fatal"},
	prompt.Suggest{Text: "--insecure-skip-tls-verify-backend", Description: "Skip verifying the identity of the kubelet that logs are requested from.  In theory, an attacker could provide invalid log content back. You might want to use this if your kubelet serving certificates have expired."},
	prompt.Suggest{Text: "--limit-bytes", Description: "Maximum bytes of logs to return. Defaults to no limit."},
	prompt.Suggest{Text: "--max-log-requests", Description: "Specify maximum number of concurrent logs to follow when using by a selector. Defaults to 5."},
	prompt.Suggest{Text: "--pod-running-timeout", Description: "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running"},
	prompt.Suggest{Text: "--prefix", Description: "Prefix each log line with the log source (pod name and container name)"},
	prompt.Suggest{Text: "-p", Description: "If true, print the logs for the previous instance of the container in a pod if it exists."},
	prompt.Suggest{Text: "--previous", Description: "If true, print the logs for the previous instance of the container in a pod if it exists."},
	prompt.Suggest{Text: "-l", Description: "Selector (label query) to filter on."},
	prompt.Suggest{Text: "--selector", Description: "Selector (label query) to filter on."},
	prompt.Suggest{Text: "--since", Description: "Only return logs newer than a relative duration like 5s, 2m, or 3h. Defaults to all logs. Only one of since-time / since may be used."},
	prompt.Suggest{Text: "--since-time", Description: "Only return logs after a specific date (RFC3339). Defaults to all logs. Only one of since-time / since may be used."},
	prompt.Suggest{Text: "--tail", Description: "Lines of recent log file to display. Defaults to -1 with no selector, showing all log lines otherwise 10, if a selector is provided."},
	prompt.Suggest{Text: "--timestamps", Description: "Include timestamps on each line in the log output"},
}
