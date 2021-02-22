#!/bin/sh

set -e

# Ensure cluster is up and running
timeout 600 sh -c "until k3s kubectl explain deployment &> /dev/null; do sleep 1; done"  || ( echo >&2 "Timed out while waiting for the Kubernetes cluster to start"; exit 1 )

# Wait for flyte deployment
k3s kubectl wait --for=condition=available deployment/datacatalog deployment/flyteadmin deployment/flyteconsole deployment/flytepropeller -n flyte --timeout=10m || ( echo >&2 "Timed out while waiting for the Flyte deployment to start"; exit 1 )

timeout 600 sh -c 'until [[ $(k3s kubectl get daemonset envoy -n flyte -o jsonpath="{.status.numberReady}") -eq 1 ]]; do sleep 1; done' || ( echo >&2 "Timed out while waiting for the Flyte envoy proxy to start"; exit 1 )
