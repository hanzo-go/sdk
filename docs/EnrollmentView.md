# EnrollmentView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **string** | ExpiresAt is when the un-used token lapses, RFC 3339. | [optional] 
**Jwt** | Pointer to **string** | JWT is the one-time enrollment token the device presents ONCE to join the fabric (zt edge enroll / zt-edge-tunnel enroll). Spent or lapsed, it authenticates nothing; this surface stores it nowhere. | [optional] 

## Methods

### NewEnrollmentView

`func NewEnrollmentView() *EnrollmentView`

NewEnrollmentView instantiates a new EnrollmentView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentViewWithDefaults

`func NewEnrollmentViewWithDefaults() *EnrollmentView`

NewEnrollmentViewWithDefaults instantiates a new EnrollmentView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *EnrollmentView) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *EnrollmentView) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *EnrollmentView) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *EnrollmentView) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetJwt

`func (o *EnrollmentView) GetJwt() string`

GetJwt returns the Jwt field if non-nil, zero value otherwise.

### GetJwtOk

`func (o *EnrollmentView) GetJwtOk() (*string, bool)`

GetJwtOk returns a tuple with the Jwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwt

`func (o *EnrollmentView) SetJwt(v string)`

SetJwt sets Jwt field to given value.

### HasJwt

`func (o *EnrollmentView) HasJwt() bool`

HasJwt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


