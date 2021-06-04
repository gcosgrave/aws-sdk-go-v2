// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// ABP device object for LoRaWAN specification v1.0.x
type AbpV1_0_x struct {

	// The DevAddr value.
	DevAddr *string

	// Session keys for ABP v1.0.x
	SessionKeys *SessionKeysAbpV1_0_x
}

// ABP device object for LoRaWAN specification v1.1
type AbpV1_1 struct {

	// The DevAddr value.
	DevAddr *string

	// Session keys for ABP v1.1
	SessionKeys *SessionKeysAbpV1_1
}

// List of sidewalk certificates.
type CertificateList struct {

	// The certificate chain algorithm provided by sidewalk.
	//
	// This member is required.
	SigningAlg SigningAlg

	// The value of the chosen sidewalk certificate.
	//
	// This member is required.
	Value *string
}

// Describes a destination.
type Destinations struct {

	// The Amazon Resource Name of the resource.
	Arn *string

	// The description of the resource.
	Description *string

	// The rule name or topic rule to send messages to.
	Expression *string

	// The type of value in Expression.
	ExpressionType ExpressionType

	// The name of the resource.
	Name *string

	// The ARN of the IAM Role that authorizes the destination.
	RoleArn *string
}

// Describes a device profile.
type DeviceProfile struct {

	// The Amazon Resource Name of the resource.
	Arn *string

	// The ID of the device profile.
	Id *string

	// The name of the resource.
	Name *string
}

// LoRaWAN object for create functions.
type LoRaWANDevice struct {

	// LoRaWAN object for create APIs
	AbpV1_0_x *AbpV1_0_x

	// ABP device object for create APIs for v1.1
	AbpV1_1 *AbpV1_1

	// The DevEUI value.
	DevEui *string

	// The ID of the device profile for the new wireless device.
	DeviceProfileId *string

	// OTAA device object for create APIs for v1.0.x
	OtaaV1_0_x *OtaaV1_0_x

	// OTAA device object for v1.1 for create APIs
	OtaaV1_1 *OtaaV1_1

	// The ID of the service profile.
	ServiceProfileId *string
}

// LoRaWAN device metatdata.
type LoRaWANDeviceMetadata struct {

	// The DataRate value.
	DataRate *int32

	// The DevEUI value.
	DevEui *string

	// The FPort value.
	FPort *int32

	// The device's channel frequency in Hz.
	Frequency *int32

	// Information about the gateways accessed by the device.
	Gateways []LoRaWANGatewayMetadata

	// The date and time of the metadata.
	Timestamp *string
}

// LoRaWANDeviceProfile object.
type LoRaWANDeviceProfile struct {

	// The ClassBTimeout value.
	ClassBTimeout *int32

	// The ClassCTimeout value.
	ClassCTimeout *int32

	// The list of values that make up the FactoryPresetFreqs value.
	FactoryPresetFreqsList []int32

	// The MAC version (such as OTAA 1.1 or OTAA 1.0.3) to use with this device
	// profile.
	MacVersion *string

	// The MaxDutyCycle value.
	MaxDutyCycle *int32

	// The MaxEIRP value.
	MaxEirp *int32

	// The PingSlotDR value.
	PingSlotDr *int32

	// The PingSlotFreq value.
	PingSlotFreq *int32

	// The PingSlotPeriod value.
	PingSlotPeriod *int32

	// The version of regional parameters.
	RegParamsRevision *string

	// The frequency band (RFRegion) value.
	RfRegion *string

	// The RXDataRate2 value.
	RxDataRate2 *int32

	// The RXDelay1 value.
	RxDelay1 *int32

	// The RXDROffset1 value.
	RxDrOffset1 *int32

	// The RXFreq2 value.
	RxFreq2 *int32

	// The Supports32BitFCnt value.
	Supports32BitFCnt bool

	// The SupportsClassB value.
	SupportsClassB bool

	// The SupportsClassC value.
	SupportsClassC bool

	// The SupportsJoin value.
	SupportsJoin *bool
}

// LoRaWANGateway object.
type LoRaWANGateway struct {

	// The gateway's EUI value.
	GatewayEui *string

	// A list of JoinEuiRange used by LoRa gateways to filter LoRa frames.
	JoinEuiFilters [][]string

	// A list of NetId values that are used by LoRa gateways to filter the uplink
	// frames.
	NetIdFilters []string

	// The frequency band (RFRegion) value.
	RfRegion *string

	// A list of integer indicating which sub bands are supported by LoRa gateway.
	SubBands []int32
}

// LoRaWANGatewayCurrentVersion object.
type LoRaWANGatewayCurrentVersion struct {

	// The version of the gateways that should receive the update.
	CurrentVersion *LoRaWANGatewayVersion
}

// LoRaWAN gateway metatdata.
type LoRaWANGatewayMetadata struct {

	// The gateway's EUI value.
	GatewayEui *string

	// The RSSI value.
	Rssi *float64

	// The SNR value.
	Snr *float64
}

// LoRaWANGatewayVersion object.
type LoRaWANGatewayVersion struct {

	// The model number of the wireless gateway.
	Model *string

	// The version of the wireless gateway firmware.
	PackageVersion *string

	// The basic station version of the wireless gateway.
	Station *string
}

// LoRaWANGetServiceProfileInfo object.
type LoRaWANGetServiceProfileInfo struct {

	// The AddGWMetaData value.
	AddGwMetadata bool

	// The ChannelMask value.
	ChannelMask *string

	// The DevStatusReqFreq value.
	DevStatusReqFreq *int32

	// The DLBucketSize value.
	DlBucketSize *int32

	// The DLRate value.
	DlRate *int32

	// The DLRatePolicy value.
	DlRatePolicy *string

	// The DRMax value.
	DrMax int32

	// The DRMin value.
	DrMin int32

	// The HRAllowed value that describes whether handover roaming is allowed.
	HrAllowed bool

	// The MinGwDiversity value.
	MinGwDiversity *int32

	// The NwkGeoLoc value.
	NwkGeoLoc bool

	// The PRAllowed value that describes whether passive roaming is allowed.
	PrAllowed bool

	// The RAAllowed value that describes whether roaming activation is allowed.
	RaAllowed bool

	// The ReportDevStatusBattery value.
	ReportDevStatusBattery bool

	// The ReportDevStatusMargin value.
	ReportDevStatusMargin bool

	// The TargetPER value.
	TargetPer int32

	// The ULBucketSize value.
	UlBucketSize *int32

	// The ULRate value.
	UlRate *int32

	// The ULRatePolicy value.
	UlRatePolicy *string
}

// LoRaWAN object for list functions.
type LoRaWANListDevice struct {

	// The DevEUI value.
	DevEui *string
}

// LoRaWAN router info.
type LoRaWANSendDataToDevice struct {

	// The Fport value.
	FPort *int32
}

// LoRaWANServiceProfile object.
type LoRaWANServiceProfile struct {

	// The AddGWMetaData value.
	AddGwMetadata bool
}

// LoRaWAN object for update functions.
type LoRaWANUpdateDevice struct {

	// The ID of the device profile for the wireless device.
	DeviceProfileId *string

	// The ID of the service profile.
	ServiceProfileId *string
}

// LoRaWANUpdateGatewayTaskCreate object.
type LoRaWANUpdateGatewayTaskCreate struct {

	// The version of the gateways that should receive the update.
	CurrentVersion *LoRaWANGatewayVersion

	// The CRC of the signature private key to check.
	SigKeyCrc *int64

	// The signature used to verify the update firmware.
	UpdateSignature *string

	// The firmware version to update the gateway to.
	UpdateVersion *LoRaWANGatewayVersion
}

// LoRaWANUpdateGatewayTaskEntry object.
type LoRaWANUpdateGatewayTaskEntry struct {

	// The version of the gateways that should receive the update.
	CurrentVersion *LoRaWANGatewayVersion

	// The firmware version to update the gateway to.
	UpdateVersion *LoRaWANGatewayVersion
}

// OTAA device object for v1.0.x
type OtaaV1_0_x struct {

	// The AppEUI value.
	AppEui *string

	// The AppKey value.
	AppKey *string
}

// OTAA device object for v1.1
type OtaaV1_1 struct {

	// The AppKey value.
	AppKey *string

	// The JoinEUI value.
	JoinEui *string

	// The NwkKey value.
	NwkKey *string
}

// Information about a service profile.
type ServiceProfile struct {

	// The Amazon Resource Name of the resource.
	Arn *string

	// The ID of the service profile.
	Id *string

	// The name of the resource.
	Name *string
}

// Session keys for ABP v1.1
type SessionKeysAbpV1_0_x struct {

	// The AppSKey value.
	AppSKey *string

	// The NwkSKey value.
	NwkSKey *string
}

// Session keys for ABP v1.1
type SessionKeysAbpV1_1 struct {

	// The AppSKey value.
	AppSKey *string

	// The FNwkSIntKey value.
	FNwkSIntKey *string

	// The NwkSEncKey value.
	NwkSEncKey *string

	// The SNwkSIntKey value.
	SNwkSIntKey *string
}

// Information about a Sidewalk account.
type SidewalkAccountInfo struct {

	// The Sidewalk Amazon ID.
	AmazonId *string

	// The Sidewalk application server private key.
	AppServerPrivateKey *string
}

// Information about a Sidewalk account.
type SidewalkAccountInfoWithFingerprint struct {

	// The Sidewalk Amazon ID.
	AmazonId *string

	// The Amazon Resource Name of the resource.
	Arn *string

	// The fingerprint of the Sidewalk application server private key.
	Fingerprint *string
}

// Sidewalk device object.
type SidewalkDevice struct {

	// The sidewalk device certificates for Ed25519 and P256r1.
	DeviceCertificates []CertificateList

	// The sidewalk device identification.
	SidewalkId *string

	// The Sidewalk manufacturing series number.
	SidewalkManufacturingSn *string
}

// MetaData for Sidewalk device.
type SidewalkDeviceMetadata struct {

	// Sidewalk device battery level.
	BatteryLevel BatteryLevel

	// Device state defines the device status of sidewalk device.
	DeviceState DeviceState

	// Sidewalk device status notification.
	Event Event

	// The RSSI value.
	Rssi *int32
}

// Sidewalk object used by list functions.
type SidewalkListDevice struct {

	// The Sidewalk Amazon ID.
	AmazonId *string

	// The sidewalk device certificates for Ed25519 and P256r1.
	DeviceCertificates []CertificateList

	// The sidewalk device identification.
	SidewalkId *string

	// The Sidewalk manufacturing series number.
	SidewalkManufacturingSn *string
}

// Information about a Sidewalk router.
type SidewalkSendDataToDevice struct {

	// Sidewalk device message type.
	MessageType MessageType

	// The sequence number.
	Seq *int32
}

// Sidewalk update.
type SidewalkUpdateAccount struct {

	// The new Sidewalk application server private key.
	AppServerPrivateKey *string
}

// A simple label consisting of a customer-defined key-value pair
type Tag struct {

	// The tag's key value.
	//
	// This member is required.
	Key *string

	// The tag's value.
	//
	// This member is required.
	Value *string
}

// UpdateWirelessGatewayTaskCreate object.
type UpdateWirelessGatewayTaskCreate struct {

	// The properties that relate to the LoRaWAN wireless gateway.
	LoRaWAN *LoRaWANUpdateGatewayTaskCreate

	// The IAM role used to read data from the S3 bucket.
	UpdateDataRole *string

	// The link to the S3 bucket.
	UpdateDataSource *string
}

// UpdateWirelessGatewayTaskEntry object.
type UpdateWirelessGatewayTaskEntry struct {

	// The Amazon Resource Name of the resource.
	Arn *string

	// The ID of the new wireless gateway task entry.
	Id *string

	// The properties that relate to the LoRaWAN wireless gateway.
	LoRaWAN *LoRaWANUpdateGatewayTaskEntry
}

// The log option for a wireless device event. Can be used to set log level for a
// specific wireless device event. For a LoRaWAN device, the possible events for a
// log messsage are: Join, Rejoin, Downlink_Data, Uplink_Data. For a Sidewalk
// device, the possible events for a log message are: Registration, Downlink_Data,
// Uplink_Data.
type WirelessDeviceEventLogOption struct {

	// The event for a log message, if the log message is tied to a wireless device.
	//
	// This member is required.
	Event WirelessDeviceEvent

	// The log level for a log message.
	//
	// This member is required.
	LogLevel LogLevel
}

// The log option for wireless devices. Can be used to set log level for a specific
// type of wireless device.
type WirelessDeviceLogOption struct {

	// The log level for a log message.
	//
	// This member is required.
	LogLevel LogLevel

	// The wireless device type.
	//
	// This member is required.
	Type WirelessDeviceType

	// The list of wireless device event log options.
	Events []WirelessDeviceEventLogOption
}

// Information about a wireless device's operation.
type WirelessDeviceStatistics struct {

	// The Amazon Resource Name of the resource.
	Arn *string

	// The name of the destination to which the device is assigned.
	DestinationName *string

	// The ID of the wireless device reporting the data.
	Id *string

	// The date and time when the most recent uplink was received.
	LastUplinkReceivedAt *string

	// LoRaWAN device info.
	LoRaWAN *LoRaWANListDevice

	// The name of the resource.
	Name *string

	// The Sidewalk account credentials.
	Sidewalk *SidewalkListDevice

	// The wireless device type.
	Type WirelessDeviceType
}

// The log option for a wireless gateway event. Can be used to set log level for a
// specific wireless gateway event. For a LoRaWAN gateway, the possible events for
// a log message are: CUPS_Request, Certificate.
type WirelessGatewayEventLogOption struct {

	// The event for a log message, if the log message is tied to a wireless gateway.
	//
	// This member is required.
	Event WirelessGatewayEvent

	// The log level for a log message.
	//
	// This member is required.
	LogLevel LogLevel
}

// The log option for wireless gateways. Can be used to set log level for a
// specific type of wireless gateway.
type WirelessGatewayLogOption struct {

	// The log level for a log message.
	//
	// This member is required.
	LogLevel LogLevel

	// The wireless gateway type.
	//
	// This member is required.
	Type WirelessGatewayType

	// The list of wireless gateway event log options.
	Events []WirelessGatewayEventLogOption
}

// Information about a wireless gateway's operation.
type WirelessGatewayStatistics struct {

	// The Amazon Resource Name of the resource.
	Arn *string

	// The description of the resource.
	Description *string

	// The ID of the wireless gateway reporting the data.
	Id *string

	// The date and time when the most recent uplink was received.
	LastUplinkReceivedAt *string

	// LoRaWAN gateway info.
	LoRaWAN *LoRaWANGateway

	// The name of the resource.
	Name *string
}

// WirelessMetadata object.
type WirelessMetadata struct {

	// LoRaWAN device info.
	LoRaWAN *LoRaWANSendDataToDevice

	// The Sidewalk account credentials.
	Sidewalk *SidewalkSendDataToDevice
}
