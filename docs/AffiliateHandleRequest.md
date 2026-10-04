# AffiliateHandleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | Pointer to **string** | Handle is the public leaderboard display name; empty opts out. Body-only: the URL cannot supply it. | [optional] 

## Methods

### NewAffiliateHandleRequest

`func NewAffiliateHandleRequest() *AffiliateHandleRequest`

NewAffiliateHandleRequest instantiates a new AffiliateHandleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffiliateHandleRequestWithDefaults

`func NewAffiliateHandleRequestWithDefaults() *AffiliateHandleRequest`

NewAffiliateHandleRequestWithDefaults instantiates a new AffiliateHandleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandle

`func (o *AffiliateHandleRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *AffiliateHandleRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *AffiliateHandleRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *AffiliateHandleRequest) HasHandle() bool`

HasHandle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


