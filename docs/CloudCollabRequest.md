# CloudCollabRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | Pointer to **string** | DocumentID addresses the document field, as \&quot;&lt;workspaceUuid&gt;|&lt;objectClass&gt;|&lt;objectId&gt;|&lt;objectAttr&gt;\&quot; — the collaborator-client encodeDocumentId shape, from the path. | [optional] 
**Method** | Pointer to **string** | Method is the verb: createContent, updateContent or getContent. | [optional] 
**Payload** | Pointer to [**CloudCollabPayload**](CloudCollabPayload.md) | Payload is the verb&#39;s argument. | [optional] 

## Methods

### NewCloudCollabRequest

`func NewCloudCollabRequest() *CloudCollabRequest`

NewCloudCollabRequest instantiates a new CloudCollabRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCollabRequestWithDefaults

`func NewCloudCollabRequestWithDefaults() *CloudCollabRequest`

NewCloudCollabRequestWithDefaults instantiates a new CloudCollabRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *CloudCollabRequest) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *CloudCollabRequest) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *CloudCollabRequest) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *CloudCollabRequest) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetMethod

`func (o *CloudCollabRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CloudCollabRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CloudCollabRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *CloudCollabRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPayload

`func (o *CloudCollabRequest) GetPayload() CloudCollabPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CloudCollabRequest) GetPayloadOk() (*CloudCollabPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CloudCollabRequest) SetPayload(v CloudCollabPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *CloudCollabRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


