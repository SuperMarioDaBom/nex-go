package nex

import (
	"strings"
	"time"
)

// StructureInterface implements all Structure methods
type StructureInterface interface {
	Hierarchy() []StructureInterface
	ExtractFromStream(*StreamIn) error
	Bytes(*StreamOut) []byte
}

// Structure represents a nex Structure type
type Structure struct {
	StructureInterface
}

// Hierarchy returns a Structure hierarchy
func (structure *Structure) Hierarchy() []StructureInterface {
	return make([]StructureInterface, 0)
}

// NullData represents a structure with no data
type NullData struct {
	*Structure
}

// NewNullData returns a new NullData Structure
func NewNullData() *NullData {
	structure := &Structure{}
	nullData := &NullData{structure}

	return nullData
}

// ExtractFromStream does nothing for NullData
func (nullData *NullData) ExtractFromStream(stream *StreamIn) error {
	// Basically do nothing. Does a relative seek with 0
	stream.SeekByte(0, true)

	return nil
}

// Bytes does nothing for NullData
func (nullData *NullData) Bytes(stream *StreamOut) []byte {
	return stream.Bytes()
}

// RVConnectionData represents a nex RVConnectionData type
type RVConnectionData struct {
	stationURL                 string
	specialProtocols           []byte
	stationURLSpecialProtocols string
	time                       uint64

	hierarchy []StructureInterface
	Structure
}

// SetStationURL sets the RVConnectionData station URL
func (rvConnectionData *RVConnectionData) SetStationURL(stationURL string) {
	rvConnectionData.stationURL = stationURL
}

// SetSpecialProtocols sets the RVConnectionData special protocol list (unused by Nintendo)
func (rvConnectionData *RVConnectionData) SetSpecialProtocols(specialProtocols []byte) {
	rvConnectionData.specialProtocols = specialProtocols
}

// SetStationURLSpecialProtocols sets the RVConnectionData special station URL (unused by Nintendo)
func (rvConnectionData *RVConnectionData) SetStationURLSpecialProtocols(stationURLSpecialProtocols string) {
	rvConnectionData.stationURLSpecialProtocols = stationURLSpecialProtocols
}

// SetTime sets the RVConnectionData time
func (rvConnectionData *RVConnectionData) SetTime(time uint64) {
	rvConnectionData.time = time
}

// Bytes encodes the RVConnectionData and returns a byte array
func (rvConnectionData *RVConnectionData) Bytes(stream *StreamOut) []byte {
	stream.WriteString(rvConnectionData.stationURL)
	stream.WriteUInt32LE(0) // Always 0
	stream.WriteString(rvConnectionData.stationURLSpecialProtocols)
	stream.WriteUInt64LE(rvConnectionData.time)

	return stream.Bytes()
}

// NewRVConnectionData returns a new RVConnectionData
func NewRVConnectionData() *RVConnectionData {
	rvConnectionData := &RVConnectionData{}

	return rvConnectionData
}

// DateTime represents a NEX DateTime type
type DateTime struct {
	value uint64
}

// Now gets current time and converts it into a format DateTime can understand
func (datetime *DateTime) Now() uint64 {
	timestamp := time.Now()

	second := timestamp.Second()
	minute := timestamp.Minute()
	hour := timestamp.Hour()
	day := timestamp.Day()
	month := int(timestamp.Month())
	year := timestamp.Year() + 1

	datetime.value = uint64(second | (minute << 6) | (hour << 12) | (day << 17) | (month << 22) | (year << 26))

	return datetime.value
}

// Value returns the stored DateTime time
func (datetime *DateTime) Value() uint64 {
	return datetime.value
}

// NewDateTime returns a new DateTime instance
func NewDateTime(value uint64) *DateTime {
	return &DateTime{value: value}
}

// StationURL contains the data for a NEX station URL
type StationURL struct {
	// Using pointers to check for nil
	scheme        *string
	address       *string
	port          *string
	stream        *string
	sid           *string
	cid           *string
	pid           *string
	transportType *string
	rvcid         *string
	natm          *string
	natf          *string
	upnp          *string
	pmp           *string
	probeinit     *string
	prid          *string
}

// SetScheme sets the StationURL scheme
func (station *StationURL) SetScheme(scheme *string) {
	station.scheme = scheme
}

// SetAddress sets the StationURL address
func (station *StationURL) SetAddress(address *string) {
	station.address = address
}

// SetPort sets the StationURL port
func (station *StationURL) SetPort(port *string) {
	station.port = port
}

// SetStream sets the StationURL stream
func (station *StationURL) SetStream(stream *string) {
	station.stream = stream
}

// SetSID sets the StationURL SID
func (station *StationURL) SetSID(sid *string) {
	station.sid = sid
}

// SetCID sets the StationURL CID
func (station *StationURL) SetCID(cid *string) {
	station.cid = cid
}

// SetPid sets the StationURL PID
func (station *StationURL) SetPid(pid *string) {
	station.pid = pid
}

// SetType sets the StationURL transportType
func (station *StationURL) SetType(transportType *string) {
	station.transportType = transportType
}

// SetRVCID sets the StationURL RVCID
func (station *StationURL) SetRVCID(rvcid *string) {
	station.rvcid = rvcid
}

// SetNatm sets the StationURL Natm
func (station *StationURL) SetNatm(natm *string) {
	station.natm = natm
}

// SetNatf sets the StationURL Natf
func (station *StationURL) SetNatf(natf *string) {
	station.natf = natf
}

// SetUpnp sets the StationURL Upnp
func (station *StationURL) SetUpnp(upnp *string) {
	station.upnp = upnp
}

// SetPmp sets the StationURL Pmp
func (station *StationURL) SetPmp(pmp *string) {
	station.pmp = pmp
}

// SetProbeInit sets the StationURL ProbeInit
func (station *StationURL) SetProbeInit(probeinit *string) {
	station.probeinit = probeinit
}

// SetPRID sets the StationURL PRID
func (station *StationURL) SetPRID(prid *string) {
	station.prid = prid
}

func (station *StationURL) Scheme() string {
	if station.scheme != nil {
		return *station.address
	}
	return ""
}

func (station *StationURL) Address() string {
	if station.address != nil {
		return *station.address
	}
	return ""
}

func (station *StationURL) Port() string {
	if station.port != nil {
		return *station.port
	}
	return ""
}

func (station *StationURL) Stream() string {
	if station.stream != nil {
		return *station.stream
	}
	return ""
}

func (station *StationURL) SID() string {
	if station.sid != nil {
		return *station.sid
	}
	return ""
}

func (station *StationURL) CID() string {
	if station.cid != nil {
		return *station.cid
	}
	return ""
}

func (station *StationURL) PID() string {
	if station.pid != nil {
		return *station.pid
	}
	return ""
}

func (station *StationURL) Type() string {
	if station.transportType != nil {
		return *station.transportType
	}
	return ""
}

func (station *StationURL) RVCID() string {
	if station.rvcid != nil {
		return *station.rvcid
	}
	return ""
}

func (station *StationURL) Natm() string {
	if station.natm != nil {
		return *station.natm
	}
	return ""
}

func (station *StationURL) Natf() string {
	if station.natf != nil {
		return *station.natf
	}
	return ""
}

func (station *StationURL) Upnp() string {
	if station.upnp != nil {
		return *station.upnp
	}
	return ""
}

func (station *StationURL) Pmp() string {
	if station.pmp != nil {
		return *station.pmp
	}
	return ""
}

func (station *StationURL) ProbeInit() string {
	if station.probeinit != nil {
		return *station.probeinit
	}
	return ""
}

func (station *StationURL) PRID() string {
	if station.prid != nil {
		return *station.prid
	}
	return ""
}

// FromString parses the StationURL data from a string
func (station *StationURL) FromString(str string) {
	split := strings.Split(str, ":/")

	station.scheme = &split[0]
	fields := split[1]

	params := strings.Split(fields, ";")

	for i := 0; i < len(params); i++ {
		param := params[i]
		split = strings.Split(param, "=")

		name := split[0]
		value := split[1]

		switch name {
		case "address":
			station.address = &value
		case "port":
			station.port = &value
		case "stream":
			station.stream = &value
		case "sid":
			station.sid = &value
		case "CID":
			station.cid = &value
		case "PID":
			station.pid = &value
		case "type":
			station.transportType = &value
		case "RVCID":
			station.rvcid = &value
		case "natm":
			station.natm = &value
		case "natf":
			station.natf = &value
		case "upnp":
			station.upnp = &value
		case "pmp":
			station.pmp = &value
		case "probeinit":
			station.probeinit = &value
		case "PRID":
			station.prid = &value
		}
	}
}

// EncodeToString encodes the StationURL into a string
func (station *StationURL) EncodeToString() string {
	fields := []string{}

	if station.address != nil {
		fields = append(fields, "address="+*station.address)
	}

	if station.port != nil {
		fields = append(fields, "port="+*station.port)
	}

	if station.stream != nil {
		fields = append(fields, "stream="+*station.stream)
	}

	if station.sid != nil {
		fields = append(fields, "sid="+*station.sid)
	}

	if station.cid != nil {
		fields = append(fields, "CID="+*station.cid)
	}

	if station.pid != nil {
		fields = append(fields, "PID="+*station.pid)
	}

	if station.transportType != nil {
		fields = append(fields, "type="+*station.transportType)
	}

	if station.rvcid != nil {
		fields = append(fields, "RVCID="+*station.rvcid)
	}

	if station.natm != nil {
		fields = append(fields, "natm="+*station.natm)
	}

	if station.natf != nil {
		fields = append(fields, "natf="+*station.natf)
	}

	if station.upnp != nil {
		fields = append(fields, "upnp="+*station.upnp)
	}

	if station.pmp != nil {
		fields = append(fields, "pmp="+*station.pmp)
	}

	if station.probeinit != nil {
		fields = append(fields, "probeinit="+*station.probeinit)
	}

	if station.prid != nil {
		fields = append(fields, "PRID="+*station.prid)
	}

	return *station.scheme + ":/" + strings.Join(fields, ";")
}

// NewStationURL returns a new StationURL instance
func NewStationURL(str string) *StationURL {
	station := &StationURL{}

	if str != "" {
		station.FromString(str)
	}

	return station
}
