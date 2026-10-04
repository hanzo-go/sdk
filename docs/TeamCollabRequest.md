# TeamCollabRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | Pointer to **string** | DocumentID addresses the document field, as \&quot;&lt;spaceUuid&gt;|&lt;objectClass&gt;|&lt;objectId&gt;|&lt;objectAttr&gt;\&quot; — the collaborator-client encodeDocumentId shape, from the path. | [optional] 
**Method** | Pointer to **string** | Method is the verb: createContent, updateContent or getContent. | [optional] 
**Payload** | Pointer to [**TeamCollabPayload**](TeamCollabPayload.md) | Payload is the verb&#39;s argument. | [optional] 

## Methods

### NewTeamCollabRequest

`func NewTeamCollabRequest() *TeamCollabRequest`

NewTeamCollabRequest instantiates a new TeamCollabRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamCollabRequestWithDefaults

`func NewTeamCollabRequestWithDefaults() *TeamCollabRequest`

NewTeamCollabRequestWithDefaults instantiates a new TeamCollabRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *TeamCollabRequest) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *TeamCollabRequest) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *TeamCollabRequest) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *TeamCollabRequest) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetMethod

`func (o *TeamCollabRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *TeamCollabRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *TeamCollabRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *TeamCollabRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPayload

`func (o *TeamCollabRequest) GetPayload() TeamCollabPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *TeamCollabRequest) GetPayloadOk() (*TeamCollabPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *TeamCollabRequest) SetPayload(v TeamCollabPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *TeamCollabRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


