# AffiliatesAttributeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The attribution edge id (e.g. &#x60;afr_&lt;hex&gt;&#x60;). | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **bool** | True if this call created a new attribution edge. | [optional] 
**CreatedAt** | Pointer to **int64** | Unix timestamp (seconds). | [optional] 

## Methods

### NewAffiliatesAttributeResponse

`func NewAffiliatesAttributeResponse() *AffiliatesAttributeResponse`

NewAffiliatesAttributeResponse instantiates a new AffiliatesAttributeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffiliatesAttributeResponseWithDefaults

`func NewAffiliatesAttributeResponseWithDefaults() *AffiliatesAttributeResponse`

NewAffiliatesAttributeResponseWithDefaults instantiates a new AffiliatesAttributeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AffiliatesAttributeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AffiliatesAttributeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AffiliatesAttributeResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AffiliatesAttributeResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *AffiliatesAttributeResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AffiliatesAttributeResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AffiliatesAttributeResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AffiliatesAttributeResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreated

`func (o *AffiliatesAttributeResponse) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AffiliatesAttributeResponse) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AffiliatesAttributeResponse) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AffiliatesAttributeResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AffiliatesAttributeResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AffiliatesAttributeResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AffiliatesAttributeResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AffiliatesAttributeResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


