# CloudAgentsRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | Pointer to **string** | Caller input appended to the agent&#39;s instructions (max 128 KiB). | [optional] 

## Methods

### NewCloudAgentsRunRequest

`func NewCloudAgentsRunRequest() *CloudAgentsRunRequest`

NewCloudAgentsRunRequest instantiates a new CloudAgentsRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAgentsRunRequestWithDefaults

`func NewCloudAgentsRunRequestWithDefaults() *CloudAgentsRunRequest`

NewCloudAgentsRunRequestWithDefaults instantiates a new CloudAgentsRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *CloudAgentsRunRequest) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *CloudAgentsRunRequest) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *CloudAgentsRunRequest) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *CloudAgentsRunRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


