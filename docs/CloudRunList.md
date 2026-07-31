# CloudRunList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Runs** | Pointer to [**[]CloudAgentRunView**](CloudAgentRunView.md) | Runs is the agent&#39;s executions, newest first. | [optional] 

## Methods

### NewCloudRunList

`func NewCloudRunList() *CloudRunList`

NewCloudRunList instantiates a new CloudRunList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRunListWithDefaults

`func NewCloudRunListWithDefaults() *CloudRunList`

NewCloudRunListWithDefaults instantiates a new CloudRunList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuns

`func (o *CloudRunList) GetRuns() []CloudAgentRunView`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *CloudRunList) GetRunsOk() (*[]CloudAgentRunView, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *CloudRunList) SetRuns(v []CloudAgentRunView)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *CloudRunList) HasRuns() bool`

HasRuns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


