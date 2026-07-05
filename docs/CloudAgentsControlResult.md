# CloudAgentsControlResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **string** |  | [optional] 
**Event** | Pointer to [**CloudAgentsEvent**](CloudAgentsEvent.md) |  | [optional] 
**Forwarded** | Pointer to **bool** | True when the command was forwarded to the hanzoai/tasks engine. | [optional] 

## Methods

### NewCloudAgentsControlResult

`func NewCloudAgentsControlResult() *CloudAgentsControlResult`

NewCloudAgentsControlResult instantiates a new CloudAgentsControlResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAgentsControlResultWithDefaults

`func NewCloudAgentsControlResultWithDefaults() *CloudAgentsControlResult`

NewCloudAgentsControlResultWithDefaults instantiates a new CloudAgentsControlResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *CloudAgentsControlResult) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *CloudAgentsControlResult) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *CloudAgentsControlResult) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *CloudAgentsControlResult) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetEvent

`func (o *CloudAgentsControlResult) GetEvent() CloudAgentsEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CloudAgentsControlResult) GetEventOk() (*CloudAgentsEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CloudAgentsControlResult) SetEvent(v CloudAgentsEvent)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *CloudAgentsControlResult) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetForwarded

`func (o *CloudAgentsControlResult) GetForwarded() bool`

GetForwarded returns the Forwarded field if non-nil, zero value otherwise.

### GetForwardedOk

`func (o *CloudAgentsControlResult) GetForwardedOk() (*bool, bool)`

GetForwardedOk returns a tuple with the Forwarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwarded

`func (o *CloudAgentsControlResult) SetForwarded(v bool)`

SetForwarded sets Forwarded field to given value.

### HasForwarded

`func (o *CloudAgentsControlResult) HasForwarded() bool`

HasForwarded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


