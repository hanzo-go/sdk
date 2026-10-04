# AffiliateCreateLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is an optional vanity code; it must be free across the whole directory, and omitting it mints a random one. Body-only. | [optional] 
**Label** | Pointer to **string** | Label is cosmetic — trimmed, stripped of control characters, capped — and never part of a code. Body-only: the URL cannot supply it. | [optional] 

## Methods

### NewAffiliateCreateLinkRequest

`func NewAffiliateCreateLinkRequest() *AffiliateCreateLinkRequest`

NewAffiliateCreateLinkRequest instantiates a new AffiliateCreateLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffiliateCreateLinkRequestWithDefaults

`func NewAffiliateCreateLinkRequestWithDefaults() *AffiliateCreateLinkRequest`

NewAffiliateCreateLinkRequestWithDefaults instantiates a new AffiliateCreateLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AffiliateCreateLinkRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AffiliateCreateLinkRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AffiliateCreateLinkRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AffiliateCreateLinkRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLabel

`func (o *AffiliateCreateLinkRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AffiliateCreateLinkRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AffiliateCreateLinkRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AffiliateCreateLinkRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


