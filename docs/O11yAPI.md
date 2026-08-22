# \O11yAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentCheckIn**](O11yAPI.md#AgentCheckIn) | **Post** /v1/o11y/cloud_integrations/{cloud_provider}/accounts/check_in | Is the deployed agent&#39;s check-in — the path consistent with the account surface, reporting the agent&#39;s account and telemetry state so the connection can be tracked.
[**AgentCheckInDeprecated**](O11yAPI.md#AgentCheckInDeprecated) | **Post** /v1/o11y/cloud-integrations/{cloud_provider}/agent-check-in | Is the deployed agent&#39;s check-in on its original hyphenated path, kept for backward compatibility with agents already running.
[**AuthzCheck**](O11yAPI.md#AuthzCheck) | **Post** /v1/o11y/authz/check | Evaluates a batch of transactions — relation plus object — for the authenticated caller and answers each with its authorization verdict, in the order they were asked.
[**CloneDashboardV2**](O11yAPI.md#CloneDashboardV2) | **Post** /v1/o11y/dashboards/{id}/clone | Clones an existing v2-shape dashboard.
[**CreateAccount**](O11yAPI.md#CreateAccount) | **Post** /v1/o11y/cloud_integrations/{cloud_provider}/accounts | Connects a new cloud-integration account for the given provider from its posted config and credentials, answering with the account and the artifact the agent deploys to complete the connection.
[**CreateAuthDomain**](O11yAPI.md#CreateAuthDomain) | **Post** /v1/o11y/domains | Claims an email domain for the org and configures how its users sign in; the answer is the new domain&#39;s id.
[**CreateBulkInvite**](O11yAPI.md#CreateBulkInvite) | **Post** /v1/o11y/invite/bulk | Invites several people to the caller&#39;s org in one call, refusing the whole batch when any email repeats.
[**CreateChannel**](O11yAPI.md#CreateChannel) | **Post** /v1/o11y/channels | Creates a notification channel, answering with the stored channel.
[**CreateDashboardV2**](O11yAPI.md#CreateDashboardV2) | **Post** /v1/o11y/dashboards | Creates a dashboard in the v2 format that follows the Perses spec and answers with the stored dashboard.
[**CreateDashboardView**](O11yAPI.md#CreateDashboardView) | **Post** /v1/o11y/dashboard_views | Persists the calling user&#39;s dashboard-listing state (query, sort, order) as a named, reusable view shared across the org.
[**CreateDowntimeSchedule**](O11yAPI.md#CreateDowntimeSchedule) | **Post** /v1/o11y/downtime_schedules | Creates a planned maintenance window, answering with the stored schedule.
[**CreateIngestionKey**](O11yAPI.md#CreateIngestionKey) | **Post** /v1/o11y/gateway/ingestion_keys | Mints an ingestion key for the workspace, answering with the created key.
[**CreateIngestionKeyLimit**](O11yAPI.md#CreateIngestionKeyLimit) | **Post** /v1/o11y/gateway/ingestion_keys/{keyId}/limits | Sets a signal limit on an ingestion key, by key id, answering with the created limit.
[**CreateInvite**](O11yAPI.md#CreateInvite) | **Post** /v1/o11y/invite | Invites one person to the caller&#39;s org by email, with the role they will hold when they accept.
[**CreateLLMAnnotation**](O11yAPI.md#CreateLLMAnnotation) | **Post** /v1/o11y/llm/annotation | Adds a human annotation to a trace or observation, optionally in a review queue.
[**CreateLLMScore**](O11yAPI.md#CreateLLMScore) | **Post** /v1/o11y/llm/scores | Attaches an eval score or human-feedback signal to a trace or a single observation.
[**CreateMetricReductionRule**](O11yAPI.md#CreateMetricReductionRule) | **Post** /v1/o11y/metric_reduction_rules | Creates a volume-control rule for a metric and returns it with its id; a metric that already has a rule is refused.
[**CreateOrUpdateLLMPricingRules**](O11yAPI.md#CreateOrUpdateLLMPricingRules) | **Put** /v1/o11y/llm_pricing_rules | Writes the pricing-rule batch — the single write endpoint used by both the user and the Zeus sync job.
[**CreatePublicDashboard**](O11yAPI.md#CreatePublicDashboard) | **Post** /v1/o11y/dashboards/{id}/public | Creates the public-sharing config for a dashboard and enables public sharing, answering with the new share&#39;s id.
[**CreateResetPasswordToken**](O11yAPI.md#CreateResetPasswordToken) | **Put** /v1/o11y/users/{id}/reset_password_tokens | Creates or regenerates a user&#39;s reset-password token: a live token is returned as it is, an expired one is replaced.
[**CreateRole**](O11yAPI.md#CreateRole) | **Post** /v1/o11y/roles | Creates a custom role in the caller&#39;s org from a name, an optional description and the transaction groups it grants, answering the new role&#39;s id.
[**CreateRoutePolicy**](O11yAPI.md#CreateRoutePolicy) | **Post** /v1/o11y/route_policies | Creates a route policy, answering with the stored policy.
[**CreateRule**](O11yAPI.md#CreateRule) | **Post** /v1/o11y/rules | Creates a new alert rule and answers with the stored rule.
[**CreateServiceAccount**](O11yAPI.md#CreateServiceAccount) | **Post** /v1/o11y/service_accounts | Creates a service account in the caller&#39;s org, answering its id.
[**CreateServiceAccountKey**](O11yAPI.md#CreateServiceAccountKey) | **Post** /v1/o11y/service_accounts/{id}/keys | Mints an API key for a service account and answers the key&#39;s id and its secret — the one time the secret is ever shown.
[**CreateServiceAccountRole**](O11yAPI.md#CreateServiceAccountRole) | **Post** /v1/o11y/service_accounts/{id}/roles | Assigns a role, named by its id, to a service account.
[**CreateSessionByEmailPassword**](O11yAPI.md#CreateSessionByEmailPassword) | **Post** /v1/o11y/sessions/email_password | Signs a user in with email and password and answers with the session&#39;s token pair.
[**CreateSpanMapper**](O11yAPI.md#CreateSpanMapper) | **Post** /v1/o11y/span_mapper_groups/{groupId}/span_mappers | Adds a mapper to a group: which field context it reads, the move or copy it performs, and whether it is on.
[**CreateSpanMapperGroup**](O11yAPI.md#CreateSpanMapperGroup) | **Post** /v1/o11y/span_mapper_groups | Creates a mapping group: the name it is known by, the span and resource attributes whose presence selects a span into it, and whether it is on.
[**CreateTraceFunnel**](O11yAPI.md#CreateTraceFunnel) | **Post** /v1/o11y/trace-funnels/new | Creates an empty funnel with a name, answering the funnel it created.
[**CreateUser**](O11yAPI.md#CreateUser) | **Post** /v1/o11y/users | Creates a member of the caller&#39;s org in the pending-invite state and mails them their invitation; the answer is the new user&#39;s id.
[**DeleteAuthDomain**](O11yAPI.md#DeleteAuthDomain) | **Delete** /v1/o11y/domains/{id} | Releases an email domain and discards its SSO configuration, by id.
[**DeleteChannelByID**](O11yAPI.md#DeleteChannelByID) | **Delete** /v1/o11y/channels/{id} | Removes a notification channel, by id.
[**DeleteDashboardV2**](O11yAPI.md#DeleteDashboardV2) | **Delete** /v1/o11y/dashboards/{id} | Deletes a v2-shape dashboard along with its tag relations.
[**DeleteDashboardView**](O11yAPI.md#DeleteDashboardView) | **Delete** /v1/o11y/dashboard_views/{id} | Removes a saved view.
[**DeleteDowntimeScheduleByID**](O11yAPI.md#DeleteDowntimeScheduleByID) | **Delete** /v1/o11y/downtime_schedules/{id} | Removes a planned maintenance window, by id.
[**DeleteIngestionKey**](O11yAPI.md#DeleteIngestionKey) | **Delete** /v1/o11y/gateway/ingestion_keys/{keyId} | Removes an ingestion key, by id.
[**DeleteIngestionKeyLimit**](O11yAPI.md#DeleteIngestionKeyLimit) | **Delete** /v1/o11y/gateway/ingestion_keys/limits/{limitId} | Removes an ingestion key limit, by limit id.
[**DeleteLLMPricingRule**](O11yAPI.md#DeleteLLMPricingRule) | **Delete** /v1/o11y/llm_pricing_rules/{id} | Hard-deletes a pricing rule by id.
[**DeleteLLMScore**](O11yAPI.md#DeleteLLMScore) | **Delete** /v1/o11y/llm/score/{id} | Hard-deletes a score by id.
[**DeleteMetricReductionRuleByID**](O11yAPI.md#DeleteMetricReductionRuleByID) | **Delete** /v1/o11y/metric_reduction_rules/{id} | Deletes a volume-control rule by its id.
[**DeleteO11yExplorerViewsByViewid**](O11yAPI.md#DeleteO11yExplorerViewsByViewid) | **Delete** /v1/o11y/explorer/views/{viewId} | Deletes one saved explorer view by id.
[**DeleteO11yReviewsById**](O11yAPI.md#DeleteO11yReviewsById) | **Delete** /v1/o11y/reviews/{id} | Removes one review queue and every item in it.
[**DeleteO11ySentinelProjectsById**](O11yAPI.md#DeleteO11ySentinelProjectsById) | **Delete** /v1/o11y/sentinel/projects/{id} | Deletes one Sentry project of the caller&#39;s org.
[**DeletePublicDashboard**](O11yAPI.md#DeletePublicDashboard) | **Delete** /v1/o11y/dashboards/{id}/public | Deletes the public-sharing config and disables public sharing of a dashboard.
[**DeleteRole**](O11yAPI.md#DeleteRole) | **Delete** /v1/o11y/roles/{id} | Deletes a custom role.
[**DeleteRoutePolicyByID**](O11yAPI.md#DeleteRoutePolicyByID) | **Delete** /v1/o11y/route_policies/{id} | Removes a route policy, by id.
[**DeleteRuleByID**](O11yAPI.md#DeleteRuleByID) | **Delete** /v1/o11y/rules/{id} | Removes an alert rule, by id.
[**DeleteServiceAccount**](O11yAPI.md#DeleteServiceAccount) | **Delete** /v1/o11y/service_accounts/{id} | Deletes a service account and revokes every key it holds.
[**DeleteServiceAccountRole**](O11yAPI.md#DeleteServiceAccountRole) | **Delete** /v1/o11y/service_accounts/{id}/roles/{rid} | Removes a role from a service account.
[**DeleteSession**](O11yAPI.md#DeleteSession) | **Delete** /v1/o11y/sessions | Signs the calling session out, invalidating its tokens.
[**DeleteSpanMapper**](O11yAPI.md#DeleteSpanMapper) | **Delete** /v1/o11y/span_mapper_groups/{groupId}/span_mappers/{mapperId} | Deletes one mapper from a group.
[**DeleteSpanMapperGroup**](O11yAPI.md#DeleteSpanMapperGroup) | **Delete** /v1/o11y/span_mapper_groups/{groupId} | Deletes a mapping group and every mapper under it.
[**DeleteTraceFunnel**](O11yAPI.md#DeleteTraceFunnel) | **Delete** /v1/o11y/trace-funnels/{funnel_id} | Deletes a funnel.
[**DeleteUser**](O11yAPI.md#DeleteUser) | **Delete** /v1/o11y/users/{id} | Removes one org member, by user id.
[**DeleteUserDeprecated**](O11yAPI.md#DeleteUserDeprecated) | **Delete** /v1/o11y/user/{id} | Removes one org member, by user id.
[**DisconnectAccount**](O11yAPI.md#DisconnectAccount) | **Delete** /v1/o11y/cloud_integrations/{cloud_provider}/accounts/{id} | Tears down a connected account for the given provider, by id.
[**ForgotPassword**](O11yAPI.md#ForgotPassword) | **Post** /v1/o11y/factor_password/forgot | Starts the forgotten-password flow: the named user is mailed a reset link.
[**GetAccount**](O11yAPI.md#GetAccount) | **Get** /v1/o11y/cloud_integrations/{cloud_provider}/accounts/{id} | Returns one connected account for the given provider, by id.
[**GetAccountService**](O11yAPI.md#GetAccountService) | **Get** /v1/o11y/cloud_integrations/{cloud_provider}/accounts/{id}/services/{service_id} | Returns one service and its configuration for a connected account of the given provider, by account id and service id.
[**GetAlerts**](O11yAPI.md#GetAlerts) | **Get** /v1/o11y/alerts | Returns the org&#39;s current alerts.
[**GetAllRoutePolicies**](O11yAPI.md#GetAllRoutePolicies) | **Get** /v1/o11y/route_policies | Lists the org&#39;s route policies.
[**GetAuthDomain**](O11yAPI.md#GetAuthDomain) | **Get** /v1/o11y/domains/{id} | Returns one auth domain with its SSO configuration, by id.
[**GetChannelByID**](O11yAPI.md#GetChannelByID) | **Get** /v1/o11y/channels/{id} | Returns one notification channel, by id.
[**GetConnectionCredentials**](O11yAPI.md#GetConnectionCredentials) | **Get** /v1/o11y/cloud_integrations/{cloud_provider}/credentials | Returns the credentials the connecting agent needs to establish the cloud integration, for the given cloud provider.
[**GetDashboardV2**](O11yAPI.md#GetDashboardV2) | **Get** /v1/o11y/dashboards/{id} | Returns a v2-shape dashboard.
[**GetDowntimeScheduleByID**](O11yAPI.md#GetDowntimeScheduleByID) | **Get** /v1/o11y/downtime_schedules/{id} | Returns one planned maintenance window, by id.
[**GetDraftFunnelErrorTraces**](O11yAPI.md#GetDraftFunnelErrorTraces) | **Post** /v1/o11y/trace-funnels/analytics/error-traces | Returns the errored traces through a step transition of a funnel described inline.
[**GetDraftFunnelOverview**](O11yAPI.md#GetDraftFunnelOverview) | **Post** /v1/o11y/trace-funnels/analytics/overview | Returns the conversion overview of a funnel described inline.
[**GetDraftFunnelSlowTraces**](O11yAPI.md#GetDraftFunnelSlowTraces) | **Post** /v1/o11y/trace-funnels/analytics/slow-traces | Returns the slowest traces through a step transition of a funnel described inline.
[**GetDraftFunnelStepMetrics**](O11yAPI.md#GetDraftFunnelStepMetrics) | **Post** /v1/o11y/trace-funnels/analytics/steps | Returns the per-step metrics of a funnel described inline.
[**GetDraftFunnelStepOverview**](O11yAPI.md#GetDraftFunnelStepOverview) | **Post** /v1/o11y/trace-funnels/analytics/steps/overview | Returns the conversion between two steps of a funnel described inline.
[**GetFlamegraph**](O11yAPI.md#GetFlamegraph) | **Post** /v1/o11y/traces/{traceId}/flamegraph | Returns a trace&#39;s flamegraph: spans bucketed by depth level, each level ordered as it is drawn, around the selected span.
[**GetHosts**](O11yAPI.md#GetHosts) | **Get** /v1/o11y/zeus/hosts | Returns the deployment&#39;s host info from Zeus.
[**GetIngestionKeys**](O11yAPI.md#GetIngestionKeys) | **Get** /v1/o11y/gateway/ingestion_keys | Lists the workspace&#39;s ingestion keys, paginated.
[**GetIntegration**](O11yAPI.md#GetIntegration) | **Get** /v1/o11y/integrations/{integrationId} | Returns one integration&#39;s full detail — its overview, configuration steps, collected data and assets — together with its installation record when the org has installed it.
[**GetIntegrationConnectionStatus**](O11yAPI.md#GetIntegrationConnectionStatus) | **Get** /v1/o11y/integrations/{integrationId}/connection_status | Reports whether the integration&#39;s logs and metrics have been received over the lookback window, so the console can show a live connection state.
[**GetLLMPricingRule**](O11yAPI.md#GetLLMPricingRule) | **Get** /v1/o11y/llm_pricing_rules/{id} | Returns a single LLM pricing rule by id.
[**GetLLMScore**](O11yAPI.md#GetLLMScore) | **Get** /v1/o11y/llm/score/{id} | Returns a single score by id.
[**GetMetricAlerts**](O11yAPI.md#GetMetricAlerts) | **Get** /v1/o11y/metrics/alerts | Lists the alert rules that reference a metric.
[**GetMetricAttributes**](O11yAPI.md#GetMetricAttributes) | **Get** /v1/o11y/metrics/attributes | Returns one metric&#39;s attribute keys, each with its unique values and their count.
[**GetMetricDashboardsV2**](O11yAPI.md#GetMetricDashboardsV2) | **Get** /v1/o11y/metrics/dashboards | Lists the dashboard panels that reference a metric.
[**GetMetricHighlights**](O11yAPI.md#GetMetricHighlights) | **Get** /v1/o11y/metrics/highlights | Returns one metric&#39;s headline numbers: data points, total and active time series, and when it was last received.
[**GetMetricMetadata**](O11yAPI.md#GetMetricMetadata) | **Get** /v1/o11y/metrics/metadata | Returns one metric&#39;s metadata: description, type, unit, temporality and monotonicity.
[**GetMetricReductionRuleByID**](O11yAPI.md#GetMetricReductionRuleByID) | **Get** /v1/o11y/metric_reduction_rules/{id} | Returns one volume-control rule by its id.
[**GetMetricReductionRuleStats**](O11yAPI.md#GetMetricReductionRuleStats) | **Get** /v1/o11y/metric_reduction_rules/stats | Returns total ingested vs retained series and samples and the estimated monthly savings across all volume-control rules.
[**GetMetricReductionRuleTimeseries**](O11yAPI.md#GetMetricReductionRuleTimeseries) | **Get** /v1/o11y/metric_reduction_rules/timeseries | Returns ingested vs retained series over time across all volume-control rules, in hourly buckets, in the query-range time-series response shape.
[**GetMetricsOnboardingStatus**](O11yAPI.md#GetMetricsOnboardingStatus) | **Get** /v1/o11y/metrics/onboarding | Reports whether any non-O11y metrics have been ingested — the lightweight check onboarding polls.
[**GetMetricsStats**](O11yAPI.md#GetMetricsStats) | **Post** /v1/o11y/metrics/stats | Lists metrics with their sample and time-series counts for a time range — the volume view of the metrics explorer, pageable and sortable.
[**GetMetricsTreemap**](O11yAPI.md#GetMetricsTreemap) | **Post** /v1/o11y/metrics/treemap | Returns the proportional distribution of metrics by sample count or time-series count, as the entries of a treemap.
[**GetMyOrganization**](O11yAPI.md#GetMyOrganization) | **Get** /v1/o11y/orgs/me | Returns the caller&#39;s own organization.
[**GetMyServiceAccount**](O11yAPI.md#GetMyServiceAccount) | **Get** /v1/o11y/service_accounts/me | Returns the calling service account itself, with the roles it holds — the self-inspection read for a key-authenticated caller.
[**GetMyUser**](O11yAPI.md#GetMyUser) | **Get** /v1/o11y/users/me | Returns the calling user together with every role they hold.
[**GetMyUserDeprecated**](O11yAPI.md#GetMyUserDeprecated) | **Get** /v1/o11y/user/me | Returns the calling user with their single legacy role.
[**GetO11yAlertsLast**](O11yAPI.md#GetO11yAlertsLast) | **Get** /v1/o11y/alerts/last | Replay the alert records this process took
[**GetO11yAutocompleteAggregateAttributes**](O11yAPI.md#GetO11yAutocompleteAggregateAttributes) | **Get** /v1/o11y/autocomplete/aggregate_attributes | Lists the attributes usable as an aggregate target for the given telemetry and operator — what a filter builder offers after the aggregation is chosen.
[**GetO11yAutocompleteAttributeKeys**](O11yAPI.md#GetO11yAutocompleteAttributeKeys) | **Get** /v1/o11y/autocomplete/attribute_keys | Lists the attribute keys available for filtering the given telemetry, each with its data type and whether it is a materialized column.
[**GetO11yAutocompleteAttributeValues**](O11yAPI.md#GetO11yAutocompleteAttributeValues) | **Get** /v1/o11y/autocomplete/attribute_values | Lists the values one attribute key has taken — string, number and bool values in their own lists — for completing a filter.
[**GetO11yAvailability**](O11yAPI.md#GetO11yAvailability) | **Get** /v1/o11y/availability | Reports how much of the Hanzo fleet is up — the current per-service inventory plus an up-versus-reporting trend across the window.
[**GetO11yClustersAttributeKeys**](O11yAPI.md#GetO11yClustersAttributeKeys) | **Get** /v1/o11y/clusters/attribute_keys | Lists the metric attribute keys Kubernetes clusters report, for building cluster filters.
[**GetO11yClustersAttributeValues**](O11yAPI.md#GetO11yClustersAttributeValues) | **Get** /v1/o11y/clusters/attribute_values | Lists the values one cluster attribute key has taken, for building cluster filters.
[**GetO11yCompleteGoogle**](O11yAPI.md#GetO11yCompleteGoogle) | **Get** /v1/o11y/complete/google | Complete a Google sign-in
[**GetO11yCompleteOidc**](O11yAPI.md#GetO11yCompleteOidc) | **Get** /v1/o11y/complete/oidc | Complete a generic OIDC sign-in
[**GetO11yDaemonsetsAttributeKeys**](O11yAPI.md#GetO11yDaemonsetsAttributeKeys) | **Get** /v1/o11y/daemonsets/attribute_keys | Lists the metric attribute keys Kubernetes daemonsets report, for building daemonset filters.
[**GetO11yDaemonsetsAttributeValues**](O11yAPI.md#GetO11yDaemonsetsAttributeValues) | **Get** /v1/o11y/daemonsets/attribute_values | Lists the values one daemonset attribute key has taken, for building daemonset filters.
[**GetO11yDeploymentsAttributeKeys**](O11yAPI.md#GetO11yDeploymentsAttributeKeys) | **Get** /v1/o11y/deployments/attribute_keys | Lists the metric attribute keys Kubernetes deployments report, for building deployment filters.
[**GetO11yDeploymentsAttributeValues**](O11yAPI.md#GetO11yDeploymentsAttributeValues) | **Get** /v1/o11y/deployments/attribute_values | Lists the values one deployment attribute key has taken, for building deployment filters.
[**GetO11yDisks**](O11yAPI.md#GetO11yDisks) | **Get** /v1/o11y/disks | Lists the storage disks the datastore reports, with their names and types.
[**GetO11yErrorfromerrorid**](O11yAPI.md#GetO11yErrorfromerrorid) | **Get** /v1/o11y/errorFromErrorID | Returns one exception instance and the span it happened on, by its error id within a group at a timestamp.
[**GetO11yErrorfromgroupid**](O11yAPI.md#GetO11yErrorfromgroupid) | **Get** /v1/o11y/errorFromGroupID | Returns the representative exception instance of a group at a timestamp, and the span it happened on.
[**GetO11yErrortrackingIssues**](O11yAPI.md#GetO11yErrortrackingIssues) | **Get** /v1/o11y/errortracking/issues | Lists the caller&#39;s org&#39;s grouped error issues (by fingerprint) with status, level, counts and first/last-seen.
[**GetO11yErrortrackingIssuesById**](O11yAPI.md#GetO11yErrortrackingIssuesById) | **Get** /v1/o11y/errortracking/issues/{id} | Returns one grouped issue with its latest occurrence sample.
[**GetO11yExplorerViews**](O11yAPI.md#GetO11yExplorerViews) | **Get** /v1/o11y/explorer/views | Lists the caller&#39;s org&#39;s saved explorer views, optionally narrowed to one source page, name or category.
[**GetO11yExplorerViewsByViewid**](O11yAPI.md#GetO11yExplorerViewsByViewid) | **Get** /v1/o11y/explorer/views/{viewId} | Returns one saved explorer view by id.
[**GetO11yFeatures**](O11yAPI.md#GetO11yFeatures) | **Get** /v1/o11y/features | Returns the supported feature flags and their resolved values for the caller&#39;s org.
[**GetO11yFieldsKeys**](O11yAPI.md#GetO11yFieldsKeys) | **Get** /v1/o11y/fields/keys | Returns the telemetry field keys matching the selector — the signal&#39;s fields grouped by name, and whether the catalog is complete.
[**GetO11yFieldsValues**](O11yAPI.md#GetO11yFieldsValues) | **Get** /v1/o11y/fields/values | Returns the values one telemetry field has taken — string, bool, number and related values — and whether the value list is complete.
[**GetO11yFilterSuggestions**](O11yAPI.md#GetO11yFilterSuggestions) | **Get** /v1/o11y/filter_suggestions | Suggests attribute keys and example filter queries for the query builder, seeded by what the org&#39;s own telemetry carries.
[**GetO11yGlobalConfig**](O11yAPI.md#GetO11yGlobalConfig) | **Get** /v1/o11y/global/config | Returns the deployment&#39;s global configuration: its public endpoints and which identity providers are enabled.
[**GetO11yHealth**](O11yAPI.md#GetO11yHealth) | **Get** /v1/o11y/health | Reports service health.
[**GetO11yHealthz**](O11yAPI.md#GetO11yHealthz) | **Get** /v1/o11y/healthz | Health of the observability runtime&#39;s services
[**GetO11yHostsAttributeKeys**](O11yAPI.md#GetO11yHostsAttributeKeys) | **Get** /v1/o11y/hosts/attribute_keys | Lists the metric attribute keys hosts report, for building host filters — each with its data type and whether it is a materialized column.
[**GetO11yHostsAttributeValues**](O11yAPI.md#GetO11yHostsAttributeValues) | **Get** /v1/o11y/hosts/attribute_values | Lists the values one host attribute key has taken, for building host filters — string, number and bool values in their own lists.
[**GetO11yInfraMonitoringChecks**](O11yAPI.md#GetO11yInfraMonitoringChecks) | **Get** /v1/o11y/infra_monitoring/checks | Reports whether the metrics and attributes an infra-monitoring section needs are being received — for each collector receiver or processor involved, what is present and what is missing, with a user-facing message and a docs link per missing piece.
[**GetO11yInfraOnboardingK8sStatus**](O11yAPI.md#GetO11yInfraOnboardingK8sStatus) | **Get** /v1/o11y/infra_onboarding/k8s/status | Reports how far Kubernetes infra onboarding has progressed: which metric families have arrived and, per pod, which required metadata labels are present.
[**GetO11yJobsAttributeKeys**](O11yAPI.md#GetO11yJobsAttributeKeys) | **Get** /v1/o11y/jobs/attribute_keys | Lists the metric attribute keys Kubernetes jobs report, for building job filters.
[**GetO11yJobsAttributeValues**](O11yAPI.md#GetO11yJobsAttributeValues) | **Get** /v1/o11y/jobs/attribute_values | Lists the values one job attribute key has taken, for building job filters.
[**GetO11yLicenses**](O11yAPI.md#GetO11yLicenses) | **Get** /v1/o11y/licenses | Lists the org&#39;s licenses.
[**GetO11yLicensesActive**](O11yAPI.md#GetO11yLicensesActive) | **Get** /v1/o11y/licenses/active | Activates the enterprise license.
[**GetO11yLivez**](O11yAPI.md#GetO11yLivez) | **Get** /v1/o11y/livez | Liveness of the observability process
[**GetO11yLogin**](O11yAPI.md#GetO11yLogin) | **Get** /v1/o11y/login | Report why an observability sign-in did not complete
[**GetO11yLogs**](O11yAPI.md#GetO11yLogs) | **Get** /v1/o11y/logs | Returns the most recent log records in the query window, newest first — each record an open object carrying its nanosecond timestamp and whatever fields the record was ingested with.
[**GetO11yLogsAggregate**](O11yAPI.md#GetO11yLogsAggregate) | **Get** /v1/o11y/logs/aggregate | Returns the logs aggregate buckets for the query window.
[**GetO11yLogsFields**](O11yAPI.md#GetO11yLogsFields) | **Get** /v1/o11y/logs/fields | Returns the log field catalog: the fields already selected as indexed columns, and the interesting ones seen in the data that could be.
[**GetO11yLogsLivetail**](O11yAPI.md#GetO11yLogsLivetail) | **Get** /v1/o11y/logs/livetail | Follow log records as they arrive
[**GetO11yLogsPipelinesByVersion**](O11yAPI.md#GetO11yLogsPipelinesByVersion) | **Get** /v1/o11y/logs/pipelines/{version} | Returns the caller&#39;s org&#39;s log parsing pipelines at one config version — \&quot;latest\&quot; for the newest — along with that version&#39;s deployment record and the recent version history.
[**GetO11yLogsPromotePaths**](O11yAPI.md#GetO11yLogsPromotePaths) | **Get** /v1/o11y/logs/promote_paths | Lists the log body paths already promoted or indexed, with the indexes each carries.
[**GetO11yMetricMetricMetadata**](O11yAPI.md#GetO11yMetricMetricMetadata) | **Get** /v1/o11y/metric/metric_metadata | Serves the OLDER /metric/metric_metadata route.
[**GetO11yNamespacesAttributeKeys**](O11yAPI.md#GetO11yNamespacesAttributeKeys) | **Get** /v1/o11y/namespaces/attribute_keys | Lists the metric attribute keys Kubernetes namespaces report, for building namespace filters.
[**GetO11yNamespacesAttributeValues**](O11yAPI.md#GetO11yNamespacesAttributeValues) | **Get** /v1/o11y/namespaces/attribute_values | Lists the values one namespace attribute key has taken, for building namespace filters.
[**GetO11yNextpreverrorids**](O11yAPI.md#GetO11yNextpreverrorids) | **Get** /v1/o11y/nextPrevErrorIDs | Returns the ids of the exception instances immediately after and before a given one within its group — the paging cursor the error detail view walks.
[**GetO11yNodesAttributeKeys**](O11yAPI.md#GetO11yNodesAttributeKeys) | **Get** /v1/o11y/nodes/attribute_keys | Lists the metric attribute keys Kubernetes nodes report, for building node filters.
[**GetO11yNodesAttributeValues**](O11yAPI.md#GetO11yNodesAttributeValues) | **Get** /v1/o11y/nodes/attribute_values | Lists the values one node attribute key has taken, for building node filters.
[**GetO11yPodsAttributeKeys**](O11yAPI.md#GetO11yPodsAttributeKeys) | **Get** /v1/o11y/pods/attribute_keys | Lists the metric attribute keys Kubernetes pods report, for building pod filters.
[**GetO11yPodsAttributeValues**](O11yAPI.md#GetO11yPodsAttributeValues) | **Get** /v1/o11y/pods/attribute_values | Lists the values one pod attribute key has taken, for building pod filters.
[**GetO11yProcessesAttributeKeys**](O11yAPI.md#GetO11yProcessesAttributeKeys) | **Get** /v1/o11y/processes/attribute_keys | Lists the metric attribute keys processes report, for building process filters.
[**GetO11yProcessesAttributeValues**](O11yAPI.md#GetO11yProcessesAttributeValues) | **Get** /v1/o11y/processes/attribute_values | Lists the values one process attribute key has taken, for building process filters.
[**GetO11yProductMetrics**](O11yAPI.md#GetO11yProductMetrics) | **Get** /v1/o11y/product/metrics | Returns one product&#39;s RED series — request rate, errors, p50 and p95 latency — for the caller&#39;s org, plus that org&#39;s LLM usage rollup over the same window.
[**GetO11yPvcsAttributeKeys**](O11yAPI.md#GetO11yPvcsAttributeKeys) | **Get** /v1/o11y/pvcs/attribute_keys | Lists the metric attribute keys persistent volume claims report, for building volume filters.
[**GetO11yPvcsAttributeValues**](O11yAPI.md#GetO11yPvcsAttributeValues) | **Get** /v1/o11y/pvcs/attribute_values | Lists the values one persistent-volume-claim attribute key has taken, for building volume filters.
[**GetO11yQuery**](O11yAPI.md#GetO11yQuery) | **Get** /v1/o11y/query | Evaluates one instant PromQL query against the org&#39;s metrics and returns the result at a single point in time.
[**GetO11yQueryProgress**](O11yAPI.md#GetO11yQueryProgress) | **Get** /v1/o11y/query_progress | Watch one running query&#39;s progress
[**GetO11yQueryRange**](O11yAPI.md#GetO11yQueryRange) | **Get** /v1/o11y/query_range | Runs a Prometheus-style range query over metrics — the legacy read that predates the v5 querier — and returns the matrix, vector or scalar the query resolved to.
[**GetO11yReadyz**](O11yAPI.md#GetO11yReadyz) | **Get** /v1/o11y/readyz | Readiness of the observability runtime to serve
[**GetO11yReviews**](O11yAPI.md#GetO11yReviews) | **Get** /v1/o11y/reviews | Returns a page of the caller org&#39;s human-review queues, newest first, narrowed to the caller&#39;s project.
[**GetO11yReviewsById**](O11yAPI.md#GetO11yReviewsById) | **Get** /v1/o11y/reviews/{id} | Returns one review queue with its pending and completed counts and its first page of items.
[**GetO11yReviewsByIdItems**](O11yAPI.md#GetO11yReviewsByIdItems) | **Get** /v1/o11y/reviews/{id}/items | Returns a page of one review queue&#39;s items, newest first, optionally filtered to PENDING or COMPLETED.
[**GetO11ySentinelEventsById**](O11yAPI.md#GetO11ySentinelEventsById) | **Get** /v1/o11y/sentinel/events/{id} | Returns one captured error event of a project, by its id.
[**GetO11ySentinelIssues**](O11yAPI.md#GetO11ySentinelIssues) | **Get** /v1/o11y/sentinel/issues | Lists the caller&#39;s org&#39;s grouped error issues, optionally narrowed to one project and one time window, and filtered by status, level, environment, service, a free-text query and a sort.
[**GetO11ySentinelIssuesById**](O11yAPI.md#GetO11ySentinelIssuesById) | **Get** /v1/o11y/sentinel/issues/{id} | Returns one grouped issue of the caller&#39;s org with its latest occurrence sample.
[**GetO11ySentinelIssuesByIdEvents**](O11yAPI.md#GetO11ySentinelIssuesByIdEvents) | **Get** /v1/o11y/sentinel/issues/{id}/events | Lists one issue&#39;s captured occurrences, scoped to a project — a project is an isolation unit, so the caller declares which project&#39;s occurrences to read.
[**GetO11ySentinelLogs**](O11yAPI.md#GetO11ySentinelLogs) | **Get** /v1/o11y/sentinel/logs | Lists a project&#39;s captured error events, newest first, optionally narrowed to those whose message or exception text contains a search string.
[**GetO11ySentinelProjects**](O11yAPI.md#GetO11ySentinelProjects) | **Get** /v1/o11y/sentinel/projects | Lists the caller&#39;s org&#39;s Sentry projects, each with its freshly-derived DSN.
[**GetO11ySentinelProjectsById**](O11yAPI.md#GetO11ySentinelProjectsById) | **Get** /v1/o11y/sentinel/projects/{id} | Returns one Sentry project of the caller&#39;s org, DSN included.
[**GetO11ySentinelStats**](O11yAPI.md#GetO11ySentinelStats) | **Get** /v1/o11y/sentinel/stats | Returns a project&#39;s event-rate timeseries: one bucket per interval over the requested period, counting the events in it.
[**GetO11ySentinelTraces**](O11yAPI.md#GetO11ySentinelTraces) | **Get** /v1/o11y/sentinel/traces | Lists the traces a project&#39;s captured errors reference, each with how many errors landed on it, when they started and stopped, and the latest message seen — the entry point for \&quot;which requests are failing\&quot;.
[**GetO11ySentinelTracesById**](O11yAPI.md#GetO11ySentinelTracesById) | **Get** /v1/o11y/sentinel/traces/{id} | Returns one trace&#39;s captured errors for a project — every error event that carried the trace id, in the order the events plane holds them.
[**GetO11yServicesList**](O11yAPI.md#GetO11yServicesList) | **Get** /v1/o11y/services/list | Lists the name of every service the trace store holds, with no window applied — the complete catalog, for pickers and autocomplete.
[**GetO11ySessions**](O11yAPI.md#GetO11ySessions) | **Get** /v1/o11y/sessions | List the caller org&#39;s LLM sessions
[**GetO11ySettingsApdex**](O11yAPI.md#GetO11ySettingsApdex) | **Get** /v1/o11y/settings/apdex | Returns apdex settings for the named services.
[**GetO11ySettingsTtl**](O11yAPI.md#GetO11ySettingsTtl) | **Get** /v1/o11y/settings/ttl | Returns the org&#39;s current retention policy: default TTL, custom per-label rules, and cold-storage settings where configured.
[**GetO11yStatefulsetsAttributeKeys**](O11yAPI.md#GetO11yStatefulsetsAttributeKeys) | **Get** /v1/o11y/statefulsets/attribute_keys | Lists the metric attribute keys Kubernetes statefulsets report, for building statefulset filters.
[**GetO11yStatefulsetsAttributeValues**](O11yAPI.md#GetO11yStatefulsetsAttributeValues) | **Get** /v1/o11y/statefulsets/attribute_values | Lists the values one statefulset attribute key has taken, for building statefulset filters.
[**GetO11yStats**](O11yAPI.md#GetO11yStats) | **Get** /v1/o11y/stats | Returns the collected usage statistics for the caller&#39;s org, as the stats reporter aggregates them — a map whose keys are the reporter&#39;s own counter names.
[**GetO11yStatus**](O11yAPI.md#GetO11yStatus) | **Get** /v1/o11y/status | Reports whether a product&#39;s service is live: an in-cluster health probe with its measured latency, fused with the per-replica up inventory.
[**GetO11ySummary**](O11yAPI.md#GetO11ySummary) | **Get** /v1/o11y/summary | Reports whether the platform is up.
[**GetO11yTraces**](O11yAPI.md#GetO11yTraces) | **Get** /v1/o11y/traces | Lists the caller org&#39;s recent traces — one row per trace with its span count and wall-clock duration, most recently active first.
[**GetO11yUsage**](O11yAPI.md#GetO11yUsage) | **Get** /v1/o11y/usage | Returns ingestion usage counts bucketed over the requested window, optionally narrowed to one service.
[**GetO11yVersion**](O11yAPI.md#GetO11yVersion) | **Get** /v1/o11y/version | Reports the running build: its version, whether an enterprise edition is present (\&quot;N\&quot; in this build), and whether first-user setup has completed.
[**GetOrgPreference**](O11yAPI.md#GetOrgPreference) | **Get** /v1/o11y/org/preferences/{name} | Returns one org-scoped preference, by name.
[**GetOverallStateTransitions**](O11yAPI.md#GetOverallStateTransitions) | **Post** /v1/o11y/rules/{id}/history/overall_status | Returns the overall firing/inactive windows for a rule, for the posted query range.
[**GetPublicDashboard**](O11yAPI.md#GetPublicDashboard) | **Get** /v1/o11y/dashboards/{id}/public | Returns the public-sharing config for a dashboard.
[**GetPublicDashboardData**](O11yAPI.md#GetPublicDashboardData) | **Get** /v1/o11y/public/dashboards/{id} | Returns the sanitized dashboard data for public access — the read a shared dashboard&#39;s public page makes.
[**GetPublicDashboardWidgetQueryRange**](O11yAPI.md#GetPublicDashboardWidgetQueryRange) | **Get** /v1/o11y/public/dashboards/{id}/widgets/{idx}/query_range | Returns the query-range result for one widget of a public dashboard.
[**GetQuickFilters**](O11yAPI.md#GetQuickFilters) | **Get** /v1/o11y/orgs/me/filters | Returns the org&#39;s quick filters for every signal — the attribute shortlists its explorers offer as one-click filters.
[**GetResetPasswordToken**](O11yAPI.md#GetResetPasswordToken) | **Get** /v1/o11y/users/{id}/reset_password_tokens | Returns the reset-password token a user already has; absent one, the answer is a not-found rather than a fresh token.
[**GetResetPasswordTokenDeprecated**](O11yAPI.md#GetResetPasswordTokenDeprecated) | **Get** /v1/o11y/getResetPasswordToken/{id} | Returns a user&#39;s password-reset token, creating one if none is live.
[**GetRole**](O11yAPI.md#GetRole) | **Get** /v1/o11y/roles/{id} | Returns one role with the transaction groups it grants.
[**GetRolesByUserID**](O11yAPI.md#GetRolesByUserID) | **Get** /v1/o11y/users/{id}/roles | Returns every role one org member holds, by user id.
[**GetRoutePolicyByID**](O11yAPI.md#GetRoutePolicyByID) | **Get** /v1/o11y/route_policies/{id} | Returns one route policy, by id.
[**GetRuleByID**](O11yAPI.md#GetRuleByID) | **Get** /v1/o11y/rules/{id} | Returns one alert rule with its evaluation state, by id.
[**GetRuleHistoryFilterKeys**](O11yAPI.md#GetRuleHistoryFilterKeys) | **Get** /v1/o11y/rules/{id}/history/filter_keys | Returns the distinct label keys present in a rule&#39;s history entries over the selected range, for building history filters.
[**GetRuleHistoryFilterValues**](O11yAPI.md#GetRuleHistoryFilterValues) | **Get** /v1/o11y/rules/{id}/history/filter_values | Returns the distinct values a given label key has taken across a rule&#39;s history entries.
[**GetRuleHistoryOverallStatus**](O11yAPI.md#GetRuleHistoryOverallStatus) | **Get** /v1/o11y/rules/{id}/history/overall_status | Returns the overall firing/inactive intervals for a rule over the selected range.
[**GetRuleHistoryStats**](O11yAPI.md#GetRuleHistoryStats) | **Get** /v1/o11y/rules/{id}/history/stats | Returns trigger and resolution statistics for a rule over the selected time range, current window against the prior one.
[**GetRuleHistoryTimeline**](O11yAPI.md#GetRuleHistoryTimeline) | **Get** /v1/o11y/rules/{id}/history/timeline | Returns paginated timeline entries for a rule&#39;s state transitions, filterable by state and a label expression, cursor-paginated.
[**GetRuleHistoryTopContributors**](O11yAPI.md#GetRuleHistoryTopContributors) | **Get** /v1/o11y/rules/{id}/history/top_contributors | Returns the label combinations that contributed most to a rule firing over the selected range.
[**GetRuleStateHistory**](O11yAPI.md#GetRuleStateHistory) | **Post** /v1/o11y/rules/{id}/history/timeline | Returns a rule&#39;s state-transition timeline for the posted query range, each entry carrying its related-logs or related-traces link.
[**GetRuleStateHistoryTopContributors**](O11yAPI.md#GetRuleStateHistoryTopContributors) | **Post** /v1/o11y/rules/{id}/history/top_contributors | Returns the label combinations that contributed most to a rule firing, for the posted query range.
[**GetRuleStats**](O11yAPI.md#GetRuleStats) | **Post** /v1/o11y/rules/{id}/history/stats | Returns trigger and resolution statistics for a rule, current window against the prior one, for the posted query range.
[**GetService**](O11yAPI.md#GetService) | **Get** /v1/o11y/cloud_integrations/{cloud_provider}/services/{service_id} | Returns one service the given provider can collect from, by service id, optionally scoped to one cloud integration.
[**GetServiceAccount**](O11yAPI.md#GetServiceAccount) | **Get** /v1/o11y/service_accounts/{id} | Returns one service account with the roles it holds.
[**GetServiceAccountRoles**](O11yAPI.md#GetServiceAccountRoles) | **Get** /v1/o11y/service_accounts/{id}/roles | Lists the roles a service account holds.
[**GetSessionContext**](O11yAPI.md#GetSessionContext) | **Get** /v1/o11y/sessions/context | Tells a sign-in page what an email address can do: which orgs the address belongs to and, per org, which password and SSO routes are open to it.
[**GetSignalFilters**](O11yAPI.md#GetSignalFilters) | **Get** /v1/o11y/orgs/me/filters/{signal} | Returns the org&#39;s quick filters for one signal — traces, logs, metrics, exceptions or api_monitoring.
[**GetTraceAggregations**](O11yAPI.md#GetTraceAggregations) | **Post** /v1/o11y/traces/{traceId}/aggregations | Computes span aggregations over one trace — span count, duration or share of execution time — grouped by the resource field each aggregation names.
[**GetTraceFields**](O11yAPI.md#GetTraceFields) | **Get** /v1/o11y/traces/fields | Returns the trace field catalog: the span fields already selected as indexed columns, and the interesting ones seen in the data that could be.
[**GetTraceFunnel**](O11yAPI.md#GetTraceFunnel) | **Get** /v1/o11y/trace-funnels/{funnel_id} | Returns one funnel with its steps.
[**GetTraceFunnelErrorTraces**](O11yAPI.md#GetTraceFunnelErrorTraces) | **Post** /v1/o11y/trace-funnels/{funnel_id}/analytics/error-traces | Returns the errored traces through a step transition of a saved funnel — the entry point for \&quot;why is this step failing\&quot;.
[**GetTraceFunnelOverview**](O11yAPI.md#GetTraceFunnelOverview) | **Post** /v1/o11y/trace-funnels/{funnel_id}/analytics/overview | Returns a saved funnel&#39;s conversion overview over a window: how many entered, how many converted, the rate and the latency.
[**GetTraceFunnelSlowTraces**](O11yAPI.md#GetTraceFunnelSlowTraces) | **Post** /v1/o11y/trace-funnels/{funnel_id}/analytics/slow-traces | Returns the slowest traces through a step transition of a saved funnel — the entry point for \&quot;why is this step slow\&quot;.
[**GetTraceFunnelStepMetrics**](O11yAPI.md#GetTraceFunnelStepMetrics) | **Post** /v1/o11y/trace-funnels/{funnel_id}/analytics/steps | Returns a saved funnel&#39;s per-step metrics over a window — the counts and latencies at each step, in step order.
[**GetTraceFunnelStepOverview**](O11yAPI.md#GetTraceFunnelStepOverview) | **Post** /v1/o11y/trace-funnels/{funnel_id}/analytics/steps/overview | Returns the conversion between two named steps of a saved funnel — the step-to-step drill-down behind the overview.
[**GetUser**](O11yAPI.md#GetUser) | **Get** /v1/o11y/users/{id} | Returns one org member together with every role they hold, by user id.
[**GetUserDeprecated**](O11yAPI.md#GetUserDeprecated) | **Get** /v1/o11y/user/{id} | Returns one org member with their single legacy role, by user id.
[**GetUserPreference**](O11yAPI.md#GetUserPreference) | **Get** /v1/o11y/user/preferences/{name} | Returns one preference of the calling user, by name.
[**GetUsersByRoleID**](O11yAPI.md#GetUsersByRoleID) | **Get** /v1/o11y/roles/{id}/users | Returns every org member holding a role, by role id.
[**GetWaterfallV4**](O11yAPI.md#GetWaterfallV4) | **Post** /v1/o11y/traces/{traceId}/waterfall | Returns a trace&#39;s waterfall: every span when the trace is small enough, a capped window around the selected span when it is not, with the uncollapsed subtrees the caller asked to keep open.
[**InspectMetrics**](O11yAPI.md#InspectMetrics) | **Post** /v1/o11y/metrics/inspect | Returns one metric&#39;s raw time series over a window of at most thirty minutes — each series with its labels and timestamp/value pairs.
[**InstallIntegration**](O11yAPI.md#InstallIntegration) | **Post** /v1/o11y/integrations/install | Installs an integration into the caller&#39;s org from its id and configuration, answering with the installed catalog item.
[**ListAccountServicesMetadata**](O11yAPI.md#ListAccountServicesMetadata) | **Get** /v1/o11y/cloud_integrations/{cloud_provider}/accounts/{id}/services | Lists the services metadata for one connected account of the given provider, by account id.
[**ListAccounts**](O11yAPI.md#ListAccounts) | **Get** /v1/o11y/cloud_integrations/{cloud_provider}/accounts | Lists the cloud-integration accounts connected for the given provider.
[**ListAuthDomains**](O11yAPI.md#ListAuthDomains) | **Get** /v1/o11y/domains | Lists the org&#39;s auth domains — the email domains whose SSO configuration this org owns.
[**ListChannels**](O11yAPI.md#ListChannels) | **Get** /v1/o11y/channels | Lists the org&#39;s notification channels.
[**ListDashboardViews**](O11yAPI.md#ListDashboardViews) | **Get** /v1/o11y/dashboard_views | Returns every saved view in the calling user&#39;s org.
[**ListDashboardsForUserV2**](O11yAPI.md#ListDashboardsForUserV2) | **Get** /v1/o11y/users/me/dashboards | Is dashboardListV2 personalized for the calling user: each dashboard carries the caller&#39;s pinned state, and pinned dashboards float to the top of the requested ordering.
[**ListDashboardsV2**](O11yAPI.md#ListDashboardsV2) | **Get** /v1/o11y/dashboards | Returns a page of v2-shape dashboards for the org.
[**ListDowntimeSchedules**](O11yAPI.md#ListDowntimeSchedules) | **Get** /v1/o11y/downtime_schedules | Lists all planned maintenance windows, optionally narrowed to the active ones or the recurring ones.
[**ListIntegrations**](O11yAPI.md#ListIntegrations) | **Get** /v1/o11y/integrations | Lists the available integrations and whether each is installed in the caller&#39;s org, optionally narrowed to installed or not-installed.
[**ListLLMAnnotations**](O11yAPI.md#ListLLMAnnotations) | **Get** /v1/o11y/llm/annotation | Lists human annotations on traces and observations, optionally scoped to one review queue.
[**ListLLMObservations**](O11yAPI.md#ListLLMObservations) | **Get** /v1/o11y/llm/observations | Lists gen_ai spans as LLM observations — each an LLM call with its model, token counts, cost and latency projected from gen_ai.* attributes, newest first, over the query window.
[**ListLLMPricingRules**](O11yAPI.md#ListLLMPricingRules) | **Get** /v1/o11y/llm_pricing_rules | Returns the LLM pricing rules for the caller&#39;s org, with pagination and an optional search and override filter.
[**ListLLMScores**](O11yAPI.md#ListLLMScores) | **Get** /v1/o11y/llm/scores | Lists eval scores and human-feedback signals attached to traces and observations, newest first.
[**ListLLMSessions**](O11yAPI.md#ListLLMSessions) | **Get** /v1/o11y/llm/sessions | Lists conversations — gen_ai spans grouped by session.id, with their trace and observation counts, tokens and cost.
[**ListLLMTraces**](O11yAPI.md#ListLLMTraces) | **Get** /v1/o11y/llm/traces | Lists LLM traces — gen_ai spans grouped by trace_id, with cost, tokens and latency rolled up across each trace.
[**ListLLMUsers**](O11yAPI.md#ListLLMUsers) | **Get** /v1/o11y/llm/users | Lists end users — gen_ai spans grouped by user.id, with their session, trace and observation counts, tokens and cost.
[**ListMetricReductionRules**](O11yAPI.md#ListMetricReductionRules) | **Get** /v1/o11y/metric_reduction_rules | Lists the org&#39;s metric volume-control (label reduction) rules, pageable and sortable by name, volume or recency.
[**ListMetrics**](O11yAPI.md#ListMetrics) | **Get** /v1/o11y/metrics | Lists the distinct metric names seen in a time range, each with its description, type, unit, temporality and monotonicity.
[**ListOrgPreferences**](O11yAPI.md#ListOrgPreferences) | **Get** /v1/o11y/org/preferences | Lists every org-scoped preference, each with its current and default value.
[**ListRoles**](O11yAPI.md#ListRoles) | **Get** /v1/o11y/roles | Lists every role in the caller&#39;s org — the managed ones the platform seeds and the custom ones its admins created.
[**ListRules**](O11yAPI.md#ListRules) | **Get** /v1/o11y/rules | Lists all alert rules with their current evaluation state.
[**ListServiceAccountKeys**](O11yAPI.md#ListServiceAccountKeys) | **Get** /v1/o11y/service_accounts/{id}/keys | Lists a service account&#39;s API keys — metadata only, never the secrets.
[**ListServiceAccounts**](O11yAPI.md#ListServiceAccounts) | **Get** /v1/o11y/service_accounts | Lists the caller&#39;s org&#39;s service accounts.
[**ListServicesMetadata**](O11yAPI.md#ListServicesMetadata) | **Get** /v1/o11y/cloud_integrations/{cloud_provider}/services | Lists the services the given provider can collect from, optionally scoped to one cloud integration.
[**ListSpanMapperGroups**](O11yAPI.md#ListSpanMapperGroups) | **Get** /v1/o11y/span_mapper_groups | Lists the caller&#39;s org&#39;s mapping groups, optionally only the enabled ones.
[**ListSpanMappers**](O11yAPI.md#ListSpanMappers) | **Get** /v1/o11y/span_mapper_groups/{groupId}/span_mappers | Lists the mappers belonging to one group, in the order they are applied.
[**ListTraceFunnels**](O11yAPI.md#ListTraceFunnels) | **Get** /v1/o11y/trace-funnels/list | Lists the caller&#39;s org&#39;s funnels, each with its steps and who last touched it.
[**ListUserPreferences**](O11yAPI.md#ListUserPreferences) | **Get** /v1/o11y/user/preferences | Lists every preference of the calling user, each with its current and default value.
[**ListUsers**](O11yAPI.md#ListUsers) | **Get** /v1/o11y/users | Lists the caller&#39;s org members.
[**ListUsersDeprecated**](O11yAPI.md#ListUsersDeprecated) | **Get** /v1/o11y/user | Lists the org&#39;s members with their single legacy role.
[**LockDashboardV2**](O11yAPI.md#LockDashboardV2) | **Put** /v1/o11y/dashboards/{id}/lock | Locks a v2-shape dashboard.
[**PatchDashboardV2**](O11yAPI.md#PatchDashboardV2) | **Patch** /v1/o11y/dashboards/{id} | Applies an RFC 6902 JSON Patch to a v2-shape dashboard.
[**PatchO11yReviewsById**](O11yAPI.md#PatchO11yReviewsById) | **Patch** /v1/o11y/reviews/{id} | Changes a review queue&#39;s name, description or score-config set.
[**PatchO11yReviewsByIdItemsByItemid**](O11yAPI.md#PatchO11yReviewsByIdItemsByItemid) | **Patch** /v1/o11y/reviews/{id}/items/{itemId} | Moves one queue item between PENDING and COMPLETED and sets its assignee.
[**PatchRuleByID**](O11yAPI.md#PatchRuleByID) | **Patch** /v1/o11y/rules/{id} | Applies a partial update to an alert rule, by id, answering with the stored rule — the common toggle for enabling or muting a rule.
[**PinDashboardV2**](O11yAPI.md#PinDashboardV2) | **Put** /v1/o11y/users/me/dashboards/{id}/pins | Pins a dashboard for the calling user.
[**PostO11yAlertsByReceiver**](O11yAPI.md#PostO11yAlertsByReceiver) | **Post** /v1/o11y/alerts/{receiver} | Take an Alertmanager notification and page a human
[**PostO11yApiByProjectIdEnvelope**](O11yAPI.md#PostO11yApiByProjectIdEnvelope) | **Post** /v1/o11y/api/{project_id}/envelope/ | Receive a Sentry envelope on the SDK&#39;s own DSN path
[**PostO11yApiByProjectIdStore**](O11yAPI.md#PostO11yApiByProjectIdStore) | **Post** /v1/o11y/api/{project_id}/store/ | Receive a single Sentry event on the SDK&#39;s own DSN path
[**PostO11yAutoCompleteAttributeValues**](O11yAPI.md#PostO11yAutoCompleteAttributeValues) | **Post** /v1/o11y/auto_complete/attribute_values | Reads the attribute-value request from the body rather than off the query string — the spelling the newer builder uses to send its filters alongside the request.
[**PostO11yClustersList**](O11yAPI.md#PostO11yClustersList) | **Post** /v1/o11y/clusters/list | Lists Kubernetes clusters over a time range, each with its CPU and memory usage against allocatable capacity and its attributes; filterable, groupable and paginated.
[**PostO11yCompleteSaml**](O11yAPI.md#PostO11yCompleteSaml) | **Post** /v1/o11y/complete/saml | Complete a SAML sign-in
[**PostO11yCounterrors**](O11yAPI.md#PostO11yCounterrors) | **Post** /v1/o11y/countErrors | Counts the grouped exceptions in the query window for the caller&#39;s org.
[**PostO11yDaemonsetsList**](O11yAPI.md#PostO11yDaemonsetsList) | **Post** /v1/o11y/daemonsets/list | Lists Kubernetes daemonsets over a time range, each with the CPU and memory its pods used against request and limit, desired and available node counts, restarts and attributes; filterable, groupable and paginated.
[**PostO11yDependencyGraph**](O11yAPI.md#PostO11yDependencyGraph) | **Post** /v1/o11y/dependency_graph | Returns the service dependency graph over the requested window: every parent→child edge observed, with call and error rates and latency percentiles per edge.
[**PostO11yDeploymentsList**](O11yAPI.md#PostO11yDeploymentsList) | **Post** /v1/o11y/deployments/list | Lists Kubernetes deployments over a time range, each with the CPU and memory its pods used against request and limit, desired and available replica counts, restarts and attributes; filterable, groupable and paginated.
[**PostO11yErrortrackingIssuesById**](O11yAPI.md#PostO11yErrortrackingIssuesById) | **Post** /v1/o11y/errortracking/issues/{id} | Changes an issue&#39;s lifecycle — resolve, ignore, reopen or assign — and returns the updated issue.
[**PostO11yEvent**](O11yAPI.md#PostO11yEvent) | **Post** /v1/o11y/event | Records one product-analytics event for the signed-in user — a track event with a name and free-form attributes.
[**PostO11yExplorerViews**](O11yAPI.md#PostO11yExplorerViews) | **Post** /v1/o11y/explorer/views | Saves a new explorer view for the caller&#39;s org and returns its id.
[**PostO11yExportRawData**](O11yAPI.md#PostO11yExportRawData) | **Post** /v1/o11y/export_raw_data | Export raw telemetry rows as a file
[**PostO11yHostsList**](O11yAPI.md#PostO11yHostsList) | **Post** /v1/o11y/hosts/list | Lists monitored hosts over a time range, each with its CPU, memory, I/O wait and 15-minute load, whether it is actively reporting, its OS and its attributes; filterable, groupable and paginated.
[**PostO11yInfraMonitoringClusters**](O11yAPI.md#PostO11yInfraMonitoringClusters) | **Post** /v1/o11y/infra_monitoring/clusters | Lists Kubernetes clusters with CPU and memory usage against allocatable capacity summed over their nodes, plus per-group node readiness and pod phase counts.
[**PostO11yInfraMonitoringDaemonsets**](O11yAPI.md#PostO11yInfraMonitoringDaemonsets) | **Post** /v1/o11y/infra_monitoring/daemonsets | Lists Kubernetes daemonsets with the CPU and memory their pods used against request and limit, the latest desired and current scheduled NODE counts (node counts, not pod counts), and per-group pod phase counts.
[**PostO11yInfraMonitoringDeployments**](O11yAPI.md#PostO11yInfraMonitoringDeployments) | **Post** /v1/o11y/infra_monitoring/deployments | Lists Kubernetes deployments with the CPU and memory their pods used against request and limit, the latest desired and available replica counts, and per-group pod phase counts.
[**PostO11yInfraMonitoringHosts**](O11yAPI.md#PostO11yInfraMonitoringHosts) | **Post** /v1/o11y/infra_monitoring/hosts | Lists hosts with key infrastructure metrics — CPU, memory, I/O wait and disk usage percentages and 15-minute load — plus an active/inactive status from whether the host reported in the last ten minutes.
[**PostO11yInfraMonitoringJobs**](O11yAPI.md#PostO11yInfraMonitoringJobs) | **Post** /v1/o11y/infra_monitoring/jobs | Lists Kubernetes jobs with the CPU and memory their pods used against request and limit, the latest desired-successful, active, failed and successful pod counters, and per-group pod phase counts — the phase counts are current state while the counters are cumulative over the job&#39;s life.
[**PostO11yInfraMonitoringNamespaces**](O11yAPI.md#PostO11yInfraMonitoringNamespaces) | **Post** /v1/o11y/infra_monitoring/namespaces | Lists Kubernetes namespaces with the CPU and memory their pods used and per-group pod phase counts.
[**PostO11yInfraMonitoringNodes**](O11yAPI.md#PostO11yInfraMonitoringNodes) | **Post** /v1/o11y/infra_monitoring/nodes | Lists Kubernetes nodes with CPU and memory usage against allocatable capacity, per-group readiness counts and per-group phase counts for the pods scheduled on them.
[**PostO11yInfraMonitoringPods**](O11yAPI.md#PostO11yInfraMonitoringPods) | **Post** /v1/o11y/infra_monitoring/pods | Lists Kubernetes pods with CPU and memory usage against request and limit, the pod&#39;s phase and its age, plus its namespace, node, owning workload and cluster attributes.
[**PostO11yInfraMonitoringPvcs**](O11yAPI.md#PostO11yInfraMonitoringPvcs) | **Post** /v1/o11y/infra_monitoring/pvcs | Lists Kubernetes persistent volume claims with available, capacity and used bytes and inode counts, plus the claim&#39;s pod, namespace, node, statefulset and cluster attributes.
[**PostO11yInfraMonitoringStatefulsets**](O11yAPI.md#PostO11yInfraMonitoringStatefulsets) | **Post** /v1/o11y/infra_monitoring/statefulsets | Lists Kubernetes statefulsets with the CPU and memory their pods used against request and limit, the latest desired and current replica counts, and per-group pod phase counts.
[**PostO11yJobsList**](O11yAPI.md#PostO11yJobsList) | **Post** /v1/o11y/jobs/list | Lists Kubernetes jobs over a time range, each with the CPU and memory its pods used against request and limit, desired-successful, active, failed and successful pod counts, restarts and attributes; filterable, groupable and paginated.
[**PostO11yListerrors**](O11yAPI.md#PostO11yListerrors) | **Post** /v1/o11y/listErrors | Lists the grouped exceptions in the query window — each an exception type with its message, count, service and first/last-seen — for the caller&#39;s org.
[**PostO11yLogsFields**](O11yAPI.md#PostO11yLogsFields) | **Post** /v1/o11y/logs/fields | Changes how one log field is stored — selects or deselects it as a materialized column and tunes its index — and echoes the setting back.
[**PostO11yLogsPipelines**](O11yAPI.md#PostO11yLogsPipelines) | **Post** /v1/o11y/logs/pipelines | Saves the given log parsing pipelines as the new config version for the caller&#39;s org and starts deploying it.
[**PostO11yLogsPipelinesPreview**](O11yAPI.md#PostO11yLogsPipelinesPreview) | **Post** /v1/o11y/logs/pipelines/preview | Runs the given log parsing pipelines over the given sample records without saving anything, and returns the transformed records plus whatever the collector logged while simulating them.
[**PostO11yLogsPromotePaths**](O11yAPI.md#PostO11yLogsPromotePaths) | **Post** /v1/o11y/logs/promote_paths | Promotes and indexes log body paths: each named path is lifted out of the JSON body into its own column, with the indexes the caller asked for.
[**PostO11yMessagingQueuesKafkaConsumerLagConsumerDetails**](O11yAPI.md#PostO11yMessagingQueuesKafkaConsumerLagConsumerDetails) | **Post** /v1/o11y/messaging-queues/kafka/consumer-lag/consumer-details | Returns the consumer side of a consumer-lag view: the consumer groups reading the topic/partition named in variables, with their throughput and latency over the window.
[**PostO11yMessagingQueuesKafkaConsumerLagNetworkLatency**](O11yAPI.md#PostO11yMessagingQueuesKafkaConsumerLagNetworkLatency) | **Post** /v1/o11y/messaging-queues/kafka/consumer-lag/network-latency | Returns consumer network latency correlated per client: a throughput pass over the window finds the consumer clients, then their fetch latency joins in as a latency column per client/instance/service.
[**PostO11yMessagingQueuesKafkaConsumerLagProducerDetails**](O11yAPI.md#PostO11yMessagingQueuesKafkaConsumerLagProducerDetails) | **Post** /v1/o11y/messaging-queues/kafka/consumer-lag/producer-details | Returns the producer side of a consumer-lag view: the producers writing to the topic/partition named in variables, with their throughput and latency over the window.
[**PostO11yMessagingQueuesKafkaOnboardingConsumers**](O11yAPI.md#PostO11yMessagingQueuesKafkaOnboardingConsumers) | **Post** /v1/o11y/messaging-queues/kafka/onboarding/consumers | Checks whether the spans the Kafka consumer views need are arriving, row for row like producersOnboarding.
[**PostO11yMessagingQueuesKafkaOnboardingKafka**](O11yAPI.md#PostO11yMessagingQueuesKafkaOnboardingKafka) | **Post** /v1/o11y/messaging-queues/kafka/onboarding/kafka | Checks whether Kafka&#39;s own metrics — consumer lag and partition telemetry — are arriving, so the lag views can be lit up.
[**PostO11yMessagingQueuesKafkaOnboardingProducers**](O11yAPI.md#PostO11yMessagingQueuesKafkaOnboardingProducers) | **Post** /v1/o11y/messaging-queues/kafka/onboarding/producers | Checks whether the spans the Kafka producer views need are arriving — one row per required span attribute, with a pass/fail status and, on failure, what is missing from the instrumentation.
[**PostO11yMessagingQueuesKafkaPartitionLatencyConsumer**](O11yAPI.md#PostO11yMessagingQueuesKafkaPartitionLatencyConsumer) | **Post** /v1/o11y/messaging-queues/kafka/partition-latency/consumer | Returns the consumer-group latency detail for the topic and partition named in the request&#39;s variables.
[**PostO11yMessagingQueuesKafkaPartitionLatencyOverview**](O11yAPI.md#PostO11yMessagingQueuesKafkaPartitionLatencyOverview) | **Post** /v1/o11y/messaging-queues/kafka/partition-latency/overview | Returns the per-partition latency overview for the window — each topic/partition with its throughput and latency profile.
[**PostO11yMessagingQueuesKafkaSpanEvaluation**](O11yAPI.md#PostO11yMessagingQueuesKafkaSpanEvaluation) | **Post** /v1/o11y/messaging-queues/kafka/span/evaluation | Correlates producer and consumer spans over the evaluation window (eval_time bounds the scan) and returns the pairings with their end-to-end delay — the check that messages produced are being consumed.
[**PostO11yMessagingQueuesKafkaTopicThroughputConsumer**](O11yAPI.md#PostO11yMessagingQueuesKafkaTopicThroughputConsumer) | **Post** /v1/o11y/messaging-queues/kafka/topic-throughput/consumer | Returns the consumer topic-throughput overview for the window — what each consumer group read, per topic.
[**PostO11yMessagingQueuesKafkaTopicThroughputConsumerDetails**](O11yAPI.md#PostO11yMessagingQueuesKafkaTopicThroughputConsumerDetails) | **Post** /v1/o11y/messaging-queues/kafka/topic-throughput/consumer-details | Breaks one consumer topic&#39;s throughput down using the topic and service named in variables.
[**PostO11yMessagingQueuesKafkaTopicThroughputProducer**](O11yAPI.md#PostO11yMessagingQueuesKafkaTopicThroughputProducer) | **Post** /v1/o11y/messaging-queues/kafka/topic-throughput/producer | Returns the producer topic-throughput overview for the window — what each producer service wrote, per topic.
[**PostO11yMessagingQueuesKafkaTopicThroughputProducerDetails**](O11yAPI.md#PostO11yMessagingQueuesKafkaTopicThroughputProducerDetails) | **Post** /v1/o11y/messaging-queues/kafka/topic-throughput/producer-details | Breaks one producer topic&#39;s throughput down using the topic and service named in variables.
[**PostO11yMessagingQueuesQueueOverview**](O11yAPI.md#PostO11yMessagingQueuesQueueOverview) | **Post** /v1/o11y/messaging-queues/queue-overview | Lists the messaging destinations observed in the window — one row per queue/destination/service combination with its throughput and latency columns.
[**PostO11yNamespacesList**](O11yAPI.md#PostO11yNamespacesList) | **Post** /v1/o11y/namespaces/list | Lists Kubernetes namespaces over a time range, each with the CPU and memory its pods used, their phase counts and its attributes; filterable, groupable and paginated.
[**PostO11yNodesList**](O11yAPI.md#PostO11yNodesList) | **Post** /v1/o11y/nodes/list | Lists Kubernetes nodes over a time range, each with its CPU and memory usage against allocatable capacity, readiness condition counts and attributes; filterable, groupable and paginated.
[**PostO11yPodsList**](O11yAPI.md#PostO11yPodsList) | **Post** /v1/o11y/pods/list | Lists Kubernetes pods over a time range, each with its CPU and memory usage against request and limit, restart count, phase counts and attributes; filterable, groupable and paginated.
[**PostO11yProcessesList**](O11yAPI.md#PostO11yProcessesList) | **Post** /v1/o11y/processes/list | Lists monitored processes over a time range, each with its name, PID, command line and CPU and memory usage; filterable, groupable and paginated.
[**PostO11yPvcsList**](O11yAPI.md#PostO11yPvcsList) | **Post** /v1/o11y/pvcs/list | Lists Kubernetes persistent volume claims over a time range, each with its available, capacity and used bytes, inode counts and attributes; filterable, groupable and paginated.
[**PostO11yQueryFilterAnalyze**](O11yAPI.md#PostO11yQueryFilterAnalyze) | **Post** /v1/o11y/query_filter/analyze | Analyzes a query and extracts the metric names it reads and the columns it groups by.
[**PostO11yQueryRange**](O11yAPI.md#PostO11yQueryRange) | **Post** /v1/o11y/query_range | Executes a composite query over a time range: builder queries over traces, logs and metrics, formulas, trace operators, PromQL and Datastore SQL, answering time series, scalars or raw records as the request type asks.
[**PostO11yQueryRangeFormat**](O11yAPI.md#PostO11yQueryRangeFormat) | **Post** /v1/o11y/query_range/format | Parses a builder query and echoes it back normalized to the v3 shape — the endpoint the UI uses to canonicalize a query without running it.
[**PostO11yQueryRangePreview**](O11yAPI.md#PostO11yQueryRangePreview) | **Post** /v1/o11y/query_range/preview | Validates a composite query and renders the Datastore statements it would run WITHOUT executing it — a dry run for agentic and tooling use.
[**PostO11yRegister**](O11yAPI.md#PostO11yRegister) | **Post** /v1/o11y/register | Creates the FIRST organization and its admin user.
[**PostO11yReviews**](O11yAPI.md#PostO11yReviews) | **Post** /v1/o11y/reviews | Creates a human-review queue in the caller&#39;s org and project.
[**PostO11yReviewsByIdItems**](O11yAPI.md#PostO11yReviewsByIdItems) | **Post** /v1/o11y/reviews/{id}/items | Enqueues traces, observations or sessions on a review queue.
[**PostO11ySentinelDiscover**](O11yAPI.md#PostO11ySentinelDiscover) | **Post** /v1/o11y/sentinel/discover | Aggregates a project&#39;s captured errors into a table — the caller names the filters, the groupings and the aggregations, and gets back the columns and rows they asked for.
[**PostO11ySentinelProjects**](O11yAPI.md#PostO11ySentinelProjects) | **Post** /v1/o11y/sentinel/projects | Creates a Sentry project under the caller&#39;s org and returns it, DSN included.
[**PostO11ySentinelProjectsByIdKeysRotate**](O11yAPI.md#PostO11ySentinelProjectsByIdKeysRotate) | **Post** /v1/o11y/sentinel/projects/{id}/keys/rotate | Rotates a project&#39;s DSN key — bumping its rotation watermark so keys below it stop verifying — and returns the project with its new DSN.
[**PostO11yServiceEntryPointOperations**](O11yAPI.md#PostO11yServiceEntryPointOperations) | **Post** /v1/o11y/service/entry_point_operations | Returns one service&#39;s entry-point operations with the same latency and error profile topOperations reports.
[**PostO11yServiceTopLevelOperations**](O11yAPI.md#PostO11yServiceTopLevelOperations) | **Post** /v1/o11y/service/top_level_operations | Maps each service to its entry-point span names — for the one service named in the request, or for every service when none is.
[**PostO11yServiceTopOperations**](O11yAPI.md#PostO11yServiceTopOperations) | **Post** /v1/o11y/service/top_operations | Returns one service&#39;s heaviest operations in the window, each with p50/p95/p99 latency, how often it ran and how often it errored.
[**PostO11yServices**](O11yAPI.md#PostO11yServices) | **Post** /v1/o11y/services | Lists the instrumented services seen in the window, each with the request profile of its entry-point spans: p99 and average latency, call and error rates, and the entry-point operations the numbers were computed over.
[**PostO11ySettingsApdex**](O11yAPI.md#PostO11ySettingsApdex) | **Post** /v1/o11y/settings/apdex | Sets one service&#39;s apdex threshold and the status codes excluded from its score.
[**PostO11ySettingsTtl**](O11yAPI.md#PostO11ySettingsTtl) | **Post** /v1/o11y/settings/ttl | Sets the org&#39;s retention policy for one signal: the default TTL in days, ordered per-label retention rules, and optional cold-storage settings.
[**PostO11ySpanPercentile**](O11yAPI.md#PostO11ySpanPercentile) | **Post** /v1/o11y/span_percentile | Places one span&#39;s duration among its peers: the p50/p90/p99 durations of like spans, and the percentile the given duration lands at.
[**PostO11yStatefulsetsList**](O11yAPI.md#PostO11yStatefulsetsList) | **Post** /v1/o11y/statefulsets/list | Lists Kubernetes statefulsets over a time range, each with the CPU and memory its pods used against request and limit, desired and available replica counts, restarts and attributes; filterable, groupable and paginated.
[**PostO11ySubstituteVars**](O11yAPI.md#PostO11ySubstituteVars) | **Post** /v1/o11y/substitute_vars | Substitutes a query&#39;s variables and returns the resolved request, without running it — what a dashboard does before it queries.
[**PostO11yThirdPartyApisOverviewDomain**](O11yAPI.md#PostO11yThirdPartyApisOverviewDomain) | **Post** /v1/o11y/third-party-apis/overview/domain | Returns one external domain&#39;s endpoint-level breakdown — each endpoint with its rate, error and latency columns over the window.
[**PostO11yThirdPartyApisOverviewList**](O11yAPI.md#PostO11yThirdPartyApisOverviewList) | **Post** /v1/o11y/third-party-apis/overview/list | Lists the external domains the instrumented services call, with request rate, error percentage and latency per domain.
[**PostO11yVariablesQuery**](O11yAPI.md#PostO11yVariablesQuery) | **Post** /v1/o11y/variables/query | Evaluates a dashboard variable query and returns the values the variable may take.
[**PreviewMetricReductionRule**](O11yAPI.md#PreviewMetricReductionRule) | **Post** /v1/o11y/metric_reduction_rules/preview | Estimates the series reduction and the dashboards and alerts a candidate volume-control rule would touch, without persisting it.
[**PutHost**](O11yAPI.md#PutHost) | **Put** /v1/o11y/zeus/hosts | Records the deployment&#39;s host in Zeus, overwriting any prior one.
[**PutO11yExplorerViewsByViewid**](O11yAPI.md#PutO11yExplorerViewsByViewid) | **Put** /v1/o11y/explorer/views/{viewId} | Replaces one saved explorer view by id with the given view and echoes it back.
[**PutO11ySentinelIssuesById**](O11yAPI.md#PutO11ySentinelIssuesById) | **Put** /v1/o11y/sentinel/issues/{id} | Changes an issue&#39;s lifecycle — resolve, ignore, reopen or assign — and returns the updated issue.
[**PutProfile**](O11yAPI.md#PutProfile) | **Put** /v1/o11y/zeus/profiles | Records the deployment&#39;s profile in Zeus — how the team uses observability today and what they plan — overwriting any prior one.
[**RemoveUserRoleByUserIDAndRoleID**](O11yAPI.md#RemoveUserRoleByUserIDAndRoleID) | **Delete** /v1/o11y/users/{id}/roles/{roleId} | Takes a role away from one org member, by user id and role id — someone else, never the caller.
[**ResetPassword**](O11yAPI.md#ResetPassword) | **Post** /v1/o11y/resetPassword | Sets a new password for whoever the reset token was minted for, consuming the token.
[**RevokeServiceAccountKey**](O11yAPI.md#RevokeServiceAccountKey) | **Delete** /v1/o11y/service_accounts/{id}/keys/{fid} | Revokes an API key.
[**RotateSession**](O11yAPI.md#RotateSession) | **Post** /v1/o11y/sessions/rotate | Exchanges a refresh token for a fresh token pair, retiring the old pair.
[**SearchIngestionKeys**](O11yAPI.md#SearchIngestionKeys) | **Get** /v1/o11y/gateway/ingestion_keys/search | Lists the workspace&#39;s ingestion keys whose name matches the search, paginated.
[**SearchTraces**](O11yAPI.md#SearchTraces) | **Get** /v1/o11y/traces/{traceId} | Returns one trace&#39;s spans as a column/row table, optionally centred on a span and walked a fixed number of levels up and down from it — the read the trace explorer opens a trace with.
[**SetRoleByUserID**](O11yAPI.md#SetRoleByUserID) | **Post** /v1/o11y/users/{id}/roles | Assigns a role, by role name, to one org member — someone else, never the caller.
[**TestChannel**](O11yAPI.md#TestChannel) | **Post** /v1/o11y/channels/test | Sends a test notification to the posted receiver.
[**TestChannelDeprecated**](O11yAPI.md#TestChannelDeprecated) | **Post** /v1/o11y/testChannel | Sends a test notification to the posted receiver.
[**TestRule**](O11yAPI.md#TestRule) | **Post** /v1/o11y/rules/test | Fires a test notification for a rule definition without saving it, answering with how many series would alert.
[**TestRuleNotification**](O11yAPI.md#TestRuleNotification) | **Post** /v1/o11y/testRule | Fires a test notification for the posted rule definition and answers with how many series alerted and a status message.
[**UninstallIntegration**](O11yAPI.md#UninstallIntegration) | **Post** /v1/o11y/integrations/uninstall | Removes an integration from the caller&#39;s org by id.
[**UnlockDashboardV2**](O11yAPI.md#UnlockDashboardV2) | **Delete** /v1/o11y/dashboards/{id}/lock | Unlocks a v2-shape dashboard.
[**UnpinDashboardV2**](O11yAPI.md#UnpinDashboardV2) | **Delete** /v1/o11y/users/me/dashboards/{id}/pins | Removes the caller&#39;s pin for a dashboard.
[**UpdateAccount**](O11yAPI.md#UpdateAccount) | **Put** /v1/o11y/cloud_integrations/{cloud_provider}/accounts/{id} | Changes a connected account&#39;s configuration for the given provider, by id.
[**UpdateAuthDomain**](O11yAPI.md#UpdateAuthDomain) | **Put** /v1/o11y/domains/{id} | Replaces one auth domain&#39;s SSO configuration, by id.
[**UpdateChannelByID**](O11yAPI.md#UpdateChannelByID) | **Put** /v1/o11y/channels/{id} | Replaces a notification channel&#39;s receiver, by id.
[**UpdateDashboardV2**](O11yAPI.md#UpdateDashboardV2) | **Put** /v1/o11y/dashboards/{id} | Updates a v2-shape dashboard&#39;s metadata, spec and tag set.
[**UpdateDashboardView**](O11yAPI.md#UpdateDashboardView) | **Put** /v1/o11y/dashboard_views/{id} | Replaces a saved view&#39;s name and data.
[**UpdateDowntimeScheduleByID**](O11yAPI.md#UpdateDowntimeScheduleByID) | **Put** /v1/o11y/downtime_schedules/{id} | Replaces a planned maintenance window, by id.
[**UpdateIngestionKey**](O11yAPI.md#UpdateIngestionKey) | **Patch** /v1/o11y/gateway/ingestion_keys/{keyId} | Changes an ingestion key, by id.
[**UpdateIngestionKeyLimit**](O11yAPI.md#UpdateIngestionKeyLimit) | **Patch** /v1/o11y/gateway/ingestion_keys/limits/{limitId} | Changes an ingestion key limit, by limit id.
[**UpdateMetricMetadata**](O11yAPI.md#UpdateMetricMetadata) | **Post** /v1/o11y/metrics/metadata | Updates one metric&#39;s metadata — description, type, unit, temporality, monotonicity — and answers with the bare success envelope.
[**UpdateMetricReductionRuleByID**](O11yAPI.md#UpdateMetricReductionRuleByID) | **Put** /v1/o11y/metric_reduction_rules/{id} | Updates the match type and labels of a volume-control rule by its id; the metric name is immutable.
[**UpdateMyOrganization**](O11yAPI.md#UpdateMyOrganization) | **Put** /v1/o11y/orgs/me | Rewrites the caller&#39;s own organization record — display name, name, alias — always addressed as \&quot;me\&quot;, never by id.
[**UpdateMyPassword**](O11yAPI.md#UpdateMyPassword) | **Put** /v1/o11y/users/me/factor_password | Replaces the calling user&#39;s password, refusing when the old one does not match.
[**UpdateMyServiceAccount**](O11yAPI.md#UpdateMyServiceAccount) | **Put** /v1/o11y/service_accounts/me | Renames the calling service account.
[**UpdateMyUserV2**](O11yAPI.md#UpdateMyUserV2) | **Put** /v1/o11y/users/me | Renames the calling user.
[**UpdateOrgPreference**](O11yAPI.md#UpdateOrgPreference) | **Put** /v1/o11y/org/preferences/{name} | Sets one org-scoped preference, by name.
[**UpdatePublicDashboard**](O11yAPI.md#UpdatePublicDashboard) | **Put** /v1/o11y/dashboards/{id}/public | Updates the public-sharing config for a dashboard.
[**UpdateQuickFilters**](O11yAPI.md#UpdateQuickFilters) | **Put** /v1/o11y/orgs/me/filters | Replaces the org&#39;s quick filters for one signal with the attribute list given.
[**UpdateRole**](O11yAPI.md#UpdateRole) | **Put** /v1/o11y/roles/{id} | Replaces a custom role&#39;s description and transaction groups.
[**UpdateRoutePolicy**](O11yAPI.md#UpdateRoutePolicy) | **Put** /v1/o11y/route_policies/{id} | Replaces a route policy, by id, answering with the stored policy.
[**UpdateRuleByID**](O11yAPI.md#UpdateRuleByID) | **Put** /v1/o11y/rules/{id} | Replaces an alert rule&#39;s definition, by id.
[**UpdateService**](O11yAPI.md#UpdateService) | **Put** /v1/o11y/cloud_integrations/{cloud_provider}/accounts/{id}/services/{service_id} | Changes a service&#39;s configuration for one connected account of the given provider, by account id and service id.
[**UpdateServiceAccount**](O11yAPI.md#UpdateServiceAccount) | **Put** /v1/o11y/service_accounts/{id} | Renames a service account.
[**UpdateServiceAccountKey**](O11yAPI.md#UpdateServiceAccountKey) | **Put** /v1/o11y/service_accounts/{id}/keys/{fid} | Renames an API key or moves its expiry.
[**UpdateSpanMapper**](O11yAPI.md#UpdateSpanMapper) | **Patch** /v1/o11y/span_mapper_groups/{groupId}/span_mappers/{mapperId} | Changes a mapper&#39;s field context, config or enabled state.
[**UpdateSpanMapperGroup**](O11yAPI.md#UpdateSpanMapperGroup) | **Patch** /v1/o11y/span_mapper_groups/{groupId} | Changes a group&#39;s name, condition or enabled state.
[**UpdateTraceField**](O11yAPI.md#UpdateTraceField) | **Post** /v1/o11y/traces/fields | Changes how one span field is stored — selects or deselects it as a materialized column and tunes its index — and echoes the setting back.
[**UpdateTraceFunnel**](O11yAPI.md#UpdateTraceFunnel) | **Put** /v1/o11y/trace-funnels/{funnel_id} | Renames a funnel or rewrites its description, answering the funnel as it now stands.
[**UpdateTraceFunnelSteps**](O11yAPI.md#UpdateTraceFunnelSteps) | **Put** /v1/o11y/trace-funnels/steps/update | Replaces a funnel&#39;s steps — the funnel is named in the body rather than the path — and answers the funnel as it now stands.
[**UpdateUser**](O11yAPI.md#UpdateUser) | **Put** /v1/o11y/users/{id} | Renames one org member, by user id — someone else, never the caller, who renames themselves through updateMyUser.
[**UpdateUserDeprecated**](O11yAPI.md#UpdateUserDeprecated) | **Put** /v1/o11y/user/{id} | Renames one org member and may move their legacy role, answering with the updated record.
[**UpdateUserPreference**](O11yAPI.md#UpdateUserPreference) | **Put** /v1/o11y/user/preferences/{name} | Sets one preference of the calling user, by name.
[**ValidateDraftFunnelTraces**](O11yAPI.md#ValidateDraftFunnelTraces) | **Post** /v1/o11y/trace-funnels/analytics/validate | Lists the traces that match a funnel described inline — the builder&#39;s \&quot;try this\&quot; before anything is saved.
[**ValidateTraceFunnelTraces**](O11yAPI.md#ValidateTraceFunnelTraces) | **Post** /v1/o11y/trace-funnels/{funnel_id}/analytics/validate | Lists the traces that match a saved funnel over a window — the read that answers \&quot;is this funnel finding anything at all\&quot;.
[**VerifyResetPasswordToken**](O11yAPI.md#VerifyResetPasswordToken) | **Post** /v1/o11y/reset_password_tokens/verify | Checks that a reset-password token exists and has not expired, without consuming it.



## AgentCheckIn

> O11yO11yAgentCheckInOut AgentCheckIn(ctx, cloudProvider).O11yO11yAgentCheckInIn(o11yO11yAgentCheckInIn).Execute()

Is the deployed agent's check-in — the path consistent with the account surface, reporting the agent's account and telemetry state so the connection can be tracked.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	cloudProvider := "cloudProvider_example" // string | 
	o11yO11yAgentCheckInIn := *openapiclient.NewO11yO11yAgentCheckInIn() // O11yO11yAgentCheckInIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.AgentCheckIn(context.Background(), cloudProvider).O11yO11yAgentCheckInIn(o11yO11yAgentCheckInIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.AgentCheckIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentCheckIn`: O11yO11yAgentCheckInOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.AgentCheckIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentCheckInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yAgentCheckInIn** | [**O11yO11yAgentCheckInIn**](O11yO11yAgentCheckInIn.md) |  | 

### Return type

[**O11yO11yAgentCheckInOut**](O11yO11yAgentCheckInOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentCheckInDeprecated

> O11yO11yAgentCheckInOut AgentCheckInDeprecated(ctx, cloudProvider).O11yO11yAgentCheckInIn(o11yO11yAgentCheckInIn).Execute()

Is the deployed agent's check-in on its original hyphenated path, kept for backward compatibility with agents already running.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	cloudProvider := "cloudProvider_example" // string | 
	o11yO11yAgentCheckInIn := *openapiclient.NewO11yO11yAgentCheckInIn() // O11yO11yAgentCheckInIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.AgentCheckInDeprecated(context.Background(), cloudProvider).O11yO11yAgentCheckInIn(o11yO11yAgentCheckInIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.AgentCheckInDeprecated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentCheckInDeprecated`: O11yO11yAgentCheckInOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.AgentCheckInDeprecated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentCheckInDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yAgentCheckInIn** | [**O11yO11yAgentCheckInIn**](O11yO11yAgentCheckInIn.md) |  | 

### Return type

[**O11yO11yAgentCheckInOut**](O11yO11yAgentCheckInOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthzCheck

> O11yO11yCheckOut AuthzCheck(ctx).O11yO11yTransaction(o11yO11yTransaction).Execute()

Evaluates a batch of transactions — relation plus object — for the authenticated caller and answers each with its authorization verdict, in the order they were asked.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yTransaction := []openapiclient.O11yO11yTransaction{*openapiclient.NewO11yO11yTransaction()} // []O11yO11yTransaction | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.AuthzCheck(context.Background()).O11yO11yTransaction(o11yO11yTransaction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.AuthzCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthzCheck`: O11yO11yCheckOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.AuthzCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthzCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yTransaction** | [**[]O11yO11yTransaction**](O11yO11yTransaction.md) |  | 

### Return type

[**O11yO11yCheckOut**](O11yO11yCheckOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloneDashboardV2

> O11yO11yDashboardOut CloneDashboardV2(ctx, id).Execute()

Clones an existing v2-shape dashboard.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the resource id from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CloneDashboardV2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CloneDashboardV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloneDashboardV2`: O11yO11yDashboardOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CloneDashboardV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the resource id from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneDashboardV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yDashboardOut**](O11yO11yDashboardOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccount

> O11yO11yCreateAccountOut CreateAccount(ctx, cloudProvider).O11yO11yCreateAccountIn(o11yO11yCreateAccountIn).Execute()

Connects a new cloud-integration account for the given provider from its posted config and credentials, answering with the account and the artifact the agent deploys to complete the connection.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	cloudProvider := "cloudProvider_example" // string | 
	o11yO11yCreateAccountIn := *openapiclient.NewO11yO11yCreateAccountIn() // O11yO11yCreateAccountIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CreateAccount(context.Background(), cloudProvider).O11yO11yCreateAccountIn(o11yO11yCreateAccountIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccount`: O11yO11yCreateAccountOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CreateAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yCreateAccountIn** | [**O11yO11yCreateAccountIn**](O11yO11yCreateAccountIn.md) |  | 

### Return type

[**O11yO11yCreateAccountOut**](O11yO11yCreateAccountOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthDomain

> O11yO11yCreatedOut CreateAuthDomain(ctx).O11yO11yPostableAuthDomain(o11yO11yPostableAuthDomain).Execute()

Claims an email domain for the org and configures how its users sign in; the answer is the new domain's id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yPostableAuthDomain := *openapiclient.NewO11yO11yPostableAuthDomain() // O11yO11yPostableAuthDomain | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CreateAuthDomain(context.Background()).O11yO11yPostableAuthDomain(o11yO11yPostableAuthDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreateAuthDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAuthDomain`: O11yO11yCreatedOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CreateAuthDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yPostableAuthDomain** | [**O11yO11yPostableAuthDomain**](O11yO11yPostableAuthDomain.md) |  | 

### Return type

[**O11yO11yCreatedOut**](O11yO11yCreatedOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBulkInvite

> O11yO11yAck CreateBulkInvite(ctx).O11yO11yBulkInviteIn(o11yO11yBulkInviteIn).Execute()

Invites several people to the caller's org in one call, refusing the whole batch when any email repeats.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yBulkInviteIn := *openapiclient.NewO11yO11yBulkInviteIn() // O11yO11yBulkInviteIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CreateBulkInvite(context.Background()).O11yO11yBulkInviteIn(o11yO11yBulkInviteIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreateBulkInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBulkInvite`: O11yO11yAck
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CreateBulkInvite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBulkInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yBulkInviteIn** | [**O11yO11yBulkInviteIn**](O11yO11yBulkInviteIn.md) |  | 

### Return type

[**O11yO11yAck**](O11yO11yAck.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateChannel

> O11yO11yChannelOut CreateChannel(ctx).O11yPostableChannel(o11yPostableChannel).Execute()

Creates a notification channel, answering with the stored channel.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yPostableChannel := *openapiclient.NewO11yPostableChannel() // O11yPostableChannel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CreateChannel(context.Background()).O11yPostableChannel(o11yPostableChannel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreateChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateChannel`: O11yO11yChannelOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CreateChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yPostableChannel** | [**O11yPostableChannel**](O11yPostableChannel.md) |  | 

### Return type

[**O11yO11yChannelOut**](O11yO11yChannelOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDashboardV2

> O11yO11yDashboardOut CreateDashboardV2(ctx).O11yO11yDashboardPostable(o11yO11yDashboardPostable).Execute()

Creates a dashboard in the v2 format that follows the Perses spec and answers with the stored dashboard.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yDashboardPostable := *openapiclient.NewO11yO11yDashboardPostable() // O11yO11yDashboardPostable | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CreateDashboardV2(context.Background()).O11yO11yDashboardPostable(o11yO11yDashboardPostable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreateDashboardV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDashboardV2`: O11yO11yDashboardOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CreateDashboardV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDashboardV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yDashboardPostable** | [**O11yO11yDashboardPostable**](O11yO11yDashboardPostable.md) |  | 

### Return type

[**O11yO11yDashboardOut**](O11yO11yDashboardOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDashboardView

> O11yO11yDashboardViewOut CreateDashboardView(ctx).O11yO11yDashboardViewPostable(o11yO11yDashboardViewPostable).Execute()

Persists the calling user's dashboard-listing state (query, sort, order) as a named, reusable view shared across the org.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yDashboardViewPostable := *openapiclient.NewO11yO11yDashboardViewPostable() // O11yO11yDashboardViewPostable | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CreateDashboardView(context.Background()).O11yO11yDashboardViewPostable(o11yO11yDashboardViewPostable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreateDashboardView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDashboardView`: O11yO11yDashboardViewOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CreateDashboardView`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDashboardViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yDashboardViewPostable** | [**O11yO11yDashboardViewPostable**](O11yO11yDashboardViewPostable.md) |  | 

### Return type

[**O11yO11yDashboardViewOut**](O11yO11yDashboardViewOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDowntimeSchedule

> O11yO11yDowntimeScheduleOut CreateDowntimeSchedule(ctx).O11yPostablePlannedMaintenance(o11yPostablePlannedMaintenance).Execute()

Creates a planned maintenance window, answering with the stored schedule.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yPostablePlannedMaintenance := *openapiclient.NewO11yPostablePlannedMaintenance() // O11yPostablePlannedMaintenance | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CreateDowntimeSchedule(context.Background()).O11yPostablePlannedMaintenance(o11yPostablePlannedMaintenance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreateDowntimeSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDowntimeSchedule`: O11yO11yDowntimeScheduleOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CreateDowntimeSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDowntimeScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yPostablePlannedMaintenance** | [**O11yPostablePlannedMaintenance**](O11yPostablePlannedMaintenance.md) |  | 

### Return type

[**O11yO11yDowntimeScheduleOut**](O11yO11yDowntimeScheduleOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIngestionKey

> O11yO11yCreatedIngestionKeyOut CreateIngestionKey(ctx).O11yPostableIngestionKey(o11yPostableIngestionKey).Execute()

Mints an ingestion key for the workspace, answering with the created key.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yPostableIngestionKey := *openapiclient.NewO11yPostableIngestionKey() // O11yPostableIngestionKey | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CreateIngestionKey(context.Background()).O11yPostableIngestionKey(o11yPostableIngestionKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreateIngestionKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIngestionKey`: O11yO11yCreatedIngestionKeyOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CreateIngestionKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIngestionKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yPostableIngestionKey** | [**O11yPostableIngestionKey**](O11yPostableIngestionKey.md) |  | 

### Return type

[**O11yO11yCreatedIngestionKeyOut**](O11yO11yCreatedIngestionKeyOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIngestionKeyLimit

> O11yO11yCreatedLimitOut CreateIngestionKeyLimit(ctx, keyId).O11yO11yCreateLimitIn(o11yO11yCreateLimitIn).Execute()

Sets a signal limit on an ingestion key, by key id, answering with the created limit.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	keyId := "keyId_example" // string | 
	o11yO11yCreateLimitIn := *openapiclient.NewO11yO11yCreateLimitIn() // O11yO11yCreateLimitIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CreateIngestionKeyLimit(context.Background(), keyId).O11yO11yCreateLimitIn(o11yO11yCreateLimitIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreateIngestionKeyLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIngestionKeyLimit`: O11yO11yCreatedLimitOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CreateIngestionKeyLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIngestionKeyLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yCreateLimitIn** | [**O11yO11yCreateLimitIn**](O11yO11yCreateLimitIn.md) |  | 

### Return type

[**O11yO11yCreatedLimitOut**](O11yO11yCreatedLimitOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvite

> O11yO11yInviteOut CreateInvite(ctx).O11yO11yInviteIn(o11yO11yInviteIn).Execute()

Invites one person to the caller's org by email, with the role they will hold when they accept.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yInviteIn := *openapiclient.NewO11yO11yInviteIn() // O11yO11yInviteIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CreateInvite(context.Background()).O11yO11yInviteIn(o11yO11yInviteIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreateInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInvite`: O11yO11yInviteOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CreateInvite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yInviteIn** | [**O11yO11yInviteIn**](O11yO11yInviteIn.md) |  | 

### Return type

[**O11yO11yInviteOut**](O11yO11yInviteOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLLMAnnotation

> O11yO11yLLMAnnotationOut CreateLLMAnnotation(ctx).O11yO11yLLMIngestAnnotation(o11yO11yLLMIngestAnnotation).Execute()

Adds a human annotation to a trace or observation, optionally in a review queue.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yLLMIngestAnnotation := *openapiclient.NewO11yO11yLLMIngestAnnotation() // O11yO11yLLMIngestAnnotation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CreateLLMAnnotation(context.Background()).O11yO11yLLMIngestAnnotation(o11yO11yLLMIngestAnnotation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreateLLMAnnotation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLLMAnnotation`: O11yO11yLLMAnnotationOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CreateLLMAnnotation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLLMAnnotationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yLLMIngestAnnotation** | [**O11yO11yLLMIngestAnnotation**](O11yO11yLLMIngestAnnotation.md) |  | 

### Return type

[**O11yO11yLLMAnnotationOut**](O11yO11yLLMAnnotationOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLLMScore

> O11yO11yLLMScoreOut CreateLLMScore(ctx).O11yO11yLLMIngestScore(o11yO11yLLMIngestScore).Execute()

Attaches an eval score or human-feedback signal to a trace or a single observation.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yLLMIngestScore := *openapiclient.NewO11yO11yLLMIngestScore() // O11yO11yLLMIngestScore | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CreateLLMScore(context.Background()).O11yO11yLLMIngestScore(o11yO11yLLMIngestScore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreateLLMScore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLLMScore`: O11yO11yLLMScoreOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CreateLLMScore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLLMScoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yLLMIngestScore** | [**O11yO11yLLMIngestScore**](O11yO11yLLMIngestScore.md) |  | 

### Return type

[**O11yO11yLLMScoreOut**](O11yO11yLLMScoreOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMetricReductionRule

> O11yO11yReductionRuleOut CreateMetricReductionRule(ctx).O11yO11yReductionRuleCreateIn(o11yO11yReductionRuleCreateIn).Execute()

Creates a volume-control rule for a metric and returns it with its id; a metric that already has a rule is refused.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yReductionRuleCreateIn := *openapiclient.NewO11yO11yReductionRuleCreateIn([]string{"Labels_example"}, "MatchType_example", "MetricName_example") // O11yO11yReductionRuleCreateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CreateMetricReductionRule(context.Background()).O11yO11yReductionRuleCreateIn(o11yO11yReductionRuleCreateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreateMetricReductionRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMetricReductionRule`: O11yO11yReductionRuleOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CreateMetricReductionRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMetricReductionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yReductionRuleCreateIn** | [**O11yO11yReductionRuleCreateIn**](O11yO11yReductionRuleCreateIn.md) |  | 

### Return type

[**O11yO11yReductionRuleOut**](O11yO11yReductionRuleOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateLLMPricingRules

> CreateOrUpdateLLMPricingRules(ctx).O11yO11yLLMUpdatablePricingRules(o11yO11yLLMUpdatablePricingRules).Execute()

Writes the pricing-rule batch — the single write endpoint used by both the user and the Zeus sync job.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yLLMUpdatablePricingRules := *openapiclient.NewO11yO11yLLMUpdatablePricingRules() // O11yO11yLLMUpdatablePricingRules | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.CreateOrUpdateLLMPricingRules(context.Background()).O11yO11yLLMUpdatablePricingRules(o11yO11yLLMUpdatablePricingRules).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreateOrUpdateLLMPricingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateLLMPricingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yLLMUpdatablePricingRules** | [**O11yO11yLLMUpdatablePricingRules**](O11yO11yLLMUpdatablePricingRules.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePublicDashboard

> O11yO11yIdentifiableOut CreatePublicDashboard(ctx, id).O11yO11yPublicDashboardWriteIn(o11yO11yPublicDashboardWriteIn).Execute()

Creates the public-sharing config for a dashboard and enables public sharing, answering with the new share's id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the dashboard id from the path.
	o11yO11yPublicDashboardWriteIn := *openapiclient.NewO11yO11yPublicDashboardWriteIn() // O11yO11yPublicDashboardWriteIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CreatePublicDashboard(context.Background(), id).O11yO11yPublicDashboardWriteIn(o11yO11yPublicDashboardWriteIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreatePublicDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePublicDashboard`: O11yO11yIdentifiableOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CreatePublicDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the dashboard id from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePublicDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yPublicDashboardWriteIn** | [**O11yO11yPublicDashboardWriteIn**](O11yO11yPublicDashboardWriteIn.md) |  | 

### Return type

[**O11yO11yIdentifiableOut**](O11yO11yIdentifiableOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateResetPasswordToken

> O11yO11yResetTokenOut CreateResetPasswordToken(ctx, id).Execute()

Creates or regenerates a user's reset-password token: a live token is returned as it is, an expired one is replaced.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CreateResetPasswordToken(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreateResetPasswordToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateResetPasswordToken`: O11yO11yResetTokenOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CreateResetPasswordToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateResetPasswordTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yResetTokenOut**](O11yO11yResetTokenOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> O11yO11yRoleCreateOut CreateRole(ctx).O11yO11yRoleCreateIn(o11yO11yRoleCreateIn).Execute()

Creates a custom role in the caller's org from a name, an optional description and the transaction groups it grants, answering the new role's id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yRoleCreateIn := *openapiclient.NewO11yO11yRoleCreateIn() // O11yO11yRoleCreateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CreateRole(context.Background()).O11yO11yRoleCreateIn(o11yO11yRoleCreateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreateRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRole`: O11yO11yRoleCreateOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CreateRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yRoleCreateIn** | [**O11yO11yRoleCreateIn**](O11yO11yRoleCreateIn.md) |  | 

### Return type

[**O11yO11yRoleCreateOut**](O11yO11yRoleCreateOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRoutePolicy

> O11yO11yRoutePolicyOut CreateRoutePolicy(ctx).O11yPostableRoutePolicy(o11yPostableRoutePolicy).Execute()

Creates a route policy, answering with the stored policy.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yPostableRoutePolicy := *openapiclient.NewO11yPostableRoutePolicy() // O11yPostableRoutePolicy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CreateRoutePolicy(context.Background()).O11yPostableRoutePolicy(o11yPostableRoutePolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreateRoutePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRoutePolicy`: O11yO11yRoutePolicyOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CreateRoutePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoutePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yPostableRoutePolicy** | [**O11yPostableRoutePolicy**](O11yPostableRoutePolicy.md) |  | 

### Return type

[**O11yO11yRoutePolicyOut**](O11yO11yRoutePolicyOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRule

> O11yO11yRuleOut CreateRule(ctx).Body(body).Execute()

Creates a new alert rule and answers with the stored rule.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := interface{}(987) // interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CreateRule(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreateRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRule`: O11yO11yRuleOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CreateRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **interface{}** |  | 

### Return type

[**O11yO11yRuleOut**](O11yO11yRuleOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServiceAccount

> O11yO11yServiceAccountCreateOut CreateServiceAccount(ctx).O11yO11yServiceAccountCreateIn(o11yO11yServiceAccountCreateIn).Execute()

Creates a service account in the caller's org, answering its id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yServiceAccountCreateIn := *openapiclient.NewO11yO11yServiceAccountCreateIn() // O11yO11yServiceAccountCreateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CreateServiceAccount(context.Background()).O11yO11yServiceAccountCreateIn(o11yO11yServiceAccountCreateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreateServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServiceAccount`: O11yO11yServiceAccountCreateOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CreateServiceAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yServiceAccountCreateIn** | [**O11yO11yServiceAccountCreateIn**](O11yO11yServiceAccountCreateIn.md) |  | 

### Return type

[**O11yO11yServiceAccountCreateOut**](O11yO11yServiceAccountCreateOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServiceAccountKey

> O11yO11yAPIKeyCreateOut CreateServiceAccountKey(ctx, id).O11yO11yAPIKeyCreateIn(o11yO11yAPIKeyCreateIn).Execute()

Mints an API key for a service account and answers the key's id and its secret — the one time the secret is ever shown.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 
	o11yO11yAPIKeyCreateIn := *openapiclient.NewO11yO11yAPIKeyCreateIn() // O11yO11yAPIKeyCreateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CreateServiceAccountKey(context.Background(), id).O11yO11yAPIKeyCreateIn(o11yO11yAPIKeyCreateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreateServiceAccountKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServiceAccountKey`: O11yO11yAPIKeyCreateOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CreateServiceAccountKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceAccountKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yAPIKeyCreateIn** | [**O11yO11yAPIKeyCreateIn**](O11yO11yAPIKeyCreateIn.md) |  | 

### Return type

[**O11yO11yAPIKeyCreateOut**](O11yO11yAPIKeyCreateOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServiceAccountRole

> CreateServiceAccountRole(ctx, id).O11yO11yServiceAccountRoleGrantIn(o11yO11yServiceAccountRoleGrantIn).Execute()

Assigns a role, named by its id, to a service account.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | RoleID is the id of the role to assign. Required.
	o11yO11yServiceAccountRoleGrantIn := *openapiclient.NewO11yO11yServiceAccountRoleGrantIn() // O11yO11yServiceAccountRoleGrantIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.CreateServiceAccountRole(context.Background(), id).O11yO11yServiceAccountRoleGrantIn(o11yO11yServiceAccountRoleGrantIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreateServiceAccountRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | RoleID is the id of the role to assign. Required. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceAccountRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yServiceAccountRoleGrantIn** | [**O11yO11yServiceAccountRoleGrantIn**](O11yO11yServiceAccountRoleGrantIn.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSessionByEmailPassword

> O11yO11yTokenOut CreateSessionByEmailPassword(ctx).O11yO11yEmailPasswordSessionIn(o11yO11yEmailPasswordSessionIn).Execute()

Signs a user in with email and password and answers with the session's token pair.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yEmailPasswordSessionIn := *openapiclient.NewO11yO11yEmailPasswordSessionIn() // O11yO11yEmailPasswordSessionIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CreateSessionByEmailPassword(context.Background()).O11yO11yEmailPasswordSessionIn(o11yO11yEmailPasswordSessionIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreateSessionByEmailPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSessionByEmailPassword`: O11yO11yTokenOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CreateSessionByEmailPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionByEmailPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yEmailPasswordSessionIn** | [**O11yO11yEmailPasswordSessionIn**](O11yO11yEmailPasswordSessionIn.md) |  | 

### Return type

[**O11yO11yTokenOut**](O11yO11yTokenOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSpanMapper

> O11yO11ySpanMapperOut CreateSpanMapper(ctx, groupId).O11yO11ySpanMapperCreateIn(o11yO11ySpanMapperCreateIn).Execute()

Adds a mapper to a group: which field context it reads, the move or copy it performs, and whether it is on.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	groupId := "groupId_example" // string | 
	o11yO11ySpanMapperCreateIn := *openapiclient.NewO11yO11ySpanMapperCreateIn() // O11yO11ySpanMapperCreateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CreateSpanMapper(context.Background(), groupId).O11yO11ySpanMapperCreateIn(o11yO11ySpanMapperCreateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreateSpanMapper``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSpanMapper`: O11yO11ySpanMapperOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CreateSpanMapper`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpanMapperRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11ySpanMapperCreateIn** | [**O11yO11ySpanMapperCreateIn**](O11yO11ySpanMapperCreateIn.md) |  | 

### Return type

[**O11yO11ySpanMapperOut**](O11yO11ySpanMapperOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSpanMapperGroup

> O11yO11ySpanMapperGroupOut CreateSpanMapperGroup(ctx).O11yPostableSpanMapperGroup(o11yPostableSpanMapperGroup).Execute()

Creates a mapping group: the name it is known by, the span and resource attributes whose presence selects a span into it, and whether it is on.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yPostableSpanMapperGroup := *openapiclient.NewO11yPostableSpanMapperGroup() // O11yPostableSpanMapperGroup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CreateSpanMapperGroup(context.Background()).O11yPostableSpanMapperGroup(o11yPostableSpanMapperGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreateSpanMapperGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSpanMapperGroup`: O11yO11ySpanMapperGroupOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CreateSpanMapperGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpanMapperGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yPostableSpanMapperGroup** | [**O11yPostableSpanMapperGroup**](O11yPostableSpanMapperGroup.md) |  | 

### Return type

[**O11yO11ySpanMapperGroupOut**](O11yO11ySpanMapperGroupOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTraceFunnel

> O11yO11yFunnelOut CreateTraceFunnel(ctx).O11yO11yFunnelCreateIn(o11yO11yFunnelCreateIn).Execute()

Creates an empty funnel with a name, answering the funnel it created.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yFunnelCreateIn := *openapiclient.NewO11yO11yFunnelCreateIn() // O11yO11yFunnelCreateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CreateTraceFunnel(context.Background()).O11yO11yFunnelCreateIn(o11yO11yFunnelCreateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreateTraceFunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTraceFunnel`: O11yO11yFunnelOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CreateTraceFunnel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTraceFunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yFunnelCreateIn** | [**O11yO11yFunnelCreateIn**](O11yO11yFunnelCreateIn.md) |  | 

### Return type

[**O11yO11yFunnelOut**](O11yO11yFunnelOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> O11yO11yCreatedOut CreateUser(ctx).O11yO11yPostableUser(o11yO11yPostableUser).Execute()

Creates a member of the caller's org in the pending-invite state and mails them their invitation; the answer is the new user's id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yPostableUser := *openapiclient.NewO11yO11yPostableUser() // O11yO11yPostableUser | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.CreateUser(context.Background()).O11yO11yPostableUser(o11yO11yPostableUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.CreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUser`: O11yO11yCreatedOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yPostableUser** | [**O11yO11yPostableUser**](O11yO11yPostableUser.md) |  | 

### Return type

[**O11yO11yCreatedOut**](O11yO11yCreatedOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthDomain

> DeleteAuthDomain(ctx, id).Execute()

Releases an email domain and discards its SSO configuration, by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.DeleteAuthDomain(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.DeleteAuthDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteChannelByID

> DeleteChannelByID(ctx, id).Execute()

Removes a notification channel, by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.DeleteChannelByID(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.DeleteChannelByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteChannelByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDashboardV2

> DeleteDashboardV2(ctx, id).Execute()

Deletes a v2-shape dashboard along with its tag relations.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the resource id from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.DeleteDashboardV2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.DeleteDashboardV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the resource id from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDashboardV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDashboardView

> DeleteDashboardView(ctx, id).Execute()

Removes a saved view.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the resource id from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.DeleteDashboardView(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.DeleteDashboardView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the resource id from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDashboardViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDowntimeScheduleByID

> DeleteDowntimeScheduleByID(ctx, id).Execute()

Removes a planned maintenance window, by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.DeleteDowntimeScheduleByID(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.DeleteDowntimeScheduleByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDowntimeScheduleByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIngestionKey

> DeleteIngestionKey(ctx, keyId).Execute()

Removes an ingestion key, by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	keyId := "keyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.DeleteIngestionKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.DeleteIngestionKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIngestionKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIngestionKeyLimit

> DeleteIngestionKeyLimit(ctx, limitId).Execute()

Removes an ingestion key limit, by limit id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	limitId := "limitId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.DeleteIngestionKeyLimit(context.Background(), limitId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.DeleteIngestionKeyLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**limitId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIngestionKeyLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLLMPricingRule

> DeleteLLMPricingRule(ctx, id).Execute()

Hard-deletes a pricing rule by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.DeleteLLMPricingRule(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.DeleteLLMPricingRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLLMPricingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLLMScore

> DeleteLLMScore(ctx, id).Execute()

Hard-deletes a score by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.DeleteLLMScore(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.DeleteLLMScore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLLMScoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMetricReductionRuleByID

> DeleteMetricReductionRuleByID(ctx, id).Execute()

Deletes a volume-control rule by its id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the rule's id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.DeleteMetricReductionRuleByID(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.DeleteMetricReductionRuleByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the rule&#39;s id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMetricReductionRuleByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteO11yExplorerViewsByViewid

> O11yO11ySavedViewDeleteOut DeleteO11yExplorerViewsByViewid(ctx, viewId).Execute()

Deletes one saved explorer view by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	viewId := "viewId_example" // string | ViewID is the view's id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.DeleteO11yExplorerViewsByViewid(context.Background(), viewId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.DeleteO11yExplorerViewsByViewid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteO11yExplorerViewsByViewid`: O11yO11ySavedViewDeleteOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.DeleteO11yExplorerViewsByViewid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **string** | ViewID is the view&#39;s id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteO11yExplorerViewsByViewidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11ySavedViewDeleteOut**](O11yO11ySavedViewDeleteOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteO11yReviewsById

> O11yAnnQueueDeleted DeleteO11yReviewsById(ctx, id).Execute()

Removes one review queue and every item in it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "annq_1" // string | ID is the annotation queue to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.DeleteO11yReviewsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.DeleteO11yReviewsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteO11yReviewsById`: O11yAnnQueueDeleted
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.DeleteO11yReviewsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the annotation queue to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteO11yReviewsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yAnnQueueDeleted**](O11yAnnQueueDeleted.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteO11ySentinelProjectsById

> DeleteO11ySentinelProjectsById(ctx, id).Execute()

Deletes one Sentry project of the caller's org.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the project id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.DeleteO11ySentinelProjectsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.DeleteO11ySentinelProjectsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the project id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteO11ySentinelProjectsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePublicDashboard

> DeletePublicDashboard(ctx, id).Execute()

Deletes the public-sharing config and disables public sharing of a dashboard.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the resource id from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.DeletePublicDashboard(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.DeletePublicDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the resource id from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePublicDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> DeleteRole(ctx, id).Execute()

Deletes a custom role.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.DeleteRole(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.DeleteRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoutePolicyByID

> DeleteRoutePolicyByID(ctx, id).Execute()

Removes a route policy, by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.DeleteRoutePolicyByID(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.DeleteRoutePolicyByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoutePolicyByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRuleByID

> DeleteRuleByID(ctx, id).Execute()

Removes an alert rule, by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.DeleteRuleByID(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.DeleteRuleByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceAccount

> DeleteServiceAccount(ctx, id).Execute()

Deletes a service account and revokes every key it holds.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.DeleteServiceAccount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.DeleteServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceAccountRole

> DeleteServiceAccountRole(ctx, id, rid).Execute()

Removes a role from a service account.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 
	rid := "rid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.DeleteServiceAccountRole(context.Background(), id, rid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.DeleteServiceAccountRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**rid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceAccountRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSession

> DeleteSession(ctx).Execute()

Signs the calling session out, invalidating its tokens.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.DeleteSession(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.DeleteSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSessionRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSpanMapper

> DeleteSpanMapper(ctx, groupId, mapperId).Execute()

Deletes one mapper from a group.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	groupId := "groupId_example" // string | 
	mapperId := "mapperId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.DeleteSpanMapper(context.Background(), groupId, mapperId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.DeleteSpanMapper``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 
**mapperId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSpanMapperRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSpanMapperGroup

> DeleteSpanMapperGroup(ctx, groupId).Execute()

Deletes a mapping group and every mapper under it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	groupId := "groupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.DeleteSpanMapperGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.DeleteSpanMapperGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSpanMapperGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTraceFunnel

> O11yO11yFunnelDeleteOut DeleteTraceFunnel(ctx, funnelId).Execute()

Deletes a funnel.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	funnelId := "funnelId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.DeleteTraceFunnel(context.Background(), funnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.DeleteTraceFunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTraceFunnel`: O11yO11yFunnelDeleteOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.DeleteTraceFunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**funnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTraceFunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yFunnelDeleteOut**](O11yO11yFunnelDeleteOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, id).Execute()

Removes one org member, by user id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.DeleteUser(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserDeprecated

> DeleteUserDeprecated(ctx, id).Execute()

Removes one org member, by user id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.DeleteUserDeprecated(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.DeleteUserDeprecated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisconnectAccount

> DisconnectAccount(ctx, cloudProvider, id).Execute()

Tears down a connected account for the given provider, by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	cloudProvider := "cloudProvider_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.DisconnectAccount(context.Background(), cloudProvider, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.DisconnectAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForgotPassword

> ForgotPassword(ctx).O11yO11yForgotPasswordIn(o11yO11yForgotPasswordIn).Execute()

Starts the forgotten-password flow: the named user is mailed a reset link.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yForgotPasswordIn := *openapiclient.NewO11yO11yForgotPasswordIn() // O11yO11yForgotPasswordIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.ForgotPassword(context.Background()).O11yO11yForgotPasswordIn(o11yO11yForgotPasswordIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ForgotPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiForgotPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yForgotPasswordIn** | [**O11yO11yForgotPasswordIn**](O11yO11yForgotPasswordIn.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount

> O11yO11yAccountOut GetAccount(ctx, cloudProvider, id).Execute()

Returns one connected account for the given provider, by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	cloudProvider := "cloudProvider_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetAccount(context.Background(), cloudProvider, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccount`: O11yO11yAccountOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**O11yO11yAccountOut**](O11yO11yAccountOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountService

> O11yO11yServiceOut GetAccountService(ctx, cloudProvider, id, serviceId).Execute()

Returns one service and its configuration for a connected account of the given provider, by account id and service id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	cloudProvider := "cloudProvider_example" // string | 
	id := "id_example" // string | 
	serviceId := "serviceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetAccountService(context.Background(), cloudProvider, id, serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetAccountService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountService`: O11yO11yServiceOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetAccountService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** |  | 
**id** | **string** |  | 
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**O11yO11yServiceOut**](O11yO11yServiceOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlerts

> O11yO11yAlertsOut GetAlerts(ctx).Execute()

Returns the org's current alerts.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetAlerts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlerts`: O11yO11yAlertsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetAlerts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsRequest struct via the builder pattern


### Return type

[**O11yO11yAlertsOut**](O11yO11yAlertsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllRoutePolicies

> O11yO11yRoutePoliciesOut GetAllRoutePolicies(ctx).Execute()

Lists the org's route policies.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetAllRoutePolicies(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetAllRoutePolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllRoutePolicies`: O11yO11yRoutePoliciesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetAllRoutePolicies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRoutePoliciesRequest struct via the builder pattern


### Return type

[**O11yO11yRoutePoliciesOut**](O11yO11yRoutePoliciesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthDomain

> O11yO11yAuthDomainOut GetAuthDomain(ctx, id).Execute()

Returns one auth domain with its SSO configuration, by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetAuthDomain(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetAuthDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthDomain`: O11yO11yAuthDomainOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetAuthDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yAuthDomainOut**](O11yO11yAuthDomainOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelByID

> O11yO11yChannelOut GetChannelByID(ctx, id).Execute()

Returns one notification channel, by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetChannelByID(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetChannelByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelByID`: O11yO11yChannelOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetChannelByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yChannelOut**](O11yO11yChannelOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionCredentials

> O11yO11yCredentialsOut GetConnectionCredentials(ctx, cloudProvider).Execute()

Returns the credentials the connecting agent needs to establish the cloud integration, for the given cloud provider.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	cloudProvider := "cloudProvider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetConnectionCredentials(context.Background(), cloudProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetConnectionCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionCredentials`: O11yO11yCredentialsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetConnectionCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yCredentialsOut**](O11yO11yCredentialsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboardV2

> O11yO11yDashboardOut GetDashboardV2(ctx, id).Execute()

Returns a v2-shape dashboard.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the resource id from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetDashboardV2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetDashboardV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDashboardV2`: O11yO11yDashboardOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetDashboardV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the resource id from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yDashboardOut**](O11yO11yDashboardOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDowntimeScheduleByID

> O11yO11yDowntimeScheduleOut GetDowntimeScheduleByID(ctx, id).Execute()

Returns one planned maintenance window, by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetDowntimeScheduleByID(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetDowntimeScheduleByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDowntimeScheduleByID`: O11yO11yDowntimeScheduleOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetDowntimeScheduleByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDowntimeScheduleByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yDowntimeScheduleOut**](O11yO11yDowntimeScheduleOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDraftFunnelErrorTraces

> O11yO11yFunnelRowsOut GetDraftFunnelErrorTraces(ctx).O11yO11yDraftFunnelIn(o11yO11yDraftFunnelIn).Execute()

Returns the errored traces through a step transition of a funnel described inline.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yDraftFunnelIn := *openapiclient.NewO11yO11yDraftFunnelIn() // O11yO11yDraftFunnelIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetDraftFunnelErrorTraces(context.Background()).O11yO11yDraftFunnelIn(o11yO11yDraftFunnelIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetDraftFunnelErrorTraces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDraftFunnelErrorTraces`: O11yO11yFunnelRowsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetDraftFunnelErrorTraces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDraftFunnelErrorTracesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yDraftFunnelIn** | [**O11yO11yDraftFunnelIn**](O11yO11yDraftFunnelIn.md) |  | 

### Return type

[**O11yO11yFunnelRowsOut**](O11yO11yFunnelRowsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDraftFunnelOverview

> O11yO11yFunnelRowsOut GetDraftFunnelOverview(ctx).O11yO11yDraftFunnelIn(o11yO11yDraftFunnelIn).Execute()

Returns the conversion overview of a funnel described inline.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yDraftFunnelIn := *openapiclient.NewO11yO11yDraftFunnelIn() // O11yO11yDraftFunnelIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetDraftFunnelOverview(context.Background()).O11yO11yDraftFunnelIn(o11yO11yDraftFunnelIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetDraftFunnelOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDraftFunnelOverview`: O11yO11yFunnelRowsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetDraftFunnelOverview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDraftFunnelOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yDraftFunnelIn** | [**O11yO11yDraftFunnelIn**](O11yO11yDraftFunnelIn.md) |  | 

### Return type

[**O11yO11yFunnelRowsOut**](O11yO11yFunnelRowsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDraftFunnelSlowTraces

> O11yO11yFunnelRowsOut GetDraftFunnelSlowTraces(ctx).O11yO11yDraftFunnelIn(o11yO11yDraftFunnelIn).Execute()

Returns the slowest traces through a step transition of a funnel described inline.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yDraftFunnelIn := *openapiclient.NewO11yO11yDraftFunnelIn() // O11yO11yDraftFunnelIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetDraftFunnelSlowTraces(context.Background()).O11yO11yDraftFunnelIn(o11yO11yDraftFunnelIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetDraftFunnelSlowTraces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDraftFunnelSlowTraces`: O11yO11yFunnelRowsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetDraftFunnelSlowTraces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDraftFunnelSlowTracesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yDraftFunnelIn** | [**O11yO11yDraftFunnelIn**](O11yO11yDraftFunnelIn.md) |  | 

### Return type

[**O11yO11yFunnelRowsOut**](O11yO11yFunnelRowsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDraftFunnelStepMetrics

> O11yO11yFunnelRowsOut GetDraftFunnelStepMetrics(ctx).O11yO11yDraftFunnelIn(o11yO11yDraftFunnelIn).Execute()

Returns the per-step metrics of a funnel described inline.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yDraftFunnelIn := *openapiclient.NewO11yO11yDraftFunnelIn() // O11yO11yDraftFunnelIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetDraftFunnelStepMetrics(context.Background()).O11yO11yDraftFunnelIn(o11yO11yDraftFunnelIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetDraftFunnelStepMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDraftFunnelStepMetrics`: O11yO11yFunnelRowsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetDraftFunnelStepMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDraftFunnelStepMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yDraftFunnelIn** | [**O11yO11yDraftFunnelIn**](O11yO11yDraftFunnelIn.md) |  | 

### Return type

[**O11yO11yFunnelRowsOut**](O11yO11yFunnelRowsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDraftFunnelStepOverview

> O11yO11yFunnelRowsOut GetDraftFunnelStepOverview(ctx).O11yO11yDraftFunnelIn(o11yO11yDraftFunnelIn).Execute()

Returns the conversion between two steps of a funnel described inline.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yDraftFunnelIn := *openapiclient.NewO11yO11yDraftFunnelIn() // O11yO11yDraftFunnelIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetDraftFunnelStepOverview(context.Background()).O11yO11yDraftFunnelIn(o11yO11yDraftFunnelIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetDraftFunnelStepOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDraftFunnelStepOverview`: O11yO11yFunnelRowsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetDraftFunnelStepOverview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDraftFunnelStepOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yDraftFunnelIn** | [**O11yO11yDraftFunnelIn**](O11yO11yDraftFunnelIn.md) |  | 

### Return type

[**O11yO11yFunnelRowsOut**](O11yO11yFunnelRowsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlamegraph

> O11yO11yTraceFlamegraphOut GetFlamegraph(ctx, traceId).O11yO11yTraceFlamegraphIn(o11yO11yTraceFlamegraphIn).Execute()

Returns a trace's flamegraph: spans bucketed by depth level, each level ordered as it is drawn, around the selected span.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	traceId := "traceId_example" // string | 
	o11yO11yTraceFlamegraphIn := *openapiclient.NewO11yO11yTraceFlamegraphIn() // O11yO11yTraceFlamegraphIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetFlamegraph(context.Background(), traceId).O11yO11yTraceFlamegraphIn(o11yO11yTraceFlamegraphIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetFlamegraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlamegraph`: O11yO11yTraceFlamegraphOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetFlamegraph`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlamegraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yTraceFlamegraphIn** | [**O11yO11yTraceFlamegraphIn**](O11yO11yTraceFlamegraphIn.md) |  | 

### Return type

[**O11yO11yTraceFlamegraphOut**](O11yO11yTraceFlamegraphOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHosts

> O11yO11yGettableHostOut GetHosts(ctx).Execute()

Returns the deployment's host info from Zeus.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetHosts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHosts`: O11yO11yGettableHostOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetHosts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostsRequest struct via the builder pattern


### Return type

[**O11yO11yGettableHostOut**](O11yO11yGettableHostOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIngestionKeys

> O11yO11yIngestionKeysOut GetIngestionKeys(ctx).Page(page).PerPage(perPage).Execute()

Lists the workspace's ingestion keys, paginated.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	page := int32(56) // int32 | Page is the 1-based page number. (optional)
	perPage := int32(56) // int32 | PerPage is the page size. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetIngestionKeys(context.Background()).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetIngestionKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIngestionKeys`: O11yO11yIngestionKeysOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetIngestionKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIngestionKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page is the 1-based page number. | 
 **perPage** | **int32** | PerPage is the page size. | 

### Return type

[**O11yO11yIngestionKeysOut**](O11yO11yIngestionKeysOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegration

> O11yO11yIntegrationOut GetIntegration(ctx, integrationId).Execute()

Returns one integration's full detail — its overview, configuration steps, collected data and assets — together with its installation record when the org has installed it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	integrationId := "integrationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetIntegration(context.Background(), integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegration`: O11yO11yIntegrationOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yIntegrationOut**](O11yO11yIntegrationOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationConnectionStatus

> O11yO11yConnectionStatusOut GetIntegrationConnectionStatus(ctx, integrationId).LookbackSeconds(lookbackSeconds).Execute()

Reports whether the integration's logs and metrics have been received over the lookback window, so the console can show a live connection state.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	integrationId := "integrationId_example" // string | 
	lookbackSeconds := int32(56) // int32 | LookbackSeconds is how far back to look for received telemetry, in seconds. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetIntegrationConnectionStatus(context.Background(), integrationId).LookbackSeconds(lookbackSeconds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetIntegrationConnectionStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationConnectionStatus`: O11yO11yConnectionStatusOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetIntegrationConnectionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationConnectionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lookbackSeconds** | **int32** | LookbackSeconds is how far back to look for received telemetry, in seconds. | 

### Return type

[**O11yO11yConnectionStatusOut**](O11yO11yConnectionStatusOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLLMPricingRule

> O11yO11yLLMPricingRuleOut GetLLMPricingRule(ctx, id).Execute()

Returns a single LLM pricing rule by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetLLMPricingRule(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetLLMPricingRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLLMPricingRule`: O11yO11yLLMPricingRuleOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetLLMPricingRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLLMPricingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yLLMPricingRuleOut**](O11yO11yLLMPricingRuleOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLLMScore

> O11yO11yLLMScoreOut GetLLMScore(ctx, id).Execute()

Returns a single score by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetLLMScore(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetLLMScore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLLMScore`: O11yO11yLLMScoreOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetLLMScore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLLMScoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yLLMScoreOut**](O11yO11yLLMScoreOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricAlerts

> O11yO11yMetricAlertsOut GetMetricAlerts(ctx).MetricName(metricName).Execute()

Lists the alert rules that reference a metric.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	metricName := "metricName_example" // string | MetricName is the metric's name; it may contain slashes, e.g. run.googleapis.com/request_latencies. Required.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetMetricAlerts(context.Background()).MetricName(metricName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetMetricAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricAlerts`: O11yO11yMetricAlertsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetMetricAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricName** | **string** | MetricName is the metric&#39;s name; it may contain slashes, e.g. run.googleapis.com/request_latencies. Required. | 

### Return type

[**O11yO11yMetricAlertsOut**](O11yO11yMetricAlertsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricAttributes

> O11yO11yMetricAttributesOut GetMetricAttributes(ctx).MetricName(metricName).Start(start).End(end).Execute()

Returns one metric's attribute keys, each with its unique values and their count.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	metricName := "metricName_example" // string | MetricName is the metric's name; it may contain slashes. Required.
	start := int32(56) // int32 | Start is the start of the window as a Unix timestamp in milliseconds. (optional)
	end := int32(56) // int32 | End is the end of the window as a Unix timestamp in milliseconds. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetMetricAttributes(context.Background()).MetricName(metricName).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetMetricAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricAttributes`: O11yO11yMetricAttributesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetMetricAttributes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricName** | **string** | MetricName is the metric&#39;s name; it may contain slashes. Required. | 
 **start** | **int32** | Start is the start of the window as a Unix timestamp in milliseconds. | 
 **end** | **int32** | End is the end of the window as a Unix timestamp in milliseconds. | 

### Return type

[**O11yO11yMetricAttributesOut**](O11yO11yMetricAttributesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricDashboardsV2

> O11yO11yMetricDashboardsOut GetMetricDashboardsV2(ctx).MetricName(metricName).Execute()

Lists the dashboard panels that reference a metric.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	metricName := "metricName_example" // string | MetricName is the metric's name; it may contain slashes, e.g. run.googleapis.com/request_latencies. Required.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetMetricDashboardsV2(context.Background()).MetricName(metricName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetMetricDashboardsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricDashboardsV2`: O11yO11yMetricDashboardsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetMetricDashboardsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricDashboardsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricName** | **string** | MetricName is the metric&#39;s name; it may contain slashes, e.g. run.googleapis.com/request_latencies. Required. | 

### Return type

[**O11yO11yMetricDashboardsOut**](O11yO11yMetricDashboardsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricHighlights

> O11yO11yMetricHighlightsOut GetMetricHighlights(ctx).MetricName(metricName).Execute()

Returns one metric's headline numbers: data points, total and active time series, and when it was last received.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	metricName := "metricName_example" // string | MetricName is the metric's name; it may contain slashes, e.g. run.googleapis.com/request_latencies. Required.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetMetricHighlights(context.Background()).MetricName(metricName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetMetricHighlights``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricHighlights`: O11yO11yMetricHighlightsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetMetricHighlights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricHighlightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricName** | **string** | MetricName is the metric&#39;s name; it may contain slashes, e.g. run.googleapis.com/request_latencies. Required. | 

### Return type

[**O11yO11yMetricHighlightsOut**](O11yO11yMetricHighlightsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricMetadata

> O11yO11yMetricMetadataOut GetMetricMetadata(ctx).MetricName(metricName).Execute()

Returns one metric's metadata: description, type, unit, temporality and monotonicity.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	metricName := "metricName_example" // string | MetricName is the metric's name; it may contain slashes, e.g. run.googleapis.com/request_latencies. Required.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetMetricMetadata(context.Background()).MetricName(metricName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetMetricMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricMetadata`: O11yO11yMetricMetadataOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetMetricMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricName** | **string** | MetricName is the metric&#39;s name; it may contain slashes, e.g. run.googleapis.com/request_latencies. Required. | 

### Return type

[**O11yO11yMetricMetadataOut**](O11yO11yMetricMetadataOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricReductionRuleByID

> O11yO11yReductionRuleOut GetMetricReductionRuleByID(ctx, id).Execute()

Returns one volume-control rule by its id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the rule's id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetMetricReductionRuleByID(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetMetricReductionRuleByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricReductionRuleByID`: O11yO11yReductionRuleOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetMetricReductionRuleByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the rule&#39;s id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricReductionRuleByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yReductionRuleOut**](O11yO11yReductionRuleOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricReductionRuleStats

> O11yO11yReductionStatsOut GetMetricReductionRuleStats(ctx).Execute()

Returns total ingested vs retained series and samples and the estimated monthly savings across all volume-control rules.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetMetricReductionRuleStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetMetricReductionRuleStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricReductionRuleStats`: O11yO11yReductionStatsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetMetricReductionRuleStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricReductionRuleStatsRequest struct via the builder pattern


### Return type

[**O11yO11yReductionStatsOut**](O11yO11yReductionStatsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricReductionRuleTimeseries

> O11yO11yReductionSeriesOut GetMetricReductionRuleTimeseries(ctx).Execute()

Returns ingested vs retained series over time across all volume-control rules, in hourly buckets, in the query-range time-series response shape.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetMetricReductionRuleTimeseries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetMetricReductionRuleTimeseries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricReductionRuleTimeseries`: O11yO11yReductionSeriesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetMetricReductionRuleTimeseries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricReductionRuleTimeseriesRequest struct via the builder pattern


### Return type

[**O11yO11yReductionSeriesOut**](O11yO11yReductionSeriesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricsOnboardingStatus

> O11yO11yMetricOnboardingOut GetMetricsOnboardingStatus(ctx).Execute()

Reports whether any non-O11y metrics have been ingested — the lightweight check onboarding polls.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetMetricsOnboardingStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetMetricsOnboardingStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricsOnboardingStatus`: O11yO11yMetricOnboardingOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetMetricsOnboardingStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsOnboardingStatusRequest struct via the builder pattern


### Return type

[**O11yO11yMetricOnboardingOut**](O11yO11yMetricOnboardingOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricsStats

> O11yO11yMetricStatsOut GetMetricsStats(ctx).O11yO11yMetricStatsIn(o11yO11yMetricStatsIn).Execute()

Lists metrics with their sample and time-series counts for a time range — the volume view of the metrics explorer, pageable and sortable.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yMetricStatsIn := *openapiclient.NewO11yO11yMetricStatsIn(int32(123), int32(123), int32(123)) // O11yO11yMetricStatsIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetMetricsStats(context.Background()).O11yO11yMetricStatsIn(o11yO11yMetricStatsIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetMetricsStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricsStats`: O11yO11yMetricStatsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetMetricsStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yMetricStatsIn** | [**O11yO11yMetricStatsIn**](O11yO11yMetricStatsIn.md) |  | 

### Return type

[**O11yO11yMetricStatsOut**](O11yO11yMetricStatsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricsTreemap

> O11yO11yMetricTreemapOut GetMetricsTreemap(ctx).O11yO11yMetricTreemapIn(o11yO11yMetricTreemapIn).Execute()

Returns the proportional distribution of metrics by sample count or time-series count, as the entries of a treemap.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yMetricTreemapIn := *openapiclient.NewO11yO11yMetricTreemapIn(int32(123), int32(123), "Mode_example", int32(123)) // O11yO11yMetricTreemapIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetMetricsTreemap(context.Background()).O11yO11yMetricTreemapIn(o11yO11yMetricTreemapIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetMetricsTreemap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricsTreemap`: O11yO11yMetricTreemapOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetMetricsTreemap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsTreemapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yMetricTreemapIn** | [**O11yO11yMetricTreemapIn**](O11yO11yMetricTreemapIn.md) |  | 

### Return type

[**O11yO11yMetricTreemapOut**](O11yO11yMetricTreemapOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyOrganization

> O11yO11yOrganizationOut GetMyOrganization(ctx).Execute()

Returns the caller's own organization.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetMyOrganization(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetMyOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyOrganization`: O11yO11yOrganizationOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetMyOrganization`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyOrganizationRequest struct via the builder pattern


### Return type

[**O11yO11yOrganizationOut**](O11yO11yOrganizationOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyServiceAccount

> O11yO11yServiceAccountOut GetMyServiceAccount(ctx).Execute()

Returns the calling service account itself, with the roles it holds — the self-inspection read for a key-authenticated caller.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetMyServiceAccount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetMyServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyServiceAccount`: O11yO11yServiceAccountOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetMyServiceAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyServiceAccountRequest struct via the builder pattern


### Return type

[**O11yO11yServiceAccountOut**](O11yO11yServiceAccountOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyUser

> O11yO11yUserWithRolesOut GetMyUser(ctx).Execute()

Returns the calling user together with every role they hold.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetMyUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetMyUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyUser`: O11yO11yUserWithRolesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetMyUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyUserRequest struct via the builder pattern


### Return type

[**O11yO11yUserWithRolesOut**](O11yO11yUserWithRolesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyUserDeprecated

> O11yO11yDeprecatedUserOut GetMyUserDeprecated(ctx).Execute()

Returns the calling user with their single legacy role.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetMyUserDeprecated(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetMyUserDeprecated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyUserDeprecated`: O11yO11yDeprecatedUserOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetMyUserDeprecated`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyUserDeprecatedRequest struct via the builder pattern


### Return type

[**O11yO11yDeprecatedUserOut**](O11yO11yDeprecatedUserOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yAlertsLast

> GetO11yAlertsLast(ctx).Execute()

Replay the alert records this process took



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.GetO11yAlertsLast(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yAlertsLast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yAlertsLastRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yAutocompleteAggregateAttributes

> O11yO11yAggregateAttributesOut GetO11yAutocompleteAggregateAttributes(ctx).DataSource(dataSource).AggregateOperator(aggregateOperator).SearchText(searchText).Limit(limit).Execute()

Lists the attributes usable as an aggregate target for the given telemetry and operator — what a filter builder offers after the aggregation is chosen.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	dataSource := "dataSource_example" // string | DataSource is the telemetry the attributes come from — traces, logs, metrics or meter. The runtime requires it. (optional)
	aggregateOperator := "aggregateOperator_example" // string | AggregateOperator is the aggregation the attribute will be used under, e.g. count, avg, sum. The runtime requires it for non-metrics sources. (optional)
	searchText := "searchText_example" // string | SearchText narrows the attributes to those containing it. (optional)
	limit := int32(56) // int32 | Limit caps how many attributes come back. Absent means 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yAutocompleteAggregateAttributes(context.Background()).DataSource(dataSource).AggregateOperator(aggregateOperator).SearchText(searchText).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yAutocompleteAggregateAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yAutocompleteAggregateAttributes`: O11yO11yAggregateAttributesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yAutocompleteAggregateAttributes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yAutocompleteAggregateAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | **string** | DataSource is the telemetry the attributes come from — traces, logs, metrics or meter. The runtime requires it. | 
 **aggregateOperator** | **string** | AggregateOperator is the aggregation the attribute will be used under, e.g. count, avg, sum. The runtime requires it for non-metrics sources. | 
 **searchText** | **string** | SearchText narrows the attributes to those containing it. | 
 **limit** | **int32** | Limit caps how many attributes come back. Absent means 50. | 

### Return type

[**O11yO11yAggregateAttributesOut**](O11yO11yAggregateAttributesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yAutocompleteAttributeKeys

> O11yO11yAttributeKeysOut GetO11yAutocompleteAttributeKeys(ctx).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).SearchText(searchText).TagType(tagType).Limit(limit).Execute()

Lists the attribute keys available for filtering the given telemetry, each with its data type and whether it is a materialized column.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	dataSource := "dataSource_example" // string | DataSource is the telemetry the keys come from — traces, logs, metrics or meter. The runtime requires it. (optional)
	aggregateOperator := "aggregateOperator_example" // string | AggregateOperator is the aggregation the keys will be used under. The runtime requires it for non-metrics sources. (optional)
	aggregateAttribute := "aggregateAttribute_example" // string | AggregateAttribute is the metric the keys must appear on. (optional)
	searchText := "searchText_example" // string | SearchText narrows the keys to those containing it. (optional)
	tagType := "tagType_example" // string | TagType narrows the keys to one kind — tag or resource. Empty means all; an invalid value reads as empty. (optional)
	limit := int32(56) // int32 | Limit caps how many keys come back. Absent means 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yAutocompleteAttributeKeys(context.Background()).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).SearchText(searchText).TagType(tagType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yAutocompleteAttributeKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yAutocompleteAttributeKeys`: O11yO11yAttributeKeysOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yAutocompleteAttributeKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yAutocompleteAttributeKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | **string** | DataSource is the telemetry the keys come from — traces, logs, metrics or meter. The runtime requires it. | 
 **aggregateOperator** | **string** | AggregateOperator is the aggregation the keys will be used under. The runtime requires it for non-metrics sources. | 
 **aggregateAttribute** | **string** | AggregateAttribute is the metric the keys must appear on. | 
 **searchText** | **string** | SearchText narrows the keys to those containing it. | 
 **tagType** | **string** | TagType narrows the keys to one kind — tag or resource. Empty means all; an invalid value reads as empty. | 
 **limit** | **int32** | Limit caps how many keys come back. Absent means 50. | 

### Return type

[**O11yO11yAttributeKeysOut**](O11yO11yAttributeKeysOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yAutocompleteAttributeValues

> O11yO11yAttributeValuesOut GetO11yAutocompleteAttributeValues(ctx).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).AttributeKey(attributeKey).FilterAttributeKeyDataType(filterAttributeKeyDataType).SearchText(searchText).TagType(tagType).Limit(limit).Execute()

Lists the values one attribute key has taken — string, number and bool values in their own lists — for completing a filter.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	dataSource := "dataSource_example" // string | DataSource is the telemetry the values come from — traces, logs or metrics. The runtime requires it. (optional)
	aggregateOperator := "aggregateOperator_example" // string | AggregateOperator is the aggregation the values will be used under. The runtime requires it for non-metrics sources. (optional)
	aggregateAttribute := "aggregateAttribute_example" // string | AggregateAttribute is the metric the values must appear on. (optional)
	attributeKey := "attributeKey_example" // string | AttributeKey is the key whose values to list. (optional)
	filterAttributeKeyDataType := "filterAttributeKeyDataType_example" // string | FilterAttributeKeyDataType is the key's data type — string, int64, float64 or bool. Empty means unspecified. (optional)
	searchText := "searchText_example" // string | SearchText narrows the values to those containing it. (optional)
	tagType := "tagType_example" // string | TagType narrows the search to one kind of key — tag or resource. (optional)
	limit := int32(56) // int32 | Limit caps how many values come back. Absent means 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yAutocompleteAttributeValues(context.Background()).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).AttributeKey(attributeKey).FilterAttributeKeyDataType(filterAttributeKeyDataType).SearchText(searchText).TagType(tagType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yAutocompleteAttributeValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yAutocompleteAttributeValues`: O11yO11yAttributeValuesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yAutocompleteAttributeValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yAutocompleteAttributeValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | **string** | DataSource is the telemetry the values come from — traces, logs or metrics. The runtime requires it. | 
 **aggregateOperator** | **string** | AggregateOperator is the aggregation the values will be used under. The runtime requires it for non-metrics sources. | 
 **aggregateAttribute** | **string** | AggregateAttribute is the metric the values must appear on. | 
 **attributeKey** | **string** | AttributeKey is the key whose values to list. | 
 **filterAttributeKeyDataType** | **string** | FilterAttributeKeyDataType is the key&#39;s data type — string, int64, float64 or bool. Empty means unspecified. | 
 **searchText** | **string** | SearchText narrows the values to those containing it. | 
 **tagType** | **string** | TagType narrows the search to one kind of key — tag or resource. | 
 **limit** | **int32** | Limit caps how many values come back. Absent means 50. | 

### Return type

[**O11yO11yAttributeValuesOut**](O11yO11yAttributeValuesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yAvailability

> O11yAvailabilityResponse GetO11yAvailability(ctx).Range_(range_).StepSec(stepSec).Execute()

Reports how much of the Hanzo fleet is up — the current per-service inventory plus an up-versus-reporting trend across the window.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	range_ := int32(3600) // int32 | Range is the trend window in seconds. Default 3600, capped at 604800 (7d). (optional)
	stepSec := int32(56) // int32 | StepSec is the bucket width in seconds, clamped to [30, 3600]. Absent picks ~60 buckets across the range. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yAvailability(context.Background()).Range_(range_).StepSec(stepSec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yAvailability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yAvailability`: O11yAvailabilityResponse
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yAvailability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **int32** | Range is the trend window in seconds. Default 3600, capped at 604800 (7d). | 
 **stepSec** | **int32** | StepSec is the bucket width in seconds, clamped to [30, 3600]. Absent picks ~60 buckets across the range. | 

### Return type

[**O11yAvailabilityResponse**](O11yAvailabilityResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yClustersAttributeKeys

> O11yO11yInfraAttributeKeysOut GetO11yClustersAttributeKeys(ctx).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).SearchText(searchText).TagType(tagType).Limit(limit).Execute()

Lists the metric attribute keys Kubernetes clusters report, for building cluster filters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	dataSource := "dataSource_example" // string | DataSource is the telemetry the keys come from — metrics for the infra faces. The runtime requires it. (optional)
	aggregateOperator := "aggregateOperator_example" // string | AggregateOperator is the aggregation the keys will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. (optional)
	aggregateAttribute := "aggregateAttribute_example" // string | AggregateAttribute is the metric the keys must appear on. (optional)
	searchText := "searchText_example" // string | SearchText narrows the keys to those containing it. (optional)
	tagType := "tagType_example" // string | TagType narrows the keys to one kind — tag or resource. Empty means all; an invalid value reads as empty. (optional)
	limit := int32(56) // int32 | Limit caps how many keys come back. Absent means 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yClustersAttributeKeys(context.Background()).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).SearchText(searchText).TagType(tagType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yClustersAttributeKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yClustersAttributeKeys`: O11yO11yInfraAttributeKeysOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yClustersAttributeKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yClustersAttributeKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | **string** | DataSource is the telemetry the keys come from — metrics for the infra faces. The runtime requires it. | 
 **aggregateOperator** | **string** | AggregateOperator is the aggregation the keys will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. | 
 **aggregateAttribute** | **string** | AggregateAttribute is the metric the keys must appear on. | 
 **searchText** | **string** | SearchText narrows the keys to those containing it. | 
 **tagType** | **string** | TagType narrows the keys to one kind — tag or resource. Empty means all; an invalid value reads as empty. | 
 **limit** | **int32** | Limit caps how many keys come back. Absent means 50. | 

### Return type

[**O11yO11yInfraAttributeKeysOut**](O11yO11yInfraAttributeKeysOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yClustersAttributeValues

> O11yO11yInfraAttributeValuesOut GetO11yClustersAttributeValues(ctx).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).AttributeKey(attributeKey).FilterAttributeKeyDataType(filterAttributeKeyDataType).SearchText(searchText).TagType(tagType).Limit(limit).Execute()

Lists the values one cluster attribute key has taken, for building cluster filters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	dataSource := "dataSource_example" // string | DataSource is the telemetry the values come from — metrics for the infra faces. The runtime requires it. (optional)
	aggregateOperator := "aggregateOperator_example" // string | AggregateOperator is the aggregation the values will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. (optional)
	aggregateAttribute := "aggregateAttribute_example" // string | AggregateAttribute is the metric the values must appear on. (optional)
	attributeKey := "attributeKey_example" // string | AttributeKey is the key whose values to list. (optional)
	filterAttributeKeyDataType := "filterAttributeKeyDataType_example" // string | FilterAttributeKeyDataType is the key's data type — string, int64, float64 or bool. Empty means unspecified. (optional)
	searchText := "searchText_example" // string | SearchText narrows the values to those containing it. (optional)
	tagType := "tagType_example" // string | TagType narrows the search to one kind of key — tag or resource. (optional)
	limit := int32(56) // int32 | Limit caps how many values come back. Absent means 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yClustersAttributeValues(context.Background()).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).AttributeKey(attributeKey).FilterAttributeKeyDataType(filterAttributeKeyDataType).SearchText(searchText).TagType(tagType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yClustersAttributeValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yClustersAttributeValues`: O11yO11yInfraAttributeValuesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yClustersAttributeValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yClustersAttributeValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | **string** | DataSource is the telemetry the values come from — metrics for the infra faces. The runtime requires it. | 
 **aggregateOperator** | **string** | AggregateOperator is the aggregation the values will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. | 
 **aggregateAttribute** | **string** | AggregateAttribute is the metric the values must appear on. | 
 **attributeKey** | **string** | AttributeKey is the key whose values to list. | 
 **filterAttributeKeyDataType** | **string** | FilterAttributeKeyDataType is the key&#39;s data type — string, int64, float64 or bool. Empty means unspecified. | 
 **searchText** | **string** | SearchText narrows the values to those containing it. | 
 **tagType** | **string** | TagType narrows the search to one kind of key — tag or resource. | 
 **limit** | **int32** | Limit caps how many values come back. Absent means 50. | 

### Return type

[**O11yO11yInfraAttributeValuesOut**](O11yO11yInfraAttributeValuesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yCompleteGoogle

> GetO11yCompleteGoogle(ctx).Execute()

Complete a Google sign-in



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.GetO11yCompleteGoogle(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yCompleteGoogle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yCompleteGoogleRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yCompleteOidc

> GetO11yCompleteOidc(ctx).Execute()

Complete a generic OIDC sign-in



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.GetO11yCompleteOidc(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yCompleteOidc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yCompleteOidcRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yDaemonsetsAttributeKeys

> O11yO11yInfraAttributeKeysOut GetO11yDaemonsetsAttributeKeys(ctx).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).SearchText(searchText).TagType(tagType).Limit(limit).Execute()

Lists the metric attribute keys Kubernetes daemonsets report, for building daemonset filters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	dataSource := "dataSource_example" // string | DataSource is the telemetry the keys come from — metrics for the infra faces. The runtime requires it. (optional)
	aggregateOperator := "aggregateOperator_example" // string | AggregateOperator is the aggregation the keys will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. (optional)
	aggregateAttribute := "aggregateAttribute_example" // string | AggregateAttribute is the metric the keys must appear on. (optional)
	searchText := "searchText_example" // string | SearchText narrows the keys to those containing it. (optional)
	tagType := "tagType_example" // string | TagType narrows the keys to one kind — tag or resource. Empty means all; an invalid value reads as empty. (optional)
	limit := int32(56) // int32 | Limit caps how many keys come back. Absent means 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yDaemonsetsAttributeKeys(context.Background()).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).SearchText(searchText).TagType(tagType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yDaemonsetsAttributeKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yDaemonsetsAttributeKeys`: O11yO11yInfraAttributeKeysOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yDaemonsetsAttributeKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yDaemonsetsAttributeKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | **string** | DataSource is the telemetry the keys come from — metrics for the infra faces. The runtime requires it. | 
 **aggregateOperator** | **string** | AggregateOperator is the aggregation the keys will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. | 
 **aggregateAttribute** | **string** | AggregateAttribute is the metric the keys must appear on. | 
 **searchText** | **string** | SearchText narrows the keys to those containing it. | 
 **tagType** | **string** | TagType narrows the keys to one kind — tag or resource. Empty means all; an invalid value reads as empty. | 
 **limit** | **int32** | Limit caps how many keys come back. Absent means 50. | 

### Return type

[**O11yO11yInfraAttributeKeysOut**](O11yO11yInfraAttributeKeysOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yDaemonsetsAttributeValues

> O11yO11yInfraAttributeValuesOut GetO11yDaemonsetsAttributeValues(ctx).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).AttributeKey(attributeKey).FilterAttributeKeyDataType(filterAttributeKeyDataType).SearchText(searchText).TagType(tagType).Limit(limit).Execute()

Lists the values one daemonset attribute key has taken, for building daemonset filters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	dataSource := "dataSource_example" // string | DataSource is the telemetry the values come from — metrics for the infra faces. The runtime requires it. (optional)
	aggregateOperator := "aggregateOperator_example" // string | AggregateOperator is the aggregation the values will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. (optional)
	aggregateAttribute := "aggregateAttribute_example" // string | AggregateAttribute is the metric the values must appear on. (optional)
	attributeKey := "attributeKey_example" // string | AttributeKey is the key whose values to list. (optional)
	filterAttributeKeyDataType := "filterAttributeKeyDataType_example" // string | FilterAttributeKeyDataType is the key's data type — string, int64, float64 or bool. Empty means unspecified. (optional)
	searchText := "searchText_example" // string | SearchText narrows the values to those containing it. (optional)
	tagType := "tagType_example" // string | TagType narrows the search to one kind of key — tag or resource. (optional)
	limit := int32(56) // int32 | Limit caps how many values come back. Absent means 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yDaemonsetsAttributeValues(context.Background()).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).AttributeKey(attributeKey).FilterAttributeKeyDataType(filterAttributeKeyDataType).SearchText(searchText).TagType(tagType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yDaemonsetsAttributeValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yDaemonsetsAttributeValues`: O11yO11yInfraAttributeValuesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yDaemonsetsAttributeValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yDaemonsetsAttributeValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | **string** | DataSource is the telemetry the values come from — metrics for the infra faces. The runtime requires it. | 
 **aggregateOperator** | **string** | AggregateOperator is the aggregation the values will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. | 
 **aggregateAttribute** | **string** | AggregateAttribute is the metric the values must appear on. | 
 **attributeKey** | **string** | AttributeKey is the key whose values to list. | 
 **filterAttributeKeyDataType** | **string** | FilterAttributeKeyDataType is the key&#39;s data type — string, int64, float64 or bool. Empty means unspecified. | 
 **searchText** | **string** | SearchText narrows the values to those containing it. | 
 **tagType** | **string** | TagType narrows the search to one kind of key — tag or resource. | 
 **limit** | **int32** | Limit caps how many values come back. Absent means 50. | 

### Return type

[**O11yO11yInfraAttributeValuesOut**](O11yO11yInfraAttributeValuesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yDeploymentsAttributeKeys

> O11yO11yInfraAttributeKeysOut GetO11yDeploymentsAttributeKeys(ctx).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).SearchText(searchText).TagType(tagType).Limit(limit).Execute()

Lists the metric attribute keys Kubernetes deployments report, for building deployment filters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	dataSource := "dataSource_example" // string | DataSource is the telemetry the keys come from — metrics for the infra faces. The runtime requires it. (optional)
	aggregateOperator := "aggregateOperator_example" // string | AggregateOperator is the aggregation the keys will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. (optional)
	aggregateAttribute := "aggregateAttribute_example" // string | AggregateAttribute is the metric the keys must appear on. (optional)
	searchText := "searchText_example" // string | SearchText narrows the keys to those containing it. (optional)
	tagType := "tagType_example" // string | TagType narrows the keys to one kind — tag or resource. Empty means all; an invalid value reads as empty. (optional)
	limit := int32(56) // int32 | Limit caps how many keys come back. Absent means 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yDeploymentsAttributeKeys(context.Background()).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).SearchText(searchText).TagType(tagType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yDeploymentsAttributeKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yDeploymentsAttributeKeys`: O11yO11yInfraAttributeKeysOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yDeploymentsAttributeKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yDeploymentsAttributeKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | **string** | DataSource is the telemetry the keys come from — metrics for the infra faces. The runtime requires it. | 
 **aggregateOperator** | **string** | AggregateOperator is the aggregation the keys will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. | 
 **aggregateAttribute** | **string** | AggregateAttribute is the metric the keys must appear on. | 
 **searchText** | **string** | SearchText narrows the keys to those containing it. | 
 **tagType** | **string** | TagType narrows the keys to one kind — tag or resource. Empty means all; an invalid value reads as empty. | 
 **limit** | **int32** | Limit caps how many keys come back. Absent means 50. | 

### Return type

[**O11yO11yInfraAttributeKeysOut**](O11yO11yInfraAttributeKeysOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yDeploymentsAttributeValues

> O11yO11yInfraAttributeValuesOut GetO11yDeploymentsAttributeValues(ctx).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).AttributeKey(attributeKey).FilterAttributeKeyDataType(filterAttributeKeyDataType).SearchText(searchText).TagType(tagType).Limit(limit).Execute()

Lists the values one deployment attribute key has taken, for building deployment filters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	dataSource := "dataSource_example" // string | DataSource is the telemetry the values come from — metrics for the infra faces. The runtime requires it. (optional)
	aggregateOperator := "aggregateOperator_example" // string | AggregateOperator is the aggregation the values will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. (optional)
	aggregateAttribute := "aggregateAttribute_example" // string | AggregateAttribute is the metric the values must appear on. (optional)
	attributeKey := "attributeKey_example" // string | AttributeKey is the key whose values to list. (optional)
	filterAttributeKeyDataType := "filterAttributeKeyDataType_example" // string | FilterAttributeKeyDataType is the key's data type — string, int64, float64 or bool. Empty means unspecified. (optional)
	searchText := "searchText_example" // string | SearchText narrows the values to those containing it. (optional)
	tagType := "tagType_example" // string | TagType narrows the search to one kind of key — tag or resource. (optional)
	limit := int32(56) // int32 | Limit caps how many values come back. Absent means 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yDeploymentsAttributeValues(context.Background()).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).AttributeKey(attributeKey).FilterAttributeKeyDataType(filterAttributeKeyDataType).SearchText(searchText).TagType(tagType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yDeploymentsAttributeValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yDeploymentsAttributeValues`: O11yO11yInfraAttributeValuesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yDeploymentsAttributeValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yDeploymentsAttributeValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | **string** | DataSource is the telemetry the values come from — metrics for the infra faces. The runtime requires it. | 
 **aggregateOperator** | **string** | AggregateOperator is the aggregation the values will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. | 
 **aggregateAttribute** | **string** | AggregateAttribute is the metric the values must appear on. | 
 **attributeKey** | **string** | AttributeKey is the key whose values to list. | 
 **filterAttributeKeyDataType** | **string** | FilterAttributeKeyDataType is the key&#39;s data type — string, int64, float64 or bool. Empty means unspecified. | 
 **searchText** | **string** | SearchText narrows the values to those containing it. | 
 **tagType** | **string** | TagType narrows the search to one kind of key — tag or resource. | 
 **limit** | **int32** | Limit caps how many values come back. Absent means 50. | 

### Return type

[**O11yO11yInfraAttributeValuesOut**](O11yO11yInfraAttributeValuesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yDisks

> []O11yO11yDisk GetO11yDisks(ctx).Execute()

Lists the storage disks the datastore reports, with their names and types.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yDisks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yDisks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yDisks`: []O11yO11yDisk
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yDisks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yDisksRequest struct via the builder pattern


### Return type

[**[]O11yO11yDisk**](O11yO11yDisk.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yErrorfromerrorid

> O11yO11yErrorWithSpan GetO11yErrorfromerrorid(ctx).Timestamp(timestamp).GroupID(groupID).ErrorID(errorID).Execute()

Returns one exception instance and the span it happened on, by its error id within a group at a timestamp.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	timestamp := "timestamp_example" // string | Timestamp is the instance's time as a nanosecond epoch spelled as a string. Required.
	groupID := "groupID_example" // string | GroupID is the exception group the instance belongs to. Required.
	errorID := "errorID_example" // string | ErrorID is the exception instance id. Required by errorFromErrorID and nextPrevErrorIDs; unused by errorFromGroupID. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yErrorfromerrorid(context.Background()).Timestamp(timestamp).GroupID(groupID).ErrorID(errorID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yErrorfromerrorid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yErrorfromerrorid`: O11yO11yErrorWithSpan
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yErrorfromerrorid`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yErrorfromerroridRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **string** | Timestamp is the instance&#39;s time as a nanosecond epoch spelled as a string. Required. | 
 **groupID** | **string** | GroupID is the exception group the instance belongs to. Required. | 
 **errorID** | **string** | ErrorID is the exception instance id. Required by errorFromErrorID and nextPrevErrorIDs; unused by errorFromGroupID. | 

### Return type

[**O11yO11yErrorWithSpan**](O11yO11yErrorWithSpan.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yErrorfromgroupid

> O11yO11yErrorWithSpan GetO11yErrorfromgroupid(ctx).Timestamp(timestamp).GroupID(groupID).ErrorID(errorID).Execute()

Returns the representative exception instance of a group at a timestamp, and the span it happened on.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	timestamp := "timestamp_example" // string | Timestamp is the instance's time as a nanosecond epoch spelled as a string. Required.
	groupID := "groupID_example" // string | GroupID is the exception group the instance belongs to. Required.
	errorID := "errorID_example" // string | ErrorID is the exception instance id. Required by errorFromErrorID and nextPrevErrorIDs; unused by errorFromGroupID. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yErrorfromgroupid(context.Background()).Timestamp(timestamp).GroupID(groupID).ErrorID(errorID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yErrorfromgroupid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yErrorfromgroupid`: O11yO11yErrorWithSpan
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yErrorfromgroupid`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yErrorfromgroupidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **string** | Timestamp is the instance&#39;s time as a nanosecond epoch spelled as a string. Required. | 
 **groupID** | **string** | GroupID is the exception group the instance belongs to. Required. | 
 **errorID** | **string** | ErrorID is the exception instance id. Required by errorFromErrorID and nextPrevErrorIDs; unused by errorFromGroupID. | 

### Return type

[**O11yO11yErrorWithSpan**](O11yO11yErrorWithSpan.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yErrortrackingIssues

> O11yO11yErrorIssuesOut GetO11yErrortrackingIssues(ctx).Status(status).Level(level).Environment(environment).ServiceName(serviceName).Query(query).Sort(sort).Offset(offset).Limit(limit).Execute()

Lists the caller's org's grouped error issues (by fingerprint) with status, level, counts and first/last-seen.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	status := "status_example" // string | Status narrows to one lifecycle state: unresolved, resolved or ignored. (optional)
	level := "level_example" // string | Level narrows to one severity, e.g. error, warning, info. (optional)
	environment := "environment_example" // string | Environment narrows to one deployment environment. (optional)
	serviceName := "serviceName_example" // string | ServiceName narrows to one reporting service. (optional)
	query := "query_example" // string | Query narrows to issues whose text contains it. (optional)
	sort := "sort_example" // string | Sort orders the page, e.g. lastSeen, firstSeen, count. (optional)
	offset := int32(56) // int32 | Offset is how many issues to skip. Zero starts at the first. (optional)
	limit := int32(56) // int32 | Limit caps how many issues come back. Zero means the default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yErrortrackingIssues(context.Background()).Status(status).Level(level).Environment(environment).ServiceName(serviceName).Query(query).Sort(sort).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yErrortrackingIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yErrortrackingIssues`: O11yO11yErrorIssuesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yErrortrackingIssues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yErrortrackingIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Status narrows to one lifecycle state: unresolved, resolved or ignored. | 
 **level** | **string** | Level narrows to one severity, e.g. error, warning, info. | 
 **environment** | **string** | Environment narrows to one deployment environment. | 
 **serviceName** | **string** | ServiceName narrows to one reporting service. | 
 **query** | **string** | Query narrows to issues whose text contains it. | 
 **sort** | **string** | Sort orders the page, e.g. lastSeen, firstSeen, count. | 
 **offset** | **int32** | Offset is how many issues to skip. Zero starts at the first. | 
 **limit** | **int32** | Limit caps how many issues come back. Zero means the default. | 

### Return type

[**O11yO11yErrorIssuesOut**](O11yO11yErrorIssuesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yErrortrackingIssuesById

> O11yO11yErrorGettableIssueOut GetO11yErrortrackingIssuesById(ctx, id).Execute()

Returns one grouped issue with its latest occurrence sample.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the issue id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yErrortrackingIssuesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yErrortrackingIssuesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yErrortrackingIssuesById`: O11yO11yErrorGettableIssueOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yErrortrackingIssuesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the issue id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yErrortrackingIssuesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yErrorGettableIssueOut**](O11yO11yErrorGettableIssueOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yExplorerViews

> O11yO11ySavedViewListOut GetO11yExplorerViews(ctx).SourcePage(sourcePage).Name(name).Category(category).Execute()

Lists the caller's org's saved explorer views, optionally narrowed to one source page, name or category.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	sourcePage := "sourcePage_example" // string | SourcePage narrows the views to one source page, e.g. traces, logs. (optional)
	name := "name_example" // string | Name narrows the views to one name. (optional)
	category := "category_example" // string | Category narrows the views to one category. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yExplorerViews(context.Background()).SourcePage(sourcePage).Name(name).Category(category).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yExplorerViews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yExplorerViews`: O11yO11ySavedViewListOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yExplorerViews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yExplorerViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourcePage** | **string** | SourcePage narrows the views to one source page, e.g. traces, logs. | 
 **name** | **string** | Name narrows the views to one name. | 
 **category** | **string** | Category narrows the views to one category. | 

### Return type

[**O11yO11ySavedViewListOut**](O11yO11ySavedViewListOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yExplorerViewsByViewid

> O11yO11ySavedViewOut GetO11yExplorerViewsByViewid(ctx, viewId).Execute()

Returns one saved explorer view by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	viewId := "viewId_example" // string | ViewID is the view's id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yExplorerViewsByViewid(context.Background(), viewId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yExplorerViewsByViewid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yExplorerViewsByViewid`: O11yO11ySavedViewOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yExplorerViewsByViewid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **string** | ViewID is the view&#39;s id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yExplorerViewsByViewidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11ySavedViewOut**](O11yO11ySavedViewOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yFeatures

> O11yO11yFeaturesOut GetO11yFeatures(ctx).Execute()

Returns the supported feature flags and their resolved values for the caller's org.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yFeatures(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yFeatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yFeatures`: O11yO11yFeaturesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yFeatures`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yFeaturesRequest struct via the builder pattern


### Return type

[**O11yO11yFeaturesOut**](O11yO11yFeaturesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yFieldsKeys

> O11yO11yFieldKeysOut GetO11yFieldsKeys(ctx).Signal(signal).Source(source).Limit(limit).StartUnixMilli(startUnixMilli).EndUnixMilli(endUnixMilli).FieldContext(fieldContext).FieldDataType(fieldDataType).MetricName(metricName).MetricNamespace(metricNamespace).SearchText(searchText).Execute()

Returns the telemetry field keys matching the selector — the signal's fields grouped by name, and whether the catalog is complete.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	signal := "signal_example" // string | Signal is the telemetry to read the fields of — traces, logs or metrics. (optional)
	source := "source_example" // string | Source narrows the fields to one source within the signal. (optional)
	limit := int32(56) // int32 | Limit caps how many keys come back. (optional)
	startUnixMilli := int32(56) // int32 | StartUnixMilli is the window start as a unix millisecond epoch. Zero reads as unset. (optional)
	endUnixMilli := int32(56) // int32 | EndUnixMilli is the window end as a unix millisecond epoch. Zero reads as unset. (optional)
	fieldContext := "fieldContext_example" // string | FieldContext narrows the keys to one context — resource, scope, attribute, span, log or metric. (optional)
	fieldDataType := "fieldDataType_example" // string | FieldDataType narrows the keys to one data type. (optional)
	metricName := "metricName_example" // string | MetricName narrows the keys to those on one metric. (optional)
	metricNamespace := "metricNamespace_example" // string | MetricNamespace narrows the keys to one metric namespace. (optional)
	searchText := "searchText_example" // string | SearchText narrows the keys to those containing it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yFieldsKeys(context.Background()).Signal(signal).Source(source).Limit(limit).StartUnixMilli(startUnixMilli).EndUnixMilli(endUnixMilli).FieldContext(fieldContext).FieldDataType(fieldDataType).MetricName(metricName).MetricNamespace(metricNamespace).SearchText(searchText).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yFieldsKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yFieldsKeys`: O11yO11yFieldKeysOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yFieldsKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yFieldsKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signal** | **string** | Signal is the telemetry to read the fields of — traces, logs or metrics. | 
 **source** | **string** | Source narrows the fields to one source within the signal. | 
 **limit** | **int32** | Limit caps how many keys come back. | 
 **startUnixMilli** | **int32** | StartUnixMilli is the window start as a unix millisecond epoch. Zero reads as unset. | 
 **endUnixMilli** | **int32** | EndUnixMilli is the window end as a unix millisecond epoch. Zero reads as unset. | 
 **fieldContext** | **string** | FieldContext narrows the keys to one context — resource, scope, attribute, span, log or metric. | 
 **fieldDataType** | **string** | FieldDataType narrows the keys to one data type. | 
 **metricName** | **string** | MetricName narrows the keys to those on one metric. | 
 **metricNamespace** | **string** | MetricNamespace narrows the keys to one metric namespace. | 
 **searchText** | **string** | SearchText narrows the keys to those containing it. | 

### Return type

[**O11yO11yFieldKeysOut**](O11yO11yFieldKeysOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yFieldsValues

> O11yO11yFieldValuesOut GetO11yFieldsValues(ctx).Signal(signal).Source(source).Limit(limit).StartUnixMilli(startUnixMilli).EndUnixMilli(endUnixMilli).FieldContext(fieldContext).FieldDataType(fieldDataType).MetricName(metricName).MetricNamespace(metricNamespace).SearchText(searchText).Name(name).ExistingQuery(existingQuery).Execute()

Returns the values one telemetry field has taken — string, bool, number and related values — and whether the value list is complete.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	signal := "signal_example" // string | Signal is the telemetry to read the field of — traces, logs or metrics. (optional)
	source := "source_example" // string | Source narrows the field to one source within the signal. (optional)
	limit := int32(56) // int32 | Limit caps how many values come back. (optional)
	startUnixMilli := int32(56) // int32 | StartUnixMilli is the window start as a unix millisecond epoch. Zero reads as unset. (optional)
	endUnixMilli := int32(56) // int32 | EndUnixMilli is the window end as a unix millisecond epoch. Zero reads as unset. (optional)
	fieldContext := "fieldContext_example" // string | FieldContext narrows the field to one context. (optional)
	fieldDataType := "fieldDataType_example" // string | FieldDataType narrows the field to one data type. (optional)
	metricName := "metricName_example" // string | MetricName narrows the field to one metric. (optional)
	metricNamespace := "metricNamespace_example" // string | MetricNamespace narrows the field to one metric namespace. (optional)
	searchText := "searchText_example" // string | SearchText narrows the values to those containing it. (optional)
	name := "name_example" // string | Name is the field whose values to read. (optional)
	existingQuery := "existingQuery_example" // string | ExistingQuery is the query the field appears in, so related values can be suggested for it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yFieldsValues(context.Background()).Signal(signal).Source(source).Limit(limit).StartUnixMilli(startUnixMilli).EndUnixMilli(endUnixMilli).FieldContext(fieldContext).FieldDataType(fieldDataType).MetricName(metricName).MetricNamespace(metricNamespace).SearchText(searchText).Name(name).ExistingQuery(existingQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yFieldsValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yFieldsValues`: O11yO11yFieldValuesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yFieldsValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yFieldsValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signal** | **string** | Signal is the telemetry to read the field of — traces, logs or metrics. | 
 **source** | **string** | Source narrows the field to one source within the signal. | 
 **limit** | **int32** | Limit caps how many values come back. | 
 **startUnixMilli** | **int32** | StartUnixMilli is the window start as a unix millisecond epoch. Zero reads as unset. | 
 **endUnixMilli** | **int32** | EndUnixMilli is the window end as a unix millisecond epoch. Zero reads as unset. | 
 **fieldContext** | **string** | FieldContext narrows the field to one context. | 
 **fieldDataType** | **string** | FieldDataType narrows the field to one data type. | 
 **metricName** | **string** | MetricName narrows the field to one metric. | 
 **metricNamespace** | **string** | MetricNamespace narrows the field to one metric namespace. | 
 **searchText** | **string** | SearchText narrows the values to those containing it. | 
 **name** | **string** | Name is the field whose values to read. | 
 **existingQuery** | **string** | ExistingQuery is the query the field appears in, so related values can be suggested for it. | 

### Return type

[**O11yO11yFieldValuesOut**](O11yO11yFieldValuesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yFilterSuggestions

> O11yO11yFilterSuggestionsOut GetO11yFilterSuggestions(ctx).DataSource(dataSource).SearchText(searchText).ExistingFilter(existingFilter).AttributesLimit(attributesLimit).ExamplesLimit(examplesLimit).Execute()

Suggests attribute keys and example filter queries for the query builder, seeded by what the org's own telemetry carries.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	dataSource := "dataSource_example" // string | DataSource is the signal suggestions are drawn from; only logs is supported today. Required.
	searchText := "searchText_example" // string | SearchText narrows attribute suggestions to keys containing it. (optional)
	existingFilter := "existingFilter_example" // string | ExistingFilter is the current filter set, JSON base64url-encoded, so example queries build on it rather than repeat it. (optional)
	attributesLimit := int32(56) // int32 | AttributesLimit caps how many attribute keys come back. (optional)
	examplesLimit := int32(56) // int32 | ExamplesLimit caps how many example queries come back. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yFilterSuggestions(context.Background()).DataSource(dataSource).SearchText(searchText).ExistingFilter(existingFilter).AttributesLimit(attributesLimit).ExamplesLimit(examplesLimit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yFilterSuggestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yFilterSuggestions`: O11yO11yFilterSuggestionsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yFilterSuggestions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yFilterSuggestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | **string** | DataSource is the signal suggestions are drawn from; only logs is supported today. Required. | 
 **searchText** | **string** | SearchText narrows attribute suggestions to keys containing it. | 
 **existingFilter** | **string** | ExistingFilter is the current filter set, JSON base64url-encoded, so example queries build on it rather than repeat it. | 
 **attributesLimit** | **int32** | AttributesLimit caps how many attribute keys come back. | 
 **examplesLimit** | **int32** | ExamplesLimit caps how many example queries come back. | 

### Return type

[**O11yO11yFilterSuggestionsOut**](O11yO11yFilterSuggestionsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yGlobalConfig

> O11yO11yGlobalConfigOut GetO11yGlobalConfig(ctx).Execute()

Returns the deployment's global configuration: its public endpoints and which identity providers are enabled.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yGlobalConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yGlobalConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yGlobalConfig`: O11yO11yGlobalConfigOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yGlobalConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yGlobalConfigRequest struct via the builder pattern


### Return type

[**O11yO11yGlobalConfigOut**](O11yO11yGlobalConfigOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yHealth

> O11yO11yHealthOut GetO11yHealth(ctx).Live(live).Execute()

Reports service health.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	live := true // bool | Live also checks the datastore connection; an unreachable store refuses with 503. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yHealth(context.Background()).Live(live).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yHealth`: O11yO11yHealthOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yHealth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **live** | **bool** | Live also checks the datastore connection; an unreachable store refuses with 503. | 

### Return type

[**O11yO11yHealthOut**](O11yO11yHealthOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yHealthz

> GetO11yHealthz(ctx).Execute()

Health of the observability runtime's services



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.GetO11yHealthz(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yHealthz``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yHealthzRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yHostsAttributeKeys

> O11yO11yInfraAttributeKeysOut GetO11yHostsAttributeKeys(ctx).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).SearchText(searchText).TagType(tagType).Limit(limit).Execute()

Lists the metric attribute keys hosts report, for building host filters — each with its data type and whether it is a materialized column.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	dataSource := "dataSource_example" // string | DataSource is the telemetry the keys come from — metrics for the infra faces. The runtime requires it. (optional)
	aggregateOperator := "aggregateOperator_example" // string | AggregateOperator is the aggregation the keys will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. (optional)
	aggregateAttribute := "aggregateAttribute_example" // string | AggregateAttribute is the metric the keys must appear on. (optional)
	searchText := "searchText_example" // string | SearchText narrows the keys to those containing it. (optional)
	tagType := "tagType_example" // string | TagType narrows the keys to one kind — tag or resource. Empty means all; an invalid value reads as empty. (optional)
	limit := int32(56) // int32 | Limit caps how many keys come back. Absent means 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yHostsAttributeKeys(context.Background()).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).SearchText(searchText).TagType(tagType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yHostsAttributeKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yHostsAttributeKeys`: O11yO11yInfraAttributeKeysOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yHostsAttributeKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yHostsAttributeKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | **string** | DataSource is the telemetry the keys come from — metrics for the infra faces. The runtime requires it. | 
 **aggregateOperator** | **string** | AggregateOperator is the aggregation the keys will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. | 
 **aggregateAttribute** | **string** | AggregateAttribute is the metric the keys must appear on. | 
 **searchText** | **string** | SearchText narrows the keys to those containing it. | 
 **tagType** | **string** | TagType narrows the keys to one kind — tag or resource. Empty means all; an invalid value reads as empty. | 
 **limit** | **int32** | Limit caps how many keys come back. Absent means 50. | 

### Return type

[**O11yO11yInfraAttributeKeysOut**](O11yO11yInfraAttributeKeysOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yHostsAttributeValues

> O11yO11yInfraAttributeValuesOut GetO11yHostsAttributeValues(ctx).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).AttributeKey(attributeKey).FilterAttributeKeyDataType(filterAttributeKeyDataType).SearchText(searchText).TagType(tagType).Limit(limit).Execute()

Lists the values one host attribute key has taken, for building host filters — string, number and bool values in their own lists.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	dataSource := "dataSource_example" // string | DataSource is the telemetry the values come from — metrics for the infra faces. The runtime requires it. (optional)
	aggregateOperator := "aggregateOperator_example" // string | AggregateOperator is the aggregation the values will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. (optional)
	aggregateAttribute := "aggregateAttribute_example" // string | AggregateAttribute is the metric the values must appear on. (optional)
	attributeKey := "attributeKey_example" // string | AttributeKey is the key whose values to list. (optional)
	filterAttributeKeyDataType := "filterAttributeKeyDataType_example" // string | FilterAttributeKeyDataType is the key's data type — string, int64, float64 or bool. Empty means unspecified. (optional)
	searchText := "searchText_example" // string | SearchText narrows the values to those containing it. (optional)
	tagType := "tagType_example" // string | TagType narrows the search to one kind of key — tag or resource. (optional)
	limit := int32(56) // int32 | Limit caps how many values come back. Absent means 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yHostsAttributeValues(context.Background()).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).AttributeKey(attributeKey).FilterAttributeKeyDataType(filterAttributeKeyDataType).SearchText(searchText).TagType(tagType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yHostsAttributeValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yHostsAttributeValues`: O11yO11yInfraAttributeValuesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yHostsAttributeValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yHostsAttributeValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | **string** | DataSource is the telemetry the values come from — metrics for the infra faces. The runtime requires it. | 
 **aggregateOperator** | **string** | AggregateOperator is the aggregation the values will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. | 
 **aggregateAttribute** | **string** | AggregateAttribute is the metric the values must appear on. | 
 **attributeKey** | **string** | AttributeKey is the key whose values to list. | 
 **filterAttributeKeyDataType** | **string** | FilterAttributeKeyDataType is the key&#39;s data type — string, int64, float64 or bool. Empty means unspecified. | 
 **searchText** | **string** | SearchText narrows the values to those containing it. | 
 **tagType** | **string** | TagType narrows the search to one kind of key — tag or resource. | 
 **limit** | **int32** | Limit caps how many values come back. Absent means 50. | 

### Return type

[**O11yO11yInfraAttributeValuesOut**](O11yO11yInfraAttributeValuesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yInfraMonitoringChecks

> O11yO11yInfraChecksOut GetO11yInfraMonitoringChecks(ctx).Type_(type_).Execute()

Reports whether the metrics and attributes an infra-monitoring section needs are being received — for each collector receiver or processor involved, what is present and what is missing, with a user-facing message and a docs link per missing piece.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	type_ := "type__example" // string | Type is the section to check — hosts, processes, pods, nodes, deployments, daemonsets, statefulsets, jobs, namespaces, clusters or volumes. Required.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yInfraMonitoringChecks(context.Background()).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yInfraMonitoringChecks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yInfraMonitoringChecks`: O11yO11yInfraChecksOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yInfraMonitoringChecks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yInfraMonitoringChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Type is the section to check — hosts, processes, pods, nodes, deployments, daemonsets, statefulsets, jobs, namespaces, clusters or volumes. Required. | 

### Return type

[**O11yO11yInfraChecksOut**](O11yO11yInfraChecksOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yInfraOnboardingK8sStatus

> O11yO11yOnboardingOut GetO11yInfraOnboardingK8sStatus(ctx).Execute()

Reports how far Kubernetes infra onboarding has progressed: which metric families have arrived and, per pod, which required metadata labels are present.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yInfraOnboardingK8sStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yInfraOnboardingK8sStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yInfraOnboardingK8sStatus`: O11yO11yOnboardingOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yInfraOnboardingK8sStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yInfraOnboardingK8sStatusRequest struct via the builder pattern


### Return type

[**O11yO11yOnboardingOut**](O11yO11yOnboardingOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yJobsAttributeKeys

> O11yO11yInfraAttributeKeysOut GetO11yJobsAttributeKeys(ctx).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).SearchText(searchText).TagType(tagType).Limit(limit).Execute()

Lists the metric attribute keys Kubernetes jobs report, for building job filters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	dataSource := "dataSource_example" // string | DataSource is the telemetry the keys come from — metrics for the infra faces. The runtime requires it. (optional)
	aggregateOperator := "aggregateOperator_example" // string | AggregateOperator is the aggregation the keys will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. (optional)
	aggregateAttribute := "aggregateAttribute_example" // string | AggregateAttribute is the metric the keys must appear on. (optional)
	searchText := "searchText_example" // string | SearchText narrows the keys to those containing it. (optional)
	tagType := "tagType_example" // string | TagType narrows the keys to one kind — tag or resource. Empty means all; an invalid value reads as empty. (optional)
	limit := int32(56) // int32 | Limit caps how many keys come back. Absent means 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yJobsAttributeKeys(context.Background()).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).SearchText(searchText).TagType(tagType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yJobsAttributeKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yJobsAttributeKeys`: O11yO11yInfraAttributeKeysOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yJobsAttributeKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yJobsAttributeKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | **string** | DataSource is the telemetry the keys come from — metrics for the infra faces. The runtime requires it. | 
 **aggregateOperator** | **string** | AggregateOperator is the aggregation the keys will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. | 
 **aggregateAttribute** | **string** | AggregateAttribute is the metric the keys must appear on. | 
 **searchText** | **string** | SearchText narrows the keys to those containing it. | 
 **tagType** | **string** | TagType narrows the keys to one kind — tag or resource. Empty means all; an invalid value reads as empty. | 
 **limit** | **int32** | Limit caps how many keys come back. Absent means 50. | 

### Return type

[**O11yO11yInfraAttributeKeysOut**](O11yO11yInfraAttributeKeysOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yJobsAttributeValues

> O11yO11yInfraAttributeValuesOut GetO11yJobsAttributeValues(ctx).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).AttributeKey(attributeKey).FilterAttributeKeyDataType(filterAttributeKeyDataType).SearchText(searchText).TagType(tagType).Limit(limit).Execute()

Lists the values one job attribute key has taken, for building job filters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	dataSource := "dataSource_example" // string | DataSource is the telemetry the values come from — metrics for the infra faces. The runtime requires it. (optional)
	aggregateOperator := "aggregateOperator_example" // string | AggregateOperator is the aggregation the values will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. (optional)
	aggregateAttribute := "aggregateAttribute_example" // string | AggregateAttribute is the metric the values must appear on. (optional)
	attributeKey := "attributeKey_example" // string | AttributeKey is the key whose values to list. (optional)
	filterAttributeKeyDataType := "filterAttributeKeyDataType_example" // string | FilterAttributeKeyDataType is the key's data type — string, int64, float64 or bool. Empty means unspecified. (optional)
	searchText := "searchText_example" // string | SearchText narrows the values to those containing it. (optional)
	tagType := "tagType_example" // string | TagType narrows the search to one kind of key — tag or resource. (optional)
	limit := int32(56) // int32 | Limit caps how many values come back. Absent means 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yJobsAttributeValues(context.Background()).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).AttributeKey(attributeKey).FilterAttributeKeyDataType(filterAttributeKeyDataType).SearchText(searchText).TagType(tagType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yJobsAttributeValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yJobsAttributeValues`: O11yO11yInfraAttributeValuesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yJobsAttributeValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yJobsAttributeValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | **string** | DataSource is the telemetry the values come from — metrics for the infra faces. The runtime requires it. | 
 **aggregateOperator** | **string** | AggregateOperator is the aggregation the values will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. | 
 **aggregateAttribute** | **string** | AggregateAttribute is the metric the values must appear on. | 
 **attributeKey** | **string** | AttributeKey is the key whose values to list. | 
 **filterAttributeKeyDataType** | **string** | FilterAttributeKeyDataType is the key&#39;s data type — string, int64, float64 or bool. Empty means unspecified. | 
 **searchText** | **string** | SearchText narrows the values to those containing it. | 
 **tagType** | **string** | TagType narrows the search to one kind of key — tag or resource. | 
 **limit** | **int32** | Limit caps how many values come back. Absent means 50. | 

### Return type

[**O11yO11yInfraAttributeValuesOut**](O11yO11yInfraAttributeValuesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yLicenses

> O11yO11yLicensesOut GetO11yLicenses(ctx).Execute()

Lists the org's licenses.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yLicenses(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yLicenses`: O11yO11yLicensesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yLicenses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yLicensesRequest struct via the builder pattern


### Return type

[**O11yO11yLicensesOut**](O11yO11yLicensesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yLicensesActive

> O11yO11yLicenseActiveOut GetO11yLicensesActive(ctx).Execute()

Activates the enterprise license.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yLicensesActive(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yLicensesActive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yLicensesActive`: O11yO11yLicenseActiveOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yLicensesActive`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yLicensesActiveRequest struct via the builder pattern


### Return type

[**O11yO11yLicenseActiveOut**](O11yO11yLicenseActiveOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yLivez

> GetO11yLivez(ctx).Execute()

Liveness of the observability process



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.GetO11yLivez(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yLivez``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yLivezRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yLogin

> GetO11yLogin(ctx).Execute()

Report why an observability sign-in did not complete



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.GetO11yLogin(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yLoginRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yLogs

> O11yO11yLogRecordsOut GetO11yLogs(ctx).Limit(limit).TimestampStart(timestampStart).TimestampEnd(timestampEnd).Execute()

Returns the most recent log records in the query window, newest first — each record an open object carrying its nanosecond timestamp and whatever fields the record was ingested with.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	limit := int32(56) // int32 | Limit caps how many records come back. Zero means the default of 100. (optional)
	timestampStart := int32(56) // int32 | TimestampStart is the start of the window as a nanosecond epoch. Zero means fifteen minutes before the end. (optional)
	timestampEnd := int32(56) // int32 | TimestampEnd is the end of the window as a nanosecond epoch. Zero means now. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yLogs(context.Background()).Limit(limit).TimestampStart(timestampStart).TimestampEnd(timestampEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yLogs`: O11yO11yLogRecordsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps how many records come back. Zero means the default of 100. | 
 **timestampStart** | **int32** | TimestampStart is the start of the window as a nanosecond epoch. Zero means fifteen minutes before the end. | 
 **timestampEnd** | **int32** | TimestampEnd is the end of the window as a nanosecond epoch. Zero means now. | 

### Return type

[**O11yO11yLogRecordsOut**](O11yO11yLogRecordsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yLogsAggregate

> O11yO11yLogAggregateOut GetO11yLogsAggregate(ctx).Execute()

Returns the logs aggregate buckets for the query window.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yLogsAggregate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yLogsAggregate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yLogsAggregate`: O11yO11yLogAggregateOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yLogsAggregate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yLogsAggregateRequest struct via the builder pattern


### Return type

[**O11yO11yLogAggregateOut**](O11yO11yLogAggregateOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yLogsFields

> O11yO11yFieldCatalogOut GetO11yLogsFields(ctx).Execute()

Returns the log field catalog: the fields already selected as indexed columns, and the interesting ones seen in the data that could be.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yLogsFields(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yLogsFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yLogsFields`: O11yO11yFieldCatalogOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yLogsFields`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yLogsFieldsRequest struct via the builder pattern


### Return type

[**O11yO11yFieldCatalogOut**](O11yO11yFieldCatalogOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yLogsLivetail

> GetO11yLogsLivetail(ctx).Execute()

Follow log records as they arrive



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.GetO11yLogsLivetail(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yLogsLivetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yLogsLivetailRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yLogsPipelinesByVersion

> O11yO11yLogPipelinesOut GetO11yLogsPipelinesByVersion(ctx, version).Execute()

Returns the caller's org's log parsing pipelines at one config version — \"latest\" for the newest — along with that version's deployment record and the recent version history.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	version := "version_example" // string | Version is the config version to read — a positive number, or \"latest\".

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yLogsPipelinesByVersion(context.Background(), version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yLogsPipelinesByVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yLogsPipelinesByVersion`: O11yO11yLogPipelinesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yLogsPipelinesByVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** | Version is the config version to read — a positive number, or \&quot;latest\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yLogsPipelinesByVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yLogPipelinesOut**](O11yO11yLogPipelinesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yLogsPromotePaths

> O11yO11yLogPromotedOut GetO11yLogsPromotePaths(ctx).Execute()

Lists the log body paths already promoted or indexed, with the indexes each carries.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yLogsPromotePaths(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yLogsPromotePaths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yLogsPromotePaths`: O11yO11yLogPromotedOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yLogsPromotePaths`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yLogsPromotePathsRequest struct via the builder pattern


### Return type

[**O11yO11yLogPromotedOut**](O11yO11yLogPromotedOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yMetricMetricMetadata

> O11yO11yMetricMetadataOut GetO11yMetricMetricMetadata(ctx).MetricName(metricName).ServiceName(serviceName).Execute()

Serves the OLDER /metric/metric_metadata route.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	metricName := "metricName_example" // string | MetricName is the metric to read. (optional)
	serviceName := "serviceName_example" // string | ServiceName scopes the metadata to the metric as one service reports it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yMetricMetricMetadata(context.Background()).MetricName(metricName).ServiceName(serviceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yMetricMetricMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yMetricMetricMetadata`: O11yO11yMetricMetadataOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yMetricMetricMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yMetricMetricMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricName** | **string** | MetricName is the metric to read. | 
 **serviceName** | **string** | ServiceName scopes the metadata to the metric as one service reports it. | 

### Return type

[**O11yO11yMetricMetadataOut**](O11yO11yMetricMetadataOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yNamespacesAttributeKeys

> O11yO11yInfraAttributeKeysOut GetO11yNamespacesAttributeKeys(ctx).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).SearchText(searchText).TagType(tagType).Limit(limit).Execute()

Lists the metric attribute keys Kubernetes namespaces report, for building namespace filters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	dataSource := "dataSource_example" // string | DataSource is the telemetry the keys come from — metrics for the infra faces. The runtime requires it. (optional)
	aggregateOperator := "aggregateOperator_example" // string | AggregateOperator is the aggregation the keys will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. (optional)
	aggregateAttribute := "aggregateAttribute_example" // string | AggregateAttribute is the metric the keys must appear on. (optional)
	searchText := "searchText_example" // string | SearchText narrows the keys to those containing it. (optional)
	tagType := "tagType_example" // string | TagType narrows the keys to one kind — tag or resource. Empty means all; an invalid value reads as empty. (optional)
	limit := int32(56) // int32 | Limit caps how many keys come back. Absent means 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yNamespacesAttributeKeys(context.Background()).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).SearchText(searchText).TagType(tagType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yNamespacesAttributeKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yNamespacesAttributeKeys`: O11yO11yInfraAttributeKeysOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yNamespacesAttributeKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yNamespacesAttributeKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | **string** | DataSource is the telemetry the keys come from — metrics for the infra faces. The runtime requires it. | 
 **aggregateOperator** | **string** | AggregateOperator is the aggregation the keys will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. | 
 **aggregateAttribute** | **string** | AggregateAttribute is the metric the keys must appear on. | 
 **searchText** | **string** | SearchText narrows the keys to those containing it. | 
 **tagType** | **string** | TagType narrows the keys to one kind — tag or resource. Empty means all; an invalid value reads as empty. | 
 **limit** | **int32** | Limit caps how many keys come back. Absent means 50. | 

### Return type

[**O11yO11yInfraAttributeKeysOut**](O11yO11yInfraAttributeKeysOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yNamespacesAttributeValues

> O11yO11yInfraAttributeValuesOut GetO11yNamespacesAttributeValues(ctx).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).AttributeKey(attributeKey).FilterAttributeKeyDataType(filterAttributeKeyDataType).SearchText(searchText).TagType(tagType).Limit(limit).Execute()

Lists the values one namespace attribute key has taken, for building namespace filters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	dataSource := "dataSource_example" // string | DataSource is the telemetry the values come from — metrics for the infra faces. The runtime requires it. (optional)
	aggregateOperator := "aggregateOperator_example" // string | AggregateOperator is the aggregation the values will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. (optional)
	aggregateAttribute := "aggregateAttribute_example" // string | AggregateAttribute is the metric the values must appear on. (optional)
	attributeKey := "attributeKey_example" // string | AttributeKey is the key whose values to list. (optional)
	filterAttributeKeyDataType := "filterAttributeKeyDataType_example" // string | FilterAttributeKeyDataType is the key's data type — string, int64, float64 or bool. Empty means unspecified. (optional)
	searchText := "searchText_example" // string | SearchText narrows the values to those containing it. (optional)
	tagType := "tagType_example" // string | TagType narrows the search to one kind of key — tag or resource. (optional)
	limit := int32(56) // int32 | Limit caps how many values come back. Absent means 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yNamespacesAttributeValues(context.Background()).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).AttributeKey(attributeKey).FilterAttributeKeyDataType(filterAttributeKeyDataType).SearchText(searchText).TagType(tagType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yNamespacesAttributeValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yNamespacesAttributeValues`: O11yO11yInfraAttributeValuesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yNamespacesAttributeValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yNamespacesAttributeValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | **string** | DataSource is the telemetry the values come from — metrics for the infra faces. The runtime requires it. | 
 **aggregateOperator** | **string** | AggregateOperator is the aggregation the values will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. | 
 **aggregateAttribute** | **string** | AggregateAttribute is the metric the values must appear on. | 
 **attributeKey** | **string** | AttributeKey is the key whose values to list. | 
 **filterAttributeKeyDataType** | **string** | FilterAttributeKeyDataType is the key&#39;s data type — string, int64, float64 or bool. Empty means unspecified. | 
 **searchText** | **string** | SearchText narrows the values to those containing it. | 
 **tagType** | **string** | TagType narrows the search to one kind of key — tag or resource. | 
 **limit** | **int32** | Limit caps how many values come back. Absent means 50. | 

### Return type

[**O11yO11yInfraAttributeValuesOut**](O11yO11yInfraAttributeValuesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yNextpreverrorids

> O11yO11yNextPrevErrorIDs GetO11yNextpreverrorids(ctx).Timestamp(timestamp).GroupID(groupID).ErrorID(errorID).Execute()

Returns the ids of the exception instances immediately after and before a given one within its group — the paging cursor the error detail view walks.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	timestamp := "timestamp_example" // string | Timestamp is the instance's time as a nanosecond epoch spelled as a string. Required.
	groupID := "groupID_example" // string | GroupID is the exception group the instance belongs to. Required.
	errorID := "errorID_example" // string | ErrorID is the exception instance id. Required by errorFromErrorID and nextPrevErrorIDs; unused by errorFromGroupID. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yNextpreverrorids(context.Background()).Timestamp(timestamp).GroupID(groupID).ErrorID(errorID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yNextpreverrorids``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yNextpreverrorids`: O11yO11yNextPrevErrorIDs
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yNextpreverrorids`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yNextpreverroridsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **string** | Timestamp is the instance&#39;s time as a nanosecond epoch spelled as a string. Required. | 
 **groupID** | **string** | GroupID is the exception group the instance belongs to. Required. | 
 **errorID** | **string** | ErrorID is the exception instance id. Required by errorFromErrorID and nextPrevErrorIDs; unused by errorFromGroupID. | 

### Return type

[**O11yO11yNextPrevErrorIDs**](O11yO11yNextPrevErrorIDs.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yNodesAttributeKeys

> O11yO11yInfraAttributeKeysOut GetO11yNodesAttributeKeys(ctx).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).SearchText(searchText).TagType(tagType).Limit(limit).Execute()

Lists the metric attribute keys Kubernetes nodes report, for building node filters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	dataSource := "dataSource_example" // string | DataSource is the telemetry the keys come from — metrics for the infra faces. The runtime requires it. (optional)
	aggregateOperator := "aggregateOperator_example" // string | AggregateOperator is the aggregation the keys will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. (optional)
	aggregateAttribute := "aggregateAttribute_example" // string | AggregateAttribute is the metric the keys must appear on. (optional)
	searchText := "searchText_example" // string | SearchText narrows the keys to those containing it. (optional)
	tagType := "tagType_example" // string | TagType narrows the keys to one kind — tag or resource. Empty means all; an invalid value reads as empty. (optional)
	limit := int32(56) // int32 | Limit caps how many keys come back. Absent means 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yNodesAttributeKeys(context.Background()).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).SearchText(searchText).TagType(tagType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yNodesAttributeKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yNodesAttributeKeys`: O11yO11yInfraAttributeKeysOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yNodesAttributeKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yNodesAttributeKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | **string** | DataSource is the telemetry the keys come from — metrics for the infra faces. The runtime requires it. | 
 **aggregateOperator** | **string** | AggregateOperator is the aggregation the keys will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. | 
 **aggregateAttribute** | **string** | AggregateAttribute is the metric the keys must appear on. | 
 **searchText** | **string** | SearchText narrows the keys to those containing it. | 
 **tagType** | **string** | TagType narrows the keys to one kind — tag or resource. Empty means all; an invalid value reads as empty. | 
 **limit** | **int32** | Limit caps how many keys come back. Absent means 50. | 

### Return type

[**O11yO11yInfraAttributeKeysOut**](O11yO11yInfraAttributeKeysOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yNodesAttributeValues

> O11yO11yInfraAttributeValuesOut GetO11yNodesAttributeValues(ctx).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).AttributeKey(attributeKey).FilterAttributeKeyDataType(filterAttributeKeyDataType).SearchText(searchText).TagType(tagType).Limit(limit).Execute()

Lists the values one node attribute key has taken, for building node filters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	dataSource := "dataSource_example" // string | DataSource is the telemetry the values come from — metrics for the infra faces. The runtime requires it. (optional)
	aggregateOperator := "aggregateOperator_example" // string | AggregateOperator is the aggregation the values will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. (optional)
	aggregateAttribute := "aggregateAttribute_example" // string | AggregateAttribute is the metric the values must appear on. (optional)
	attributeKey := "attributeKey_example" // string | AttributeKey is the key whose values to list. (optional)
	filterAttributeKeyDataType := "filterAttributeKeyDataType_example" // string | FilterAttributeKeyDataType is the key's data type — string, int64, float64 or bool. Empty means unspecified. (optional)
	searchText := "searchText_example" // string | SearchText narrows the values to those containing it. (optional)
	tagType := "tagType_example" // string | TagType narrows the search to one kind of key — tag or resource. (optional)
	limit := int32(56) // int32 | Limit caps how many values come back. Absent means 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yNodesAttributeValues(context.Background()).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).AttributeKey(attributeKey).FilterAttributeKeyDataType(filterAttributeKeyDataType).SearchText(searchText).TagType(tagType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yNodesAttributeValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yNodesAttributeValues`: O11yO11yInfraAttributeValuesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yNodesAttributeValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yNodesAttributeValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | **string** | DataSource is the telemetry the values come from — metrics for the infra faces. The runtime requires it. | 
 **aggregateOperator** | **string** | AggregateOperator is the aggregation the values will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. | 
 **aggregateAttribute** | **string** | AggregateAttribute is the metric the values must appear on. | 
 **attributeKey** | **string** | AttributeKey is the key whose values to list. | 
 **filterAttributeKeyDataType** | **string** | FilterAttributeKeyDataType is the key&#39;s data type — string, int64, float64 or bool. Empty means unspecified. | 
 **searchText** | **string** | SearchText narrows the values to those containing it. | 
 **tagType** | **string** | TagType narrows the search to one kind of key — tag or resource. | 
 **limit** | **int32** | Limit caps how many values come back. Absent means 50. | 

### Return type

[**O11yO11yInfraAttributeValuesOut**](O11yO11yInfraAttributeValuesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yPodsAttributeKeys

> O11yO11yInfraAttributeKeysOut GetO11yPodsAttributeKeys(ctx).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).SearchText(searchText).TagType(tagType).Limit(limit).Execute()

Lists the metric attribute keys Kubernetes pods report, for building pod filters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	dataSource := "dataSource_example" // string | DataSource is the telemetry the keys come from — metrics for the infra faces. The runtime requires it. (optional)
	aggregateOperator := "aggregateOperator_example" // string | AggregateOperator is the aggregation the keys will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. (optional)
	aggregateAttribute := "aggregateAttribute_example" // string | AggregateAttribute is the metric the keys must appear on. (optional)
	searchText := "searchText_example" // string | SearchText narrows the keys to those containing it. (optional)
	tagType := "tagType_example" // string | TagType narrows the keys to one kind — tag or resource. Empty means all; an invalid value reads as empty. (optional)
	limit := int32(56) // int32 | Limit caps how many keys come back. Absent means 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yPodsAttributeKeys(context.Background()).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).SearchText(searchText).TagType(tagType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yPodsAttributeKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yPodsAttributeKeys`: O11yO11yInfraAttributeKeysOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yPodsAttributeKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yPodsAttributeKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | **string** | DataSource is the telemetry the keys come from — metrics for the infra faces. The runtime requires it. | 
 **aggregateOperator** | **string** | AggregateOperator is the aggregation the keys will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. | 
 **aggregateAttribute** | **string** | AggregateAttribute is the metric the keys must appear on. | 
 **searchText** | **string** | SearchText narrows the keys to those containing it. | 
 **tagType** | **string** | TagType narrows the keys to one kind — tag or resource. Empty means all; an invalid value reads as empty. | 
 **limit** | **int32** | Limit caps how many keys come back. Absent means 50. | 

### Return type

[**O11yO11yInfraAttributeKeysOut**](O11yO11yInfraAttributeKeysOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yPodsAttributeValues

> O11yO11yInfraAttributeValuesOut GetO11yPodsAttributeValues(ctx).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).AttributeKey(attributeKey).FilterAttributeKeyDataType(filterAttributeKeyDataType).SearchText(searchText).TagType(tagType).Limit(limit).Execute()

Lists the values one pod attribute key has taken, for building pod filters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	dataSource := "dataSource_example" // string | DataSource is the telemetry the values come from — metrics for the infra faces. The runtime requires it. (optional)
	aggregateOperator := "aggregateOperator_example" // string | AggregateOperator is the aggregation the values will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. (optional)
	aggregateAttribute := "aggregateAttribute_example" // string | AggregateAttribute is the metric the values must appear on. (optional)
	attributeKey := "attributeKey_example" // string | AttributeKey is the key whose values to list. (optional)
	filterAttributeKeyDataType := "filterAttributeKeyDataType_example" // string | FilterAttributeKeyDataType is the key's data type — string, int64, float64 or bool. Empty means unspecified. (optional)
	searchText := "searchText_example" // string | SearchText narrows the values to those containing it. (optional)
	tagType := "tagType_example" // string | TagType narrows the search to one kind of key — tag or resource. (optional)
	limit := int32(56) // int32 | Limit caps how many values come back. Absent means 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yPodsAttributeValues(context.Background()).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).AttributeKey(attributeKey).FilterAttributeKeyDataType(filterAttributeKeyDataType).SearchText(searchText).TagType(tagType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yPodsAttributeValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yPodsAttributeValues`: O11yO11yInfraAttributeValuesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yPodsAttributeValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yPodsAttributeValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | **string** | DataSource is the telemetry the values come from — metrics for the infra faces. The runtime requires it. | 
 **aggregateOperator** | **string** | AggregateOperator is the aggregation the values will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. | 
 **aggregateAttribute** | **string** | AggregateAttribute is the metric the values must appear on. | 
 **attributeKey** | **string** | AttributeKey is the key whose values to list. | 
 **filterAttributeKeyDataType** | **string** | FilterAttributeKeyDataType is the key&#39;s data type — string, int64, float64 or bool. Empty means unspecified. | 
 **searchText** | **string** | SearchText narrows the values to those containing it. | 
 **tagType** | **string** | TagType narrows the search to one kind of key — tag or resource. | 
 **limit** | **int32** | Limit caps how many values come back. Absent means 50. | 

### Return type

[**O11yO11yInfraAttributeValuesOut**](O11yO11yInfraAttributeValuesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yProcessesAttributeKeys

> O11yO11yInfraAttributeKeysOut GetO11yProcessesAttributeKeys(ctx).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).SearchText(searchText).TagType(tagType).Limit(limit).Execute()

Lists the metric attribute keys processes report, for building process filters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	dataSource := "dataSource_example" // string | DataSource is the telemetry the keys come from — metrics for the infra faces. The runtime requires it. (optional)
	aggregateOperator := "aggregateOperator_example" // string | AggregateOperator is the aggregation the keys will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. (optional)
	aggregateAttribute := "aggregateAttribute_example" // string | AggregateAttribute is the metric the keys must appear on. (optional)
	searchText := "searchText_example" // string | SearchText narrows the keys to those containing it. (optional)
	tagType := "tagType_example" // string | TagType narrows the keys to one kind — tag or resource. Empty means all; an invalid value reads as empty. (optional)
	limit := int32(56) // int32 | Limit caps how many keys come back. Absent means 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yProcessesAttributeKeys(context.Background()).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).SearchText(searchText).TagType(tagType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yProcessesAttributeKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yProcessesAttributeKeys`: O11yO11yInfraAttributeKeysOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yProcessesAttributeKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yProcessesAttributeKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | **string** | DataSource is the telemetry the keys come from — metrics for the infra faces. The runtime requires it. | 
 **aggregateOperator** | **string** | AggregateOperator is the aggregation the keys will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. | 
 **aggregateAttribute** | **string** | AggregateAttribute is the metric the keys must appear on. | 
 **searchText** | **string** | SearchText narrows the keys to those containing it. | 
 **tagType** | **string** | TagType narrows the keys to one kind — tag or resource. Empty means all; an invalid value reads as empty. | 
 **limit** | **int32** | Limit caps how many keys come back. Absent means 50. | 

### Return type

[**O11yO11yInfraAttributeKeysOut**](O11yO11yInfraAttributeKeysOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yProcessesAttributeValues

> O11yO11yInfraAttributeValuesOut GetO11yProcessesAttributeValues(ctx).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).AttributeKey(attributeKey).FilterAttributeKeyDataType(filterAttributeKeyDataType).SearchText(searchText).TagType(tagType).Limit(limit).Execute()

Lists the values one process attribute key has taken, for building process filters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	dataSource := "dataSource_example" // string | DataSource is the telemetry the values come from — metrics for the infra faces. The runtime requires it. (optional)
	aggregateOperator := "aggregateOperator_example" // string | AggregateOperator is the aggregation the values will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. (optional)
	aggregateAttribute := "aggregateAttribute_example" // string | AggregateAttribute is the metric the values must appear on. (optional)
	attributeKey := "attributeKey_example" // string | AttributeKey is the key whose values to list. (optional)
	filterAttributeKeyDataType := "filterAttributeKeyDataType_example" // string | FilterAttributeKeyDataType is the key's data type — string, int64, float64 or bool. Empty means unspecified. (optional)
	searchText := "searchText_example" // string | SearchText narrows the values to those containing it. (optional)
	tagType := "tagType_example" // string | TagType narrows the search to one kind of key — tag or resource. (optional)
	limit := int32(56) // int32 | Limit caps how many values come back. Absent means 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yProcessesAttributeValues(context.Background()).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).AttributeKey(attributeKey).FilterAttributeKeyDataType(filterAttributeKeyDataType).SearchText(searchText).TagType(tagType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yProcessesAttributeValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yProcessesAttributeValues`: O11yO11yInfraAttributeValuesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yProcessesAttributeValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yProcessesAttributeValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | **string** | DataSource is the telemetry the values come from — metrics for the infra faces. The runtime requires it. | 
 **aggregateOperator** | **string** | AggregateOperator is the aggregation the values will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. | 
 **aggregateAttribute** | **string** | AggregateAttribute is the metric the values must appear on. | 
 **attributeKey** | **string** | AttributeKey is the key whose values to list. | 
 **filterAttributeKeyDataType** | **string** | FilterAttributeKeyDataType is the key&#39;s data type — string, int64, float64 or bool. Empty means unspecified. | 
 **searchText** | **string** | SearchText narrows the values to those containing it. | 
 **tagType** | **string** | TagType narrows the search to one kind of key — tag or resource. | 
 **limit** | **int32** | Limit caps how many values come back. Absent means 50. | 

### Return type

[**O11yO11yInfraAttributeValuesOut**](O11yO11yInfraAttributeValuesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yProductMetrics

> O11yMetricsResponse GetO11yProductMetrics(ctx).Product(product).Range_(range_).StepSec(stepSec).Execute()

Returns one product's RED series — request rate, errors, p50 and p95 latency — for the caller's org, plus that org's LLM usage rollup over the same window.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	product := "kms" // string | Product is the console product slug to read, e.g. \"kms\". Required. (optional)
	range_ := int32(3600) // int32 | Range is the window in seconds. Default 3600, capped at 604800 (7d). (optional)
	stepSec := int32(56) // int32 | StepSec is the bucket width in seconds, clamped to [30, 3600]. Absent picks ~60 buckets across the range. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yProductMetrics(context.Background()).Product(product).Range_(range_).StepSec(stepSec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yProductMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yProductMetrics`: O11yMetricsResponse
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yProductMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yProductMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **product** | **string** | Product is the console product slug to read, e.g. \&quot;kms\&quot;. Required. | 
 **range_** | **int32** | Range is the window in seconds. Default 3600, capped at 604800 (7d). | 
 **stepSec** | **int32** | StepSec is the bucket width in seconds, clamped to [30, 3600]. Absent picks ~60 buckets across the range. | 

### Return type

[**O11yMetricsResponse**](O11yMetricsResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yPvcsAttributeKeys

> O11yO11yInfraAttributeKeysOut GetO11yPvcsAttributeKeys(ctx).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).SearchText(searchText).TagType(tagType).Limit(limit).Execute()

Lists the metric attribute keys persistent volume claims report, for building volume filters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	dataSource := "dataSource_example" // string | DataSource is the telemetry the keys come from — metrics for the infra faces. The runtime requires it. (optional)
	aggregateOperator := "aggregateOperator_example" // string | AggregateOperator is the aggregation the keys will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. (optional)
	aggregateAttribute := "aggregateAttribute_example" // string | AggregateAttribute is the metric the keys must appear on. (optional)
	searchText := "searchText_example" // string | SearchText narrows the keys to those containing it. (optional)
	tagType := "tagType_example" // string | TagType narrows the keys to one kind — tag or resource. Empty means all; an invalid value reads as empty. (optional)
	limit := int32(56) // int32 | Limit caps how many keys come back. Absent means 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yPvcsAttributeKeys(context.Background()).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).SearchText(searchText).TagType(tagType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yPvcsAttributeKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yPvcsAttributeKeys`: O11yO11yInfraAttributeKeysOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yPvcsAttributeKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yPvcsAttributeKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | **string** | DataSource is the telemetry the keys come from — metrics for the infra faces. The runtime requires it. | 
 **aggregateOperator** | **string** | AggregateOperator is the aggregation the keys will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. | 
 **aggregateAttribute** | **string** | AggregateAttribute is the metric the keys must appear on. | 
 **searchText** | **string** | SearchText narrows the keys to those containing it. | 
 **tagType** | **string** | TagType narrows the keys to one kind — tag or resource. Empty means all; an invalid value reads as empty. | 
 **limit** | **int32** | Limit caps how many keys come back. Absent means 50. | 

### Return type

[**O11yO11yInfraAttributeKeysOut**](O11yO11yInfraAttributeKeysOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yPvcsAttributeValues

> O11yO11yInfraAttributeValuesOut GetO11yPvcsAttributeValues(ctx).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).AttributeKey(attributeKey).FilterAttributeKeyDataType(filterAttributeKeyDataType).SearchText(searchText).TagType(tagType).Limit(limit).Execute()

Lists the values one persistent-volume-claim attribute key has taken, for building volume filters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	dataSource := "dataSource_example" // string | DataSource is the telemetry the values come from — metrics for the infra faces. The runtime requires it. (optional)
	aggregateOperator := "aggregateOperator_example" // string | AggregateOperator is the aggregation the values will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. (optional)
	aggregateAttribute := "aggregateAttribute_example" // string | AggregateAttribute is the metric the values must appear on. (optional)
	attributeKey := "attributeKey_example" // string | AttributeKey is the key whose values to list. (optional)
	filterAttributeKeyDataType := "filterAttributeKeyDataType_example" // string | FilterAttributeKeyDataType is the key's data type — string, int64, float64 or bool. Empty means unspecified. (optional)
	searchText := "searchText_example" // string | SearchText narrows the values to those containing it. (optional)
	tagType := "tagType_example" // string | TagType narrows the search to one kind of key — tag or resource. (optional)
	limit := int32(56) // int32 | Limit caps how many values come back. Absent means 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yPvcsAttributeValues(context.Background()).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).AttributeKey(attributeKey).FilterAttributeKeyDataType(filterAttributeKeyDataType).SearchText(searchText).TagType(tagType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yPvcsAttributeValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yPvcsAttributeValues`: O11yO11yInfraAttributeValuesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yPvcsAttributeValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yPvcsAttributeValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | **string** | DataSource is the telemetry the values come from — metrics for the infra faces. The runtime requires it. | 
 **aggregateOperator** | **string** | AggregateOperator is the aggregation the values will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. | 
 **aggregateAttribute** | **string** | AggregateAttribute is the metric the values must appear on. | 
 **attributeKey** | **string** | AttributeKey is the key whose values to list. | 
 **filterAttributeKeyDataType** | **string** | FilterAttributeKeyDataType is the key&#39;s data type — string, int64, float64 or bool. Empty means unspecified. | 
 **searchText** | **string** | SearchText narrows the values to those containing it. | 
 **tagType** | **string** | TagType narrows the search to one kind of key — tag or resource. | 
 **limit** | **int32** | Limit caps how many values come back. Absent means 50. | 

### Return type

[**O11yO11yInfraAttributeValuesOut**](O11yO11yInfraAttributeValuesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yQuery

> O11yO11yPromQueryOut GetO11yQuery(ctx).Query(query).Time(time).Stats(stats).Timeout(timeout).Execute()

Evaluates one instant PromQL query against the org's metrics and returns the result at a single point in time.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	query := "query_example" // string | Query is the PromQL expression to evaluate. Required.
	time := "time_example" // string | Time is the evaluation timestamp — epoch seconds or RFC3339. Empty evaluates at now. (optional)
	stats := "stats_example" // string | Stats set to any non-empty value includes query statistics in the answer. (optional)
	timeout := "timeout_example" // string | Timeout caps evaluation time, as a duration in seconds. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yQuery(context.Background()).Query(query).Time(time).Stats(stats).Timeout(timeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yQuery`: O11yO11yPromQueryOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query is the PromQL expression to evaluate. Required. | 
 **time** | **string** | Time is the evaluation timestamp — epoch seconds or RFC3339. Empty evaluates at now. | 
 **stats** | **string** | Stats set to any non-empty value includes query statistics in the answer. | 
 **timeout** | **string** | Timeout caps evaluation time, as a duration in seconds. | 

### Return type

[**O11yO11yPromQueryOut**](O11yO11yPromQueryOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yQueryProgress

> GetO11yQueryProgress(ctx).Execute()

Watch one running query's progress



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.GetO11yQueryProgress(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yQueryProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yQueryProgressRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yQueryRange

> O11yO11yMetricsQueryRangeOut GetO11yQueryRange(ctx).Start(start).End(end).Step(step).Query(query).Stats(stats).Timeout(timeout).Execute()

Runs a Prometheus-style range query over metrics — the legacy read that predates the v5 querier — and returns the matrix, vector or scalar the query resolved to.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	start := "start_example" // string | Start is the window start — a unix timestamp (seconds, with optional fraction) or an RFC 3339 time. Required.
	end := "end_example" // string | End is the window end, in the same form as Start, and not before it. Required.
	step := "step_example" // string | Step is the query resolution, e.g. 60s, 1m, 1h — a positive duration. Required.
	query := "query_example" // string | Query is the PromQL expression to evaluate. Required.
	stats := "stats_example" // string | Stats, when \"all\", asks for query statistics alongside the result. (optional)
	timeout := "timeout_example" // string | Timeout caps how long the query may run, e.g. 30s, 1m — a positive duration. Absent means the server default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yQueryRange(context.Background()).Start(start).End(end).Step(step).Query(query).Stats(stats).Timeout(timeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yQueryRange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yQueryRange`: O11yO11yMetricsQueryRangeOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yQueryRange`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yQueryRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | Start is the window start — a unix timestamp (seconds, with optional fraction) or an RFC 3339 time. Required. | 
 **end** | **string** | End is the window end, in the same form as Start, and not before it. Required. | 
 **step** | **string** | Step is the query resolution, e.g. 60s, 1m, 1h — a positive duration. Required. | 
 **query** | **string** | Query is the PromQL expression to evaluate. Required. | 
 **stats** | **string** | Stats, when \&quot;all\&quot;, asks for query statistics alongside the result. | 
 **timeout** | **string** | Timeout caps how long the query may run, e.g. 30s, 1m — a positive duration. Absent means the server default. | 

### Return type

[**O11yO11yMetricsQueryRangeOut**](O11yO11yMetricsQueryRangeOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yReadyz

> GetO11yReadyz(ctx).Execute()

Readiness of the observability runtime to serve



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.GetO11yReadyz(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yReadyz``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yReadyzRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yReviews

> O11yAnnQueueList GetO11yReviews(ctx).Page(page).Limit(limit).Execute()

Returns a page of the caller org's human-review queues, newest first, narrowed to the caller's project.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	page := int32(1) // int32 | Page is the 1-based page to read. Default 1. (optional)
	limit := int32(20) // int32 | Limit is how many rows to return. Default 20, capped at 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yReviews(context.Background()).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yReviews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yReviews`: O11yAnnQueueList
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yReviews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yReviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page is the 1-based page to read. Default 1. | 
 **limit** | **int32** | Limit is how many rows to return. Default 20, capped at 100. | 

### Return type

[**O11yAnnQueueList**](O11yAnnQueueList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yReviewsById

> O11yAnnQueueDetailView GetO11yReviewsById(ctx, id).Execute()

Returns one review queue with its pending and completed counts and its first page of items.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "annq_1" // string | ID is the annotation queue to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yReviewsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yReviewsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yReviewsById`: O11yAnnQueueDetailView
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yReviewsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the annotation queue to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yReviewsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yAnnQueueDetailView**](O11yAnnQueueDetailView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yReviewsByIdItems

> O11yAnnItemList GetO11yReviewsByIdItems(ctx, id).Status(status).Page(page).Limit(limit).Execute()

Returns a page of one review queue's items, newest first, optionally filtered to PENDING or COMPLETED.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "annq_1" // string | ID is the annotation queue whose items to list, from the path.
	status := "PENDING" // string | Status filters to PENDING or COMPLETED items. Absent returns both. (optional)
	page := int32(56) // int32 | Page is the 1-based page to read. Default 1. (optional)
	limit := int32(56) // int32 | Limit is how many rows to return. Default 20, capped at 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yReviewsByIdItems(context.Background(), id).Status(status).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yReviewsByIdItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yReviewsByIdItems`: O11yAnnItemList
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yReviewsByIdItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the annotation queue whose items to list, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yReviewsByIdItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Status filters to PENDING or COMPLETED items. Absent returns both. | 
 **page** | **int32** | Page is the 1-based page to read. Default 1. | 
 **limit** | **int32** | Limit is how many rows to return. Default 20, capped at 100. | 

### Return type

[**O11yAnnItemList**](O11yAnnItemList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11ySentinelEventsById

> O11yO11ySentryEventOut GetO11ySentinelEventsById(ctx, id).Project(project).Execute()

Returns one captured error event of a project, by its id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the event id.
	project := "project_example" // string | Project is the project the event belongs to, by its id. Required.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11ySentinelEventsById(context.Background(), id).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11ySentinelEventsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11ySentinelEventsById`: O11yO11ySentryEventOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11ySentinelEventsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the event id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11ySentinelEventsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | **string** | Project is the project the event belongs to, by its id. Required. | 

### Return type

[**O11yO11ySentryEventOut**](O11yO11ySentryEventOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11ySentinelIssues

> O11yO11yErrorIssuesOut GetO11ySentinelIssues(ctx).Status(status).Level(level).Environment(environment).ServiceName(serviceName).Query(query).Sort(sort).Offset(offset).Limit(limit).Project(project).Period(period).Execute()

Lists the caller's org's grouped error issues, optionally narrowed to one project and one time window, and filtered by status, level, environment, service, a free-text query and a sort.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	status := "status_example" // string | Status narrows to one lifecycle state: unresolved, resolved or ignored. (optional)
	level := "level_example" // string | Level narrows to one severity, e.g. error, warning, info. (optional)
	environment := "environment_example" // string | Environment narrows to one deployment environment. (optional)
	serviceName := "serviceName_example" // string | ServiceName narrows to one reporting service. (optional)
	query := "query_example" // string | Query narrows to issues whose text contains it. (optional)
	sort := "sort_example" // string | Sort orders the page, e.g. lastSeen, firstSeen, count. (optional)
	offset := int32(56) // int32 | Offset is how many issues to skip. Zero starts at the first. (optional)
	limit := int32(56) // int32 | Limit caps how many issues come back. Zero means the default. (optional)
	project := "project_example" // string | Project narrows the org's issues to one project, by its id. (optional)
	period := "period_example" // string | Period is the window to read, relative to now — 1h, 24h, 7d, 14d, 30d. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11ySentinelIssues(context.Background()).Status(status).Level(level).Environment(environment).ServiceName(serviceName).Query(query).Sort(sort).Offset(offset).Limit(limit).Project(project).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11ySentinelIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11ySentinelIssues`: O11yO11yErrorIssuesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11ySentinelIssues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11ySentinelIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Status narrows to one lifecycle state: unresolved, resolved or ignored. | 
 **level** | **string** | Level narrows to one severity, e.g. error, warning, info. | 
 **environment** | **string** | Environment narrows to one deployment environment. | 
 **serviceName** | **string** | ServiceName narrows to one reporting service. | 
 **query** | **string** | Query narrows to issues whose text contains it. | 
 **sort** | **string** | Sort orders the page, e.g. lastSeen, firstSeen, count. | 
 **offset** | **int32** | Offset is how many issues to skip. Zero starts at the first. | 
 **limit** | **int32** | Limit caps how many issues come back. Zero means the default. | 
 **project** | **string** | Project narrows the org&#39;s issues to one project, by its id. | 
 **period** | **string** | Period is the window to read, relative to now — 1h, 24h, 7d, 14d, 30d. | 

### Return type

[**O11yO11yErrorIssuesOut**](O11yO11yErrorIssuesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11ySentinelIssuesById

> O11yO11yErrorGettableIssueOut GetO11ySentinelIssuesById(ctx, id).Execute()

Returns one grouped issue of the caller's org with its latest occurrence sample.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the issue id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11ySentinelIssuesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11ySentinelIssuesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11ySentinelIssuesById`: O11yO11yErrorGettableIssueOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11ySentinelIssuesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the issue id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11ySentinelIssuesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yErrorGettableIssueOut**](O11yO11yErrorGettableIssueOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11ySentinelIssuesByIdEvents

> O11yO11ySentryIssueEventsOut GetO11ySentinelIssuesByIdEvents(ctx, id).Project(project).Limit(limit).Execute()

Lists one issue's captured occurrences, scoped to a project — a project is an isolation unit, so the caller declares which project's occurrences to read.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the issue id.
	project := "project_example" // string | Project is the project whose occurrences to read, by its id. Required.
	limit := int32(56) // int32 | Limit caps how many occurrences come back. Zero means the default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11ySentinelIssuesByIdEvents(context.Background(), id).Project(project).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11ySentinelIssuesByIdEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11ySentinelIssuesByIdEvents`: O11yO11ySentryIssueEventsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11ySentinelIssuesByIdEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the issue id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11ySentinelIssuesByIdEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | **string** | Project is the project whose occurrences to read, by its id. Required. | 
 **limit** | **int32** | Limit caps how many occurrences come back. Zero means the default. | 

### Return type

[**O11yO11ySentryIssueEventsOut**](O11yO11ySentryIssueEventsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11ySentinelLogs

> O11yO11yLogsOut GetO11ySentinelLogs(ctx).Project(project).Query(query).Period(period).Limit(limit).Execute()

Lists a project's captured error events, newest first, optionally narrowed to those whose message or exception text contains a search string.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	project := "project_example" // string | Project is the project to read, as its id. Required.
	query := "query_example" // string | Query narrows the page to events whose text contains it. (optional)
	period := "period_example" // string | Period is the window to read, relative to now — 1h, 24h, 7d, 14d, 30d. (optional)
	limit := int32(56) // int32 | Limit caps how many events come back. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11ySentinelLogs(context.Background()).Project(project).Query(query).Period(period).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11ySentinelLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11ySentinelLogs`: O11yO11yLogsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11ySentinelLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11ySentinelLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Project is the project to read, as its id. Required. | 
 **query** | **string** | Query narrows the page to events whose text contains it. | 
 **period** | **string** | Period is the window to read, relative to now — 1h, 24h, 7d, 14d, 30d. | 
 **limit** | **int32** | Limit caps how many events come back. | 

### Return type

[**O11yO11yLogsOut**](O11yO11yLogsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11ySentinelProjects

> O11yO11ySentryProjectsOut GetO11ySentinelProjects(ctx).Execute()

Lists the caller's org's Sentry projects, each with its freshly-derived DSN.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11ySentinelProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11ySentinelProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11ySentinelProjects`: O11yO11ySentryProjectsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11ySentinelProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11ySentinelProjectsRequest struct via the builder pattern


### Return type

[**O11yO11ySentryProjectsOut**](O11yO11ySentryProjectsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11ySentinelProjectsById

> O11yO11ySentryProjectOut GetO11ySentinelProjectsById(ctx, id).Execute()

Returns one Sentry project of the caller's org, DSN included.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the project id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11ySentinelProjectsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11ySentinelProjectsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11ySentinelProjectsById`: O11yO11ySentryProjectOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11ySentinelProjectsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the project id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11ySentinelProjectsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11ySentryProjectOut**](O11yO11ySentryProjectOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11ySentinelStats

> O11yO11yStatsOut GetO11ySentinelStats(ctx).Project(project).Field(field).Period(period).Execute()

Returns a project's event-rate timeseries: one bucket per interval over the requested period, counting the events in it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	project := "project_example" // string | Project is the project to read, as its id. Required.
	field := "field_example" // string | Field is the dimension to count over. Empty counts all events. (optional)
	period := "period_example" // string | Period is the window to read, relative to now — 1h, 24h, 7d, 14d, 30d. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11ySentinelStats(context.Background()).Project(project).Field(field).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11ySentinelStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11ySentinelStats`: O11yO11yStatsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11ySentinelStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11ySentinelStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Project is the project to read, as its id. Required. | 
 **field** | **string** | Field is the dimension to count over. Empty counts all events. | 
 **period** | **string** | Period is the window to read, relative to now — 1h, 24h, 7d, 14d, 30d. | 

### Return type

[**O11yO11yStatsOut**](O11yO11yStatsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11ySentinelTraces

> O11yO11yTracesOut GetO11ySentinelTraces(ctx).Project(project).Period(period).Limit(limit).Execute()

Lists the traces a project's captured errors reference, each with how many errors landed on it, when they started and stopped, and the latest message seen — the entry point for \"which requests are failing\".



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	project := "project_example" // string | Project is the project to read, as its id. Required.
	period := "period_example" // string | Period is the window to read, relative to now — 1h, 24h, 7d, 14d, 30d. (optional)
	limit := int32(56) // int32 | Limit caps how many traces come back. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11ySentinelTraces(context.Background()).Project(project).Period(period).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11ySentinelTraces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11ySentinelTraces`: O11yO11yTracesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11ySentinelTraces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11ySentinelTracesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Project is the project to read, as its id. Required. | 
 **period** | **string** | Period is the window to read, relative to now — 1h, 24h, 7d, 14d, 30d. | 
 **limit** | **int32** | Limit caps how many traces come back. | 

### Return type

[**O11yO11yTracesOut**](O11yO11yTracesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11ySentinelTracesById

> O11yO11yTraceOut GetO11ySentinelTracesById(ctx, id).Project(project).Execute()

Returns one trace's captured errors for a project — every error event that carried the trace id, in the order the events plane holds them.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the trace id.
	project := "project_example" // string | Project is the project the trace's errors belong to. Required.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11ySentinelTracesById(context.Background(), id).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11ySentinelTracesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11ySentinelTracesById`: O11yO11yTraceOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11ySentinelTracesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the trace id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11ySentinelTracesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | **string** | Project is the project the trace&#39;s errors belong to. Required. | 

### Return type

[**O11yO11yTraceOut**](O11yO11yTraceOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yServicesList

> []string GetO11yServicesList(ctx).Execute()

Lists the name of every service the trace store holds, with no window applied — the complete catalog, for pickers and autocomplete.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yServicesList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yServicesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yServicesList`: []string
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yServicesList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yServicesListRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11ySessions

> GetO11ySessions(ctx).Execute()

List the caller org's LLM sessions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.GetO11ySessions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11ySessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11ySessionsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11ySettingsApdex

> O11yO11yApdexOut GetO11ySettingsApdex(ctx).Services(services).Execute()

Returns apdex settings for the named services.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	services := "services_example" // string | Services are the service names, comma separated. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11ySettingsApdex(context.Background()).Services(services).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11ySettingsApdex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11ySettingsApdex`: O11yO11yApdexOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11ySettingsApdex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11ySettingsApdexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **services** | **string** | Services are the service names, comma separated. | 

### Return type

[**O11yO11yApdexOut**](O11yO11yApdexOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11ySettingsTtl

> O11yO11yRetentionOut GetO11ySettingsTtl(ctx).Execute()

Returns the org's current retention policy: default TTL, custom per-label rules, and cold-storage settings where configured.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11ySettingsTtl(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11ySettingsTtl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11ySettingsTtl`: O11yO11yRetentionOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11ySettingsTtl`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11ySettingsTtlRequest struct via the builder pattern


### Return type

[**O11yO11yRetentionOut**](O11yO11yRetentionOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yStatefulsetsAttributeKeys

> O11yO11yInfraAttributeKeysOut GetO11yStatefulsetsAttributeKeys(ctx).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).SearchText(searchText).TagType(tagType).Limit(limit).Execute()

Lists the metric attribute keys Kubernetes statefulsets report, for building statefulset filters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	dataSource := "dataSource_example" // string | DataSource is the telemetry the keys come from — metrics for the infra faces. The runtime requires it. (optional)
	aggregateOperator := "aggregateOperator_example" // string | AggregateOperator is the aggregation the keys will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. (optional)
	aggregateAttribute := "aggregateAttribute_example" // string | AggregateAttribute is the metric the keys must appear on. (optional)
	searchText := "searchText_example" // string | SearchText narrows the keys to those containing it. (optional)
	tagType := "tagType_example" // string | TagType narrows the keys to one kind — tag or resource. Empty means all; an invalid value reads as empty. (optional)
	limit := int32(56) // int32 | Limit caps how many keys come back. Absent means 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yStatefulsetsAttributeKeys(context.Background()).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).SearchText(searchText).TagType(tagType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yStatefulsetsAttributeKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yStatefulsetsAttributeKeys`: O11yO11yInfraAttributeKeysOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yStatefulsetsAttributeKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yStatefulsetsAttributeKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | **string** | DataSource is the telemetry the keys come from — metrics for the infra faces. The runtime requires it. | 
 **aggregateOperator** | **string** | AggregateOperator is the aggregation the keys will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. | 
 **aggregateAttribute** | **string** | AggregateAttribute is the metric the keys must appear on. | 
 **searchText** | **string** | SearchText narrows the keys to those containing it. | 
 **tagType** | **string** | TagType narrows the keys to one kind — tag or resource. Empty means all; an invalid value reads as empty. | 
 **limit** | **int32** | Limit caps how many keys come back. Absent means 50. | 

### Return type

[**O11yO11yInfraAttributeKeysOut**](O11yO11yInfraAttributeKeysOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yStatefulsetsAttributeValues

> O11yO11yInfraAttributeValuesOut GetO11yStatefulsetsAttributeValues(ctx).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).AttributeKey(attributeKey).FilterAttributeKeyDataType(filterAttributeKeyDataType).SearchText(searchText).TagType(tagType).Limit(limit).Execute()

Lists the values one statefulset attribute key has taken, for building statefulset filters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	dataSource := "dataSource_example" // string | DataSource is the telemetry the values come from — metrics for the infra faces. The runtime requires it. (optional)
	aggregateOperator := "aggregateOperator_example" // string | AggregateOperator is the aggregation the values will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. (optional)
	aggregateAttribute := "aggregateAttribute_example" // string | AggregateAttribute is the metric the values must appear on. (optional)
	attributeKey := "attributeKey_example" // string | AttributeKey is the key whose values to list. (optional)
	filterAttributeKeyDataType := "filterAttributeKeyDataType_example" // string | FilterAttributeKeyDataType is the key's data type — string, int64, float64 or bool. Empty means unspecified. (optional)
	searchText := "searchText_example" // string | SearchText narrows the values to those containing it. (optional)
	tagType := "tagType_example" // string | TagType narrows the search to one kind of key — tag or resource. (optional)
	limit := int32(56) // int32 | Limit caps how many values come back. Absent means 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yStatefulsetsAttributeValues(context.Background()).DataSource(dataSource).AggregateOperator(aggregateOperator).AggregateAttribute(aggregateAttribute).AttributeKey(attributeKey).FilterAttributeKeyDataType(filterAttributeKeyDataType).SearchText(searchText).TagType(tagType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yStatefulsetsAttributeValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yStatefulsetsAttributeValues`: O11yO11yInfraAttributeValuesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yStatefulsetsAttributeValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yStatefulsetsAttributeValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSource** | **string** | DataSource is the telemetry the values come from — metrics for the infra faces. The runtime requires it. | 
 **aggregateOperator** | **string** | AggregateOperator is the aggregation the values will be used under, e.g. noop, count, avg. The runtime requires it for non-metrics sources. | 
 **aggregateAttribute** | **string** | AggregateAttribute is the metric the values must appear on. | 
 **attributeKey** | **string** | AttributeKey is the key whose values to list. | 
 **filterAttributeKeyDataType** | **string** | FilterAttributeKeyDataType is the key&#39;s data type — string, int64, float64 or bool. Empty means unspecified. | 
 **searchText** | **string** | SearchText narrows the values to those containing it. | 
 **tagType** | **string** | TagType narrows the search to one kind of key — tag or resource. | 
 **limit** | **int32** | Limit caps how many values come back. Absent means 50. | 

### Return type

[**O11yO11yInfraAttributeValuesOut**](O11yO11yInfraAttributeValuesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yStats

> O11yO11yOrgStatsOut GetO11yStats(ctx).Execute()

Returns the collected usage statistics for the caller's org, as the stats reporter aggregates them — a map whose keys are the reporter's own counter names.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yStats`: O11yO11yOrgStatsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yStatsRequest struct via the builder pattern


### Return type

[**O11yO11yOrgStatsOut**](O11yO11yOrgStatsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yStatus

> O11yStatusResult GetO11yStatus(ctx).Product(product).Execute()

Reports whether a product's service is live: an in-cluster health probe with its measured latency, fused with the per-replica up inventory.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	product := "kms" // string | Product is the console product slug to probe, e.g. \"kms\". Required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yStatus(context.Background()).Product(product).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yStatus`: O11yStatusResult
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **product** | **string** | Product is the console product slug to probe, e.g. \&quot;kms\&quot;. Required. | 

### Return type

[**O11yStatusResult**](O11yStatusResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11ySummary

> O11yStatusSummary GetO11ySummary(ctx).Execute()

Reports whether the platform is up.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11ySummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11ySummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11ySummary`: O11yStatusSummary
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11ySummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11ySummaryRequest struct via the builder pattern


### Return type

[**O11yStatusSummary**](O11yStatusSummary.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yTraces

> O11yTracesOut GetO11yTraces(ctx).Range_(range_).Limit(limit).MinDurationMs(minDurationMs).Execute()

Lists the caller org's recent traces — one row per trace with its span count and wall-clock duration, most recently active first.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	range_ := int32(3600) // int32 | Range is the window in seconds, counted back from now over each trace's last activity. Default 3600, capped at 604800 (7d). (optional)
	limit := int32(50) // int32 | Limit is how many traces to return. Default 50, capped at 500. (optional)
	minDurationMs := int32(56) // int32 | MinDurationMs keeps only traces that lasted at least this many milliseconds. Zero or absent keeps every trace in the window. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yTraces(context.Background()).Range_(range_).Limit(limit).MinDurationMs(minDurationMs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yTraces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yTraces`: O11yTracesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yTraces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yTracesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **int32** | Range is the window in seconds, counted back from now over each trace&#39;s last activity. Default 3600, capped at 604800 (7d). | 
 **limit** | **int32** | Limit is how many traces to return. Default 50, capped at 500. | 
 **minDurationMs** | **int32** | MinDurationMs keeps only traces that lasted at least this many milliseconds. Zero or absent keeps every trace in the window. | 

### Return type

[**O11yTracesOut**](O11yTracesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yUsage

> []O11yO11yUsageItem GetO11yUsage(ctx).Start(start).End(end).Step(step).Service(service).Execute()

Returns ingestion usage counts bucketed over the requested window, optionally narrowed to one service.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	start := "start_example" // string | Start is the window start, as epoch nanoseconds. Required.
	end := "end_example" // string | End is the window end, as epoch nanoseconds. Required.
	step := int32(56) // int32 | Step is the bucket width in seconds. The runtime requires it. (optional)
	service := "service_example" // string | Service narrows usage to one service. Empty covers all. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yUsage(context.Background()).Start(start).End(end).Step(step).Service(service).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yUsage`: []O11yO11yUsageItem
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | Start is the window start, as epoch nanoseconds. Required. | 
 **end** | **string** | End is the window end, as epoch nanoseconds. Required. | 
 **step** | **int32** | Step is the bucket width in seconds. The runtime requires it. | 
 **service** | **string** | Service narrows usage to one service. Empty covers all. | 

### Return type

[**[]O11yO11yUsageItem**](O11yO11yUsageItem.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetO11yVersion

> O11yO11yVersionOut GetO11yVersion(ctx).Execute()

Reports the running build: its version, whether an enterprise edition is present (\"N\" in this build), and whether first-user setup has completed.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetO11yVersion(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetO11yVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetO11yVersion`: O11yO11yVersionOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetO11yVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetO11yVersionRequest struct via the builder pattern


### Return type

[**O11yO11yVersionOut**](O11yO11yVersionOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgPreference

> O11yO11yPreferenceOut GetOrgPreference(ctx, name).Execute()

Returns one org-scoped preference, by name.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetOrgPreference(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetOrgPreference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgPreference`: O11yO11yPreferenceOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetOrgPreference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgPreferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yPreferenceOut**](O11yO11yPreferenceOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOverallStateTransitions

> O11yO11yOverallStateTransitionsOut GetOverallStateTransitions(ctx, id).O11yO11yRuleHistoryQueryIn(o11yO11yRuleHistoryQueryIn).Execute()

Returns the overall firing/inactive windows for a rule, for the posted query range.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 
	o11yO11yRuleHistoryQueryIn := *openapiclient.NewO11yO11yRuleHistoryQueryIn() // O11yO11yRuleHistoryQueryIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetOverallStateTransitions(context.Background(), id).O11yO11yRuleHistoryQueryIn(o11yO11yRuleHistoryQueryIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetOverallStateTransitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOverallStateTransitions`: O11yO11yOverallStateTransitionsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetOverallStateTransitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOverallStateTransitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yRuleHistoryQueryIn** | [**O11yO11yRuleHistoryQueryIn**](O11yO11yRuleHistoryQueryIn.md) |  | 

### Return type

[**O11yO11yOverallStateTransitionsOut**](O11yO11yOverallStateTransitionsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicDashboard

> O11yO11yPublicDashboardOut GetPublicDashboard(ctx, id).Execute()

Returns the public-sharing config for a dashboard.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the resource id from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetPublicDashboard(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetPublicDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicDashboard`: O11yO11yPublicDashboardOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetPublicDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the resource id from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yPublicDashboardOut**](O11yO11yPublicDashboardOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicDashboardData

> O11yO11yPublicDashboardDataOut GetPublicDashboardData(ctx, id).Execute()

Returns the sanitized dashboard data for public access — the read a shared dashboard's public page makes.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the resource id from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetPublicDashboardData(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetPublicDashboardData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicDashboardData`: O11yO11yPublicDashboardDataOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetPublicDashboardData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the resource id from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicDashboardDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yPublicDashboardDataOut**](O11yO11yPublicDashboardDataOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicDashboardWidgetQueryRange

> O11yO11yWidgetQueryRangeOut GetPublicDashboardWidgetQueryRange(ctx, id, idx).StartTime(startTime).EndTime(endTime).Execute()

Returns the query-range result for one widget of a public dashboard.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the public dashboard id from the path.
	idx := "idx_example" // string | Idx is the widget's index from the path.
	startTime := "startTime_example" // string | StartTime is the window start as a millisecond epoch. Used only when the share enables a caller-chosen time range. (optional)
	endTime := "endTime_example" // string | EndTime is the window end as a millisecond epoch. Used only when the share enables a caller-chosen time range. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetPublicDashboardWidgetQueryRange(context.Background(), id, idx).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetPublicDashboardWidgetQueryRange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicDashboardWidgetQueryRange`: O11yO11yWidgetQueryRangeOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetPublicDashboardWidgetQueryRange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the public dashboard id from the path. | 
**idx** | **string** | Idx is the widget&#39;s index from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicDashboardWidgetQueryRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startTime** | **string** | StartTime is the window start as a millisecond epoch. Used only when the share enables a caller-chosen time range. | 
 **endTime** | **string** | EndTime is the window end as a millisecond epoch. Used only when the share enables a caller-chosen time range. | 

### Return type

[**O11yO11yWidgetQueryRangeOut**](O11yO11yWidgetQueryRangeOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuickFilters

> O11yO11yQuickFiltersOut GetQuickFilters(ctx).Execute()

Returns the org's quick filters for every signal — the attribute shortlists its explorers offer as one-click filters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetQuickFilters(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetQuickFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuickFilters`: O11yO11yQuickFiltersOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetQuickFilters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuickFiltersRequest struct via the builder pattern


### Return type

[**O11yO11yQuickFiltersOut**](O11yO11yQuickFiltersOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResetPasswordToken

> O11yO11yResetTokenOut GetResetPasswordToken(ctx, id).Execute()

Returns the reset-password token a user already has; absent one, the answer is a not-found rather than a fresh token.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetResetPasswordToken(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetResetPasswordToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResetPasswordToken`: O11yO11yResetTokenOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetResetPasswordToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResetPasswordTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yResetTokenOut**](O11yO11yResetTokenOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResetPasswordTokenDeprecated

> O11yO11yResetTokenOut GetResetPasswordTokenDeprecated(ctx, id).Execute()

Returns a user's password-reset token, creating one if none is live.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetResetPasswordTokenDeprecated(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetResetPasswordTokenDeprecated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResetPasswordTokenDeprecated`: O11yO11yResetTokenOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetResetPasswordTokenDeprecated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResetPasswordTokenDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yResetTokenOut**](O11yO11yResetTokenOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> O11yO11yRoleOut GetRole(ctx, id).Execute()

Returns one role with the transaction groups it grants.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetRole(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRole`: O11yO11yRoleOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yRoleOut**](O11yO11yRoleOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRolesByUserID

> O11yO11yRolesOut GetRolesByUserID(ctx, id).Execute()

Returns every role one org member holds, by user id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetRolesByUserID(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetRolesByUserID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRolesByUserID`: O11yO11yRolesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetRolesByUserID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRolesByUserIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yRolesOut**](O11yO11yRolesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoutePolicyByID

> O11yO11yRoutePolicyOut GetRoutePolicyByID(ctx, id).Execute()

Returns one route policy, by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetRoutePolicyByID(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetRoutePolicyByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoutePolicyByID`: O11yO11yRoutePolicyOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetRoutePolicyByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoutePolicyByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yRoutePolicyOut**](O11yO11yRoutePolicyOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleByID

> O11yO11yRuleOut GetRuleByID(ctx, id).Execute()

Returns one alert rule with its evaluation state, by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetRuleByID(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetRuleByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleByID`: O11yO11yRuleOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetRuleByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yRuleOut**](O11yO11yRuleOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleHistoryFilterKeys

> O11yO11yRuleHistoryFilterKeysOut GetRuleHistoryFilterKeys(ctx, id).StartUnixMilli(startUnixMilli).EndUnixMilli(endUnixMilli).SearchText(searchText).Limit(limit).Execute()

Returns the distinct label keys present in a rule's history entries over the selected range, for building history filters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 
	startUnixMilli := int32(56) // int32 | StartUnixMilli is the window start, unix milliseconds. (optional)
	endUnixMilli := int32(56) // int32 | EndUnixMilli is the window end, unix milliseconds. (optional)
	searchText := "searchText_example" // string | SearchText narrows the keys to those containing it. (optional)
	limit := int32(56) // int32 | Limit caps how many keys come back. Absent means 50, capped at 200. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetRuleHistoryFilterKeys(context.Background(), id).StartUnixMilli(startUnixMilli).EndUnixMilli(endUnixMilli).SearchText(searchText).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetRuleHistoryFilterKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleHistoryFilterKeys`: O11yO11yRuleHistoryFilterKeysOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetRuleHistoryFilterKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleHistoryFilterKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startUnixMilli** | **int32** | StartUnixMilli is the window start, unix milliseconds. | 
 **endUnixMilli** | **int32** | EndUnixMilli is the window end, unix milliseconds. | 
 **searchText** | **string** | SearchText narrows the keys to those containing it. | 
 **limit** | **int32** | Limit caps how many keys come back. Absent means 50, capped at 200. | 

### Return type

[**O11yO11yRuleHistoryFilterKeysOut**](O11yO11yRuleHistoryFilterKeysOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleHistoryFilterValues

> O11yO11yRuleHistoryFilterValuesOut GetRuleHistoryFilterValues(ctx, id).Name(name).StartUnixMilli(startUnixMilli).EndUnixMilli(endUnixMilli).SearchText(searchText).Limit(limit).ExistingQuery(existingQuery).Execute()

Returns the distinct values a given label key has taken across a rule's history entries.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 
	name := "name_example" // string | Name is the label key whose values to list. Required.
	startUnixMilli := int32(56) // int32 |  (optional)
	endUnixMilli := int32(56) // int32 |  (optional)
	searchText := "searchText_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	existingQuery := "existingQuery_example" // string | ExistingQuery is a filter expression scoping which values appear. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetRuleHistoryFilterValues(context.Background(), id).Name(name).StartUnixMilli(startUnixMilli).EndUnixMilli(endUnixMilli).SearchText(searchText).Limit(limit).ExistingQuery(existingQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetRuleHistoryFilterValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleHistoryFilterValues`: O11yO11yRuleHistoryFilterValuesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetRuleHistoryFilterValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleHistoryFilterValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Name is the label key whose values to list. Required. | 
 **startUnixMilli** | **int32** |  | 
 **endUnixMilli** | **int32** |  | 
 **searchText** | **string** |  | 
 **limit** | **int32** |  | 
 **existingQuery** | **string** | ExistingQuery is a filter expression scoping which values appear. | 

### Return type

[**O11yO11yRuleHistoryFilterValuesOut**](O11yO11yRuleHistoryFilterValuesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleHistoryOverallStatus

> O11yO11yRuleHistoryOverallStatusOut GetRuleHistoryOverallStatus(ctx, id).Start(start).End(end).Execute()

Returns the overall firing/inactive intervals for a rule over the selected range.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 
	start := int32(56) // int32 | Start is the window start, unix milliseconds. Required by the runtime. (optional)
	end := int32(56) // int32 | End is the window end, unix milliseconds. Required by the runtime. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetRuleHistoryOverallStatus(context.Background(), id).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetRuleHistoryOverallStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleHistoryOverallStatus`: O11yO11yRuleHistoryOverallStatusOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetRuleHistoryOverallStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleHistoryOverallStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int32** | Start is the window start, unix milliseconds. Required by the runtime. | 
 **end** | **int32** | End is the window end, unix milliseconds. Required by the runtime. | 

### Return type

[**O11yO11yRuleHistoryOverallStatusOut**](O11yO11yRuleHistoryOverallStatusOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleHistoryStats

> O11yO11yRuleHistoryStatsOut GetRuleHistoryStats(ctx, id).Start(start).End(end).Execute()

Returns trigger and resolution statistics for a rule over the selected time range, current window against the prior one.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 
	start := int32(56) // int32 | Start is the window start, unix milliseconds. Required by the runtime. (optional)
	end := int32(56) // int32 | End is the window end, unix milliseconds. Required by the runtime. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetRuleHistoryStats(context.Background(), id).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetRuleHistoryStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleHistoryStats`: O11yO11yRuleHistoryStatsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetRuleHistoryStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleHistoryStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int32** | Start is the window start, unix milliseconds. Required by the runtime. | 
 **end** | **int32** | End is the window end, unix milliseconds. Required by the runtime. | 

### Return type

[**O11yO11yRuleHistoryStatsOut**](O11yO11yRuleHistoryStatsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleHistoryTimeline

> O11yO11yRuleHistoryTimelineOut GetRuleHistoryTimeline(ctx, id).Start(start).End(end).State(state).FilterExpression(filterExpression).Limit(limit).Order(order).Cursor(cursor).Execute()

Returns paginated timeline entries for a rule's state transitions, filterable by state and a label expression, cursor-paginated.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 
	start := int32(56) // int32 | Start is the window start, unix milliseconds. Required by the runtime. (optional)
	end := int32(56) // int32 | End is the window end, unix milliseconds. Required by the runtime. (optional)
	state := "state_example" // string | State keeps only entries in one alert state, e.g. firing or normal. (optional)
	filterExpression := "filterExpression_example" // string | FilterExpression narrows entries to those whose labels match it. (optional)
	limit := int32(56) // int32 | Limit caps how many entries come back. Absent means 50. (optional)
	order := "order_example" // string | Order sorts by time, asc or desc. (optional)
	cursor := "cursor_example" // string | Cursor resumes a previous page; opaque, returned as nextCursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetRuleHistoryTimeline(context.Background(), id).Start(start).End(end).State(state).FilterExpression(filterExpression).Limit(limit).Order(order).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetRuleHistoryTimeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleHistoryTimeline`: O11yO11yRuleHistoryTimelineOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetRuleHistoryTimeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleHistoryTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int32** | Start is the window start, unix milliseconds. Required by the runtime. | 
 **end** | **int32** | End is the window end, unix milliseconds. Required by the runtime. | 
 **state** | **string** | State keeps only entries in one alert state, e.g. firing or normal. | 
 **filterExpression** | **string** | FilterExpression narrows entries to those whose labels match it. | 
 **limit** | **int32** | Limit caps how many entries come back. Absent means 50. | 
 **order** | **string** | Order sorts by time, asc or desc. | 
 **cursor** | **string** | Cursor resumes a previous page; opaque, returned as nextCursor. | 

### Return type

[**O11yO11yRuleHistoryTimelineOut**](O11yO11yRuleHistoryTimelineOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleHistoryTopContributors

> O11yO11yRuleHistoryContributorsOut GetRuleHistoryTopContributors(ctx, id).Start(start).End(end).Execute()

Returns the label combinations that contributed most to a rule firing over the selected range.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 
	start := int32(56) // int32 | Start is the window start, unix milliseconds. Required by the runtime. (optional)
	end := int32(56) // int32 | End is the window end, unix milliseconds. Required by the runtime. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetRuleHistoryTopContributors(context.Background(), id).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetRuleHistoryTopContributors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleHistoryTopContributors`: O11yO11yRuleHistoryContributorsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetRuleHistoryTopContributors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleHistoryTopContributorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int32** | Start is the window start, unix milliseconds. Required by the runtime. | 
 **end** | **int32** | End is the window end, unix milliseconds. Required by the runtime. | 

### Return type

[**O11yO11yRuleHistoryContributorsOut**](O11yO11yRuleHistoryContributorsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleStateHistory

> O11yO11yRuleStateTimelineOut GetRuleStateHistory(ctx, id).O11yO11yRuleHistoryQueryIn(o11yO11yRuleHistoryQueryIn).Execute()

Returns a rule's state-transition timeline for the posted query range, each entry carrying its related-logs or related-traces link.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 
	o11yO11yRuleHistoryQueryIn := *openapiclient.NewO11yO11yRuleHistoryQueryIn() // O11yO11yRuleHistoryQueryIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetRuleStateHistory(context.Background(), id).O11yO11yRuleHistoryQueryIn(o11yO11yRuleHistoryQueryIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetRuleStateHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleStateHistory`: O11yO11yRuleStateTimelineOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetRuleStateHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleStateHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yRuleHistoryQueryIn** | [**O11yO11yRuleHistoryQueryIn**](O11yO11yRuleHistoryQueryIn.md) |  | 

### Return type

[**O11yO11yRuleStateTimelineOut**](O11yO11yRuleStateTimelineOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleStateHistoryTopContributors

> O11yO11yRuleStateContributorsOut GetRuleStateHistoryTopContributors(ctx, id).O11yO11yRuleHistoryQueryIn(o11yO11yRuleHistoryQueryIn).Execute()

Returns the label combinations that contributed most to a rule firing, for the posted query range.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 
	o11yO11yRuleHistoryQueryIn := *openapiclient.NewO11yO11yRuleHistoryQueryIn() // O11yO11yRuleHistoryQueryIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetRuleStateHistoryTopContributors(context.Background(), id).O11yO11yRuleHistoryQueryIn(o11yO11yRuleHistoryQueryIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetRuleStateHistoryTopContributors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleStateHistoryTopContributors`: O11yO11yRuleStateContributorsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetRuleStateHistoryTopContributors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleStateHistoryTopContributorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yRuleHistoryQueryIn** | [**O11yO11yRuleHistoryQueryIn**](O11yO11yRuleHistoryQueryIn.md) |  | 

### Return type

[**O11yO11yRuleStateContributorsOut**](O11yO11yRuleStateContributorsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleStats

> O11yO11yRuleStatsOut GetRuleStats(ctx, id).O11yO11yRuleHistoryQueryIn(o11yO11yRuleHistoryQueryIn).Execute()

Returns trigger and resolution statistics for a rule, current window against the prior one, for the posted query range.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 
	o11yO11yRuleHistoryQueryIn := *openapiclient.NewO11yO11yRuleHistoryQueryIn() // O11yO11yRuleHistoryQueryIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetRuleStats(context.Background(), id).O11yO11yRuleHistoryQueryIn(o11yO11yRuleHistoryQueryIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetRuleStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleStats`: O11yO11yRuleStatsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetRuleStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yRuleHistoryQueryIn** | [**O11yO11yRuleHistoryQueryIn**](O11yO11yRuleHistoryQueryIn.md) |  | 

### Return type

[**O11yO11yRuleStatsOut**](O11yO11yRuleStatsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetService

> O11yO11yServiceOut GetService(ctx, cloudProvider, serviceId).CloudIntegrationId(cloudIntegrationId).Execute()

Returns one service the given provider can collect from, by service id, optionally scoped to one cloud integration.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	cloudProvider := "cloudProvider_example" // string | 
	serviceId := "serviceId_example" // string | 
	cloudIntegrationId := "cloudIntegrationId_example" // string | CloudIntegrationID, when set, scopes the service to one cloud integration. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetService(context.Background(), cloudProvider, serviceId).CloudIntegrationId(cloudIntegrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetService`: O11yO11yServiceOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** |  | 
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cloudIntegrationId** | **string** | CloudIntegrationID, when set, scopes the service to one cloud integration. | 

### Return type

[**O11yO11yServiceOut**](O11yO11yServiceOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceAccount

> O11yO11yServiceAccountOut GetServiceAccount(ctx, id).Execute()

Returns one service account with the roles it holds.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetServiceAccount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceAccount`: O11yO11yServiceAccountOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yServiceAccountOut**](O11yO11yServiceAccountOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceAccountRoles

> O11yO11yServiceAccountRolesOut GetServiceAccountRoles(ctx, id).Execute()

Lists the roles a service account holds.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetServiceAccountRoles(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetServiceAccountRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceAccountRoles`: O11yO11yServiceAccountRolesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetServiceAccountRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAccountRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yServiceAccountRolesOut**](O11yO11yServiceAccountRolesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionContext

> O11yO11ySessionContextOut GetSessionContext(ctx).Email(email).Ref(ref).Execute()

Tells a sign-in page what an email address can do: which orgs the address belongs to and, per org, which password and SSO routes are open to it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	email := "email_example" // string | Email is the address about to sign in. Required. (optional)
	ref := "ref_example" // string | Ref is the page the sign-in started from, carried into SSO redirects. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetSessionContext(context.Background()).Email(email).Ref(ref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetSessionContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSessionContext`: O11yO11ySessionContextOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetSessionContext`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Email is the address about to sign in. Required. | 
 **ref** | **string** | Ref is the page the sign-in started from, carried into SSO redirects. | 

### Return type

[**O11yO11ySessionContextOut**](O11yO11ySessionContextOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignalFilters

> O11yO11ySignalFiltersOut GetSignalFilters(ctx, signal).Execute()

Returns the org's quick filters for one signal — traces, logs, metrics, exceptions or api_monitoring.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	signal := "signal_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetSignalFilters(context.Background(), signal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetSignalFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSignalFilters`: O11yO11ySignalFiltersOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetSignalFilters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**signal** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSignalFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11ySignalFiltersOut**](O11yO11ySignalFiltersOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTraceAggregations

> O11yO11yTraceAggregationsOut GetTraceAggregations(ctx, traceId).O11yO11yTraceAggregationsIn(o11yO11yTraceAggregationsIn).Execute()

Computes span aggregations over one trace — span count, duration or share of execution time — grouped by the resource field each aggregation names.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	traceId := "traceId_example" // string | 
	o11yO11yTraceAggregationsIn := *openapiclient.NewO11yO11yTraceAggregationsIn() // O11yO11yTraceAggregationsIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetTraceAggregations(context.Background(), traceId).O11yO11yTraceAggregationsIn(o11yO11yTraceAggregationsIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetTraceAggregations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTraceAggregations`: O11yO11yTraceAggregationsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetTraceAggregations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceAggregationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yTraceAggregationsIn** | [**O11yO11yTraceAggregationsIn**](O11yO11yTraceAggregationsIn.md) |  | 

### Return type

[**O11yO11yTraceAggregationsOut**](O11yO11yTraceAggregationsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTraceFields

> O11yO11yFieldCatalogOut GetTraceFields(ctx).Execute()

Returns the trace field catalog: the span fields already selected as indexed columns, and the interesting ones seen in the data that could be.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetTraceFields(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetTraceFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTraceFields`: O11yO11yFieldCatalogOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetTraceFields`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceFieldsRequest struct via the builder pattern


### Return type

[**O11yO11yFieldCatalogOut**](O11yO11yFieldCatalogOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTraceFunnel

> O11yO11yFunnelOut GetTraceFunnel(ctx, funnelId).Execute()

Returns one funnel with its steps.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	funnelId := "funnelId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetTraceFunnel(context.Background(), funnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetTraceFunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTraceFunnel`: O11yO11yFunnelOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetTraceFunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**funnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceFunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yFunnelOut**](O11yO11yFunnelOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTraceFunnelErrorTraces

> O11yO11yFunnelRowsOut GetTraceFunnelErrorTraces(ctx, funnelId).O11yO11yFunnelStepWindowIn(o11yO11yFunnelStepWindowIn).Execute()

Returns the errored traces through a step transition of a saved funnel — the entry point for \"why is this step failing\".



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	funnelId := "funnelId_example" // string | 
	o11yO11yFunnelStepWindowIn := *openapiclient.NewO11yO11yFunnelStepWindowIn() // O11yO11yFunnelStepWindowIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetTraceFunnelErrorTraces(context.Background(), funnelId).O11yO11yFunnelStepWindowIn(o11yO11yFunnelStepWindowIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetTraceFunnelErrorTraces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTraceFunnelErrorTraces`: O11yO11yFunnelRowsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetTraceFunnelErrorTraces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**funnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceFunnelErrorTracesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yFunnelStepWindowIn** | [**O11yO11yFunnelStepWindowIn**](O11yO11yFunnelStepWindowIn.md) |  | 

### Return type

[**O11yO11yFunnelRowsOut**](O11yO11yFunnelRowsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTraceFunnelOverview

> O11yO11yFunnelRowsOut GetTraceFunnelOverview(ctx, funnelId).O11yO11yFunnelStepWindowIn(o11yO11yFunnelStepWindowIn).Execute()

Returns a saved funnel's conversion overview over a window: how many entered, how many converted, the rate and the latency.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	funnelId := "funnelId_example" // string | 
	o11yO11yFunnelStepWindowIn := *openapiclient.NewO11yO11yFunnelStepWindowIn() // O11yO11yFunnelStepWindowIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetTraceFunnelOverview(context.Background(), funnelId).O11yO11yFunnelStepWindowIn(o11yO11yFunnelStepWindowIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetTraceFunnelOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTraceFunnelOverview`: O11yO11yFunnelRowsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetTraceFunnelOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**funnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceFunnelOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yFunnelStepWindowIn** | [**O11yO11yFunnelStepWindowIn**](O11yO11yFunnelStepWindowIn.md) |  | 

### Return type

[**O11yO11yFunnelRowsOut**](O11yO11yFunnelRowsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTraceFunnelSlowTraces

> O11yO11yFunnelRowsOut GetTraceFunnelSlowTraces(ctx, funnelId).O11yO11yFunnelStepWindowIn(o11yO11yFunnelStepWindowIn).Execute()

Returns the slowest traces through a step transition of a saved funnel — the entry point for \"why is this step slow\".



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	funnelId := "funnelId_example" // string | 
	o11yO11yFunnelStepWindowIn := *openapiclient.NewO11yO11yFunnelStepWindowIn() // O11yO11yFunnelStepWindowIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetTraceFunnelSlowTraces(context.Background(), funnelId).O11yO11yFunnelStepWindowIn(o11yO11yFunnelStepWindowIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetTraceFunnelSlowTraces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTraceFunnelSlowTraces`: O11yO11yFunnelRowsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetTraceFunnelSlowTraces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**funnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceFunnelSlowTracesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yFunnelStepWindowIn** | [**O11yO11yFunnelStepWindowIn**](O11yO11yFunnelStepWindowIn.md) |  | 

### Return type

[**O11yO11yFunnelRowsOut**](O11yO11yFunnelRowsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTraceFunnelStepMetrics

> O11yO11yFunnelRowsOut GetTraceFunnelStepMetrics(ctx, funnelId).O11yO11yFunnelWindowIn(o11yO11yFunnelWindowIn).Execute()

Returns a saved funnel's per-step metrics over a window — the counts and latencies at each step, in step order.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	funnelId := "funnelId_example" // string | 
	o11yO11yFunnelWindowIn := *openapiclient.NewO11yO11yFunnelWindowIn() // O11yO11yFunnelWindowIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetTraceFunnelStepMetrics(context.Background(), funnelId).O11yO11yFunnelWindowIn(o11yO11yFunnelWindowIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetTraceFunnelStepMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTraceFunnelStepMetrics`: O11yO11yFunnelRowsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetTraceFunnelStepMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**funnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceFunnelStepMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yFunnelWindowIn** | [**O11yO11yFunnelWindowIn**](O11yO11yFunnelWindowIn.md) |  | 

### Return type

[**O11yO11yFunnelRowsOut**](O11yO11yFunnelRowsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTraceFunnelStepOverview

> O11yO11yFunnelRowsOut GetTraceFunnelStepOverview(ctx, funnelId).O11yO11yFunnelStepWindowIn(o11yO11yFunnelStepWindowIn).Execute()

Returns the conversion between two named steps of a saved funnel — the step-to-step drill-down behind the overview.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	funnelId := "funnelId_example" // string | 
	o11yO11yFunnelStepWindowIn := *openapiclient.NewO11yO11yFunnelStepWindowIn() // O11yO11yFunnelStepWindowIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetTraceFunnelStepOverview(context.Background(), funnelId).O11yO11yFunnelStepWindowIn(o11yO11yFunnelStepWindowIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetTraceFunnelStepOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTraceFunnelStepOverview`: O11yO11yFunnelRowsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetTraceFunnelStepOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**funnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceFunnelStepOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yFunnelStepWindowIn** | [**O11yO11yFunnelStepWindowIn**](O11yO11yFunnelStepWindowIn.md) |  | 

### Return type

[**O11yO11yFunnelRowsOut**](O11yO11yFunnelRowsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> O11yO11yUserWithRolesOut GetUser(ctx, id).Execute()

Returns one org member together with every role they hold, by user id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetUser(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: O11yO11yUserWithRolesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yUserWithRolesOut**](O11yO11yUserWithRolesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserDeprecated

> O11yO11yDeprecatedUserOut GetUserDeprecated(ctx, id).Execute()

Returns one org member with their single legacy role, by user id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetUserDeprecated(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetUserDeprecated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserDeprecated`: O11yO11yDeprecatedUserOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetUserDeprecated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yDeprecatedUserOut**](O11yO11yDeprecatedUserOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserPreference

> O11yO11yPreferenceOut GetUserPreference(ctx, name).Execute()

Returns one preference of the calling user, by name.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetUserPreference(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetUserPreference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserPreference`: O11yO11yPreferenceOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetUserPreference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPreferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yPreferenceOut**](O11yO11yPreferenceOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersByRoleID

> O11yO11yUsersOut GetUsersByRoleID(ctx, id).Execute()

Returns every org member holding a role, by role id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetUsersByRoleID(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetUsersByRoleID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersByRoleID`: O11yO11yUsersOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetUsersByRoleID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersByRoleIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yUsersOut**](O11yO11yUsersOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWaterfallV4

> O11yO11yTraceWaterfallOut GetWaterfallV4(ctx, traceId).O11yO11yTraceWaterfallIn(o11yO11yTraceWaterfallIn).Execute()

Returns a trace's waterfall: every span when the trace is small enough, a capped window around the selected span when it is not, with the uncollapsed subtrees the caller asked to keep open.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	traceId := "traceId_example" // string | 
	o11yO11yTraceWaterfallIn := *openapiclient.NewO11yO11yTraceWaterfallIn() // O11yO11yTraceWaterfallIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.GetWaterfallV4(context.Background(), traceId).O11yO11yTraceWaterfallIn(o11yO11yTraceWaterfallIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.GetWaterfallV4``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWaterfallV4`: O11yO11yTraceWaterfallOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.GetWaterfallV4`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWaterfallV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yTraceWaterfallIn** | [**O11yO11yTraceWaterfallIn**](O11yO11yTraceWaterfallIn.md) |  | 

### Return type

[**O11yO11yTraceWaterfallOut**](O11yO11yTraceWaterfallOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InspectMetrics

> O11yO11yMetricInspectOut InspectMetrics(ctx).O11yO11yMetricInspectIn(o11yO11yMetricInspectIn).Execute()

Returns one metric's raw time series over a window of at most thirty minutes — each series with its labels and timestamp/value pairs.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yMetricInspectIn := *openapiclient.NewO11yO11yMetricInspectIn(int32(123), "MetricName_example", int32(123)) // O11yO11yMetricInspectIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.InspectMetrics(context.Background()).O11yO11yMetricInspectIn(o11yO11yMetricInspectIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.InspectMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InspectMetrics`: O11yO11yMetricInspectOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.InspectMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInspectMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yMetricInspectIn** | [**O11yO11yMetricInspectIn**](O11yO11yMetricInspectIn.md) |  | 

### Return type

[**O11yO11yMetricInspectOut**](O11yO11yMetricInspectOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallIntegration

> O11yO11yInstallOut InstallIntegration(ctx).O11yInstallIntegrationRequest(o11yInstallIntegrationRequest).Execute()

Installs an integration into the caller's org from its id and configuration, answering with the installed catalog item.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yInstallIntegrationRequest := *openapiclient.NewO11yInstallIntegrationRequest() // O11yInstallIntegrationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.InstallIntegration(context.Background()).O11yInstallIntegrationRequest(o11yInstallIntegrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.InstallIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstallIntegration`: O11yO11yInstallOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.InstallIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstallIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yInstallIntegrationRequest** | [**O11yInstallIntegrationRequest**](O11yInstallIntegrationRequest.md) |  | 

### Return type

[**O11yO11yInstallOut**](O11yO11yInstallOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountServicesMetadata

> O11yO11yServicesMetadataOut ListAccountServicesMetadata(ctx, cloudProvider, id).Execute()

Lists the services metadata for one connected account of the given provider, by account id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	cloudProvider := "cloudProvider_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListAccountServicesMetadata(context.Background(), cloudProvider, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListAccountServicesMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccountServicesMetadata`: O11yO11yServicesMetadataOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListAccountServicesMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountServicesMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**O11yO11yServicesMetadataOut**](O11yO11yServicesMetadataOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccounts

> O11yO11yAccountsOut ListAccounts(ctx, cloudProvider).Execute()

Lists the cloud-integration accounts connected for the given provider.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	cloudProvider := "cloudProvider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListAccounts(context.Background(), cloudProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccounts`: O11yO11yAccountsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yAccountsOut**](O11yO11yAccountsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthDomains

> O11yO11yAuthDomainsOut ListAuthDomains(ctx).Execute()

Lists the org's auth domains — the email domains whose SSO configuration this org owns.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListAuthDomains(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListAuthDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAuthDomains`: O11yO11yAuthDomainsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListAuthDomains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAuthDomainsRequest struct via the builder pattern


### Return type

[**O11yO11yAuthDomainsOut**](O11yO11yAuthDomainsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChannels

> O11yO11yChannelsOut ListChannels(ctx).Execute()

Lists the org's notification channels.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListChannels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListChannels`: O11yO11yChannelsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListChannels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListChannelsRequest struct via the builder pattern


### Return type

[**O11yO11yChannelsOut**](O11yO11yChannelsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDashboardViews

> O11yO11yDashboardViewListOut ListDashboardViews(ctx).Execute()

Returns every saved view in the calling user's org.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListDashboardViews(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListDashboardViews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDashboardViews`: O11yO11yDashboardViewListOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListDashboardViews`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDashboardViewsRequest struct via the builder pattern


### Return type

[**O11yO11yDashboardViewListOut**](O11yO11yDashboardViewListOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDashboardsForUserV2

> O11yO11yDashboardListForUserOut ListDashboardsForUserV2(ctx).Query(query).Sort(sort).Order(order).Limit(limit).Offset(offset).Execute()

Is dashboardListV2 personalized for the calling user: each dashboard carries the caller's pinned state, and pinned dashboards float to the top of the requested ordering.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	query := "query_example" // string | Query is the filter DSL over dashboard columns and tags, e.g. `name:cpu source:user`. Empty lists everything. (optional)
	sort := "sort_example" // string | Sort is the sort field: updated_at, created_at or name. Empty sorts by updated_at. (optional)
	order := "order_example" // string | Order is the sort direction: asc or desc. Empty orders desc. (optional)
	limit := int32(56) // int32 | Limit caps how many dashboards come back. Zero means the default of 20; the runtime caps it at 200. (optional)
	offset := int32(56) // int32 | Offset is how many dashboards to skip for pagination. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListDashboardsForUserV2(context.Background()).Query(query).Sort(sort).Order(order).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListDashboardsForUserV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDashboardsForUserV2`: O11yO11yDashboardListForUserOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListDashboardsForUserV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDashboardsForUserV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query is the filter DSL over dashboard columns and tags, e.g. &#x60;name:cpu source:user&#x60;. Empty lists everything. | 
 **sort** | **string** | Sort is the sort field: updated_at, created_at or name. Empty sorts by updated_at. | 
 **order** | **string** | Order is the sort direction: asc or desc. Empty orders desc. | 
 **limit** | **int32** | Limit caps how many dashboards come back. Zero means the default of 20; the runtime caps it at 200. | 
 **offset** | **int32** | Offset is how many dashboards to skip for pagination. | 

### Return type

[**O11yO11yDashboardListForUserOut**](O11yO11yDashboardListForUserOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDashboardsV2

> O11yO11yDashboardListOut ListDashboardsV2(ctx).Query(query).Sort(sort).Order(order).Limit(limit).Offset(offset).Execute()

Returns a page of v2-shape dashboards for the org.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	query := "query_example" // string | Query is the filter DSL over dashboard columns and tags, e.g. `name:cpu source:user`. Empty lists everything. (optional)
	sort := "sort_example" // string | Sort is the sort field: updated_at, created_at or name. Empty sorts by updated_at. (optional)
	order := "order_example" // string | Order is the sort direction: asc or desc. Empty orders desc. (optional)
	limit := int32(56) // int32 | Limit caps how many dashboards come back. Zero means the default of 20; the runtime caps it at 200. (optional)
	offset := int32(56) // int32 | Offset is how many dashboards to skip for pagination. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListDashboardsV2(context.Background()).Query(query).Sort(sort).Order(order).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListDashboardsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDashboardsV2`: O11yO11yDashboardListOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListDashboardsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDashboardsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query is the filter DSL over dashboard columns and tags, e.g. &#x60;name:cpu source:user&#x60;. Empty lists everything. | 
 **sort** | **string** | Sort is the sort field: updated_at, created_at or name. Empty sorts by updated_at. | 
 **order** | **string** | Order is the sort direction: asc or desc. Empty orders desc. | 
 **limit** | **int32** | Limit caps how many dashboards come back. Zero means the default of 20; the runtime caps it at 200. | 
 **offset** | **int32** | Offset is how many dashboards to skip for pagination. | 

### Return type

[**O11yO11yDashboardListOut**](O11yO11yDashboardListOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDowntimeSchedules

> O11yO11yDowntimeSchedulesOut ListDowntimeSchedules(ctx).Active(active).Recurring(recurring).Execute()

Lists all planned maintenance windows, optionally narrowed to the active ones or the recurring ones.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	active := "active_example" // string | Active, when \"true\" or \"false\", keeps only the active or inactive windows. Absent lists all. (optional)
	recurring := "recurring_example" // string | Recurring, when \"true\" or \"false\", keeps only the recurring or one-off windows. Absent lists all. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListDowntimeSchedules(context.Background()).Active(active).Recurring(recurring).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListDowntimeSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDowntimeSchedules`: O11yO11yDowntimeSchedulesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListDowntimeSchedules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDowntimeSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **active** | **string** | Active, when \&quot;true\&quot; or \&quot;false\&quot;, keeps only the active or inactive windows. Absent lists all. | 
 **recurring** | **string** | Recurring, when \&quot;true\&quot; or \&quot;false\&quot;, keeps only the recurring or one-off windows. Absent lists all. | 

### Return type

[**O11yO11yDowntimeSchedulesOut**](O11yO11yDowntimeSchedulesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntegrations

> O11yO11yIntegrationsListOut ListIntegrations(ctx).IsInstalled(isInstalled).Execute()

Lists the available integrations and whether each is installed in the caller's org, optionally narrowed to installed or not-installed.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	isInstalled := "isInstalled_example" // string | IsInstalled, when \"true\" or \"false\", keeps only integrations in that installed state; empty lists them all. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListIntegrations(context.Background()).IsInstalled(isInstalled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIntegrations`: O11yO11yIntegrationsListOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListIntegrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isInstalled** | **string** | IsInstalled, when \&quot;true\&quot; or \&quot;false\&quot;, keeps only integrations in that installed state; empty lists them all. | 

### Return type

[**O11yO11yIntegrationsListOut**](O11yO11yIntegrationsListOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLLMAnnotations

> O11yO11yLLMAnnotationsOut ListLLMAnnotations(ctx).TraceId(traceId).Queue(queue).Status(status).Offset(offset).Limit(limit).Execute()

Lists human annotations on traces and observations, optionally scoped to one review queue.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	traceId := "traceId_example" // string | TraceID narrows to annotations on one trace. (optional)
	queue := "queue_example" // string | Queue narrows to one review queue. (optional)
	status := "status_example" // string | Status narrows to one review status, e.g. PENDING. (optional)
	offset := int32(56) // int32 | Offset is how many rows to skip, for paging. (optional)
	limit := int32(56) // int32 | Limit caps how many rows come back. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListLLMAnnotations(context.Background()).TraceId(traceId).Queue(queue).Status(status).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListLLMAnnotations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLLMAnnotations`: O11yO11yLLMAnnotationsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListLLMAnnotations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLLMAnnotationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **traceId** | **string** | TraceID narrows to annotations on one trace. | 
 **queue** | **string** | Queue narrows to one review queue. | 
 **status** | **string** | Status narrows to one review status, e.g. PENDING. | 
 **offset** | **int32** | Offset is how many rows to skip, for paging. | 
 **limit** | **int32** | Limit caps how many rows come back. | 

### Return type

[**O11yO11yLLMAnnotationsOut**](O11yO11yLLMAnnotationsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLLMObservations

> O11yO11yLLMObservationsOut ListLLMObservations(ctx).Start(start).End(end).TraceId(traceId).SessionId(sessionId).UserId(userId).Name(name).Model(model).Offset(offset).Limit(limit).Execute()

Lists gen_ai spans as LLM observations — each an LLM call with its model, token counts, cost and latency projected from gen_ai.* attributes, newest first, over the query window.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	start := int32(56) // int32 | Start is the start of the window as a unix-millisecond epoch. Zero means 24h before the end. (optional)
	end := int32(56) // int32 | End is the end of the window as a unix-millisecond epoch. Zero means now. (optional)
	traceId := "traceId_example" // string | TraceID narrows the view to one trace. (optional)
	sessionId := "sessionId_example" // string | SessionID narrows the view to one conversation. (optional)
	userId := "userId_example" // string | UserID narrows the view to one end user. (optional)
	name := "name_example" // string | Name narrows the view to observations of one name. (optional)
	model := "model_example" // string | Model narrows the view to one model. (optional)
	offset := int32(56) // int32 | Offset is how many rows to skip, for paging. (optional)
	limit := int32(56) // int32 | Limit caps how many rows come back. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListLLMObservations(context.Background()).Start(start).End(end).TraceId(traceId).SessionId(sessionId).UserId(userId).Name(name).Model(model).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListLLMObservations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLLMObservations`: O11yO11yLLMObservationsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListLLMObservations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLLMObservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | Start is the start of the window as a unix-millisecond epoch. Zero means 24h before the end. | 
 **end** | **int32** | End is the end of the window as a unix-millisecond epoch. Zero means now. | 
 **traceId** | **string** | TraceID narrows the view to one trace. | 
 **sessionId** | **string** | SessionID narrows the view to one conversation. | 
 **userId** | **string** | UserID narrows the view to one end user. | 
 **name** | **string** | Name narrows the view to observations of one name. | 
 **model** | **string** | Model narrows the view to one model. | 
 **offset** | **int32** | Offset is how many rows to skip, for paging. | 
 **limit** | **int32** | Limit caps how many rows come back. | 

### Return type

[**O11yO11yLLMObservationsOut**](O11yO11yLLMObservationsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLLMPricingRules

> O11yO11yLLMPricingRulesOut ListLLMPricingRules(ctx).Q(q).IsOverride(isOverride).Offset(offset).Limit(limit).Execute()

Returns the LLM pricing rules for the caller's org, with pagination and an optional search and override filter.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	q := "q_example" // string | Search matches rules by model or provider. (optional)
	isOverride := "isOverride_example" // string | IsOverride, when \"true\" or \"false\", narrows to user-pinned rules or to synced ones; empty returns both. It is a string because a query param is a string on the wire, and the runtime reads absent as \"no filter\". (optional)
	offset := int32(56) // int32 | Offset is how many rows to skip, for paging. (optional)
	limit := int32(56) // int32 | Limit caps how many rows come back. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListLLMPricingRules(context.Background()).Q(q).IsOverride(isOverride).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListLLMPricingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLLMPricingRules`: O11yO11yLLMPricingRulesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListLLMPricingRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLLMPricingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Search matches rules by model or provider. | 
 **isOverride** | **string** | IsOverride, when \&quot;true\&quot; or \&quot;false\&quot;, narrows to user-pinned rules or to synced ones; empty returns both. It is a string because a query param is a string on the wire, and the runtime reads absent as \&quot;no filter\&quot;. | 
 **offset** | **int32** | Offset is how many rows to skip, for paging. | 
 **limit** | **int32** | Limit caps how many rows come back. | 

### Return type

[**O11yO11yLLMPricingRulesOut**](O11yO11yLLMPricingRulesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLLMScores

> O11yO11yLLMScoresOut ListLLMScores(ctx).TraceId(traceId).ObservationId(observationId).Name(name).Source(source).Offset(offset).Limit(limit).Execute()

Lists eval scores and human-feedback signals attached to traces and observations, newest first.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	traceId := "traceId_example" // string | TraceID narrows to scores on one trace. (optional)
	observationId := "observationId_example" // string | ObservationID narrows to scores on one observation. (optional)
	name := "name_example" // string | Name narrows to scores of one name. (optional)
	source := "source_example" // string | Source narrows to scores from one source, e.g. API, EVAL. (optional)
	offset := int32(56) // int32 | Offset is how many rows to skip, for paging. (optional)
	limit := int32(56) // int32 | Limit caps how many rows come back. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListLLMScores(context.Background()).TraceId(traceId).ObservationId(observationId).Name(name).Source(source).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListLLMScores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLLMScores`: O11yO11yLLMScoresOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListLLMScores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLLMScoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **traceId** | **string** | TraceID narrows to scores on one trace. | 
 **observationId** | **string** | ObservationID narrows to scores on one observation. | 
 **name** | **string** | Name narrows to scores of one name. | 
 **source** | **string** | Source narrows to scores from one source, e.g. API, EVAL. | 
 **offset** | **int32** | Offset is how many rows to skip, for paging. | 
 **limit** | **int32** | Limit caps how many rows come back. | 

### Return type

[**O11yO11yLLMScoresOut**](O11yO11yLLMScoresOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLLMSessions

> O11yO11yLLMSessionsOut ListLLMSessions(ctx).Start(start).End(end).TraceId(traceId).SessionId(sessionId).UserId(userId).Name(name).Model(model).Offset(offset).Limit(limit).Execute()

Lists conversations — gen_ai spans grouped by session.id, with their trace and observation counts, tokens and cost.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	start := int32(56) // int32 | Start is the start of the window as a unix-millisecond epoch. Zero means 24h before the end. (optional)
	end := int32(56) // int32 | End is the end of the window as a unix-millisecond epoch. Zero means now. (optional)
	traceId := "traceId_example" // string | TraceID narrows the view to one trace. (optional)
	sessionId := "sessionId_example" // string | SessionID narrows the view to one conversation. (optional)
	userId := "userId_example" // string | UserID narrows the view to one end user. (optional)
	name := "name_example" // string | Name narrows the view to observations of one name. (optional)
	model := "model_example" // string | Model narrows the view to one model. (optional)
	offset := int32(56) // int32 | Offset is how many rows to skip, for paging. (optional)
	limit := int32(56) // int32 | Limit caps how many rows come back. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListLLMSessions(context.Background()).Start(start).End(end).TraceId(traceId).SessionId(sessionId).UserId(userId).Name(name).Model(model).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListLLMSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLLMSessions`: O11yO11yLLMSessionsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListLLMSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLLMSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | Start is the start of the window as a unix-millisecond epoch. Zero means 24h before the end. | 
 **end** | **int32** | End is the end of the window as a unix-millisecond epoch. Zero means now. | 
 **traceId** | **string** | TraceID narrows the view to one trace. | 
 **sessionId** | **string** | SessionID narrows the view to one conversation. | 
 **userId** | **string** | UserID narrows the view to one end user. | 
 **name** | **string** | Name narrows the view to observations of one name. | 
 **model** | **string** | Model narrows the view to one model. | 
 **offset** | **int32** | Offset is how many rows to skip, for paging. | 
 **limit** | **int32** | Limit caps how many rows come back. | 

### Return type

[**O11yO11yLLMSessionsOut**](O11yO11yLLMSessionsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLLMTraces

> O11yO11yLLMTracesOut ListLLMTraces(ctx).Start(start).End(end).TraceId(traceId).SessionId(sessionId).UserId(userId).Name(name).Model(model).Offset(offset).Limit(limit).Execute()

Lists LLM traces — gen_ai spans grouped by trace_id, with cost, tokens and latency rolled up across each trace.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	start := int32(56) // int32 | Start is the start of the window as a unix-millisecond epoch. Zero means 24h before the end. (optional)
	end := int32(56) // int32 | End is the end of the window as a unix-millisecond epoch. Zero means now. (optional)
	traceId := "traceId_example" // string | TraceID narrows the view to one trace. (optional)
	sessionId := "sessionId_example" // string | SessionID narrows the view to one conversation. (optional)
	userId := "userId_example" // string | UserID narrows the view to one end user. (optional)
	name := "name_example" // string | Name narrows the view to observations of one name. (optional)
	model := "model_example" // string | Model narrows the view to one model. (optional)
	offset := int32(56) // int32 | Offset is how many rows to skip, for paging. (optional)
	limit := int32(56) // int32 | Limit caps how many rows come back. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListLLMTraces(context.Background()).Start(start).End(end).TraceId(traceId).SessionId(sessionId).UserId(userId).Name(name).Model(model).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListLLMTraces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLLMTraces`: O11yO11yLLMTracesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListLLMTraces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLLMTracesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | Start is the start of the window as a unix-millisecond epoch. Zero means 24h before the end. | 
 **end** | **int32** | End is the end of the window as a unix-millisecond epoch. Zero means now. | 
 **traceId** | **string** | TraceID narrows the view to one trace. | 
 **sessionId** | **string** | SessionID narrows the view to one conversation. | 
 **userId** | **string** | UserID narrows the view to one end user. | 
 **name** | **string** | Name narrows the view to observations of one name. | 
 **model** | **string** | Model narrows the view to one model. | 
 **offset** | **int32** | Offset is how many rows to skip, for paging. | 
 **limit** | **int32** | Limit caps how many rows come back. | 

### Return type

[**O11yO11yLLMTracesOut**](O11yO11yLLMTracesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLLMUsers

> O11yO11yLLMUsersOut ListLLMUsers(ctx).Start(start).End(end).TraceId(traceId).SessionId(sessionId).UserId(userId).Name(name).Model(model).Offset(offset).Limit(limit).Execute()

Lists end users — gen_ai spans grouped by user.id, with their session, trace and observation counts, tokens and cost.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	start := int32(56) // int32 | Start is the start of the window as a unix-millisecond epoch. Zero means 24h before the end. (optional)
	end := int32(56) // int32 | End is the end of the window as a unix-millisecond epoch. Zero means now. (optional)
	traceId := "traceId_example" // string | TraceID narrows the view to one trace. (optional)
	sessionId := "sessionId_example" // string | SessionID narrows the view to one conversation. (optional)
	userId := "userId_example" // string | UserID narrows the view to one end user. (optional)
	name := "name_example" // string | Name narrows the view to observations of one name. (optional)
	model := "model_example" // string | Model narrows the view to one model. (optional)
	offset := int32(56) // int32 | Offset is how many rows to skip, for paging. (optional)
	limit := int32(56) // int32 | Limit caps how many rows come back. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListLLMUsers(context.Background()).Start(start).End(end).TraceId(traceId).SessionId(sessionId).UserId(userId).Name(name).Model(model).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListLLMUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLLMUsers`: O11yO11yLLMUsersOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListLLMUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLLMUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | Start is the start of the window as a unix-millisecond epoch. Zero means 24h before the end. | 
 **end** | **int32** | End is the end of the window as a unix-millisecond epoch. Zero means now. | 
 **traceId** | **string** | TraceID narrows the view to one trace. | 
 **sessionId** | **string** | SessionID narrows the view to one conversation. | 
 **userId** | **string** | UserID narrows the view to one end user. | 
 **name** | **string** | Name narrows the view to observations of one name. | 
 **model** | **string** | Model narrows the view to one model. | 
 **offset** | **int32** | Offset is how many rows to skip, for paging. | 
 **limit** | **int32** | Limit caps how many rows come back. | 

### Return type

[**O11yO11yLLMUsersOut**](O11yO11yLLMUsersOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetricReductionRules

> O11yO11yReductionRuleListOut ListMetricReductionRules(ctx).OrderBy(orderBy).Order(order).Search(search).MetricName(metricName).Offset(offset).Limit(limit).Execute()

Lists the org's metric volume-control (label reduction) rules, pageable and sortable by name, volume or recency.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	orderBy := "orderBy_example" // string | OrderBy sorts the page: metric, ingested_volume, reduced_volume or last_updated. Unset means ingested_volume. (optional)
	order := "order_example" // string | Order is asc or desc. Unset means desc. (optional)
	search := "search_example" // string | Search narrows the page to rules whose metric name contains it. (optional)
	metricName := "metricName_example" // string | MetricName narrows the page to one metric's rule. (optional)
	offset := int32(56) // int32 | Offset is how many rules to skip, for paging. (optional)
	limit := int32(56) // int32 | Limit caps how many rules come back, at most 1000. Unset means 10. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListMetricReductionRules(context.Background()).OrderBy(orderBy).Order(order).Search(search).MetricName(metricName).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListMetricReductionRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMetricReductionRules`: O11yO11yReductionRuleListOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListMetricReductionRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMetricReductionRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderBy** | **string** | OrderBy sorts the page: metric, ingested_volume, reduced_volume or last_updated. Unset means ingested_volume. | 
 **order** | **string** | Order is asc or desc. Unset means desc. | 
 **search** | **string** | Search narrows the page to rules whose metric name contains it. | 
 **metricName** | **string** | MetricName narrows the page to one metric&#39;s rule. | 
 **offset** | **int32** | Offset is how many rules to skip, for paging. | 
 **limit** | **int32** | Limit caps how many rules come back, at most 1000. Unset means 10. | 

### Return type

[**O11yO11yReductionRuleListOut**](O11yO11yReductionRuleListOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetrics

> O11yO11yMetricListOut ListMetrics(ctx).Start(start).End(end).Limit(limit).SearchText(searchText).Source(source).Execute()

Lists the distinct metric names seen in a time range, each with its description, type, unit, temporality and monotonicity.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	start := int32(56) // int32 | Start is the start of the window as a Unix timestamp in milliseconds. (optional)
	end := int32(56) // int32 | End is the end of the window as a Unix timestamp in milliseconds. (optional)
	limit := int32(56) // int32 | Limit caps how many metrics come back; unset means 100, at most 5000. (optional)
	searchText := "searchText_example" // string | SearchText narrows the page to metric names containing it. (optional)
	source := "source_example" // string | Source narrows the page by ingestion source. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListMetrics(context.Background()).Start(start).End(end).Limit(limit).SearchText(searchText).Source(source).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMetrics`: O11yO11yMetricListOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | Start is the start of the window as a Unix timestamp in milliseconds. | 
 **end** | **int32** | End is the end of the window as a Unix timestamp in milliseconds. | 
 **limit** | **int32** | Limit caps how many metrics come back; unset means 100, at most 5000. | 
 **searchText** | **string** | SearchText narrows the page to metric names containing it. | 
 **source** | **string** | Source narrows the page by ingestion source. | 

### Return type

[**O11yO11yMetricListOut**](O11yO11yMetricListOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgPreferences

> O11yO11yPreferencesOut ListOrgPreferences(ctx).Execute()

Lists every org-scoped preference, each with its current and default value.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListOrgPreferences(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListOrgPreferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgPreferences`: O11yO11yPreferencesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListOrgPreferences`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgPreferencesRequest struct via the builder pattern


### Return type

[**O11yO11yPreferencesOut**](O11yO11yPreferencesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoles

> O11yO11yRolesOut ListRoles(ctx).Execute()

Lists every role in the caller's org — the managed ones the platform seeds and the custom ones its admins created.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListRoles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRoles`: O11yO11yRolesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListRoles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRolesRequest struct via the builder pattern


### Return type

[**O11yO11yRolesOut**](O11yO11yRolesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRules

> O11yO11yRulesOut ListRules(ctx).Execute()

Lists all alert rules with their current evaluation state.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListRules(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRules`: O11yO11yRulesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRulesRequest struct via the builder pattern


### Return type

[**O11yO11yRulesOut**](O11yO11yRulesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceAccountKeys

> O11yO11yAPIKeysOut ListServiceAccountKeys(ctx, id).Execute()

Lists a service account's API keys — metadata only, never the secrets.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListServiceAccountKeys(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListServiceAccountKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServiceAccountKeys`: O11yO11yAPIKeysOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListServiceAccountKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceAccountKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11yAPIKeysOut**](O11yO11yAPIKeysOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceAccounts

> O11yO11yServiceAccountsOut ListServiceAccounts(ctx).Execute()

Lists the caller's org's service accounts.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListServiceAccounts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListServiceAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServiceAccounts`: O11yO11yServiceAccountsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListServiceAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceAccountsRequest struct via the builder pattern


### Return type

[**O11yO11yServiceAccountsOut**](O11yO11yServiceAccountsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServicesMetadata

> O11yO11yServicesMetadataOut ListServicesMetadata(ctx, cloudProvider).CloudIntegrationId(cloudIntegrationId).Execute()

Lists the services the given provider can collect from, optionally scoped to one cloud integration.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	cloudProvider := "cloudProvider_example" // string | 
	cloudIntegrationId := "cloudIntegrationId_example" // string | CloudIntegrationID, when set, scopes the listing to one cloud integration. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListServicesMetadata(context.Background(), cloudProvider).CloudIntegrationId(cloudIntegrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListServicesMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServicesMetadata`: O11yO11yServicesMetadataOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListServicesMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServicesMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudIntegrationId** | **string** | CloudIntegrationID, when set, scopes the listing to one cloud integration. | 

### Return type

[**O11yO11yServicesMetadataOut**](O11yO11yServicesMetadataOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSpanMapperGroups

> O11yO11ySpanMapperGroupsOut ListSpanMapperGroups(ctx).Enabled(enabled).Execute()

Lists the caller's org's mapping groups, optionally only the enabled ones.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	enabled := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListSpanMapperGroups(context.Background()).Enabled(enabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListSpanMapperGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSpanMapperGroups`: O11yO11ySpanMapperGroupsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListSpanMapperGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSpanMapperGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enabled** | **bool** |  | 

### Return type

[**O11yO11ySpanMapperGroupsOut**](O11yO11ySpanMapperGroupsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSpanMappers

> O11yO11ySpanMappersOut ListSpanMappers(ctx, groupId).Execute()

Lists the mappers belonging to one group, in the order they are applied.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	groupId := "groupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListSpanMappers(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListSpanMappers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSpanMappers`: O11yO11ySpanMappersOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListSpanMappers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSpanMappersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11ySpanMappersOut**](O11yO11ySpanMappersOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTraceFunnels

> O11yO11yFunnelsOut ListTraceFunnels(ctx).Execute()

Lists the caller's org's funnels, each with its steps and who last touched it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListTraceFunnels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListTraceFunnels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTraceFunnels`: O11yO11yFunnelsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListTraceFunnels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTraceFunnelsRequest struct via the builder pattern


### Return type

[**O11yO11yFunnelsOut**](O11yO11yFunnelsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserPreferences

> O11yO11yPreferencesOut ListUserPreferences(ctx).Execute()

Lists every preference of the calling user, each with its current and default value.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListUserPreferences(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListUserPreferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserPreferences`: O11yO11yPreferencesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListUserPreferences`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListUserPreferencesRequest struct via the builder pattern


### Return type

[**O11yO11yPreferencesOut**](O11yO11yPreferencesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> O11yO11yUsersOut ListUsers(ctx).Execute()

Lists the caller's org members.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListUsers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsers`: O11yO11yUsersOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


### Return type

[**O11yO11yUsersOut**](O11yO11yUsersOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsersDeprecated

> O11yO11yDeprecatedUsersOut ListUsersDeprecated(ctx).Execute()

Lists the org's members with their single legacy role.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ListUsersDeprecated(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ListUsersDeprecated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsersDeprecated`: O11yO11yDeprecatedUsersOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ListUsersDeprecated`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListUsersDeprecatedRequest struct via the builder pattern


### Return type

[**O11yO11yDeprecatedUsersOut**](O11yO11yDeprecatedUsersOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LockDashboardV2

> LockDashboardV2(ctx, id).Execute()

Locks a v2-shape dashboard.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the resource id from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.LockDashboardV2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.LockDashboardV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the resource id from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockDashboardV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDashboardV2

> O11yO11yDashboardOut PatchDashboardV2(ctx, id).O11yO11yDashboardPatchIn(o11yO11yDashboardPatchIn).Execute()

Applies an RFC 6902 JSON Patch to a v2-shape dashboard.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the dashboard id from the path.
	o11yO11yDashboardPatchIn := *openapiclient.NewO11yO11yDashboardPatchIn() // O11yO11yDashboardPatchIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PatchDashboardV2(context.Background(), id).O11yO11yDashboardPatchIn(o11yO11yDashboardPatchIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PatchDashboardV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchDashboardV2`: O11yO11yDashboardOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PatchDashboardV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the dashboard id from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDashboardV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yDashboardPatchIn** | [**O11yO11yDashboardPatchIn**](O11yO11yDashboardPatchIn.md) |  | 

### Return type

[**O11yO11yDashboardOut**](O11yO11yDashboardOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchO11yReviewsById

> O11yAnnQueueView PatchO11yReviewsById(ctx, id).O11yUpdateQueueIn(o11yUpdateQueueIn).Execute()

Changes a review queue's name, description or score-config set.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "annq_1" // string | ID is the annotation queue to update, from the path.
	o11yUpdateQueueIn := *openapiclient.NewO11yUpdateQueueIn() // O11yUpdateQueueIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PatchO11yReviewsById(context.Background(), id).O11yUpdateQueueIn(o11yUpdateQueueIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PatchO11yReviewsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchO11yReviewsById`: O11yAnnQueueView
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PatchO11yReviewsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the annotation queue to update, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchO11yReviewsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yUpdateQueueIn** | [**O11yUpdateQueueIn**](O11yUpdateQueueIn.md) |  | 

### Return type

[**O11yAnnQueueView**](O11yAnnQueueView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchO11yReviewsByIdItemsByItemid

> O11yAnnItemView PatchO11yReviewsByIdItemsByItemid(ctx, id, itemId).O11yUpdateItemIn(o11yUpdateItemIn).Execute()

Moves one queue item between PENDING and COMPLETED and sets its assignee.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "annq_1" // string | ID is the annotation queue the item belongs to, from the path.
	itemId := "annqi_1" // string | ItemID is the item to update, from the path.
	o11yUpdateItemIn := *openapiclient.NewO11yUpdateItemIn() // O11yUpdateItemIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PatchO11yReviewsByIdItemsByItemid(context.Background(), id, itemId).O11yUpdateItemIn(o11yUpdateItemIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PatchO11yReviewsByIdItemsByItemid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchO11yReviewsByIdItemsByItemid`: O11yAnnItemView
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PatchO11yReviewsByIdItemsByItemid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the annotation queue the item belongs to, from the path. | 
**itemId** | **string** | ItemID is the item to update, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchO11yReviewsByIdItemsByItemidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **o11yUpdateItemIn** | [**O11yUpdateItemIn**](O11yUpdateItemIn.md) |  | 

### Return type

[**O11yAnnItemView**](O11yAnnItemView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRuleByID

> O11yO11yRuleOut PatchRuleByID(ctx, id).Body(body).Execute()

Applies a partial update to an alert rule, by id, answering with the stored rule — the common toggle for enabling or muting a rule.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 
	body := interface{}(987) // interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PatchRuleByID(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PatchRuleByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchRuleByID`: O11yO11yRuleOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PatchRuleByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRuleByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** |  | 

### Return type

[**O11yO11yRuleOut**](O11yO11yRuleOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PinDashboardV2

> PinDashboardV2(ctx, id).Execute()

Pins a dashboard for the calling user.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the resource id from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.PinDashboardV2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PinDashboardV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the resource id from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPinDashboardV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yAlertsByReceiver

> PostO11yAlertsByReceiver(ctx, receiver).Execute()

Take an Alertmanager notification and page a human



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	receiver := "receiver_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.PostO11yAlertsByReceiver(context.Background(), receiver).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yAlertsByReceiver``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**receiver** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yAlertsByReceiverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yApiByProjectIdEnvelope

> PostO11yApiByProjectIdEnvelope(ctx, projectId).Execute()

Receive a Sentry envelope on the SDK's own DSN path



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	projectId := "projectId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.PostO11yApiByProjectIdEnvelope(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yApiByProjectIdEnvelope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yApiByProjectIdEnvelopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yApiByProjectIdStore

> PostO11yApiByProjectIdStore(ctx, projectId).Execute()

Receive a single Sentry event on the SDK's own DSN path



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	projectId := "projectId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.PostO11yApiByProjectIdStore(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yApiByProjectIdStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yApiByProjectIdStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yAutoCompleteAttributeValues

> O11yO11yAttributeValuesOut PostO11yAutoCompleteAttributeValues(ctx).O11yFilterAttributeValueRequest(o11yFilterAttributeValueRequest).Execute()

Reads the attribute-value request from the body rather than off the query string — the spelling the newer builder uses to send its filters alongside the request.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yFilterAttributeValueRequest := *openapiclient.NewO11yFilterAttributeValueRequest() // O11yFilterAttributeValueRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yAutoCompleteAttributeValues(context.Background()).O11yFilterAttributeValueRequest(o11yFilterAttributeValueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yAutoCompleteAttributeValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yAutoCompleteAttributeValues`: O11yO11yAttributeValuesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yAutoCompleteAttributeValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yAutoCompleteAttributeValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yFilterAttributeValueRequest** | [**O11yFilterAttributeValueRequest**](O11yFilterAttributeValueRequest.md) |  | 

### Return type

[**O11yO11yAttributeValuesOut**](O11yO11yAttributeValuesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yClustersList

> O11yO11yClusterListOut PostO11yClustersList(ctx).O11yClusterListRequest(o11yClusterListRequest).Execute()

Lists Kubernetes clusters over a time range, each with its CPU and memory usage against allocatable capacity and its attributes; filterable, groupable and paginated.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yClusterListRequest := *openapiclient.NewO11yClusterListRequest() // O11yClusterListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yClustersList(context.Background()).O11yClusterListRequest(o11yClusterListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yClustersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yClustersList`: O11yO11yClusterListOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yClustersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yClustersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yClusterListRequest** | [**O11yClusterListRequest**](O11yClusterListRequest.md) |  | 

### Return type

[**O11yO11yClusterListOut**](O11yO11yClusterListOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yCompleteSaml

> PostO11yCompleteSaml(ctx).Execute()

Complete a SAML sign-in



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.PostO11yCompleteSaml(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yCompleteSaml``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yCompleteSamlRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yCounterrors

> int32 PostO11yCounterrors(ctx).O11yO11yErrorsCountIn(o11yO11yErrorsCountIn).Execute()

Counts the grouped exceptions in the query window for the caller's org.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yErrorsCountIn := *openapiclient.NewO11yO11yErrorsCountIn() // O11yO11yErrorsCountIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yCounterrors(context.Background()).O11yO11yErrorsCountIn(o11yO11yErrorsCountIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yCounterrors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yCounterrors`: int32
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yCounterrors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yCounterrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yErrorsCountIn** | [**O11yO11yErrorsCountIn**](O11yO11yErrorsCountIn.md) |  | 

### Return type

**int32**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yDaemonsetsList

> O11yO11yDaemonSetListOut PostO11yDaemonsetsList(ctx).O11yDaemonSetListRequest(o11yDaemonSetListRequest).Execute()

Lists Kubernetes daemonsets over a time range, each with the CPU and memory its pods used against request and limit, desired and available node counts, restarts and attributes; filterable, groupable and paginated.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yDaemonSetListRequest := *openapiclient.NewO11yDaemonSetListRequest() // O11yDaemonSetListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yDaemonsetsList(context.Background()).O11yDaemonSetListRequest(o11yDaemonSetListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yDaemonsetsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yDaemonsetsList`: O11yO11yDaemonSetListOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yDaemonsetsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yDaemonsetsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yDaemonSetListRequest** | [**O11yDaemonSetListRequest**](O11yDaemonSetListRequest.md) |  | 

### Return type

[**O11yO11yDaemonSetListOut**](O11yO11yDaemonSetListOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yDependencyGraph

> []O11yO11yDependency PostO11yDependencyGraph(ctx).O11yO11yDependencyGraphIn(o11yO11yDependencyGraphIn).Execute()

Returns the service dependency graph over the requested window: every parent→child edge observed, with call and error rates and latency percentiles per edge.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yDependencyGraphIn := *openapiclient.NewO11yO11yDependencyGraphIn("End_example", "Start_example") // O11yO11yDependencyGraphIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yDependencyGraph(context.Background()).O11yO11yDependencyGraphIn(o11yO11yDependencyGraphIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yDependencyGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yDependencyGraph`: []O11yO11yDependency
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yDependencyGraph`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yDependencyGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yDependencyGraphIn** | [**O11yO11yDependencyGraphIn**](O11yO11yDependencyGraphIn.md) |  | 

### Return type

[**[]O11yO11yDependency**](O11yO11yDependency.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yDeploymentsList

> O11yO11yDeploymentListOut PostO11yDeploymentsList(ctx).O11yDeploymentListRequest(o11yDeploymentListRequest).Execute()

Lists Kubernetes deployments over a time range, each with the CPU and memory its pods used against request and limit, desired and available replica counts, restarts and attributes; filterable, groupable and paginated.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yDeploymentListRequest := *openapiclient.NewO11yDeploymentListRequest() // O11yDeploymentListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yDeploymentsList(context.Background()).O11yDeploymentListRequest(o11yDeploymentListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yDeploymentsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yDeploymentsList`: O11yO11yDeploymentListOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yDeploymentsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yDeploymentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yDeploymentListRequest** | [**O11yDeploymentListRequest**](O11yDeploymentListRequest.md) |  | 

### Return type

[**O11yO11yDeploymentListOut**](O11yO11yDeploymentListOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yErrortrackingIssuesById

> O11yO11yErrorIssueOut PostO11yErrortrackingIssuesById(ctx, id).O11yO11yErrorUpdateIssueIn(o11yO11yErrorUpdateIssueIn).Execute()

Changes an issue's lifecycle — resolve, ignore, reopen or assign — and returns the updated issue.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the issue id.
	o11yO11yErrorUpdateIssueIn := *openapiclient.NewO11yO11yErrorUpdateIssueIn("Id_example") // O11yO11yErrorUpdateIssueIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yErrortrackingIssuesById(context.Background(), id).O11yO11yErrorUpdateIssueIn(o11yO11yErrorUpdateIssueIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yErrortrackingIssuesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yErrortrackingIssuesById`: O11yO11yErrorIssueOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yErrortrackingIssuesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the issue id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yErrortrackingIssuesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yErrorUpdateIssueIn** | [**O11yO11yErrorUpdateIssueIn**](O11yO11yErrorUpdateIssueIn.md) |  | 

### Return type

[**O11yO11yErrorIssueOut**](O11yO11yErrorIssueOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yEvent

> O11yO11yMessage PostO11yEvent(ctx).O11yO11yEventIn(o11yO11yEventIn).Execute()

Records one product-analytics event for the signed-in user — a track event with a name and free-form attributes.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yEventIn := *openapiclient.NewO11yO11yEventIn("EventType_example") // O11yO11yEventIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yEvent(context.Background()).O11yO11yEventIn(o11yO11yEventIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yEvent`: O11yO11yMessage
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yEventIn** | [**O11yO11yEventIn**](O11yO11yEventIn.md) |  | 

### Return type

[**O11yO11yMessage**](O11yO11yMessage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yExplorerViews

> O11yO11ySavedViewCreateOut PostO11yExplorerViews(ctx).O11ySavedView(o11ySavedView).Execute()

Saves a new explorer view for the caller's org and returns its id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11ySavedView := *openapiclient.NewO11ySavedView() // O11ySavedView | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yExplorerViews(context.Background()).O11ySavedView(o11ySavedView).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yExplorerViews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yExplorerViews`: O11yO11ySavedViewCreateOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yExplorerViews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yExplorerViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11ySavedView** | [**O11ySavedView**](O11ySavedView.md) |  | 

### Return type

[**O11yO11ySavedViewCreateOut**](O11yO11ySavedViewCreateOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yExportRawData

> PostO11yExportRawData(ctx).Execute()

Export raw telemetry rows as a file



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.PostO11yExportRawData(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yExportRawData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yExportRawDataRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yHostsList

> O11yO11yHostListOut PostO11yHostsList(ctx).O11yHostListRequest(o11yHostListRequest).Execute()

Lists monitored hosts over a time range, each with its CPU, memory, I/O wait and 15-minute load, whether it is actively reporting, its OS and its attributes; filterable, groupable and paginated.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yHostListRequest := *openapiclient.NewO11yHostListRequest() // O11yHostListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yHostsList(context.Background()).O11yHostListRequest(o11yHostListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yHostsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yHostsList`: O11yO11yHostListOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yHostsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yHostsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yHostListRequest** | [**O11yHostListRequest**](O11yHostListRequest.md) |  | 

### Return type

[**O11yO11yHostListOut**](O11yO11yHostListOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yInfraMonitoringClusters

> O11yO11yInfraClustersOut PostO11yInfraMonitoringClusters(ctx).O11yPostableClusters(o11yPostableClusters).Execute()

Lists Kubernetes clusters with CPU and memory usage against allocatable capacity summed over their nodes, plus per-group node readiness and pod phase counts.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yPostableClusters := *openapiclient.NewO11yPostableClusters() // O11yPostableClusters | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yInfraMonitoringClusters(context.Background()).O11yPostableClusters(o11yPostableClusters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yInfraMonitoringClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yInfraMonitoringClusters`: O11yO11yInfraClustersOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yInfraMonitoringClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yInfraMonitoringClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yPostableClusters** | [**O11yPostableClusters**](O11yPostableClusters.md) |  | 

### Return type

[**O11yO11yInfraClustersOut**](O11yO11yInfraClustersOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yInfraMonitoringDaemonsets

> O11yO11yInfraDaemonSetsOut PostO11yInfraMonitoringDaemonsets(ctx).O11yPostableDaemonSets(o11yPostableDaemonSets).Execute()

Lists Kubernetes daemonsets with the CPU and memory their pods used against request and limit, the latest desired and current scheduled NODE counts (node counts, not pod counts), and per-group pod phase counts.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yPostableDaemonSets := *openapiclient.NewO11yPostableDaemonSets() // O11yPostableDaemonSets | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yInfraMonitoringDaemonsets(context.Background()).O11yPostableDaemonSets(o11yPostableDaemonSets).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yInfraMonitoringDaemonsets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yInfraMonitoringDaemonsets`: O11yO11yInfraDaemonSetsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yInfraMonitoringDaemonsets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yInfraMonitoringDaemonsetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yPostableDaemonSets** | [**O11yPostableDaemonSets**](O11yPostableDaemonSets.md) |  | 

### Return type

[**O11yO11yInfraDaemonSetsOut**](O11yO11yInfraDaemonSetsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yInfraMonitoringDeployments

> O11yO11yInfraDeploymentsOut PostO11yInfraMonitoringDeployments(ctx).O11yPostableDeployments(o11yPostableDeployments).Execute()

Lists Kubernetes deployments with the CPU and memory their pods used against request and limit, the latest desired and available replica counts, and per-group pod phase counts.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yPostableDeployments := *openapiclient.NewO11yPostableDeployments() // O11yPostableDeployments | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yInfraMonitoringDeployments(context.Background()).O11yPostableDeployments(o11yPostableDeployments).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yInfraMonitoringDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yInfraMonitoringDeployments`: O11yO11yInfraDeploymentsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yInfraMonitoringDeployments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yInfraMonitoringDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yPostableDeployments** | [**O11yPostableDeployments**](O11yPostableDeployments.md) |  | 

### Return type

[**O11yO11yInfraDeploymentsOut**](O11yO11yInfraDeploymentsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yInfraMonitoringHosts

> O11yO11yInfraHostsOut PostO11yInfraMonitoringHosts(ctx).O11yPostableHosts(o11yPostableHosts).Execute()

Lists hosts with key infrastructure metrics — CPU, memory, I/O wait and disk usage percentages and 15-minute load — plus an active/inactive status from whether the host reported in the last ten minutes.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yPostableHosts := *openapiclient.NewO11yPostableHosts() // O11yPostableHosts | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yInfraMonitoringHosts(context.Background()).O11yPostableHosts(o11yPostableHosts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yInfraMonitoringHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yInfraMonitoringHosts`: O11yO11yInfraHostsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yInfraMonitoringHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yInfraMonitoringHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yPostableHosts** | [**O11yPostableHosts**](O11yPostableHosts.md) |  | 

### Return type

[**O11yO11yInfraHostsOut**](O11yO11yInfraHostsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yInfraMonitoringJobs

> O11yO11yInfraJobsOut PostO11yInfraMonitoringJobs(ctx).O11yPostableJobs(o11yPostableJobs).Execute()

Lists Kubernetes jobs with the CPU and memory their pods used against request and limit, the latest desired-successful, active, failed and successful pod counters, and per-group pod phase counts — the phase counts are current state while the counters are cumulative over the job's life.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yPostableJobs := *openapiclient.NewO11yPostableJobs() // O11yPostableJobs | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yInfraMonitoringJobs(context.Background()).O11yPostableJobs(o11yPostableJobs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yInfraMonitoringJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yInfraMonitoringJobs`: O11yO11yInfraJobsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yInfraMonitoringJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yInfraMonitoringJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yPostableJobs** | [**O11yPostableJobs**](O11yPostableJobs.md) |  | 

### Return type

[**O11yO11yInfraJobsOut**](O11yO11yInfraJobsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yInfraMonitoringNamespaces

> O11yO11yInfraNamespacesOut PostO11yInfraMonitoringNamespaces(ctx).O11yPostableNamespaces(o11yPostableNamespaces).Execute()

Lists Kubernetes namespaces with the CPU and memory their pods used and per-group pod phase counts.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yPostableNamespaces := *openapiclient.NewO11yPostableNamespaces() // O11yPostableNamespaces | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yInfraMonitoringNamespaces(context.Background()).O11yPostableNamespaces(o11yPostableNamespaces).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yInfraMonitoringNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yInfraMonitoringNamespaces`: O11yO11yInfraNamespacesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yInfraMonitoringNamespaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yInfraMonitoringNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yPostableNamespaces** | [**O11yPostableNamespaces**](O11yPostableNamespaces.md) |  | 

### Return type

[**O11yO11yInfraNamespacesOut**](O11yO11yInfraNamespacesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yInfraMonitoringNodes

> O11yO11yInfraNodesOut PostO11yInfraMonitoringNodes(ctx).O11yPostableNodes(o11yPostableNodes).Execute()

Lists Kubernetes nodes with CPU and memory usage against allocatable capacity, per-group readiness counts and per-group phase counts for the pods scheduled on them.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yPostableNodes := *openapiclient.NewO11yPostableNodes() // O11yPostableNodes | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yInfraMonitoringNodes(context.Background()).O11yPostableNodes(o11yPostableNodes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yInfraMonitoringNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yInfraMonitoringNodes`: O11yO11yInfraNodesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yInfraMonitoringNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yInfraMonitoringNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yPostableNodes** | [**O11yPostableNodes**](O11yPostableNodes.md) |  | 

### Return type

[**O11yO11yInfraNodesOut**](O11yO11yInfraNodesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yInfraMonitoringPods

> O11yO11yInfraPodsOut PostO11yInfraMonitoringPods(ctx).O11yPostablePods(o11yPostablePods).Execute()

Lists Kubernetes pods with CPU and memory usage against request and limit, the pod's phase and its age, plus its namespace, node, owning workload and cluster attributes.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yPostablePods := *openapiclient.NewO11yPostablePods() // O11yPostablePods | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yInfraMonitoringPods(context.Background()).O11yPostablePods(o11yPostablePods).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yInfraMonitoringPods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yInfraMonitoringPods`: O11yO11yInfraPodsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yInfraMonitoringPods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yInfraMonitoringPodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yPostablePods** | [**O11yPostablePods**](O11yPostablePods.md) |  | 

### Return type

[**O11yO11yInfraPodsOut**](O11yO11yInfraPodsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yInfraMonitoringPvcs

> O11yO11yInfraVolumesOut PostO11yInfraMonitoringPvcs(ctx).O11yPostableVolumes(o11yPostableVolumes).Execute()

Lists Kubernetes persistent volume claims with available, capacity and used bytes and inode counts, plus the claim's pod, namespace, node, statefulset and cluster attributes.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yPostableVolumes := *openapiclient.NewO11yPostableVolumes() // O11yPostableVolumes | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yInfraMonitoringPvcs(context.Background()).O11yPostableVolumes(o11yPostableVolumes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yInfraMonitoringPvcs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yInfraMonitoringPvcs`: O11yO11yInfraVolumesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yInfraMonitoringPvcs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yInfraMonitoringPvcsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yPostableVolumes** | [**O11yPostableVolumes**](O11yPostableVolumes.md) |  | 

### Return type

[**O11yO11yInfraVolumesOut**](O11yO11yInfraVolumesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yInfraMonitoringStatefulsets

> O11yO11yInfraStatefulSetsOut PostO11yInfraMonitoringStatefulsets(ctx).O11yPostableStatefulSets(o11yPostableStatefulSets).Execute()

Lists Kubernetes statefulsets with the CPU and memory their pods used against request and limit, the latest desired and current replica counts, and per-group pod phase counts.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yPostableStatefulSets := *openapiclient.NewO11yPostableStatefulSets() // O11yPostableStatefulSets | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yInfraMonitoringStatefulsets(context.Background()).O11yPostableStatefulSets(o11yPostableStatefulSets).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yInfraMonitoringStatefulsets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yInfraMonitoringStatefulsets`: O11yO11yInfraStatefulSetsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yInfraMonitoringStatefulsets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yInfraMonitoringStatefulsetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yPostableStatefulSets** | [**O11yPostableStatefulSets**](O11yPostableStatefulSets.md) |  | 

### Return type

[**O11yO11yInfraStatefulSetsOut**](O11yO11yInfraStatefulSetsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yJobsList

> O11yO11yJobListOut PostO11yJobsList(ctx).O11yJobListRequest(o11yJobListRequest).Execute()

Lists Kubernetes jobs over a time range, each with the CPU and memory its pods used against request and limit, desired-successful, active, failed and successful pod counts, restarts and attributes; filterable, groupable and paginated.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yJobListRequest := *openapiclient.NewO11yJobListRequest() // O11yJobListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yJobsList(context.Background()).O11yJobListRequest(o11yJobListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yJobsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yJobsList`: O11yO11yJobListOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yJobsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yJobsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yJobListRequest** | [**O11yJobListRequest**](O11yJobListRequest.md) |  | 

### Return type

[**O11yO11yJobListOut**](O11yO11yJobListOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yListerrors

> []O11yO11yListError PostO11yListerrors(ctx).O11yO11yErrorsListIn(o11yO11yErrorsListIn).Execute()

Lists the grouped exceptions in the query window — each an exception type with its message, count, service and first/last-seen — for the caller's org.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yErrorsListIn := *openapiclient.NewO11yO11yErrorsListIn() // O11yO11yErrorsListIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yListerrors(context.Background()).O11yO11yErrorsListIn(o11yO11yErrorsListIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yListerrors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yListerrors`: []O11yO11yListError
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yListerrors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yListerrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yErrorsListIn** | [**O11yO11yErrorsListIn**](O11yO11yErrorsListIn.md) |  | 

### Return type

[**[]O11yO11yListError**](O11yO11yListError.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yLogsFields

> O11yO11yFieldSetting PostO11yLogsFields(ctx).O11yO11yFieldSetting(o11yO11yFieldSetting).Execute()

Changes how one log field is stored — selects or deselects it as a materialized column and tunes its index — and echoes the setting back.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yFieldSetting := *openapiclient.NewO11yO11yFieldSetting("DataType_example", "Name_example", "Type_example") // O11yO11yFieldSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yLogsFields(context.Background()).O11yO11yFieldSetting(o11yO11yFieldSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yLogsFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yLogsFields`: O11yO11yFieldSetting
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yLogsFields`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yLogsFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yFieldSetting** | [**O11yO11yFieldSetting**](O11yO11yFieldSetting.md) |  | 

### Return type

[**O11yO11yFieldSetting**](O11yO11yFieldSetting.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yLogsPipelines

> O11yO11yLogPipelinesOut PostO11yLogsPipelines(ctx).O11yO11yLogPipelineCreateIn(o11yO11yLogPipelineCreateIn).Execute()

Saves the given log parsing pipelines as the new config version for the caller's org and starts deploying it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yLogPipelineCreateIn := *openapiclient.NewO11yO11yLogPipelineCreateIn() // O11yO11yLogPipelineCreateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yLogsPipelines(context.Background()).O11yO11yLogPipelineCreateIn(o11yO11yLogPipelineCreateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yLogsPipelines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yLogsPipelines`: O11yO11yLogPipelinesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yLogsPipelines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yLogsPipelinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yLogPipelineCreateIn** | [**O11yO11yLogPipelineCreateIn**](O11yO11yLogPipelineCreateIn.md) |  | 

### Return type

[**O11yO11yLogPipelinesOut**](O11yO11yLogPipelinesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yLogsPipelinesPreview

> O11yO11yLogPipelinePreviewOut PostO11yLogsPipelinesPreview(ctx).O11yO11yLogPipelinePreviewIn(o11yO11yLogPipelinePreviewIn).Execute()

Runs the given log parsing pipelines over the given sample records without saving anything, and returns the transformed records plus whatever the collector logged while simulating them.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yLogPipelinePreviewIn := *openapiclient.NewO11yO11yLogPipelinePreviewIn() // O11yO11yLogPipelinePreviewIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yLogsPipelinesPreview(context.Background()).O11yO11yLogPipelinePreviewIn(o11yO11yLogPipelinePreviewIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yLogsPipelinesPreview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yLogsPipelinesPreview`: O11yO11yLogPipelinePreviewOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yLogsPipelinesPreview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yLogsPipelinesPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yLogPipelinePreviewIn** | [**O11yO11yLogPipelinePreviewIn**](O11yO11yLogPipelinePreviewIn.md) |  | 

### Return type

[**O11yO11yLogPipelinePreviewOut**](O11yO11yLogPipelinePreviewOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yLogsPromotePaths

> O11yO11yLogPromoteOut PostO11yLogsPromotePaths(ctx).O11yO11yLogPromotePath(o11yO11yLogPromotePath).Execute()

Promotes and indexes log body paths: each named path is lifted out of the JSON body into its own column, with the indexes the caller asked for.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yLogPromotePath := []openapiclient.O11yO11yLogPromotePath{*openapiclient.NewO11yO11yLogPromotePath()} // []O11yO11yLogPromotePath | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yLogsPromotePaths(context.Background()).O11yO11yLogPromotePath(o11yO11yLogPromotePath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yLogsPromotePaths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yLogsPromotePaths`: O11yO11yLogPromoteOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yLogsPromotePaths`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yLogsPromotePathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yLogPromotePath** | [**[]O11yO11yLogPromotePath**](O11yO11yLogPromotePath.md) |  | 

### Return type

[**O11yO11yLogPromoteOut**](O11yO11yLogPromoteOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yMessagingQueuesKafkaConsumerLagConsumerDetails

> O11yO11yQueryRangeOut PostO11yMessagingQueuesKafkaConsumerLagConsumerDetails(ctx).O11yO11yQueueIn(o11yO11yQueueIn).Execute()

Returns the consumer side of a consumer-lag view: the consumer groups reading the topic/partition named in variables, with their throughput and latency over the window.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yQueueIn := *openapiclient.NewO11yO11yQueueIn() // O11yO11yQueueIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yMessagingQueuesKafkaConsumerLagConsumerDetails(context.Background()).O11yO11yQueueIn(o11yO11yQueueIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yMessagingQueuesKafkaConsumerLagConsumerDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yMessagingQueuesKafkaConsumerLagConsumerDetails`: O11yO11yQueryRangeOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yMessagingQueuesKafkaConsumerLagConsumerDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yMessagingQueuesKafkaConsumerLagConsumerDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yQueueIn** | [**O11yO11yQueueIn**](O11yO11yQueueIn.md) |  | 

### Return type

[**O11yO11yQueryRangeOut**](O11yO11yQueryRangeOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yMessagingQueuesKafkaConsumerLagNetworkLatency

> O11yO11yQueryRangeOut PostO11yMessagingQueuesKafkaConsumerLagNetworkLatency(ctx).O11yO11yQueueIn(o11yO11yQueueIn).Execute()

Returns consumer network latency correlated per client: a throughput pass over the window finds the consumer clients, then their fetch latency joins in as a latency column per client/instance/service.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yQueueIn := *openapiclient.NewO11yO11yQueueIn() // O11yO11yQueueIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yMessagingQueuesKafkaConsumerLagNetworkLatency(context.Background()).O11yO11yQueueIn(o11yO11yQueueIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yMessagingQueuesKafkaConsumerLagNetworkLatency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yMessagingQueuesKafkaConsumerLagNetworkLatency`: O11yO11yQueryRangeOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yMessagingQueuesKafkaConsumerLagNetworkLatency`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yMessagingQueuesKafkaConsumerLagNetworkLatencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yQueueIn** | [**O11yO11yQueueIn**](O11yO11yQueueIn.md) |  | 

### Return type

[**O11yO11yQueryRangeOut**](O11yO11yQueryRangeOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yMessagingQueuesKafkaConsumerLagProducerDetails

> O11yO11yQueryRangeOut PostO11yMessagingQueuesKafkaConsumerLagProducerDetails(ctx).O11yO11yQueueIn(o11yO11yQueueIn).Execute()

Returns the producer side of a consumer-lag view: the producers writing to the topic/partition named in variables, with their throughput and latency over the window.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yQueueIn := *openapiclient.NewO11yO11yQueueIn() // O11yO11yQueueIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yMessagingQueuesKafkaConsumerLagProducerDetails(context.Background()).O11yO11yQueueIn(o11yO11yQueueIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yMessagingQueuesKafkaConsumerLagProducerDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yMessagingQueuesKafkaConsumerLagProducerDetails`: O11yO11yQueryRangeOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yMessagingQueuesKafkaConsumerLagProducerDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yMessagingQueuesKafkaConsumerLagProducerDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yQueueIn** | [**O11yO11yQueueIn**](O11yO11yQueueIn.md) |  | 

### Return type

[**O11yO11yQueryRangeOut**](O11yO11yQueryRangeOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yMessagingQueuesKafkaOnboardingConsumers

> O11yO11yQueueChecksOut PostO11yMessagingQueuesKafkaOnboardingConsumers(ctx).O11yO11yQueueIn(o11yO11yQueueIn).Execute()

Checks whether the spans the Kafka consumer views need are arriving, row for row like producersOnboarding.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yQueueIn := *openapiclient.NewO11yO11yQueueIn() // O11yO11yQueueIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yMessagingQueuesKafkaOnboardingConsumers(context.Background()).O11yO11yQueueIn(o11yO11yQueueIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yMessagingQueuesKafkaOnboardingConsumers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yMessagingQueuesKafkaOnboardingConsumers`: O11yO11yQueueChecksOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yMessagingQueuesKafkaOnboardingConsumers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yMessagingQueuesKafkaOnboardingConsumersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yQueueIn** | [**O11yO11yQueueIn**](O11yO11yQueueIn.md) |  | 

### Return type

[**O11yO11yQueueChecksOut**](O11yO11yQueueChecksOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yMessagingQueuesKafkaOnboardingKafka

> O11yO11yQueueChecksOut PostO11yMessagingQueuesKafkaOnboardingKafka(ctx).O11yO11yQueueIn(o11yO11yQueueIn).Execute()

Checks whether Kafka's own metrics — consumer lag and partition telemetry — are arriving, so the lag views can be lit up.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yQueueIn := *openapiclient.NewO11yO11yQueueIn() // O11yO11yQueueIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yMessagingQueuesKafkaOnboardingKafka(context.Background()).O11yO11yQueueIn(o11yO11yQueueIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yMessagingQueuesKafkaOnboardingKafka``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yMessagingQueuesKafkaOnboardingKafka`: O11yO11yQueueChecksOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yMessagingQueuesKafkaOnboardingKafka`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yMessagingQueuesKafkaOnboardingKafkaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yQueueIn** | [**O11yO11yQueueIn**](O11yO11yQueueIn.md) |  | 

### Return type

[**O11yO11yQueueChecksOut**](O11yO11yQueueChecksOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yMessagingQueuesKafkaOnboardingProducers

> O11yO11yQueueChecksOut PostO11yMessagingQueuesKafkaOnboardingProducers(ctx).O11yO11yQueueIn(o11yO11yQueueIn).Execute()

Checks whether the spans the Kafka producer views need are arriving — one row per required span attribute, with a pass/fail status and, on failure, what is missing from the instrumentation.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yQueueIn := *openapiclient.NewO11yO11yQueueIn() // O11yO11yQueueIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yMessagingQueuesKafkaOnboardingProducers(context.Background()).O11yO11yQueueIn(o11yO11yQueueIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yMessagingQueuesKafkaOnboardingProducers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yMessagingQueuesKafkaOnboardingProducers`: O11yO11yQueueChecksOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yMessagingQueuesKafkaOnboardingProducers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yMessagingQueuesKafkaOnboardingProducersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yQueueIn** | [**O11yO11yQueueIn**](O11yO11yQueueIn.md) |  | 

### Return type

[**O11yO11yQueueChecksOut**](O11yO11yQueueChecksOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yMessagingQueuesKafkaPartitionLatencyConsumer

> O11yO11yQueryRangeOut PostO11yMessagingQueuesKafkaPartitionLatencyConsumer(ctx).O11yO11yQueueIn(o11yO11yQueueIn).Execute()

Returns the consumer-group latency detail for the topic and partition named in the request's variables.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yQueueIn := *openapiclient.NewO11yO11yQueueIn() // O11yO11yQueueIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yMessagingQueuesKafkaPartitionLatencyConsumer(context.Background()).O11yO11yQueueIn(o11yO11yQueueIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yMessagingQueuesKafkaPartitionLatencyConsumer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yMessagingQueuesKafkaPartitionLatencyConsumer`: O11yO11yQueryRangeOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yMessagingQueuesKafkaPartitionLatencyConsumer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yMessagingQueuesKafkaPartitionLatencyConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yQueueIn** | [**O11yO11yQueueIn**](O11yO11yQueueIn.md) |  | 

### Return type

[**O11yO11yQueryRangeOut**](O11yO11yQueryRangeOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yMessagingQueuesKafkaPartitionLatencyOverview

> O11yO11yQueryRangeOut PostO11yMessagingQueuesKafkaPartitionLatencyOverview(ctx).O11yO11yQueueIn(o11yO11yQueueIn).Execute()

Returns the per-partition latency overview for the window — each topic/partition with its throughput and latency profile.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yQueueIn := *openapiclient.NewO11yO11yQueueIn() // O11yO11yQueueIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yMessagingQueuesKafkaPartitionLatencyOverview(context.Background()).O11yO11yQueueIn(o11yO11yQueueIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yMessagingQueuesKafkaPartitionLatencyOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yMessagingQueuesKafkaPartitionLatencyOverview`: O11yO11yQueryRangeOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yMessagingQueuesKafkaPartitionLatencyOverview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yMessagingQueuesKafkaPartitionLatencyOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yQueueIn** | [**O11yO11yQueueIn**](O11yO11yQueueIn.md) |  | 

### Return type

[**O11yO11yQueryRangeOut**](O11yO11yQueryRangeOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yMessagingQueuesKafkaSpanEvaluation

> O11yO11yQueryRangeOut PostO11yMessagingQueuesKafkaSpanEvaluation(ctx).O11yO11yQueueIn(o11yO11yQueueIn).Execute()

Correlates producer and consumer spans over the evaluation window (eval_time bounds the scan) and returns the pairings with their end-to-end delay — the check that messages produced are being consumed.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yQueueIn := *openapiclient.NewO11yO11yQueueIn() // O11yO11yQueueIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yMessagingQueuesKafkaSpanEvaluation(context.Background()).O11yO11yQueueIn(o11yO11yQueueIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yMessagingQueuesKafkaSpanEvaluation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yMessagingQueuesKafkaSpanEvaluation`: O11yO11yQueryRangeOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yMessagingQueuesKafkaSpanEvaluation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yMessagingQueuesKafkaSpanEvaluationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yQueueIn** | [**O11yO11yQueueIn**](O11yO11yQueueIn.md) |  | 

### Return type

[**O11yO11yQueryRangeOut**](O11yO11yQueryRangeOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yMessagingQueuesKafkaTopicThroughputConsumer

> O11yO11yQueryRangeOut PostO11yMessagingQueuesKafkaTopicThroughputConsumer(ctx).O11yO11yQueueIn(o11yO11yQueueIn).Execute()

Returns the consumer topic-throughput overview for the window — what each consumer group read, per topic.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yQueueIn := *openapiclient.NewO11yO11yQueueIn() // O11yO11yQueueIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yMessagingQueuesKafkaTopicThroughputConsumer(context.Background()).O11yO11yQueueIn(o11yO11yQueueIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yMessagingQueuesKafkaTopicThroughputConsumer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yMessagingQueuesKafkaTopicThroughputConsumer`: O11yO11yQueryRangeOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yMessagingQueuesKafkaTopicThroughputConsumer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yMessagingQueuesKafkaTopicThroughputConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yQueueIn** | [**O11yO11yQueueIn**](O11yO11yQueueIn.md) |  | 

### Return type

[**O11yO11yQueryRangeOut**](O11yO11yQueryRangeOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yMessagingQueuesKafkaTopicThroughputConsumerDetails

> O11yO11yQueryRangeOut PostO11yMessagingQueuesKafkaTopicThroughputConsumerDetails(ctx).O11yO11yQueueIn(o11yO11yQueueIn).Execute()

Breaks one consumer topic's throughput down using the topic and service named in variables.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yQueueIn := *openapiclient.NewO11yO11yQueueIn() // O11yO11yQueueIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yMessagingQueuesKafkaTopicThroughputConsumerDetails(context.Background()).O11yO11yQueueIn(o11yO11yQueueIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yMessagingQueuesKafkaTopicThroughputConsumerDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yMessagingQueuesKafkaTopicThroughputConsumerDetails`: O11yO11yQueryRangeOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yMessagingQueuesKafkaTopicThroughputConsumerDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yMessagingQueuesKafkaTopicThroughputConsumerDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yQueueIn** | [**O11yO11yQueueIn**](O11yO11yQueueIn.md) |  | 

### Return type

[**O11yO11yQueryRangeOut**](O11yO11yQueryRangeOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yMessagingQueuesKafkaTopicThroughputProducer

> O11yO11yQueryRangeOut PostO11yMessagingQueuesKafkaTopicThroughputProducer(ctx).O11yO11yQueueIn(o11yO11yQueueIn).Execute()

Returns the producer topic-throughput overview for the window — what each producer service wrote, per topic.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yQueueIn := *openapiclient.NewO11yO11yQueueIn() // O11yO11yQueueIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yMessagingQueuesKafkaTopicThroughputProducer(context.Background()).O11yO11yQueueIn(o11yO11yQueueIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yMessagingQueuesKafkaTopicThroughputProducer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yMessagingQueuesKafkaTopicThroughputProducer`: O11yO11yQueryRangeOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yMessagingQueuesKafkaTopicThroughputProducer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yMessagingQueuesKafkaTopicThroughputProducerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yQueueIn** | [**O11yO11yQueueIn**](O11yO11yQueueIn.md) |  | 

### Return type

[**O11yO11yQueryRangeOut**](O11yO11yQueryRangeOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yMessagingQueuesKafkaTopicThroughputProducerDetails

> O11yO11yQueryRangeOut PostO11yMessagingQueuesKafkaTopicThroughputProducerDetails(ctx).O11yO11yQueueIn(o11yO11yQueueIn).Execute()

Breaks one producer topic's throughput down using the topic and service named in variables.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yQueueIn := *openapiclient.NewO11yO11yQueueIn() // O11yO11yQueueIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yMessagingQueuesKafkaTopicThroughputProducerDetails(context.Background()).O11yO11yQueueIn(o11yO11yQueueIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yMessagingQueuesKafkaTopicThroughputProducerDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yMessagingQueuesKafkaTopicThroughputProducerDetails`: O11yO11yQueryRangeOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yMessagingQueuesKafkaTopicThroughputProducerDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yMessagingQueuesKafkaTopicThroughputProducerDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yQueueIn** | [**O11yO11yQueueIn**](O11yO11yQueueIn.md) |  | 

### Return type

[**O11yO11yQueryRangeOut**](O11yO11yQueryRangeOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yMessagingQueuesQueueOverview

> O11yO11yQueueRowsOut PostO11yMessagingQueuesQueueOverview(ctx).O11yO11yQueueListIn(o11yO11yQueueListIn).Execute()

Lists the messaging destinations observed in the window — one row per queue/destination/service combination with its throughput and latency columns.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yQueueListIn := *openapiclient.NewO11yO11yQueueListIn() // O11yO11yQueueListIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yMessagingQueuesQueueOverview(context.Background()).O11yO11yQueueListIn(o11yO11yQueueListIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yMessagingQueuesQueueOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yMessagingQueuesQueueOverview`: O11yO11yQueueRowsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yMessagingQueuesQueueOverview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yMessagingQueuesQueueOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yQueueListIn** | [**O11yO11yQueueListIn**](O11yO11yQueueListIn.md) |  | 

### Return type

[**O11yO11yQueueRowsOut**](O11yO11yQueueRowsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yNamespacesList

> O11yO11yNamespaceListOut PostO11yNamespacesList(ctx).O11yNamespaceListRequest(o11yNamespaceListRequest).Execute()

Lists Kubernetes namespaces over a time range, each with the CPU and memory its pods used, their phase counts and its attributes; filterable, groupable and paginated.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yNamespaceListRequest := *openapiclient.NewO11yNamespaceListRequest() // O11yNamespaceListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yNamespacesList(context.Background()).O11yNamespaceListRequest(o11yNamespaceListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yNamespacesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yNamespacesList`: O11yO11yNamespaceListOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yNamespacesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yNamespacesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yNamespaceListRequest** | [**O11yNamespaceListRequest**](O11yNamespaceListRequest.md) |  | 

### Return type

[**O11yO11yNamespaceListOut**](O11yO11yNamespaceListOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yNodesList

> O11yO11yNodeListOut PostO11yNodesList(ctx).O11yNodeListRequest(o11yNodeListRequest).Execute()

Lists Kubernetes nodes over a time range, each with its CPU and memory usage against allocatable capacity, readiness condition counts and attributes; filterable, groupable and paginated.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yNodeListRequest := *openapiclient.NewO11yNodeListRequest() // O11yNodeListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yNodesList(context.Background()).O11yNodeListRequest(o11yNodeListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yNodesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yNodesList`: O11yO11yNodeListOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yNodesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yNodesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yNodeListRequest** | [**O11yNodeListRequest**](O11yNodeListRequest.md) |  | 

### Return type

[**O11yO11yNodeListOut**](O11yO11yNodeListOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yPodsList

> O11yO11yPodListOut PostO11yPodsList(ctx).O11yPodListRequest(o11yPodListRequest).Execute()

Lists Kubernetes pods over a time range, each with its CPU and memory usage against request and limit, restart count, phase counts and attributes; filterable, groupable and paginated.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yPodListRequest := *openapiclient.NewO11yPodListRequest() // O11yPodListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yPodsList(context.Background()).O11yPodListRequest(o11yPodListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yPodsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yPodsList`: O11yO11yPodListOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yPodsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yPodsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yPodListRequest** | [**O11yPodListRequest**](O11yPodListRequest.md) |  | 

### Return type

[**O11yO11yPodListOut**](O11yO11yPodListOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yProcessesList

> O11yO11yProcessListOut PostO11yProcessesList(ctx).O11yProcessListRequest(o11yProcessListRequest).Execute()

Lists monitored processes over a time range, each with its name, PID, command line and CPU and memory usage; filterable, groupable and paginated.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yProcessListRequest := *openapiclient.NewO11yProcessListRequest() // O11yProcessListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yProcessesList(context.Background()).O11yProcessListRequest(o11yProcessListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yProcessesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yProcessesList`: O11yO11yProcessListOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yProcessesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yProcessesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yProcessListRequest** | [**O11yProcessListRequest**](O11yProcessListRequest.md) |  | 

### Return type

[**O11yO11yProcessListOut**](O11yO11yProcessListOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yPvcsList

> O11yO11yPvcListOut PostO11yPvcsList(ctx).O11yVolumeListRequest(o11yVolumeListRequest).Execute()

Lists Kubernetes persistent volume claims over a time range, each with its available, capacity and used bytes, inode counts and attributes; filterable, groupable and paginated.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yVolumeListRequest := *openapiclient.NewO11yVolumeListRequest() // O11yVolumeListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yPvcsList(context.Background()).O11yVolumeListRequest(o11yVolumeListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yPvcsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yPvcsList`: O11yO11yPvcListOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yPvcsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yPvcsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yVolumeListRequest** | [**O11yVolumeListRequest**](O11yVolumeListRequest.md) |  | 

### Return type

[**O11yO11yPvcListOut**](O11yO11yPvcListOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yQueryFilterAnalyze

> O11yO11yAnalyzeOut PostO11yQueryFilterAnalyze(ctx).O11yO11yAnalyzeIn(o11yO11yAnalyzeIn).Execute()

Analyzes a query and extracts the metric names it reads and the columns it groups by.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yAnalyzeIn := *openapiclient.NewO11yO11yAnalyzeIn("Query_example", "QueryType_example") // O11yO11yAnalyzeIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yQueryFilterAnalyze(context.Background()).O11yO11yAnalyzeIn(o11yO11yAnalyzeIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yQueryFilterAnalyze``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yQueryFilterAnalyze`: O11yO11yAnalyzeOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yQueryFilterAnalyze`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yQueryFilterAnalyzeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yAnalyzeIn** | [**O11yO11yAnalyzeIn**](O11yO11yAnalyzeIn.md) |  | 

### Return type

[**O11yO11yAnalyzeOut**](O11yO11yAnalyzeOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yQueryRange

> O11yO11yQueryRangeOut PostO11yQueryRange(ctx).O11yQueryRangeRequest(o11yQueryRangeRequest).Execute()

Executes a composite query over a time range: builder queries over traces, logs and metrics, formulas, trace operators, PromQL and Datastore SQL, answering time series, scalars or raw records as the request type asks.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yQueryRangeRequest := *openapiclient.NewO11yQueryRangeRequest() // O11yQueryRangeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yQueryRange(context.Background()).O11yQueryRangeRequest(o11yQueryRangeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yQueryRange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yQueryRange`: O11yO11yQueryRangeOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yQueryRange`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yQueryRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yQueryRangeRequest** | [**O11yQueryRangeRequest**](O11yQueryRangeRequest.md) |  | 

### Return type

[**O11yO11yQueryRangeOut**](O11yO11yQueryRangeOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yQueryRangeFormat

> O11yO11yQueryRangeFormatOut PostO11yQueryRangeFormat(ctx).O11yQueryRangeParamsV3(o11yQueryRangeParamsV3).Execute()

Parses a builder query and echoes it back normalized to the v3 shape — the endpoint the UI uses to canonicalize a query without running it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yQueryRangeParamsV3 := *openapiclient.NewO11yQueryRangeParamsV3() // O11yQueryRangeParamsV3 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yQueryRangeFormat(context.Background()).O11yQueryRangeParamsV3(o11yQueryRangeParamsV3).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yQueryRangeFormat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yQueryRangeFormat`: O11yO11yQueryRangeFormatOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yQueryRangeFormat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yQueryRangeFormatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yQueryRangeParamsV3** | [**O11yQueryRangeParamsV3**](O11yQueryRangeParamsV3.md) |  | 

### Return type

[**O11yO11yQueryRangeFormatOut**](O11yO11yQueryRangeFormatOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yQueryRangePreview

> O11yO11yQueryRangePreviewOut PostO11yQueryRangePreview(ctx).O11yO11yQueryRangePreviewIn(o11yO11yQueryRangePreviewIn).Execute()

Validates a composite query and renders the Datastore statements it would run WITHOUT executing it — a dry run for agentic and tooling use.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yQueryRangePreviewIn := *openapiclient.NewO11yO11yQueryRangePreviewIn() // O11yO11yQueryRangePreviewIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yQueryRangePreview(context.Background()).O11yO11yQueryRangePreviewIn(o11yO11yQueryRangePreviewIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yQueryRangePreview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yQueryRangePreview`: O11yO11yQueryRangePreviewOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yQueryRangePreview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yQueryRangePreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yQueryRangePreviewIn** | [**O11yO11yQueryRangePreviewIn**](O11yO11yQueryRangePreviewIn.md) |  | 

### Return type

[**O11yO11yQueryRangePreviewOut**](O11yO11yQueryRangePreviewOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yRegister

> O11yO11yRegisterOut PostO11yRegister(ctx).O11yO11yRegisterIn(o11yO11yRegisterIn).Execute()

Creates the FIRST organization and its admin user.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yRegisterIn := *openapiclient.NewO11yO11yRegisterIn("Email_example") // O11yO11yRegisterIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yRegister(context.Background()).O11yO11yRegisterIn(o11yO11yRegisterIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yRegister``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yRegister`: O11yO11yRegisterOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yRegister`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yRegisterIn** | [**O11yO11yRegisterIn**](O11yO11yRegisterIn.md) |  | 

### Return type

[**O11yO11yRegisterOut**](O11yO11yRegisterOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yReviews

> O11yAnnQueueView PostO11yReviews(ctx).O11yCreateQueueReq(o11yCreateQueueReq).Execute()

Creates a human-review queue in the caller's org and project.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yCreateQueueReq := *openapiclient.NewO11yCreateQueueReq() // O11yCreateQueueReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yReviews(context.Background()).O11yCreateQueueReq(o11yCreateQueueReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yReviews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yReviews`: O11yAnnQueueView
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yReviews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yReviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yCreateQueueReq** | [**O11yCreateQueueReq**](O11yCreateQueueReq.md) |  | 

### Return type

[**O11yAnnQueueView**](O11yAnnQueueView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yReviewsByIdItems

> O11yAnnItemsCreated PostO11yReviewsByIdItems(ctx, id).O11yAddItemsIn(o11yAddItemsIn).Execute()

Enqueues traces, observations or sessions on a review queue.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "annq_1" // string | ID is the annotation queue to add to, from the path.
	o11yAddItemsIn := *openapiclient.NewO11yAddItemsIn() // O11yAddItemsIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yReviewsByIdItems(context.Background(), id).O11yAddItemsIn(o11yAddItemsIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yReviewsByIdItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yReviewsByIdItems`: O11yAnnItemsCreated
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yReviewsByIdItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the annotation queue to add to, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yReviewsByIdItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yAddItemsIn** | [**O11yAddItemsIn**](O11yAddItemsIn.md) |  | 

### Return type

[**O11yAnnItemsCreated**](O11yAnnItemsCreated.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11ySentinelDiscover

> O11yO11yDiscoverOut PostO11ySentinelDiscover(ctx).O11yO11yDiscoverIn(o11yO11yDiscoverIn).Execute()

Aggregates a project's captured errors into a table — the caller names the filters, the groupings and the aggregations, and gets back the columns and rows they asked for.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yDiscoverIn := *openapiclient.NewO11yO11yDiscoverIn("Project_example") // O11yO11yDiscoverIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11ySentinelDiscover(context.Background()).O11yO11yDiscoverIn(o11yO11yDiscoverIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11ySentinelDiscover``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11ySentinelDiscover`: O11yO11yDiscoverOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11ySentinelDiscover`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11ySentinelDiscoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yDiscoverIn** | [**O11yO11yDiscoverIn**](O11yO11yDiscoverIn.md) |  | 

### Return type

[**O11yO11yDiscoverOut**](O11yO11yDiscoverOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11ySentinelProjects

> O11yO11ySentryProjectOut PostO11ySentinelProjects(ctx).O11yO11ySentryPostableProject(o11yO11ySentryPostableProject).Execute()

Creates a Sentry project under the caller's org and returns it, DSN included.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11ySentryPostableProject := *openapiclient.NewO11yO11ySentryPostableProject("Name_example") // O11yO11ySentryPostableProject | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11ySentinelProjects(context.Background()).O11yO11ySentryPostableProject(o11yO11ySentryPostableProject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11ySentinelProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11ySentinelProjects`: O11yO11ySentryProjectOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11ySentinelProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11ySentinelProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11ySentryPostableProject** | [**O11yO11ySentryPostableProject**](O11yO11ySentryPostableProject.md) |  | 

### Return type

[**O11yO11ySentryProjectOut**](O11yO11ySentryProjectOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11ySentinelProjectsByIdKeysRotate

> O11yO11ySentryProjectOut PostO11ySentinelProjectsByIdKeysRotate(ctx, id).Execute()

Rotates a project's DSN key — bumping its rotation watermark so keys below it stop verifying — and returns the project with its new DSN.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the project id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11ySentinelProjectsByIdKeysRotate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11ySentinelProjectsByIdKeysRotate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11ySentinelProjectsByIdKeysRotate`: O11yO11ySentryProjectOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11ySentinelProjectsByIdKeysRotate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the project id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostO11ySentinelProjectsByIdKeysRotateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yO11ySentryProjectOut**](O11yO11ySentryProjectOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yServiceEntryPointOperations

> O11yO11yOperationsOut PostO11yServiceEntryPointOperations(ctx).O11yO11yOperationsIn(o11yO11yOperationsIn).Execute()

Returns one service's entry-point operations with the same latency and error profile topOperations reports.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yOperationsIn := *openapiclient.NewO11yO11yOperationsIn() // O11yO11yOperationsIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yServiceEntryPointOperations(context.Background()).O11yO11yOperationsIn(o11yO11yOperationsIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yServiceEntryPointOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yServiceEntryPointOperations`: O11yO11yOperationsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yServiceEntryPointOperations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yServiceEntryPointOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yOperationsIn** | [**O11yO11yOperationsIn**](O11yO11yOperationsIn.md) |  | 

### Return type

[**O11yO11yOperationsOut**](O11yO11yOperationsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yServiceTopLevelOperations

> map[string][]string PostO11yServiceTopLevelOperations(ctx).O11yO11yTopLevelOpsIn(o11yO11yTopLevelOpsIn).Execute()

Maps each service to its entry-point span names — for the one service named in the request, or for every service when none is.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yTopLevelOpsIn := *openapiclient.NewO11yO11yTopLevelOpsIn() // O11yO11yTopLevelOpsIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yServiceTopLevelOperations(context.Background()).O11yO11yTopLevelOpsIn(o11yO11yTopLevelOpsIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yServiceTopLevelOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yServiceTopLevelOperations`: map[string][]string
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yServiceTopLevelOperations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yServiceTopLevelOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yTopLevelOpsIn** | [**O11yO11yTopLevelOpsIn**](O11yO11yTopLevelOpsIn.md) |  | 

### Return type

[**map[string][]string**](array.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yServiceTopOperations

> O11yO11yOperationsOut PostO11yServiceTopOperations(ctx).O11yO11yOperationsIn(o11yO11yOperationsIn).Execute()

Returns one service's heaviest operations in the window, each with p50/p95/p99 latency, how often it ran and how often it errored.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yOperationsIn := *openapiclient.NewO11yO11yOperationsIn() // O11yO11yOperationsIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yServiceTopOperations(context.Background()).O11yO11yOperationsIn(o11yO11yOperationsIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yServiceTopOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yServiceTopOperations`: O11yO11yOperationsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yServiceTopOperations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yServiceTopOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yOperationsIn** | [**O11yO11yOperationsIn**](O11yO11yOperationsIn.md) |  | 

### Return type

[**O11yO11yOperationsOut**](O11yO11yOperationsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yServices

> O11yO11yServicesOut PostO11yServices(ctx).O11yO11yServicesIn(o11yO11yServicesIn).Execute()

Lists the instrumented services seen in the window, each with the request profile of its entry-point spans: p99 and average latency, call and error rates, and the entry-point operations the numbers were computed over.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yServicesIn := *openapiclient.NewO11yO11yServicesIn() // O11yO11yServicesIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yServices(context.Background()).O11yO11yServicesIn(o11yO11yServicesIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yServices`: O11yO11yServicesOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yServicesIn** | [**O11yO11yServicesIn**](O11yO11yServicesIn.md) |  | 

### Return type

[**O11yO11yServicesOut**](O11yO11yServicesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11ySettingsApdex

> O11yO11yApdexSetOut PostO11ySettingsApdex(ctx).O11yO11yApdexSetIn(o11yO11yApdexSetIn).Execute()

Sets one service's apdex threshold and the status codes excluded from its score.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yApdexSetIn := *openapiclient.NewO11yO11yApdexSetIn() // O11yO11yApdexSetIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11ySettingsApdex(context.Background()).O11yO11yApdexSetIn(o11yO11yApdexSetIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11ySettingsApdex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11ySettingsApdex`: O11yO11yApdexSetOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11ySettingsApdex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11ySettingsApdexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yApdexSetIn** | [**O11yO11yApdexSetIn**](O11yO11yApdexSetIn.md) |  | 

### Return type

[**O11yO11yApdexSetOut**](O11yO11yApdexSetOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11ySettingsTtl

> O11yO11yRetentionSetOut PostO11ySettingsTtl(ctx).O11yO11yRetentionSetIn(o11yO11yRetentionSetIn).Execute()

Sets the org's retention policy for one signal: the default TTL in days, ordered per-label retention rules, and optional cold-storage settings.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yRetentionSetIn := *openapiclient.NewO11yO11yRetentionSetIn() // O11yO11yRetentionSetIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11ySettingsTtl(context.Background()).O11yO11yRetentionSetIn(o11yO11yRetentionSetIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11ySettingsTtl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11ySettingsTtl`: O11yO11yRetentionSetOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11ySettingsTtl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11ySettingsTtlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yRetentionSetIn** | [**O11yO11yRetentionSetIn**](O11yO11yRetentionSetIn.md) |  | 

### Return type

[**O11yO11yRetentionSetOut**](O11yO11yRetentionSetOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11ySpanPercentile

> O11yO11ySpanPercentileOut PostO11ySpanPercentile(ctx).O11yO11ySpanPercentileIn(o11yO11ySpanPercentileIn).Execute()

Places one span's duration among its peers: the p50/p90/p99 durations of like spans, and the percentile the given duration lands at.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11ySpanPercentileIn := *openapiclient.NewO11yO11ySpanPercentileIn("Name_example", "ServiceName_example") // O11yO11ySpanPercentileIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11ySpanPercentile(context.Background()).O11yO11ySpanPercentileIn(o11yO11ySpanPercentileIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11ySpanPercentile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11ySpanPercentile`: O11yO11ySpanPercentileOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11ySpanPercentile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11ySpanPercentileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11ySpanPercentileIn** | [**O11yO11ySpanPercentileIn**](O11yO11ySpanPercentileIn.md) |  | 

### Return type

[**O11yO11ySpanPercentileOut**](O11yO11ySpanPercentileOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yStatefulsetsList

> O11yO11yStatefulSetListOut PostO11yStatefulsetsList(ctx).O11yStatefulSetListRequest(o11yStatefulSetListRequest).Execute()

Lists Kubernetes statefulsets over a time range, each with the CPU and memory its pods used against request and limit, desired and available replica counts, restarts and attributes; filterable, groupable and paginated.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yStatefulSetListRequest := *openapiclient.NewO11yStatefulSetListRequest() // O11yStatefulSetListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yStatefulsetsList(context.Background()).O11yStatefulSetListRequest(o11yStatefulSetListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yStatefulsetsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yStatefulsetsList`: O11yO11yStatefulSetListOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yStatefulsetsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yStatefulsetsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yStatefulSetListRequest** | [**O11yStatefulSetListRequest**](O11yStatefulSetListRequest.md) |  | 

### Return type

[**O11yO11yStatefulSetListOut**](O11yO11yStatefulSetListOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11ySubstituteVars

> O11yO11ySubstituteVarsOut PostO11ySubstituteVars(ctx).O11yQueryRangeRequest(o11yQueryRangeRequest).Execute()

Substitutes a query's variables and returns the resolved request, without running it — what a dashboard does before it queries.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yQueryRangeRequest := *openapiclient.NewO11yQueryRangeRequest() // O11yQueryRangeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11ySubstituteVars(context.Background()).O11yQueryRangeRequest(o11yQueryRangeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11ySubstituteVars``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11ySubstituteVars`: O11yO11ySubstituteVarsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11ySubstituteVars`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11ySubstituteVarsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yQueryRangeRequest** | [**O11yQueryRangeRequest**](O11yQueryRangeRequest.md) |  | 

### Return type

[**O11yO11ySubstituteVarsOut**](O11yO11ySubstituteVarsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yThirdPartyApisOverviewDomain

> O11yO11yDomainsOut PostO11yThirdPartyApisOverviewDomain(ctx).O11yO11yDomainsIn(o11yO11yDomainsIn).Execute()

Returns one external domain's endpoint-level breakdown — each endpoint with its rate, error and latency columns over the window.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yDomainsIn := *openapiclient.NewO11yO11yDomainsIn() // O11yO11yDomainsIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yThirdPartyApisOverviewDomain(context.Background()).O11yO11yDomainsIn(o11yO11yDomainsIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yThirdPartyApisOverviewDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yThirdPartyApisOverviewDomain`: O11yO11yDomainsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yThirdPartyApisOverviewDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yThirdPartyApisOverviewDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yDomainsIn** | [**O11yO11yDomainsIn**](O11yO11yDomainsIn.md) |  | 

### Return type

[**O11yO11yDomainsOut**](O11yO11yDomainsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yThirdPartyApisOverviewList

> O11yO11yDomainsOut PostO11yThirdPartyApisOverviewList(ctx).O11yO11yDomainsIn(o11yO11yDomainsIn).Execute()

Lists the external domains the instrumented services call, with request rate, error percentage and latency per domain.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yDomainsIn := *openapiclient.NewO11yO11yDomainsIn() // O11yO11yDomainsIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yThirdPartyApisOverviewList(context.Background()).O11yO11yDomainsIn(o11yO11yDomainsIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yThirdPartyApisOverviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yThirdPartyApisOverviewList`: O11yO11yDomainsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yThirdPartyApisOverviewList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yThirdPartyApisOverviewListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yDomainsIn** | [**O11yO11yDomainsIn**](O11yO11yDomainsIn.md) |  | 

### Return type

[**O11yO11yDomainsOut**](O11yO11yDomainsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostO11yVariablesQuery

> O11yO11yDashboardVarsOut PostO11yVariablesQuery(ctx).O11yO11yDashboardVarsIn(o11yO11yDashboardVarsIn).Execute()

Evaluates a dashboard variable query and returns the values the variable may take.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yDashboardVarsIn := *openapiclient.NewO11yO11yDashboardVarsIn("Query_example") // O11yO11yDashboardVarsIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PostO11yVariablesQuery(context.Background()).O11yO11yDashboardVarsIn(o11yO11yDashboardVarsIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PostO11yVariablesQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostO11yVariablesQuery`: O11yO11yDashboardVarsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PostO11yVariablesQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostO11yVariablesQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yDashboardVarsIn** | [**O11yO11yDashboardVarsIn**](O11yO11yDashboardVarsIn.md) |  | 

### Return type

[**O11yO11yDashboardVarsOut**](O11yO11yDashboardVarsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewMetricReductionRule

> O11yO11yReductionRulePreviewOut PreviewMetricReductionRule(ctx).O11yO11yReductionRulePreviewIn(o11yO11yReductionRulePreviewIn).Execute()

Estimates the series reduction and the dashboards and alerts a candidate volume-control rule would touch, without persisting it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yReductionRulePreviewIn := *openapiclient.NewO11yO11yReductionRulePreviewIn([]string{"Labels_example"}, "MatchType_example", "MetricName_example") // O11yO11yReductionRulePreviewIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PreviewMetricReductionRule(context.Background()).O11yO11yReductionRulePreviewIn(o11yO11yReductionRulePreviewIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PreviewMetricReductionRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreviewMetricReductionRule`: O11yO11yReductionRulePreviewOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PreviewMetricReductionRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewMetricReductionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yReductionRulePreviewIn** | [**O11yO11yReductionRulePreviewIn**](O11yO11yReductionRulePreviewIn.md) |  | 

### Return type

[**O11yO11yReductionRulePreviewOut**](O11yO11yReductionRulePreviewOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutHost

> PutHost(ctx).O11yPostableHost(o11yPostableHost).Execute()

Records the deployment's host in Zeus, overwriting any prior one.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yPostableHost := *openapiclient.NewO11yPostableHost() // O11yPostableHost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.PutHost(context.Background()).O11yPostableHost(o11yPostableHost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PutHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yPostableHost** | [**O11yPostableHost**](O11yPostableHost.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutO11yExplorerViewsByViewid

> O11yO11ySavedViewOut PutO11yExplorerViewsByViewid(ctx, viewId).O11yO11ySavedViewUpdateIn(o11yO11ySavedViewUpdateIn).Execute()

Replaces one saved explorer view by id with the given view and echoes it back.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	viewId := "viewId_example" // string | ViewID is the id of the view to replace, taken from the URL.
	o11yO11ySavedViewUpdateIn := *openapiclient.NewO11yO11ySavedViewUpdateIn() // O11yO11ySavedViewUpdateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PutO11yExplorerViewsByViewid(context.Background(), viewId).O11yO11ySavedViewUpdateIn(o11yO11ySavedViewUpdateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PutO11yExplorerViewsByViewid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutO11yExplorerViewsByViewid`: O11yO11ySavedViewOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PutO11yExplorerViewsByViewid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **string** | ViewID is the id of the view to replace, taken from the URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutO11yExplorerViewsByViewidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11ySavedViewUpdateIn** | [**O11yO11ySavedViewUpdateIn**](O11yO11ySavedViewUpdateIn.md) |  | 

### Return type

[**O11yO11ySavedViewOut**](O11yO11ySavedViewOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutO11ySentinelIssuesById

> O11yO11yErrorIssueOut PutO11ySentinelIssuesById(ctx, id).O11yO11ySentryUpdateIssueIn(o11yO11ySentryUpdateIssueIn).Execute()

Changes an issue's lifecycle — resolve, ignore, reopen or assign — and returns the updated issue.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the issue id.
	o11yO11ySentryUpdateIssueIn := *openapiclient.NewO11yO11ySentryUpdateIssueIn("Id_example") // O11yO11ySentryUpdateIssueIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.PutO11ySentinelIssuesById(context.Background(), id).O11yO11ySentryUpdateIssueIn(o11yO11ySentryUpdateIssueIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PutO11ySentinelIssuesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutO11ySentinelIssuesById`: O11yO11yErrorIssueOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.PutO11ySentinelIssuesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the issue id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutO11ySentinelIssuesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11ySentryUpdateIssueIn** | [**O11yO11ySentryUpdateIssueIn**](O11yO11ySentryUpdateIssueIn.md) |  | 

### Return type

[**O11yO11yErrorIssueOut**](O11yO11yErrorIssueOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProfile

> PutProfile(ctx).O11yPostableProfile(o11yPostableProfile).Execute()

Records the deployment's profile in Zeus — how the team uses observability today and what they plan — overwriting any prior one.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yPostableProfile := *openapiclient.NewO11yPostableProfile() // O11yPostableProfile | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.PutProfile(context.Background()).O11yPostableProfile(o11yPostableProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.PutProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yPostableProfile** | [**O11yPostableProfile**](O11yPostableProfile.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUserRoleByUserIDAndRoleID

> RemoveUserRoleByUserIDAndRoleID(ctx, id, roleId).Execute()

Takes a role away from one org member, by user id and role id — someone else, never the caller.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 
	roleId := "roleId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.RemoveUserRoleByUserIDAndRoleID(context.Background(), id, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.RemoveUserRoleByUserIDAndRoleID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserRoleByUserIDAndRoleIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPassword

> ResetPassword(ctx).O11yO11yResetPasswordIn(o11yO11yResetPasswordIn).Execute()

Sets a new password for whoever the reset token was minted for, consuming the token.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yResetPasswordIn := *openapiclient.NewO11yO11yResetPasswordIn() // O11yO11yResetPasswordIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.ResetPassword(context.Background()).O11yO11yResetPasswordIn(o11yO11yResetPasswordIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ResetPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yResetPasswordIn** | [**O11yO11yResetPasswordIn**](O11yO11yResetPasswordIn.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeServiceAccountKey

> RevokeServiceAccountKey(ctx, id, fid).Execute()

Revokes an API key.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 
	fid := "fid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.RevokeServiceAccountKey(context.Background(), id, fid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.RevokeServiceAccountKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**fid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeServiceAccountKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateSession

> O11yO11yTokenOut RotateSession(ctx).O11yO11yRotateSessionIn(o11yO11yRotateSessionIn).Execute()

Exchanges a refresh token for a fresh token pair, retiring the old pair.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yRotateSessionIn := *openapiclient.NewO11yO11yRotateSessionIn() // O11yO11yRotateSessionIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.RotateSession(context.Background()).O11yO11yRotateSessionIn(o11yO11yRotateSessionIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.RotateSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateSession`: O11yO11yTokenOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.RotateSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRotateSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yRotateSessionIn** | [**O11yO11yRotateSessionIn**](O11yO11yRotateSessionIn.md) |  | 

### Return type

[**O11yO11yTokenOut**](O11yO11yTokenOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchIngestionKeys

> O11yO11yIngestionKeysOut SearchIngestionKeys(ctx).Name(name).Page(page).PerPage(perPage).Execute()

Lists the workspace's ingestion keys whose name matches the search, paginated.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	name := "name_example" // string | Name is the substring to match ingestion-key names against. (optional)
	page := int32(56) // int32 | Page is the 1-based page number. (optional)
	perPage := int32(56) // int32 | PerPage is the page size. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.SearchIngestionKeys(context.Background()).Name(name).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.SearchIngestionKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchIngestionKeys`: O11yO11yIngestionKeysOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.SearchIngestionKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchIngestionKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name is the substring to match ingestion-key names against. | 
 **page** | **int32** | Page is the 1-based page number. | 
 **perPage** | **int32** | PerPage is the page size. | 

### Return type

[**O11yO11yIngestionKeysOut**](O11yO11yIngestionKeysOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchTraces

> []O11yO11yTraceSpanWindow SearchTraces(ctx, traceId).SpanId(spanId).LevelUp(levelUp).LevelDown(levelDown).SpanRenderLimit(spanRenderLimit).Execute()

Returns one trace's spans as a column/row table, optionally centred on a span and walked a fixed number of levels up and down from it — the read the trace explorer opens a trace with.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	traceId := "traceId_example" // string | 
	spanId := "spanId_example" // string |  (optional)
	levelUp := int32(56) // int32 |  (optional)
	levelDown := int32(56) // int32 |  (optional)
	spanRenderLimit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.SearchTraces(context.Background(), traceId).SpanId(spanId).LevelUp(levelUp).LevelDown(levelDown).SpanRenderLimit(spanRenderLimit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.SearchTraces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchTraces`: []O11yO11yTraceSpanWindow
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.SearchTraces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchTracesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spanId** | **string** |  | 
 **levelUp** | **int32** |  | 
 **levelDown** | **int32** |  | 
 **spanRenderLimit** | **int32** |  | 

### Return type

[**[]O11yO11yTraceSpanWindow**](O11yO11yTraceSpanWindow.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRoleByUserID

> O11yO11yAck SetRoleByUserID(ctx, id).O11yO11ySetRoleIn(o11yO11ySetRoleIn).Execute()

Assigns a role, by role name, to one org member — someone else, never the caller.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 
	o11yO11ySetRoleIn := *openapiclient.NewO11yO11ySetRoleIn() // O11yO11ySetRoleIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.SetRoleByUserID(context.Background(), id).O11yO11ySetRoleIn(o11yO11ySetRoleIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.SetRoleByUserID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetRoleByUserID`: O11yO11yAck
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.SetRoleByUserID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetRoleByUserIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11ySetRoleIn** | [**O11yO11ySetRoleIn**](O11yO11ySetRoleIn.md) |  | 

### Return type

[**O11yO11yAck**](O11yO11yAck.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestChannel

> TestChannel(ctx).O11yAlertmanagertypesReceiver(o11yAlertmanagertypesReceiver).Execute()

Sends a test notification to the posted receiver.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yAlertmanagertypesReceiver := *openapiclient.NewO11yAlertmanagertypesReceiver() // O11yAlertmanagertypesReceiver | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.TestChannel(context.Background()).O11yAlertmanagertypesReceiver(o11yAlertmanagertypesReceiver).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.TestChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yAlertmanagertypesReceiver** | [**O11yAlertmanagertypesReceiver**](O11yAlertmanagertypesReceiver.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestChannelDeprecated

> TestChannelDeprecated(ctx).O11yAlertmanagertypesReceiver(o11yAlertmanagertypesReceiver).Execute()

Sends a test notification to the posted receiver.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yAlertmanagertypesReceiver := *openapiclient.NewO11yAlertmanagertypesReceiver() // O11yAlertmanagertypesReceiver | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.TestChannelDeprecated(context.Background()).O11yAlertmanagertypesReceiver(o11yAlertmanagertypesReceiver).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.TestChannelDeprecated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestChannelDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yAlertmanagertypesReceiver** | [**O11yAlertmanagertypesReceiver**](O11yAlertmanagertypesReceiver.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRule

> O11yO11yTestRuleOut TestRule(ctx).Body(body).Execute()

Fires a test notification for a rule definition without saving it, answering with how many series would alert.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := interface{}(987) // interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.TestRule(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.TestRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestRule`: O11yO11yTestRuleOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.TestRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **interface{}** |  | 

### Return type

[**O11yO11yTestRuleOut**](O11yO11yTestRuleOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRuleNotification

> O11yO11yTestNotificationOut TestRuleNotification(ctx).Body(body).Execute()

Fires a test notification for the posted rule definition and answers with how many series alerted and a status message.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := interface{}(987) // interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.TestRuleNotification(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.TestRuleNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestRuleNotification`: O11yO11yTestNotificationOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.TestRuleNotification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestRuleNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **interface{}** |  | 

### Return type

[**O11yO11yTestNotificationOut**](O11yO11yTestNotificationOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UninstallIntegration

> O11yO11yIntegrationAck UninstallIntegration(ctx).O11yUninstallIntegrationRequest(o11yUninstallIntegrationRequest).Execute()

Removes an integration from the caller's org by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yUninstallIntegrationRequest := *openapiclient.NewO11yUninstallIntegrationRequest() // O11yUninstallIntegrationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.UninstallIntegration(context.Background()).O11yUninstallIntegrationRequest(o11yUninstallIntegrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UninstallIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UninstallIntegration`: O11yO11yIntegrationAck
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.UninstallIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUninstallIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yUninstallIntegrationRequest** | [**O11yUninstallIntegrationRequest**](O11yUninstallIntegrationRequest.md) |  | 

### Return type

[**O11yO11yIntegrationAck**](O11yO11yIntegrationAck.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlockDashboardV2

> UnlockDashboardV2(ctx, id).Execute()

Unlocks a v2-shape dashboard.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the resource id from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.UnlockDashboardV2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UnlockDashboardV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the resource id from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlockDashboardV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnpinDashboardV2

> UnpinDashboardV2(ctx, id).Execute()

Removes the caller's pin for a dashboard.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the resource id from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.UnpinDashboardV2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UnpinDashboardV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the resource id from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnpinDashboardV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccount

> UpdateAccount(ctx, cloudProvider, id).O11yO11yUpdateAccountIn(o11yO11yUpdateAccountIn).Execute()

Changes a connected account's configuration for the given provider, by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	cloudProvider := "cloudProvider_example" // string | 
	id := "id_example" // string | 
	o11yO11yUpdateAccountIn := *openapiclient.NewO11yO11yUpdateAccountIn() // O11yO11yUpdateAccountIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.UpdateAccount(context.Background(), cloudProvider, id).O11yO11yUpdateAccountIn(o11yO11yUpdateAccountIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **o11yO11yUpdateAccountIn** | [**O11yO11yUpdateAccountIn**](O11yO11yUpdateAccountIn.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthDomain

> UpdateAuthDomain(ctx, id).O11yO11yUpdatableAuthDomain(o11yO11yUpdatableAuthDomain).Execute()

Replaces one auth domain's SSO configuration, by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 
	o11yO11yUpdatableAuthDomain := *openapiclient.NewO11yO11yUpdatableAuthDomain() // O11yO11yUpdatableAuthDomain | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.UpdateAuthDomain(context.Background(), id).O11yO11yUpdatableAuthDomain(o11yO11yUpdatableAuthDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateAuthDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yUpdatableAuthDomain** | [**O11yO11yUpdatableAuthDomain**](O11yO11yUpdatableAuthDomain.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChannelByID

> UpdateChannelByID(ctx, id).O11yO11yChannelUpdateIn(o11yO11yChannelUpdateIn).Execute()

Replaces a notification channel's receiver, by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 
	o11yO11yChannelUpdateIn := *openapiclient.NewO11yO11yChannelUpdateIn() // O11yO11yChannelUpdateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.UpdateChannelByID(context.Background(), id).O11yO11yChannelUpdateIn(o11yO11yChannelUpdateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateChannelByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChannelByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yChannelUpdateIn** | [**O11yO11yChannelUpdateIn**](O11yO11yChannelUpdateIn.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDashboardV2

> O11yO11yDashboardOut UpdateDashboardV2(ctx, id).O11yO11yDashboardUpdateIn(o11yO11yDashboardUpdateIn).Execute()

Updates a v2-shape dashboard's metadata, spec and tag set.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the dashboard id from the path.
	o11yO11yDashboardUpdateIn := *openapiclient.NewO11yO11yDashboardUpdateIn() // O11yO11yDashboardUpdateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.UpdateDashboardV2(context.Background(), id).O11yO11yDashboardUpdateIn(o11yO11yDashboardUpdateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateDashboardV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDashboardV2`: O11yO11yDashboardOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.UpdateDashboardV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the dashboard id from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDashboardV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yDashboardUpdateIn** | [**O11yO11yDashboardUpdateIn**](O11yO11yDashboardUpdateIn.md) |  | 

### Return type

[**O11yO11yDashboardOut**](O11yO11yDashboardOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDashboardView

> O11yO11yDashboardViewOut UpdateDashboardView(ctx, id).O11yO11yDashboardViewUpdateIn(o11yO11yDashboardViewUpdateIn).Execute()

Replaces a saved view's name and data.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the saved view id from the path.
	o11yO11yDashboardViewUpdateIn := *openapiclient.NewO11yO11yDashboardViewUpdateIn() // O11yO11yDashboardViewUpdateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.UpdateDashboardView(context.Background(), id).O11yO11yDashboardViewUpdateIn(o11yO11yDashboardViewUpdateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateDashboardView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDashboardView`: O11yO11yDashboardViewOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.UpdateDashboardView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the saved view id from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDashboardViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yDashboardViewUpdateIn** | [**O11yO11yDashboardViewUpdateIn**](O11yO11yDashboardViewUpdateIn.md) |  | 

### Return type

[**O11yO11yDashboardViewOut**](O11yO11yDashboardViewOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDowntimeScheduleByID

> UpdateDowntimeScheduleByID(ctx, id).O11yO11yDowntimeUpdateIn(o11yO11yDowntimeUpdateIn).Execute()

Replaces a planned maintenance window, by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 
	o11yO11yDowntimeUpdateIn := *openapiclient.NewO11yO11yDowntimeUpdateIn() // O11yO11yDowntimeUpdateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.UpdateDowntimeScheduleByID(context.Background(), id).O11yO11yDowntimeUpdateIn(o11yO11yDowntimeUpdateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateDowntimeScheduleByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDowntimeScheduleByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yDowntimeUpdateIn** | [**O11yO11yDowntimeUpdateIn**](O11yO11yDowntimeUpdateIn.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIngestionKey

> UpdateIngestionKey(ctx, keyId).O11yO11yUpdateIngestionKeyIn(o11yO11yUpdateIngestionKeyIn).Execute()

Changes an ingestion key, by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	keyId := "keyId_example" // string | 
	o11yO11yUpdateIngestionKeyIn := *openapiclient.NewO11yO11yUpdateIngestionKeyIn() // O11yO11yUpdateIngestionKeyIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.UpdateIngestionKey(context.Background(), keyId).O11yO11yUpdateIngestionKeyIn(o11yO11yUpdateIngestionKeyIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateIngestionKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIngestionKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yUpdateIngestionKeyIn** | [**O11yO11yUpdateIngestionKeyIn**](O11yO11yUpdateIngestionKeyIn.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIngestionKeyLimit

> UpdateIngestionKeyLimit(ctx, limitId).O11yO11yUpdateLimitIn(o11yO11yUpdateLimitIn).Execute()

Changes an ingestion key limit, by limit id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	limitId := "limitId_example" // string | 
	o11yO11yUpdateLimitIn := *openapiclient.NewO11yO11yUpdateLimitIn() // O11yO11yUpdateLimitIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.UpdateIngestionKeyLimit(context.Background(), limitId).O11yO11yUpdateLimitIn(o11yO11yUpdateLimitIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateIngestionKeyLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**limitId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIngestionKeyLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yUpdateLimitIn** | [**O11yO11yUpdateLimitIn**](O11yO11yUpdateLimitIn.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMetricMetadata

> O11yO11yMetricAckOut UpdateMetricMetadata(ctx).O11yO11yMetricMetadataSaveIn(o11yO11yMetricMetadataSaveIn).Execute()

Updates one metric's metadata — description, type, unit, temporality, monotonicity — and answers with the bare success envelope.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yMetricMetadataSaveIn := *openapiclient.NewO11yO11yMetricMetadataSaveIn("MetricName_example") // O11yO11yMetricMetadataSaveIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.UpdateMetricMetadata(context.Background()).O11yO11yMetricMetadataSaveIn(o11yO11yMetricMetadataSaveIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateMetricMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMetricMetadata`: O11yO11yMetricAckOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.UpdateMetricMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMetricMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yMetricMetadataSaveIn** | [**O11yO11yMetricMetadataSaveIn**](O11yO11yMetricMetadataSaveIn.md) |  | 

### Return type

[**O11yO11yMetricAckOut**](O11yO11yMetricAckOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMetricReductionRuleByID

> O11yO11yReductionRuleOut UpdateMetricReductionRuleByID(ctx, id).O11yO11yReductionRuleSaveIn(o11yO11yReductionRuleSaveIn).Execute()

Updates the match type and labels of a volume-control rule by its id; the metric name is immutable.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the rule's id.
	o11yO11yReductionRuleSaveIn := *openapiclient.NewO11yO11yReductionRuleSaveIn("Id_example", []string{"Labels_example"}, "MatchType_example") // O11yO11yReductionRuleSaveIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.UpdateMetricReductionRuleByID(context.Background(), id).O11yO11yReductionRuleSaveIn(o11yO11yReductionRuleSaveIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateMetricReductionRuleByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMetricReductionRuleByID`: O11yO11yReductionRuleOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.UpdateMetricReductionRuleByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the rule&#39;s id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMetricReductionRuleByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yReductionRuleSaveIn** | [**O11yO11yReductionRuleSaveIn**](O11yO11yReductionRuleSaveIn.md) |  | 

### Return type

[**O11yO11yReductionRuleOut**](O11yO11yReductionRuleOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMyOrganization

> UpdateMyOrganization(ctx).O11yO11yOrganization(o11yO11yOrganization).Execute()

Rewrites the caller's own organization record — display name, name, alias — always addressed as \"me\", never by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yOrganization := *openapiclient.NewO11yO11yOrganization() // O11yO11yOrganization | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.UpdateMyOrganization(context.Background()).O11yO11yOrganization(o11yO11yOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateMyOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMyOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yOrganization** | [**O11yO11yOrganization**](O11yO11yOrganization.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMyPassword

> UpdateMyPassword(ctx).O11yO11yChangePasswordIn(o11yO11yChangePasswordIn).Execute()

Replaces the calling user's password, refusing when the old one does not match.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yChangePasswordIn := *openapiclient.NewO11yO11yChangePasswordIn() // O11yO11yChangePasswordIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.UpdateMyPassword(context.Background()).O11yO11yChangePasswordIn(o11yO11yChangePasswordIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateMyPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMyPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yChangePasswordIn** | [**O11yO11yChangePasswordIn**](O11yO11yChangePasswordIn.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMyServiceAccount

> UpdateMyServiceAccount(ctx).O11yO11yMyServiceAccountUpdateIn(o11yO11yMyServiceAccountUpdateIn).Execute()

Renames the calling service account.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yMyServiceAccountUpdateIn := *openapiclient.NewO11yO11yMyServiceAccountUpdateIn() // O11yO11yMyServiceAccountUpdateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.UpdateMyServiceAccount(context.Background()).O11yO11yMyServiceAccountUpdateIn(o11yO11yMyServiceAccountUpdateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateMyServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMyServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yMyServiceAccountUpdateIn** | [**O11yO11yMyServiceAccountUpdateIn**](O11yO11yMyServiceAccountUpdateIn.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMyUserV2

> UpdateMyUserV2(ctx).O11yO11yUpdatableUser(o11yO11yUpdatableUser).Execute()

Renames the calling user.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yUpdatableUser := *openapiclient.NewO11yO11yUpdatableUser() // O11yO11yUpdatableUser | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.UpdateMyUserV2(context.Background()).O11yO11yUpdatableUser(o11yO11yUpdatableUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateMyUserV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMyUserV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yUpdatableUser** | [**O11yO11yUpdatableUser**](O11yO11yUpdatableUser.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgPreference

> UpdateOrgPreference(ctx, name).O11yO11yUpdatablePreference(o11yO11yUpdatablePreference).Execute()

Sets one org-scoped preference, by name.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	name := "name_example" // string | 
	o11yO11yUpdatablePreference := *openapiclient.NewO11yO11yUpdatablePreference() // O11yO11yUpdatablePreference | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.UpdateOrgPreference(context.Background(), name).O11yO11yUpdatablePreference(o11yO11yUpdatablePreference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateOrgPreference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgPreferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yUpdatablePreference** | [**O11yO11yUpdatablePreference**](O11yO11yUpdatablePreference.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePublicDashboard

> UpdatePublicDashboard(ctx, id).O11yO11yPublicDashboardWriteIn(o11yO11yPublicDashboardWriteIn).Execute()

Updates the public-sharing config for a dashboard.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | ID is the dashboard id from the path.
	o11yO11yPublicDashboardWriteIn := *openapiclient.NewO11yO11yPublicDashboardWriteIn() // O11yO11yPublicDashboardWriteIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.UpdatePublicDashboard(context.Background(), id).O11yO11yPublicDashboardWriteIn(o11yO11yPublicDashboardWriteIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdatePublicDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the dashboard id from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePublicDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yPublicDashboardWriteIn** | [**O11yO11yPublicDashboardWriteIn**](O11yO11yPublicDashboardWriteIn.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateQuickFilters

> UpdateQuickFilters(ctx).O11yO11yUpdatableQuickFilters(o11yO11yUpdatableQuickFilters).Execute()

Replaces the org's quick filters for one signal with the attribute list given.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yUpdatableQuickFilters := *openapiclient.NewO11yO11yUpdatableQuickFilters() // O11yO11yUpdatableQuickFilters | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.UpdateQuickFilters(context.Background()).O11yO11yUpdatableQuickFilters(o11yO11yUpdatableQuickFilters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateQuickFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateQuickFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yUpdatableQuickFilters** | [**O11yO11yUpdatableQuickFilters**](O11yO11yUpdatableQuickFilters.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> UpdateRole(ctx, id).O11yO11yRoleUpdateIn(o11yO11yRoleUpdateIn).Execute()

Replaces a custom role's description and transaction groups.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 
	o11yO11yRoleUpdateIn := *openapiclient.NewO11yO11yRoleUpdateIn() // O11yO11yRoleUpdateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.UpdateRole(context.Background(), id).O11yO11yRoleUpdateIn(o11yO11yRoleUpdateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yRoleUpdateIn** | [**O11yO11yRoleUpdateIn**](O11yO11yRoleUpdateIn.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoutePolicy

> O11yO11yRoutePolicyOut UpdateRoutePolicy(ctx, id).O11yO11yRoutePolicyUpdateIn(o11yO11yRoutePolicyUpdateIn).Execute()

Replaces a route policy, by id, answering with the stored policy.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 
	o11yO11yRoutePolicyUpdateIn := *openapiclient.NewO11yO11yRoutePolicyUpdateIn() // O11yO11yRoutePolicyUpdateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.UpdateRoutePolicy(context.Background(), id).O11yO11yRoutePolicyUpdateIn(o11yO11yRoutePolicyUpdateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateRoutePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoutePolicy`: O11yO11yRoutePolicyOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.UpdateRoutePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoutePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yRoutePolicyUpdateIn** | [**O11yO11yRoutePolicyUpdateIn**](O11yO11yRoutePolicyUpdateIn.md) |  | 

### Return type

[**O11yO11yRoutePolicyOut**](O11yO11yRoutePolicyOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRuleByID

> UpdateRuleByID(ctx, id).Body(body).Execute()

Replaces an alert rule's definition, by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 
	body := interface{}(987) // interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.UpdateRuleByID(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateRuleByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRuleByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> UpdateService(ctx, cloudProvider, id, serviceId).O11yO11yUpdateServiceIn(o11yO11yUpdateServiceIn).Execute()

Changes a service's configuration for one connected account of the given provider, by account id and service id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	cloudProvider := "cloudProvider_example" // string | 
	id := "id_example" // string | 
	serviceId := "serviceId_example" // string | 
	o11yO11yUpdateServiceIn := *openapiclient.NewO11yO11yUpdateServiceIn() // O11yO11yUpdateServiceIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.UpdateService(context.Background(), cloudProvider, id, serviceId).O11yO11yUpdateServiceIn(o11yO11yUpdateServiceIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** |  | 
**id** | **string** |  | 
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **o11yO11yUpdateServiceIn** | [**O11yO11yUpdateServiceIn**](O11yO11yUpdateServiceIn.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceAccount

> UpdateServiceAccount(ctx, id).O11yO11yServiceAccountUpdateIn(o11yO11yServiceAccountUpdateIn).Execute()

Renames a service account.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 
	o11yO11yServiceAccountUpdateIn := *openapiclient.NewO11yO11yServiceAccountUpdateIn() // O11yO11yServiceAccountUpdateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.UpdateServiceAccount(context.Background(), id).O11yO11yServiceAccountUpdateIn(o11yO11yServiceAccountUpdateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yServiceAccountUpdateIn** | [**O11yO11yServiceAccountUpdateIn**](O11yO11yServiceAccountUpdateIn.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceAccountKey

> UpdateServiceAccountKey(ctx, id, fid).O11yO11yAPIKeyUpdateIn(o11yO11yAPIKeyUpdateIn).Execute()

Renames an API key or moves its expiry.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 
	fid := "fid_example" // string | 
	o11yO11yAPIKeyUpdateIn := *openapiclient.NewO11yO11yAPIKeyUpdateIn() // O11yO11yAPIKeyUpdateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.UpdateServiceAccountKey(context.Background(), id, fid).O11yO11yAPIKeyUpdateIn(o11yO11yAPIKeyUpdateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateServiceAccountKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**fid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceAccountKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **o11yO11yAPIKeyUpdateIn** | [**O11yO11yAPIKeyUpdateIn**](O11yO11yAPIKeyUpdateIn.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSpanMapper

> UpdateSpanMapper(ctx, groupId, mapperId).O11yO11ySpanMapperUpdateIn(o11yO11ySpanMapperUpdateIn).Execute()

Changes a mapper's field context, config or enabled state.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	groupId := "groupId_example" // string | 
	mapperId := "mapperId_example" // string | 
	o11yO11ySpanMapperUpdateIn := *openapiclient.NewO11yO11ySpanMapperUpdateIn() // O11yO11ySpanMapperUpdateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.UpdateSpanMapper(context.Background(), groupId, mapperId).O11yO11ySpanMapperUpdateIn(o11yO11ySpanMapperUpdateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateSpanMapper``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 
**mapperId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpanMapperRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **o11yO11ySpanMapperUpdateIn** | [**O11yO11ySpanMapperUpdateIn**](O11yO11ySpanMapperUpdateIn.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSpanMapperGroup

> UpdateSpanMapperGroup(ctx, groupId).O11yO11ySpanMapperGroupUpdateIn(o11yO11ySpanMapperGroupUpdateIn).Execute()

Changes a group's name, condition or enabled state.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	groupId := "groupId_example" // string | 
	o11yO11ySpanMapperGroupUpdateIn := *openapiclient.NewO11yO11ySpanMapperGroupUpdateIn() // O11yO11ySpanMapperGroupUpdateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.UpdateSpanMapperGroup(context.Background(), groupId).O11yO11ySpanMapperGroupUpdateIn(o11yO11ySpanMapperGroupUpdateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateSpanMapperGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpanMapperGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11ySpanMapperGroupUpdateIn** | [**O11yO11ySpanMapperGroupUpdateIn**](O11yO11ySpanMapperGroupUpdateIn.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTraceField

> O11yO11yFieldSetting UpdateTraceField(ctx).O11yO11yFieldSetting(o11yO11yFieldSetting).Execute()

Changes how one span field is stored — selects or deselects it as a materialized column and tunes its index — and echoes the setting back.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yFieldSetting := *openapiclient.NewO11yO11yFieldSetting("DataType_example", "Name_example", "Type_example") // O11yO11yFieldSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.UpdateTraceField(context.Background()).O11yO11yFieldSetting(o11yO11yFieldSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateTraceField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTraceField`: O11yO11yFieldSetting
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.UpdateTraceField`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTraceFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yFieldSetting** | [**O11yO11yFieldSetting**](O11yO11yFieldSetting.md) |  | 

### Return type

[**O11yO11yFieldSetting**](O11yO11yFieldSetting.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTraceFunnel

> O11yO11yFunnelOut UpdateTraceFunnel(ctx, funnelId).O11yO11yFunnelUpdateIn(o11yO11yFunnelUpdateIn).Execute()

Renames a funnel or rewrites its description, answering the funnel as it now stands.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	funnelId := "funnelId_example" // string | 
	o11yO11yFunnelUpdateIn := *openapiclient.NewO11yO11yFunnelUpdateIn() // O11yO11yFunnelUpdateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.UpdateTraceFunnel(context.Background(), funnelId).O11yO11yFunnelUpdateIn(o11yO11yFunnelUpdateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateTraceFunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTraceFunnel`: O11yO11yFunnelOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.UpdateTraceFunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**funnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTraceFunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yFunnelUpdateIn** | [**O11yO11yFunnelUpdateIn**](O11yO11yFunnelUpdateIn.md) |  | 

### Return type

[**O11yO11yFunnelOut**](O11yO11yFunnelOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTraceFunnelSteps

> O11yO11yFunnelOut UpdateTraceFunnelSteps(ctx).O11yO11yFunnelStepsUpdateIn(o11yO11yFunnelStepsUpdateIn).Execute()

Replaces a funnel's steps — the funnel is named in the body rather than the path — and answers the funnel as it now stands.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yFunnelStepsUpdateIn := *openapiclient.NewO11yO11yFunnelStepsUpdateIn() // O11yO11yFunnelStepsUpdateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.UpdateTraceFunnelSteps(context.Background()).O11yO11yFunnelStepsUpdateIn(o11yO11yFunnelStepsUpdateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateTraceFunnelSteps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTraceFunnelSteps`: O11yO11yFunnelOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.UpdateTraceFunnelSteps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTraceFunnelStepsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yFunnelStepsUpdateIn** | [**O11yO11yFunnelStepsUpdateIn**](O11yO11yFunnelStepsUpdateIn.md) |  | 

### Return type

[**O11yO11yFunnelOut**](O11yO11yFunnelOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> UpdateUser(ctx, id).O11yO11yUserUpdate(o11yO11yUserUpdate).Execute()

Renames one org member, by user id — someone else, never the caller, who renames themselves through updateMyUser.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 
	o11yO11yUserUpdate := *openapiclient.NewO11yO11yUserUpdate() // O11yO11yUserUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.UpdateUser(context.Background(), id).O11yO11yUserUpdate(o11yO11yUserUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yUserUpdate** | [**O11yO11yUserUpdate**](O11yO11yUserUpdate.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserDeprecated

> O11yO11yDeprecatedUserOut UpdateUserDeprecated(ctx, id).O11yO11yDeprecatedUserUpdate(o11yO11yDeprecatedUserUpdate).Execute()

Renames one org member and may move their legacy role, answering with the updated record.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	id := "id_example" // string | 
	o11yO11yDeprecatedUserUpdate := *openapiclient.NewO11yO11yDeprecatedUserUpdate() // O11yO11yDeprecatedUserUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.UpdateUserDeprecated(context.Background(), id).O11yO11yDeprecatedUserUpdate(o11yO11yDeprecatedUserUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateUserDeprecated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserDeprecated`: O11yO11yDeprecatedUserOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.UpdateUserDeprecated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yDeprecatedUserUpdate** | [**O11yO11yDeprecatedUserUpdate**](O11yO11yDeprecatedUserUpdate.md) |  | 

### Return type

[**O11yO11yDeprecatedUserOut**](O11yO11yDeprecatedUserOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPreference

> UpdateUserPreference(ctx, name).O11yO11yUpdatablePreference(o11yO11yUpdatablePreference).Execute()

Sets one preference of the calling user, by name.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	name := "name_example" // string | 
	o11yO11yUpdatablePreference := *openapiclient.NewO11yO11yUpdatablePreference() // O11yO11yUpdatablePreference | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.UpdateUserPreference(context.Background(), name).O11yO11yUpdatablePreference(o11yO11yUpdatablePreference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.UpdateUserPreference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPreferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yUpdatablePreference** | [**O11yO11yUpdatablePreference**](O11yO11yUpdatablePreference.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateDraftFunnelTraces

> O11yO11yFunnelRowsOut ValidateDraftFunnelTraces(ctx).O11yO11yDraftFunnelIn(o11yO11yDraftFunnelIn).Execute()

Lists the traces that match a funnel described inline — the builder's \"try this\" before anything is saved.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yDraftFunnelIn := *openapiclient.NewO11yO11yDraftFunnelIn() // O11yO11yDraftFunnelIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ValidateDraftFunnelTraces(context.Background()).O11yO11yDraftFunnelIn(o11yO11yDraftFunnelIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ValidateDraftFunnelTraces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateDraftFunnelTraces`: O11yO11yFunnelRowsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ValidateDraftFunnelTraces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateDraftFunnelTracesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yDraftFunnelIn** | [**O11yO11yDraftFunnelIn**](O11yO11yDraftFunnelIn.md) |  | 

### Return type

[**O11yO11yFunnelRowsOut**](O11yO11yFunnelRowsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateTraceFunnelTraces

> O11yO11yFunnelRowsOut ValidateTraceFunnelTraces(ctx, funnelId).O11yO11yFunnelWindowIn(o11yO11yFunnelWindowIn).Execute()

Lists the traces that match a saved funnel over a window — the read that answers \"is this funnel finding anything at all\".



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	funnelId := "funnelId_example" // string | 
	o11yO11yFunnelWindowIn := *openapiclient.NewO11yO11yFunnelWindowIn() // O11yO11yFunnelWindowIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.ValidateTraceFunnelTraces(context.Background(), funnelId).O11yO11yFunnelWindowIn(o11yO11yFunnelWindowIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.ValidateTraceFunnelTraces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateTraceFunnelTraces`: O11yO11yFunnelRowsOut
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.ValidateTraceFunnelTraces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**funnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateTraceFunnelTracesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **o11yO11yFunnelWindowIn** | [**O11yO11yFunnelWindowIn**](O11yO11yFunnelWindowIn.md) |  | 

### Return type

[**O11yO11yFunnelRowsOut**](O11yO11yFunnelRowsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyResetPasswordToken

> VerifyResetPasswordToken(ctx).O11yO11yResetTokenRef(o11yO11yResetTokenRef).Execute()

Checks that a reset-password token exists and has not expired, without consuming it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	o11yO11yResetTokenRef := *openapiclient.NewO11yO11yResetTokenRef() // O11yO11yResetTokenRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.O11yAPI.VerifyResetPasswordToken(context.Background()).O11yO11yResetTokenRef(o11yO11yResetTokenRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.VerifyResetPasswordToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyResetPasswordTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yO11yResetTokenRef** | [**O11yO11yResetTokenRef**](O11yO11yResetTokenRef.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

