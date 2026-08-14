# CapIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the cap to edit or remove, from the path. Unused by the list and create ops. | [optional] 
**Org** | Pointer to **string** | Org is the tenant to act on. Required for a SuperAdmin — they must name their target; ignored for a white-label admin, who always acts on their own org. | [optional] 

## Methods

### NewCapIn

`func NewCapIn() *CapIn`

NewCapIn instantiates a new CapIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapInWithDefaults

`func NewCapInWithDefaults() *CapIn`

NewCapInWithDefaults instantiates a new CapIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CapIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CapIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CapIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CapIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrg

`func (o *CapIn) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CapIn) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CapIn) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CapIn) HasOrg() bool`

HasOrg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


