# AffiliatesApplyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**AffiliatesAffiliateStatus**](AffiliatesAffiliateStatus.md) |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**RequestedCode** | Pointer to **string** |  | [optional] 
**RateBps** | Pointer to **int64** |  | [optional] 
**Created** | Pointer to **bool** | True if this apply created a new affiliate record. | [optional] 

## Methods

### NewAffiliatesApplyResponse

`func NewAffiliatesApplyResponse() *AffiliatesApplyResponse`

NewAffiliatesApplyResponse instantiates a new AffiliatesApplyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffiliatesApplyResponseWithDefaults

`func NewAffiliatesApplyResponseWithDefaults() *AffiliatesApplyResponse`

NewAffiliatesApplyResponseWithDefaults instantiates a new AffiliatesApplyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AffiliatesApplyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AffiliatesApplyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AffiliatesApplyResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AffiliatesApplyResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *AffiliatesApplyResponse) GetStatus() AffiliatesAffiliateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AffiliatesApplyResponse) GetStatusOk() (*AffiliatesAffiliateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AffiliatesApplyResponse) SetStatus(v AffiliatesAffiliateStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AffiliatesApplyResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCode

`func (o *AffiliatesApplyResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AffiliatesApplyResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AffiliatesApplyResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AffiliatesApplyResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetRequestedCode

`func (o *AffiliatesApplyResponse) GetRequestedCode() string`

GetRequestedCode returns the RequestedCode field if non-nil, zero value otherwise.

### GetRequestedCodeOk

`func (o *AffiliatesApplyResponse) GetRequestedCodeOk() (*string, bool)`

GetRequestedCodeOk returns a tuple with the RequestedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedCode

`func (o *AffiliatesApplyResponse) SetRequestedCode(v string)`

SetRequestedCode sets RequestedCode field to given value.

### HasRequestedCode

`func (o *AffiliatesApplyResponse) HasRequestedCode() bool`

HasRequestedCode returns a boolean if a field has been set.

### GetRateBps

`func (o *AffiliatesApplyResponse) GetRateBps() int64`

GetRateBps returns the RateBps field if non-nil, zero value otherwise.

### GetRateBpsOk

`func (o *AffiliatesApplyResponse) GetRateBpsOk() (*int64, bool)`

GetRateBpsOk returns a tuple with the RateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateBps

`func (o *AffiliatesApplyResponse) SetRateBps(v int64)`

SetRateBps sets RateBps field to given value.

### HasRateBps

`func (o *AffiliatesApplyResponse) HasRateBps() bool`

HasRateBps returns a boolean if a field has been set.

### GetCreated

`func (o *AffiliatesApplyResponse) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AffiliatesApplyResponse) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AffiliatesApplyResponse) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AffiliatesApplyResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


