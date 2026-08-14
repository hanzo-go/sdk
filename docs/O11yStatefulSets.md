# O11yStatefulSets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTimeBeforeRetention** | Pointer to **bool** |  | [optional] 
**Records** | Pointer to [**[]O11yStatefulSetRecord**](O11yStatefulSetRecord.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **interface{}** |  | [optional] 
**Warning** | Pointer to [**O11yQueryWarnData**](O11yQueryWarnData.md) |  | [optional] 

## Methods

### NewO11yStatefulSets

`func NewO11yStatefulSets() *O11yStatefulSets`

NewO11yStatefulSets instantiates a new O11yStatefulSets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yStatefulSetsWithDefaults

`func NewO11yStatefulSetsWithDefaults() *O11yStatefulSets`

NewO11yStatefulSetsWithDefaults instantiates a new O11yStatefulSets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTimeBeforeRetention

`func (o *O11yStatefulSets) GetEndTimeBeforeRetention() bool`

GetEndTimeBeforeRetention returns the EndTimeBeforeRetention field if non-nil, zero value otherwise.

### GetEndTimeBeforeRetentionOk

`func (o *O11yStatefulSets) GetEndTimeBeforeRetentionOk() (*bool, bool)`

GetEndTimeBeforeRetentionOk returns a tuple with the EndTimeBeforeRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeBeforeRetention

`func (o *O11yStatefulSets) SetEndTimeBeforeRetention(v bool)`

SetEndTimeBeforeRetention sets EndTimeBeforeRetention field to given value.

### HasEndTimeBeforeRetention

`func (o *O11yStatefulSets) HasEndTimeBeforeRetention() bool`

HasEndTimeBeforeRetention returns a boolean if a field has been set.

### GetRecords

`func (o *O11yStatefulSets) GetRecords() []O11yStatefulSetRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *O11yStatefulSets) GetRecordsOk() (*[]O11yStatefulSetRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *O11yStatefulSets) SetRecords(v []O11yStatefulSetRecord)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *O11yStatefulSets) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetTotal

`func (o *O11yStatefulSets) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *O11yStatefulSets) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *O11yStatefulSets) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *O11yStatefulSets) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetType

`func (o *O11yStatefulSets) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yStatefulSets) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yStatefulSets) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *O11yStatefulSets) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *O11yStatefulSets) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *O11yStatefulSets) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetWarning

`func (o *O11yStatefulSets) GetWarning() O11yQueryWarnData`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *O11yStatefulSets) GetWarningOk() (*O11yQueryWarnData, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *O11yStatefulSets) SetWarning(v O11yQueryWarnData)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *O11yStatefulSets) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


