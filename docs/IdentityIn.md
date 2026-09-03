# IdentityIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the device&#39;s name within the org — a DNS label. The fabric knows the identity as \&quot;&lt;name&gt;.&lt;org&gt;\&quot;; every answer here uses the caller&#39;s name. | [optional] 
**Roles** | Pointer to **[]string** | Roles are extra role attributes for the identity, each scoped to the caller&#39;s org on the way in (\&quot;k3s-host\&quot; is written as \&quot;k3s-host.&lt;org&gt;\&quot;) so no caller can claim an attribute another tenant&#39;s policy selects. A role of the form \&quot;&lt;service&gt;-host\&quot; makes this identity a HOST of that published service — the bind policy from POST /v1/network/services selects exactly that attribute — and is refused when the org has no such service. | [optional] 

## Methods

### NewIdentityIn

`func NewIdentityIn() *IdentityIn`

NewIdentityIn instantiates a new IdentityIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityInWithDefaults

`func NewIdentityInWithDefaults() *IdentityIn`

NewIdentityInWithDefaults instantiates a new IdentityIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IdentityIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoles

`func (o *IdentityIn) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IdentityIn) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IdentityIn) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *IdentityIn) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


