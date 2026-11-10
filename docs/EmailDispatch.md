# EmailDispatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]EmailMessageRecord**](EmailMessageRecord.md) | Messages is one message per recipient that was sent, each with its own unsubscribe link; a message the provider refused is here with status failed. | [optional] 
**Refused** | Pointer to [**[]EmailRefused**](EmailRefused.md) | Refused is every recipient the suppression check refused. Nothing was sent to them. | [optional] 

## Methods

### NewEmailDispatch

`func NewEmailDispatch() *EmailDispatch`

NewEmailDispatch instantiates a new EmailDispatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailDispatchWithDefaults

`func NewEmailDispatchWithDefaults() *EmailDispatch`

NewEmailDispatchWithDefaults instantiates a new EmailDispatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *EmailDispatch) GetMessages() []EmailMessageRecord`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *EmailDispatch) GetMessagesOk() (*[]EmailMessageRecord, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *EmailDispatch) SetMessages(v []EmailMessageRecord)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *EmailDispatch) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetRefused

`func (o *EmailDispatch) GetRefused() []EmailRefused`

GetRefused returns the Refused field if non-nil, zero value otherwise.

### GetRefusedOk

`func (o *EmailDispatch) GetRefusedOk() (*[]EmailRefused, bool)`

GetRefusedOk returns a tuple with the Refused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefused

`func (o *EmailDispatch) SetRefused(v []EmailRefused)`

SetRefused sets Refused field to given value.

### HasRefused

`func (o *EmailDispatch) HasRefused() bool`

HasRefused returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


