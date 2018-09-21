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
package fake

import (
	user "github.com/openshift/origin/pkg/user/apis/user"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIdentities implements IdentityInterface
type FakeIdentities struct {
	Fake *FakeUser
}

var identitiesResource = schema.GroupVersionResource{Group: "user.openshift.io", Version: "", Resource: "identities"}

var identitiesKind = schema.GroupVersionKind{Group: "user.openshift.io", Version: "", Kind: "Identity"}

// Get takes name of the identity, and returns the corresponding identity object, and an error if there is any.
func (c *FakeIdentities) Get(name string, options v1.GetOptions) (result *user.Identity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(identitiesResource, name), &user.Identity{})
	if obj == nil {
		return nil, err
	}
	return obj.(*user.Identity), err
}

// List takes label and field selectors, and returns the list of Identities that match those selectors.
func (c *FakeIdentities) List(opts v1.ListOptions) (result *user.IdentityList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(identitiesResource, identitiesKind, opts), &user.IdentityList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &user.IdentityList{}
	for _, item := range obj.(*user.IdentityList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested identities.
func (c *FakeIdentities) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(identitiesResource, opts))
}

// Create takes the representation of a identity and creates it.  Returns the server's representation of the identity, and an error, if there is any.
func (c *FakeIdentities) Create(identity *user.Identity) (result *user.Identity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(identitiesResource, identity), &user.Identity{})
	if obj == nil {
		return nil, err
	}
	return obj.(*user.Identity), err
}

// Update takes the representation of a identity and updates it. Returns the server's representation of the identity, and an error, if there is any.
func (c *FakeIdentities) Update(identity *user.Identity) (result *user.Identity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(identitiesResource, identity), &user.Identity{})
	if obj == nil {
		return nil, err
	}
	return obj.(*user.Identity), err
}

// Delete takes name of the identity and deletes it. Returns an error if one occurs.
func (c *FakeIdentities) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(identitiesResource, name), &user.Identity{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIdentities) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(identitiesResource, listOptions)

	_, err := c.Fake.Invokes(action, &user.IdentityList{})
	return err
}

// Patch applies the patch and returns the patched identity.
func (c *FakeIdentities) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *user.Identity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(identitiesResource, name, data, subresources...), &user.Identity{})
	if obj == nil {
		return nil, err
	}
	return obj.(*user.Identity), err
}
