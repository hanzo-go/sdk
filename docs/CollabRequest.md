# CollabRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | Pointer to **string** | DocumentID addresses the document field, as \&quot;&lt;workspaceUuid&gt;|&lt;objectClass&gt;|&lt;objectId&gt;|&lt;objectAttr&gt;\&quot; — the collaborator-client encodeDocumentId shape, from the path. | [optional] 
**Method** | Pointer to **string** | Method is the verb: createContent, updateContent or getContent. | [optional] 
**Payload** | Pointer to [**CollabPayload**](CollabPayload.md) | Payload is the verb&#39;s argument. | [optional] 

## Methods

### NewCollabRequest

`func NewCollabRequest() *CollabRequest`

NewCollabRequest instantiates a new CollabRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollabRequestWithDefaults

`func NewCollabRequestWithDefaults() *CollabRequest`

NewCollabRequestWithDefaults instantiates a new CollabRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *CollabRequest) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *CollabRequest) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *CollabRequest) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *CollabRequest) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetMethod

`func (o *CollabRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CollabRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CollabRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *CollabRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPayload

`func (o *CollabRequest) GetPayload() CollabPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CollabRequest) GetPayloadOk() (*CollabPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CollabRequest) SetPayload(v CollabPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *CollabRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


