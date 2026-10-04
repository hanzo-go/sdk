# PrefPrefsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefs** | Pointer to **interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is when the document was last written, unix seconds. Absent when nothing has been saved. | [optional] 

## Methods

### NewPrefPrefsView

`func NewPrefPrefsView() *PrefPrefsView`

NewPrefPrefsView instantiates a new PrefPrefsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrefPrefsViewWithDefaults

`func NewPrefPrefsViewWithDefaults() *PrefPrefsView`

NewPrefPrefsViewWithDefaults instantiates a new PrefPrefsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefs

`func (o *PrefPrefsView) GetPrefs() interface{}`

GetPrefs returns the Prefs field if non-nil, zero value otherwise.

### GetPrefsOk

`func (o *PrefPrefsView) GetPrefsOk() (*interface{}, bool)`

GetPrefsOk returns a tuple with the Prefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefs

`func (o *PrefPrefsView) SetPrefs(v interface{})`

SetPrefs sets Prefs field to given value.

### HasPrefs

`func (o *PrefPrefsView) HasPrefs() bool`

HasPrefs returns a boolean if a field has been set.

### SetPrefsNil

`func (o *PrefPrefsView) SetPrefsNil(b bool)`

 SetPrefsNil sets the value for Prefs to be an explicit nil

### UnsetPrefs
`func (o *PrefPrefsView) UnsetPrefs()`

UnsetPrefs ensures that no value is present for Prefs, not even an explicit nil
### GetUpdatedAt

`func (o *PrefPrefsView) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PrefPrefsView) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PrefPrefsView) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PrefPrefsView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


