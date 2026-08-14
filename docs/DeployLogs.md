# DeployLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | Pointer to **string** | DeploymentID is the deployment these logs belong to. | [optional] 
**Logs** | Pointer to **string** | Logs is the recorded status timeline followed by the streamed pod output, newline-separated. | [optional] 
**Source** | Pointer to **string** | Source says which pod the log body carries — &#x60;build&#x60;, &#x60;app&#x60;, or &#x60;none&#x60; when neither pod was reachable — so a console can label the pane honestly. | [optional] 

## Methods

### NewDeployLogs

`func NewDeployLogs() *DeployLogs`

NewDeployLogs instantiates a new DeployLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployLogsWithDefaults

`func NewDeployLogsWithDefaults() *DeployLogs`

NewDeployLogsWithDefaults instantiates a new DeployLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *DeployLogs) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *DeployLogs) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *DeployLogs) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *DeployLogs) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetLogs

`func (o *DeployLogs) GetLogs() string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *DeployLogs) GetLogsOk() (*string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *DeployLogs) SetLogs(v string)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *DeployLogs) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetSource

`func (o *DeployLogs) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DeployLogs) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DeployLogs) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DeployLogs) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


