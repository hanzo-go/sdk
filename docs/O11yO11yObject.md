# O11yO11yObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | Pointer to [**O11yO11yResourceRef**](O11yO11yResourceRef.md) | Resource is the resource&#39;s type and kind. | [optional] 
**Selector** | Pointer to **string** | Selector picks the instance — an FGA object string, wildcard allowed. | [optional] 

## Methods

### NewO11yO11yObject

`func NewO11yO11yObject() *O11yO11yObject`

NewO11yO11yObject instantiates a new O11yO11yObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yObjectWithDefaults

`func NewO11yO11yObjectWithDefaults() *O11yO11yObject`

NewO11yO11yObjectWithDefaults instantiates a new O11yO11yObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *O11yO11yObject) GetResource() O11yO11yResourceRef`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *O11yO11yObject) GetResourceOk() (*O11yO11yResourceRef, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *O11yO11yObject) SetResource(v O11yO11yResourceRef)`

SetResource sets Resource field to given value.

### HasResource

`func (o *O11yO11yObject) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetSelector

`func (o *O11yO11yObject) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *O11yO11yObject) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *O11yO11yObject) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *O11yO11yObject) HasSelector() bool`

HasSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


