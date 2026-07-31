# CloudVerificationTally

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByStatus** | Pointer to **map[string]int32** | ByStatus tallies the org&#39;s verifications by provider-reported status. | [optional] 
**Total** | Pointer to **int32** | Total is the sum over every status. | [optional] 

## Methods

### NewCloudVerificationTally

`func NewCloudVerificationTally() *CloudVerificationTally`

NewCloudVerificationTally instantiates a new CloudVerificationTally object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudVerificationTallyWithDefaults

`func NewCloudVerificationTallyWithDefaults() *CloudVerificationTally`

NewCloudVerificationTallyWithDefaults instantiates a new CloudVerificationTally object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByStatus

`func (o *CloudVerificationTally) GetByStatus() map[string]int32`

GetByStatus returns the ByStatus field if non-nil, zero value otherwise.

### GetByStatusOk

`func (o *CloudVerificationTally) GetByStatusOk() (*map[string]int32, bool)`

GetByStatusOk returns a tuple with the ByStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByStatus

`func (o *CloudVerificationTally) SetByStatus(v map[string]int32)`

SetByStatus sets ByStatus field to given value.

### HasByStatus

`func (o *CloudVerificationTally) HasByStatus() bool`

HasByStatus returns a boolean if a field has been set.

### GetTotal

`func (o *CloudVerificationTally) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CloudVerificationTally) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CloudVerificationTally) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CloudVerificationTally) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


