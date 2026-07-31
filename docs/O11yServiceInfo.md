# O11yServiceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**SpanCount** | Pointer to **int64** |  | [optional] 
**ErrorCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewO11yServiceInfo

`func NewO11yServiceInfo() *O11yServiceInfo`

NewO11yServiceInfo instantiates a new O11yServiceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yServiceInfoWithDefaults

`func NewO11yServiceInfoWithDefaults() *O11yServiceInfo`

NewO11yServiceInfoWithDefaults instantiates a new O11yServiceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *O11yServiceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yServiceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yServiceInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yServiceInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpanCount

`func (o *O11yServiceInfo) GetSpanCount() int64`

GetSpanCount returns the SpanCount field if non-nil, zero value otherwise.

### GetSpanCountOk

`func (o *O11yServiceInfo) GetSpanCountOk() (*int64, bool)`

GetSpanCountOk returns a tuple with the SpanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanCount

`func (o *O11yServiceInfo) SetSpanCount(v int64)`

SetSpanCount sets SpanCount field to given value.

### HasSpanCount

`func (o *O11yServiceInfo) HasSpanCount() bool`

HasSpanCount returns a boolean if a field has been set.

### GetErrorCount

`func (o *O11yServiceInfo) GetErrorCount() int64`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *O11yServiceInfo) GetErrorCountOk() (*int64, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *O11yServiceInfo) SetErrorCount(v int64)`

SetErrorCount sets ErrorCount field to given value.

### HasErrorCount

`func (o *O11yServiceInfo) HasErrorCount() bool`

HasErrorCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


