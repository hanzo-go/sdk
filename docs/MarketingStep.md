# MarketingStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Body is the message text. Required. The signed one-click unsubscribe link is appended to it at send time. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is unix seconds, server-assigned. | [optional] 
**DelaySeconds** | Pointer to **int64** | DelaySeconds is how long after the previous step this one sends (after enrollment, for step 0). | [optional] 
**Id** | Pointer to **string** | ID is the server-assigned step id (\&quot;step_\&quot; + 128 random bits). | [optional] 
**Idx** | Pointer to **int64** | Idx is the step&#39;s 0-based position, assigned by appending: a new step always lands after the last one. | [optional] 
**SequenceId** | Pointer to **string** | SequenceID is the sequence this step belongs to. | [optional] 
**Subject** | Pointer to **string** | Subject is the email subject line, capped at 1024 bytes. | [optional] 

## Methods

### NewMarketingStep

`func NewMarketingStep() *MarketingStep`

NewMarketingStep instantiates a new MarketingStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingStepWithDefaults

`func NewMarketingStepWithDefaults() *MarketingStep`

NewMarketingStepWithDefaults instantiates a new MarketingStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *MarketingStep) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *MarketingStep) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *MarketingStep) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *MarketingStep) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MarketingStep) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MarketingStep) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MarketingStep) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MarketingStep) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDelaySeconds

`func (o *MarketingStep) GetDelaySeconds() int64`

GetDelaySeconds returns the DelaySeconds field if non-nil, zero value otherwise.

### GetDelaySecondsOk

`func (o *MarketingStep) GetDelaySecondsOk() (*int64, bool)`

GetDelaySecondsOk returns a tuple with the DelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelaySeconds

`func (o *MarketingStep) SetDelaySeconds(v int64)`

SetDelaySeconds sets DelaySeconds field to given value.

### HasDelaySeconds

`func (o *MarketingStep) HasDelaySeconds() bool`

HasDelaySeconds returns a boolean if a field has been set.

### GetId

`func (o *MarketingStep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketingStep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketingStep) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketingStep) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdx

`func (o *MarketingStep) GetIdx() int64`

GetIdx returns the Idx field if non-nil, zero value otherwise.

### GetIdxOk

`func (o *MarketingStep) GetIdxOk() (*int64, bool)`

GetIdxOk returns a tuple with the Idx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdx

`func (o *MarketingStep) SetIdx(v int64)`

SetIdx sets Idx field to given value.

### HasIdx

`func (o *MarketingStep) HasIdx() bool`

HasIdx returns a boolean if a field has been set.

### GetSequenceId

`func (o *MarketingStep) GetSequenceId() string`

GetSequenceId returns the SequenceId field if non-nil, zero value otherwise.

### GetSequenceIdOk

`func (o *MarketingStep) GetSequenceIdOk() (*string, bool)`

GetSequenceIdOk returns a tuple with the SequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceId

`func (o *MarketingStep) SetSequenceId(v string)`

SetSequenceId sets SequenceId field to given value.

### HasSequenceId

`func (o *MarketingStep) HasSequenceId() bool`

HasSequenceId returns a boolean if a field has been set.

### GetSubject

`func (o *MarketingStep) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MarketingStep) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MarketingStep) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *MarketingStep) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


