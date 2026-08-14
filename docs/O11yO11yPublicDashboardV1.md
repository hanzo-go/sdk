# O11yO11yPublicDashboardV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | CreatedAt is when the dashboard was created. | [optional] 
**CreatedBy** | Pointer to **string** | CreatedBy is who created it. | [optional] 
**Data** | Pointer to **interface{}** |  | [optional] 
**Id** | Pointer to **string** | ID is the dashboard&#39;s id. | [optional] 
**Locked** | Pointer to **bool** | Locked reports whether the dashboard is locked. | [optional] 
**OrgId** | Pointer to **string** | OrgID is the org the dashboard belongs to. | [optional] 
**Source** | Pointer to **string** | Source is where the dashboard came from. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is when it last changed. | [optional] 
**UpdatedBy** | Pointer to **string** | UpdatedBy is who last changed it. | [optional] 

## Methods

### NewO11yO11yPublicDashboardV1

`func NewO11yO11yPublicDashboardV1() *O11yO11yPublicDashboardV1`

NewO11yO11yPublicDashboardV1 instantiates a new O11yO11yPublicDashboardV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yPublicDashboardV1WithDefaults

`func NewO11yO11yPublicDashboardV1WithDefaults() *O11yO11yPublicDashboardV1`

NewO11yO11yPublicDashboardV1WithDefaults instantiates a new O11yO11yPublicDashboardV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *O11yO11yPublicDashboardV1) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yO11yPublicDashboardV1) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yO11yPublicDashboardV1) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yO11yPublicDashboardV1) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *O11yO11yPublicDashboardV1) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *O11yO11yPublicDashboardV1) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *O11yO11yPublicDashboardV1) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *O11yO11yPublicDashboardV1) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetData

`func (o *O11yO11yPublicDashboardV1) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yPublicDashboardV1) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yPublicDashboardV1) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yPublicDashboardV1) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *O11yO11yPublicDashboardV1) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *O11yO11yPublicDashboardV1) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetId

`func (o *O11yO11yPublicDashboardV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yPublicDashboardV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yPublicDashboardV1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yPublicDashboardV1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocked

`func (o *O11yO11yPublicDashboardV1) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *O11yO11yPublicDashboardV1) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *O11yO11yPublicDashboardV1) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *O11yO11yPublicDashboardV1) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetOrgId

`func (o *O11yO11yPublicDashboardV1) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *O11yO11yPublicDashboardV1) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *O11yO11yPublicDashboardV1) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *O11yO11yPublicDashboardV1) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetSource

`func (o *O11yO11yPublicDashboardV1) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *O11yO11yPublicDashboardV1) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *O11yO11yPublicDashboardV1) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *O11yO11yPublicDashboardV1) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yO11yPublicDashboardV1) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yO11yPublicDashboardV1) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yO11yPublicDashboardV1) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yO11yPublicDashboardV1) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *O11yO11yPublicDashboardV1) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *O11yO11yPublicDashboardV1) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *O11yO11yPublicDashboardV1) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *O11yO11yPublicDashboardV1) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


