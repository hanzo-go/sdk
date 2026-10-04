# ContentGenerateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brief** | Pointer to **string** | the brief/goal driving copy generation | [optional] 
**Channels** | Pointer to **string** | target channels (SocialPost) | [optional] 
**Design** | Pointer to **string** | studio design slug (asset source) | [optional] 
**Doctype** | Pointer to **string** | marketing.Campaign | marketing.SocialPost | marketing.Asset | [optional] 
**Kind** | Pointer to **string** | asset kind: ecom|product|lifestyle|hover|hero | [optional] 
**Model** | Pointer to **string** | optional zen model override (copy) | [optional] 
**Product** | Pointer to **string** | commerce product handle (copy context) | [optional] 
**Project** | Pointer to **string** | brand/site sub-scope (billing + tenancy axis) | [optional] 
**SourceMedia** | Pointer to **string** | asset source image (design CAD/photo) | [optional] 
**Title** | Pointer to **string** | optional explicit title | [optional] 
**Tone** | Pointer to **string** | tone override for a single draft | [optional] 
**Voice** | Pointer to **string** | brand-voice guidance for the copy director | [optional] 

## Methods

### NewContentGenerateInput

`func NewContentGenerateInput() *ContentGenerateInput`

NewContentGenerateInput instantiates a new ContentGenerateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentGenerateInputWithDefaults

`func NewContentGenerateInputWithDefaults() *ContentGenerateInput`

NewContentGenerateInputWithDefaults instantiates a new ContentGenerateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrief

`func (o *ContentGenerateInput) GetBrief() string`

GetBrief returns the Brief field if non-nil, zero value otherwise.

### GetBriefOk

`func (o *ContentGenerateInput) GetBriefOk() (*string, bool)`

GetBriefOk returns a tuple with the Brief field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrief

`func (o *ContentGenerateInput) SetBrief(v string)`

SetBrief sets Brief field to given value.

### HasBrief

`func (o *ContentGenerateInput) HasBrief() bool`

HasBrief returns a boolean if a field has been set.

### GetChannels

`func (o *ContentGenerateInput) GetChannels() string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ContentGenerateInput) GetChannelsOk() (*string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ContentGenerateInput) SetChannels(v string)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *ContentGenerateInput) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetDesign

`func (o *ContentGenerateInput) GetDesign() string`

GetDesign returns the Design field if non-nil, zero value otherwise.

### GetDesignOk

`func (o *ContentGenerateInput) GetDesignOk() (*string, bool)`

GetDesignOk returns a tuple with the Design field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesign

`func (o *ContentGenerateInput) SetDesign(v string)`

SetDesign sets Design field to given value.

### HasDesign

`func (o *ContentGenerateInput) HasDesign() bool`

HasDesign returns a boolean if a field has been set.

### GetDoctype

`func (o *ContentGenerateInput) GetDoctype() string`

GetDoctype returns the Doctype field if non-nil, zero value otherwise.

### GetDoctypeOk

`func (o *ContentGenerateInput) GetDoctypeOk() (*string, bool)`

GetDoctypeOk returns a tuple with the Doctype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoctype

`func (o *ContentGenerateInput) SetDoctype(v string)`

SetDoctype sets Doctype field to given value.

### HasDoctype

`func (o *ContentGenerateInput) HasDoctype() bool`

HasDoctype returns a boolean if a field has been set.

### GetKind

`func (o *ContentGenerateInput) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ContentGenerateInput) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ContentGenerateInput) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ContentGenerateInput) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModel

`func (o *ContentGenerateInput) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ContentGenerateInput) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ContentGenerateInput) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ContentGenerateInput) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetProduct

`func (o *ContentGenerateInput) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ContentGenerateInput) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ContentGenerateInput) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *ContentGenerateInput) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetProject

`func (o *ContentGenerateInput) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ContentGenerateInput) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ContentGenerateInput) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ContentGenerateInput) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetSourceMedia

`func (o *ContentGenerateInput) GetSourceMedia() string`

GetSourceMedia returns the SourceMedia field if non-nil, zero value otherwise.

### GetSourceMediaOk

`func (o *ContentGenerateInput) GetSourceMediaOk() (*string, bool)`

GetSourceMediaOk returns a tuple with the SourceMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMedia

`func (o *ContentGenerateInput) SetSourceMedia(v string)`

SetSourceMedia sets SourceMedia field to given value.

### HasSourceMedia

`func (o *ContentGenerateInput) HasSourceMedia() bool`

HasSourceMedia returns a boolean if a field has been set.

### GetTitle

`func (o *ContentGenerateInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ContentGenerateInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ContentGenerateInput) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ContentGenerateInput) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTone

`func (o *ContentGenerateInput) GetTone() string`

GetTone returns the Tone field if non-nil, zero value otherwise.

### GetToneOk

`func (o *ContentGenerateInput) GetToneOk() (*string, bool)`

GetToneOk returns a tuple with the Tone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTone

`func (o *ContentGenerateInput) SetTone(v string)`

SetTone sets Tone field to given value.

### HasTone

`func (o *ContentGenerateInput) HasTone() bool`

HasTone returns a boolean if a field has been set.

### GetVoice

`func (o *ContentGenerateInput) GetVoice() string`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *ContentGenerateInput) GetVoiceOk() (*string, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *ContentGenerateInput) SetVoice(v string)`

SetVoice sets Voice field to given value.

### HasVoice

`func (o *ContentGenerateInput) HasVoice() bool`

HasVoice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


