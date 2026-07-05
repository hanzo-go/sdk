# DidListProfiles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profiles** | Pointer to [**[]DidProfile**](DidProfile.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**NextCursor** | Pointer to **string** |  | [optional] 

## Methods

### NewDidListProfiles200Response

`func NewDidListProfiles200Response() *DidListProfiles200Response`

NewDidListProfiles200Response instantiates a new DidListProfiles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDidListProfiles200ResponseWithDefaults

`func NewDidListProfiles200ResponseWithDefaults() *DidListProfiles200Response`

NewDidListProfiles200ResponseWithDefaults instantiates a new DidListProfiles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfiles

`func (o *DidListProfiles200Response) GetProfiles() []DidProfile`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *DidListProfiles200Response) GetProfilesOk() (*[]DidProfile, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *DidListProfiles200Response) SetProfiles(v []DidProfile)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *DidListProfiles200Response) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetTotal

`func (o *DidListProfiles200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DidListProfiles200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DidListProfiles200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DidListProfiles200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetNextCursor

`func (o *DidListProfiles200Response) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *DidListProfiles200Response) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *DidListProfiles200Response) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *DidListProfiles200Response) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


