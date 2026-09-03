# O11yNodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTimeBeforeRetention** | Pointer to **bool** |  | [optional] 
**Records** | Pointer to [**[]O11yNodeRecord**](O11yNodeRecord.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **interface{}** |  | [optional] 
**Warning** | Pointer to [**O11yQueryWarnData**](O11yQueryWarnData.md) |  | [optional] 

## Methods

### NewO11yNodes

`func NewO11yNodes() *O11yNodes`

NewO11yNodes instantiates a new O11yNodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yNodesWithDefaults

`func NewO11yNodesWithDefaults() *O11yNodes`

NewO11yNodesWithDefaults instantiates a new O11yNodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTimeBeforeRetention

`func (o *O11yNodes) GetEndTimeBeforeRetention() bool`

GetEndTimeBeforeRetention returns the EndTimeBeforeRetention field if non-nil, zero value otherwise.

### GetEndTimeBeforeRetentionOk

`func (o *O11yNodes) GetEndTimeBeforeRetentionOk() (*bool, bool)`

GetEndTimeBeforeRetentionOk returns a tuple with the EndTimeBeforeRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeBeforeRetention

`func (o *O11yNodes) SetEndTimeBeforeRetention(v bool)`

SetEndTimeBeforeRetention sets EndTimeBeforeRetention field to given value.

### HasEndTimeBeforeRetention

`func (o *O11yNodes) HasEndTimeBeforeRetention() bool`

HasEndTimeBeforeRetention returns a boolean if a field has been set.

### GetRecords

`func (o *O11yNodes) GetRecords() []O11yNodeRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *O11yNodes) GetRecordsOk() (*[]O11yNodeRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *O11yNodes) SetRecords(v []O11yNodeRecord)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *O11yNodes) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetTotal

`func (o *O11yNodes) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *O11yNodes) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *O11yNodes) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *O11yNodes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetType

`func (o *O11yNodes) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yNodes) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yNodes) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *O11yNodes) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *O11yNodes) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *O11yNodes) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetWarning

`func (o *O11yNodes) GetWarning() O11yQueryWarnData`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *O11yNodes) GetWarningOk() (*O11yQueryWarnData, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *O11yNodes) SetWarning(v O11yQueryWarnData)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *O11yNodes) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


