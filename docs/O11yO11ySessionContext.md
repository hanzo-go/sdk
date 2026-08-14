# O11yO11ySessionContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exists** | Pointer to **bool** | Exists says whether any account carries the address. | [optional] 
**Orgs** | Pointer to [**[]O11yO11ySessionOrg**](O11yO11ySessionOrg.md) | Orgs are the orgs the address belongs to, each with its sign-in routes. | [optional] 

## Methods

### NewO11yO11ySessionContext

`func NewO11yO11ySessionContext() *O11yO11ySessionContext`

NewO11yO11ySessionContext instantiates a new O11yO11ySessionContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11ySessionContextWithDefaults

`func NewO11yO11ySessionContextWithDefaults() *O11yO11ySessionContext`

NewO11yO11ySessionContextWithDefaults instantiates a new O11yO11ySessionContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExists

`func (o *O11yO11ySessionContext) GetExists() bool`

GetExists returns the Exists field if non-nil, zero value otherwise.

### GetExistsOk

`func (o *O11yO11ySessionContext) GetExistsOk() (*bool, bool)`

GetExistsOk returns a tuple with the Exists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExists

`func (o *O11yO11ySessionContext) SetExists(v bool)`

SetExists sets Exists field to given value.

### HasExists

`func (o *O11yO11ySessionContext) HasExists() bool`

HasExists returns a boolean if a field has been set.

### GetOrgs

`func (o *O11yO11ySessionContext) GetOrgs() []O11yO11ySessionOrg`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *O11yO11ySessionContext) GetOrgsOk() (*[]O11yO11ySessionOrg, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *O11yO11ySessionContext) SetOrgs(v []O11yO11ySessionOrg)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *O11yO11ySessionContext) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


