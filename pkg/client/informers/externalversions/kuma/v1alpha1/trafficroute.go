/*
Copyright 2020 The Flux authors

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	kumav1alpha1 "github.com/fluxcd/flagger/pkg/apis/kuma/v1alpha1"
	versioned "github.com/fluxcd/flagger/pkg/client/clientset/versioned"
	internalinterfaces "github.com/fluxcd/flagger/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/fluxcd/flagger/pkg/client/listers/kuma/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TrafficRouteInformer provides access to a shared informer and lister for
// TrafficRoutes.
type TrafficRouteInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.TrafficRouteLister
}

type trafficRouteInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewTrafficRouteInformer constructs a new informer for TrafficRoute type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTrafficRouteInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTrafficRouteInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredTrafficRouteInformer constructs a new informer for TrafficRoute type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTrafficRouteInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KumaV1alpha1().TrafficRoutes().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KumaV1alpha1().TrafficRoutes().Watch(context.TODO(), options)
			},
		},
		&kumav1alpha1.TrafficRoute{},
		resyncPeriod,
		indexers,
	)
}

func (f *trafficRouteInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTrafficRouteInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *trafficRouteInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kumav1alpha1.TrafficRoute{}, f.defaultInformer)
}

func (f *trafficRouteInformer) Lister() v1alpha1.TrafficRouteLister {
	return v1alpha1.NewTrafficRouteLister(f.Informer().GetIndexer())
}
