# Attribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is the affiliate code the edge was recorded under, normalized to lower case. On a re-post it is the code of the STANDING edge, which may differ from the one just sent — first touch wins. | [optional] 
**Created** | Pointer to **bool** | Created says whether THIS call made the edge. false means the caller org was already attributed and nothing moved. The HTTP status says the same: 201 when true, 200 when false. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the edge was FIRST recorded, Unix seconds UTC. On a re-post it is the original time, not now. | [optional] 
**Id** | Pointer to **string** | ID is the attribution edge&#39;s server-minted handle, \&quot;afr_\&quot;-prefixed. | [optional] 

## Methods

### NewAttribution

`func NewAttribution() *Attribution`

NewAttribution instantiates a new Attribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributionWithDefaults

`func NewAttributionWithDefaults() *Attribution`

NewAttributionWithDefaults instantiates a new Attribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Attribution) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Attribution) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Attribution) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Attribution) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreated

`func (o *Attribution) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Attribution) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Attribution) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Attribution) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Attribution) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Attribution) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Attribution) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Attribution) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *Attribution) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Attribution) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Attribution) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Attribution) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


