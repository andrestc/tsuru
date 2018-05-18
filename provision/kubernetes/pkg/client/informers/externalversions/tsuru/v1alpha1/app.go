// Copyright 2018 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was automatically generated by informer-gen

package v1alpha1

import (
	tsuru_v1alpha1 "github.com/tsuru/tsuru/provision/kubernetes/pkg/apis/tsuru/v1alpha1"
	versioned "github.com/tsuru/tsuru/provision/kubernetes/pkg/client/clientset/versioned"
	internalinterfaces "github.com/tsuru/tsuru/provision/kubernetes/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/tsuru/tsuru/provision/kubernetes/pkg/client/listers/tsuru/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// AppInformer provides access to a shared informer and lister for
// Apps.
type AppInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AppLister
}

type appInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAppInformer constructs a new informer for App type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAppInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAppInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAppInformer constructs a new informer for App type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAppInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TsuruV1alpha1().Apps(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TsuruV1alpha1().Apps(namespace).Watch(options)
			},
		},
		&tsuru_v1alpha1.App{},
		resyncPeriod,
		indexers,
	)
}

func (f *appInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAppInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *appInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&tsuru_v1alpha1.App{}, f.defaultInformer)
}

func (f *appInformer) Lister() v1alpha1.AppLister {
	return v1alpha1.NewAppLister(f.Informer().GetIndexer())
}
