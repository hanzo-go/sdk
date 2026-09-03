# IdentityList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identities** | Pointer to [**[]IdentityView**](IdentityView.md) | Identities is one row per fabric identity tagged with the caller&#39;s org role. | [optional] 

## Methods

### NewIdentityList

`func NewIdentityList() *IdentityList`

NewIdentityList instantiates a new IdentityList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityListWithDefaults

`func NewIdentityListWithDefaults() *IdentityList`

NewIdentityListWithDefaults instantiates a new IdentityList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentities

`func (o *IdentityList) GetIdentities() []IdentityView`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *IdentityList) GetIdentitiesOk() (*[]IdentityView, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *IdentityList) SetIdentities(v []IdentityView)`

SetIdentities sets Identities field to given value.

### HasIdentities

`func (o *IdentityList) HasIdentities() bool`

HasIdentities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


