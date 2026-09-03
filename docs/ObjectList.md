# ObjectList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | Bucket is the bucket that was listed, friendly name. | [optional] 
**Objects** | Pointer to [**[]ObjectItem**](ObjectItem.md) | Objects are the entries at this level, keys RELATIVE to Prefix. | [optional] 
**Prefix** | Pointer to **string** | Prefix is the sub-folder the listing was scoped to, cleaned. Empty for the bucket root. | [optional] 
**Total** | Pointer to **int64** | Total is how many entries came back. The listing is BOUNDED, so a bucket with more keys than the cap answers the cap and this says so — it is not a count of what the bucket holds. | [optional] 

## Methods

### NewObjectList

`func NewObjectList() *ObjectList`

NewObjectList instantiates a new ObjectList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectListWithDefaults

`func NewObjectListWithDefaults() *ObjectList`

NewObjectListWithDefaults instantiates a new ObjectList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *ObjectList) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *ObjectList) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *ObjectList) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *ObjectList) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetObjects

`func (o *ObjectList) GetObjects() []ObjectItem`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ObjectList) GetObjectsOk() (*[]ObjectItem, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ObjectList) SetObjects(v []ObjectItem)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *ObjectList) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetPrefix

`func (o *ObjectList) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ObjectList) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ObjectList) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ObjectList) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetTotal

`func (o *ObjectList) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ObjectList) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ObjectList) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ObjectList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


