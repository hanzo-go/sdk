# DocumentView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Body is the rendered document. It is sealed at rest and returned only to the owning org. When the template is counsel-review it opens with the counsel notice, which the engine prepends and no caller can suppress. | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** | ContentType is the rendered body&#39;s media type — text/markdown. | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**EsignProvider** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**SignedAt** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TemplateId** | Pointer to **string** |  | [optional] 
**TemplateVersion** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **int32** |  | [optional] 

## Methods

### NewDocumentView

`func NewDocumentView() *DocumentView`

NewDocumentView instantiates a new DocumentView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentViewWithDefaults

`func NewDocumentViewWithDefaults() *DocumentView`

NewDocumentViewWithDefaults instantiates a new DocumentView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *DocumentView) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *DocumentView) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *DocumentView) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *DocumentView) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCategory

`func (o *DocumentView) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *DocumentView) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *DocumentView) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *DocumentView) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetContentType

`func (o *DocumentView) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *DocumentView) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *DocumentView) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *DocumentView) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DocumentView) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DocumentView) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DocumentView) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DocumentView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEsignProvider

`func (o *DocumentView) GetEsignProvider() string`

GetEsignProvider returns the EsignProvider field if non-nil, zero value otherwise.

### GetEsignProviderOk

`func (o *DocumentView) GetEsignProviderOk() (*string, bool)`

GetEsignProviderOk returns a tuple with the EsignProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsignProvider

`func (o *DocumentView) SetEsignProvider(v string)`

SetEsignProvider sets EsignProvider field to given value.

### HasEsignProvider

`func (o *DocumentView) HasEsignProvider() bool`

HasEsignProvider returns a boolean if a field has been set.

### GetId

`func (o *DocumentView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DocumentView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSignedAt

`func (o *DocumentView) GetSignedAt() int32`

GetSignedAt returns the SignedAt field if non-nil, zero value otherwise.

### GetSignedAtOk

`func (o *DocumentView) GetSignedAtOk() (*int32, bool)`

GetSignedAtOk returns a tuple with the SignedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedAt

`func (o *DocumentView) SetSignedAt(v int32)`

SetSignedAt sets SignedAt field to given value.

### HasSignedAt

`func (o *DocumentView) HasSignedAt() bool`

HasSignedAt returns a boolean if a field has been set.

### GetStatus

`func (o *DocumentView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DocumentView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DocumentView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DocumentView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTemplateId

`func (o *DocumentView) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *DocumentView) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *DocumentView) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *DocumentView) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTemplateVersion

`func (o *DocumentView) GetTemplateVersion() int32`

GetTemplateVersion returns the TemplateVersion field if non-nil, zero value otherwise.

### GetTemplateVersionOk

`func (o *DocumentView) GetTemplateVersionOk() (*int32, bool)`

GetTemplateVersionOk returns a tuple with the TemplateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVersion

`func (o *DocumentView) SetTemplateVersion(v int32)`

SetTemplateVersion sets TemplateVersion field to given value.

### HasTemplateVersion

`func (o *DocumentView) HasTemplateVersion() bool`

HasTemplateVersion returns a boolean if a field has been set.

### GetTitle

`func (o *DocumentView) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DocumentView) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DocumentView) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DocumentView) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DocumentView) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DocumentView) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DocumentView) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DocumentView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


