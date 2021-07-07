// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	monitoringv1alpha1 "github.com/GoogleCloudPlatform/prometheus-engine/pkg/operator/apis/monitoring/v1alpha1"
	versioned "github.com/GoogleCloudPlatform/prometheus-engine/pkg/operator/generated/clientset/versioned"
	internalinterfaces "github.com/GoogleCloudPlatform/prometheus-engine/pkg/operator/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/GoogleCloudPlatform/prometheus-engine/pkg/operator/generated/listers/monitoring/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// RulesInformer provides access to a shared informer and lister for
// Rules.
type RulesInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.RulesLister
}

type rulesInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewRulesInformer constructs a new informer for Rules type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRulesInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRulesInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredRulesInformer constructs a new informer for Rules type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRulesInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MonitoringV1alpha1().Rules(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MonitoringV1alpha1().Rules(namespace).Watch(context.TODO(), options)
			},
		},
		&monitoringv1alpha1.Rules{},
		resyncPeriod,
		indexers,
	)
}

func (f *rulesInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRulesInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *rulesInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&monitoringv1alpha1.Rules{}, f.defaultInformer)
}

func (f *rulesInformer) Lister() v1alpha1.RulesLister {
	return v1alpha1.NewRulesLister(f.Informer().GetIndexer())
}