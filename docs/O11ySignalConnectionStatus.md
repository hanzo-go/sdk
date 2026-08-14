# O11ySignalConnectionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastReceivedFrom** | Pointer to **string** | resource identifier | [optional] 
**LastReceivedTsMs** | Pointer to **int32** | epoch milliseconds | [optional] 

## Methods

### NewO11ySignalConnectionStatus

`func NewO11ySignalConnectionStatus() *O11ySignalConnectionStatus`

NewO11ySignalConnectionStatus instantiates a new O11ySignalConnectionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11ySignalConnectionStatusWithDefaults

`func NewO11ySignalConnectionStatusWithDefaults() *O11ySignalConnectionStatus`

NewO11ySignalConnectionStatusWithDefaults instantiates a new O11ySignalConnectionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastReceivedFrom

`func (o *O11ySignalConnectionStatus) GetLastReceivedFrom() string`

GetLastReceivedFrom returns the LastReceivedFrom field if non-nil, zero value otherwise.

### GetLastReceivedFromOk

`func (o *O11ySignalConnectionStatus) GetLastReceivedFromOk() (*string, bool)`

GetLastReceivedFromOk returns a tuple with the LastReceivedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReceivedFrom

`func (o *O11ySignalConnectionStatus) SetLastReceivedFrom(v string)`

SetLastReceivedFrom sets LastReceivedFrom field to given value.

### HasLastReceivedFrom

`func (o *O11ySignalConnectionStatus) HasLastReceivedFrom() bool`

HasLastReceivedFrom returns a boolean if a field has been set.

### GetLastReceivedTsMs

`func (o *O11ySignalConnectionStatus) GetLastReceivedTsMs() int32`

GetLastReceivedTsMs returns the LastReceivedTsMs field if non-nil, zero value otherwise.

### GetLastReceivedTsMsOk

`func (o *O11ySignalConnectionStatus) GetLastReceivedTsMsOk() (*int32, bool)`

GetLastReceivedTsMsOk returns a tuple with the LastReceivedTsMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReceivedTsMs

`func (o *O11ySignalConnectionStatus) SetLastReceivedTsMs(v int32)`

SetLastReceivedTsMs sets LastReceivedTsMs field to given value.

### HasLastReceivedTsMs

`func (o *O11ySignalConnectionStatus) HasLastReceivedTsMs() bool`

HasLastReceivedTsMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


