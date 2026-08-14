# CallInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to **string** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**Record** | Pointer to **bool** | Record is a per-call flag rather than a product. Where a recording lands and how long it is kept is the org&#39;s retention policy, not this call&#39;s. | [optional] 
**To** | Pointer to **string** |  | [optional] 
**Webhook** | Pointer to **string** |  | [optional] 

## Methods

### NewCallInput

`func NewCallInput() *CallInput`

NewCallInput instantiates a new CallInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallInputWithDefaults

`func NewCallInputWithDefaults() *CallInput`

NewCallInputWithDefaults instantiates a new CallInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *CallInput) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *CallInput) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *CallInput) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *CallInput) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetFrom

`func (o *CallInput) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CallInput) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CallInput) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CallInput) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetRecord

`func (o *CallInput) GetRecord() bool`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *CallInput) GetRecordOk() (*bool, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *CallInput) SetRecord(v bool)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *CallInput) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetTo

`func (o *CallInput) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CallInput) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CallInput) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *CallInput) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetWebhook

`func (o *CallInput) GetWebhook() string`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *CallInput) GetWebhookOk() (*string, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *CallInput) SetWebhook(v string)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *CallInput) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


