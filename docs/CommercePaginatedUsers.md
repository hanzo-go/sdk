# CommercePaginatedUsers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **string** |  | [optional] 
**Display** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Models** | Pointer to [**[]CommerceUser**](CommerceUser.md) |  | [optional] 

## Methods

### NewCommercePaginatedUsers

`func NewCommercePaginatedUsers() *CommercePaginatedUsers`

NewCommercePaginatedUsers instantiates a new CommercePaginatedUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommercePaginatedUsersWithDefaults

`func NewCommercePaginatedUsersWithDefaults() *CommercePaginatedUsers`

NewCommercePaginatedUsersWithDefaults instantiates a new CommercePaginatedUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *CommercePaginatedUsers) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *CommercePaginatedUsers) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *CommercePaginatedUsers) SetPage(v string)`

SetPage sets Page field to given value.

### HasPage

`func (o *CommercePaginatedUsers) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetDisplay

`func (o *CommercePaginatedUsers) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CommercePaginatedUsers) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CommercePaginatedUsers) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *CommercePaginatedUsers) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetCount

`func (o *CommercePaginatedUsers) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CommercePaginatedUsers) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CommercePaginatedUsers) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CommercePaginatedUsers) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetModels

`func (o *CommercePaginatedUsers) GetModels() []CommerceUser`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *CommercePaginatedUsers) GetModelsOk() (*[]CommerceUser, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *CommercePaginatedUsers) SetModels(v []CommerceUser)`

SetModels sets Models field to given value.

### HasModels

`func (o *CommercePaginatedUsers) HasModels() bool`

HasModels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


