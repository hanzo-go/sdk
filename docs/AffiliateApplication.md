# AffiliateApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is the minted referral code. Empty on a first apply — applying does not mint a code, approval does; a re-apply echoes whatever the row already holds. | [optional] 
**Created** | Pointer to **bool** | Created says whether THIS call made the row. false means the org had already applied and nothing changed — no second row, no reset of an existing approval. The HTTP status states the same fact: 201 when true, 200 when false. | [optional] 
**Id** | Pointer to **string** | ID is the affiliate&#39;s server-minted handle, \&quot;aff_\&quot;-prefixed — the id staff approve, suspend, re-rate and pay against. | [optional] 
**RateBps** | Pointer to **int64** | RateBps is the direct (level 1) commission rate the row carries, in basis points OF Hanzo&#39;s margin (2000 &#x3D; 20% of margin, never of the customer&#39;s bill). | [optional] 
**RequestedCode** | Pointer to **string** | RequestedCode echoes the vanity code asked for, normalized to lower case. It is a request only: approval mints a different slug if this one is taken. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;applied\&quot; for a row this call created. A re-apply echoes the existing row&#39;s status, which may already be \&quot;approved\&quot; or \&quot;suspended\&quot;. | [optional] 

## Methods

### NewAffiliateApplication

`func NewAffiliateApplication() *AffiliateApplication`

NewAffiliateApplication instantiates a new AffiliateApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffiliateApplicationWithDefaults

`func NewAffiliateApplicationWithDefaults() *AffiliateApplication`

NewAffiliateApplicationWithDefaults instantiates a new AffiliateApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AffiliateApplication) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AffiliateApplication) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AffiliateApplication) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AffiliateApplication) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreated

`func (o *AffiliateApplication) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AffiliateApplication) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AffiliateApplication) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AffiliateApplication) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *AffiliateApplication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AffiliateApplication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AffiliateApplication) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AffiliateApplication) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRateBps

`func (o *AffiliateApplication) GetRateBps() int64`

GetRateBps returns the RateBps field if non-nil, zero value otherwise.

### GetRateBpsOk

`func (o *AffiliateApplication) GetRateBpsOk() (*int64, bool)`

GetRateBpsOk returns a tuple with the RateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateBps

`func (o *AffiliateApplication) SetRateBps(v int64)`

SetRateBps sets RateBps field to given value.

### HasRateBps

`func (o *AffiliateApplication) HasRateBps() bool`

HasRateBps returns a boolean if a field has been set.

### GetRequestedCode

`func (o *AffiliateApplication) GetRequestedCode() string`

GetRequestedCode returns the RequestedCode field if non-nil, zero value otherwise.

### GetRequestedCodeOk

`func (o *AffiliateApplication) GetRequestedCodeOk() (*string, bool)`

GetRequestedCodeOk returns a tuple with the RequestedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedCode

`func (o *AffiliateApplication) SetRequestedCode(v string)`

SetRequestedCode sets RequestedCode field to given value.

### HasRequestedCode

`func (o *AffiliateApplication) HasRequestedCode() bool`

HasRequestedCode returns a boolean if a field has been set.

### GetStatus

`func (o *AffiliateApplication) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AffiliateApplication) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AffiliateApplication) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AffiliateApplication) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


