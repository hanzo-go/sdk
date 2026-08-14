# O11yO11yObjectGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | Pointer to [**O11yO11yResourceRef**](O11yO11yResourceRef.md) | Resource is the objects&#39; type and kind. | [optional] 
**Selectors** | Pointer to **[]string** | Selectors pick the instances; a wildcard selects them all. | [optional] 

## Methods

### NewO11yO11yObjectGroup

`func NewO11yO11yObjectGroup() *O11yO11yObjectGroup`

NewO11yO11yObjectGroup instantiates a new O11yO11yObjectGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yObjectGroupWithDefaults

`func NewO11yO11yObjectGroupWithDefaults() *O11yO11yObjectGroup`

NewO11yO11yObjectGroupWithDefaults instantiates a new O11yO11yObjectGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *O11yO11yObjectGroup) GetResource() O11yO11yResourceRef`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *O11yO11yObjectGroup) GetResourceOk() (*O11yO11yResourceRef, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *O11yO11yObjectGroup) SetResource(v O11yO11yResourceRef)`

SetResource sets Resource field to given value.

### HasResource

`func (o *O11yO11yObjectGroup) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetSelectors

`func (o *O11yO11yObjectGroup) GetSelectors() []string`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *O11yO11yObjectGroup) GetSelectorsOk() (*[]string, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *O11yO11yObjectGroup) SetSelectors(v []string)`

SetSelectors sets Selectors field to given value.

### HasSelectors

`func (o *O11yO11yObjectGroup) HasSelectors() bool`

HasSelectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


