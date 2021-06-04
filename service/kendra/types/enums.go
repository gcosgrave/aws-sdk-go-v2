// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AdditionalResultAttributeValueType string

// Enum values for AdditionalResultAttributeValueType
const (
	AdditionalResultAttributeValueTypeTextWithHighlightsValue AdditionalResultAttributeValueType = "TEXT_WITH_HIGHLIGHTS_VALUE"
)

// Values returns all known values for AdditionalResultAttributeValueType. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (AdditionalResultAttributeValueType) Values() []AdditionalResultAttributeValueType {
	return []AdditionalResultAttributeValueType{
		"TEXT_WITH_HIGHLIGHTS_VALUE",
	}
}

type ConfluenceAttachmentFieldName string

// Enum values for ConfluenceAttachmentFieldName
const (
	ConfluenceAttachmentFieldNameAuthor      ConfluenceAttachmentFieldName = "AUTHOR"
	ConfluenceAttachmentFieldNameContentType ConfluenceAttachmentFieldName = "CONTENT_TYPE"
	ConfluenceAttachmentFieldNameCreatedDate ConfluenceAttachmentFieldName = "CREATED_DATE"
	ConfluenceAttachmentFieldNameDisplayUrl  ConfluenceAttachmentFieldName = "DISPLAY_URL"
	ConfluenceAttachmentFieldNameFileSize    ConfluenceAttachmentFieldName = "FILE_SIZE"
	ConfluenceAttachmentFieldNameItemType    ConfluenceAttachmentFieldName = "ITEM_TYPE"
	ConfluenceAttachmentFieldNameParentId    ConfluenceAttachmentFieldName = "PARENT_ID"
	ConfluenceAttachmentFieldNameSpaceKey    ConfluenceAttachmentFieldName = "SPACE_KEY"
	ConfluenceAttachmentFieldNameSpaceName   ConfluenceAttachmentFieldName = "SPACE_NAME"
	ConfluenceAttachmentFieldNameUrl         ConfluenceAttachmentFieldName = "URL"
	ConfluenceAttachmentFieldNameVersion     ConfluenceAttachmentFieldName = "VERSION"
)

// Values returns all known values for ConfluenceAttachmentFieldName. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ConfluenceAttachmentFieldName) Values() []ConfluenceAttachmentFieldName {
	return []ConfluenceAttachmentFieldName{
		"AUTHOR",
		"CONTENT_TYPE",
		"CREATED_DATE",
		"DISPLAY_URL",
		"FILE_SIZE",
		"ITEM_TYPE",
		"PARENT_ID",
		"SPACE_KEY",
		"SPACE_NAME",
		"URL",
		"VERSION",
	}
}

type ConfluenceBlogFieldName string

// Enum values for ConfluenceBlogFieldName
const (
	ConfluenceBlogFieldNameAuthor      ConfluenceBlogFieldName = "AUTHOR"
	ConfluenceBlogFieldNameDisplayUrl  ConfluenceBlogFieldName = "DISPLAY_URL"
	ConfluenceBlogFieldNameItemType    ConfluenceBlogFieldName = "ITEM_TYPE"
	ConfluenceBlogFieldNameLabels      ConfluenceBlogFieldName = "LABELS"
	ConfluenceBlogFieldNamePublishDate ConfluenceBlogFieldName = "PUBLISH_DATE"
	ConfluenceBlogFieldNameSpaceKey    ConfluenceBlogFieldName = "SPACE_KEY"
	ConfluenceBlogFieldNameSpaceName   ConfluenceBlogFieldName = "SPACE_NAME"
	ConfluenceBlogFieldNameUrl         ConfluenceBlogFieldName = "URL"
	ConfluenceBlogFieldNameVersion     ConfluenceBlogFieldName = "VERSION"
)

// Values returns all known values for ConfluenceBlogFieldName. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConfluenceBlogFieldName) Values() []ConfluenceBlogFieldName {
	return []ConfluenceBlogFieldName{
		"AUTHOR",
		"DISPLAY_URL",
		"ITEM_TYPE",
		"LABELS",
		"PUBLISH_DATE",
		"SPACE_KEY",
		"SPACE_NAME",
		"URL",
		"VERSION",
	}
}

type ConfluencePageFieldName string

// Enum values for ConfluencePageFieldName
const (
	ConfluencePageFieldNameAuthor        ConfluencePageFieldName = "AUTHOR"
	ConfluencePageFieldNameContentStatus ConfluencePageFieldName = "CONTENT_STATUS"
	ConfluencePageFieldNameCreatedDate   ConfluencePageFieldName = "CREATED_DATE"
	ConfluencePageFieldNameDisplayUrl    ConfluencePageFieldName = "DISPLAY_URL"
	ConfluencePageFieldNameItemType      ConfluencePageFieldName = "ITEM_TYPE"
	ConfluencePageFieldNameLabels        ConfluencePageFieldName = "LABELS"
	ConfluencePageFieldNameModifiedDate  ConfluencePageFieldName = "MODIFIED_DATE"
	ConfluencePageFieldNameParentId      ConfluencePageFieldName = "PARENT_ID"
	ConfluencePageFieldNameSpaceKey      ConfluencePageFieldName = "SPACE_KEY"
	ConfluencePageFieldNameSpaceName     ConfluencePageFieldName = "SPACE_NAME"
	ConfluencePageFieldNameUrl           ConfluencePageFieldName = "URL"
	ConfluencePageFieldNameVersion       ConfluencePageFieldName = "VERSION"
)

// Values returns all known values for ConfluencePageFieldName. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConfluencePageFieldName) Values() []ConfluencePageFieldName {
	return []ConfluencePageFieldName{
		"AUTHOR",
		"CONTENT_STATUS",
		"CREATED_DATE",
		"DISPLAY_URL",
		"ITEM_TYPE",
		"LABELS",
		"MODIFIED_DATE",
		"PARENT_ID",
		"SPACE_KEY",
		"SPACE_NAME",
		"URL",
		"VERSION",
	}
}

type ConfluenceSpaceFieldName string

// Enum values for ConfluenceSpaceFieldName
const (
	ConfluenceSpaceFieldNameDisplayUrl ConfluenceSpaceFieldName = "DISPLAY_URL"
	ConfluenceSpaceFieldNameItemType   ConfluenceSpaceFieldName = "ITEM_TYPE"
	ConfluenceSpaceFieldNameSpaceKey   ConfluenceSpaceFieldName = "SPACE_KEY"
	ConfluenceSpaceFieldNameUrl        ConfluenceSpaceFieldName = "URL"
)

// Values returns all known values for ConfluenceSpaceFieldName. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConfluenceSpaceFieldName) Values() []ConfluenceSpaceFieldName {
	return []ConfluenceSpaceFieldName{
		"DISPLAY_URL",
		"ITEM_TYPE",
		"SPACE_KEY",
		"URL",
	}
}

type ConfluenceVersion string

// Enum values for ConfluenceVersion
const (
	ConfluenceVersionCloud  ConfluenceVersion = "CLOUD"
	ConfluenceVersionServer ConfluenceVersion = "SERVER"
)

// Values returns all known values for ConfluenceVersion. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConfluenceVersion) Values() []ConfluenceVersion {
	return []ConfluenceVersion{
		"CLOUD",
		"SERVER",
	}
}

type ContentType string

// Enum values for ContentType
const (
	ContentTypePdf       ContentType = "PDF"
	ContentTypeHtml      ContentType = "HTML"
	ContentTypeMsWord    ContentType = "MS_WORD"
	ContentTypePlainText ContentType = "PLAIN_TEXT"
	ContentTypePpt       ContentType = "PPT"
)

// Values returns all known values for ContentType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ContentType) Values() []ContentType {
	return []ContentType{
		"PDF",
		"HTML",
		"MS_WORD",
		"PLAIN_TEXT",
		"PPT",
	}
}

type DatabaseEngineType string

// Enum values for DatabaseEngineType
const (
	DatabaseEngineTypeRdsAuroraMysql      DatabaseEngineType = "RDS_AURORA_MYSQL"
	DatabaseEngineTypeRdsAuroraPostgresql DatabaseEngineType = "RDS_AURORA_POSTGRESQL"
	DatabaseEngineTypeRdsMysql            DatabaseEngineType = "RDS_MYSQL"
	DatabaseEngineTypeRdsPostgresql       DatabaseEngineType = "RDS_POSTGRESQL"
)

// Values returns all known values for DatabaseEngineType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DatabaseEngineType) Values() []DatabaseEngineType {
	return []DatabaseEngineType{
		"RDS_AURORA_MYSQL",
		"RDS_AURORA_POSTGRESQL",
		"RDS_MYSQL",
		"RDS_POSTGRESQL",
	}
}

type DataSourceStatus string

// Enum values for DataSourceStatus
const (
	DataSourceStatusCreating DataSourceStatus = "CREATING"
	DataSourceStatusDeleting DataSourceStatus = "DELETING"
	DataSourceStatusFailed   DataSourceStatus = "FAILED"
	DataSourceStatusUpdating DataSourceStatus = "UPDATING"
	DataSourceStatusActive   DataSourceStatus = "ACTIVE"
)

// Values returns all known values for DataSourceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataSourceStatus) Values() []DataSourceStatus {
	return []DataSourceStatus{
		"CREATING",
		"DELETING",
		"FAILED",
		"UPDATING",
		"ACTIVE",
	}
}

type DataSourceSyncJobStatus string

// Enum values for DataSourceSyncJobStatus
const (
	DataSourceSyncJobStatusFailed          DataSourceSyncJobStatus = "FAILED"
	DataSourceSyncJobStatusSucceeded       DataSourceSyncJobStatus = "SUCCEEDED"
	DataSourceSyncJobStatusSyncing         DataSourceSyncJobStatus = "SYNCING"
	DataSourceSyncJobStatusIncomplete      DataSourceSyncJobStatus = "INCOMPLETE"
	DataSourceSyncJobStatusStopping        DataSourceSyncJobStatus = "STOPPING"
	DataSourceSyncJobStatusAborted         DataSourceSyncJobStatus = "ABORTED"
	DataSourceSyncJobStatusSyncingIndexing DataSourceSyncJobStatus = "SYNCING_INDEXING"
)

// Values returns all known values for DataSourceSyncJobStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataSourceSyncJobStatus) Values() []DataSourceSyncJobStatus {
	return []DataSourceSyncJobStatus{
		"FAILED",
		"SUCCEEDED",
		"SYNCING",
		"INCOMPLETE",
		"STOPPING",
		"ABORTED",
		"SYNCING_INDEXING",
	}
}

type DataSourceType string

// Enum values for DataSourceType
const (
	DataSourceTypeS3          DataSourceType = "S3"
	DataSourceTypeSharepoint  DataSourceType = "SHAREPOINT"
	DataSourceTypeDatabase    DataSourceType = "DATABASE"
	DataSourceTypeSalesforce  DataSourceType = "SALESFORCE"
	DataSourceTypeOnedrive    DataSourceType = "ONEDRIVE"
	DataSourceTypeServicenow  DataSourceType = "SERVICENOW"
	DataSourceTypeCustom      DataSourceType = "CUSTOM"
	DataSourceTypeConfluence  DataSourceType = "CONFLUENCE"
	DataSourceTypeGoogledrive DataSourceType = "GOOGLEDRIVE"
)

// Values returns all known values for DataSourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataSourceType) Values() []DataSourceType {
	return []DataSourceType{
		"S3",
		"SHAREPOINT",
		"DATABASE",
		"SALESFORCE",
		"ONEDRIVE",
		"SERVICENOW",
		"CUSTOM",
		"CONFLUENCE",
		"GOOGLEDRIVE",
	}
}

type DocumentAttributeValueType string

// Enum values for DocumentAttributeValueType
const (
	DocumentAttributeValueTypeStringValue     DocumentAttributeValueType = "STRING_VALUE"
	DocumentAttributeValueTypeStringListValue DocumentAttributeValueType = "STRING_LIST_VALUE"
	DocumentAttributeValueTypeLongValue       DocumentAttributeValueType = "LONG_VALUE"
	DocumentAttributeValueTypeDateValue       DocumentAttributeValueType = "DATE_VALUE"
)

// Values returns all known values for DocumentAttributeValueType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DocumentAttributeValueType) Values() []DocumentAttributeValueType {
	return []DocumentAttributeValueType{
		"STRING_VALUE",
		"STRING_LIST_VALUE",
		"LONG_VALUE",
		"DATE_VALUE",
	}
}

type ErrorCode string

// Enum values for ErrorCode
const (
	ErrorCodeInternalError  ErrorCode = "InternalError"
	ErrorCodeInvalidRequest ErrorCode = "InvalidRequest"
)

// Values returns all known values for ErrorCode. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (ErrorCode) Values() []ErrorCode {
	return []ErrorCode{
		"InternalError",
		"InvalidRequest",
	}
}

type FaqFileFormat string

// Enum values for FaqFileFormat
const (
	FaqFileFormatCsv           FaqFileFormat = "CSV"
	FaqFileFormatCsvWithHeader FaqFileFormat = "CSV_WITH_HEADER"
	FaqFileFormatJson          FaqFileFormat = "JSON"
)

// Values returns all known values for FaqFileFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FaqFileFormat) Values() []FaqFileFormat {
	return []FaqFileFormat{
		"CSV",
		"CSV_WITH_HEADER",
		"JSON",
	}
}

type FaqStatus string

// Enum values for FaqStatus
const (
	FaqStatusCreating FaqStatus = "CREATING"
	FaqStatusUpdating FaqStatus = "UPDATING"
	FaqStatusActive   FaqStatus = "ACTIVE"
	FaqStatusDeleting FaqStatus = "DELETING"
	FaqStatusFailed   FaqStatus = "FAILED"
)

// Values returns all known values for FaqStatus. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (FaqStatus) Values() []FaqStatus {
	return []FaqStatus{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"FAILED",
	}
}

type HighlightType string

// Enum values for HighlightType
const (
	HighlightTypeStandard         HighlightType = "STANDARD"
	HighlightTypeThesaurusSynonym HighlightType = "THESAURUS_SYNONYM"
)

// Values returns all known values for HighlightType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (HighlightType) Values() []HighlightType {
	return []HighlightType{
		"STANDARD",
		"THESAURUS_SYNONYM",
	}
}

type IndexEdition string

// Enum values for IndexEdition
const (
	IndexEditionDeveloperEdition  IndexEdition = "DEVELOPER_EDITION"
	IndexEditionEnterpriseEdition IndexEdition = "ENTERPRISE_EDITION"
)

// Values returns all known values for IndexEdition. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (IndexEdition) Values() []IndexEdition {
	return []IndexEdition{
		"DEVELOPER_EDITION",
		"ENTERPRISE_EDITION",
	}
}

type IndexStatus string

// Enum values for IndexStatus
const (
	IndexStatusCreating       IndexStatus = "CREATING"
	IndexStatusActive         IndexStatus = "ACTIVE"
	IndexStatusDeleting       IndexStatus = "DELETING"
	IndexStatusFailed         IndexStatus = "FAILED"
	IndexStatusUpdating       IndexStatus = "UPDATING"
	IndexStatusSystemUpdating IndexStatus = "SYSTEM_UPDATING"
)

// Values returns all known values for IndexStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (IndexStatus) Values() []IndexStatus {
	return []IndexStatus{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"FAILED",
		"UPDATING",
		"SYSTEM_UPDATING",
	}
}

type KeyLocation string

// Enum values for KeyLocation
const (
	KeyLocationUrl           KeyLocation = "URL"
	KeyLocationSecretManager KeyLocation = "SECRET_MANAGER"
)

// Values returns all known values for KeyLocation. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (KeyLocation) Values() []KeyLocation {
	return []KeyLocation{
		"URL",
		"SECRET_MANAGER",
	}
}

type Mode string

// Enum values for Mode
const (
	ModeEnabled   Mode = "ENABLED"
	ModeLearnOnly Mode = "LEARN_ONLY"
)

// Values returns all known values for Mode. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Mode) Values() []Mode {
	return []Mode{
		"ENABLED",
		"LEARN_ONLY",
	}
}

type Order string

// Enum values for Order
const (
	OrderAscending  Order = "ASCENDING"
	OrderDescending Order = "DESCENDING"
)

// Values returns all known values for Order. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Order) Values() []Order {
	return []Order{
		"ASCENDING",
		"DESCENDING",
	}
}

type PrincipalType string

// Enum values for PrincipalType
const (
	PrincipalTypeUser  PrincipalType = "USER"
	PrincipalTypeGroup PrincipalType = "GROUP"
)

// Values returns all known values for PrincipalType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PrincipalType) Values() []PrincipalType {
	return []PrincipalType{
		"USER",
		"GROUP",
	}
}

type QueryIdentifiersEnclosingOption string

// Enum values for QueryIdentifiersEnclosingOption
const (
	QueryIdentifiersEnclosingOptionDoubleQuotes QueryIdentifiersEnclosingOption = "DOUBLE_QUOTES"
	QueryIdentifiersEnclosingOptionNone         QueryIdentifiersEnclosingOption = "NONE"
)

// Values returns all known values for QueryIdentifiersEnclosingOption. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (QueryIdentifiersEnclosingOption) Values() []QueryIdentifiersEnclosingOption {
	return []QueryIdentifiersEnclosingOption{
		"DOUBLE_QUOTES",
		"NONE",
	}
}

type QueryResultType string

// Enum values for QueryResultType
const (
	QueryResultTypeDocument       QueryResultType = "DOCUMENT"
	QueryResultTypeQuestionAnswer QueryResultType = "QUESTION_ANSWER"
	QueryResultTypeAnswer         QueryResultType = "ANSWER"
)

// Values returns all known values for QueryResultType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (QueryResultType) Values() []QueryResultType {
	return []QueryResultType{
		"DOCUMENT",
		"QUESTION_ANSWER",
		"ANSWER",
	}
}

type QuerySuggestionsBlockListStatus string

// Enum values for QuerySuggestionsBlockListStatus
const (
	QuerySuggestionsBlockListStatusActive                QuerySuggestionsBlockListStatus = "ACTIVE"
	QuerySuggestionsBlockListStatusCreating              QuerySuggestionsBlockListStatus = "CREATING"
	QuerySuggestionsBlockListStatusDeleting              QuerySuggestionsBlockListStatus = "DELETING"
	QuerySuggestionsBlockListStatusUpdating              QuerySuggestionsBlockListStatus = "UPDATING"
	QuerySuggestionsBlockListStatusActiveButUpdateFailed QuerySuggestionsBlockListStatus = "ACTIVE_BUT_UPDATE_FAILED"
	QuerySuggestionsBlockListStatusFailed                QuerySuggestionsBlockListStatus = "FAILED"
)

// Values returns all known values for QuerySuggestionsBlockListStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (QuerySuggestionsBlockListStatus) Values() []QuerySuggestionsBlockListStatus {
	return []QuerySuggestionsBlockListStatus{
		"ACTIVE",
		"CREATING",
		"DELETING",
		"UPDATING",
		"ACTIVE_BUT_UPDATE_FAILED",
		"FAILED",
	}
}

type QuerySuggestionsStatus string

// Enum values for QuerySuggestionsStatus
const (
	QuerySuggestionsStatusActive   QuerySuggestionsStatus = "ACTIVE"
	QuerySuggestionsStatusUpdating QuerySuggestionsStatus = "UPDATING"
)

// Values returns all known values for QuerySuggestionsStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (QuerySuggestionsStatus) Values() []QuerySuggestionsStatus {
	return []QuerySuggestionsStatus{
		"ACTIVE",
		"UPDATING",
	}
}

type ReadAccessType string

// Enum values for ReadAccessType
const (
	ReadAccessTypeAllow ReadAccessType = "ALLOW"
	ReadAccessTypeDeny  ReadAccessType = "DENY"
)

// Values returns all known values for ReadAccessType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReadAccessType) Values() []ReadAccessType {
	return []ReadAccessType{
		"ALLOW",
		"DENY",
	}
}

type RelevanceType string

// Enum values for RelevanceType
const (
	RelevanceTypeRelevant    RelevanceType = "RELEVANT"
	RelevanceTypeNotRelevant RelevanceType = "NOT_RELEVANT"
)

// Values returns all known values for RelevanceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RelevanceType) Values() []RelevanceType {
	return []RelevanceType{
		"RELEVANT",
		"NOT_RELEVANT",
	}
}

type SalesforceChatterFeedIncludeFilterType string

// Enum values for SalesforceChatterFeedIncludeFilterType
const (
	SalesforceChatterFeedIncludeFilterTypeActiveUser   SalesforceChatterFeedIncludeFilterType = "ACTIVE_USER"
	SalesforceChatterFeedIncludeFilterTypeStandardUser SalesforceChatterFeedIncludeFilterType = "STANDARD_USER"
)

// Values returns all known values for SalesforceChatterFeedIncludeFilterType. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (SalesforceChatterFeedIncludeFilterType) Values() []SalesforceChatterFeedIncludeFilterType {
	return []SalesforceChatterFeedIncludeFilterType{
		"ACTIVE_USER",
		"STANDARD_USER",
	}
}

type SalesforceKnowledgeArticleState string

// Enum values for SalesforceKnowledgeArticleState
const (
	SalesforceKnowledgeArticleStateDraft     SalesforceKnowledgeArticleState = "DRAFT"
	SalesforceKnowledgeArticleStatePublished SalesforceKnowledgeArticleState = "PUBLISHED"
	SalesforceKnowledgeArticleStateArchived  SalesforceKnowledgeArticleState = "ARCHIVED"
)

// Values returns all known values for SalesforceKnowledgeArticleState. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (SalesforceKnowledgeArticleState) Values() []SalesforceKnowledgeArticleState {
	return []SalesforceKnowledgeArticleState{
		"DRAFT",
		"PUBLISHED",
		"ARCHIVED",
	}
}

type SalesforceStandardObjectName string

// Enum values for SalesforceStandardObjectName
const (
	SalesforceStandardObjectNameAccount     SalesforceStandardObjectName = "ACCOUNT"
	SalesforceStandardObjectNameCampaign    SalesforceStandardObjectName = "CAMPAIGN"
	SalesforceStandardObjectNameCase        SalesforceStandardObjectName = "CASE"
	SalesforceStandardObjectNameContact     SalesforceStandardObjectName = "CONTACT"
	SalesforceStandardObjectNameContract    SalesforceStandardObjectName = "CONTRACT"
	SalesforceStandardObjectNameDocument    SalesforceStandardObjectName = "DOCUMENT"
	SalesforceStandardObjectNameGroup       SalesforceStandardObjectName = "GROUP"
	SalesforceStandardObjectNameIdea        SalesforceStandardObjectName = "IDEA"
	SalesforceStandardObjectNameLead        SalesforceStandardObjectName = "LEAD"
	SalesforceStandardObjectNameOpportunity SalesforceStandardObjectName = "OPPORTUNITY"
	SalesforceStandardObjectNamePartner     SalesforceStandardObjectName = "PARTNER"
	SalesforceStandardObjectNamePricebook   SalesforceStandardObjectName = "PRICEBOOK"
	SalesforceStandardObjectNameProduct     SalesforceStandardObjectName = "PRODUCT"
	SalesforceStandardObjectNameProfile     SalesforceStandardObjectName = "PROFILE"
	SalesforceStandardObjectNameSolution    SalesforceStandardObjectName = "SOLUTION"
	SalesforceStandardObjectNameTask        SalesforceStandardObjectName = "TASK"
	SalesforceStandardObjectNameUser        SalesforceStandardObjectName = "USER"
)

// Values returns all known values for SalesforceStandardObjectName. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (SalesforceStandardObjectName) Values() []SalesforceStandardObjectName {
	return []SalesforceStandardObjectName{
		"ACCOUNT",
		"CAMPAIGN",
		"CASE",
		"CONTACT",
		"CONTRACT",
		"DOCUMENT",
		"GROUP",
		"IDEA",
		"LEAD",
		"OPPORTUNITY",
		"PARTNER",
		"PRICEBOOK",
		"PRODUCT",
		"PROFILE",
		"SOLUTION",
		"TASK",
		"USER",
	}
}

type ScoreConfidence string

// Enum values for ScoreConfidence
const (
	ScoreConfidenceVeryHigh ScoreConfidence = "VERY_HIGH"
	ScoreConfidenceHigh     ScoreConfidence = "HIGH"
	ScoreConfidenceMedium   ScoreConfidence = "MEDIUM"
	ScoreConfidenceLow      ScoreConfidence = "LOW"
)

// Values returns all known values for ScoreConfidence. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ScoreConfidence) Values() []ScoreConfidence {
	return []ScoreConfidence{
		"VERY_HIGH",
		"HIGH",
		"MEDIUM",
		"LOW",
	}
}

type ServiceNowAuthenticationType string

// Enum values for ServiceNowAuthenticationType
const (
	ServiceNowAuthenticationTypeHttpBasic ServiceNowAuthenticationType = "HTTP_BASIC"
	ServiceNowAuthenticationTypeOauth2    ServiceNowAuthenticationType = "OAUTH2"
)

// Values returns all known values for ServiceNowAuthenticationType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ServiceNowAuthenticationType) Values() []ServiceNowAuthenticationType {
	return []ServiceNowAuthenticationType{
		"HTTP_BASIC",
		"OAUTH2",
	}
}

type ServiceNowBuildVersionType string

// Enum values for ServiceNowBuildVersionType
const (
	ServiceNowBuildVersionTypeLondon ServiceNowBuildVersionType = "LONDON"
	ServiceNowBuildVersionTypeOthers ServiceNowBuildVersionType = "OTHERS"
)

// Values returns all known values for ServiceNowBuildVersionType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ServiceNowBuildVersionType) Values() []ServiceNowBuildVersionType {
	return []ServiceNowBuildVersionType{
		"LONDON",
		"OTHERS",
	}
}

type SharePointVersion string

// Enum values for SharePointVersion
const (
	SharePointVersionSharepointOnline SharePointVersion = "SHAREPOINT_ONLINE"
)

// Values returns all known values for SharePointVersion. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SharePointVersion) Values() []SharePointVersion {
	return []SharePointVersion{
		"SHAREPOINT_ONLINE",
	}
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderDesc SortOrder = "DESC"
	SortOrderAsc  SortOrder = "ASC"
)

// Values returns all known values for SortOrder. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (SortOrder) Values() []SortOrder {
	return []SortOrder{
		"DESC",
		"ASC",
	}
}

type ThesaurusStatus string

// Enum values for ThesaurusStatus
const (
	ThesaurusStatusCreating              ThesaurusStatus = "CREATING"
	ThesaurusStatusActive                ThesaurusStatus = "ACTIVE"
	ThesaurusStatusDeleting              ThesaurusStatus = "DELETING"
	ThesaurusStatusUpdating              ThesaurusStatus = "UPDATING"
	ThesaurusStatusActiveButUpdateFailed ThesaurusStatus = "ACTIVE_BUT_UPDATE_FAILED"
	ThesaurusStatusFailed                ThesaurusStatus = "FAILED"
)

// Values returns all known values for ThesaurusStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ThesaurusStatus) Values() []ThesaurusStatus {
	return []ThesaurusStatus{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"UPDATING",
		"ACTIVE_BUT_UPDATE_FAILED",
		"FAILED",
	}
}

type UserContextPolicy string

// Enum values for UserContextPolicy
const (
	UserContextPolicyAttributeFilter UserContextPolicy = "ATTRIBUTE_FILTER"
	UserContextPolicyUserToken       UserContextPolicy = "USER_TOKEN"
)

// Values returns all known values for UserContextPolicy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UserContextPolicy) Values() []UserContextPolicy {
	return []UserContextPolicy{
		"ATTRIBUTE_FILTER",
		"USER_TOKEN",
	}
}
