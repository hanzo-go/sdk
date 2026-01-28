// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"
)

type UnionTime time.Time

func (UnionTime) ImplementsKeyListResponseKeysUserAPIKeyAuthExpiresUnion()                  {}
func (UnionTime) ImplementsKeyListResponseKeysLiteLlmDeletedVerificationTokenExpiresUnion() {}
func (UnionTime) ImplementsKeyBlockResponseExpiresUnion()                                   {}
func (UnionTime) ImplementsSpendListLogsResponseEndTimeUnion()                              {}
func (UnionTime) ImplementsSpendListLogsResponseStartTimeUnion()                            {}
func (UnionTime) ImplementsSpendListTagsResponseEndTimeUnion()                              {}
func (UnionTime) ImplementsSpendListTagsResponseStartTimeUnion()                            {}
func (UnionTime) ImplementsGlobalSpendListTagsResponseEndTimeUnion()                        {}
func (UnionTime) ImplementsGlobalSpendListTagsResponseStartTimeUnion()                      {}
func (UnionTime) ImplementsGlobalSpendGetReportResponseEndTimeUnion()                       {}
func (UnionTime) ImplementsGlobalSpendGetReportResponseStartTimeUnion()                     {}

type UnionString string

func (UnionString) ImplementsOpenAIDeploymentEmbedParamsCustomLlmProviderUnion() {}
func (UnionString) ImplementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentUnion() {
}
func (UnionString) ImplementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLUnion() {
}
func (UnionString) ImplementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLUnion() {
}
func (UnionString) ImplementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionAssistantMessageContentUnion() {
}
func (UnionString) ImplementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionToolMessageContentUnion() {
}
func (UnionString) ImplementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionSystemMessageContentUnion() {
}
func (UnionString) ImplementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionFunctionMessageContentUnion() {
}
func (UnionString) ImplementsOpenAIDeploymentChatCompleteParamsMessagesChatCompletionDeveloperMessageContentUnion() {
}
func (UnionString) ImplementsOpenAIDeploymentChatCompleteParamsFunctionCallUnion() {}
func (UnionString) ImplementsOpenAIDeploymentChatCompleteParamsStopUnion()         {}
func (UnionString) ImplementsOpenAIDeploymentChatCompleteParamsToolChoiceUnion()   {}
func (UnionString) ImplementsEngineEmbedParamsCustomLlmProviderUnion()             {}
func (UnionString) ImplementsEngineChatCompleteParamsMessagesChatCompletionUserMessageContentUnion() {
}
func (UnionString) ImplementsEngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLUnion() {
}
func (UnionString) ImplementsEngineChatCompleteParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLUnion() {
}
func (UnionString) ImplementsEngineChatCompleteParamsMessagesChatCompletionAssistantMessageContentUnion() {
}
func (UnionString) ImplementsEngineChatCompleteParamsMessagesChatCompletionToolMessageContentUnion() {
}
func (UnionString) ImplementsEngineChatCompleteParamsMessagesChatCompletionSystemMessageContentUnion() {
}
func (UnionString) ImplementsEngineChatCompleteParamsMessagesChatCompletionFunctionMessageContentUnion() {
}
func (UnionString) ImplementsEngineChatCompleteParamsMessagesChatCompletionDeveloperMessageContentUnion() {
}
func (UnionString) ImplementsEngineChatCompleteParamsFunctionCallUnion()                            {}
func (UnionString) ImplementsEngineChatCompleteParamsStopUnion()                                    {}
func (UnionString) ImplementsEngineChatCompleteParamsToolChoiceUnion()                              {}
func (UnionString) ImplementsChatCompletionNewParamsMessagesChatCompletionUserMessageContentUnion() {}
func (UnionString) ImplementsChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionImageObjectImageURLUnion() {
}
func (UnionString) ImplementsChatCompletionNewParamsMessagesChatCompletionUserMessageContentArrayChatCompletionVideoObjectVideoURLUnion() {
}
func (UnionString) ImplementsChatCompletionNewParamsMessagesChatCompletionAssistantMessageContentUnion() {
}
func (UnionString) ImplementsChatCompletionNewParamsMessagesChatCompletionToolMessageContentUnion() {}
func (UnionString) ImplementsChatCompletionNewParamsMessagesChatCompletionSystemMessageContentUnion() {
}
func (UnionString) ImplementsChatCompletionNewParamsMessagesChatCompletionFunctionMessageContentUnion() {
}
func (UnionString) ImplementsChatCompletionNewParamsMessagesChatCompletionDeveloperMessageContentUnion() {
}
func (UnionString) ImplementsChatCompletionNewParamsFunctionCallUnion()                        {}
func (UnionString) ImplementsChatCompletionNewParamsStopUnion()                                {}
func (UnionString) ImplementsChatCompletionNewParamsToolChoiceUnion()                          {}
func (UnionString) ImplementsEmbeddingNewParamsCustomLlmProviderUnion()                        {}
func (UnionString) ImplementsModelNewParamsLitellmParamsConfigurableClientsideAuthParamUnion() {}
func (UnionString) ImplementsModelNewParamsLitellmParamsMockResponseUnion()                    {}
func (UnionString) ImplementsModelNewParamsLitellmParamsStreamTimeoutUnion()                   {}
func (UnionString) ImplementsModelNewParamsLitellmParamsTimeoutUnion()                         {}
func (UnionString) ImplementsModelNewParamsLitellmParamsVertexCredentialsUnion()               {}
func (UnionString) ImplementsUpdateDeploymentLitellmParamsConfigurableClientsideAuthParamsUnionParam() {
}
func (UnionString) ImplementsUpdateDeploymentLitellmParamsMockResponseUnionParam()          {}
func (UnionString) ImplementsUpdateDeploymentLitellmParamsStreamTimeoutUnionParam()         {}
func (UnionString) ImplementsUpdateDeploymentLitellmParamsTimeoutUnionParam()               {}
func (UnionString) ImplementsUpdateDeploymentLitellmParamsVertexCredentialsUnionParam()     {}
func (UnionString) ImplementsLiteLlmFineTuningJobCreateHyperparametersBatchSizeUnionParam() {}
func (UnionString) ImplementsLiteLlmFineTuningJobCreateHyperparametersLearningRateMultiplierUnionParam() {
}
func (UnionString) ImplementsLiteLlmFineTuningJobCreateHyperparametersNEpochsUnionParam()     {}
func (UnionString) ImplementsGenerateKeyResponseRouterSettingsModelGroupAliasUnion()          {}
func (UnionString) ImplementsKeyListResponseKeysUnion()                                       {}
func (UnionString) ImplementsKeyListResponseKeysUserAPIKeyAuthExpiresUnion()                  {}
func (UnionString) ImplementsKeyListResponseKeysLiteLlmDeletedVerificationTokenExpiresUnion() {}
func (UnionString) ImplementsKeyBlockResponseExpiresUnion()                                   {}
func (UnionString) ImplementsKeyUpdateParamsRouterSettingsModelGroupAliasUnion()              {}
func (UnionString) ImplementsKeyGenerateParamsRouterSettingsModelGroupAliasUnion()            {}
func (UnionString) ImplementsRegenerateKeyRequestRouterSettingsModelGroupAliasUnionParam()    {}
func (UnionString) ImplementsUserNewResponseRouterSettingsModelGroupAliasUnion()              {}
func (UnionString) ImplementsTeamAddMemberResponseLitellmModelTableModelAliasesUnion()        {}
func (UnionString) ImplementsSpendListLogsResponseEndTimeUnion()                              {}
func (UnionString) ImplementsSpendListLogsResponseMessagesUnion()                             {}
func (UnionString) ImplementsSpendListLogsResponseResponseUnion()                             {}
func (UnionString) ImplementsSpendListLogsResponseStartTimeUnion()                            {}
func (UnionString) ImplementsSpendListTagsResponseEndTimeUnion()                              {}
func (UnionString) ImplementsSpendListTagsResponseMessagesUnion()                             {}
func (UnionString) ImplementsSpendListTagsResponseResponseUnion()                             {}
func (UnionString) ImplementsSpendListTagsResponseStartTimeUnion()                            {}
func (UnionString) ImplementsGlobalSpendListTagsResponseEndTimeUnion()                        {}
func (UnionString) ImplementsGlobalSpendListTagsResponseMessagesUnion()                       {}
func (UnionString) ImplementsGlobalSpendListTagsResponseResponseUnion()                       {}
func (UnionString) ImplementsGlobalSpendListTagsResponseStartTimeUnion()                      {}
func (UnionString) ImplementsGlobalSpendGetReportResponseEndTimeUnion()                       {}
func (UnionString) ImplementsGlobalSpendGetReportResponseMessagesUnion()                      {}
func (UnionString) ImplementsGlobalSpendGetReportResponseResponseUnion()                      {}
func (UnionString) ImplementsGlobalSpendGetReportResponseStartTimeUnion()                     {}

type UnionInt int64

func (UnionInt) ImplementsLiteLlmFineTuningJobCreateHyperparametersBatchSizeUnionParam() {}
func (UnionInt) ImplementsLiteLlmFineTuningJobCreateHyperparametersNEpochsUnionParam()   {}

type UnionFloat float64

func (UnionFloat) ImplementsModelNewParamsLitellmParamsStreamTimeoutUnion()        {}
func (UnionFloat) ImplementsModelNewParamsLitellmParamsTimeoutUnion()              {}
func (UnionFloat) ImplementsUpdateDeploymentLitellmParamsStreamTimeoutUnionParam() {}
func (UnionFloat) ImplementsUpdateDeploymentLitellmParamsTimeoutUnionParam()       {}
func (UnionFloat) ImplementsLiteLlmFineTuningJobCreateHyperparametersLearningRateMultiplierUnionParam() {
}
