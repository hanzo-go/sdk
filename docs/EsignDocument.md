# EsignDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | Pointer to **int32** | CompletedAt is when the document sealed, in unix milliseconds; null until it does. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is when the document was uploaded, in unix milliseconds. | [optional] 
**ExternalId** | Pointer to **string** | ExternalID is the caller&#39;s own identifier for this document, echoed back as it was given; null when none was. | [optional] 
**Fields** | Pointer to [**[]EsignField**](EsignField.md) | Fields is every field on the document, ordered by page and then by when it was placed. | [optional] 
**Id** | Pointer to **string** | ID is the document id. | [optional] 
**Message** | Pointer to **string** | Message is the covering message stored with the document; null when none was given. Nothing in this surface sends it. | [optional] 
**Recipients** | Pointer to [**[]EsignRecipient**](EsignRecipient.md) | Recipients is everyone on the document, ordered by signing order and then by when they were added — which is also the order a SEQUENTIAL document enforces. | [optional] 
**SigningOrder** | Pointer to **string** | SigningOrder is PARALLEL or SEQUENTIAL, fixed when the document was created. | [optional] 
**Source** | Pointer to **string** | Source is how the document came to exist. It is DOCUMENT for everything this surface creates. | [optional] 
**Status** | Pointer to **string** | Status is DRAFT while recipients and fields may still be added, PENDING once it has gone out, then COMPLETED when every signer has finished or REJECTED if any one of them declined. | [optional] 
**Subject** | Pointer to **string** | Subject is the covering subject line stored with the document; null when none was given. | [optional] 
**Title** | Pointer to **string** | Title is the document&#39;s name, and the stem of the download filename. | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is when the document last changed, in unix milliseconds. | [optional] 

## Methods

### NewEsignDocument

`func NewEsignDocument() *EsignDocument`

NewEsignDocument instantiates a new EsignDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsignDocumentWithDefaults

`func NewEsignDocumentWithDefaults() *EsignDocument`

NewEsignDocumentWithDefaults instantiates a new EsignDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *EsignDocument) GetCompletedAt() int32`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *EsignDocument) GetCompletedAtOk() (*int32, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *EsignDocument) SetCompletedAt(v int32)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *EsignDocument) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EsignDocument) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EsignDocument) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EsignDocument) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EsignDocument) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExternalId

`func (o *EsignDocument) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *EsignDocument) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *EsignDocument) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *EsignDocument) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetFields

`func (o *EsignDocument) GetFields() []EsignField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *EsignDocument) GetFieldsOk() (*[]EsignField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *EsignDocument) SetFields(v []EsignField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *EsignDocument) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetId

`func (o *EsignDocument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsignDocument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsignDocument) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EsignDocument) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *EsignDocument) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EsignDocument) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EsignDocument) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EsignDocument) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRecipients

`func (o *EsignDocument) GetRecipients() []EsignRecipient`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *EsignDocument) GetRecipientsOk() (*[]EsignRecipient, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *EsignDocument) SetRecipients(v []EsignRecipient)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *EsignDocument) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetSigningOrder

`func (o *EsignDocument) GetSigningOrder() string`

GetSigningOrder returns the SigningOrder field if non-nil, zero value otherwise.

### GetSigningOrderOk

`func (o *EsignDocument) GetSigningOrderOk() (*string, bool)`

GetSigningOrderOk returns a tuple with the SigningOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningOrder

`func (o *EsignDocument) SetSigningOrder(v string)`

SetSigningOrder sets SigningOrder field to given value.

### HasSigningOrder

`func (o *EsignDocument) HasSigningOrder() bool`

HasSigningOrder returns a boolean if a field has been set.

### GetSource

`func (o *EsignDocument) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EsignDocument) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EsignDocument) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *EsignDocument) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *EsignDocument) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EsignDocument) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EsignDocument) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EsignDocument) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubject

`func (o *EsignDocument) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EsignDocument) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EsignDocument) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EsignDocument) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTitle

`func (o *EsignDocument) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EsignDocument) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EsignDocument) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EsignDocument) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EsignDocument) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EsignDocument) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EsignDocument) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EsignDocument) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


