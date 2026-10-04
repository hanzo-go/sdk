# AffiliateAttribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is the affiliate code the edge was recorded under, normalized to lower case. On a re-post it is the code of the STANDING edge, which may differ from the one just sent — first touch wins. | [optional] 
**Created** | Pointer to **bool** | Created says whether THIS call made the edge. false means the caller org was already attributed and nothing moved. The HTTP status says the same: 201 when true, 200 when false. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the edge was FIRST recorded, Unix seconds UTC. On a re-post it is the original time, not now. | [optional] 
**Id** | Pointer to **string** | ID is the attribution edge&#39;s server-minted handle, \&quot;afr_\&quot;-prefixed. | [optional] 

## Methods

### NewAffiliateAttribution

`func NewAffiliateAttribution() *AffiliateAttribution`

NewAffiliateAttribution instantiates a new AffiliateAttribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffiliateAttributionWithDefaults

`func NewAffiliateAttributionWithDefaults() *AffiliateAttribution`

NewAffiliateAttributionWithDefaults instantiates a new AffiliateAttribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AffiliateAttribution) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AffiliateAttribution) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AffiliateAttribution) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AffiliateAttribution) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreated

`func (o *AffiliateAttribution) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AffiliateAttribution) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AffiliateAttribution) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AffiliateAttribution) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AffiliateAttribution) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AffiliateAttribution) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AffiliateAttribution) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AffiliateAttribution) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *AffiliateAttribution) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AffiliateAttribution) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AffiliateAttribution) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AffiliateAttribution) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


