# O11yO11yQueueChecksOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]O11yO11yQueueCheck**](O11yO11yQueueCheck.md) | Data holds one check per required attribute, sorted by attribute. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;success\&quot;. | [optional] 

## Methods

### NewO11yO11yQueueChecksOut

`func NewO11yO11yQueueChecksOut() *O11yO11yQueueChecksOut`

NewO11yO11yQueueChecksOut instantiates a new O11yO11yQueueChecksOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yQueueChecksOutWithDefaults

`func NewO11yO11yQueueChecksOutWithDefaults() *O11yO11yQueueChecksOut`

NewO11yO11yQueueChecksOutWithDefaults instantiates a new O11yO11yQueueChecksOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11yQueueChecksOut) GetData() []O11yO11yQueueCheck`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yQueueChecksOut) GetDataOk() (*[]O11yO11yQueueCheck, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yQueueChecksOut) SetData(v []O11yO11yQueueCheck)`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yQueueChecksOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yQueueChecksOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yQueueChecksOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yQueueChecksOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yQueueChecksOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


