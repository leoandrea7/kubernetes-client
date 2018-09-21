/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was automatically generated by informer-gen

package internalversion

import (
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	internalclientset "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "github.com/kubernetes-incubator/service-catalog/pkg/client/informers_generated/internalversion/internalinterfaces"
	internalversion "github.com/kubernetes-incubator/service-catalog/pkg/client/listers_generated/servicecatalog/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ClusterServiceBrokerInformer provides access to a shared informer and lister for
// ClusterServiceBrokers.
type ClusterServiceBrokerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.ClusterServiceBrokerLister
}

type clusterServiceBrokerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterServiceBrokerInformer constructs a new informer for ClusterServiceBroker type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterServiceBrokerInformer(client internalclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterServiceBrokerInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterServiceBrokerInformer constructs a new informer for ClusterServiceBroker type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterServiceBrokerInformer(client internalclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Servicecatalog().ClusterServiceBrokers().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Servicecatalog().ClusterServiceBrokers().Watch(options)
			},
		},
		&servicecatalog.ClusterServiceBroker{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterServiceBrokerInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterServiceBrokerInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterServiceBrokerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&servicecatalog.ClusterServiceBroker{}, f.defaultInformer)
}

func (f *clusterServiceBrokerInformer) Lister() internalversion.ClusterServiceBrokerLister {
	return internalversion.NewClusterServiceBrokerLister(f.Informer().GetIndexer())
}
