# O11yHostListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterNames** | Pointer to **[]string** |  | [optional] 
**EndTimeBeforeRetention** | Pointer to **bool** |  | [optional] 
**IsSendingK8SAgentMetrics** | Pointer to **bool** |  | [optional] 
**NodeNames** | Pointer to **[]string** |  | [optional] 
**Records** | Pointer to [**[]O11yHostListRecord**](O11yHostListRecord.md) |  | [optional] 
**SentAnyHostMetricsData** | Pointer to **bool** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yHostListResponse

`func NewO11yHostListResponse() *O11yHostListResponse`

NewO11yHostListResponse instantiates a new O11yHostListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yHostListResponseWithDefaults

`func NewO11yHostListResponseWithDefaults() *O11yHostListResponse`

NewO11yHostListResponseWithDefaults instantiates a new O11yHostListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterNames

`func (o *O11yHostListResponse) GetClusterNames() []string`

GetClusterNames returns the ClusterNames field if non-nil, zero value otherwise.

### GetClusterNamesOk

`func (o *O11yHostListResponse) GetClusterNamesOk() (*[]string, bool)`

GetClusterNamesOk returns a tuple with the ClusterNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNames

`func (o *O11yHostListResponse) SetClusterNames(v []string)`

SetClusterNames sets ClusterNames field to given value.

### HasClusterNames

`func (o *O11yHostListResponse) HasClusterNames() bool`

HasClusterNames returns a boolean if a field has been set.

### GetEndTimeBeforeRetention

`func (o *O11yHostListResponse) GetEndTimeBeforeRetention() bool`

GetEndTimeBeforeRetention returns the EndTimeBeforeRetention field if non-nil, zero value otherwise.

### GetEndTimeBeforeRetentionOk

`func (o *O11yHostListResponse) GetEndTimeBeforeRetentionOk() (*bool, bool)`

GetEndTimeBeforeRetentionOk returns a tuple with the EndTimeBeforeRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeBeforeRetention

`func (o *O11yHostListResponse) SetEndTimeBeforeRetention(v bool)`

SetEndTimeBeforeRetention sets EndTimeBeforeRetention field to given value.

### HasEndTimeBeforeRetention

`func (o *O11yHostListResponse) HasEndTimeBeforeRetention() bool`

HasEndTimeBeforeRetention returns a boolean if a field has been set.

### GetIsSendingK8SAgentMetrics

`func (o *O11yHostListResponse) GetIsSendingK8SAgentMetrics() bool`

GetIsSendingK8SAgentMetrics returns the IsSendingK8SAgentMetrics field if non-nil, zero value otherwise.

### GetIsSendingK8SAgentMetricsOk

`func (o *O11yHostListResponse) GetIsSendingK8SAgentMetricsOk() (*bool, bool)`

GetIsSendingK8SAgentMetricsOk returns a tuple with the IsSendingK8SAgentMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSendingK8SAgentMetrics

`func (o *O11yHostListResponse) SetIsSendingK8SAgentMetrics(v bool)`

SetIsSendingK8SAgentMetrics sets IsSendingK8SAgentMetrics field to given value.

### HasIsSendingK8SAgentMetrics

`func (o *O11yHostListResponse) HasIsSendingK8SAgentMetrics() bool`

HasIsSendingK8SAgentMetrics returns a boolean if a field has been set.

### GetNodeNames

`func (o *O11yHostListResponse) GetNodeNames() []string`

GetNodeNames returns the NodeNames field if non-nil, zero value otherwise.

### GetNodeNamesOk

`func (o *O11yHostListResponse) GetNodeNamesOk() (*[]string, bool)`

GetNodeNamesOk returns a tuple with the NodeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeNames

`func (o *O11yHostListResponse) SetNodeNames(v []string)`

SetNodeNames sets NodeNames field to given value.

### HasNodeNames

`func (o *O11yHostListResponse) HasNodeNames() bool`

HasNodeNames returns a boolean if a field has been set.

### GetRecords

`func (o *O11yHostListResponse) GetRecords() []O11yHostListRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *O11yHostListResponse) GetRecordsOk() (*[]O11yHostListRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *O11yHostListResponse) SetRecords(v []O11yHostListRecord)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *O11yHostListResponse) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetSentAnyHostMetricsData

`func (o *O11yHostListResponse) GetSentAnyHostMetricsData() bool`

GetSentAnyHostMetricsData returns the SentAnyHostMetricsData field if non-nil, zero value otherwise.

### GetSentAnyHostMetricsDataOk

`func (o *O11yHostListResponse) GetSentAnyHostMetricsDataOk() (*bool, bool)`

GetSentAnyHostMetricsDataOk returns a tuple with the SentAnyHostMetricsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAnyHostMetricsData

`func (o *O11yHostListResponse) SetSentAnyHostMetricsData(v bool)`

SetSentAnyHostMetricsData sets SentAnyHostMetricsData field to given value.

### HasSentAnyHostMetricsData

`func (o *O11yHostListResponse) HasSentAnyHostMetricsData() bool`

HasSentAnyHostMetricsData returns a boolean if a field has been set.

### GetTotal

`func (o *O11yHostListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *O11yHostListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *O11yHostListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *O11yHostListResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetType

`func (o *O11yHostListResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yHostListResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yHostListResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yHostListResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


