# CloudKycRefreshOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Formation** | Pointer to [**CloudFormation**](CloudFormation.md) | Formation is the org&#39;s incorporation record with each founder&#39;s reconciled status. | [optional] 
**Provider** | Pointer to **string** | Provider is the identity-verification provider that was consulted. | [optional] 

## Methods

### NewCloudKycRefreshOut

`func NewCloudKycRefreshOut() *CloudKycRefreshOut`

NewCloudKycRefreshOut instantiates a new CloudKycRefreshOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudKycRefreshOutWithDefaults

`func NewCloudKycRefreshOutWithDefaults() *CloudKycRefreshOut`

NewCloudKycRefreshOutWithDefaults instantiates a new CloudKycRefreshOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormation

`func (o *CloudKycRefreshOut) GetFormation() CloudFormation`

GetFormation returns the Formation field if non-nil, zero value otherwise.

### GetFormationOk

`func (o *CloudKycRefreshOut) GetFormationOk() (*CloudFormation, bool)`

GetFormationOk returns a tuple with the Formation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormation

`func (o *CloudKycRefreshOut) SetFormation(v CloudFormation)`

SetFormation sets Formation field to given value.

### HasFormation

`func (o *CloudKycRefreshOut) HasFormation() bool`

HasFormation returns a boolean if a field has been set.

### GetProvider

`func (o *CloudKycRefreshOut) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudKycRefreshOut) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudKycRefreshOut) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudKycRefreshOut) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


