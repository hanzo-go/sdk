# PrefsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefs** | Pointer to **interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is when the document was last written, unix seconds. Absent when nothing has been saved. | [optional] 

## Methods

### NewPrefsView

`func NewPrefsView() *PrefsView`

NewPrefsView instantiates a new PrefsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrefsViewWithDefaults

`func NewPrefsViewWithDefaults() *PrefsView`

NewPrefsViewWithDefaults instantiates a new PrefsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefs

`func (o *PrefsView) GetPrefs() interface{}`

GetPrefs returns the Prefs field if non-nil, zero value otherwise.

### GetPrefsOk

`func (o *PrefsView) GetPrefsOk() (*interface{}, bool)`

GetPrefsOk returns a tuple with the Prefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefs

`func (o *PrefsView) SetPrefs(v interface{})`

SetPrefs sets Prefs field to given value.

### HasPrefs

`func (o *PrefsView) HasPrefs() bool`

HasPrefs returns a boolean if a field has been set.

### SetPrefsNil

`func (o *PrefsView) SetPrefsNil(b bool)`

 SetPrefsNil sets the value for Prefs to be an explicit nil

### UnsetPrefs
`func (o *PrefsView) UnsetPrefs()`

UnsetPrefs ensures that no value is present for Prefs, not even an explicit nil
### GetUpdatedAt

`func (o *PrefsView) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PrefsView) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PrefsView) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PrefsView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


