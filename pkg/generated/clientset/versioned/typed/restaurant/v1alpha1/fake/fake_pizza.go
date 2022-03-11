/*
Copyright The Kubernetes Authors.

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
	v1alpha1 "github.com/yang-wang11/pizza-apiserver/pkg/apis/restaurant/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePizzas implements PizzaInterface
type FakePizzas struct {
	Fake *FakeRestaurantV1alpha1
	ns   string
}

var pizzasResource = schema.GroupVersionResource{Group: "restaurant.yang-wang11.info", Version: "v1alpha1", Resource: "pizzas"}

var pizzasKind = schema.GroupVersionKind{Group: "restaurant.yang-wang11.info", Version: "v1alpha1", Kind: "Pizza"}

// Get takes name of the pizza, and returns the corresponding pizza object, and an error if there is any.
func (c *FakePizzas) Get(name string, options v1.GetOptions) (result *v1alpha1.Pizza, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pizzasResource, c.ns, name), &v1alpha1.Pizza{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Pizza), err
}

// List takes label and field selectors, and returns the list of Pizzas that match those selectors.
func (c *FakePizzas) List(opts v1.ListOptions) (result *v1alpha1.PizzaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pizzasResource, pizzasKind, c.ns, opts), &v1alpha1.PizzaList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PizzaList{ListMeta: obj.(*v1alpha1.PizzaList).ListMeta}
	for _, item := range obj.(*v1alpha1.PizzaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pizzas.
func (c *FakePizzas) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pizzasResource, c.ns, opts))

}

// Create takes the representation of a pizza and creates it.  Returns the server's representation of the pizza, and an error, if there is any.
func (c *FakePizzas) Create(pizza *v1alpha1.Pizza) (result *v1alpha1.Pizza, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pizzasResource, c.ns, pizza), &v1alpha1.Pizza{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Pizza), err
}

// Update takes the representation of a pizza and updates it. Returns the server's representation of the pizza, and an error, if there is any.
func (c *FakePizzas) Update(pizza *v1alpha1.Pizza) (result *v1alpha1.Pizza, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pizzasResource, c.ns, pizza), &v1alpha1.Pizza{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Pizza), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePizzas) UpdateStatus(pizza *v1alpha1.Pizza) (*v1alpha1.Pizza, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(pizzasResource, "status", c.ns, pizza), &v1alpha1.Pizza{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Pizza), err
}

// Delete takes name of the pizza and deletes it. Returns an error if one occurs.
func (c *FakePizzas) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pizzasResource, c.ns, name), &v1alpha1.Pizza{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePizzas) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pizzasResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PizzaList{})
	return err
}

// Patch applies the patch and returns the patched pizza.
func (c *FakePizzas) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Pizza, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pizzasResource, c.ns, name, pt, data, subresources...), &v1alpha1.Pizza{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Pizza), err
}
