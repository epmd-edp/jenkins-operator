// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	"github.com/go-openapi/spec"
	"k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.Jenkins":                     schema_pkg_apis_v2_v1alpha1_Jenkins(ref),
		"github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsScript":               schema_pkg_apis_v2_v1alpha1_JenkinsScript(ref),
		"github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsScriptStatus":         schema_pkg_apis_v2_v1alpha1_JenkinsScriptStatus(ref),
		"github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsServiceAccount":       schema_pkg_apis_v2_v1alpha1_JenkinsServiceAccount(ref),
		"github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsServiceAccountSpec":   schema_pkg_apis_v2_v1alpha1_JenkinsServiceAccountSpec(ref),
		"github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsServiceAccountStatus": schema_pkg_apis_v2_v1alpha1_JenkinsServiceAccountStatus(ref),
		"github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsStatus":               schema_pkg_apis_v2_v1alpha1_JenkinsStatus(ref),
		"github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsFolder":               schema_pkg_apis_v2_v1alpha1_JenkinsFolder(ref),
		"github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsFolderStatus":         schema_pkg_apis_v2_v1alpha1_JenkinsFolderStatus(ref),
		"github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsJob":                  schema_pkg_apis_v2_v1alpha1_JenkinsJob(ref),
		"github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsJobSpec":              schema_pkg_apis_v2_v1alpha1_JenkinsJobSpec(ref),
		"github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsJobStatus":            schema_pkg_apis_v2_v1alpha1_JenkinsJobStatus(ref),
	}
}

func schema_pkg_apis_v2_v1alpha1_Jenkins(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Jenkins is the Schema for the jenkins API",
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
							Ref: ref("github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsSpec", "github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_v2_v1alpha1_JenkinsScript(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JenkinsScript is the Schema for the jenkinsscripts API",
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
							Ref: ref("github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsScriptSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsScriptStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsScriptSpec", "github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsScriptStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_v2_v1alpha1_JenkinsScriptStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JenkinsScriptStatus defines the observed state of JenkinsScript",
				Properties: map[string]spec.Schema{
					"available": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"executed": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"lastTimeUpdated": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "date-time",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_v2_v1alpha1_JenkinsServiceAccount(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JenkinsServiceAccount is the Schema for the jenkinsserviceaccounts API",
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
							Ref: ref("github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsServiceAccountSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsServiceAccountStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsServiceAccountSpec", "github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsServiceAccountStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_v2_v1alpha1_JenkinsServiceAccountSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JenkinsServiceAccountSpec defines the desired state of JenkinsServiceAccount",
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"credentials": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"ownerName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"type", "credentials"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_v2_v1alpha1_JenkinsServiceAccountStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JenkinsServiceAccountStatus defines the observed state of JenkinsServiceAccount",
				Properties: map[string]spec.Schema{
					"available": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"created": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"lastTimeUpdated": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "date-time",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_v2_v1alpha1_JenkinsStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JenkinsStatus defines the observed state of Jenkins",
				Properties: map[string]spec.Schema{
					"available": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"lastTimeUpdated": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "date-time",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"adminSecretName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"slaves": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.Slave"),
									},
								},
							},
						},
					},
					"jobProvisions": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JobProvision"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JobProvision", "github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.Slave"},
	}
}

func schema_pkg_apis_v2_v1alpha1_JenkinsFolder(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JenkinsFolder is the Schema for the jenkins folder API",
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
							Ref: ref("github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsFolderSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsFolderStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsFolderSpec", "github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsFolderStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_v2_v1alpha1_JenkinsFolderStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JenkinsFolderStatus defines the observed state of JenkinsFolder",
				Properties: map[string]spec.Schema{
					"available": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"lastTimeUpdated": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "date-time",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_v2_v1alpha1_JenkinsJob(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JenkinsJob is the Schema for the jenkinsjob API",
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
							Ref: ref("github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsJobSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsJobStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta", "github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsJobSpec", "github.com/epmd-edp/jenkins-operator/pkg/apis/v2/v1alpha1.JenkinsJobStatus"},
	}
}

func schema_pkg_apis_v2_v1alpha1_JenkinsJobSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JenkinsJobSpec defines the desired state of jenkins job",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_v2_v1alpha1_JenkinsJobStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JenkinsJobStatus defines the desired state of jenkins job",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}