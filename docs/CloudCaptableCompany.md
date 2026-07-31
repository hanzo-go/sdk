# CloudCaptableCompany

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | CreatedAt is when the company row was seeded, in unix milliseconds. | [optional] 
**Id** | Pointer to **string** | ID is the company id, which is the tenant&#39;s own org id. | [optional] 
**IncorporationCountry** | Pointer to **string** | IncorporationCountry is the ISO country the entity is incorporated in. | [optional] 
**IncorporationState** | Pointer to **string** | IncorporationState is the state or province of incorporation. | [optional] 
**IncorporationType** | Pointer to **string** | IncorporationType is the entity kind, e.g. LLC or C_CORP. | [optional] 
**Name** | Pointer to **string** | Name is the company&#39;s legal name. | [optional] 
**PublicId** | Pointer to **string** | PublicID is the company&#39;s shareable public identifier. | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is when the company row last changed, in unix milliseconds. | [optional] 

## Methods

### NewCloudCaptableCompany

`func NewCloudCaptableCompany() *CloudCaptableCompany`

NewCloudCaptableCompany instantiates a new CloudCaptableCompany object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCaptableCompanyWithDefaults

`func NewCloudCaptableCompanyWithDefaults() *CloudCaptableCompany`

NewCloudCaptableCompanyWithDefaults instantiates a new CloudCaptableCompany object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CloudCaptableCompany) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudCaptableCompany) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudCaptableCompany) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudCaptableCompany) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *CloudCaptableCompany) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudCaptableCompany) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudCaptableCompany) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudCaptableCompany) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncorporationCountry

`func (o *CloudCaptableCompany) GetIncorporationCountry() string`

GetIncorporationCountry returns the IncorporationCountry field if non-nil, zero value otherwise.

### GetIncorporationCountryOk

`func (o *CloudCaptableCompany) GetIncorporationCountryOk() (*string, bool)`

GetIncorporationCountryOk returns a tuple with the IncorporationCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncorporationCountry

`func (o *CloudCaptableCompany) SetIncorporationCountry(v string)`

SetIncorporationCountry sets IncorporationCountry field to given value.

### HasIncorporationCountry

`func (o *CloudCaptableCompany) HasIncorporationCountry() bool`

HasIncorporationCountry returns a boolean if a field has been set.

### GetIncorporationState

`func (o *CloudCaptableCompany) GetIncorporationState() string`

GetIncorporationState returns the IncorporationState field if non-nil, zero value otherwise.

### GetIncorporationStateOk

`func (o *CloudCaptableCompany) GetIncorporationStateOk() (*string, bool)`

GetIncorporationStateOk returns a tuple with the IncorporationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncorporationState

`func (o *CloudCaptableCompany) SetIncorporationState(v string)`

SetIncorporationState sets IncorporationState field to given value.

### HasIncorporationState

`func (o *CloudCaptableCompany) HasIncorporationState() bool`

HasIncorporationState returns a boolean if a field has been set.

### GetIncorporationType

`func (o *CloudCaptableCompany) GetIncorporationType() string`

GetIncorporationType returns the IncorporationType field if non-nil, zero value otherwise.

### GetIncorporationTypeOk

`func (o *CloudCaptableCompany) GetIncorporationTypeOk() (*string, bool)`

GetIncorporationTypeOk returns a tuple with the IncorporationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncorporationType

`func (o *CloudCaptableCompany) SetIncorporationType(v string)`

SetIncorporationType sets IncorporationType field to given value.

### HasIncorporationType

`func (o *CloudCaptableCompany) HasIncorporationType() bool`

HasIncorporationType returns a boolean if a field has been set.

### GetName

`func (o *CloudCaptableCompany) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCaptableCompany) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCaptableCompany) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudCaptableCompany) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicId

`func (o *CloudCaptableCompany) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *CloudCaptableCompany) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *CloudCaptableCompany) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *CloudCaptableCompany) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudCaptableCompany) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudCaptableCompany) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudCaptableCompany) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudCaptableCompany) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


