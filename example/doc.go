// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

//go:generate sh -c "bash $GARDENER_HACK_DIR/generate-extension.sh --name=extension-shoot-cert-service --provider-type=shoot-cert-service --component-name=extension --extension-oci-repository=europe-docker.pkg.dev/gardener-project/releases/gardener/extensions/shoot-cert-service:$(cat ../VERSION) --destination=\"$REPO_ROOT/example/extension/extension.yaml\""
//go:generate kustomize build "$REPO_ROOT/example/extension/" -o "$REPO_ROOT/example/extension.yaml"

// Package example contains generated manifests for all CRDs and other examples.
// Useful for development purposes.
package example
