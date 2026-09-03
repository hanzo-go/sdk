# O11yAgentReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**TimestampMillis** | Pointer to **int64** |  | [optional] 

## Methods

### NewO11yAgentReport

`func NewO11yAgentReport() *O11yAgentReport`

NewO11yAgentReport instantiates a new O11yAgentReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAgentReportWithDefaults

`func NewO11yAgentReportWithDefaults() *O11yAgentReport`

NewO11yAgentReportWithDefaults instantiates a new O11yAgentReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yAgentReport) GetData() map[string]map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yAgentReport) GetDataOk() (*map[string]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yAgentReport) SetData(v map[string]map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *O11yAgentReport) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTimestampMillis

`func (o *O11yAgentReport) GetTimestampMillis() int64`

GetTimestampMillis returns the TimestampMillis field if non-nil, zero value otherwise.

### GetTimestampMillisOk

`func (o *O11yAgentReport) GetTimestampMillisOk() (*int64, bool)`

GetTimestampMillisOk returns a tuple with the TimestampMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampMillis

`func (o *O11yAgentReport) SetTimestampMillis(v int64)`

SetTimestampMillis sets TimestampMillis field to given value.

### HasTimestampMillis

`func (o *O11yAgentReport) HasTimestampMillis() bool`

HasTimestampMillis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


