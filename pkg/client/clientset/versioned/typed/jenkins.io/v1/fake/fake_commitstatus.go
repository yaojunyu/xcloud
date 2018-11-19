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

// FakeCommitStatuses implements CommitStatusInterface
type FakeCommitStatuses struct {
	Fake *FakeJenkinsV1
	ns   string
}

var commitstatusesResource = schema.GroupVersionResource{Group: "jenkins.io", Version: "v1", Resource: "commitstatuses"}

var commitstatusesKind = schema.GroupVersionKind{Group: "jenkins.io", Version: "v1", Kind: "CommitStatus"}

// Get takes name of the commitStatus, and returns the corresponding commitStatus object, and an error if there is any.
func (c *FakeCommitStatuses) Get(name string, options v1.GetOptions) (result *jenkinsiov1.CommitStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(commitstatusesResource, c.ns, name), &jenkinsiov1.CommitStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.CommitStatus), err
}

// List takes label and field selectors, and returns the list of CommitStatuses that match those selectors.
func (c *FakeCommitStatuses) List(opts v1.ListOptions) (result *jenkinsiov1.CommitStatusList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(commitstatusesResource, commitstatusesKind, c.ns, opts), &jenkinsiov1.CommitStatusList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &jenkinsiov1.CommitStatusList{ListMeta: obj.(*jenkinsiov1.CommitStatusList).ListMeta}
	for _, item := range obj.(*jenkinsiov1.CommitStatusList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested commitStatuses.
func (c *FakeCommitStatuses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(commitstatusesResource, c.ns, opts))

}

// Create takes the representation of a commitStatus and creates it.  Returns the server's representation of the commitStatus, and an error, if there is any.
func (c *FakeCommitStatuses) Create(commitStatus *jenkinsiov1.CommitStatus) (result *jenkinsiov1.CommitStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(commitstatusesResource, c.ns, commitStatus), &jenkinsiov1.CommitStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.CommitStatus), err
}

// Update takes the representation of a commitStatus and updates it. Returns the server's representation of the commitStatus, and an error, if there is any.
func (c *FakeCommitStatuses) Update(commitStatus *jenkinsiov1.CommitStatus) (result *jenkinsiov1.CommitStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(commitstatusesResource, c.ns, commitStatus), &jenkinsiov1.CommitStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.CommitStatus), err
}

// Delete takes name of the commitStatus and deletes it. Returns an error if one occurs.
func (c *FakeCommitStatuses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(commitstatusesResource, c.ns, name), &jenkinsiov1.CommitStatus{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCommitStatuses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(commitstatusesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &jenkinsiov1.CommitStatusList{})
	return err
}

// Patch applies the patch and returns the patched commitStatus.
func (c *FakeCommitStatuses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *jenkinsiov1.CommitStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(commitstatusesResource, c.ns, name, pt, data, subresources...), &jenkinsiov1.CommitStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.CommitStatus), err
}
