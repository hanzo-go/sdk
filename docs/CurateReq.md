# CurateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Featured** | Pointer to **bool** | Featured puts the listing on the front of the shelf, or takes it off. | [optional] 
**Hidden** | Pointer to **bool** | Hidden takes the listing off the org-visible shelf, or puts it back. | [optional] 
**Id** | Pointer to **string** | ID is the listing to curate, from the path. | [optional] 
**Logo** | Pointer to **string** | Logo is the brand mark to render, an https URL. Empty clears ours and lets the next sync adopt the publisher&#39;s own icon again. | [optional] 
**Official** | Pointer to **bool** | Official overrides the derivation: setting it makes this answer FINAL, so no later sync re-derives over it. That is the difference between a default and a decision — the derivation can only tell that a domain-verified publisher serves the endpoint, not that the product is theirs. | [optional] 

## Methods

### NewCurateReq

`func NewCurateReq() *CurateReq`

NewCurateReq instantiates a new CurateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurateReqWithDefaults

`func NewCurateReqWithDefaults() *CurateReq`

NewCurateReqWithDefaults instantiates a new CurateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatured

`func (o *CurateReq) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *CurateReq) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *CurateReq) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *CurateReq) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetHidden

`func (o *CurateReq) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *CurateReq) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *CurateReq) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *CurateReq) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetId

`func (o *CurateReq) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CurateReq) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CurateReq) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CurateReq) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogo

`func (o *CurateReq) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *CurateReq) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *CurateReq) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *CurateReq) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetOfficial

`func (o *CurateReq) GetOfficial() bool`

GetOfficial returns the Official field if non-nil, zero value otherwise.

### GetOfficialOk

`func (o *CurateReq) GetOfficialOk() (*bool, bool)`

GetOfficialOk returns a tuple with the Official field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficial

`func (o *CurateReq) SetOfficial(v bool)`

SetOfficial sets Official field to given value.

### HasOfficial

`func (o *CurateReq) HasOfficial() bool`

HasOfficial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


