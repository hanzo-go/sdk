# AgentCodingCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to **string** | Author is the commit author&#39;s name as git recorded it; a run&#39;s own commits are authored by its harness. | [optional] 
**Date** | Pointer to **string** | Date is the author date, RFC 3339 in UTC. | [optional] 
**Message** | Pointer to **string** | Message is the commit&#39;s subject — its first line only. | [optional] 
**Sha** | Pointer to **string** | SHA is the full commit hash. | [optional] 

## Methods

### NewAgentCodingCommit

`func NewAgentCodingCommit() *AgentCodingCommit`

NewAgentCodingCommit instantiates a new AgentCodingCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentCodingCommitWithDefaults

`func NewAgentCodingCommitWithDefaults() *AgentCodingCommit`

NewAgentCodingCommitWithDefaults instantiates a new AgentCodingCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *AgentCodingCommit) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *AgentCodingCommit) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *AgentCodingCommit) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *AgentCodingCommit) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetDate

`func (o *AgentCodingCommit) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AgentCodingCommit) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AgentCodingCommit) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *AgentCodingCommit) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetMessage

`func (o *AgentCodingCommit) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AgentCodingCommit) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AgentCodingCommit) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AgentCodingCommit) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSha

`func (o *AgentCodingCommit) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *AgentCodingCommit) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *AgentCodingCommit) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *AgentCodingCommit) HasSha() bool`

HasSha returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


