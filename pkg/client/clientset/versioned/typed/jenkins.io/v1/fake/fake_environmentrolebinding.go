// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	jenkinsiov1 "github.com/xcloudnative/xcloud/pkg/apis/jenkins.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEnvironmentRoleBindings implements EnvironmentRoleBindingInterface
type FakeEnvironmentRoleBindings struct {
	Fake *FakeJenkinsV1
	ns   string
}

var environmentrolebindingsResource = schema.GroupVersionResource{Group: "jenkins.io", Version: "v1", Resource: "environmentrolebindings"}

var environmentrolebindingsKind = schema.GroupVersionKind{Group: "jenkins.io", Version: "v1", Kind: "EnvironmentRoleBinding"}

// Get takes name of the environmentRoleBinding, and returns the corresponding environmentRoleBinding object, and an error if there is any.
func (c *FakeEnvironmentRoleBindings) Get(name string, options v1.GetOptions) (result *jenkinsiov1.EnvironmentRoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(environmentrolebindingsResource, c.ns, name), &jenkinsiov1.EnvironmentRoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.EnvironmentRoleBinding), err
}

// List takes label and field selectors, and returns the list of EnvironmentRoleBindings that match those selectors.
func (c *FakeEnvironmentRoleBindings) List(opts v1.ListOptions) (result *jenkinsiov1.EnvironmentRoleBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(environmentrolebindingsResource, environmentrolebindingsKind, c.ns, opts), &jenkinsiov1.EnvironmentRoleBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &jenkinsiov1.EnvironmentRoleBindingList{ListMeta: obj.(*jenkinsiov1.EnvironmentRoleBindingList).ListMeta}
	for _, item := range obj.(*jenkinsiov1.EnvironmentRoleBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested environmentRoleBindings.
func (c *FakeEnvironmentRoleBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(environmentrolebindingsResource, c.ns, opts))

}

// Create takes the representation of a environmentRoleBinding and creates it.  Returns the server's representation of the environmentRoleBinding, and an error, if there is any.
func (c *FakeEnvironmentRoleBindings) Create(environmentRoleBinding *jenkinsiov1.EnvironmentRoleBinding) (result *jenkinsiov1.EnvironmentRoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(environmentrolebindingsResource, c.ns, environmentRoleBinding), &jenkinsiov1.EnvironmentRoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.EnvironmentRoleBinding), err
}

// Update takes the representation of a environmentRoleBinding and updates it. Returns the server's representation of the environmentRoleBinding, and an error, if there is any.
func (c *FakeEnvironmentRoleBindings) Update(environmentRoleBinding *jenkinsiov1.EnvironmentRoleBinding) (result *jenkinsiov1.EnvironmentRoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(environmentrolebindingsResource, c.ns, environmentRoleBinding), &jenkinsiov1.EnvironmentRoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.EnvironmentRoleBinding), err
}

// Delete takes name of the environmentRoleBinding and deletes it. Returns an error if one occurs.
func (c *FakeEnvironmentRoleBindings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(environmentrolebindingsResource, c.ns, name), &jenkinsiov1.EnvironmentRoleBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEnvironmentRoleBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(environmentrolebindingsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &jenkinsiov1.EnvironmentRoleBindingList{})
	return err
}

// Patch applies the patch and returns the patched environmentRoleBinding.
func (c *FakeEnvironmentRoleBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *jenkinsiov1.EnvironmentRoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(environmentrolebindingsResource, c.ns, name, pt, data, subresources...), &jenkinsiov1.EnvironmentRoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.EnvironmentRoleBinding), err
}
