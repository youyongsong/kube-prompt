// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var portForwardOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--address", Description: "Addresses to listen on (comma separated). Only accepts IP addresses or localhost as a value. When localhost is supplied, kubectl will try to bind on both 127.0.0.1 and ::1 and will fail if neither of these addresses are available to bind."},
	prompt.Suggest{Text: "--pod-running-timeout", Description: "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running"},
}
