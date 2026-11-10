# LegalLegalTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Body is the text/template source, markdown, referencing each declared field as {{.key}}. Rendering is pure — no clock, no I/O — so a template version and a data map always produce identical bytes. Only the single-template read carries it; the catalog listing does not. | [optional] 
**Category** | Pointer to **string** | Category is the corporate need the template serves: formation, equity, ops or sales. Formation and equity are the securities-class categories, which is what forces CounselReview. | [optional] 
**CounselReview** | Pointer to **bool** | CounselReview marks a template whose rendered documents open with CounselNotice. It can be raised but never lowered, and the engine also raises it for any formation or equity template whatever the flag says — so no org can render a securities-class document without the notice. | [optional] 
**Fields** | Pointer to [**[]LegalField**](LegalField.md) | Fields declares the merge fields Body consumes, each with its human label. Every one is REQUIRED at generation: a missing or blank value is refused rather than rendered as a silent blank into a contract. | [optional] 
**Id** | Pointer to **string** | ID is the template&#39;s stable id and the path segment that addresses it — \&quot;nda\&quot;, \&quot;msa\&quot;, \&quot;safe\&quot;. An override keeps the built-in&#39;s id, so the id names the DOCUMENT KIND rather than a particular version of it. | [optional] 
**Origin** | Pointer to **string** | Origin is \&quot;builtin\&quot; for a template the platform ships or \&quot;org\&quot; for one this org saved. It separates the catalog every tenant sees from this tenant&#39;s own. | [optional] 
**Title** | Pointer to **string** | Title is the display name, e.g. \&quot;Mutual Non-Disclosure Agreement\&quot;. A generated document inherits it. | [optional] 
**Version** | Pointer to **int64** | Version is which version of this template the caller resolved to, and what a generated document records so it stays reproducible. A built-in is version 1; an org&#39;s first override is 2 and each save increments, so an override version never collides with the built-in&#39;s. | [optional] 

## Methods

### NewLegalLegalTemplate

`func NewLegalLegalTemplate() *LegalLegalTemplate`

NewLegalLegalTemplate instantiates a new LegalLegalTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalLegalTemplateWithDefaults

`func NewLegalLegalTemplateWithDefaults() *LegalLegalTemplate`

NewLegalLegalTemplateWithDefaults instantiates a new LegalLegalTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *LegalLegalTemplate) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *LegalLegalTemplate) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *LegalLegalTemplate) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *LegalLegalTemplate) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCategory

`func (o *LegalLegalTemplate) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *LegalLegalTemplate) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *LegalLegalTemplate) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *LegalLegalTemplate) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCounselReview

`func (o *LegalLegalTemplate) GetCounselReview() bool`

GetCounselReview returns the CounselReview field if non-nil, zero value otherwise.

### GetCounselReviewOk

`func (o *LegalLegalTemplate) GetCounselReviewOk() (*bool, bool)`

GetCounselReviewOk returns a tuple with the CounselReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounselReview

`func (o *LegalLegalTemplate) SetCounselReview(v bool)`

SetCounselReview sets CounselReview field to given value.

### HasCounselReview

`func (o *LegalLegalTemplate) HasCounselReview() bool`

HasCounselReview returns a boolean if a field has been set.

### GetFields

`func (o *LegalLegalTemplate) GetFields() []LegalField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *LegalLegalTemplate) GetFieldsOk() (*[]LegalField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *LegalLegalTemplate) SetFields(v []LegalField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *LegalLegalTemplate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetId

`func (o *LegalLegalTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LegalLegalTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LegalLegalTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LegalLegalTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrigin

`func (o *LegalLegalTemplate) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *LegalLegalTemplate) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *LegalLegalTemplate) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *LegalLegalTemplate) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetTitle

`func (o *LegalLegalTemplate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LegalLegalTemplate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LegalLegalTemplate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *LegalLegalTemplate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVersion

`func (o *LegalLegalTemplate) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LegalLegalTemplate) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LegalLegalTemplate) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LegalLegalTemplate) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


