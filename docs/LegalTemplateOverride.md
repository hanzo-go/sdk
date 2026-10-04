# LegalTemplateOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Body is the text/template source. Required. Every {{.key}} it references must be declared in Fields, or the save is refused rather than rendering a blank into a contract later. | [optional] 
**Category** | Pointer to **string** | Category groups the template: formation, equity, ops or sales. Optional when overriding a built-in, which supplies its own. | [optional] 
**CounselReview** | Pointer to **bool** | CounselReview marks a template whose documents must carry the counsel notice. It can be raised but never lowered: a formation or equity template is always counsel-review, and an override of a counsel-review built-in stays one. | [optional] 
**Fields** | Pointer to [**[]LegalField**](LegalField.md) | Fields declares the merge fields the body consumes. Every declared field is REQUIRED at generation — the engine fails closed on a missing one. | [optional] 
**Id** | Pointer to **string** | ID is the template to override, from the path. Overriding a built-in id inherits that built-in&#39;s category, title and counsel-review posture. | [optional] 
**Title** | Pointer to **string** | Title is the template&#39;s display name. Required unless a built-in supplies it. | [optional] 

## Methods

### NewLegalTemplateOverride

`func NewLegalTemplateOverride() *LegalTemplateOverride`

NewLegalTemplateOverride instantiates a new LegalTemplateOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalTemplateOverrideWithDefaults

`func NewLegalTemplateOverrideWithDefaults() *LegalTemplateOverride`

NewLegalTemplateOverrideWithDefaults instantiates a new LegalTemplateOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *LegalTemplateOverride) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *LegalTemplateOverride) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *LegalTemplateOverride) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *LegalTemplateOverride) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCategory

`func (o *LegalTemplateOverride) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *LegalTemplateOverride) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *LegalTemplateOverride) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *LegalTemplateOverride) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCounselReview

`func (o *LegalTemplateOverride) GetCounselReview() bool`

GetCounselReview returns the CounselReview field if non-nil, zero value otherwise.

### GetCounselReviewOk

`func (o *LegalTemplateOverride) GetCounselReviewOk() (*bool, bool)`

GetCounselReviewOk returns a tuple with the CounselReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounselReview

`func (o *LegalTemplateOverride) SetCounselReview(v bool)`

SetCounselReview sets CounselReview field to given value.

### HasCounselReview

`func (o *LegalTemplateOverride) HasCounselReview() bool`

HasCounselReview returns a boolean if a field has been set.

### GetFields

`func (o *LegalTemplateOverride) GetFields() []LegalField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *LegalTemplateOverride) GetFieldsOk() (*[]LegalField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *LegalTemplateOverride) SetFields(v []LegalField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *LegalTemplateOverride) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetId

`func (o *LegalTemplateOverride) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LegalTemplateOverride) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LegalTemplateOverride) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LegalTemplateOverride) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *LegalTemplateOverride) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LegalTemplateOverride) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LegalTemplateOverride) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *LegalTemplateOverride) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


