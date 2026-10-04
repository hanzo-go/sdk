# AffiliateClickRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is the share-link code that was clicked. Body-only: the URL cannot supply it. | [optional] 

## Methods

### NewAffiliateClickRequest

`func NewAffiliateClickRequest() *AffiliateClickRequest`

NewAffiliateClickRequest instantiates a new AffiliateClickRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffiliateClickRequestWithDefaults

`func NewAffiliateClickRequestWithDefaults() *AffiliateClickRequest`

NewAffiliateClickRequestWithDefaults instantiates a new AffiliateClickRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AffiliateClickRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AffiliateClickRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AffiliateClickRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AffiliateClickRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


