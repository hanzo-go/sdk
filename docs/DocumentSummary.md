# DocumentSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Category is the template&#39;s category: formation, equity, ops or sales. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is when the document was generated, in unix seconds. | [optional] 
**EsignProvider** | Pointer to **string** | EsignProvider names the e-signature provider handling it, absent until a signature has been requested. | [optional] 
**Id** | Pointer to **string** | ID is the document&#39;s server-minted handle, \&quot;doc_\&quot;-prefixed. | [optional] 
**SignedAt** | Pointer to **int32** | SignedAt is when the provider reported completion, in unix seconds. Absent until then. | [optional] 
**Status** | Pointer to **string** | Status is the lifecycle state: draft, out_for_signature, signed or voided. There is deliberately no \&quot;legally valid\&quot; state — that is counsel&#39;s determination, not the platform&#39;s. | [optional] 
**TemplateId** | Pointer to **string** | TemplateID is the template it was rendered from. | [optional] 
**TemplateVersion** | Pointer to **int32** | TemplateVersion is WHICH version of that template rendered it, so the document is reproducible and auditable. | [optional] 
**Title** | Pointer to **string** | Title is the document&#39;s title, inherited from the template. | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is when it last changed, in unix seconds. | [optional] 

## Methods

### NewDocumentSummary

`func NewDocumentSummary() *DocumentSummary`

NewDocumentSummary instantiates a new DocumentSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentSummaryWithDefaults

`func NewDocumentSummaryWithDefaults() *DocumentSummary`

NewDocumentSummaryWithDefaults instantiates a new DocumentSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *DocumentSummary) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *DocumentSummary) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *DocumentSummary) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *DocumentSummary) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DocumentSummary) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DocumentSummary) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DocumentSummary) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DocumentSummary) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEsignProvider

`func (o *DocumentSummary) GetEsignProvider() string`

GetEsignProvider returns the EsignProvider field if non-nil, zero value otherwise.

### GetEsignProviderOk

`func (o *DocumentSummary) GetEsignProviderOk() (*string, bool)`

GetEsignProviderOk returns a tuple with the EsignProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsignProvider

`func (o *DocumentSummary) SetEsignProvider(v string)`

SetEsignProvider sets EsignProvider field to given value.

### HasEsignProvider

`func (o *DocumentSummary) HasEsignProvider() bool`

HasEsignProvider returns a boolean if a field has been set.

### GetId

`func (o *DocumentSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DocumentSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSignedAt

`func (o *DocumentSummary) GetSignedAt() int32`

GetSignedAt returns the SignedAt field if non-nil, zero value otherwise.

### GetSignedAtOk

`func (o *DocumentSummary) GetSignedAtOk() (*int32, bool)`

GetSignedAtOk returns a tuple with the SignedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedAt

`func (o *DocumentSummary) SetSignedAt(v int32)`

SetSignedAt sets SignedAt field to given value.

### HasSignedAt

`func (o *DocumentSummary) HasSignedAt() bool`

HasSignedAt returns a boolean if a field has been set.

### GetStatus

`func (o *DocumentSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DocumentSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DocumentSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DocumentSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTemplateId

`func (o *DocumentSummary) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *DocumentSummary) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *DocumentSummary) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *DocumentSummary) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTemplateVersion

`func (o *DocumentSummary) GetTemplateVersion() int32`

GetTemplateVersion returns the TemplateVersion field if non-nil, zero value otherwise.

### GetTemplateVersionOk

`func (o *DocumentSummary) GetTemplateVersionOk() (*int32, bool)`

GetTemplateVersionOk returns a tuple with the TemplateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVersion

`func (o *DocumentSummary) SetTemplateVersion(v int32)`

SetTemplateVersion sets TemplateVersion field to given value.

### HasTemplateVersion

`func (o *DocumentSummary) HasTemplateVersion() bool`

HasTemplateVersion returns a boolean if a field has been set.

### GetTitle

`func (o *DocumentSummary) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DocumentSummary) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DocumentSummary) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DocumentSummary) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DocumentSummary) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DocumentSummary) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DocumentSummary) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DocumentSummary) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


