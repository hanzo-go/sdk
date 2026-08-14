# O11yGettableFieldKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Complete** | Pointer to **bool** |  | [optional] 
**Keys** | Pointer to [**map[string][]O11yTelemetryFieldKey**](array.md) |  | [optional] 

## Methods

### NewO11yGettableFieldKeys

`func NewO11yGettableFieldKeys() *O11yGettableFieldKeys`

NewO11yGettableFieldKeys instantiates a new O11yGettableFieldKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yGettableFieldKeysWithDefaults

`func NewO11yGettableFieldKeysWithDefaults() *O11yGettableFieldKeys`

NewO11yGettableFieldKeysWithDefaults instantiates a new O11yGettableFieldKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComplete

`func (o *O11yGettableFieldKeys) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *O11yGettableFieldKeys) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *O11yGettableFieldKeys) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *O11yGettableFieldKeys) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetKeys

`func (o *O11yGettableFieldKeys) GetKeys() map[string][]O11yTelemetryFieldKey`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *O11yGettableFieldKeys) GetKeysOk() (*map[string][]O11yTelemetryFieldKey, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *O11yGettableFieldKeys) SetKeys(v map[string][]O11yTelemetryFieldKey)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *O11yGettableFieldKeys) HasKeys() bool`

HasKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


