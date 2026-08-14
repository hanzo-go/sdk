# CaptableShares

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CaptableShare**](CaptableShare.md) | Data is every issued share certificate, newest first. | [optional] 

## Methods

### NewCaptableShares

`func NewCaptableShares() *CaptableShares`

NewCaptableShares instantiates a new CaptableShares object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableSharesWithDefaults

`func NewCaptableSharesWithDefaults() *CaptableShares`

NewCaptableSharesWithDefaults instantiates a new CaptableShares object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CaptableShares) GetData() []CaptableShare`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CaptableShares) GetDataOk() (*[]CaptableShare, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CaptableShares) SetData(v []CaptableShare)`

SetData sets Data field to given value.

### HasData

`func (o *CaptableShares) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


