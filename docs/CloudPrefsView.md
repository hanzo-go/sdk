# CloudPrefsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefs** | Pointer to **interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is when the document was last written, unix seconds. Absent when nothing has been saved. | [optional] 

## Methods

### NewCloudPrefsView

`func NewCloudPrefsView() *CloudPrefsView`

NewCloudPrefsView instantiates a new CloudPrefsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPrefsViewWithDefaults

`func NewCloudPrefsViewWithDefaults() *CloudPrefsView`

NewCloudPrefsViewWithDefaults instantiates a new CloudPrefsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefs

`func (o *CloudPrefsView) GetPrefs() interface{}`

GetPrefs returns the Prefs field if non-nil, zero value otherwise.

### GetPrefsOk

`func (o *CloudPrefsView) GetPrefsOk() (*interface{}, bool)`

GetPrefsOk returns a tuple with the Prefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefs

`func (o *CloudPrefsView) SetPrefs(v interface{})`

SetPrefs sets Prefs field to given value.

### HasPrefs

`func (o *CloudPrefsView) HasPrefs() bool`

HasPrefs returns a boolean if a field has been set.

### SetPrefsNil

`func (o *CloudPrefsView) SetPrefsNil(b bool)`

 SetPrefsNil sets the value for Prefs to be an explicit nil

### UnsetPrefs
`func (o *CloudPrefsView) UnsetPrefs()`

UnsetPrefs ensures that no value is present for Prefs, not even an explicit nil
### GetUpdatedAt

`func (o *CloudPrefsView) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudPrefsView) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudPrefsView) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudPrefsView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


