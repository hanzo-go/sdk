# CloudPlanVocab

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EngineFeatures** | Pointer to **[]string** | EngineFeatures are the inference-engine capabilities a license can grant: inference, embeddings, rerank, training, vision, audio, tools. | [optional] 
**Keys** | Pointer to **map[string]interface{}** | Keys maps every entitlement key to its descriptor — key, namespace, JSON type(s), nullability, unit, enum and title, as the schema declares them. | [optional] 
**Namespaces** | Pointer to **[]string** | Namespaces are the entitlement key namespaces: the prefix before the dot in \&quot;ai.tokens_per_min\&quot;. | [optional] 

## Methods

### NewCloudPlanVocab

`func NewCloudPlanVocab() *CloudPlanVocab`

NewCloudPlanVocab instantiates a new CloudPlanVocab object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPlanVocabWithDefaults

`func NewCloudPlanVocabWithDefaults() *CloudPlanVocab`

NewCloudPlanVocabWithDefaults instantiates a new CloudPlanVocab object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineFeatures

`func (o *CloudPlanVocab) GetEngineFeatures() []string`

GetEngineFeatures returns the EngineFeatures field if non-nil, zero value otherwise.

### GetEngineFeaturesOk

`func (o *CloudPlanVocab) GetEngineFeaturesOk() (*[]string, bool)`

GetEngineFeaturesOk returns a tuple with the EngineFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineFeatures

`func (o *CloudPlanVocab) SetEngineFeatures(v []string)`

SetEngineFeatures sets EngineFeatures field to given value.

### HasEngineFeatures

`func (o *CloudPlanVocab) HasEngineFeatures() bool`

HasEngineFeatures returns a boolean if a field has been set.

### GetKeys

`func (o *CloudPlanVocab) GetKeys() map[string]interface{}`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *CloudPlanVocab) GetKeysOk() (*map[string]interface{}, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *CloudPlanVocab) SetKeys(v map[string]interface{})`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *CloudPlanVocab) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetNamespaces

`func (o *CloudPlanVocab) GetNamespaces() []string`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *CloudPlanVocab) GetNamespacesOk() (*[]string, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *CloudPlanVocab) SetNamespaces(v []string)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *CloudPlanVocab) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


