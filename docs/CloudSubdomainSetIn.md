# CloudSubdomainSetIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enabled publishes the script on &lt;script&gt;.&lt;subdomain&gt;.workers.dev when true, and withdraws it when false. | [optional] 
**Script** | Pointer to **string** | Script is the Worker script name, from the path. | [optional] 

## Methods

### NewCloudSubdomainSetIn

`func NewCloudSubdomainSetIn() *CloudSubdomainSetIn`

NewCloudSubdomainSetIn instantiates a new CloudSubdomainSetIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSubdomainSetInWithDefaults

`func NewCloudSubdomainSetInWithDefaults() *CloudSubdomainSetIn`

NewCloudSubdomainSetInWithDefaults instantiates a new CloudSubdomainSetIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *CloudSubdomainSetIn) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CloudSubdomainSetIn) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CloudSubdomainSetIn) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CloudSubdomainSetIn) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetScript

`func (o *CloudSubdomainSetIn) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *CloudSubdomainSetIn) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *CloudSubdomainSetIn) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *CloudSubdomainSetIn) HasScript() bool`

HasScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


