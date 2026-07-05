# GuardGetAuditLog200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]GuardAuditEntry**](GuardAuditEntry.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 

## Methods

### NewGuardGetAuditLog200Response

`func NewGuardGetAuditLog200Response() *GuardGetAuditLog200Response`

NewGuardGetAuditLog200Response instantiates a new GuardGetAuditLog200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuardGetAuditLog200ResponseWithDefaults

`func NewGuardGetAuditLog200ResponseWithDefaults() *GuardGetAuditLog200Response`

NewGuardGetAuditLog200ResponseWithDefaults instantiates a new GuardGetAuditLog200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *GuardGetAuditLog200Response) GetEntries() []GuardAuditEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *GuardGetAuditLog200Response) GetEntriesOk() (*[]GuardAuditEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *GuardGetAuditLog200Response) SetEntries(v []GuardAuditEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *GuardGetAuditLog200Response) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetTotal

`func (o *GuardGetAuditLog200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GuardGetAuditLog200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GuardGetAuditLog200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GuardGetAuditLog200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetHasMore

`func (o *GuardGetAuditLog200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GuardGetAuditLog200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GuardGetAuditLog200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *GuardGetAuditLog200Response) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


