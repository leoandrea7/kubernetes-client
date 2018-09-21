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
package internalversion

import (
	security "github.com/openshift/origin/pkg/security/apis/security"
	rest "k8s.io/client-go/rest"
)

// PodSecurityPolicyReviewsGetter has a method to return a PodSecurityPolicyReviewInterface.
// A group's client should implement this interface.
type PodSecurityPolicyReviewsGetter interface {
	PodSecurityPolicyReviews(namespace string) PodSecurityPolicyReviewInterface
}

// PodSecurityPolicyReviewInterface has methods to work with PodSecurityPolicyReview resources.
type PodSecurityPolicyReviewInterface interface {
	Create(*security.PodSecurityPolicyReview) (*security.PodSecurityPolicyReview, error)
	PodSecurityPolicyReviewExpansion
}

// podSecurityPolicyReviews implements PodSecurityPolicyReviewInterface
type podSecurityPolicyReviews struct {
	client rest.Interface
	ns     string
}

// newPodSecurityPolicyReviews returns a PodSecurityPolicyReviews
func newPodSecurityPolicyReviews(c *SecurityClient, namespace string) *podSecurityPolicyReviews {
	return &podSecurityPolicyReviews{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a podSecurityPolicyReview and creates it.  Returns the server's representation of the podSecurityPolicyReview, and an error, if there is any.
func (c *podSecurityPolicyReviews) Create(podSecurityPolicyReview *security.PodSecurityPolicyReview) (result *security.PodSecurityPolicyReview, err error) {
	result = &security.PodSecurityPolicyReview{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("podsecuritypolicyreviews").
		Body(podSecurityPolicyReview).
		Do().
		Into(result)
	return
}
