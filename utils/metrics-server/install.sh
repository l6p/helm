#!/usr/bin/env bash
SRC_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
helm install metrics-server -n kube-system -f "${SRC_PATH}"/values.yaml stable/metrics-server
