# VisorBotLaunchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Required for a real launch (not dryRun) | [optional] 
**Agent** | Pointer to **string** | The cloud Agent the bot runs (defaults to the bot name) | [optional] 
**Size** | Pointer to **string** | Machine size slug (or use instanceType) | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**BotVersion** | Pointer to **string** |  | [optional] 
**DryRun** | Pointer to **bool** |  | [optional] 

## Methods

### NewVisorBotLaunchRequest

`func NewVisorBotLaunchRequest() *VisorBotLaunchRequest`

NewVisorBotLaunchRequest instantiates a new VisorBotLaunchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisorBotLaunchRequestWithDefaults

`func NewVisorBotLaunchRequestWithDefaults() *VisorBotLaunchRequest`

NewVisorBotLaunchRequestWithDefaults instantiates a new VisorBotLaunchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VisorBotLaunchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VisorBotLaunchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VisorBotLaunchRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VisorBotLaunchRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAgent

`func (o *VisorBotLaunchRequest) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *VisorBotLaunchRequest) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *VisorBotLaunchRequest) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *VisorBotLaunchRequest) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetSize

`func (o *VisorBotLaunchRequest) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VisorBotLaunchRequest) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VisorBotLaunchRequest) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *VisorBotLaunchRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetInstanceType

`func (o *VisorBotLaunchRequest) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *VisorBotLaunchRequest) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *VisorBotLaunchRequest) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *VisorBotLaunchRequest) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetRegion

`func (o *VisorBotLaunchRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *VisorBotLaunchRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *VisorBotLaunchRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *VisorBotLaunchRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetBotVersion

`func (o *VisorBotLaunchRequest) GetBotVersion() string`

GetBotVersion returns the BotVersion field if non-nil, zero value otherwise.

### GetBotVersionOk

`func (o *VisorBotLaunchRequest) GetBotVersionOk() (*string, bool)`

GetBotVersionOk returns a tuple with the BotVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotVersion

`func (o *VisorBotLaunchRequest) SetBotVersion(v string)`

SetBotVersion sets BotVersion field to given value.

### HasBotVersion

`func (o *VisorBotLaunchRequest) HasBotVersion() bool`

HasBotVersion returns a boolean if a field has been set.

### GetDryRun

`func (o *VisorBotLaunchRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *VisorBotLaunchRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *VisorBotLaunchRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *VisorBotLaunchRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


