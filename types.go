package smarthome

// EMeter response changed between firmware 1.0 and 1.1. Populate both sets of fields.
func normalizeEmeterResponse(emeter *GetRealtimeResponse) {
	if emeter.CurrentMa == 0 {
		emeter.CurrentMa = int(emeter.CurrentA * 1000)
	} else {
		emeter.CurrentA = float64(emeter.CurrentMa) / 1000
	}

	if emeter.PowerMw == 0 {
		emeter.PowerMw = int(emeter.PowerW * 1000)
	} else {
		emeter.PowerW = float64(emeter.PowerMw) / 1000
	}

	if emeter.TotalWh == 0 {
		emeter.TotalWh = int(emeter.TotalWhFloat)
	} else {
		emeter.TotalWhFloat = float64(emeter.TotalWh)
	}

	if emeter.VoltageMv == 0 {
		emeter.VoltageMv = int(emeter.VoltageV * 1000)
	} else {
		emeter.VoltageV = float64(emeter.VoltageMv) / 1000
	}
}

// Action is when the next action will occur.
type Action struct {
	Type int `json:"type,omitempty"`
}

// SysInfoChild is a child device returned with the SysInfo response.
// For example, each plug on an HS300 is represented as a SysInfoChild entry.
type SysInfoChild struct {
	ID         string `json:"id,omitempty"`
	State      int    `json:"state,omitempty"`
	Alias      string `json:"alias,omitempty"`
	OnTime     int    `json:"on_time,omitempty"`
	NextAction Action `json:"next_action,omitempty"`
}

// SysInfoResponse is the main device information response.
type SysInfoResponse struct {
	ErrCode    int            `json:"err_code,omitempty"`
	SwVer      string         `json:"sw_ver,omitempty"`
	HwVer      string         `json:"hw_ver,omitempty"`
	Type       string         `json:"type,omitempty"`
	Model      string         `json:"model,omitempty"`
	Mac        string         `json:"mac,omitempty"`
	DeviceID   string         `json:"deviceId,omitempty"`
	HwID       string         `json:"hwId,omitempty"`
	FwID       string         `json:"fwId,omitempty"`
	OemID      string         `json:"oemId,omitempty"`
	Alias      string         `json:"alias,omitempty"`
	DevName    string         `json:"dev_name,omitempty"`
	IconHash   string         `json:"icon_hash,omitempty"`
	RelayState int            `json:"relay_state,omitempty"`
	OnTime     int            `json:"on_time,omitempty"`
	ActiveMode string         `json:"active_mode,omitempty"`
	Feature    string         `json:"feature,omitempty"`
	Updating   int            `json:"updating,omitempty"`
	Rssi       int            `json:"rssi,omitempty"`
	LedOff     int            `json:"led_off,omitempty"`
	Latitude   int            `json:"latitude,omitempty"`
	Longitude  int            `json:"longitude,omitempty"`
	MicType    string         `json:"mic_type,omitempty"`
	Children   []SysInfoChild `json:"children,omitempty"`
	ChildNum   int            `json:"child_num,omitempty"`
}

// GetRealtimeResponse is the response for EMeter Get Realtime request.
type GetRealtimeResponse struct {
	VoltageMv int `json:"voltage_mv,omitempty"`
	CurrentMa int `json:"current_ma,omitempty"`
	PowerMw   int `json:"power_mw,omitempty"`
	TotalWh   int `json:"total_wh,omitempty"`
	ErrCode   int `json:"err_code,omitempty"`

	// Older firmware
	CurrentA     float64 `json:"current,omitempty"`
	VoltageV     float64 `json:"voltage,omitempty"`
	PowerW       float64 `json:"power,omitempty"`
	TotalWhFloat float64 `json:"total,omitempty"`
}

// EMeterResponse encapsulates all emeter category responses.
type EMeterResponse struct {
	GetRealtime *GetRealtimeResponse `json:"get_realtime,omitempty"`
}

// SystemResponse encapsulates all system category responses.
type SystemResponse struct {
	SysInfo *SysInfoResponse `json:"get_sysinfo,omitempty"`
}

// GenericResponse is the root object for all responses.
type GenericResponse struct {
	System *SystemResponse `json:"system,omitempty"`
	EMeter *EMeterResponse `json:"emeter,omitempty"`
}

// GetSysinfoRequest requests system info.
type GetSysinfoRequest struct {
}

// ResetRequest requests a device reset after Delay.
type ResetRequest struct {
	Delay int `json:"delay,omitempty"`
}

// SystemRequests encapsulates all system requests.
type SystemRequests struct {
	GetSysinfo *GetSysinfoRequest `json:"get_sysinfo,omitempty"`
	Reset      *ResetRequest      `json:"reset,omitempty"`
}

// GetRealtimeRequest requests current power information.
type GetRealtimeRequest struct {
}

// EMeterRequests encapsulates all emeter requests.
type EMeterRequests struct {
	GetRealtime *GetRealtimeRequest `json:"get_realtime,omitempty"`
}

// Context is used to query child devices.
type Context struct {
	ChildIds []string `json:"child_ids,omitempty"`
}

// GenericRequest is the root for all requests.
type GenericRequest struct {
	Context *Context        `json:"context,omitempty"`
	System  *SystemRequests `json:"system,omitempty"`
	EMeter  *EMeterRequests `json:"emeter,omitempty"`
}
