/*
Copyright The Space Cloud Authors.

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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/spacecloud-io/space-cloud/pkg/apis/core/v1alpha1"
	scheme "github.com/spacecloud-io/space-cloud/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CompiledGraphqlSourcesGetter has a method to return a CompiledGraphqlSourceInterface.
// A group's client should implement this interface.
type CompiledGraphqlSourcesGetter interface {
	CompiledGraphqlSources(namespace string) CompiledGraphqlSourceInterface
}

// CompiledGraphqlSourceInterface has methods to work with CompiledGraphqlSource resources.
type CompiledGraphqlSourceInterface interface {
	Create(ctx context.Context, compiledGraphqlSource *v1alpha1.CompiledGraphqlSource, opts v1.CreateOptions) (*v1alpha1.CompiledGraphqlSource, error)
	Update(ctx context.Context, compiledGraphqlSource *v1alpha1.CompiledGraphqlSource, opts v1.UpdateOptions) (*v1alpha1.CompiledGraphqlSource, error)
	UpdateStatus(ctx context.Context, compiledGraphqlSource *v1alpha1.CompiledGraphqlSource, opts v1.UpdateOptions) (*v1alpha1.CompiledGraphqlSource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.CompiledGraphqlSource, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.CompiledGraphqlSourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CompiledGraphqlSource, err error)
	CompiledGraphqlSourceExpansion
}

// compiledGraphqlSources implements CompiledGraphqlSourceInterface
type compiledGraphqlSources struct {
	client rest.Interface
	ns     string
}

// newCompiledGraphqlSources returns a CompiledGraphqlSources
func newCompiledGraphqlSources(c *CoreV1alpha1Client, namespace string) *compiledGraphqlSources {
	return &compiledGraphqlSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the compiledGraphqlSource, and returns the corresponding compiledGraphqlSource object, and an error if there is any.
func (c *compiledGraphqlSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CompiledGraphqlSource, err error) {
	result = &v1alpha1.CompiledGraphqlSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("compiledgraphqlsources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CompiledGraphqlSources that match those selectors.
func (c *compiledGraphqlSources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CompiledGraphqlSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CompiledGraphqlSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("compiledgraphqlsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested compiledGraphqlSources.
func (c *compiledGraphqlSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("compiledgraphqlsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a compiledGraphqlSource and creates it.  Returns the server's representation of the compiledGraphqlSource, and an error, if there is any.
func (c *compiledGraphqlSources) Create(ctx context.Context, compiledGraphqlSource *v1alpha1.CompiledGraphqlSource, opts v1.CreateOptions) (result *v1alpha1.CompiledGraphqlSource, err error) {
	result = &v1alpha1.CompiledGraphqlSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("compiledgraphqlsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(compiledGraphqlSource).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a compiledGraphqlSource and updates it. Returns the server's representation of the compiledGraphqlSource, and an error, if there is any.
func (c *compiledGraphqlSources) Update(ctx context.Context, compiledGraphqlSource *v1alpha1.CompiledGraphqlSource, opts v1.UpdateOptions) (result *v1alpha1.CompiledGraphqlSource, err error) {
	result = &v1alpha1.CompiledGraphqlSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("compiledgraphqlsources").
		Name(compiledGraphqlSource.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(compiledGraphqlSource).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *compiledGraphqlSources) UpdateStatus(ctx context.Context, compiledGraphqlSource *v1alpha1.CompiledGraphqlSource, opts v1.UpdateOptions) (result *v1alpha1.CompiledGraphqlSource, err error) {
	result = &v1alpha1.CompiledGraphqlSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("compiledgraphqlsources").
		Name(compiledGraphqlSource.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(compiledGraphqlSource).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the compiledGraphqlSource and deletes it. Returns an error if one occurs.
func (c *compiledGraphqlSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("compiledgraphqlsources").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *compiledGraphqlSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("compiledgraphqlsources").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched compiledGraphqlSource.
func (c *compiledGraphqlSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CompiledGraphqlSource, err error) {
	result = &v1alpha1.CompiledGraphqlSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("compiledgraphqlsources").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
