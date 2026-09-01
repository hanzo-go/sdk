# LinearClaimOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the organization&#39;s URL key. | [optional] 
**Organization** | Pointer to **string** | Organization is the Linear organization id now bound to this org. | [optional] 
**Path** | Pointer to **string** | Path is the address to configure in Linear, on this deployment&#39;s origin. | [optional] 

## Methods

### NewLinearClaimOut

`func NewLinearClaimOut() *LinearClaimOut`

NewLinearClaimOut instantiates a new LinearClaimOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinearClaimOutWithDefaults

`func NewLinearClaimOutWithDefaults() *LinearClaimOut`

NewLinearClaimOutWithDefaults instantiates a new LinearClaimOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LinearClaimOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LinearClaimOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LinearClaimOut) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LinearClaimOut) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *LinearClaimOut) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *LinearClaimOut) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *LinearClaimOut) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *LinearClaimOut) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPath

`func (o *LinearClaimOut) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LinearClaimOut) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LinearClaimOut) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LinearClaimOut) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


