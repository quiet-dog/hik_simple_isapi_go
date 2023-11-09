package hik

import (
	"encoding/xml"
	"mime/multipart"
	"time"
)

type AddPersonInfoReq struct {
	UserInfo UserInfo `json:"UserInfo"`
}

type UserInfo struct {
	EmployeeNo      string      `json:"employeeNo"`
	Name            string      `json:"name"`
	UserType        string      `json:"userType"`
	Valid           Valid       `json:"Valid"`
	DoorRight       string      `json:"doorRight"`
	RightPlan       []RightPlan `json:"RightPlan"`
	MaxOpenDoorTime int         `json:"maxOpenDoorTime"`
	LocalUIRight    bool        `json:"localUIRight"`
	UserVerifyMode  string      `json:"userVerifyMode"`
	Gender          string      `json:"gender"`
	CallNumbers     []string    `json:"callNumbers"`
	FloorNumbers    []int       `json:"floorNumbers"`
	GroupID         int         `json:"groupId"`
}

type Valid struct {
	Enable    bool   `json:"enable"`
	BeginTime string `json:"beginTime"`
	EndTime   string `json:"endTime"`
	TimeType  string `json:"timeType"`
}

type RightPlan struct {
	DoorNo         int    `json:"doorNo"`
	PlanTemplateNo string `json:"planTemplateNo"`
}

type SetFaceInfo struct {
	FaceDataRecord FaceDataRecordReq     `json:"FaceDataRecord"`
	Img            *multipart.FileHeader `json:"img"`
}

type FaceDataRecordReq struct {
	FaceLibType string `json:"faceLibType"`
	FDID        string `json:"FDID"`
	FPID        string `json:"FPID"`
}

type DelPersonInfo struct {
	UserInfoDelCond UserInfoDelCond
}

type UserInfoDelCond struct {
	EmployeeNoList []EmployeeNo `json:"EmployeeNoList"`
	OperateType    string       `json:"operateType"`
	TerminalNoList []int        `json:"terminalNoList"`
}
type EmployeeNo struct {
	EmployeeNo string `json:"employeeNo"`
}

type GetPersonInfoList struct {
	UserInfoSearchCond UserInfoSearchCond `json:"UserInfoSearchCond"`
}

type UserInfoSearchCond struct {
	SearchID             string `json:"searchID"`
	MaxResults           int    `json:"maxResults"`
	SearchResultPosition int    `json:"searchResultPosition"`
	FuzzySearch          string `json:"fuzzySearch"`
}

type GetPersonInfoListRes struct {
}

type UserInfoSearch struct {
	SearchID           string     `json:"searchID"`
	ResponseStatusStrg string     `json:"ResponseStatusStrg"`
	NumOfMatches       int        `json:"numOfMatches"`
	TotalMatches       int        `json:"totalMatches"`
	UserInfo           []UserInfo `json:"UserInfo"`
}

type LogSearch struct {
	CMSearchDescription CMSearchDescription `json:"CMSearchDescription"`
}

type CMSearchDescription struct {
	SearchID                   string         `json:"searchID"`
	TrackIDList                []Track        `json:"trackIDList"`
	MetaID                     string         `json:"metaId"`
	MetaIDList                 []MetaID       `json:"metaIdList"`
	TimeSpanList               []TimeSpanList `json:"timeSpanList"`
	SearchResultPostion        int            `json:"searchResultPostion"`
	MaxResults                 int            `json:"maxResults"`
	OnlySmart                  bool           `json:"onlySmart"`
	LogLevel                   string         `json:"logLevel"`
	AlarmLevel                 string         `json:"alarmLevel"`
	LogCmd                     string         `json:"logCmd"`
	UserType                   string         `json:"userType"`
	PersonnelChannelModuleType string         `json:"personnelChannelModuleType"`
}

type Track struct {
	TrackID int `json:"trackID"`
}

type MetaID struct {
	MetaID string `json:"metaId"`
}
type TimeSpanList struct {
	TimeSpan TimeSpan `json:"timeSpan"`
}

type TimeSpan struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

type LogSearchRes struct {
	XMLNAME            xml.Name             `xml:"CMSearchResult"`
	SearchID           string               `xml:"searchID"`
	ResponseStatus     bool                 `xml:"responseStatus"`
	ResponseStatusStrg string               `xml:"ResponseStatusStrg"`
	TotalMatches       int                  `xml:"totalMatches"`
	NumOfMatches       int                  `xml:"numOfMatches"`
	MatchList          []SearchLogMatchList `xml:"matchList"`
}

type TrackID struct {
	TrackID int `xml:"trackID"`
}
type MetaId struct {
	MetaId string `xml:"metaId"`
}

type SearchLogMatchList struct {
	LogDescriptor LogDescriptor `xml:"logDescriptor"`
}
type LogDescriptor struct {
	MetaId              string  `xml:"metaId"`
	StartDateTime       string  `xml:"startDateTime"`
	LocalID             string  `xml:"localID"`
	ParaType            string  `xml:"paraType"`
	UserName            string  `xml:"userName"`
	InfoContent         string  `xml:"infoContent"`
	LogInfo             LogInfo `xml:"logInfo"`
	IpAddress           string  `xml:"ipAddress"`
	Object              string  `xml:"object"`
	Params              string  `xml:"params"`
	Seq                 string  `xml:"seq"`
	AdditionInformation string  `xml:"additionInformation"`
	PanelUser           string  `xml:"panelUser"`
	DiskNumber          int     `xml:"diskNumber"`
	AlarmInPort         int     `xml:"alarmInPort"`
	AlarmOutPort        int     `xml:"alarmOutPort"`
	RemoteHostIPAddress string  `xml:"remoteHostIPAddress"`
	LogLevel            string  `xml:"logLevel"`
	AlarmLevel          string  `xml:"alarmLevel"`
	ModuleName          string  `xml:"moduleName"`
}

type LogInfo struct {
	OpenDoorRecord OpenDoorRecord `xml:"openDoorRecord"`
	VisAlarmRecord VisAlarmRecord `xml:"visAlarmRecord"`
}

type OpenDoorRecord struct {
	Type string `xml:"type"`
}

type VisAlarmRecord struct {
	Type string `xml:"type"`
}

type GetPersonInfoCountRes struct {
	UserInfoCount UserInfoCount `json:"UserInfoCount"`
}

type UserInfoCount struct {
	UserNumber                int `json:"userNumber"`
	BindFaceUserNumber        int `json:"bindFaceUserNumber"`
	BindFingerprintUserNumber int `json:"bindFingerprintUserNumber"`
	BindCardUserNumber        int `json:"bindCardUserNumber"`
	BindIrisUserNumber        int `json:"bindIrisUserNumber"`
	BindVoiceprintUserNumber  int `json:"bindVoiceprintUserNumber"`
	BindRemoteControlNumber   int `json:"bindRemoteControlNumber"`
	NormalNumber              int `json:"normalNumber"`
	VisitorNumber             int `json:"visitorNumber"`
}

type GetFaceInfoCountRes struct {
	RequestURL       string             `json:"requestURL"`
	StatusCode       int                `json:"statusCode"`
	StatusString     string             `json:"statusString"`
	SubStatusCode    string             `json:"subStatusCode"`
	ErrorCode        int                `json:"errorCode"`
	ErrorMsg         string             `json:"errorMsg"`
	FDRecordDataInfo []FDRecordDataInfo `json:"FDRecordDataInfo"`
}

type FDRecordDataInfo struct {
	FDID             string `json:"FDID"`
	FaceLibType      string `json:"faceLibType"`
	Name             string `json:"name"`
	RecordDataNumber int    `json:"recordDataNumber"`
}

type GetCardInfoCountRes struct {
	CardInfoCount CardInfoCount `json:"CardInfoCount"`
}

type CardInfoCount struct {
	CardNumber int `json:"cardNumber"`
}

type GetAcsEventTotalNumReq struct {
	AcsEventTotalNumCond AcsEventTotalNumCond `json:"AcsEventTotalNumCond"`
}

type AcsEventTotalNumCond struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	// =====
	StartTime        string `json:"startTime"`
	EndTime          string `json:"endTime"`
	CardNo           string `json:"cardNo"`
	Name             string `json:"name"`
	PicEnable        bool   `json:"picEnable"`
	BeginSerialNo    int    `json:"beginSerialNo"`
	EndSerialNo      int    `json:"endSerialNo"`
	EmployeeNoString string `json:"employeeNoString"`
}

type GetAcsEventTotalNumRes struct {
	AcsEventTotalNumCond AcsEventTotalNumCond `json:"AcsEventTotalNumCond"`
}

type GetAcsWorkStatusRes struct {
	AcsWorkStatus AcsWorkStatus `json:"AcsWorkStatus"`
}

type AcsWorkStatus struct {
	DoorLockStatus                     []int                `json:"doorLockStatus"`
	DoorStatus                         []int                `json:"doorStatus"`
	MagneticStatus                     []int                `json:"magneticStatus"`
	CaseStatus                         []int                `json:"caseStatus"`
	BatteryVoltage                     int                  `json:"batteryVoltage"`
	BatteryLowVoltage                  int                  `json:"batteryLowVoltage"`
	PowerSupplyStatus                  string               `json:"powerSupplyStatus"`
	MultiDoorInterlockStatus           string               `json:"multiDoorInterlockStatus"`
	AntiSneakStatus                    string               `json:"antiSneakStatus"`
	HostAntiDismantleStatus            string               `json:"hostAntiDismantleStatus"`
	IndicatorLightStatus               string               `json:"indicatorLightStatus"`
	CardReaderOnlineStatus             []int                `json:"cardReaderOnlineStatus"`
	NetReaderOnlineStatus              []int                `json:"netReaderOnlineStatus"`
	POEPortList                        []POEPort            `json:"POEPortList"`
	CardReaderAntiDismantleStatus      []int                `json:"cardReaderAntiDismantleStatus"`
	CardReaderVerifyMode               []int                `json:"cardReaderVerifyMode"`
	SetupAlarmStatus                   []int                `json:"setupAlarmStatus"`
	AlarmInStatus                      []int                `json:"alarmInStatus"`
	AlarmOutStatus                     []int                `json:"alarmOutStatus"`
	CardNum                            int                  `json:"cardNum"`
	FireAlarmStatus                    string               `json:"fireAlarmStatus"`
	BatteryChargeStatus                string               `json:"batteryChargeStatus"`
	MasterChannelControllerStatus      string               `json:"masterChannelControllerStatus"`
	SlaveChannelControllerStatus       string               `json:"slaveChannelControllerStatus"`
	AntiSneakServerStatus              string               `json:"antiSneakServerStatus"`
	NetStatus                          string               `json:"netStatus"`
	InterfaceStatusList                []InterfaceStatus    `json:"interfaceStatusList"`
	SignalStatus                       string               `json:"signalStatus"`
	SignalStrength                     int                  `json:"signalStrength"`
	SipStatus                          string               `json:"sipStatus"`
	EzvizStatus                        string               `json:"ezvizStatus"`
	VoipStatus                         string               `json:"voipStatus"`
	DhcpStatus                         string               `json:"dhcpStatus"`
	WifiStatus                         string               `json:"wifiStatus"`
	TFCardStatus                       string               `json:"TFCardStatus"`
	AcrossHostInterlockStatus          string               `json:"acrossHostInterlockStatus"`
	FaceReceivingStatus                string               `json:"faceReceivingStatus"`
	QRCodeReaderStatusList             []QRCodeReaderStatus `json:"QRCodeReaderStatusList"`
	GooseneckStatus                    string               `json:"GooseneckStatus"`
	WiFiSignalStrength                 int                  `json:"wiFiSignalStrength"`
	ThermopileStatus                   string               `json:"thermopileStatus"`
	ModelStatus                        string               `json:"modelStatus"`
	ISUPRegisterStatus                 string               `json:"ISUPRegisterStatus"`
	NewSipStatus                       string               `json:"newSipStatus"`
	AnalogAudioAccessStatus            string               `json:"analogAudioAccessStatus"`
	RealTimeBroadcastStatus            string               `json:"realTimeBroadcastStatus"`
	ConsumeMachineExKeypadStatus       string               `json:"consumeMachineExKeypadStatus"`
	DevWorkStatus                      string               `json:"devWorkStatus"`
	AbnormalReason                     []int                `json:"abnormalReason"`
	LocalControllerOnlineStatus        []int                `json:"localControllerOnlineStatus"`
	CollectionCheckCommunicationStatus string               `json:"collectionCheckCommunicationStatus"`
	OutdoorSipRegisterStatus           string               `json:"outdoorSipRegisterStatus"`
	USBCameraStatus                    string               `json:"USBCameraStatus"`
	DoorOnlineStatus                   []int                `json:"doorOnlineStatus"`
}

type POEPort struct {
	Port     int   `json:"port"`
	ReaderID []int `json:"readerID"`
}

type InterfaceStatus struct {
	ID        int `json:"id"`
	NetType   int `json:"netType"`
	NetStatus int `json:"netStatus"`
}

type QRCodeReaderStatus struct {
	ID                 int    `json:"id"`
	QRCodeReaderStatus string `json:"QRCodeReaderStatus"`
}

type AddCardInfoReq struct {
	CardInfo CardInfo `json:"CardInfo"`
}

type CardInfo struct {
	EmployeeNo     string `json:"employeeNo"`
	CardNo         string `json:"cardNo"`
	CardType       string `json:"cardType"`
	LeaderCard     string `json:"leaderCard"`
	EmployeeNoType string `json:"employeeNoType"`
}

type DelCardInfoReq struct {
	CardInfoDelCond CardInfoDelCond `json:"CardInfoDelCond"`
}

type CardInfoDelCond struct {
	// EmployeeNoList []EmployeeNoList `json:"EmployeeNoList"`
	CardNoList []CardNo `json:"CardNoList"`
	// OperateType    string           `json:"operateType"`
	// TerminalNoList []int            `json:"terminalNoList"`
}
type CardNo struct {
	CardNo string `json:"cardNo"`
}

type GetCardInfoReq struct {
	CardInfoSearchCond CardInfoSearchCond `json:"CardInfoSearchCond"`
}

type CardInfoSearchCond struct {
	SearchID             string       `json:"searchID"`
	MaxResults           int          `json:"maxResults"`
	SearchResultPosition int          `json:"searchResultPosition"`
	EmployeeNoList       []EmployeeNo `json:"EmployeeNoList"`
}

type CardInfoSearch struct {
	SearchID           string `json:"searchID"`
	ResponseStatusStrg string `json:"responseStatusStrg"`
	NumOfMatches       int    `json:"numOfMatches"`
	CotalMatches       int    `json:"totalMatches"`
	CardInfo           []CardInfo
}

type GetCardInfoRes struct {
	CardInfoSearch CardInfoSearch `json:"CardInfoSearch"`
}

type GetAcsEventReq struct {
	AcsEventCond AcsEventCond `json:"AcsEventCond"`
}

type AcsEventCond struct {
	SearchID             string `json:"searchID"`
	SearchResultPosition int    `json:"searchResultPosition"`
	MaxResults           int    `json:"maxResults"`
	Major                int    `json:"major"`
	Minor                int    `json:"minor"`
	StartTime            string `json:"startTime"`
	EndTime              string `json:"endTime"`
	CardNo               string `json:"cardNo"`
	Name                 string `json:"name"`
	EmployeeNoString     string `json:"employeeNoString"`
	BeginSerialNo        int    `json:"beginSerialNo"`
	EndSerialNo          int    `json:"endSerialNo"`
}

type AcsEvent struct {
	SearchID           string      `json:"searchID"`
	ResponseStatusStrg string      `json:"responseStatusStrg"`
	NumOfMatches       int         `json:"numOfMatches"`
	TotalMatches       int         `json:"totalMatches"`
	InfoList           []EventInfo `json:"InfoList"`
}

type EventInfo struct {
	Major          int    `json:"major"`
	Minor          int    `json:"minor"`
	Time           string `json:"time"`
	NetUser        string `json:"netUser"`
	RemoteHostAddr string `json:"remoteHostAddr"`
	VideoChannel   int    `json:"videoChannel"`
	CardNo         string `json:"cardNo"`
	CardType       int    `json:"cardType"`
	WhiteListNo    int    `json:"whiteListNo"`
}

type GetAcsEventRes struct {
	AcsEvent AcsEvent `json:"AcsEvent"`
}

type SetAcsCfgReq struct {
	AcsCfg AcsCfg `json:"AcsCfg"`
}

type AcsCfg struct {
	UploadCapPic                        bool   `json:"uploadCapPic"`
	SaveCapPic                          bool   `json:"saveCapPic"`
	Protocol                            string `json:"protocol"`
	VoicePrompt                         bool   `json:"voicePrompt"`
	ShowPicture                         bool   `json:"showPicture"`
	ShowEmployeeNo                      bool   `json:"showEmployeeNo"`
	ShowName                            bool   `json:"showName"`
	DesensitiseEmployeeNo               bool   `json:"desensitiseEmployeeNo"`
	DesensitiseName                     bool   `json:"desensitiseName"`
	UploadVerificationPic               bool   `json:"uploadVerificationPic"`
	SaveVerificationPic                 bool   `json:"saveVerificationPic"`
	SaveFacePic                         bool   `json:"saveFacePic"`
	RemoteCheckDoorEnabled              bool   `json:"remoteCheckDoorEnabled"`
	CheckChannelType                    string `json:"checkChannelType"`
	ExternalCardReaderEnabled           bool   `json:"externalCardReaderEnabled"`
	CombinationAuthenticationTimeout    int    `json:"combinationAuthenticationTimeout"`
	CombinationAuthenticationLimitOrder bool   `json:"combinationAuthenticationLimitOrder"`
}

type SetStorageCfgReq struct {
	EventStorageCfg EventStorageCfg `json:"EventStorageCfg"`
}
type EventStorageCfg struct {
	Mode      string `json:"mode"`
	CheckTime string `json:"checkTime"`
	Period    int    `json:"period"`
}

type GetUserAndRightShowReq struct {
	ShowCardNo             string `json:"showCardNo"`
	ShowAuthenticationList bool   `json:"showAuthenticationList"`
	ShowDuration           int    `json:"showDuration"`
}

type GetPictureReq struct {
	URL string `json:"url"`
}

type DoorID struct {
	DoorID string `json:"doorID"`
}

type RemoteControlDoorReq struct {
	DoorID
	RemoteControlDoor RemoteControlDoor `json:"RemoteControlDoor"`
}

type RemoteControlDoor struct {
	Cmd                           string                      `json:"cmd" xml:"cmd"`
	Password                      string                      `json:"password" xml:"password"`
	EmployeeNo                    string                      `json:"employeeNo" xml:"employeeNo"`
	ChannelNo                     int                         `json:"channelNo" xml:"channelNo"`
	ControlType                   string                      `json:"controlType" xml:"controlType"`
	PersonnelChannelGroupInfoList []PersonnelChannelGroupInfo `json:"personnelChannelGroupInfoList" xml:"personnelChannelGroupInfoList"`
}

type PersonnelChannelGroupInfo struct {
	PersonnelChannelGroupID  int                    `json:"personnelChannelGroupID" xml:"personnelChannelGroupID"`
	PersonnelChannelInfoList []PersonnelChannelInfo `json:"personnelChannelInfoList" xml:"personnelChannelInfoList"`
}

type PersonnelChannelInfo struct {
	PersonnelChannelID int `json:"personnelChannelID" xml:"personnelChannelID"`
}

type DeviceInfo struct {
	XMLNAME                        xml.Name                       `xml:"DeviceInfo"`
	DeviceName                     string                         `json:"deviceName" xml:"deviceName"`
	DeviceID                       string                         `json:"deviceID" xml:"deviceID"`
	DeviceDescription              string                         `json:"deviceDescription" xml:"deviceDescription"`
	DeviceLocation                 string                         `json:"deviceLocation" xml:"deviceLocation"`
	DeviceStatus                   string                         `json:"deviceStatus" xml:"deviceStatus"`
	DetailAbnormalStatus           DetailAbnormalStatus           `json:"detailAbnormalStatus" xml:"detailAbnormalStatus"`
	SystemContact                  string                         `json:"systemContact" xml:"systemContact"`
	Model                          string                         `json:"model" xml:"model"`
	SerialNumber                   string                         `json:"serialNumber" xml:"serialNumber"`
	MacAddress                     string                         `json:"macAddress" xml:"macAddress"`
	FirmwareVersion                string                         `json:"firmwareVersion" xml:"firmwareVersion"`
	FirmwareReleasedDate           string                         `json:"firmwareReleasedDate" xml:"firmwareReleasedDate"`
	EncoderVersion                 string                         `json:"encoderVersion" xml:"encoderVersion"`
	EncoderReleasedDate            string                         `json:"encoderReleasedDate" xml:"encoderReleasedDate"`
	BootVersion                    string                         `json:"bootVersion" xml:"bootVersion"`
	BootReleasedDate               string                         `json:"bootReleasedDate" xml:"bootReleasedDate"`
	PanelVersion                   string                         `json:"panelVersion" xml:"panelVersion"`
	HardwareVersion                string                         `json:"hardwareVersion" xml:"hardwareVersion"`
	DecoderVersion                 string                         `json:"decoderVersion" xml:"decoderVersion"`
	DecoderReleasedDate            string                         `json:"decoderReleasedDate" xml:"decoderReleasedDate"`
	SoftwareVersion                string                         `json:"softwareVersion" xml:"softwareVersion"`
	Capacity                       int                            `json:"capacity" xml:"capacity"`
	UsedCapacity                   int                            `json:"usedCapacity" xml:"usedCapacity"`
	DeviceType                     string                         `json:"deviceType" xml:"deviceType"`
	SubDeviceType                  string                         `json:"subDeviceType" xml:"subDeviceType"`
	TelecontrolID                  int                            `json:"telecontrolID" xml:"telecontrolID"`
	SupportBeep                    bool                           `json:"supportBeep" xml:"supportBeep"`
	SupportVideoLoss               bool                           `json:"supportVideoLoss" xml:"supportVideoLoss"`
	FirmwareVersionInfo            string                         `json:"firmwareVersionInfo" xml:"firmwareVersionInfo"`
	ActualFloorNum                 int                            `json:"actualFloorNum" xml:"actualFloorNum"`
	LocalZoneNum                   int                            `json:"localZoneNum" xml:"localZoneNum"`
	AlarmOutNum                    int                            `json:"alarmOutNum" xml:"alarmOutNum"`
	AlarmInNum                     int                            `json:"alarmInNum" xml:"alarmInNum"`
	DistanceResolution             float64                        `json:"distanceResolution" xml:"distanceResolution"`
	AngleResolution                float64                        `json:"angleResolution" xml:"angleResolution"`
	SpeedResolution                float64                        `json:"speedResolution" xml:"speedResolution"`
	DetectDistance                 float64                        `json:"detectDistance" xml:"detectDistance"`
	RelayNum                       int                            `json:"relayNum" xml:"relayNum"`
	ElectroLockNum                 int                            `json:"electroLockNum" xml:"electroLockNum"`
	SirenNum                       int                            `json:"sirenNum" xml:"sirenNum"`
	AlarmLamp                      int                            `json:"alarmLamp" xml:"alarmLamp"`
	RS485Num                       int                            `json:"RS485Num" xml:"RS485Num"`
	RadarVersion                   string                         `json:"radarVersion" xml:"radarVersion"`
	CameraModuleVersion            string                         `json:"cameraModuleVersion" xml:"cameraModuleVersion"`
	Mainversion                    string                         `json:"mainversion" xml:"mainversion"`
	Subversion                     string                         `json:"subversion" xml:"subversion"`
	Upgradeversion                 string                         `json:"upgradeversion" xml:"upgradeversion"`
	Customizeversion               string                         `json:"customizeversion" xml:"customizeversion"`
	CompanyName                    string                         `json:"companyName" xml:"companyName"`
	Copyright                      string                         `json:"copyRight" xml:"copyRight"`
	SystemName                     string                         `json:"systemName" xml:"systemName"`
	SystemStatus                   string                         `json:"systemStatus" xml:"systemStatus"`
	IsLeaderDevice                 bool                           `json:"isLeaderDevice" xml:"isLeaderDevice"`
	ClusterVersion                 string                         `json:"clusterVersion" xml:"clusterVersion"`
	CentralStorageVersion          string                         `json:"centralStorageVersion" xml:"centralStorageVersion"`
	PowerOnMode                    string                         `json:"powerOnMode" xml:"powerOnMode"`
	CustomizedInfo                 string                         `json:"customizedInfo" xml:"customizedInfo"`
	VerificationCode               string                         `json:"verificationCode" xml:"verificationCode"`
	SupportUrl                     string                         `json:"supportUrl" xml:"supportUrl"`
	SubSerialNumber                string                         `json:"subSerialNumber" xml:"subSerialNumber"`
	LanguageType                   string                         `json:"languageType" xml:"languageType"`
	DockStation                    Platform                       `json:"dockStation" xml:"dockStation"`
	WebVersion                     string                         `json:"webVersion" xml:"webVersion"`
	DeviceRFProgramVersion         string                         `json:"deviceRFProgramVersion" xml:"deviceRFProgramVersion"`
	SecurityModuleSerialNo         string                         `json:"securityModuleSerialNo" xml:"securityModuleSerialNo"`
	SecurityModuleVersion          string                         `json:"securityModuleVersion" xml:"securityModuleVersion"`
	SecurityChipVersion            string                         `json:"securityChipVersion" xml:"securityChipVersion"`
	SecurityModuleKeyVersion       string                         `json:"securityModuleKeyVersion" xml:"securityModuleKeyVersion"`
	UIDLampRecognition             UIDLampRecognition             `json:"UIDLampRecognition" xml:"UIDLampRecognition"`
	ConfDeviceIdPrefix             bool                           `json:"confDeviceIdPrefix" xml:"confDeviceIdPrefix"`
	OEMCode                        int                            `json:"oemCode" xml:"oemCode"`
	SimpleAlgorithmVersion         string                         `json:"simpleAlgorithmVersion" xml:"simpleAlgorithmVersion"`
	BootTime                       string                         `json:"bootTime" xml:"bootTime"`
	IntelligentAnalysisEngineModel string                         `json:"intelligentAnalysisEngineModel" xml:"intelligentAnalysisEngineModel"`
	MarketType                     int                            `json:"marketType" xml:"marketType"`
	AlgorithmVersion               string                         `json:"algorithmVersion" xml:"algorithmVersion"`
	Firmware                       string                         `json:"firmware" xml:"firmware"`
	EngineList                     Engine                         `json:"engine" xml:"engine"`
	Platform                       int                            `json:"platform" xml:"platform"`
	PlatformName                   string                         `json:"platformName" xml:"platformName"`
	TouchScreenVersionInfo         string                         `json:"touchScreenVersionInfo" xml:"touchScreenVersionInfo"`
	ProtocolFileURL                string                         `json:"protocolFileURL" xml:"protocolFileURL"`
	RecycleRecordEnabled           bool                           `json:"recycleRecordEnabled" xml:"recycleRecordEnabled"`
	DecordChannelNums              int                            `json:"decordChannelNums" xml:"decordChannelNums"`
	VGANums                        int                            `json:"vGANums" xml:"vGANums"`
	USBNums                        int                            `json:"uSBNums" xml:"uSBNums"`
	AuxoutNums                     int                            `json:"auxoutNums" xml:"auxoutNums"`
	ExpansionBoardVersion          string                         `json:"expansionBoardVersion" xml:"expansionBoardVersion"`
	InitWizzardDisplay             bool                           `json:"initWizzardDisplay" xml:"initWizzardDisplay"`
	BeaconID                       string                         `json:"beaconID" xml:"beaconID"`
	IsResetDeviceLanguage          bool                           `json:"isResetDeviceLanguage" xml:"isResetDeviceLanguage"`
	DispalyNum                     int                            `json:"dispalyNum" xml:"dispalyNum"`
	BspVersion                     string                         `json:"bspVersion" xml:"bspVersion"`
	DspVersion                     string                         `json:"dspVersion" xml:"dspVersion"`
	LocalUIVersion                 string                         `json:"localUIVersion" xml:"localUIVersion"`
	OPCASubType                    string                         `json:"oPCASubType" xml:"oPCASubType"`
	WiegandOutNum                  int                            `json:"wiegandOutNum" xml:"wiegandOutNum"`
	ChipVersionInfoList            []ChipVersionInfo              `json:"ChipVersionInfoList" xml:"ChipVersionInfoList"`
	PersonBagLinkAlgoEngineVersion string                         `json:"personBagLinkAlgoEngineVersion" xml:"personBagLinkAlgoEngineVersion"`
	BIOSVersion                    string                         `json:"BIOSVersion" xml:"BIOSVersion"`
	ContactInformation             string                         `json:"contactInformation" xml:"contactInformation"`
	TemperatureModuleVersionInfo   string                         `json:"temperatureModuleVersionInfo" xml:"temperatureModuleVersionInfo"`
	PedestrianWarningModuleVersion PedestrianWarningModuleVersion `json:"pedestrianWarningModuleVersion" xml:"pedestrianWarningModuleVersion"`
	EncryptionModel                string                         `json:"encryptionModel" xml:"encryptionModel"`
	UWBVersion                     string                         `json:"UWBVersion" xml:"UWBVersion"`
	AudioBoard                     AudioBoard                     `json:"audioBoard" xml:"audioBoard"`
	MaterialScanAlgorithmVersion   string                         `json:"materialScanAlgorithmVersion" xml:"materialScanAlgorithmVersion"`
	RegionVersion                  string                         `json:"regionVersion" xml:"regionVersion"`
	ProductionDate                 string                         `json:"productionDate" xml:"productionDate"`
	WifiModuleMACAddress           string                         `json:"wifiModuleMACAddress" xml:"wifiModuleMACAddress"`
	DisplayInterfaceSize           float64                        `json:"displayInterfaceSize" xml:"displayInterfaceSize"`
	ReleaseRegion                  string                         `json:"releaseRegion" xml:"releaseRegion"`
	ShortSerialNumber              string                         `json:"shortSerialNumber" xml:"shortSerialNumber"`
	AudioVersion                   string                         `json:"audioVersion" xml:"audioVersion"`
	CommunicationFrequency         string                         `json:"communicationFrequency" xml:"communicationFrequency"`
	OSCoreVersionInfo              OSCoreVersionInfo              `json:"OSCoreVersionInfo" xml:"OSCoreVersionInfo"`
	XTransVersion                  string                         `json:"xTransVersion" xml:"xTransVersion"`
	DeviceMaintenanceInfoQRCode    string                         `json:"deviceMaintenanceInfoQRCode" xml:"deviceMaintenanceInfoQRCode"`
	BatteryFirmwareVersion         string                         `json:"batteryFirmwareVersion" xml:"batteryFirmwareVersion"`
	RadarVideoMapInfoList          RadarVideoMapInfo              `json:"radarVideoMapInfoList" xml:"radarVideoMapInfoList"`
}

type RadarVideoMapInfo struct {
	DevIndex             string `json:"devIndex" xml:"devIndex"`
	RadarVideoMapVersion string `json:"radarVideoMapVersion" xml:"radarVideoMapVersion"`
}

type OSCoreVersionInfo struct {
	OSCoreVersion         string `json:"OSCoreVersion" xml:"OSCoreVersion"`
	MinisysVersion        string `json:"minisysVersion" xml:"minisysVersion"`
	NetworkServiceVersion string `json:"networkServiceVersion" xml:"networkServiceVersion"`
	UpgradeServiceVersion string `json:"upgradeServiceVersion" xml:"upgradeServiceVersion"`
}

type AudioBoard struct {
	AudioBoardModel   string `json:"audioBoardModel" xml:"audioBoardModel"`
	AudioBoardVersion string `json:"audioBoardVersion" xml:"audioBoardVersion"`
}

type PedestrianWarningModuleVersion struct {
	PedestrianWarningMCUVersion    string `json:"pedestrianWarningMCUVersion" xml:"pedestrianWarningMCUVersion"`
	PedestrianWarningRadarVersion  string `json:"pedestrianWarningRadarVersion" xml:"pedestrianWarningRadarVersion"`
	PedestrianRangingModuleVersion string `json:"pedestrianRangingModuleVersion" xml:"pedestrianRangingModuleVersion"`
}

type ChipVersionInfo struct {
	ID               int    `json:"id" xml:"id"`
	Type             string `json:"type" xml:"type"`
	FirmwareVersion  string `json:"firmwareVersion" xml:"firmwareVersion"`
	AlgorithmVersion string `json:"algorithmVersion" xml:"algorithmVersion"`
	ChipName         string `json:"chipName" xml:"chipName"`
}

type Engine struct {
	Engine int `json:"engine" xml:"engine"`
}

type UIDLampRecognition struct {
	Enabled bool `json:"enabled" xml:"enabled"`
}

type DockStation struct {
	Platform                    Platform `json:"platform" xml:"platform"`
	CentralStorageBackupEnabled bool     `json:"centralStorageBackupEnabled" xml:"centralStorageBackupEnabled"`
}

type Platform struct {
	Type     string `json:"type" xml:"type"`
	IP       string `json:"ip" xml:"ip"`
	Port     int    `json:"port" xml:"port"`
	Username string `json:"username" xml:"username"`
}

type DetailAbnormalStatus struct {
	HardDiskFull         bool `json:"hardDiskFull" xml:"hardDiskFull"`
	HardDiskError        bool `json:"hardDiskError" xml:"hardDiskError"`
	EthernetBroken       bool `json:"ethernetBroken" xml:"ethernetBroken"`
	IpaddrConflict       bool `json:"ipaddrConflict" xml:"ipaddrConflict"`
	IllegalAccess        bool `json:"illegalAccess" xml:"illegalAccess"`
	RecordError          bool `json:"recordError" xml:"recordError"`
	RaidLogicDiskError   bool `json:"raidLogicDiskError" xml:"raidLogicDiskError"`
	SpareWorkDeviceError bool `json:"spareWorkDeviceError" xml:"spareWorkDeviceError"`
}
