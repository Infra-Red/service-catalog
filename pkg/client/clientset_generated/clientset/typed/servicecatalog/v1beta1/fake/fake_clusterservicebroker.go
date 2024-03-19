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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/drycc-addons/service-catalog/pkg/apis/servicecatalog/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterServiceBrokers implements ClusterServiceBrokerInterface
type FakeClusterServiceBrokers struct {
	Fake *FakeServicecatalogV1beta1
}

var clusterservicebrokersResource = v1beta1.SchemeGroupVersion.WithResource("clusterservicebrokers")

var clusterservicebrokersKind = v1beta1.SchemeGroupVersion.WithKind("ClusterServiceBroker")

// Get takes name of the clusterServiceBroker, and returns the corresponding clusterServiceBroker object, and an error if there is any.
func (c *FakeClusterServiceBrokers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ClusterServiceBroker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterservicebrokersResource, name), &v1beta1.ClusterServiceBroker{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterServiceBroker), err
}

// List takes label and field selectors, and returns the list of ClusterServiceBrokers that match those selectors.
func (c *FakeClusterServiceBrokers) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ClusterServiceBrokerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterservicebrokersResource, clusterservicebrokersKind, opts), &v1beta1.ClusterServiceBrokerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ClusterServiceBrokerList{ListMeta: obj.(*v1beta1.ClusterServiceBrokerList).ListMeta}
	for _, item := range obj.(*v1beta1.ClusterServiceBrokerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterServiceBrokers.
func (c *FakeClusterServiceBrokers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterservicebrokersResource, opts))
}

// Create takes the representation of a clusterServiceBroker and creates it.  Returns the server's representation of the clusterServiceBroker, and an error, if there is any.
func (c *FakeClusterServiceBrokers) Create(ctx context.Context, clusterServiceBroker *v1beta1.ClusterServiceBroker, opts v1.CreateOptions) (result *v1beta1.ClusterServiceBroker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterservicebrokersResource, clusterServiceBroker), &v1beta1.ClusterServiceBroker{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterServiceBroker), err
}

// Update takes the representation of a clusterServiceBroker and updates it. Returns the server's representation of the clusterServiceBroker, and an error, if there is any.
func (c *FakeClusterServiceBrokers) Update(ctx context.Context, clusterServiceBroker *v1beta1.ClusterServiceBroker, opts v1.UpdateOptions) (result *v1beta1.ClusterServiceBroker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterservicebrokersResource, clusterServiceBroker), &v1beta1.ClusterServiceBroker{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterServiceBroker), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterServiceBrokers) UpdateStatus(ctx context.Context, clusterServiceBroker *v1beta1.ClusterServiceBroker, opts v1.UpdateOptions) (*v1beta1.ClusterServiceBroker, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterservicebrokersResource, "status", clusterServiceBroker), &v1beta1.ClusterServiceBroker{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterServiceBroker), err
}

// Delete takes name of the clusterServiceBroker and deletes it. Returns an error if one occurs.
func (c *FakeClusterServiceBrokers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusterservicebrokersResource, name, opts), &v1beta1.ClusterServiceBroker{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterServiceBrokers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterservicebrokersResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ClusterServiceBrokerList{})
	return err
}

// Patch applies the patch and returns the patched clusterServiceBroker.
func (c *FakeClusterServiceBrokers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ClusterServiceBroker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterservicebrokersResource, name, pt, data, subresources...), &v1beta1.ClusterServiceBroker{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterServiceBroker), err
}
