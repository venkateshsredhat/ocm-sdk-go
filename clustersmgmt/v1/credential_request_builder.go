/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// CredentialRequestBuilder contains the data and logic needed to build 'credential_request' objects.
//
// Contains the necessary attributes to allow each operator to access the necessary AWS resources
type CredentialRequestBuilder struct {
	bitmap_           uint32
	name              string
	namespace         string
	policyPermissions []string
	serviceAccount    string
}

// NewCredentialRequest creates a new builder of 'credential_request' objects.
func NewCredentialRequest() *CredentialRequestBuilder {
	return &CredentialRequestBuilder{}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *CredentialRequestBuilder) Empty() bool {
	return b == nil || b.bitmap_ == 0
}

// Name sets the value of the 'name' attribute to the given value.
func (b *CredentialRequestBuilder) Name(value string) *CredentialRequestBuilder {
	b.name = value
	b.bitmap_ |= 1
	return b
}

// Namespace sets the value of the 'namespace' attribute to the given value.
func (b *CredentialRequestBuilder) Namespace(value string) *CredentialRequestBuilder {
	b.namespace = value
	b.bitmap_ |= 2
	return b
}

// PolicyPermissions sets the value of the 'policy_permissions' attribute to the given values.
func (b *CredentialRequestBuilder) PolicyPermissions(values ...string) *CredentialRequestBuilder {
	b.policyPermissions = make([]string, len(values))
	copy(b.policyPermissions, values)
	b.bitmap_ |= 4
	return b
}

// ServiceAccount sets the value of the 'service_account' attribute to the given value.
func (b *CredentialRequestBuilder) ServiceAccount(value string) *CredentialRequestBuilder {
	b.serviceAccount = value
	b.bitmap_ |= 8
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *CredentialRequestBuilder) Copy(object *CredentialRequest) *CredentialRequestBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.name = object.name
	b.namespace = object.namespace
	if object.policyPermissions != nil {
		b.policyPermissions = make([]string, len(object.policyPermissions))
		copy(b.policyPermissions, object.policyPermissions)
	} else {
		b.policyPermissions = nil
	}
	b.serviceAccount = object.serviceAccount
	return b
}

// Build creates a 'credential_request' object using the configuration stored in the builder.
func (b *CredentialRequestBuilder) Build() (object *CredentialRequest, err error) {
	object = new(CredentialRequest)
	object.bitmap_ = b.bitmap_
	object.name = b.name
	object.namespace = b.namespace
	if b.policyPermissions != nil {
		object.policyPermissions = make([]string, len(b.policyPermissions))
		copy(object.policyPermissions, b.policyPermissions)
	}
	object.serviceAccount = b.serviceAccount
	return
}
