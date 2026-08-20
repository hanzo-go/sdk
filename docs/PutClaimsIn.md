# PutClaimsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**By** | Pointer to **string** | By is who is recording them — a person, or the importer&#39;s name. | [optional] 
**Data** | Pointer to [**[]PublishedClaim**](PublishedClaim.md) | Data is the claims to record. One row is a correction; many is an import. There is no separate bulk endpoint because there is no separate operation: importing a leaderboard and fixing one number are the same write. | [optional] 

## Methods

### NewPutClaimsIn

`func NewPutClaimsIn() *PutClaimsIn`

NewPutClaimsIn instantiates a new PutClaimsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutClaimsInWithDefaults

`func NewPutClaimsInWithDefaults() *PutClaimsIn`

NewPutClaimsInWithDefaults instantiates a new PutClaimsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBy

`func (o *PutClaimsIn) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *PutClaimsIn) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *PutClaimsIn) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *PutClaimsIn) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetData

`func (o *PutClaimsIn) GetData() []PublishedClaim`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PutClaimsIn) GetDataOk() (*[]PublishedClaim, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PutClaimsIn) SetData(v []PublishedClaim)`

SetData sets Data field to given value.

### HasData

`func (o *PutClaimsIn) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


