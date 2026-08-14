# O11yO11yDomainsAnswer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**O11yO11yDomainsData**](O11yO11yDomainsData.md) | Data holds the per-query results, emitted LAST — see the field-order note on the type. | [optional] 
**Meta** | Pointer to [**O11yO11yQueryStats**](O11yO11yQueryStats.md) | Meta reports what the read cost. | [optional] 
**Type** | Pointer to **string** | Type names the result shape: scalar, time_series or raw. | [optional] 
**Warning** | Pointer to [**O11yO11yQueryWarning**](O11yO11yQueryWarning.md) | Warning carries the store&#39;s warning for this read, when it raised one. | [optional] 

## Methods

### NewO11yO11yDomainsAnswer

`func NewO11yO11yDomainsAnswer() *O11yO11yDomainsAnswer`

NewO11yO11yDomainsAnswer instantiates a new O11yO11yDomainsAnswer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yDomainsAnswerWithDefaults

`func NewO11yO11yDomainsAnswerWithDefaults() *O11yO11yDomainsAnswer`

NewO11yO11yDomainsAnswerWithDefaults instantiates a new O11yO11yDomainsAnswer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11yDomainsAnswer) GetData() O11yO11yDomainsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yDomainsAnswer) GetDataOk() (*O11yO11yDomainsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yDomainsAnswer) SetData(v O11yO11yDomainsData)`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yDomainsAnswer) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *O11yO11yDomainsAnswer) GetMeta() O11yO11yQueryStats`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *O11yO11yDomainsAnswer) GetMetaOk() (*O11yO11yQueryStats, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *O11yO11yDomainsAnswer) SetMeta(v O11yO11yQueryStats)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *O11yO11yDomainsAnswer) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetType

`func (o *O11yO11yDomainsAnswer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yO11yDomainsAnswer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yO11yDomainsAnswer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yO11yDomainsAnswer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWarning

`func (o *O11yO11yDomainsAnswer) GetWarning() O11yO11yQueryWarning`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *O11yO11yDomainsAnswer) GetWarningOk() (*O11yO11yQueryWarning, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *O11yO11yDomainsAnswer) SetWarning(v O11yO11yQueryWarning)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *O11yO11yDomainsAnswer) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


