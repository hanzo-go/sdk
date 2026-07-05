# DbDatabase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**BranchId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OwnerName** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDbDatabase

`func NewDbDatabase() *DbDatabase`

NewDbDatabase instantiates a new DbDatabase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbDatabaseWithDefaults

`func NewDbDatabaseWithDefaults() *DbDatabase`

NewDbDatabaseWithDefaults instantiates a new DbDatabase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbDatabase) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbDatabase) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbDatabase) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DbDatabase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBranchId

`func (o *DbDatabase) GetBranchId() string`

GetBranchId returns the BranchId field if non-nil, zero value otherwise.

### GetBranchIdOk

`func (o *DbDatabase) GetBranchIdOk() (*string, bool)`

GetBranchIdOk returns a tuple with the BranchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchId

`func (o *DbDatabase) SetBranchId(v string)`

SetBranchId sets BranchId field to given value.

### HasBranchId

`func (o *DbDatabase) HasBranchId() bool`

HasBranchId returns a boolean if a field has been set.

### GetName

`func (o *DbDatabase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DbDatabase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DbDatabase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DbDatabase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerName

`func (o *DbDatabase) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *DbDatabase) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *DbDatabase) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *DbDatabase) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DbDatabase) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DbDatabase) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DbDatabase) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DbDatabase) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DbDatabase) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DbDatabase) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DbDatabase) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DbDatabase) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


