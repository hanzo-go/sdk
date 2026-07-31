# CloudCapturedError

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

### NewCloudCapturedError

`func NewCloudCapturedError() *CloudCapturedError`

NewCloudCapturedError instantiates a new CloudCapturedError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCapturedErrorWithDefaults

`func NewCloudCapturedErrorWithDefaults() *CloudCapturedError`

NewCloudCapturedErrorWithDefaults instantiates a new CloudCapturedError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistinctId

`func (o *CloudCapturedError) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *CloudCapturedError) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *CloudCapturedError) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.

### HasDistinctId

`func (o *CloudCapturedError) HasDistinctId() bool`

HasDistinctId returns a boolean if a field has been set.

### GetEvent

`func (o *CloudCapturedError) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CloudCapturedError) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CloudCapturedError) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *CloudCapturedError) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetException

`func (o *CloudCapturedError) GetException() interface{}`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *CloudCapturedError) GetExceptionOk() (*interface{}, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *CloudCapturedError) SetException(v interface{})`

SetException sets Exception field to given value.

### HasException

`func (o *CloudCapturedError) HasException() bool`

HasException returns a boolean if a field has been set.

### SetExceptionNil

`func (o *CloudCapturedError) SetExceptionNil(b bool)`

 SetExceptionNil sets the value for Exception to be an explicit nil

### UnsetException
`func (o *CloudCapturedError) UnsetException()`

UnsetException ensures that no value is present for Exception, not even an explicit nil
### GetId

`func (o *CloudCapturedError) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudCapturedError) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudCapturedError) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudCapturedError) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLibrary

`func (o *CloudCapturedError) GetLibrary() string`

GetLibrary returns the Library field if non-nil, zero value otherwise.

### GetLibraryOk

`func (o *CloudCapturedError) GetLibraryOk() (*string, bool)`

GetLibraryOk returns a tuple with the Library field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibrary

`func (o *CloudCapturedError) SetLibrary(v string)`

SetLibrary sets Library field to given value.

### HasLibrary

`func (o *CloudCapturedError) HasLibrary() bool`

HasLibrary returns a boolean if a field has been set.

### GetLibraryVersion

`func (o *CloudCapturedError) GetLibraryVersion() string`

GetLibraryVersion returns the LibraryVersion field if non-nil, zero value otherwise.

### GetLibraryVersionOk

`func (o *CloudCapturedError) GetLibraryVersionOk() (*string, bool)`

GetLibraryVersionOk returns a tuple with the LibraryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryVersion

`func (o *CloudCapturedError) SetLibraryVersion(v string)`

SetLibraryVersion sets LibraryVersion field to given value.

### HasLibraryVersion

`func (o *CloudCapturedError) HasLibraryVersion() bool`

HasLibraryVersion returns a boolean if a field has been set.

### GetPath

`func (o *CloudCapturedError) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CloudCapturedError) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CloudCapturedError) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CloudCapturedError) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProduct

`func (o *CloudCapturedError) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CloudCapturedError) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *CloudCapturedError) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *CloudCapturedError) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetProperties

`func (o *CloudCapturedError) GetProperties() interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CloudCapturedError) GetPropertiesOk() (*interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CloudCapturedError) SetProperties(v interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CloudCapturedError) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *CloudCapturedError) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *CloudCapturedError) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetSessionId

`func (o *CloudCapturedError) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CloudCapturedError) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *CloudCapturedError) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *CloudCapturedError) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CloudCapturedError) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CloudCapturedError) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CloudCapturedError) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CloudCapturedError) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUrl

`func (o *CloudCapturedError) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CloudCapturedError) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CloudCapturedError) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CloudCapturedError) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


