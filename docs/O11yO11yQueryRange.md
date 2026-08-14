# O11yO11yQueryRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**O11yO11yQueryRangeData**](O11yO11yQueryRangeData.md) | Data holds the results. | [optional] 
**Meta** | Pointer to [**O11yO11yQueryStats**](O11yO11yQueryStats.md) | Meta reports what the query scanned. | [optional] 
**Type** | Pointer to **string** | Type is the result kind; time_series here. | [optional] 
**Warning** | Pointer to [**O11yO11yQueryWarning**](O11yO11yQueryWarning.md) | Warning carries a non-fatal warning, when the query raised one. | [optional] 

## Methods

### NewO11yO11yQueryRange

`func NewO11yO11yQueryRange() *O11yO11yQueryRange`

NewO11yO11yQueryRange instantiates a new O11yO11yQueryRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yQueryRangeWithDefaults

`func NewO11yO11yQueryRangeWithDefaults() *O11yO11yQueryRange`

NewO11yO11yQueryRangeWithDefaults instantiates a new O11yO11yQueryRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11yQueryRange) GetData() O11yO11yQueryRangeData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yQueryRange) GetDataOk() (*O11yO11yQueryRangeData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yQueryRange) SetData(v O11yO11yQueryRangeData)`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yQueryRange) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *O11yO11yQueryRange) GetMeta() O11yO11yQueryStats`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *O11yO11yQueryRange) GetMetaOk() (*O11yO11yQueryStats, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *O11yO11yQueryRange) SetMeta(v O11yO11yQueryStats)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *O11yO11yQueryRange) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetType

`func (o *O11yO11yQueryRange) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yO11yQueryRange) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yO11yQueryRange) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yO11yQueryRange) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWarning

`func (o *O11yO11yQueryRange) GetWarning() O11yO11yQueryWarning`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *O11yO11yQueryRange) GetWarningOk() (*O11yO11yQueryWarning, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *O11yO11yQueryRange) SetWarning(v O11yO11yQueryWarning)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *O11yO11yQueryRange) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


