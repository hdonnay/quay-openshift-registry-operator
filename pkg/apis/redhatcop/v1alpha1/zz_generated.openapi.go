// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/redhat-cop/quay-openshift-registry-operator/pkg/apis/redhatcop/v1alpha1.QuayIntegration":       schema_pkg_apis_redhatcop_v1alpha1_QuayIntegration(ref),
		"github.com/redhat-cop/quay-openshift-registry-operator/pkg/apis/redhatcop/v1alpha1.QuayIntegrationSpec":   schema_pkg_apis_redhatcop_v1alpha1_QuayIntegrationSpec(ref),
		"github.com/redhat-cop/quay-openshift-registry-operator/pkg/apis/redhatcop/v1alpha1.QuayIntegrationStatus": schema_pkg_apis_redhatcop_v1alpha1_QuayIntegrationStatus(ref),
	}
}

func schema_pkg_apis_redhatcop_v1alpha1_QuayIntegration(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "QuayIntegration is the Schema for the quayintegrations API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/redhat-cop/quay-openshift-registry-operator/pkg/apis/redhatcop/v1alpha1.QuayIntegrationSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/redhat-cop/quay-openshift-registry-operator/pkg/apis/redhatcop/v1alpha1.QuayIntegrationStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/redhat-cop/quay-openshift-registry-operator/pkg/apis/redhatcop/v1alpha1.QuayIntegrationSpec", "github.com/redhat-cop/quay-openshift-registry-operator/pkg/apis/redhatcop/v1alpha1.QuayIntegrationStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_redhatcop_v1alpha1_QuayIntegrationSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "QuayIntegrationSpec defines the desired state of QuayIntegration",
				Properties: map[string]spec.Schema{
					"clusterID": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"credentialsSecretName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"organizationPrefix": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"quayHostname": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"insecureRegistry": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"scheduledImageStreamImport": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"blacklistNamespaces": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"whitelistNamespaces": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"clusterID", "credentialsSecretName", "quayHostname"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_redhatcop_v1alpha1_QuayIntegrationStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "QuayIntegrationStatus defines the observed state of QuayIntegration",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
