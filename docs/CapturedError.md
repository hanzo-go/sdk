# CapturedError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistinctId** | Pointer to **string** | DistinctID is the person/visitor the error is attributed to. Omitted when the row carries none. | [optional] 
**Event** | Pointer to **string** | Event is the event name the error was stored under, e.g. $error. | [optional] 
**Exception** | Pointer to **interface{}** |  | [optional] 
**Id** | Pointer to **string** | ID is the row&#39;s stable event id — the client&#39;s own idempotency id when it sent one, else the server-minted one. | [optional] 
**Library** | Pointer to **string** | Library is the client SDK that reported the error. Omitted when absent. | [optional] 
**LibraryVersion** | Pointer to **string** | LibraryVer is that SDK&#39;s version. Omitted when absent. | [optional] 
**Path** | Pointer to **string** | Path is the URL&#39;s path component. Omitted when absent. | [optional] 
**Product** | Pointer to **string** | Product is the surface that emitted the error. Omitted when absent. | [optional] 
**Properties** | Pointer to **interface{}** |  | [optional] 
**SessionId** | Pointer to **string** | SessionID groups the events of one visit. Omitted when the client sent none. | [optional] 
**Timestamp** | Pointer to **string** | Timestamp is when the error was captured, RFC3339 UTC. | [optional] 
**Url** | Pointer to **string** | URL is the full page address the error fired on. Omitted when absent. | [optional] 

## Methods

### NewCapturedError

`func NewCapturedError() *CapturedError`

NewCapturedError instantiates a new CapturedError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapturedErrorWithDefaults

`func NewCapturedErrorWithDefaults() *CapturedError`

NewCapturedErrorWithDefaults instantiates a new CapturedError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistinctId

`func (o *CapturedError) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *CapturedError) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *CapturedError) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.

### HasDistinctId

`func (o *CapturedError) HasDistinctId() bool`

HasDistinctId returns a boolean if a field has been set.

### GetEvent

`func (o *CapturedError) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CapturedError) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CapturedError) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *CapturedError) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetException

`func (o *CapturedError) GetException() interface{}`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *CapturedError) GetExceptionOk() (*interface{}, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *CapturedError) SetException(v interface{})`

SetException sets Exception field to given value.

### HasException

`func (o *CapturedError) HasException() bool`

HasException returns a boolean if a field has been set.

### SetExceptionNil

`func (o *CapturedError) SetExceptionNil(b bool)`

 SetExceptionNil sets the value for Exception to be an explicit nil

### UnsetException
`func (o *CapturedError) UnsetException()`

UnsetException ensures that no value is present for Exception, not even an explicit nil
### GetId

`func (o *CapturedError) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CapturedError) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CapturedError) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CapturedError) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLibrary

`func (o *CapturedError) GetLibrary() string`

GetLibrary returns the Library field if non-nil, zero value otherwise.

### GetLibraryOk

`func (o *CapturedError) GetLibraryOk() (*string, bool)`

GetLibraryOk returns a tuple with the Library field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibrary

`func (o *CapturedError) SetLibrary(v string)`

SetLibrary sets Library field to given value.

### HasLibrary

`func (o *CapturedError) HasLibrary() bool`

HasLibrary returns a boolean if a field has been set.

### GetLibraryVersion

`func (o *CapturedError) GetLibraryVersion() string`

GetLibraryVersion returns the LibraryVersion field if non-nil, zero value otherwise.

### GetLibraryVersionOk

`func (o *CapturedError) GetLibraryVersionOk() (*string, bool)`

GetLibraryVersionOk returns a tuple with the LibraryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryVersion

`func (o *CapturedError) SetLibraryVersion(v string)`

SetLibraryVersion sets LibraryVersion field to given value.

### HasLibraryVersion

`func (o *CapturedError) HasLibraryVersion() bool`

HasLibraryVersion returns a boolean if a field has been set.

### GetPath

`func (o *CapturedError) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CapturedError) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CapturedError) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CapturedError) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProduct

`func (o *CapturedError) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CapturedError) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *CapturedError) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *CapturedError) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetProperties

`func (o *CapturedError) GetProperties() interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CapturedError) GetPropertiesOk() (*interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CapturedError) SetProperties(v interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CapturedError) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *CapturedError) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *CapturedError) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetSessionId

`func (o *CapturedError) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CapturedError) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *CapturedError) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *CapturedError) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CapturedError) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CapturedError) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CapturedError) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CapturedError) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUrl

`func (o *CapturedError) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CapturedError) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CapturedError) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CapturedError) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


