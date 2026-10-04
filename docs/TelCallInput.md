# TelCallInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to **string** | Agent hands the answered call to a Hanzo assistant by name instead of connecting it to a person. Empty places an ordinary call. | [optional] 
**From** | Pointer to **string** | From is the number to call FROM, in E.164. It must be one this org holds. | [optional] 
**Record** | Pointer to **bool** | Record is a per-call flag rather than a product. Where a recording lands and how long it is kept is the org&#39;s retention policy, not this call&#39;s. | [optional] 
**To** | Pointer to **string** | To is the number to call, in E.164. | [optional] 
**Webhook** | Pointer to **string** | Webhook is a URL the carrier posts this call&#39;s events to as it progresses. Empty means the call&#39;s outcome is only visible by reading it back. | [optional] 

## Methods

### NewTelCallInput

`func NewTelCallInput() *TelCallInput`

NewTelCallInput instantiates a new TelCallInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelCallInputWithDefaults

`func NewTelCallInputWithDefaults() *TelCallInput`

NewTelCallInputWithDefaults instantiates a new TelCallInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *TelCallInput) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *TelCallInput) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *TelCallInput) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *TelCallInput) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetFrom

`func (o *TelCallInput) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TelCallInput) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TelCallInput) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *TelCallInput) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetRecord

`func (o *TelCallInput) GetRecord() bool`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *TelCallInput) GetRecordOk() (*bool, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *TelCallInput) SetRecord(v bool)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *TelCallInput) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetTo

`func (o *TelCallInput) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *TelCallInput) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *TelCallInput) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *TelCallInput) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetWebhook

`func (o *TelCallInput) GetWebhook() string`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *TelCallInput) GetWebhookOk() (*string, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *TelCallInput) SetWebhook(v string)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *TelCallInput) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


