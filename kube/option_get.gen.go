// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var getOptions = []prompt.Suggest{
	prompt.Suggest{Text: "-A", Description: "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace."},
	prompt.Suggest{Text: "--all-namespaces", Description: "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace."},
	prompt.Suggest{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	prompt.Suggest{Text: "--chunk-size", Description: "Return large lists in chunks rather than all at once. Pass 0 to disable. This flag is beta and may change in the future."},
	prompt.Suggest{Text: "--field-selector", Description: "Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1,key2=value2). The server only supports a limited number of field queries per type."},
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files identifying the resource to get from a server."},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files identifying the resource to get from a server."},
	prompt.Suggest{Text: "--ignore-not-found", Description: "If the requested object does not exist the command will return exit code 0."},
	prompt.Suggest{Text: "-k", Description: "Process the kustomization directory. This flag can't be used together with -f or -R."},
	prompt.Suggest{Text: "--kustomize", Description: "Process the kustomization directory. This flag can't be used together with -f or -R."},
	prompt.Suggest{Text: "-L", Description: "Accepts a comma separated list of labels that are going to be presented as columns. Names are case-sensitive. You can also use multiple flag options like -L label1 -L label2..."},
	prompt.Suggest{Text: "--label-columns", Description: "Accepts a comma separated list of labels that are going to be presented as columns. Names are case-sensitive. You can also use multiple flag options like -L label1 -L label2..."},
	prompt.Suggest{Text: "--no-headers", Description: "When using the default or custom-column output format, don't print headers (default print headers)."},
	prompt.Suggest{Text: "-o", Description: "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-as-json|jsonpath-file|custom-columns-file|custom-columns|wide See custom columns [https://kubernetes.io/docs/reference/kubectl/overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [https://kubernetes.io/docs/reference/kubectl/jsonpath/]."},
	prompt.Suggest{Text: "--output", Description: "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-as-json|jsonpath-file|custom-columns-file|custom-columns|wide See custom columns [https://kubernetes.io/docs/reference/kubectl/overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [https://kubernetes.io/docs/reference/kubectl/jsonpath/]."},
	prompt.Suggest{Text: "--output-watch-events", Description: "Output watch event objects when --watch or --watch-only is used. Existing objects are output as initial ADDED events."},
	prompt.Suggest{Text: "--raw", Description: "Raw URI to request from the server.  Uses the transport specified by the kubeconfig file."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "-l", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
	prompt.Suggest{Text: "--selector", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
	prompt.Suggest{Text: "--server-print", Description: "If true, have the server return the appropriate table output. Supports extension APIs and CRDs."},
	prompt.Suggest{Text: "--show-kind", Description: "If present, list the resource type for the requested object(s)."},
	prompt.Suggest{Text: "--show-labels", Description: "When printing, show all labels as the last column (default hide labels column)"},
	prompt.Suggest{Text: "--show-managed-fields", Description: "If true, keep the managedFields when printing objects in JSON or YAML format."},
	prompt.Suggest{Text: "--sort-by", Description: "If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string."},
	prompt.Suggest{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
	prompt.Suggest{Text: "-w", Description: "After listing/getting the requested object, watch for changes. Uninitialized objects are excluded if no object name is provided."},
	prompt.Suggest{Text: "--watch", Description: "After listing/getting the requested object, watch for changes. Uninitialized objects are excluded if no object name is provided."},
	prompt.Suggest{Text: "--watch-only", Description: "Watch for changes to the requested object(s), without listing/getting first."},
}
