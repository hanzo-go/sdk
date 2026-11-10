# EmailListMessagesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]EmailMessageRecord**](EmailMessageRecord.md) | Messages is the page. | [optional] 
**Next** | Pointer to **string** | Next is the cursor of the page after this one; empty on the last page. | [optional] 

## Methods

### NewEmailListMessagesResult

`func NewEmailListMessagesResult() *EmailListMessagesResult`

NewEmailListMessagesResult instantiates a new EmailListMessagesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailListMessagesResultWithDefaults

`func NewEmailListMessagesResultWithDefaults() *EmailListMessagesResult`

NewEmailListMessagesResultWithDefaults instantiates a new EmailListMessagesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *EmailListMessagesResult) GetMessages() []EmailMessageRecord`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *EmailListMessagesResult) GetMessagesOk() (*[]EmailMessageRecord, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *EmailListMessagesResult) SetMessages(v []EmailMessageRecord)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *EmailListMessagesResult) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetNext

`func (o *EmailListMessagesResult) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *EmailListMessagesResult) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *EmailListMessagesResult) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *EmailListMessagesResult) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


