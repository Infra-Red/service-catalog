/*
Copyright 2024 The Kubernetes Authors.

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

package v1beta1

import (
	"context"
	time "time"

	servicecatalogv1beta1 "github.com/drycc-addons/service-catalog/pkg/apis/servicecatalog/v1beta1"
	clientset "github.com/drycc-addons/service-catalog/pkg/client/clientset_generated/clientset"
	internalinterfaces "github.com/drycc-addons/service-catalog/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1beta1 "github.com/drycc-addons/service-catalog/pkg/client/listers_generated/servicecatalog/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ServiceClassInformer provides access to a shared informer and lister for
// ServiceClasses.
type ServiceClassInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ServiceClassLister
}

type serviceClassInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewServiceClassInformer constructs a new informer for ServiceClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServiceClassInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServiceClassInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredServiceClassInformer constructs a new informer for ServiceClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServiceClassInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServicecatalogV1beta1().ServiceClasses(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServicecatalogV1beta1().ServiceClasses(namespace).Watch(context.TODO(), options)
			},
		},
		&servicecatalogv1beta1.ServiceClass{},
		resyncPeriod,
		indexers,
	)
}

func (f *serviceClassInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServiceClassInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serviceClassInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&servicecatalogv1beta1.ServiceClass{}, f.defaultInformer)
}

func (f *serviceClassInformer) Lister() v1beta1.ServiceClassLister {
	return v1beta1.NewServiceClassLister(f.Informer().GetIndexer())
}
