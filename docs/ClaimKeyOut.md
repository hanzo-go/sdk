# ClaimKeyOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimKey** | Pointer to **string** | ClaimKey is the capability itself. It is returned ONCE and never again — only its SHA-256 hash is stored — so a daemon that loses it mints a new one. | [optional] 
**TargetId** | Pointer to **string** | TargetID is the machine the key authenticates. | [optional] 

## Methods

### NewClaimKeyOut

`func NewClaimKeyOut() *ClaimKeyOut`

NewClaimKeyOut instantiates a new ClaimKeyOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClaimKeyOutWithDefaults

`func NewClaimKeyOutWithDefaults() *ClaimKeyOut`

NewClaimKeyOutWithDefaults instantiates a new ClaimKeyOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimKey

`func (o *ClaimKeyOut) GetClaimKey() string`

GetClaimKey returns the ClaimKey field if non-nil, zero value otherwise.

### GetClaimKeyOk

`func (o *ClaimKeyOut) GetClaimKeyOk() (*string, bool)`

GetClaimKeyOk returns a tuple with the ClaimKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimKey

`func (o *ClaimKeyOut) SetClaimKey(v string)`

SetClaimKey sets ClaimKey field to given value.

### HasClaimKey

`func (o *ClaimKeyOut) HasClaimKey() bool`

HasClaimKey returns a boolean if a field has been set.

### GetTargetId

`func (o *ClaimKeyOut) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *ClaimKeyOut) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *ClaimKeyOut) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *ClaimKeyOut) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


