# KitList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]StarterKit**](StarterKit.md) | Data is the public catalog followed by the caller org&#39;s own kits. | [optional] 

## Methods

### NewKitList

`func NewKitList() *KitList`

NewKitList instantiates a new KitList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKitListWithDefaults

`func NewKitListWithDefaults() *KitList`

NewKitListWithDefaults instantiates a new KitList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *KitList) GetData() []StarterKit`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *KitList) GetDataOk() (*[]StarterKit, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *KitList) SetData(v []StarterKit)`

SetData sets Data field to given value.

### HasData

`func (o *KitList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


