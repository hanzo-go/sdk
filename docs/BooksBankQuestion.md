# BooksBankQuestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connector** | Pointer to **string** | Connector names the feed the unplaceable line arrived on. With externalId it identifies both the question and the bank line it is about, so re-syncing the same deposit never asks twice. | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is when the question was raised. | [optional] 
**ExternalId** | Pointer to **string** | ExternalID is the bank&#39;s own id for the line in question. | [optional] 
**Prompt** | Pointer to **string** | Prompt is the question put to the founder in plain language — what this money was, since the books cannot place it on their own. | [optional] 
**Status** | Pointer to **string** | Status is whether the question is still open or has been answered. | [optional] 

## Methods

### NewBooksBankQuestion

`func NewBooksBankQuestion() *BooksBankQuestion`

NewBooksBankQuestion instantiates a new BooksBankQuestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooksBankQuestionWithDefaults

`func NewBooksBankQuestionWithDefaults() *BooksBankQuestion`

NewBooksBankQuestionWithDefaults instantiates a new BooksBankQuestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnector

`func (o *BooksBankQuestion) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *BooksBankQuestion) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *BooksBankQuestion) SetConnector(v string)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *BooksBankQuestion) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BooksBankQuestion) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BooksBankQuestion) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BooksBankQuestion) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BooksBankQuestion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExternalId

`func (o *BooksBankQuestion) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *BooksBankQuestion) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *BooksBankQuestion) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *BooksBankQuestion) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetPrompt

`func (o *BooksBankQuestion) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *BooksBankQuestion) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *BooksBankQuestion) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *BooksBankQuestion) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetStatus

`func (o *BooksBankQuestion) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BooksBankQuestion) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BooksBankQuestion) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BooksBankQuestion) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


