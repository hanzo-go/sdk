# O11yGettableFieldValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Complete** | Pointer to **bool** |  | [optional] 
**Values** | Pointer to [**O11yTelemetryFieldValues**](O11yTelemetryFieldValues.md) |  | [optional] 

## Methods

### NewO11yGettableFieldValues

`func NewO11yGettableFieldValues() *O11yGettableFieldValues`

NewO11yGettableFieldValues instantiates a new O11yGettableFieldValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yGettableFieldValuesWithDefaults

`func NewO11yGettableFieldValuesWithDefaults() *O11yGettableFieldValues`

NewO11yGettableFieldValuesWithDefaults instantiates a new O11yGettableFieldValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComplete

`func (o *O11yGettableFieldValues) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *O11yGettableFieldValues) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *O11yGettableFieldValues) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *O11yGettableFieldValues) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetValues

`func (o *O11yGettableFieldValues) GetValues() O11yTelemetryFieldValues`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *O11yGettableFieldValues) GetValuesOk() (*O11yTelemetryFieldValues, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *O11yGettableFieldValues) SetValues(v O11yTelemetryFieldValues)`

SetValues sets Values field to given value.

### HasValues

`func (o *O11yGettableFieldValues) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


