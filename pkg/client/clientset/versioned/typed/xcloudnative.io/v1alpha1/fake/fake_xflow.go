// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/xcloudnative/xcloud/pkg/apis/xcloudnative.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeXflows implements XflowInterface
type FakeXflows struct {
	Fake *FakeXcloudnativeV1alpha1
	ns   string
}

var xflowsResource = schema.GroupVersionResource{Group: "xcloudnative.io", Version: "v1alpha1", Resource: "xflows"}

var xflowsKind = schema.GroupVersionKind{Group: "xcloudnative.io", Version: "v1alpha1", Kind: "Xflow"}

// Get takes name of the xflow, and returns the corresponding xflow object, and an error if there is any.
func (c *FakeXflows) Get(name string, options v1.GetOptions) (result *v1alpha1.Xflow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(xflowsResource, c.ns, name), &v1alpha1.Xflow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Xflow), err
}

// List takes label and field selectors, and returns the list of Xflows that match those selectors.
func (c *FakeXflows) List(opts v1.ListOptions) (result *v1alpha1.XflowList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(xflowsResource, xflowsKind, c.ns, opts), &v1alpha1.XflowList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.XflowList{ListMeta: obj.(*v1alpha1.XflowList).ListMeta}
	for _, item := range obj.(*v1alpha1.XflowList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested xflows.
func (c *FakeXflows) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(xflowsResource, c.ns, opts))

}

// Create takes the representation of a xflow and creates it.  Returns the server's representation of the xflow, and an error, if there is any.
func (c *FakeXflows) Create(xflow *v1alpha1.Xflow) (result *v1alpha1.Xflow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(xflowsResource, c.ns, xflow), &v1alpha1.Xflow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Xflow), err
}

// Update takes the representation of a xflow and updates it. Returns the server's representation of the xflow, and an error, if there is any.
func (c *FakeXflows) Update(xflow *v1alpha1.Xflow) (result *v1alpha1.Xflow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(xflowsResource, c.ns, xflow), &v1alpha1.Xflow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Xflow), err
}

// Delete takes name of the xflow and deletes it. Returns an error if one occurs.
func (c *FakeXflows) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(xflowsResource, c.ns, name), &v1alpha1.Xflow{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeXflows) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(xflowsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.XflowList{})
	return err
}

// Patch applies the patch and returns the patched xflow.
func (c *FakeXflows) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Xflow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(xflowsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Xflow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Xflow), err
}
