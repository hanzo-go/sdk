# \AiAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiMCPTools**](AiAPI.md#AiMCPTools) | **Get** /v1/ai/mcp/tools | Tools reports what THIS PROCESS&#39;s MCP door carries: how many tools its own registry projects, optionally their names, and which subsystems this process composed.
[**DeleteAiArticlesByOwnerByName**](AiAPI.md#DeleteAiArticlesByOwnerByName) | **Delete** /v1/ai/articles/{owner}/{name} | Delete a article
[**DeleteAiAssetsByOwnerByName**](AiAPI.md#DeleteAiAssetsByOwnerByName) | **Delete** /v1/ai/assets/{owner}/{name} | Delete a asset
[**DeleteAiChatsByOwnerByName**](AiAPI.md#DeleteAiChatsByOwnerByName) | **Delete** /v1/ai/chats/{owner}/{name} | Delete a chat
[**DeleteAiConnectionsByProvider**](AiAPI.md#DeleteAiConnectionsByProvider) | **Delete** /v1/ai/connections/{provider} | Disconnects a third-party AI account: it deactivates the org&#39;s row so completion resolution falls back to the global Hanzo account (no BYO), and best-effort tombstones the sealed secret.
[**DeleteAiDeploymentsByOwnerByName**](AiAPI.md#DeleteAiDeploymentsByOwnerByName) | **Delete** /v1/ai/deployments/{owner}/{name} | Delete a application
[**DeleteAiFilesByOwnerByName**](AiAPI.md#DeleteAiFilesByOwnerByName) | **Delete** /v1/ai/files/{owner}/{name} | Delete a file
[**DeleteAiFormsByOwnerByName**](AiAPI.md#DeleteAiFormsByOwnerByName) | **Delete** /v1/ai/forms/{owner}/{name} | Delete a form
[**DeleteAiGraphsByOwnerByName**](AiAPI.md#DeleteAiGraphsByOwnerByName) | **Delete** /v1/ai/graphs/{owner}/{name} | Delete a graph
[**DeleteAiMessagesByOwnerByName**](AiAPI.md#DeleteAiMessagesByOwnerByName) | **Delete** /v1/ai/messages/{owner}/{name} | Delete a message
[**DeleteAiMessagesWelcome**](AiAPI.md#DeleteAiMessagesWelcome) | **Delete** /v1/ai/messages/welcome | Welcome (message)
[**DeleteAiNodesByOwnerByName**](AiAPI.md#DeleteAiNodesByOwnerByName) | **Delete** /v1/ai/nodes/{owner}/{name} | Delete a node
[**DeleteAiProvidersByOwnerByName**](AiAPI.md#DeleteAiProvidersByOwnerByName) | **Delete** /v1/ai/providers/{owner}/{name} | Delete a provider
[**DeleteAiRecordsByOwnerByName**](AiAPI.md#DeleteAiRecordsByOwnerByName) | **Delete** /v1/ai/records/{owner}/{name} | Delete a record
[**DeleteAiRemoteConnectionsByOwnerByName**](AiAPI.md#DeleteAiRemoteConnectionsByOwnerByName) | **Delete** /v1/ai/remote-connections/{owner}/{name} | Delete a connection
[**DeleteAiRoutesByOwnerByName**](AiAPI.md#DeleteAiRoutesByOwnerByName) | **Delete** /v1/ai/routes/{owner}/{name} | Delete a model-route
[**DeleteAiScalesByOwnerByName**](AiAPI.md#DeleteAiScalesByOwnerByName) | **Delete** /v1/ai/scales/{owner}/{name} | Delete a scale
[**DeleteAiScansByOwnerByName**](AiAPI.md#DeleteAiScansByOwnerByName) | **Delete** /v1/ai/scans/{owner}/{name} | Delete a scan
[**DeleteAiSigninSessionsByOwnerByName**](AiAPI.md#DeleteAiSigninSessionsByOwnerByName) | **Delete** /v1/ai/signin-sessions/{owner}/{name} | Delete a session
[**DeleteAiStoresByOwnerByName**](AiAPI.md#DeleteAiStoresByOwnerByName) | **Delete** /v1/ai/stores/{owner}/{name} | Delete a store
[**DeleteAiTasksByOwnerByName**](AiAPI.md#DeleteAiTasksByOwnerByName) | **Delete** /v1/ai/tasks/{owner}/{name} | Delete a task
[**DeleteAiTemplatesByOwnerByName**](AiAPI.md#DeleteAiTemplatesByOwnerByName) | **Delete** /v1/ai/templates/{owner}/{name} | Delete a template
[**DeleteAiTreeFilesByOwnerByName**](AiAPI.md#DeleteAiTreeFilesByOwnerByName) | **Delete** /v1/ai/tree-files/{owner}/{name} | Delete a tree-file
[**DeleteAiVectorsAll**](AiAPI.md#DeleteAiVectorsAll) | **Delete** /v1/ai/vectors/all | All (vector)
[**DeleteAiVectorsByOwnerByName**](AiAPI.md#DeleteAiVectorsByOwnerByName) | **Delete** /v1/ai/vectors/{owner}/{name} | Delete a vector
[**DeleteAiVideosByOwnerByName**](AiAPI.md#DeleteAiVideosByOwnerByName) | **Delete** /v1/ai/videos/{owner}/{name} | Delete a video
[**DeleteAiWorkflowsByOwnerByName**](AiAPI.md#DeleteAiWorkflowsByOwnerByName) | **Delete** /v1/ai/workflows/{owner}/{name} | Delete a workflow
[**GetAiAccount**](AiAPI.md#GetAiAccount) | **Get** /v1/ai/account | Account
[**GetAiActivities**](AiAPI.md#GetAiActivities) | **Get** /v1/ai/activities | List activities
[**GetAiAnswer**](AiAPI.md#GetAiAnswer) | **Get** /v1/ai/answer | Answer
[**GetAiArticles**](AiAPI.md#GetAiArticles) | **Get** /v1/ai/articles | List articles
[**GetAiArticlesByOwnerByName**](AiAPI.md#GetAiArticlesByOwnerByName) | **Get** /v1/ai/articles/{owner}/{name} | Retrieve a article
[**GetAiArticlesGlobal**](AiAPI.md#GetAiArticlesGlobal) | **Get** /v1/ai/articles/global | List articles across tenants
[**GetAiAssets**](AiAPI.md#GetAiAssets) | **Get** /v1/ai/assets | List assets
[**GetAiAssetsByOwnerByName**](AiAPI.md#GetAiAssetsByOwnerByName) | **Get** /v1/ai/assets/{owner}/{name} | Retrieve a asset
[**GetAiChats**](AiAPI.md#GetAiChats) | **Get** /v1/ai/chats | List chats
[**GetAiChatsByOwnerByName**](AiAPI.md#GetAiChatsByOwnerByName) | **Get** /v1/ai/chats/{owner}/{name} | Retrieve a chat
[**GetAiChatsGlobal**](AiAPI.md#GetAiChatsGlobal) | **Get** /v1/ai/chats/global | List chats across tenants
[**GetAiConnections**](AiAPI.md#GetAiConnections) | **Get** /v1/ai/connections | Lists the org&#39;s connectable AI accounts and whether each is currently connected.
[**GetAiConnectionsByProviderAuthorize**](AiAPI.md#GetAiConnectionsByProviderAuthorize) | **Get** /v1/ai/connections/{provider}/authorize | Begins an OAuth connection for the caller&#39;s org: it binds the org into a signed state and sends the caller to the provider&#39;s authorize URL.
[**GetAiConnectionsByProviderCallback**](AiAPI.md#GetAiConnectionsByProviderCallback) | **Get** /v1/ai/connections/{provider}/callback | Completes OAuth: the org is recovered from the SIGNED state (not a header), the code is exchanged for a token, the token is SEALED into KMS (never the row/logs) through the same path as a BYOK key, and the org&#39;s provider row is upserted to \&quot;connected\&quot;.
[**GetAiConnectionsByProviderUsage**](AiAPI.md#GetAiConnectionsByProviderUsage) | **Get** /v1/ai/connections/{provider}/usage | Imports the caller org&#39;s usage for a connected third-party account.
[**GetAiDashboardsAgents**](AiAPI.md#GetAiDashboardsAgents) | **Get** /v1/ai/dashboards/agents | Dashboards Agents
[**GetAiDashboardsVm**](AiAPI.md#GetAiDashboardsVm) | **Get** /v1/ai/dashboards/vm | Dashboards Vm
[**GetAiDeployments**](AiAPI.md#GetAiDeployments) | **Get** /v1/ai/deployments | List deployments
[**GetAiDeploymentsByOwnerByName**](AiAPI.md#GetAiDeploymentsByOwnerByName) | **Get** /v1/ai/deployments/{owner}/{name} | Retrieve a application
[**GetAiFiles**](AiAPI.md#GetAiFiles) | **Get** /v1/ai/files | List files
[**GetAiFilesActive**](AiAPI.md#GetAiFilesActive) | **Get** /v1/ai/files/active | Active (file)
[**GetAiFilesByOwnerByName**](AiAPI.md#GetAiFilesByOwnerByName) | **Get** /v1/ai/files/{owner}/{name} | Retrieve a file
[**GetAiFilesGlobal**](AiAPI.md#GetAiFilesGlobal) | **Get** /v1/ai/files/global | List files across tenants
[**GetAiForms**](AiAPI.md#GetAiForms) | **Get** /v1/ai/forms | List forms
[**GetAiFormsByOwnerByName**](AiAPI.md#GetAiFormsByOwnerByName) | **Get** /v1/ai/forms/{owner}/{name} | Retrieve a form
[**GetAiFormsData**](AiAPI.md#GetAiFormsData) | **Get** /v1/ai/forms/data | Data (form)
[**GetAiFormsGlobal**](AiAPI.md#GetAiFormsGlobal) | **Get** /v1/ai/forms/global | List forms across tenants
[**GetAiGraphs**](AiAPI.md#GetAiGraphs) | **Get** /v1/ai/graphs | List graphs
[**GetAiGraphsByOwnerByName**](AiAPI.md#GetAiGraphsByOwnerByName) | **Get** /v1/ai/graphs/{owner}/{name} | Retrieve a graph
[**GetAiGraphsGlobal**](AiAPI.md#GetAiGraphsGlobal) | **Get** /v1/ai/graphs/global | List graphs across tenants
[**GetAiK8sStatus**](AiAPI.md#GetAiK8sStatus) | **Get** /v1/ai/k8s-status | K8s Status
[**GetAiMessages**](AiAPI.md#GetAiMessages) | **Get** /v1/ai/messages | List messages
[**GetAiMessagesByOwnerByName**](AiAPI.md#GetAiMessagesByOwnerByName) | **Get** /v1/ai/messages/{owner}/{name} | Retrieve a message
[**GetAiMessagesByOwnerByNameAnswer**](AiAPI.md#GetAiMessagesByOwnerByNameAnswer) | **Get** /v1/ai/messages/{owner}/{name}/answer | Answer (message)
[**GetAiMessagesGlobal**](AiAPI.md#GetAiMessagesGlobal) | **Get** /v1/ai/messages/global | List messages across tenants
[**GetAiNodes**](AiAPI.md#GetAiNodes) | **Get** /v1/ai/nodes | List nodes
[**GetAiNodesByOwnerByName**](AiAPI.md#GetAiNodesByOwnerByName) | **Get** /v1/ai/nodes/{owner}/{name} | Retrieve a node
[**GetAiNodesByOwnerByNameTunnel**](AiAPI.md#GetAiNodesByOwnerByNameTunnel) | **Get** /v1/ai/nodes/{owner}/{name}/tunnel | Tunnel (node)
[**GetAiPrometheus**](AiAPI.md#GetAiPrometheus) | **Get** /v1/ai/prometheus | Prometheus
[**GetAiProviders**](AiAPI.md#GetAiProviders) | **Get** /v1/ai/providers | List providers
[**GetAiProvidersByOwnerByName**](AiAPI.md#GetAiProvidersByOwnerByName) | **Get** /v1/ai/providers/{owner}/{name} | Retrieve a provider
[**GetAiProvidersGlobal**](AiAPI.md#GetAiProvidersGlobal) | **Get** /v1/ai/providers/global | List providers across tenants
[**GetAiRecords**](AiAPI.md#GetAiRecords) | **Get** /v1/ai/records | List records
[**GetAiRecordsByOwnerByName**](AiAPI.md#GetAiRecordsByOwnerByName) | **Get** /v1/ai/records/{owner}/{name} | Retrieve a record
[**GetAiRecordsQuery**](AiAPI.md#GetAiRecordsQuery) | **Get** /v1/ai/records/query | Query (record)
[**GetAiRecordsQuerySecond**](AiAPI.md#GetAiRecordsQuerySecond) | **Get** /v1/ai/records/query-second | Query Second (record)
[**GetAiRemoteConnections**](AiAPI.md#GetAiRemoteConnections) | **Get** /v1/ai/remote-connections | List remote-connections
[**GetAiRemoteConnectionsByOwnerByName**](AiAPI.md#GetAiRemoteConnectionsByOwnerByName) | **Get** /v1/ai/remote-connections/{owner}/{name} | Retrieve a connection
[**GetAiRoutes**](AiAPI.md#GetAiRoutes) | **Get** /v1/ai/routes | List routes
[**GetAiRoutesByOwnerByName**](AiAPI.md#GetAiRoutesByOwnerByName) | **Get** /v1/ai/routes/{owner}/{name} | Retrieve a model-route
[**GetAiScales**](AiAPI.md#GetAiScales) | **Get** /v1/ai/scales | List scales
[**GetAiScalesByOwnerByName**](AiAPI.md#GetAiScalesByOwnerByName) | **Get** /v1/ai/scales/{owner}/{name} | Retrieve a scale
[**GetAiScalesGlobal**](AiAPI.md#GetAiScalesGlobal) | **Get** /v1/ai/scales/global | List scales across tenants
[**GetAiScalesPublic**](AiAPI.md#GetAiScalesPublic) | **Get** /v1/ai/scales/public | Public (scale)
[**GetAiScans**](AiAPI.md#GetAiScans) | **Get** /v1/ai/scans | List scans
[**GetAiScansByOwnerByName**](AiAPI.md#GetAiScansByOwnerByName) | **Get** /v1/ai/scans/{owner}/{name} | Retrieve a scan
[**GetAiSigninSessions**](AiAPI.md#GetAiSigninSessions) | **Get** /v1/ai/signin-sessions | List signin-sessions
[**GetAiSigninSessionsByOwnerByName**](AiAPI.md#GetAiSigninSessionsByOwnerByName) | **Get** /v1/ai/signin-sessions/{owner}/{name} | Retrieve a session
[**GetAiSigninSessionsDuplicated**](AiAPI.md#GetAiSigninSessionsDuplicated) | **Get** /v1/ai/signin-sessions/duplicated | Duplicated (session)
[**GetAiStores**](AiAPI.md#GetAiStores) | **Get** /v1/ai/stores | List stores
[**GetAiStoresByOwnerByName**](AiAPI.md#GetAiStoresByOwnerByName) | **Get** /v1/ai/stores/{owner}/{name} | Retrieve a store
[**GetAiStoresGlobal**](AiAPI.md#GetAiStoresGlobal) | **Get** /v1/ai/stores/global | List stores across tenants
[**GetAiStoresNames**](AiAPI.md#GetAiStoresNames) | **Get** /v1/ai/stores/names | Names (store)
[**GetAiStoresProviders**](AiAPI.md#GetAiStoresProviders) | **Get** /v1/ai/stores/providers | Providers (store)
[**GetAiSystem**](AiAPI.md#GetAiSystem) | **Get** /v1/ai/system | System
[**GetAiTasks**](AiAPI.md#GetAiTasks) | **Get** /v1/ai/tasks | List tasks
[**GetAiTasksByOwnerByName**](AiAPI.md#GetAiTasksByOwnerByName) | **Get** /v1/ai/tasks/{owner}/{name} | Retrieve a task
[**GetAiTasksGlobal**](AiAPI.md#GetAiTasksGlobal) | **Get** /v1/ai/tasks/global | List tasks across tenants
[**GetAiTemplates**](AiAPI.md#GetAiTemplates) | **Get** /v1/ai/templates | List templates
[**GetAiTemplatesByOwnerByName**](AiAPI.md#GetAiTemplatesByOwnerByName) | **Get** /v1/ai/templates/{owner}/{name} | Retrieve a template
[**GetAiTrainingContribution**](AiAPI.md#GetAiTrainingContribution) | **Get** /v1/ai/training-contribution | Training Contribution
[**GetAiUsages**](AiAPI.md#GetAiUsages) | **Get** /v1/ai/usages | List usages
[**GetAiUsagesByUser**](AiAPI.md#GetAiUsagesByUser) | **Get** /v1/ai/usages/by-user | By User (usage)
[**GetAiUsagesCloud**](AiAPI.md#GetAiUsagesCloud) | **Get** /v1/ai/usages/cloud | Cloud (usage)
[**GetAiUsagesRange**](AiAPI.md#GetAiUsagesRange) | **Get** /v1/ai/usages/range | Range (usage)
[**GetAiUsagesUserNames**](AiAPI.md#GetAiUsagesUserNames) | **Get** /v1/ai/usages/user-names | User Names (usage)
[**GetAiVectors**](AiAPI.md#GetAiVectors) | **Get** /v1/ai/vectors | List vectors
[**GetAiVectorsByOwnerByName**](AiAPI.md#GetAiVectorsByOwnerByName) | **Get** /v1/ai/vectors/{owner}/{name} | Retrieve a vector
[**GetAiVectorsGlobal**](AiAPI.md#GetAiVectorsGlobal) | **Get** /v1/ai/vectors/global | List vectors across tenants
[**GetAiVersion**](AiAPI.md#GetAiVersion) | **Get** /v1/ai/version | Version
[**GetAiVideos**](AiAPI.md#GetAiVideos) | **Get** /v1/ai/videos | List videos
[**GetAiVideosByOwnerByName**](AiAPI.md#GetAiVideosByOwnerByName) | **Get** /v1/ai/videos/{owner}/{name} | Retrieve a video
[**GetAiVideosGlobal**](AiAPI.md#GetAiVideosGlobal) | **Get** /v1/ai/videos/global | List videos across tenants
[**GetAiWorkflows**](AiAPI.md#GetAiWorkflows) | **Get** /v1/ai/workflows | List workflows
[**GetAiWorkflowsByOwnerByName**](AiAPI.md#GetAiWorkflowsByOwnerByName) | **Get** /v1/ai/workflows/{owner}/{name} | Retrieve a workflow
[**GetAiWorkflowsGlobal**](AiAPI.md#GetAiWorkflowsGlobal) | **Get** /v1/ai/workflows/global | List workflows across tenants
[**PatchAiArticlesByOwnerByName**](AiAPI.md#PatchAiArticlesByOwnerByName) | **Patch** /v1/ai/articles/{owner}/{name} | Update a article
[**PatchAiAssetsByOwnerByName**](AiAPI.md#PatchAiAssetsByOwnerByName) | **Patch** /v1/ai/assets/{owner}/{name} | Update a asset
[**PatchAiChatsByOwnerByName**](AiAPI.md#PatchAiChatsByOwnerByName) | **Patch** /v1/ai/chats/{owner}/{name} | Update a chat
[**PatchAiDeploymentsByOwnerByName**](AiAPI.md#PatchAiDeploymentsByOwnerByName) | **Patch** /v1/ai/deployments/{owner}/{name} | Update a application
[**PatchAiFilesByOwnerByName**](AiAPI.md#PatchAiFilesByOwnerByName) | **Patch** /v1/ai/files/{owner}/{name} | Update a file
[**PatchAiFormsByOwnerByName**](AiAPI.md#PatchAiFormsByOwnerByName) | **Patch** /v1/ai/forms/{owner}/{name} | Update a form
[**PatchAiGraphsByOwnerByName**](AiAPI.md#PatchAiGraphsByOwnerByName) | **Patch** /v1/ai/graphs/{owner}/{name} | Update a graph
[**PatchAiMessagesByOwnerByName**](AiAPI.md#PatchAiMessagesByOwnerByName) | **Patch** /v1/ai/messages/{owner}/{name} | Update a message
[**PatchAiNodesByOwnerByName**](AiAPI.md#PatchAiNodesByOwnerByName) | **Patch** /v1/ai/nodes/{owner}/{name} | Update a node
[**PatchAiPreferences**](AiAPI.md#PatchAiPreferences) | **Patch** /v1/ai/preferences | Preferences
[**PatchAiProvidersByOwnerByName**](AiAPI.md#PatchAiProvidersByOwnerByName) | **Patch** /v1/ai/providers/{owner}/{name} | Update a provider
[**PatchAiRecordsByOwnerByName**](AiAPI.md#PatchAiRecordsByOwnerByName) | **Patch** /v1/ai/records/{owner}/{name} | Update a record
[**PatchAiRemoteConnectionsByOwnerByName**](AiAPI.md#PatchAiRemoteConnectionsByOwnerByName) | **Patch** /v1/ai/remote-connections/{owner}/{name} | Update a connection
[**PatchAiRoutesByOwnerByName**](AiAPI.md#PatchAiRoutesByOwnerByName) | **Patch** /v1/ai/routes/{owner}/{name} | Update a model-route
[**PatchAiScalesByOwnerByName**](AiAPI.md#PatchAiScalesByOwnerByName) | **Patch** /v1/ai/scales/{owner}/{name} | Update a scale
[**PatchAiScansByOwnerByName**](AiAPI.md#PatchAiScansByOwnerByName) | **Patch** /v1/ai/scans/{owner}/{name} | Update a scan
[**PatchAiSigninSessionsByOwnerByName**](AiAPI.md#PatchAiSigninSessionsByOwnerByName) | **Patch** /v1/ai/signin-sessions/{owner}/{name} | Update a session
[**PatchAiStoresByOwnerByName**](AiAPI.md#PatchAiStoresByOwnerByName) | **Patch** /v1/ai/stores/{owner}/{name} | Update a store
[**PatchAiTasksByOwnerByName**](AiAPI.md#PatchAiTasksByOwnerByName) | **Patch** /v1/ai/tasks/{owner}/{name} | Update a task
[**PatchAiTemplatesByOwnerByName**](AiAPI.md#PatchAiTemplatesByOwnerByName) | **Patch** /v1/ai/templates/{owner}/{name} | Update a template
[**PatchAiTrainingContribution**](AiAPI.md#PatchAiTrainingContribution) | **Patch** /v1/ai/training-contribution | Training Contribution
[**PatchAiTreeFilesByOwnerByName**](AiAPI.md#PatchAiTreeFilesByOwnerByName) | **Patch** /v1/ai/tree-files/{owner}/{name} | Update a tree-file
[**PatchAiVectorsByOwnerByName**](AiAPI.md#PatchAiVectorsByOwnerByName) | **Patch** /v1/ai/vectors/{owner}/{name} | Update a vector
[**PatchAiVideosByOwnerByName**](AiAPI.md#PatchAiVideosByOwnerByName) | **Patch** /v1/ai/videos/{owner}/{name} | Update a video
[**PatchAiWorkflowsByOwnerByName**](AiAPI.md#PatchAiWorkflowsByOwnerByName) | **Patch** /v1/ai/workflows/{owner}/{name} | Update a workflow
[**PostAiArticles**](AiAPI.md#PostAiArticles) | **Post** /v1/ai/articles | Create a article
[**PostAiAssets**](AiAPI.md#PostAiAssets) | **Post** /v1/ai/assets | Create a asset
[**PostAiAssetsByOwnerByNameScan**](AiAPI.md#PostAiAssetsByOwnerByNameScan) | **Post** /v1/ai/assets/{owner}/{name}/scan | Scan (asset)
[**PostAiAssetsScan**](AiAPI.md#PostAiAssetsScan) | **Post** /v1/ai/assets/scan | Scan (asset)
[**PostAiChats**](AiAPI.md#PostAiChats) | **Post** /v1/ai/chats | Create a chat
[**PostAiConnections**](AiAPI.md#PostAiConnections) | **Post** /v1/ai/connections | Connects (or reconnects) a third-party AI account for the org by sealing the supplied key into KMS and upserting the org&#39;s provider row.
[**PostAiConnectionsByProvider**](AiAPI.md#PostAiConnectionsByProvider) | **Post** /v1/ai/connections/{provider} | Disconnects a third-party AI account: it deactivates the org&#39;s row so completion resolution falls back to the global Hanzo account (no BYO), and best-effort tombstones the sealed secret.
[**PostAiDeployments**](AiAPI.md#PostAiDeployments) | **Post** /v1/ai/deployments | Create a application
[**PostAiDeploymentsByOwnerByNameDeploy**](AiAPI.md#PostAiDeploymentsByOwnerByNameDeploy) | **Post** /v1/ai/deployments/{owner}/{name}/deploy | Deploy (application)
[**PostAiDeploymentsByOwnerByNameUndeploy**](AiAPI.md#PostAiDeploymentsByOwnerByNameUndeploy) | **Post** /v1/ai/deployments/{owner}/{name}/undeploy | Undeploy (application)
[**PostAiFiles**](AiAPI.md#PostAiFiles) | **Post** /v1/ai/files | Create a file
[**PostAiFilesActivate**](AiAPI.md#PostAiFilesActivate) | **Post** /v1/ai/files/activate | Activate (file)
[**PostAiFilesByOwnerByNameVectors**](AiAPI.md#PostAiFilesByOwnerByNameVectors) | **Post** /v1/ai/files/{owner}/{name}/vectors | Vectors (file)
[**PostAiFilesUpload**](AiAPI.md#PostAiFilesUpload) | **Post** /v1/ai/files/upload | Upload (file)
[**PostAiForms**](AiAPI.md#PostAiForms) | **Post** /v1/ai/forms | Create a form
[**PostAiGraphs**](AiAPI.md#PostAiGraphs) | **Post** /v1/ai/graphs | Create a graph
[**PostAiMessages**](AiAPI.md#PostAiMessages) | **Post** /v1/ai/messages | Create a message
[**PostAiNodes**](AiAPI.md#PostAiNodes) | **Post** /v1/ai/nodes | Create a node
[**PostAiNodesByOwnerByNameTunnel**](AiAPI.md#PostAiNodesByOwnerByNameTunnel) | **Post** /v1/ai/nodes/{owner}/{name}/tunnel | Tunnel (node)
[**PostAiProviders**](AiAPI.md#PostAiProviders) | **Post** /v1/ai/providers | Create a provider
[**PostAiProvidersMcpTools**](AiAPI.md#PostAiProvidersMcpTools) | **Post** /v1/ai/providers/mcp-tools | Mcp Tools (provider)
[**PostAiRecords**](AiAPI.md#PostAiRecords) | **Post** /v1/ai/records | Create a record
[**PostAiRecordsBatch**](AiAPI.md#PostAiRecordsBatch) | **Post** /v1/ai/records/batch | Batch (record)
[**PostAiRecordsCommit**](AiAPI.md#PostAiRecordsCommit) | **Post** /v1/ai/records/commit | Commit (record)
[**PostAiRecordsCommitSecond**](AiAPI.md#PostAiRecordsCommitSecond) | **Post** /v1/ai/records/commit-second | Commit Second (record)
[**PostAiRemoteConnections**](AiAPI.md#PostAiRemoteConnections) | **Post** /v1/ai/remote-connections | Create a connection
[**PostAiRemoteConnectionsByOwnerByNameStart**](AiAPI.md#PostAiRemoteConnectionsByOwnerByNameStart) | **Post** /v1/ai/remote-connections/{owner}/{name}/start | Start (connection)
[**PostAiRemoteConnectionsByOwnerByNameStop**](AiAPI.md#PostAiRemoteConnectionsByOwnerByNameStop) | **Post** /v1/ai/remote-connections/{owner}/{name}/stop | Stop (connection)
[**PostAiRoutes**](AiAPI.md#PostAiRoutes) | **Post** /v1/ai/routes | Create a model-route
[**PostAiScales**](AiAPI.md#PostAiScales) | **Post** /v1/ai/scales | Create a scale
[**PostAiScans**](AiAPI.md#PostAiScans) | **Post** /v1/ai/scans | Create a scan
[**PostAiSignin**](AiAPI.md#PostAiSignin) | **Post** /v1/ai/signin | Signin
[**PostAiSigninSessions**](AiAPI.md#PostAiSigninSessions) | **Post** /v1/ai/signin-sessions | Create a session
[**PostAiSignout**](AiAPI.md#PostAiSignout) | **Post** /v1/ai/signout | Signout
[**PostAiStores**](AiAPI.md#PostAiStores) | **Post** /v1/ai/stores | Create a store
[**PostAiStoresByOwnerByNameVectors**](AiAPI.md#PostAiStoresByOwnerByNameVectors) | **Post** /v1/ai/stores/{owner}/{name}/vectors | Vectors (store)
[**PostAiTasks**](AiAPI.md#PostAiTasks) | **Post** /v1/ai/tasks | Create a task
[**PostAiTasksByOwnerByNameAnalyze**](AiAPI.md#PostAiTasksByOwnerByNameAnalyze) | **Post** /v1/ai/tasks/{owner}/{name}/analyze | Analyze (task)
[**PostAiTasksByOwnerByNameDocument**](AiAPI.md#PostAiTasksByOwnerByNameDocument) | **Post** /v1/ai/tasks/{owner}/{name}/document | Document (task)
[**PostAiTemplates**](AiAPI.md#PostAiTemplates) | **Post** /v1/ai/templates | Create a template
[**PostAiTreeFiles**](AiAPI.md#PostAiTreeFiles) | **Post** /v1/ai/tree-files | Create a tree-file
[**PostAiVectors**](AiAPI.md#PostAiVectors) | **Post** /v1/ai/vectors | Create a vector
[**PostAiVideos**](AiAPI.md#PostAiVideos) | **Post** /v1/ai/videos | Create a video
[**PostAiVideosUpload**](AiAPI.md#PostAiVideosUpload) | **Post** /v1/ai/videos/upload | Upload (video)
[**PostAiWorkflows**](AiAPI.md#PostAiWorkflows) | **Post** /v1/ai/workflows | Create a workflow
[**PutAiArticlesByOwnerByName**](AiAPI.md#PutAiArticlesByOwnerByName) | **Put** /v1/ai/articles/{owner}/{name} | Replace a article
[**PutAiAssetsByOwnerByName**](AiAPI.md#PutAiAssetsByOwnerByName) | **Put** /v1/ai/assets/{owner}/{name} | Replace a asset
[**PutAiChatsByOwnerByName**](AiAPI.md#PutAiChatsByOwnerByName) | **Put** /v1/ai/chats/{owner}/{name} | Replace a chat
[**PutAiDeploymentsByOwnerByName**](AiAPI.md#PutAiDeploymentsByOwnerByName) | **Put** /v1/ai/deployments/{owner}/{name} | Replace a application
[**PutAiFilesByOwnerByName**](AiAPI.md#PutAiFilesByOwnerByName) | **Put** /v1/ai/files/{owner}/{name} | Replace a file
[**PutAiFormsByOwnerByName**](AiAPI.md#PutAiFormsByOwnerByName) | **Put** /v1/ai/forms/{owner}/{name} | Replace a form
[**PutAiGraphsByOwnerByName**](AiAPI.md#PutAiGraphsByOwnerByName) | **Put** /v1/ai/graphs/{owner}/{name} | Replace a graph
[**PutAiMessagesByOwnerByName**](AiAPI.md#PutAiMessagesByOwnerByName) | **Put** /v1/ai/messages/{owner}/{name} | Replace a message
[**PutAiNodesByOwnerByName**](AiAPI.md#PutAiNodesByOwnerByName) | **Put** /v1/ai/nodes/{owner}/{name} | Replace a node
[**PutAiPreferences**](AiAPI.md#PutAiPreferences) | **Put** /v1/ai/preferences | Preferences
[**PutAiProvidersByOwnerByName**](AiAPI.md#PutAiProvidersByOwnerByName) | **Put** /v1/ai/providers/{owner}/{name} | Replace a provider
[**PutAiRecordsByOwnerByName**](AiAPI.md#PutAiRecordsByOwnerByName) | **Put** /v1/ai/records/{owner}/{name} | Replace a record
[**PutAiRemoteConnectionsByOwnerByName**](AiAPI.md#PutAiRemoteConnectionsByOwnerByName) | **Put** /v1/ai/remote-connections/{owner}/{name} | Replace a connection
[**PutAiRoutesByOwnerByName**](AiAPI.md#PutAiRoutesByOwnerByName) | **Put** /v1/ai/routes/{owner}/{name} | Replace a model-route
[**PutAiScalesByOwnerByName**](AiAPI.md#PutAiScalesByOwnerByName) | **Put** /v1/ai/scales/{owner}/{name} | Replace a scale
[**PutAiScansByOwnerByName**](AiAPI.md#PutAiScansByOwnerByName) | **Put** /v1/ai/scans/{owner}/{name} | Replace a scan
[**PutAiSigninSessionsByOwnerByName**](AiAPI.md#PutAiSigninSessionsByOwnerByName) | **Put** /v1/ai/signin-sessions/{owner}/{name} | Replace a session
[**PutAiStoresByOwnerByName**](AiAPI.md#PutAiStoresByOwnerByName) | **Put** /v1/ai/stores/{owner}/{name} | Replace a store
[**PutAiTasksByOwnerByName**](AiAPI.md#PutAiTasksByOwnerByName) | **Put** /v1/ai/tasks/{owner}/{name} | Replace a task
[**PutAiTemplatesByOwnerByName**](AiAPI.md#PutAiTemplatesByOwnerByName) | **Put** /v1/ai/templates/{owner}/{name} | Replace a template
[**PutAiTrainingContribution**](AiAPI.md#PutAiTrainingContribution) | **Put** /v1/ai/training-contribution | Training Contribution
[**PutAiTreeFilesByOwnerByName**](AiAPI.md#PutAiTreeFilesByOwnerByName) | **Put** /v1/ai/tree-files/{owner}/{name} | Replace a tree-file
[**PutAiVectorsByOwnerByName**](AiAPI.md#PutAiVectorsByOwnerByName) | **Put** /v1/ai/vectors/{owner}/{name} | Replace a vector
[**PutAiVideosByOwnerByName**](AiAPI.md#PutAiVideosByOwnerByName) | **Put** /v1/ai/videos/{owner}/{name} | Replace a video
[**PutAiWorkflowsByOwnerByName**](AiAPI.md#PutAiWorkflowsByOwnerByName) | **Put** /v1/ai/workflows/{owner}/{name} | Replace a workflow



## AiMCPTools

> AiMCPSurface AiMCPTools(ctx).Names(names).Execute()

Tools reports what THIS PROCESS's MCP door carries: how many tools its own registry projects, optionally their names, and which subsystems this process composed.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	names := true // bool | Names asks for this process's tool NAMES and not only how many there are. Off by default: a list of names is a page, and the question this op exists to answer (\"is the door up and does it have anything behind it\") is answered by the count. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.AiMCPTools(context.Background()).Names(names).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.AiMCPTools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiMCPTools`: AiMCPSurface
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.AiMCPTools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiMCPToolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **names** | **bool** | Names asks for this process&#39;s tool NAMES and not only how many there are. Off by default: a list of names is a page, and the question this op exists to answer (\&quot;is the door up and does it have anything behind it\&quot;) is answered by the count. | 

### Return type

[**AiMCPSurface**](AiMCPSurface.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiArticlesByOwnerByName

> PostAiArticles200Response DeleteAiArticlesByOwnerByName(ctx, owner, name).Execute()

Delete a article



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.DeleteAiArticlesByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.DeleteAiArticlesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiArticlesByOwnerByName`: PostAiArticles200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.DeleteAiArticlesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiArticlesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiArticles200Response**](PostAiArticles200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiAssetsByOwnerByName

> PostAiAssets200Response DeleteAiAssetsByOwnerByName(ctx, owner, name).Execute()

Delete a asset



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.DeleteAiAssetsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.DeleteAiAssetsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiAssetsByOwnerByName`: PostAiAssets200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.DeleteAiAssetsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiAssetsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiAssets200Response**](PostAiAssets200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiChatsByOwnerByName

> PostAiChats200Response DeleteAiChatsByOwnerByName(ctx, owner, name).Execute()

Delete a chat



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.DeleteAiChatsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.DeleteAiChatsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiChatsByOwnerByName`: PostAiChats200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.DeleteAiChatsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiChatsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiChats200Response**](PostAiChats200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiConnectionsByProvider

> DeleteAiConnectionsByProvider(ctx).Execute()

Disconnects a third-party AI account: it deactivates the org's row so completion resolution falls back to the global Hanzo account (no BYO), and best-effort tombstones the sealed secret.



### Example

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
	r, err := apiClient.AiAPI.DeleteAiConnectionsByProvider(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.DeleteAiConnectionsByProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiConnectionsByProviderRequest struct via the builder pattern


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


## DeleteAiDeploymentsByOwnerByName

> PostAiDeployments200Response DeleteAiDeploymentsByOwnerByName(ctx, owner, name).Execute()

Delete a application



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.DeleteAiDeploymentsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.DeleteAiDeploymentsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiDeploymentsByOwnerByName`: PostAiDeployments200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.DeleteAiDeploymentsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiDeploymentsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiDeployments200Response**](PostAiDeployments200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiFilesByOwnerByName

> PostAiFiles200Response DeleteAiFilesByOwnerByName(ctx, owner, name).Execute()

Delete a file



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.DeleteAiFilesByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.DeleteAiFilesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiFilesByOwnerByName`: PostAiFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.DeleteAiFilesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiFilesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiFiles200Response**](PostAiFiles200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiFormsByOwnerByName

> PostAiForms200Response DeleteAiFormsByOwnerByName(ctx, owner, name).Execute()

Delete a form



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.DeleteAiFormsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.DeleteAiFormsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiFormsByOwnerByName`: PostAiForms200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.DeleteAiFormsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiFormsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiForms200Response**](PostAiForms200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiGraphsByOwnerByName

> PostAiGraphs200Response DeleteAiGraphsByOwnerByName(ctx, owner, name).Execute()

Delete a graph



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.DeleteAiGraphsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.DeleteAiGraphsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiGraphsByOwnerByName`: PostAiGraphs200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.DeleteAiGraphsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiGraphsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiGraphs200Response**](PostAiGraphs200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiMessagesByOwnerByName

> PostAiMessages200Response DeleteAiMessagesByOwnerByName(ctx, owner, name).Execute()

Delete a message



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.DeleteAiMessagesByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.DeleteAiMessagesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiMessagesByOwnerByName`: PostAiMessages200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.DeleteAiMessagesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiMessagesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiMessages200Response**](PostAiMessages200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiMessagesWelcome

> Envelope DeleteAiMessagesWelcome(ctx).Execute()

Welcome (message)

### Example

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
	resp, r, err := apiClient.AiAPI.DeleteAiMessagesWelcome(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.DeleteAiMessagesWelcome``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiMessagesWelcome`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.DeleteAiMessagesWelcome`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiMessagesWelcomeRequest struct via the builder pattern


### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiNodesByOwnerByName

> PostAiNodes200Response DeleteAiNodesByOwnerByName(ctx, owner, name).Execute()

Delete a node



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.DeleteAiNodesByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.DeleteAiNodesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiNodesByOwnerByName`: PostAiNodes200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.DeleteAiNodesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiNodesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiNodes200Response**](PostAiNodes200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiProvidersByOwnerByName

> PostAiProviders200Response DeleteAiProvidersByOwnerByName(ctx, owner, name).Execute()

Delete a provider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.DeleteAiProvidersByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.DeleteAiProvidersByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiProvidersByOwnerByName`: PostAiProviders200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.DeleteAiProvidersByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiProvidersByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiProviders200Response**](PostAiProviders200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiRecordsByOwnerByName

> PostAiRecords200Response DeleteAiRecordsByOwnerByName(ctx, owner, name).Execute()

Delete a record



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.DeleteAiRecordsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.DeleteAiRecordsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiRecordsByOwnerByName`: PostAiRecords200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.DeleteAiRecordsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiRecordsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiRecords200Response**](PostAiRecords200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiRemoteConnectionsByOwnerByName

> PostAiRemoteConnections200Response DeleteAiRemoteConnectionsByOwnerByName(ctx, owner, name).Execute()

Delete a connection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.DeleteAiRemoteConnectionsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.DeleteAiRemoteConnectionsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiRemoteConnectionsByOwnerByName`: PostAiRemoteConnections200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.DeleteAiRemoteConnectionsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiRemoteConnectionsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiRemoteConnections200Response**](PostAiRemoteConnections200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiRoutesByOwnerByName

> PostAiRoutes200Response DeleteAiRoutesByOwnerByName(ctx, owner, name).Execute()

Delete a model-route



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.DeleteAiRoutesByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.DeleteAiRoutesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiRoutesByOwnerByName`: PostAiRoutes200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.DeleteAiRoutesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiRoutesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiRoutes200Response**](PostAiRoutes200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiScalesByOwnerByName

> PostAiScales200Response DeleteAiScalesByOwnerByName(ctx, owner, name).Execute()

Delete a scale



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.DeleteAiScalesByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.DeleteAiScalesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiScalesByOwnerByName`: PostAiScales200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.DeleteAiScalesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiScalesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiScales200Response**](PostAiScales200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiScansByOwnerByName

> PostAiScans200Response DeleteAiScansByOwnerByName(ctx, owner, name).Execute()

Delete a scan



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.DeleteAiScansByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.DeleteAiScansByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiScansByOwnerByName`: PostAiScans200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.DeleteAiScansByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiScansByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiScans200Response**](PostAiScans200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiSigninSessionsByOwnerByName

> PostAiSigninSessions200Response DeleteAiSigninSessionsByOwnerByName(ctx, owner, name).Execute()

Delete a session



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.DeleteAiSigninSessionsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.DeleteAiSigninSessionsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiSigninSessionsByOwnerByName`: PostAiSigninSessions200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.DeleteAiSigninSessionsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiSigninSessionsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiSigninSessions200Response**](PostAiSigninSessions200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiStoresByOwnerByName

> PostAiStores200Response DeleteAiStoresByOwnerByName(ctx, owner, name).Execute()

Delete a store



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.DeleteAiStoresByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.DeleteAiStoresByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiStoresByOwnerByName`: PostAiStores200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.DeleteAiStoresByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiStoresByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiStores200Response**](PostAiStores200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiTasksByOwnerByName

> PostAiTasks200Response DeleteAiTasksByOwnerByName(ctx, owner, name).Execute()

Delete a task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.DeleteAiTasksByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.DeleteAiTasksByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiTasksByOwnerByName`: PostAiTasks200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.DeleteAiTasksByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiTasksByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiTasks200Response**](PostAiTasks200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiTemplatesByOwnerByName

> PostAiTemplates200Response DeleteAiTemplatesByOwnerByName(ctx, owner, name).Execute()

Delete a template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.DeleteAiTemplatesByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.DeleteAiTemplatesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiTemplatesByOwnerByName`: PostAiTemplates200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.DeleteAiTemplatesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiTemplatesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiTemplates200Response**](PostAiTemplates200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiTreeFilesByOwnerByName

> PostAiTreeFiles200Response DeleteAiTreeFilesByOwnerByName(ctx, owner, name).Execute()

Delete a tree-file



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.DeleteAiTreeFilesByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.DeleteAiTreeFilesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiTreeFilesByOwnerByName`: PostAiTreeFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.DeleteAiTreeFilesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiTreeFilesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiTreeFiles200Response**](PostAiTreeFiles200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiVectorsAll

> Envelope DeleteAiVectorsAll(ctx).Execute()

All (vector)

### Example

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
	resp, r, err := apiClient.AiAPI.DeleteAiVectorsAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.DeleteAiVectorsAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiVectorsAll`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.DeleteAiVectorsAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiVectorsAllRequest struct via the builder pattern


### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiVectorsByOwnerByName

> PostAiVectors200Response DeleteAiVectorsByOwnerByName(ctx, owner, name).Execute()

Delete a vector



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.DeleteAiVectorsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.DeleteAiVectorsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiVectorsByOwnerByName`: PostAiVectors200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.DeleteAiVectorsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiVectorsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiVectors200Response**](PostAiVectors200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiVideosByOwnerByName

> PostAiVideos200Response DeleteAiVideosByOwnerByName(ctx, owner, name).Execute()

Delete a video



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.DeleteAiVideosByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.DeleteAiVideosByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiVideosByOwnerByName`: PostAiVideos200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.DeleteAiVideosByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiVideosByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiVideos200Response**](PostAiVideos200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiWorkflowsByOwnerByName

> PostAiWorkflows200Response DeleteAiWorkflowsByOwnerByName(ctx, owner, name).Execute()

Delete a workflow



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.DeleteAiWorkflowsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.DeleteAiWorkflowsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiWorkflowsByOwnerByName`: PostAiWorkflows200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.DeleteAiWorkflowsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiWorkflowsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiWorkflows200Response**](PostAiWorkflows200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiAccount

> Envelope GetAiAccount(ctx).Execute()

Account

### Example

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
	resp, r, err := apiClient.AiAPI.GetAiAccount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiAccount`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiAccountRequest struct via the builder pattern


### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiActivities

> GetAiActivities200Response GetAiActivities(ctx).Execute()

List activities



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiActivities(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiActivities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiActivities`: GetAiActivities200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiActivities`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiActivitiesRequest struct via the builder pattern


### Return type

[**GetAiActivities200Response**](GetAiActivities200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiAnswer

> Envelope GetAiAnswer(ctx).Execute()

Answer

### Example

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
	resp, r, err := apiClient.AiAPI.GetAiAnswer(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiAnswer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiAnswer`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiAnswer`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiAnswerRequest struct via the builder pattern


### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiArticles

> GetAiArticles200Response GetAiArticles(ctx).Execute()

List articles



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiArticles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiArticles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiArticles`: GetAiArticles200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiArticles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiArticlesRequest struct via the builder pattern


### Return type

[**GetAiArticles200Response**](GetAiArticles200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiArticlesByOwnerByName

> PostAiArticles200Response GetAiArticlesByOwnerByName(ctx, owner, name).Execute()

Retrieve a article



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.GetAiArticlesByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiArticlesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiArticlesByOwnerByName`: PostAiArticles200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiArticlesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiArticlesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiArticles200Response**](PostAiArticles200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiArticlesGlobal

> GetAiArticles200Response GetAiArticlesGlobal(ctx).Execute()

List articles across tenants



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiArticlesGlobal(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiArticlesGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiArticlesGlobal`: GetAiArticles200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiArticlesGlobal`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiArticlesGlobalRequest struct via the builder pattern


### Return type

[**GetAiArticles200Response**](GetAiArticles200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiAssets

> GetAiAssets200Response GetAiAssets(ctx).Execute()

List assets



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiAssets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiAssets`: GetAiAssets200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiAssets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiAssetsRequest struct via the builder pattern


### Return type

[**GetAiAssets200Response**](GetAiAssets200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiAssetsByOwnerByName

> PostAiAssets200Response GetAiAssetsByOwnerByName(ctx, owner, name).Execute()

Retrieve a asset



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.GetAiAssetsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiAssetsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiAssetsByOwnerByName`: PostAiAssets200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiAssetsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiAssetsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiAssets200Response**](PostAiAssets200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiChats

> GetAiChats200Response GetAiChats(ctx).Execute()

List chats



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiChats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiChats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiChats`: GetAiChats200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiChats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiChatsRequest struct via the builder pattern


### Return type

[**GetAiChats200Response**](GetAiChats200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiChatsByOwnerByName

> PostAiChats200Response GetAiChatsByOwnerByName(ctx, owner, name).Execute()

Retrieve a chat



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.GetAiChatsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiChatsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiChatsByOwnerByName`: PostAiChats200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiChatsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiChatsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiChats200Response**](PostAiChats200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiChatsGlobal

> GetAiChats200Response GetAiChatsGlobal(ctx).Execute()

List chats across tenants



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiChatsGlobal(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiChatsGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiChatsGlobal`: GetAiChats200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiChatsGlobal`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiChatsGlobalRequest struct via the builder pattern


### Return type

[**GetAiChats200Response**](GetAiChats200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiConnections

> GetAiConnections(ctx).Execute()

Lists the org's connectable AI accounts and whether each is currently connected.



### Example

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
	r, err := apiClient.AiAPI.GetAiConnections(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiConnectionsRequest struct via the builder pattern


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


## GetAiConnectionsByProviderAuthorize

> GetAiConnectionsByProviderAuthorize(ctx).Execute()

Begins an OAuth connection for the caller's org: it binds the org into a signed state and sends the caller to the provider's authorize URL.



### Example

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
	r, err := apiClient.AiAPI.GetAiConnectionsByProviderAuthorize(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiConnectionsByProviderAuthorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiConnectionsByProviderAuthorizeRequest struct via the builder pattern


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


## GetAiConnectionsByProviderCallback

> GetAiConnectionsByProviderCallback(ctx).Execute()

Completes OAuth: the org is recovered from the SIGNED state (not a header), the code is exchanged for a token, the token is SEALED into KMS (never the row/logs) through the same path as a BYOK key, and the org's provider row is upserted to \"connected\".



### Example

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
	r, err := apiClient.AiAPI.GetAiConnectionsByProviderCallback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiConnectionsByProviderCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiConnectionsByProviderCallbackRequest struct via the builder pattern


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


## GetAiConnectionsByProviderUsage

> GetAiConnectionsByProviderUsage(ctx).Execute()

Imports the caller org's usage for a connected third-party account.



### Example

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
	r, err := apiClient.AiAPI.GetAiConnectionsByProviderUsage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiConnectionsByProviderUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiConnectionsByProviderUsageRequest struct via the builder pattern


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


## GetAiDashboardsAgents

> Envelope GetAiDashboardsAgents(ctx).Execute()

Dashboards Agents

### Example

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
	resp, r, err := apiClient.AiAPI.GetAiDashboardsAgents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiDashboardsAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiDashboardsAgents`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiDashboardsAgents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiDashboardsAgentsRequest struct via the builder pattern


### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiDashboardsVm

> Envelope GetAiDashboardsVm(ctx).Execute()

Dashboards Vm

### Example

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
	resp, r, err := apiClient.AiAPI.GetAiDashboardsVm(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiDashboardsVm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiDashboardsVm`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiDashboardsVm`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiDashboardsVmRequest struct via the builder pattern


### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiDeployments

> GetAiDeployments200Response GetAiDeployments(ctx).Execute()

List deployments



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiDeployments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiDeployments`: GetAiDeployments200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiDeployments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiDeploymentsRequest struct via the builder pattern


### Return type

[**GetAiDeployments200Response**](GetAiDeployments200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiDeploymentsByOwnerByName

> PostAiDeployments200Response GetAiDeploymentsByOwnerByName(ctx, owner, name).Execute()

Retrieve a application



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.GetAiDeploymentsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiDeploymentsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiDeploymentsByOwnerByName`: PostAiDeployments200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiDeploymentsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiDeploymentsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiDeployments200Response**](PostAiDeployments200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiFiles

> GetAiFiles200Response GetAiFiles(ctx).Execute()

List files



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiFiles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiFiles`: GetAiFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiFiles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiFilesRequest struct via the builder pattern


### Return type

[**GetAiFiles200Response**](GetAiFiles200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiFilesActive

> Envelope GetAiFilesActive(ctx).Execute()

Active (file)

### Example

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
	resp, r, err := apiClient.AiAPI.GetAiFilesActive(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiFilesActive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiFilesActive`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiFilesActive`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiFilesActiveRequest struct via the builder pattern


### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiFilesByOwnerByName

> PostAiFiles200Response GetAiFilesByOwnerByName(ctx, owner, name).Execute()

Retrieve a file



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.GetAiFilesByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiFilesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiFilesByOwnerByName`: PostAiFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiFilesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiFilesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiFiles200Response**](PostAiFiles200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiFilesGlobal

> GetAiFiles200Response GetAiFilesGlobal(ctx).Execute()

List files across tenants



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiFilesGlobal(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiFilesGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiFilesGlobal`: GetAiFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiFilesGlobal`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiFilesGlobalRequest struct via the builder pattern


### Return type

[**GetAiFiles200Response**](GetAiFiles200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiForms

> GetAiForms200Response GetAiForms(ctx).Execute()

List forms



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiForms(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiForms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiForms`: GetAiForms200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiForms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiFormsRequest struct via the builder pattern


### Return type

[**GetAiForms200Response**](GetAiForms200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiFormsByOwnerByName

> PostAiForms200Response GetAiFormsByOwnerByName(ctx, owner, name).Execute()

Retrieve a form



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.GetAiFormsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiFormsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiFormsByOwnerByName`: PostAiForms200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiFormsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiFormsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiForms200Response**](PostAiForms200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiFormsData

> Envelope GetAiFormsData(ctx).Execute()

Data (form)

### Example

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
	resp, r, err := apiClient.AiAPI.GetAiFormsData(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiFormsData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiFormsData`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiFormsData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiFormsDataRequest struct via the builder pattern


### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiFormsGlobal

> GetAiForms200Response GetAiFormsGlobal(ctx).Execute()

List forms across tenants



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiFormsGlobal(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiFormsGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiFormsGlobal`: GetAiForms200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiFormsGlobal`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiFormsGlobalRequest struct via the builder pattern


### Return type

[**GetAiForms200Response**](GetAiForms200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiGraphs

> GetAiGraphs200Response GetAiGraphs(ctx).Execute()

List graphs



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiGraphs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiGraphs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiGraphs`: GetAiGraphs200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiGraphs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiGraphsRequest struct via the builder pattern


### Return type

[**GetAiGraphs200Response**](GetAiGraphs200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiGraphsByOwnerByName

> PostAiGraphs200Response GetAiGraphsByOwnerByName(ctx, owner, name).Execute()

Retrieve a graph



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.GetAiGraphsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiGraphsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiGraphsByOwnerByName`: PostAiGraphs200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiGraphsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiGraphsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiGraphs200Response**](PostAiGraphs200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiGraphsGlobal

> GetAiGraphs200Response GetAiGraphsGlobal(ctx).Execute()

List graphs across tenants



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiGraphsGlobal(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiGraphsGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiGraphsGlobal`: GetAiGraphs200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiGraphsGlobal`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiGraphsGlobalRequest struct via the builder pattern


### Return type

[**GetAiGraphs200Response**](GetAiGraphs200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiK8sStatus

> Envelope GetAiK8sStatus(ctx).Execute()

K8s Status

### Example

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
	resp, r, err := apiClient.AiAPI.GetAiK8sStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiK8sStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiK8sStatus`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiK8sStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiK8sStatusRequest struct via the builder pattern


### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiMessages

> GetAiMessages200Response GetAiMessages(ctx).Execute()

List messages



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiMessages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiMessages`: GetAiMessages200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiMessages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiMessagesRequest struct via the builder pattern


### Return type

[**GetAiMessages200Response**](GetAiMessages200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiMessagesByOwnerByName

> PostAiMessages200Response GetAiMessagesByOwnerByName(ctx, owner, name).Execute()

Retrieve a message



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.GetAiMessagesByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiMessagesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiMessagesByOwnerByName`: PostAiMessages200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiMessagesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiMessagesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiMessages200Response**](PostAiMessages200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiMessagesByOwnerByNameAnswer

> Envelope GetAiMessagesByOwnerByNameAnswer(ctx, owner, name).Execute()

Answer (message)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.GetAiMessagesByOwnerByNameAnswer(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiMessagesByOwnerByNameAnswer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiMessagesByOwnerByNameAnswer`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiMessagesByOwnerByNameAnswer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiMessagesByOwnerByNameAnswerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiMessagesGlobal

> GetAiMessages200Response GetAiMessagesGlobal(ctx).Execute()

List messages across tenants



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiMessagesGlobal(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiMessagesGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiMessagesGlobal`: GetAiMessages200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiMessagesGlobal`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiMessagesGlobalRequest struct via the builder pattern


### Return type

[**GetAiMessages200Response**](GetAiMessages200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiNodes

> GetAiNodes200Response GetAiNodes(ctx).Execute()

List nodes



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiNodes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiNodes`: GetAiNodes200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiNodesRequest struct via the builder pattern


### Return type

[**GetAiNodes200Response**](GetAiNodes200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiNodesByOwnerByName

> PostAiNodes200Response GetAiNodesByOwnerByName(ctx, owner, name).Execute()

Retrieve a node



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.GetAiNodesByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiNodesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiNodesByOwnerByName`: PostAiNodes200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiNodesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiNodesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiNodes200Response**](PostAiNodes200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiNodesByOwnerByNameTunnel

> Envelope GetAiNodesByOwnerByNameTunnel(ctx, owner, name).Execute()

Tunnel (node)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.GetAiNodesByOwnerByNameTunnel(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiNodesByOwnerByNameTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiNodesByOwnerByNameTunnel`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiNodesByOwnerByNameTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiNodesByOwnerByNameTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiPrometheus

> Envelope GetAiPrometheus(ctx).Execute()

Prometheus

### Example

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
	resp, r, err := apiClient.AiAPI.GetAiPrometheus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiPrometheus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiPrometheus`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiPrometheus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiPrometheusRequest struct via the builder pattern


### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiProviders

> GetAiProviders200Response GetAiProviders(ctx).Execute()

List providers



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiProviders`: GetAiProviders200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiProvidersRequest struct via the builder pattern


### Return type

[**GetAiProviders200Response**](GetAiProviders200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiProvidersByOwnerByName

> PostAiProviders200Response GetAiProvidersByOwnerByName(ctx, owner, name).Execute()

Retrieve a provider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.GetAiProvidersByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiProvidersByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiProvidersByOwnerByName`: PostAiProviders200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiProvidersByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiProvidersByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiProviders200Response**](PostAiProviders200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiProvidersGlobal

> GetAiProviders200Response GetAiProvidersGlobal(ctx).Execute()

List providers across tenants



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiProvidersGlobal(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiProvidersGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiProvidersGlobal`: GetAiProviders200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiProvidersGlobal`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiProvidersGlobalRequest struct via the builder pattern


### Return type

[**GetAiProviders200Response**](GetAiProviders200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiRecords

> GetAiRecords200Response GetAiRecords(ctx).Execute()

List records



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiRecords(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiRecords`: GetAiRecords200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiRecords`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiRecordsRequest struct via the builder pattern


### Return type

[**GetAiRecords200Response**](GetAiRecords200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiRecordsByOwnerByName

> PostAiRecords200Response GetAiRecordsByOwnerByName(ctx, owner, name).Execute()

Retrieve a record



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.GetAiRecordsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiRecordsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiRecordsByOwnerByName`: PostAiRecords200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiRecordsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiRecordsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiRecords200Response**](PostAiRecords200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiRecordsQuery

> Envelope GetAiRecordsQuery(ctx).Execute()

Query (record)

### Example

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
	resp, r, err := apiClient.AiAPI.GetAiRecordsQuery(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiRecordsQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiRecordsQuery`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiRecordsQuery`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiRecordsQueryRequest struct via the builder pattern


### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiRecordsQuerySecond

> Envelope GetAiRecordsQuerySecond(ctx).Execute()

Query Second (record)

### Example

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
	resp, r, err := apiClient.AiAPI.GetAiRecordsQuerySecond(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiRecordsQuerySecond``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiRecordsQuerySecond`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiRecordsQuerySecond`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiRecordsQuerySecondRequest struct via the builder pattern


### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiRemoteConnections

> GetAiRemoteConnections200Response GetAiRemoteConnections(ctx).Execute()

List remote-connections



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiRemoteConnections(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiRemoteConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiRemoteConnections`: GetAiRemoteConnections200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiRemoteConnections`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiRemoteConnectionsRequest struct via the builder pattern


### Return type

[**GetAiRemoteConnections200Response**](GetAiRemoteConnections200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiRemoteConnectionsByOwnerByName

> PostAiRemoteConnections200Response GetAiRemoteConnectionsByOwnerByName(ctx, owner, name).Execute()

Retrieve a connection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.GetAiRemoteConnectionsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiRemoteConnectionsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiRemoteConnectionsByOwnerByName`: PostAiRemoteConnections200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiRemoteConnectionsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiRemoteConnectionsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiRemoteConnections200Response**](PostAiRemoteConnections200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiRoutes

> GetAiRoutes200Response GetAiRoutes(ctx).Execute()

List routes



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiRoutes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiRoutes`: GetAiRoutes200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiRoutes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiRoutesRequest struct via the builder pattern


### Return type

[**GetAiRoutes200Response**](GetAiRoutes200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiRoutesByOwnerByName

> PostAiRoutes200Response GetAiRoutesByOwnerByName(ctx, owner, name).Execute()

Retrieve a model-route



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.GetAiRoutesByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiRoutesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiRoutesByOwnerByName`: PostAiRoutes200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiRoutesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiRoutesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiRoutes200Response**](PostAiRoutes200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiScales

> GetAiScales200Response GetAiScales(ctx).Execute()

List scales



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiScales(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiScales``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiScales`: GetAiScales200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiScales`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiScalesRequest struct via the builder pattern


### Return type

[**GetAiScales200Response**](GetAiScales200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiScalesByOwnerByName

> PostAiScales200Response GetAiScalesByOwnerByName(ctx, owner, name).Execute()

Retrieve a scale



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.GetAiScalesByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiScalesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiScalesByOwnerByName`: PostAiScales200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiScalesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiScalesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiScales200Response**](PostAiScales200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiScalesGlobal

> GetAiScales200Response GetAiScalesGlobal(ctx).Execute()

List scales across tenants



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiScalesGlobal(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiScalesGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiScalesGlobal`: GetAiScales200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiScalesGlobal`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiScalesGlobalRequest struct via the builder pattern


### Return type

[**GetAiScales200Response**](GetAiScales200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiScalesPublic

> Envelope GetAiScalesPublic(ctx).Execute()

Public (scale)

### Example

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
	resp, r, err := apiClient.AiAPI.GetAiScalesPublic(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiScalesPublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiScalesPublic`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiScalesPublic`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiScalesPublicRequest struct via the builder pattern


### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiScans

> GetAiScans200Response GetAiScans(ctx).Execute()

List scans



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiScans(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiScans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiScans`: GetAiScans200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiScans`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiScansRequest struct via the builder pattern


### Return type

[**GetAiScans200Response**](GetAiScans200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiScansByOwnerByName

> PostAiScans200Response GetAiScansByOwnerByName(ctx, owner, name).Execute()

Retrieve a scan



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.GetAiScansByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiScansByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiScansByOwnerByName`: PostAiScans200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiScansByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiScansByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiScans200Response**](PostAiScans200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiSigninSessions

> GetAiSigninSessions200Response GetAiSigninSessions(ctx).Execute()

List signin-sessions



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiSigninSessions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiSigninSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiSigninSessions`: GetAiSigninSessions200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiSigninSessions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiSigninSessionsRequest struct via the builder pattern


### Return type

[**GetAiSigninSessions200Response**](GetAiSigninSessions200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiSigninSessionsByOwnerByName

> PostAiSigninSessions200Response GetAiSigninSessionsByOwnerByName(ctx, owner, name).Execute()

Retrieve a session



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.GetAiSigninSessionsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiSigninSessionsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiSigninSessionsByOwnerByName`: PostAiSigninSessions200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiSigninSessionsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiSigninSessionsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiSigninSessions200Response**](PostAiSigninSessions200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiSigninSessionsDuplicated

> Envelope GetAiSigninSessionsDuplicated(ctx).Execute()

Duplicated (session)

### Example

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
	resp, r, err := apiClient.AiAPI.GetAiSigninSessionsDuplicated(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiSigninSessionsDuplicated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiSigninSessionsDuplicated`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiSigninSessionsDuplicated`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiSigninSessionsDuplicatedRequest struct via the builder pattern


### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiStores

> GetAiStores200Response GetAiStores(ctx).Execute()

List stores



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiStores(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiStores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiStores`: GetAiStores200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiStores`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiStoresRequest struct via the builder pattern


### Return type

[**GetAiStores200Response**](GetAiStores200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiStoresByOwnerByName

> PostAiStores200Response GetAiStoresByOwnerByName(ctx, owner, name).Execute()

Retrieve a store



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.GetAiStoresByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiStoresByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiStoresByOwnerByName`: PostAiStores200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiStoresByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiStoresByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiStores200Response**](PostAiStores200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiStoresGlobal

> GetAiStores200Response GetAiStoresGlobal(ctx).Execute()

List stores across tenants



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiStoresGlobal(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiStoresGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiStoresGlobal`: GetAiStores200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiStoresGlobal`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiStoresGlobalRequest struct via the builder pattern


### Return type

[**GetAiStores200Response**](GetAiStores200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiStoresNames

> Envelope GetAiStoresNames(ctx).Execute()

Names (store)

### Example

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
	resp, r, err := apiClient.AiAPI.GetAiStoresNames(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiStoresNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiStoresNames`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiStoresNames`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiStoresNamesRequest struct via the builder pattern


### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiStoresProviders

> Envelope GetAiStoresProviders(ctx).Execute()

Providers (store)

### Example

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
	resp, r, err := apiClient.AiAPI.GetAiStoresProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiStoresProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiStoresProviders`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiStoresProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiStoresProvidersRequest struct via the builder pattern


### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiSystem

> Envelope GetAiSystem(ctx).Execute()

System

### Example

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
	resp, r, err := apiClient.AiAPI.GetAiSystem(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiSystem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiSystem`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiSystem`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiSystemRequest struct via the builder pattern


### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiTasks

> GetAiTasks200Response GetAiTasks(ctx).Execute()

List tasks



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiTasks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiTasks`: GetAiTasks200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiTasks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiTasksRequest struct via the builder pattern


### Return type

[**GetAiTasks200Response**](GetAiTasks200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiTasksByOwnerByName

> PostAiTasks200Response GetAiTasksByOwnerByName(ctx, owner, name).Execute()

Retrieve a task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.GetAiTasksByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiTasksByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiTasksByOwnerByName`: PostAiTasks200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiTasksByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiTasksByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiTasks200Response**](PostAiTasks200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiTasksGlobal

> GetAiTasks200Response GetAiTasksGlobal(ctx).Execute()

List tasks across tenants



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiTasksGlobal(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiTasksGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiTasksGlobal`: GetAiTasks200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiTasksGlobal`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiTasksGlobalRequest struct via the builder pattern


### Return type

[**GetAiTasks200Response**](GetAiTasks200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiTemplates

> GetAiTemplates200Response GetAiTemplates(ctx).Execute()

List templates



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiTemplates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiTemplates`: GetAiTemplates200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiTemplatesRequest struct via the builder pattern


### Return type

[**GetAiTemplates200Response**](GetAiTemplates200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiTemplatesByOwnerByName

> PostAiTemplates200Response GetAiTemplatesByOwnerByName(ctx, owner, name).Execute()

Retrieve a template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.GetAiTemplatesByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiTemplatesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiTemplatesByOwnerByName`: PostAiTemplates200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiTemplatesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiTemplatesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiTemplates200Response**](PostAiTemplates200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiTrainingContribution

> Envelope GetAiTrainingContribution(ctx).Execute()

Training Contribution

### Example

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
	resp, r, err := apiClient.AiAPI.GetAiTrainingContribution(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiTrainingContribution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiTrainingContribution`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiTrainingContribution`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiTrainingContributionRequest struct via the builder pattern


### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiUsages

> GetAiUsages200Response GetAiUsages(ctx).Execute()

List usages



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiUsages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiUsages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiUsages`: GetAiUsages200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiUsages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiUsagesRequest struct via the builder pattern


### Return type

[**GetAiUsages200Response**](GetAiUsages200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiUsagesByUser

> Envelope GetAiUsagesByUser(ctx).Execute()

By User (usage)

### Example

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
	resp, r, err := apiClient.AiAPI.GetAiUsagesByUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiUsagesByUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiUsagesByUser`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiUsagesByUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiUsagesByUserRequest struct via the builder pattern


### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiUsagesCloud

> Envelope GetAiUsagesCloud(ctx).Execute()

Cloud (usage)

### Example

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
	resp, r, err := apiClient.AiAPI.GetAiUsagesCloud(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiUsagesCloud``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiUsagesCloud`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiUsagesCloud`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiUsagesCloudRequest struct via the builder pattern


### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiUsagesRange

> Envelope GetAiUsagesRange(ctx).Execute()

Range (usage)

### Example

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
	resp, r, err := apiClient.AiAPI.GetAiUsagesRange(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiUsagesRange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiUsagesRange`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiUsagesRange`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiUsagesRangeRequest struct via the builder pattern


### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiUsagesUserNames

> Envelope GetAiUsagesUserNames(ctx).Execute()

User Names (usage)

### Example

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
	resp, r, err := apiClient.AiAPI.GetAiUsagesUserNames(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiUsagesUserNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiUsagesUserNames`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiUsagesUserNames`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiUsagesUserNamesRequest struct via the builder pattern


### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiVectors

> GetAiVectors200Response GetAiVectors(ctx).Execute()

List vectors



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiVectors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiVectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiVectors`: GetAiVectors200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiVectors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiVectorsRequest struct via the builder pattern


### Return type

[**GetAiVectors200Response**](GetAiVectors200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiVectorsByOwnerByName

> PostAiVectors200Response GetAiVectorsByOwnerByName(ctx, owner, name).Execute()

Retrieve a vector



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.GetAiVectorsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiVectorsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiVectorsByOwnerByName`: PostAiVectors200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiVectorsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiVectorsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiVectors200Response**](PostAiVectors200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiVectorsGlobal

> GetAiVectors200Response GetAiVectorsGlobal(ctx).Execute()

List vectors across tenants



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiVectorsGlobal(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiVectorsGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiVectorsGlobal`: GetAiVectors200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiVectorsGlobal`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiVectorsGlobalRequest struct via the builder pattern


### Return type

[**GetAiVectors200Response**](GetAiVectors200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiVersion

> Envelope GetAiVersion(ctx).Execute()

Version

### Example

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
	resp, r, err := apiClient.AiAPI.GetAiVersion(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiVersion`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiVersionRequest struct via the builder pattern


### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiVideos

> GetAiVideos200Response GetAiVideos(ctx).Execute()

List videos



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiVideos(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiVideos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiVideos`: GetAiVideos200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiVideos`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiVideosRequest struct via the builder pattern


### Return type

[**GetAiVideos200Response**](GetAiVideos200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiVideosByOwnerByName

> PostAiVideos200Response GetAiVideosByOwnerByName(ctx, owner, name).Execute()

Retrieve a video



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.GetAiVideosByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiVideosByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiVideosByOwnerByName`: PostAiVideos200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiVideosByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiVideosByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiVideos200Response**](PostAiVideos200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiVideosGlobal

> GetAiVideos200Response GetAiVideosGlobal(ctx).Execute()

List videos across tenants



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiVideosGlobal(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiVideosGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiVideosGlobal`: GetAiVideos200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiVideosGlobal`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiVideosGlobalRequest struct via the builder pattern


### Return type

[**GetAiVideos200Response**](GetAiVideos200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiWorkflows

> GetAiWorkflows200Response GetAiWorkflows(ctx).Execute()

List workflows



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiWorkflows(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiWorkflows`: GetAiWorkflows200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiWorkflows`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiWorkflowsRequest struct via the builder pattern


### Return type

[**GetAiWorkflows200Response**](GetAiWorkflows200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiWorkflowsByOwnerByName

> PostAiWorkflows200Response GetAiWorkflowsByOwnerByName(ctx, owner, name).Execute()

Retrieve a workflow



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.GetAiWorkflowsByOwnerByName(context.Background(), owner, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiWorkflowsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiWorkflowsByOwnerByName`: PostAiWorkflows200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiWorkflowsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiWorkflowsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostAiWorkflows200Response**](PostAiWorkflows200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiWorkflowsGlobal

> GetAiWorkflows200Response GetAiWorkflowsGlobal(ctx).Execute()

List workflows across tenants



### Example

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
	resp, r, err := apiClient.AiAPI.GetAiWorkflowsGlobal(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.GetAiWorkflowsGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiWorkflowsGlobal`: GetAiWorkflows200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.GetAiWorkflowsGlobal`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiWorkflowsGlobalRequest struct via the builder pattern


### Return type

[**GetAiWorkflows200Response**](GetAiWorkflows200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAiArticlesByOwnerByName

> PostAiArticles200Response PatchAiArticlesByOwnerByName(ctx, owner, name).Body(body).Execute()

Update a article



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PatchAiArticlesByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PatchAiArticlesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAiArticlesByOwnerByName`: PostAiArticles200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PatchAiArticlesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAiArticlesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiArticles200Response**](PostAiArticles200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAiAssetsByOwnerByName

> PostAiAssets200Response PatchAiAssetsByOwnerByName(ctx, owner, name).Body(body).Execute()

Update a asset



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PatchAiAssetsByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PatchAiAssetsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAiAssetsByOwnerByName`: PostAiAssets200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PatchAiAssetsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAiAssetsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiAssets200Response**](PostAiAssets200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAiChatsByOwnerByName

> PostAiChats200Response PatchAiChatsByOwnerByName(ctx, owner, name).Body(body).Execute()

Update a chat



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PatchAiChatsByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PatchAiChatsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAiChatsByOwnerByName`: PostAiChats200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PatchAiChatsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAiChatsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiChats200Response**](PostAiChats200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAiDeploymentsByOwnerByName

> PostAiDeployments200Response PatchAiDeploymentsByOwnerByName(ctx, owner, name).Body(body).Execute()

Update a application



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PatchAiDeploymentsByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PatchAiDeploymentsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAiDeploymentsByOwnerByName`: PostAiDeployments200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PatchAiDeploymentsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAiDeploymentsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiDeployments200Response**](PostAiDeployments200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAiFilesByOwnerByName

> PostAiFiles200Response PatchAiFilesByOwnerByName(ctx, owner, name).Body(body).Execute()

Update a file



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PatchAiFilesByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PatchAiFilesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAiFilesByOwnerByName`: PostAiFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PatchAiFilesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAiFilesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiFiles200Response**](PostAiFiles200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAiFormsByOwnerByName

> PostAiForms200Response PatchAiFormsByOwnerByName(ctx, owner, name).Body(body).Execute()

Update a form



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PatchAiFormsByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PatchAiFormsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAiFormsByOwnerByName`: PostAiForms200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PatchAiFormsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAiFormsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiForms200Response**](PostAiForms200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAiGraphsByOwnerByName

> PostAiGraphs200Response PatchAiGraphsByOwnerByName(ctx, owner, name).Body(body).Execute()

Update a graph



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PatchAiGraphsByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PatchAiGraphsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAiGraphsByOwnerByName`: PostAiGraphs200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PatchAiGraphsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAiGraphsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiGraphs200Response**](PostAiGraphs200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAiMessagesByOwnerByName

> PostAiMessages200Response PatchAiMessagesByOwnerByName(ctx, owner, name).Body(body).Execute()

Update a message



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PatchAiMessagesByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PatchAiMessagesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAiMessagesByOwnerByName`: PostAiMessages200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PatchAiMessagesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAiMessagesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiMessages200Response**](PostAiMessages200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAiNodesByOwnerByName

> PostAiNodes200Response PatchAiNodesByOwnerByName(ctx, owner, name).Body(body).Execute()

Update a node



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PatchAiNodesByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PatchAiNodesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAiNodesByOwnerByName`: PostAiNodes200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PatchAiNodesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAiNodesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiNodes200Response**](PostAiNodes200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAiPreferences

> Envelope PatchAiPreferences(ctx).Body(body).Execute()

Preferences

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PatchAiPreferences(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PatchAiPreferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAiPreferences`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PatchAiPreferences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchAiPreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAiProvidersByOwnerByName

> PostAiProviders200Response PatchAiProvidersByOwnerByName(ctx, owner, name).Body(body).Execute()

Update a provider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PatchAiProvidersByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PatchAiProvidersByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAiProvidersByOwnerByName`: PostAiProviders200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PatchAiProvidersByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAiProvidersByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiProviders200Response**](PostAiProviders200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAiRecordsByOwnerByName

> PostAiRecords200Response PatchAiRecordsByOwnerByName(ctx, owner, name).Body(body).Execute()

Update a record



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PatchAiRecordsByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PatchAiRecordsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAiRecordsByOwnerByName`: PostAiRecords200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PatchAiRecordsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAiRecordsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiRecords200Response**](PostAiRecords200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAiRemoteConnectionsByOwnerByName

> PostAiRemoteConnections200Response PatchAiRemoteConnectionsByOwnerByName(ctx, owner, name).Body(body).Execute()

Update a connection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PatchAiRemoteConnectionsByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PatchAiRemoteConnectionsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAiRemoteConnectionsByOwnerByName`: PostAiRemoteConnections200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PatchAiRemoteConnectionsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAiRemoteConnectionsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiRemoteConnections200Response**](PostAiRemoteConnections200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAiRoutesByOwnerByName

> PostAiRoutes200Response PatchAiRoutesByOwnerByName(ctx, owner, name).Body(body).Execute()

Update a model-route



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PatchAiRoutesByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PatchAiRoutesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAiRoutesByOwnerByName`: PostAiRoutes200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PatchAiRoutesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAiRoutesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiRoutes200Response**](PostAiRoutes200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAiScalesByOwnerByName

> PostAiScales200Response PatchAiScalesByOwnerByName(ctx, owner, name).Body(body).Execute()

Update a scale



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PatchAiScalesByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PatchAiScalesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAiScalesByOwnerByName`: PostAiScales200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PatchAiScalesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAiScalesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiScales200Response**](PostAiScales200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAiScansByOwnerByName

> PostAiScans200Response PatchAiScansByOwnerByName(ctx, owner, name).Body(body).Execute()

Update a scan



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PatchAiScansByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PatchAiScansByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAiScansByOwnerByName`: PostAiScans200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PatchAiScansByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAiScansByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiScans200Response**](PostAiScans200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAiSigninSessionsByOwnerByName

> PostAiSigninSessions200Response PatchAiSigninSessionsByOwnerByName(ctx, owner, name).Body(body).Execute()

Update a session



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PatchAiSigninSessionsByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PatchAiSigninSessionsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAiSigninSessionsByOwnerByName`: PostAiSigninSessions200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PatchAiSigninSessionsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAiSigninSessionsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiSigninSessions200Response**](PostAiSigninSessions200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAiStoresByOwnerByName

> PostAiStores200Response PatchAiStoresByOwnerByName(ctx, owner, name).Body(body).Execute()

Update a store



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PatchAiStoresByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PatchAiStoresByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAiStoresByOwnerByName`: PostAiStores200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PatchAiStoresByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAiStoresByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiStores200Response**](PostAiStores200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAiTasksByOwnerByName

> PostAiTasks200Response PatchAiTasksByOwnerByName(ctx, owner, name).Body(body).Execute()

Update a task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PatchAiTasksByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PatchAiTasksByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAiTasksByOwnerByName`: PostAiTasks200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PatchAiTasksByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAiTasksByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiTasks200Response**](PostAiTasks200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAiTemplatesByOwnerByName

> PostAiTemplates200Response PatchAiTemplatesByOwnerByName(ctx, owner, name).Body(body).Execute()

Update a template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PatchAiTemplatesByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PatchAiTemplatesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAiTemplatesByOwnerByName`: PostAiTemplates200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PatchAiTemplatesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAiTemplatesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiTemplates200Response**](PostAiTemplates200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAiTrainingContribution

> Envelope PatchAiTrainingContribution(ctx).Body(body).Execute()

Training Contribution

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PatchAiTrainingContribution(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PatchAiTrainingContribution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAiTrainingContribution`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PatchAiTrainingContribution`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchAiTrainingContributionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAiTreeFilesByOwnerByName

> PostAiTreeFiles200Response PatchAiTreeFilesByOwnerByName(ctx, owner, name).Body(body).Execute()

Update a tree-file



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PatchAiTreeFilesByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PatchAiTreeFilesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAiTreeFilesByOwnerByName`: PostAiTreeFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PatchAiTreeFilesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAiTreeFilesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiTreeFiles200Response**](PostAiTreeFiles200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAiVectorsByOwnerByName

> PostAiVectors200Response PatchAiVectorsByOwnerByName(ctx, owner, name).Body(body).Execute()

Update a vector



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PatchAiVectorsByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PatchAiVectorsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAiVectorsByOwnerByName`: PostAiVectors200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PatchAiVectorsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAiVectorsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiVectors200Response**](PostAiVectors200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAiVideosByOwnerByName

> PostAiVideos200Response PatchAiVideosByOwnerByName(ctx, owner, name).Body(body).Execute()

Update a video



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PatchAiVideosByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PatchAiVideosByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAiVideosByOwnerByName`: PostAiVideos200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PatchAiVideosByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAiVideosByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiVideos200Response**](PostAiVideos200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAiWorkflowsByOwnerByName

> PostAiWorkflows200Response PatchAiWorkflowsByOwnerByName(ctx, owner, name).Body(body).Execute()

Update a workflow



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PatchAiWorkflowsByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PatchAiWorkflowsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAiWorkflowsByOwnerByName`: PostAiWorkflows200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PatchAiWorkflowsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAiWorkflowsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiWorkflows200Response**](PostAiWorkflows200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiArticles

> PostAiArticles200Response PostAiArticles(ctx).Body(body).Execute()

Create a article



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiArticles(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiArticles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiArticles`: PostAiArticles200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiArticles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiArticlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiArticles200Response**](PostAiArticles200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiAssets

> PostAiAssets200Response PostAiAssets(ctx).Body(body).Execute()

Create a asset



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiAssets(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiAssets`: PostAiAssets200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiAssets200Response**](PostAiAssets200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiAssetsByOwnerByNameScan

> Envelope PostAiAssetsByOwnerByNameScan(ctx, owner, name).Body(body).Execute()

Scan (asset)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiAssetsByOwnerByNameScan(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiAssetsByOwnerByNameScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiAssetsByOwnerByNameScan`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiAssetsByOwnerByNameScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAiAssetsByOwnerByNameScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiAssetsScan

> Envelope PostAiAssetsScan(ctx).Body(body).Execute()

Scan (asset)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiAssetsScan(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiAssetsScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiAssetsScan`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiAssetsScan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiAssetsScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiChats

> PostAiChats200Response PostAiChats(ctx).Body(body).Execute()

Create a chat



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiChats(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiChats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiChats`: PostAiChats200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiChats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiChatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiChats200Response**](PostAiChats200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiConnections

> PostAiConnections(ctx).Execute()

Connects (or reconnects) a third-party AI account for the org by sealing the supplied key into KMS and upserting the org's provider row.



### Example

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
	r, err := apiClient.AiAPI.PostAiConnections(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAiConnectionsRequest struct via the builder pattern


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


## PostAiConnectionsByProvider

> PostAiConnectionsByProvider(ctx).Execute()

Disconnects a third-party AI account: it deactivates the org's row so completion resolution falls back to the global Hanzo account (no BYO), and best-effort tombstones the sealed secret.



### Example

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
	r, err := apiClient.AiAPI.PostAiConnectionsByProvider(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiConnectionsByProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAiConnectionsByProviderRequest struct via the builder pattern


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


## PostAiDeployments

> PostAiDeployments200Response PostAiDeployments(ctx).Body(body).Execute()

Create a application



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiDeployments(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiDeployments`: PostAiDeployments200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiDeployments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiDeployments200Response**](PostAiDeployments200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiDeploymentsByOwnerByNameDeploy

> Envelope PostAiDeploymentsByOwnerByNameDeploy(ctx, owner, name).Body(body).Execute()

Deploy (application)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiDeploymentsByOwnerByNameDeploy(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiDeploymentsByOwnerByNameDeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiDeploymentsByOwnerByNameDeploy`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiDeploymentsByOwnerByNameDeploy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAiDeploymentsByOwnerByNameDeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiDeploymentsByOwnerByNameUndeploy

> Envelope PostAiDeploymentsByOwnerByNameUndeploy(ctx, owner, name).Body(body).Execute()

Undeploy (application)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiDeploymentsByOwnerByNameUndeploy(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiDeploymentsByOwnerByNameUndeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiDeploymentsByOwnerByNameUndeploy`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiDeploymentsByOwnerByNameUndeploy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAiDeploymentsByOwnerByNameUndeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiFiles

> PostAiFiles200Response PostAiFiles(ctx).Body(body).Execute()

Create a file



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiFiles(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiFiles`: PostAiFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiFiles200Response**](PostAiFiles200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiFilesActivate

> Envelope PostAiFilesActivate(ctx).Body(body).Execute()

Activate (file)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiFilesActivate(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiFilesActivate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiFilesActivate`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiFilesActivate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiFilesActivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiFilesByOwnerByNameVectors

> Envelope PostAiFilesByOwnerByNameVectors(ctx, owner, name).Body(body).Execute()

Vectors (file)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiFilesByOwnerByNameVectors(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiFilesByOwnerByNameVectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiFilesByOwnerByNameVectors`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiFilesByOwnerByNameVectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAiFilesByOwnerByNameVectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiFilesUpload

> Envelope PostAiFilesUpload(ctx).Body(body).Execute()

Upload (file)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiFilesUpload(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiFilesUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiFilesUpload`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiFilesUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiFilesUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiForms

> PostAiForms200Response PostAiForms(ctx).Body(body).Execute()

Create a form



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiForms(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiForms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiForms`: PostAiForms200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiForms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiFormsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiForms200Response**](PostAiForms200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiGraphs

> PostAiGraphs200Response PostAiGraphs(ctx).Body(body).Execute()

Create a graph



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiGraphs(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiGraphs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiGraphs`: PostAiGraphs200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiGraphs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiGraphsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiGraphs200Response**](PostAiGraphs200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiMessages

> PostAiMessages200Response PostAiMessages(ctx).Body(body).Execute()

Create a message



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiMessages(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiMessages`: PostAiMessages200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiMessages200Response**](PostAiMessages200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiNodes

> PostAiNodes200Response PostAiNodes(ctx).Body(body).Execute()

Create a node



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiNodes(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiNodes`: PostAiNodes200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiNodes200Response**](PostAiNodes200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiNodesByOwnerByNameTunnel

> Envelope PostAiNodesByOwnerByNameTunnel(ctx, owner, name).Body(body).Execute()

Tunnel (node)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiNodesByOwnerByNameTunnel(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiNodesByOwnerByNameTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiNodesByOwnerByNameTunnel`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiNodesByOwnerByNameTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAiNodesByOwnerByNameTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiProviders

> PostAiProviders200Response PostAiProviders(ctx).Body(body).Execute()

Create a provider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiProviders(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiProviders`: PostAiProviders200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiProviders200Response**](PostAiProviders200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiProvidersMcpTools

> Envelope PostAiProvidersMcpTools(ctx).Body(body).Execute()

Mcp Tools (provider)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiProvidersMcpTools(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiProvidersMcpTools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiProvidersMcpTools`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiProvidersMcpTools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiProvidersMcpToolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiRecords

> PostAiRecords200Response PostAiRecords(ctx).Body(body).Execute()

Create a record



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiRecords(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiRecords`: PostAiRecords200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiRecords200Response**](PostAiRecords200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiRecordsBatch

> Envelope PostAiRecordsBatch(ctx).Body(body).Execute()

Batch (record)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiRecordsBatch(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiRecordsBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiRecordsBatch`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiRecordsBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiRecordsBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiRecordsCommit

> Envelope PostAiRecordsCommit(ctx).Body(body).Execute()

Commit (record)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiRecordsCommit(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiRecordsCommit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiRecordsCommit`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiRecordsCommit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiRecordsCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiRecordsCommitSecond

> Envelope PostAiRecordsCommitSecond(ctx).Body(body).Execute()

Commit Second (record)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiRecordsCommitSecond(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiRecordsCommitSecond``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiRecordsCommitSecond`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiRecordsCommitSecond`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiRecordsCommitSecondRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiRemoteConnections

> PostAiRemoteConnections200Response PostAiRemoteConnections(ctx).Body(body).Execute()

Create a connection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiRemoteConnections(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiRemoteConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiRemoteConnections`: PostAiRemoteConnections200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiRemoteConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiRemoteConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiRemoteConnections200Response**](PostAiRemoteConnections200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiRemoteConnectionsByOwnerByNameStart

> Envelope PostAiRemoteConnectionsByOwnerByNameStart(ctx, owner, name).Body(body).Execute()

Start (connection)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiRemoteConnectionsByOwnerByNameStart(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiRemoteConnectionsByOwnerByNameStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiRemoteConnectionsByOwnerByNameStart`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiRemoteConnectionsByOwnerByNameStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAiRemoteConnectionsByOwnerByNameStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiRemoteConnectionsByOwnerByNameStop

> Envelope PostAiRemoteConnectionsByOwnerByNameStop(ctx, owner, name).Body(body).Execute()

Stop (connection)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiRemoteConnectionsByOwnerByNameStop(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiRemoteConnectionsByOwnerByNameStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiRemoteConnectionsByOwnerByNameStop`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiRemoteConnectionsByOwnerByNameStop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAiRemoteConnectionsByOwnerByNameStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiRoutes

> PostAiRoutes200Response PostAiRoutes(ctx).Body(body).Execute()

Create a model-route



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiRoutes(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiRoutes`: PostAiRoutes200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiRoutes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiRoutes200Response**](PostAiRoutes200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiScales

> PostAiScales200Response PostAiScales(ctx).Body(body).Execute()

Create a scale



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiScales(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiScales``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiScales`: PostAiScales200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiScales`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiScalesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiScales200Response**](PostAiScales200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiScans

> PostAiScans200Response PostAiScans(ctx).Body(body).Execute()

Create a scan



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiScans(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiScans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiScans`: PostAiScans200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiScans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiScansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiScans200Response**](PostAiScans200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiSignin

> Envelope PostAiSignin(ctx).Body(body).Execute()

Signin

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiSignin(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiSignin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiSignin`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiSignin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiSigninRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiSigninSessions

> PostAiSigninSessions200Response PostAiSigninSessions(ctx).Body(body).Execute()

Create a session



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiSigninSessions(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiSigninSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiSigninSessions`: PostAiSigninSessions200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiSigninSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiSigninSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiSigninSessions200Response**](PostAiSigninSessions200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiSignout

> Envelope PostAiSignout(ctx).Body(body).Execute()

Signout

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiSignout(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiSignout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiSignout`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiSignout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiSignoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiStores

> PostAiStores200Response PostAiStores(ctx).Body(body).Execute()

Create a store



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiStores(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiStores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiStores`: PostAiStores200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiStores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiStores200Response**](PostAiStores200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiStoresByOwnerByNameVectors

> Envelope PostAiStoresByOwnerByNameVectors(ctx, owner, name).Body(body).Execute()

Vectors (store)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiStoresByOwnerByNameVectors(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiStoresByOwnerByNameVectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiStoresByOwnerByNameVectors`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiStoresByOwnerByNameVectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAiStoresByOwnerByNameVectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiTasks

> PostAiTasks200Response PostAiTasks(ctx).Body(body).Execute()

Create a task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiTasks(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiTasks`: PostAiTasks200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiTasks200Response**](PostAiTasks200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiTasksByOwnerByNameAnalyze

> Envelope PostAiTasksByOwnerByNameAnalyze(ctx, owner, name).Body(body).Execute()

Analyze (task)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiTasksByOwnerByNameAnalyze(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiTasksByOwnerByNameAnalyze``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiTasksByOwnerByNameAnalyze`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiTasksByOwnerByNameAnalyze`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAiTasksByOwnerByNameAnalyzeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiTasksByOwnerByNameDocument

> Envelope PostAiTasksByOwnerByNameDocument(ctx, owner, name).Body(body).Execute()

Document (task)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiTasksByOwnerByNameDocument(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiTasksByOwnerByNameDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiTasksByOwnerByNameDocument`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiTasksByOwnerByNameDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAiTasksByOwnerByNameDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiTemplates

> PostAiTemplates200Response PostAiTemplates(ctx).Body(body).Execute()

Create a template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiTemplates(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiTemplates`: PostAiTemplates200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiTemplates200Response**](PostAiTemplates200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiTreeFiles

> PostAiTreeFiles200Response PostAiTreeFiles(ctx).Body(body).Execute()

Create a tree-file



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiTreeFiles(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiTreeFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiTreeFiles`: PostAiTreeFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiTreeFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiTreeFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiTreeFiles200Response**](PostAiTreeFiles200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiVectors

> PostAiVectors200Response PostAiVectors(ctx).Body(body).Execute()

Create a vector



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiVectors(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiVectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiVectors`: PostAiVectors200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiVectors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiVectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiVectors200Response**](PostAiVectors200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiVideos

> PostAiVideos200Response PostAiVideos(ctx).Body(body).Execute()

Create a video



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiVideos(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiVideos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiVideos`: PostAiVideos200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiVideos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiVideosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiVideos200Response**](PostAiVideos200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiVideosUpload

> Envelope PostAiVideosUpload(ctx).Body(body).Execute()

Upload (video)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiVideosUpload(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiVideosUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiVideosUpload`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiVideosUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiVideosUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiWorkflows

> PostAiWorkflows200Response PostAiWorkflows(ctx).Body(body).Execute()

Create a workflow



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PostAiWorkflows(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PostAiWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiWorkflows`: PostAiWorkflows200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PostAiWorkflows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiWorkflows200Response**](PostAiWorkflows200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAiArticlesByOwnerByName

> PostAiArticles200Response PutAiArticlesByOwnerByName(ctx, owner, name).Body(body).Execute()

Replace a article



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PutAiArticlesByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PutAiArticlesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAiArticlesByOwnerByName`: PostAiArticles200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PutAiArticlesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAiArticlesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiArticles200Response**](PostAiArticles200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAiAssetsByOwnerByName

> PostAiAssets200Response PutAiAssetsByOwnerByName(ctx, owner, name).Body(body).Execute()

Replace a asset



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PutAiAssetsByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PutAiAssetsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAiAssetsByOwnerByName`: PostAiAssets200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PutAiAssetsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAiAssetsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiAssets200Response**](PostAiAssets200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAiChatsByOwnerByName

> PostAiChats200Response PutAiChatsByOwnerByName(ctx, owner, name).Body(body).Execute()

Replace a chat



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PutAiChatsByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PutAiChatsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAiChatsByOwnerByName`: PostAiChats200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PutAiChatsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAiChatsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiChats200Response**](PostAiChats200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAiDeploymentsByOwnerByName

> PostAiDeployments200Response PutAiDeploymentsByOwnerByName(ctx, owner, name).Body(body).Execute()

Replace a application



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PutAiDeploymentsByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PutAiDeploymentsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAiDeploymentsByOwnerByName`: PostAiDeployments200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PutAiDeploymentsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAiDeploymentsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiDeployments200Response**](PostAiDeployments200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAiFilesByOwnerByName

> PostAiFiles200Response PutAiFilesByOwnerByName(ctx, owner, name).Body(body).Execute()

Replace a file



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PutAiFilesByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PutAiFilesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAiFilesByOwnerByName`: PostAiFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PutAiFilesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAiFilesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiFiles200Response**](PostAiFiles200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAiFormsByOwnerByName

> PostAiForms200Response PutAiFormsByOwnerByName(ctx, owner, name).Body(body).Execute()

Replace a form



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PutAiFormsByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PutAiFormsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAiFormsByOwnerByName`: PostAiForms200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PutAiFormsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAiFormsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiForms200Response**](PostAiForms200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAiGraphsByOwnerByName

> PostAiGraphs200Response PutAiGraphsByOwnerByName(ctx, owner, name).Body(body).Execute()

Replace a graph



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PutAiGraphsByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PutAiGraphsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAiGraphsByOwnerByName`: PostAiGraphs200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PutAiGraphsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAiGraphsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiGraphs200Response**](PostAiGraphs200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAiMessagesByOwnerByName

> PostAiMessages200Response PutAiMessagesByOwnerByName(ctx, owner, name).Body(body).Execute()

Replace a message



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PutAiMessagesByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PutAiMessagesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAiMessagesByOwnerByName`: PostAiMessages200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PutAiMessagesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAiMessagesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiMessages200Response**](PostAiMessages200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAiNodesByOwnerByName

> PostAiNodes200Response PutAiNodesByOwnerByName(ctx, owner, name).Body(body).Execute()

Replace a node



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PutAiNodesByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PutAiNodesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAiNodesByOwnerByName`: PostAiNodes200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PutAiNodesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAiNodesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiNodes200Response**](PostAiNodes200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAiPreferences

> Envelope PutAiPreferences(ctx).Body(body).Execute()

Preferences

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PutAiPreferences(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PutAiPreferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAiPreferences`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PutAiPreferences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutAiPreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAiProvidersByOwnerByName

> PostAiProviders200Response PutAiProvidersByOwnerByName(ctx, owner, name).Body(body).Execute()

Replace a provider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PutAiProvidersByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PutAiProvidersByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAiProvidersByOwnerByName`: PostAiProviders200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PutAiProvidersByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAiProvidersByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiProviders200Response**](PostAiProviders200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAiRecordsByOwnerByName

> PostAiRecords200Response PutAiRecordsByOwnerByName(ctx, owner, name).Body(body).Execute()

Replace a record



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PutAiRecordsByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PutAiRecordsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAiRecordsByOwnerByName`: PostAiRecords200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PutAiRecordsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAiRecordsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiRecords200Response**](PostAiRecords200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAiRemoteConnectionsByOwnerByName

> PostAiRemoteConnections200Response PutAiRemoteConnectionsByOwnerByName(ctx, owner, name).Body(body).Execute()

Replace a connection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PutAiRemoteConnectionsByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PutAiRemoteConnectionsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAiRemoteConnectionsByOwnerByName`: PostAiRemoteConnections200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PutAiRemoteConnectionsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAiRemoteConnectionsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiRemoteConnections200Response**](PostAiRemoteConnections200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAiRoutesByOwnerByName

> PostAiRoutes200Response PutAiRoutesByOwnerByName(ctx, owner, name).Body(body).Execute()

Replace a model-route



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PutAiRoutesByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PutAiRoutesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAiRoutesByOwnerByName`: PostAiRoutes200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PutAiRoutesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAiRoutesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiRoutes200Response**](PostAiRoutes200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAiScalesByOwnerByName

> PostAiScales200Response PutAiScalesByOwnerByName(ctx, owner, name).Body(body).Execute()

Replace a scale



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PutAiScalesByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PutAiScalesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAiScalesByOwnerByName`: PostAiScales200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PutAiScalesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAiScalesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiScales200Response**](PostAiScales200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAiScansByOwnerByName

> PostAiScans200Response PutAiScansByOwnerByName(ctx, owner, name).Body(body).Execute()

Replace a scan



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PutAiScansByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PutAiScansByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAiScansByOwnerByName`: PostAiScans200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PutAiScansByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAiScansByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiScans200Response**](PostAiScans200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAiSigninSessionsByOwnerByName

> PostAiSigninSessions200Response PutAiSigninSessionsByOwnerByName(ctx, owner, name).Body(body).Execute()

Replace a session



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PutAiSigninSessionsByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PutAiSigninSessionsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAiSigninSessionsByOwnerByName`: PostAiSigninSessions200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PutAiSigninSessionsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAiSigninSessionsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiSigninSessions200Response**](PostAiSigninSessions200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAiStoresByOwnerByName

> PostAiStores200Response PutAiStoresByOwnerByName(ctx, owner, name).Body(body).Execute()

Replace a store



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PutAiStoresByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PutAiStoresByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAiStoresByOwnerByName`: PostAiStores200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PutAiStoresByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAiStoresByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiStores200Response**](PostAiStores200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAiTasksByOwnerByName

> PostAiTasks200Response PutAiTasksByOwnerByName(ctx, owner, name).Body(body).Execute()

Replace a task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PutAiTasksByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PutAiTasksByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAiTasksByOwnerByName`: PostAiTasks200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PutAiTasksByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAiTasksByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiTasks200Response**](PostAiTasks200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAiTemplatesByOwnerByName

> PostAiTemplates200Response PutAiTemplatesByOwnerByName(ctx, owner, name).Body(body).Execute()

Replace a template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PutAiTemplatesByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PutAiTemplatesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAiTemplatesByOwnerByName`: PostAiTemplates200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PutAiTemplatesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAiTemplatesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiTemplates200Response**](PostAiTemplates200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAiTrainingContribution

> Envelope PutAiTrainingContribution(ctx).Body(body).Execute()

Training Contribution

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PutAiTrainingContribution(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PutAiTrainingContribution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAiTrainingContribution`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PutAiTrainingContribution`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutAiTrainingContributionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAiTreeFilesByOwnerByName

> PostAiTreeFiles200Response PutAiTreeFilesByOwnerByName(ctx, owner, name).Body(body).Execute()

Replace a tree-file



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PutAiTreeFilesByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PutAiTreeFilesByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAiTreeFilesByOwnerByName`: PostAiTreeFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PutAiTreeFilesByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAiTreeFilesByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiTreeFiles200Response**](PostAiTreeFiles200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAiVectorsByOwnerByName

> PostAiVectors200Response PutAiVectorsByOwnerByName(ctx, owner, name).Body(body).Execute()

Replace a vector



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PutAiVectorsByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PutAiVectorsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAiVectorsByOwnerByName`: PostAiVectors200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PutAiVectorsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAiVectorsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiVectors200Response**](PostAiVectors200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAiVideosByOwnerByName

> PostAiVideos200Response PutAiVideosByOwnerByName(ctx, owner, name).Body(body).Execute()

Replace a video



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PutAiVideosByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PutAiVideosByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAiVideosByOwnerByName`: PostAiVideos200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PutAiVideosByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAiVideosByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiVideos200Response**](PostAiVideos200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAiWorkflowsByOwnerByName

> PostAiWorkflows200Response PutAiWorkflowsByOwnerByName(ctx, owner, name).Body(body).Execute()

Replace a workflow



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	owner := "owner_example" // string | Owning organization.
	name := "name_example" // string | Resource name, unique within the owner.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiAPI.PutAiWorkflowsByOwnerByName(context.Background(), owner, name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiAPI.PutAiWorkflowsByOwnerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAiWorkflowsByOwnerByName`: PostAiWorkflows200Response
	fmt.Fprintf(os.Stdout, "Response from `AiAPI.PutAiWorkflowsByOwnerByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owning organization. | 
**name** | **string** | Resource name, unique within the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAiWorkflowsByOwnerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**PostAiWorkflows200Response**](PostAiWorkflows200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

