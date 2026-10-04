# RiskRiskAdoptIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address is one of YOUR organisation&#39;s own published values (GET /v1/risk/state reports them, and a search reports the one it fitted for you). An address your organisation has not published is NOT FOUND — including one another organisation published, because an address names a value and never authorises reading it. | [optional] 

## Methods

### NewRiskRiskAdoptIn

`func NewRiskRiskAdoptIn() *RiskRiskAdoptIn`

NewRiskRiskAdoptIn instantiates a new RiskRiskAdoptIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRiskAdoptInWithDefaults

`func NewRiskRiskAdoptInWithDefaults() *RiskRiskAdoptIn`

NewRiskRiskAdoptInWithDefaults instantiates a new RiskRiskAdoptIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *RiskRiskAdoptIn) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *RiskRiskAdoptIn) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *RiskRiskAdoptIn) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *RiskRiskAdoptIn) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


