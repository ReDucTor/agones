// Copyright 2018 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was autogenerated. Do not edit directly.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "agones.dev/agones/pkg/apis/stable/v1alpha1"
	scheme "agones.dev/agones/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GameServerSetsGetter has a method to return a GameServerSetInterface.
// A group's client should implement this interface.
type GameServerSetsGetter interface {
	GameServerSets(namespace string) GameServerSetInterface
}

// GameServerSetInterface has methods to work with GameServerSet resources.
type GameServerSetInterface interface {
	Create(*v1alpha1.GameServerSet) (*v1alpha1.GameServerSet, error)
	Update(*v1alpha1.GameServerSet) (*v1alpha1.GameServerSet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GameServerSet, error)
	List(opts v1.ListOptions) (*v1alpha1.GameServerSetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GameServerSet, err error)
	GameServerSetExpansion
}

// gameServerSets implements GameServerSetInterface
type gameServerSets struct {
	client rest.Interface
	ns     string
}

// newGameServerSets returns a GameServerSets
func newGameServerSets(c *StableV1alpha1Client, namespace string) *gameServerSets {
	return &gameServerSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the gameServerSet, and returns the corresponding gameServerSet object, and an error if there is any.
func (c *gameServerSets) Get(name string, options v1.GetOptions) (result *v1alpha1.GameServerSet, err error) {
	result = &v1alpha1.GameServerSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gameserversets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GameServerSets that match those selectors.
func (c *gameServerSets) List(opts v1.ListOptions) (result *v1alpha1.GameServerSetList, err error) {
	result = &v1alpha1.GameServerSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gameserversets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gameServerSets.
func (c *gameServerSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("gameserversets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a gameServerSet and creates it.  Returns the server's representation of the gameServerSet, and an error, if there is any.
func (c *gameServerSets) Create(gameServerSet *v1alpha1.GameServerSet) (result *v1alpha1.GameServerSet, err error) {
	result = &v1alpha1.GameServerSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gameserversets").
		Body(gameServerSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a gameServerSet and updates it. Returns the server's representation of the gameServerSet, and an error, if there is any.
func (c *gameServerSets) Update(gameServerSet *v1alpha1.GameServerSet) (result *v1alpha1.GameServerSet, err error) {
	result = &v1alpha1.GameServerSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gameserversets").
		Name(gameServerSet.Name).
		Body(gameServerSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the gameServerSet and deletes it. Returns an error if one occurs.
func (c *gameServerSets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gameserversets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gameServerSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gameserversets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched gameServerSet.
func (c *gameServerSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GameServerSet, err error) {
	result = &v1alpha1.GameServerSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("gameserversets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
