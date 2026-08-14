# O11yEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeMap** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**IsError** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**TimeUnixNano** | Pointer to **int32** |  | [optional] 

## Methods

### NewO11yEvent

`func NewO11yEvent() *O11yEvent`

NewO11yEvent instantiates a new O11yEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yEventWithDefaults

`func NewO11yEventWithDefaults() *O11yEvent`

NewO11yEventWithDefaults instantiates a new O11yEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeMap

`func (o *O11yEvent) GetAttributeMap() map[string]map[string]interface{}`

GetAttributeMap returns the AttributeMap field if non-nil, zero value otherwise.

### GetAttributeMapOk

`func (o *O11yEvent) GetAttributeMapOk() (*map[string]map[string]interface{}, bool)`

GetAttributeMapOk returns a tuple with the AttributeMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeMap

`func (o *O11yEvent) SetAttributeMap(v map[string]map[string]interface{})`

SetAttributeMap sets AttributeMap field to given value.

### HasAttributeMap

`func (o *O11yEvent) HasAttributeMap() bool`

HasAttributeMap returns a boolean if a field has been set.

### GetIsError

`func (o *O11yEvent) GetIsError() bool`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *O11yEvent) GetIsErrorOk() (*bool, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *O11yEvent) SetIsError(v bool)`

SetIsError sets IsError field to given value.

### HasIsError

`func (o *O11yEvent) HasIsError() bool`

HasIsError returns a boolean if a field has been set.

### GetName

`func (o *O11yEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimeUnixNano

`func (o *O11yEvent) GetTimeUnixNano() int32`

GetTimeUnixNano returns the TimeUnixNano field if non-nil, zero value otherwise.

### GetTimeUnixNanoOk

`func (o *O11yEvent) GetTimeUnixNanoOk() (*int32, bool)`

GetTimeUnixNanoOk returns a tuple with the TimeUnixNano field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnixNano

`func (o *O11yEvent) SetTimeUnixNano(v int32)`

SetTimeUnixNano sets TimeUnixNano field to given value.

### HasTimeUnixNano

`func (o *O11yEvent) HasTimeUnixNano() bool`

HasTimeUnixNano returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


