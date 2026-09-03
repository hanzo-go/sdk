# Section

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **string** | Detail is what this phase of the journey is for, in prose. | [optional] 
**Enabled** | Pointer to **bool** | Enabled is the admin lever. Absent reads as ON, so only an explicit false turns a phase off — and it takes every step filed under it out of the journey, not just the heading. | [optional] 
**Id** | Pointer to **string** | ID is the slug a step&#39;s &#x60;section&#x60; names to file itself under this phase. | [optional] 
**Order** | Pointer to **int64** | Order places the phase in the journey, ascending. Ties fall back to authoring order, and an omitted order sorts as 0 — ahead of everything. | [optional] 
**Title** | Pointer to **string** | Title is the phase heading a person reads above its steps. | [optional] 

## Methods

### NewSection

`func NewSection() *Section`

NewSection instantiates a new Section object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectionWithDefaults

`func NewSectionWithDefaults() *Section`

NewSectionWithDefaults instantiates a new Section object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *Section) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *Section) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *Section) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *Section) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetEnabled

`func (o *Section) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Section) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Section) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Section) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *Section) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Section) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Section) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Section) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrder

`func (o *Section) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Section) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Section) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Section) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetTitle

`func (o *Section) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Section) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Section) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Section) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


