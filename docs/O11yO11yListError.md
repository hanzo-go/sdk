# O11yO11yListError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExceptionCount** | Pointer to **int32** | ExceptionCount is how many instances the group holds in the window. | [optional] 
**ExceptionMessage** | Pointer to **string** | ExceptionMsg is its message. | [optional] 
**ExceptionType** | Pointer to **string** | ExceptionType is the exception&#39;s type. | [optional] 
**FirstSeen** | Pointer to **time.Time** | FirstSeen is when the earliest was. | [optional] 
**GroupID** | Pointer to **string** | GroupID is the group&#39;s id. | [optional] 
**LastSeen** | Pointer to **time.Time** | LastSeen is when the latest instance was recorded. | [optional] 
**ServiceName** | Pointer to **string** | ServiceName is the service that reported them. | [optional] 

## Methods

### NewO11yO11yListError

`func NewO11yO11yListError() *O11yO11yListError`

NewO11yO11yListError instantiates a new O11yO11yListError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yListErrorWithDefaults

`func NewO11yO11yListErrorWithDefaults() *O11yO11yListError`

NewO11yO11yListErrorWithDefaults instantiates a new O11yO11yListError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExceptionCount

`func (o *O11yO11yListError) GetExceptionCount() int32`

GetExceptionCount returns the ExceptionCount field if non-nil, zero value otherwise.

### GetExceptionCountOk

`func (o *O11yO11yListError) GetExceptionCountOk() (*int32, bool)`

GetExceptionCountOk returns a tuple with the ExceptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionCount

`func (o *O11yO11yListError) SetExceptionCount(v int32)`

SetExceptionCount sets ExceptionCount field to given value.

### HasExceptionCount

`func (o *O11yO11yListError) HasExceptionCount() bool`

HasExceptionCount returns a boolean if a field has been set.

### GetExceptionMessage

`func (o *O11yO11yListError) GetExceptionMessage() string`

GetExceptionMessage returns the ExceptionMessage field if non-nil, zero value otherwise.

### GetExceptionMessageOk

`func (o *O11yO11yListError) GetExceptionMessageOk() (*string, bool)`

GetExceptionMessageOk returns a tuple with the ExceptionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionMessage

`func (o *O11yO11yListError) SetExceptionMessage(v string)`

SetExceptionMessage sets ExceptionMessage field to given value.

### HasExceptionMessage

`func (o *O11yO11yListError) HasExceptionMessage() bool`

HasExceptionMessage returns a boolean if a field has been set.

### GetExceptionType

`func (o *O11yO11yListError) GetExceptionType() string`

GetExceptionType returns the ExceptionType field if non-nil, zero value otherwise.

### GetExceptionTypeOk

`func (o *O11yO11yListError) GetExceptionTypeOk() (*string, bool)`

GetExceptionTypeOk returns a tuple with the ExceptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionType

`func (o *O11yO11yListError) SetExceptionType(v string)`

SetExceptionType sets ExceptionType field to given value.

### HasExceptionType

`func (o *O11yO11yListError) HasExceptionType() bool`

HasExceptionType returns a boolean if a field has been set.

### GetFirstSeen

`func (o *O11yO11yListError) GetFirstSeen() time.Time`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *O11yO11yListError) GetFirstSeenOk() (*time.Time, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *O11yO11yListError) SetFirstSeen(v time.Time)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *O11yO11yListError) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetGroupID

`func (o *O11yO11yListError) GetGroupID() string`

GetGroupID returns the GroupID field if non-nil, zero value otherwise.

### GetGroupIDOk

`func (o *O11yO11yListError) GetGroupIDOk() (*string, bool)`

GetGroupIDOk returns a tuple with the GroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupID

`func (o *O11yO11yListError) SetGroupID(v string)`

SetGroupID sets GroupID field to given value.

### HasGroupID

`func (o *O11yO11yListError) HasGroupID() bool`

HasGroupID returns a boolean if a field has been set.

### GetLastSeen

`func (o *O11yO11yListError) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *O11yO11yListError) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *O11yO11yListError) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *O11yO11yListError) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetServiceName

`func (o *O11yO11yListError) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *O11yO11yListError) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *O11yO11yListError) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *O11yO11yListError) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


