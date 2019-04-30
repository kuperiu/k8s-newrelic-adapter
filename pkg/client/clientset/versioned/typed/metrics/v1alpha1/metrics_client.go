// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// A copy of the License is located at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/awslabs/k8s-cloudwatch-adapter/pkg/apis/metrics/v1alpha1"
	"github.com/awslabs/k8s-cloudwatch-adapter/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type MetricsV1alpha1Interface interface {
	RESTClient() rest.Interface
	ExternalMetricsGetter
}

// MetricsV1alpha1Client is used to interact with features provided by the metrics.aws group.
type MetricsV1alpha1Client struct {
	restClient rest.Interface
}

func (c *MetricsV1alpha1Client) ExternalMetrics(namespace string) ExternalMetricInterface {
	return newExternalMetrics(c, namespace)
}

// NewForConfig creates a new MetricsV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*MetricsV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &MetricsV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new MetricsV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *MetricsV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new MetricsV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *MetricsV1alpha1Client {
	return &MetricsV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *MetricsV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
