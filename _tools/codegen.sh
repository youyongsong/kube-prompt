#!/bin/bash

DIR=$(cd $(dirname $0); pwd)
KUBE_DIR=$(cd $(dirname $(dirname $0)); pwd)/kube

# clean generated files
rm ${KUBE_DIR}/*.gen.go
mkdir -p bin

set -e

go build -o ./bin/option-gen ./_tools/option-gen/main.go

subcmds=(
    "get"
    "describe"
    "create"
    "replace"
    "patch"
    "delete"
    "edit"
    "apply"
    "set env"
    "set image"
    "set resources"
    "set selector"
    "set serviceaccount"
    "set subject"
    "logs"
    "scale"
    "attach"
    "exec"
    "port-forward"
    "proxy"
    "run"
    "expose"
    "autoscale"
    "rollout history"
    "rollout pause"
    "rollout resume"
    "rollout status"
    "rollout undo"
    "label"
    "explain"
    "cordon"
    "drain"
    "uncordon"
    "annotate"
    "top node"
    "top pod"
    "cluster-info dump"
    "config get-contexts"
    "config set"
    "config set-cluster"
    "config set-credentials"
    "config view"
    "certificate approve"
    "certificate deny"
    "diff"
)
set -xv

for cmd in "${subcmds[@]}"; do
  camelized=`echo ${cmd} | gsed -r 's/[- ](.)/\U\1\E/g'`
  snaked=`echo ${cmd} | gsed -r 's/[- ]/_/g'`
  kubectl ${cmd} --help | ./bin/option-gen -o ${KUBE_DIR}/option_${snaked}.gen.go -var ${camelized}Options
  goimports -w ${KUBE_DIR}/option_${snaked}.gen.go
done
