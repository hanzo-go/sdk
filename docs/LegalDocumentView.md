# LegalDocumentView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Body is the rendered document. It is sealed at rest and returned only to the owning org. When the template is counsel-review it opens with the counsel notice, which the engine prepends and no caller can suppress. | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** | ContentType is the rendered body&#39;s media type — text/markdown. | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**EsignProvider** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**SignedAt** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TemplateId** | Pointer to **string** |  | [optional] 
**TemplateVersion** | Pointer to **int64** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewLegalDocumentView

`func NewLegalDocumentView() *LegalDocumentView`

NewLegalDocumentView instantiates a new LegalDocumentView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalDocumentViewWithDefaults

`func NewLegalDocumentViewWithDefaults() *LegalDocumentView`

NewLegalDocumentViewWithDefaults instantiates a new LegalDocumentView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *LegalDocumentView) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *LegalDocumentView) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *LegalDocumentView) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *LegalDocumentView) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCategory

`func (o *LegalDocumentView) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *LegalDocumentView) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *LegalDocumentView) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *LegalDocumentView) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetContentType

`func (o *LegalDocumentView) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *LegalDocumentView) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *LegalDocumentView) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *LegalDocumentView) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LegalDocumentView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LegalDocumentView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LegalDocumentView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LegalDocumentView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEsignProvider

`func (o *LegalDocumentView) GetEsignProvider() string`

GetEsignProvider returns the EsignProvider field if non-nil, zero value otherwise.

### GetEsignProviderOk

`func (o *LegalDocumentView) GetEsignProviderOk() (*string, bool)`

GetEsignProviderOk returns a tuple with the EsignProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsignProvider

`func (o *LegalDocumentView) SetEsignProvider(v string)`

SetEsignProvider sets EsignProvider field to given value.

### HasEsignProvider

`func (o *LegalDocumentView) HasEsignProvider() bool`

HasEsignProvider returns a boolean if a field has been set.

### GetId

`func (o *LegalDocumentView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LegalDocumentView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LegalDocumentView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LegalDocumentView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSignedAt

`func (o *LegalDocumentView) GetSignedAt() int64`

GetSignedAt returns the SignedAt field if non-nil, zero value otherwise.

### GetSignedAtOk

`func (o *LegalDocumentView) GetSignedAtOk() (*int64, bool)`

GetSignedAtOk returns a tuple with the SignedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedAt

`func (o *LegalDocumentView) SetSignedAt(v int64)`

SetSignedAt sets SignedAt field to given value.

### HasSignedAt

`func (o *LegalDocumentView) HasSignedAt() bool`

HasSignedAt returns a boolean if a field has been set.

### GetStatus

`func (o *LegalDocumentView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LegalDocumentView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LegalDocumentView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LegalDocumentView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTemplateId

`func (o *LegalDocumentView) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *LegalDocumentView) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *LegalDocumentView) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *LegalDocumentView) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTemplateVersion

`func (o *LegalDocumentView) GetTemplateVersion() int64`

GetTemplateVersion returns the TemplateVersion field if non-nil, zero value otherwise.

### GetTemplateVersionOk

`func (o *LegalDocumentView) GetTemplateVersionOk() (*int64, bool)`

GetTemplateVersionOk returns a tuple with the TemplateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVersion

`func (o *LegalDocumentView) SetTemplateVersion(v int64)`

SetTemplateVersion sets TemplateVersion field to given value.

### HasTemplateVersion

`func (o *LegalDocumentView) HasTemplateVersion() bool`

HasTemplateVersion returns a boolean if a field has been set.

### GetTitle

`func (o *LegalDocumentView) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LegalDocumentView) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LegalDocumentView) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *LegalDocumentView) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LegalDocumentView) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LegalDocumentView) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LegalDocumentView) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LegalDocumentView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


