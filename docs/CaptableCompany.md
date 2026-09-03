# CaptableCompany

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | CreatedAt is when the company row was seeded, in unix milliseconds. | [optional] 
**Id** | Pointer to **string** | ID is the company id, which is the tenant&#39;s own org id. | [optional] 
**IncorporationCountry** | Pointer to **string** | IncorporationCountry is the ISO country the entity is incorporated in. | [optional] 
**IncorporationState** | Pointer to **string** | IncorporationState is the state or province of incorporation. | [optional] 
**IncorporationType** | Pointer to **string** | IncorporationType is the entity kind, e.g. LLC or C_CORP. | [optional] 
**Name** | Pointer to **string** | Name is the company&#39;s legal name. | [optional] 
**PublicId** | Pointer to **string** | PublicID is the company&#39;s shareable public identifier. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is when the company row last changed, in unix milliseconds. | [optional] 

## Methods

### NewCaptableCompany

`func NewCaptableCompany() *CaptableCompany`

NewCaptableCompany instantiates a new CaptableCompany object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableCompanyWithDefaults

`func NewCaptableCompanyWithDefaults() *CaptableCompany`

NewCaptableCompanyWithDefaults instantiates a new CaptableCompany object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CaptableCompany) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CaptableCompany) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CaptableCompany) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CaptableCompany) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *CaptableCompany) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaptableCompany) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaptableCompany) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CaptableCompany) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncorporationCountry

`func (o *CaptableCompany) GetIncorporationCountry() string`

GetIncorporationCountry returns the IncorporationCountry field if non-nil, zero value otherwise.

### GetIncorporationCountryOk

`func (o *CaptableCompany) GetIncorporationCountryOk() (*string, bool)`

GetIncorporationCountryOk returns a tuple with the IncorporationCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncorporationCountry

`func (o *CaptableCompany) SetIncorporationCountry(v string)`

SetIncorporationCountry sets IncorporationCountry field to given value.

### HasIncorporationCountry

`func (o *CaptableCompany) HasIncorporationCountry() bool`

HasIncorporationCountry returns a boolean if a field has been set.

### GetIncorporationState

`func (o *CaptableCompany) GetIncorporationState() string`

GetIncorporationState returns the IncorporationState field if non-nil, zero value otherwise.

### GetIncorporationStateOk

`func (o *CaptableCompany) GetIncorporationStateOk() (*string, bool)`

GetIncorporationStateOk returns a tuple with the IncorporationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncorporationState

`func (o *CaptableCompany) SetIncorporationState(v string)`

SetIncorporationState sets IncorporationState field to given value.

### HasIncorporationState

`func (o *CaptableCompany) HasIncorporationState() bool`

HasIncorporationState returns a boolean if a field has been set.

### GetIncorporationType

`func (o *CaptableCompany) GetIncorporationType() string`

GetIncorporationType returns the IncorporationType field if non-nil, zero value otherwise.

### GetIncorporationTypeOk

`func (o *CaptableCompany) GetIncorporationTypeOk() (*string, bool)`

GetIncorporationTypeOk returns a tuple with the IncorporationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncorporationType

`func (o *CaptableCompany) SetIncorporationType(v string)`

SetIncorporationType sets IncorporationType field to given value.

### HasIncorporationType

`func (o *CaptableCompany) HasIncorporationType() bool`

HasIncorporationType returns a boolean if a field has been set.

### GetName

`func (o *CaptableCompany) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CaptableCompany) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CaptableCompany) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CaptableCompany) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicId

`func (o *CaptableCompany) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *CaptableCompany) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *CaptableCompany) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *CaptableCompany) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CaptableCompany) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CaptableCompany) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CaptableCompany) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CaptableCompany) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


