# O11yO11yNamespaceListOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**O11yNamespaceListResponse**](O11yNamespaceListResponse.md) | Data holds the namespace records. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;success\&quot;. | [optional] 

## Methods

### NewO11yO11yNamespaceListOut

`func NewO11yO11yNamespaceListOut() *O11yO11yNamespaceListOut`

NewO11yO11yNamespaceListOut instantiates a new O11yO11yNamespaceListOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yNamespaceListOutWithDefaults

`func NewO11yO11yNamespaceListOutWithDefaults() *O11yO11yNamespaceListOut`

NewO11yO11yNamespaceListOutWithDefaults instantiates a new O11yO11yNamespaceListOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11yNamespaceListOut) GetData() O11yNamespaceListResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yNamespaceListOut) GetDataOk() (*O11yNamespaceListResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yNamespaceListOut) SetData(v O11yNamespaceListResponse)`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yNamespaceListOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yNamespaceListOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yNamespaceListOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yNamespaceListOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yNamespaceListOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


