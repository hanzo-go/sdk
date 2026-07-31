# CloudAudience

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | CreatedAt and UpdatedAt are unix seconds, both server-assigned. | [optional] 
**Event** | Pointer to **string** | Event is the analytics event a member must have fired. EMPTY MEANS NO FILTER: the audience is then every mailable customer in the org, and no warehouse is consulted. | [optional] 
**Id** | Pointer to **string** | ID is the server-assigned audience id (\&quot;aud_\&quot; + 128 random bits). | [optional] 
**Name** | Pointer to **string** | Name is the audience&#39;s label. Required, trimmed, capped at 1024 bytes. | [optional] 
**UpdatedAt** | Pointer to **int32** |  | [optional] 
**WindowDays** | Pointer to **int32** | WindowDays is how far back the event counts, ending now. 0 means 30 and nothing above 3650 is honoured. Ignored when Event is empty. | [optional] 

## Methods

### NewCloudAudience

`func NewCloudAudience() *CloudAudience`

NewCloudAudience instantiates a new CloudAudience object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAudienceWithDefaults

`func NewCloudAudienceWithDefaults() *CloudAudience`

NewCloudAudienceWithDefaults instantiates a new CloudAudience object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CloudAudience) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudAudience) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudAudience) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudAudience) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEvent

`func (o *CloudAudience) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CloudAudience) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CloudAudience) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *CloudAudience) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetId

`func (o *CloudAudience) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudAudience) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudAudience) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudAudience) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CloudAudience) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudAudience) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudAudience) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudAudience) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudAudience) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudAudience) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudAudience) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudAudience) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWindowDays

`func (o *CloudAudience) GetWindowDays() int32`

GetWindowDays returns the WindowDays field if non-nil, zero value otherwise.

### GetWindowDaysOk

`func (o *CloudAudience) GetWindowDaysOk() (*int32, bool)`

GetWindowDaysOk returns a tuple with the WindowDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowDays

`func (o *CloudAudience) SetWindowDays(v int32)`

SetWindowDays sets WindowDays field to given value.

### HasWindowDays

`func (o *CloudAudience) HasWindowDays() bool`

HasWindowDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


