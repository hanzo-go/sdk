# CloudCaptableRoundCloseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloseDate** | Pointer to **interface{}** |  | [optional] 
**Id** | Pointer to **string** | ID is the round to close. It is the path segment: the URL is the addressing authority, and the org it is resolved in comes from the caller&#39;s principal, so an id from another tenant is simply not found. | [optional] 

## Methods

### NewCloudCaptableRoundCloseRequest

`func NewCloudCaptableRoundCloseRequest() *CloudCaptableRoundCloseRequest`

NewCloudCaptableRoundCloseRequest instantiates a new CloudCaptableRoundCloseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCaptableRoundCloseRequestWithDefaults

`func NewCloudCaptableRoundCloseRequestWithDefaults() *CloudCaptableRoundCloseRequest`

NewCloudCaptableRoundCloseRequestWithDefaults instantiates a new CloudCaptableRoundCloseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloseDate

`func (o *CloudCaptableRoundCloseRequest) GetCloseDate() interface{}`

GetCloseDate returns the CloseDate field if non-nil, zero value otherwise.

### GetCloseDateOk

`func (o *CloudCaptableRoundCloseRequest) GetCloseDateOk() (*interface{}, bool)`

GetCloseDateOk returns a tuple with the CloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseDate

`func (o *CloudCaptableRoundCloseRequest) SetCloseDate(v interface{})`

SetCloseDate sets CloseDate field to given value.

### HasCloseDate

`func (o *CloudCaptableRoundCloseRequest) HasCloseDate() bool`

HasCloseDate returns a boolean if a field has been set.

### SetCloseDateNil

`func (o *CloudCaptableRoundCloseRequest) SetCloseDateNil(b bool)`

 SetCloseDateNil sets the value for CloseDate to be an explicit nil

### UnsetCloseDate
`func (o *CloudCaptableRoundCloseRequest) UnsetCloseDate()`

UnsetCloseDate ensures that no value is present for CloseDate, not even an explicit nil
### GetId

`func (o *CloudCaptableRoundCloseRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudCaptableRoundCloseRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudCaptableRoundCloseRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudCaptableRoundCloseRequest) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


