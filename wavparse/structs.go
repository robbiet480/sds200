package wavparse

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Recording struct {
	File     string        `json:",omitempty"`
	Public   *ListChunk    `json:",omitempty"`
	Private  *UnidenChunk  `json:",omitempty"`
	Duration time.Duration `json:",omitempty"`
}
type ListChunk struct {
	System           string     `json:",omitempty"` // IART
	Department       string     `json:",omitempty"` // IGNR
	Channel          string     `json:",omitempty"` // INAM
	TGIDFreq         string     `json:",omitempty"` // ICMT
	Product          string     `json:",omitempty"` // IPRD
	Unknown          string     `json:",omitempty"` // IKEY
	Timestamp        *time.Time `json:",omitempty"` // ICRD
	Tone             string     `json:",omitempty"` // ISRC
	UnitID           string     `json:",omitempty"` // ITCH
	FavoriteListName string     `json:",omitempty"` // ISBJ
	Reserved         string     `json:",omitempty"` // ICOP
}

type FavoriteInfo struct {
	Name            string `json:",omitempty"`
	File            string `json:",omitempty"`
	LocationControl bool
	Monitor         bool
	QuickKey        string `json:",omitempty"`
	NumberTag       string `json:",omitempty"`
	ConfigKey0      string `json:",omitempty"`
	ConfigKey1      string `json:",omitempty"`
	ConfigKey2      string `json:",omitempty"`
	ConfigKey3      string `json:",omitempty"`
	ConfigKey4      string `json:",omitempty"`
	ConfigKey5      string `json:",omitempty"`
	ConfigKey6      string `json:",omitempty"`
	ConfigKey7      string `json:",omitempty"`
	ConfigKey8      string `json:",omitempty"`
	ConfigKey9      string `json:",omitempty"`
}

func (f *FavoriteInfo) UnmarshalBinary(data []byte) error {
	split := strings.Split(string(data), "\x00")

	if len(split) >= 1 && split[0] != "" {
		f.Name = split[0]
	}
	if len(split) >= 2 && split[1] != "" {
		f.File = split[1]
	}
	if len(split) >= 3 && split[2] != "" {
		toggleBool, toggleBoolErr := parseBool(split[2])
		if toggleBoolErr != nil {
			return fmt.Errorf("error when parsing favorite location control toggle to bool: %v", toggleBoolErr)
		}
		f.LocationControl = toggleBool
	}
	if len(split) >= 4 && split[3] != "" {
		toggleBool, toggleBoolErr := parseBool(split[3])
		if toggleBoolErr != nil {
			return fmt.Errorf("error when parsing favorite monitor toggle to bool: %v", toggleBoolErr)
		}
		f.Monitor = toggleBool
	}
	if len(split) >= 5 && split[4] != "" {
		f.QuickKey = split[4]
	}
	if len(split) >= 6 && split[5] != "" {
		f.NumberTag = split[5]
	}
	if len(split) >= 7 && split[6] != "" {
		f.ConfigKey0 = split[6]
	}
	if len(split) >= 8 && split[7] != "" {
		f.ConfigKey1 = split[7]
	}
	if len(split) >= 9 && split[8] != "" {
		f.ConfigKey2 = split[8]
	}
	if len(split) >= 10 && split[9] != "" {
		f.ConfigKey3 = split[9]
	}
	if len(split) >= 11 && split[10] != "" {
		f.ConfigKey4 = split[10]
	}
	if len(split) >= 12 && split[11] != "" {
		f.ConfigKey5 = split[11]
	}
	if len(split) >= 13 && split[12] != "" {
		f.ConfigKey6 = split[12]
	}
	if len(split) >= 14 && split[13] != "" {
		f.ConfigKey7 = split[13]
	}
	if len(split) >= 15 && split[14] != "" {
		f.ConfigKey8 = split[14]
	}
	if len(split) >= 16 && split[15] != "" {
		f.ConfigKey9 = split[15]
	}

	return nil
}

type SiteInfo struct {
	Name             string `json:",omitempty"`
	Avoid            bool
	Latitude         float64
	Longitude        float64
	Range            float64
	Modulation       string `json:",omitempty"`
	MotorolaBandPlan string `json:",omitempty"`
	EDACS            string `json:",omitempty"`
	Shape            string `json:",omitempty"`
	Attenuator       bool
}

func (s *SiteInfo) UnmarshalBinary(data []byte) error {
	split := strings.Split(string(data), "\x00")

	if len(split) >= 1 && split[0] != "" {
		s.Name = split[0]
	}
	if len(split) >= 2 && split[1] != "" {
		var parseErr error
		s.Avoid, parseErr = parseBool(split[1])
		if parseErr != nil {
			return fmt.Errorf("error when parsing site avoid toggle to bool: %v", parseErr)
		}
	}
	if len(split) >= 3 && split[2] != "" {
		var parseErr error
		s.Latitude, parseErr = strconv.ParseFloat(split[2], 64)
		if parseErr != nil {
			return fmt.Errorf("error when parsing site latitude to float64: %v", parseErr)
		}
	}
	if len(split) >= 4 && split[3] != "" {
		var parseErr error
		s.Longitude, parseErr = strconv.ParseFloat(split[3], 64)
		if parseErr != nil {
			return fmt.Errorf("error when parsing site longitude to float64: %v", parseErr)
		}
	}
	if len(split) >= 5 && split[4] != "" {
		var parseErr error
		s.Range, parseErr = strconv.ParseFloat(split[4], 64)
		if parseErr != nil {
			return fmt.Errorf("error when parsing site range to float64: %v", parseErr)
		}
	}
	if len(split) >= 6 && split[5] != "" {
		s.Modulation = split[5]
	}
	if len(split) >= 7 && split[6] != "" {
		s.MotorolaBandPlan = split[6]
	}
	if len(split) >= 8 && split[7] != "" {
		s.EDACS = split[7]
	}
	if len(split) >= 9 && split[8] != "" {
		s.Shape = split[8]
	}
	if len(split) >= 10 && split[9] != "" {
		var parseErr error
		s.Attenuator, parseErr = parseBool(split[9])
		if parseErr != nil {
			return fmt.Errorf("error when parsing site attenuator toggle to bool: %v", parseErr)
		}
	}
	return nil
}

type SystemInfo struct {
	Name                     string `json:",omitempty"`
	Avoid                    bool
	Blank                    string `json:",omitempty"`
	Type                     string `json:",omitempty"`
	IDSearch                 bool
	EmergencyAlertType       string `json:",omitempty"`
	AlertVolume              string `json:",omitempty"`
	MotorolaStatusBit        string `json:",omitempty"`
	P25NAC                   string `json:",omitempty"`
	QuickKey                 string `json:",omitempty"`
	NumberTag                string `json:",omitempty"`
	HoldTime                 string `json:",omitempty"`
	AnalogAGC                string `json:",omitempty"`
	DigitalAGC               string `json:",omitempty"`
	EndCode                  string `json:",omitempty"`
	PriorityID               string `json:",omitempty"`
	EmergencyAlertLightColor string `json:",omitempty"`
	EmergencyAlertCondition  string `json:",omitempty"`
}

func (s *SystemInfo) UnmarshalBinary(data []byte) error {
	split := strings.Split(string(data), "\x00")

	if len(split) >= 1 && split[0] != "" {
		s.Name = split[0]
	}
	if len(split) >= 2 && split[1] != "" {
		var parseErr error
		s.Avoid, parseErr = parseBool(split[1])
		if parseErr != nil {
			return fmt.Errorf("error when parsing system avoid toggle to bool: %v", parseErr)
		}
	}
	if len(split) >= 3 && split[2] != "" {
		s.Blank = split[2]
	}
	if len(split) >= 4 && split[3] != "" {
		s.Type = split[3]
	}
	if len(split) >= 5 && split[4] != "" {
		var parseErr error
		s.IDSearch, parseErr = parseBool(split[4])
		if parseErr != nil {
			return fmt.Errorf("error when parsing system id search toggle to bool: %v", parseErr)
		}
	}
	if len(split) >= 6 && split[5] != "" {
		s.EmergencyAlertType = split[5]
	}
	if len(split) >= 7 && split[6] != "" {
		s.AlertVolume = split[6]
	}
	if len(split) >= 8 && split[7] != "" {
		s.MotorolaStatusBit = split[7]
	}
	if len(split) >= 9 && split[8] != "" {
		s.P25NAC = split[8]
	}
	if len(split) >= 10 && split[9] != "" {
		s.QuickKey = split[9]
	}
	if len(split) >= 11 && split[10] != "" {
		s.NumberTag = split[10]
	}
	if len(split) >= 12 && split[11] != "" {
		s.HoldTime = split[11]
	}
	if len(split) >= 13 && split[12] != "" {
		s.AnalogAGC = split[12]
	}
	if len(split) >= 14 && split[13] != "" {
		s.DigitalAGC = split[13]
	}
	if len(split) >= 15 && split[14] != "" {
		s.EndCode = split[14]
	}
	if len(split) >= 16 && split[15] != "" {
		s.PriorityID = split[15]
	}
	if len(split) >= 17 && split[17] != "" {
		s.EmergencyAlertLightColor = split[16]
	}
	if len(split) >= 18 && split[18] != "" {
		s.EmergencyAlertCondition = split[17]
	}

	return nil
}

type DepartmentInfo struct {
	Name      string `json:",omitempty"`
	Avoid     bool
	Latitude  float64
	Longitude float64
	Range     float64
	Shape     string `json:",omitempty"`
	NumberTag string `json:",omitempty"`
}

func (d *DepartmentInfo) UnmarshalBinary(data []byte) error {
	split := strings.Split(string(data), "\x00")

	if len(split) >= 1 && split[0] != "" {
		d.Name = split[0]
	}
	if len(split) >= 2 && split[1] != "" {
		var parseErr error
		d.Avoid, parseErr = parseBool(split[1])
		if parseErr != nil {
			return fmt.Errorf("error when parsing department avoid toggle to bool: %v", parseErr)
		}
	}
	if len(split) >= 3 && split[2] != "" {
		var parseErr error
		d.Latitude, parseErr = strconv.ParseFloat(split[2], 64)
		if parseErr != nil {
			return fmt.Errorf("error when parsing department latitude to float64: %v", parseErr)
		}
	}
	if len(split) >= 4 && split[3] != "" {
		var parseErr error
		d.Longitude, parseErr = strconv.ParseFloat(split[3], 64)
		if parseErr != nil {
			return fmt.Errorf("error when parsing department longitude to float64: %v", parseErr)
		}
	}
	if len(split) >= 5 && split[4] != "" {
		var parseErr error
		d.Range, parseErr = strconv.ParseFloat(split[4], 64)
		if parseErr != nil {
			return fmt.Errorf("error when parsing department range to float64: %v", parseErr)
		}
	}
	if len(split) >= 6 && split[5] != "" {
		d.Shape = split[5]
	}
	if len(split) >= 7 && split[6] != "" {
		d.NumberTag = split[6]
	}

	return nil
}

type ServiceType int

func (s ServiceType) String() string {
	switch int(s) {
	case 1:
		return "Multi Dispatch"
	case 2:
		return "Law Dispatch"
	case 3:
		return "Fire Dispatch"
	case 4:
		return "EMS Dispatch"
	case 5:
		return "Reserved"
	case 6:
		return "Multi Tac"
	case 7:
		return "Law Tac"
	case 8:
		return "Fire Tac"
	case 9:
		return "EMS Tac"
	case 10:
		return "Reserved"
	case 11:
		return "Interop"
	case 12:
		return "Hospital"
	case 13:
		return "Ham"
	case 14:
		return "Public Works"
	case 15:
		return "Aircraft"
	case 16:
		return "Federal"
	case 17:
		return "Business"
	case 18:
		return "Reserved"
	case 19:
		return "Reserved"
	case 20:
		return "Railroad"
	case 21:
		return "Other"
	case 22:
		return "Multi Talk"
	case 23:
		return "Law Talk"
	case 24:
		return "Fire Talk"
	case 25:
		return "EMS Talk"
	case 26:
		return "Transportation"
	case 27:
		return "Reserved"
	case 28:
		return "Reserved"
	case 29:
		return "Emergency Ops"
	case 30:
		return "Military"
	case 31:
		return "Media"
	case 32:
		return "Schools"
	case 33:
		return "Security"
	case 34:
		return "Utilities"
	case 35:
		return "Reserved"
	case 36:
		return "Reserved"
	case 37:
		return "Corrections"
	case 208:
		return "Custom 1"
	case 209:
		return "Custom 2"
	case 210:
		return "Custom 3"
	case 211:
		return "Custom 4"
	case 212:
		return "Custom 5"
	case 213:
		return "Custom 6"
	case 214:
		return "Custom 7"
	case 215:
		return "Custom 8"
	case 216:
		return "Racing Officials"
	case 217:
		return "Racing Teams"
	case 255:
		return "Unspecified"
	default: // or 0
		return "Unknown"
	}
}

type ChannelInfo struct {
	Name            string `json:",omitempty"`
	Avoid           bool
	TGIDFrequency   string `json:",omitempty"`
	Mode            string `json:",omitempty"`
	ToneCode        string `json:",omitempty"`
	ServiceType     ServiceType
	Attenuator      int    // Conventional systems only
	DelayValue      string `json:",omitempty"`
	VolumeOffset    string `json:",omitempty"`
	AlertToneType   string `json:",omitempty"`
	AlertToneVolume string `json:",omitempty"`
	AlertLightColor string `json:",omitempty"`
	AlertLightType  string `json:",omitempty"`
	NumberTag       string `json:",omitempty"`
	Priority        string `json:",omitempty"`
}

func (c *ChannelInfo) UnmarshalBinary(data []byte) error {
	split := strings.Split(string(data), "\x00")

	if len(split) >= 1 && split[0] != "" {
		c.Name = split[0]
	}
	if len(split) >= 2 && split[1] != "" {
		var parseErr error
		c.Avoid, parseErr = parseBool(split[1])
		if parseErr != nil {
			return fmt.Errorf("error when parsing channel avoid toggle to bool: %v", parseErr)
		}
	}
	if len(split) >= 3 && split[2] != "" {
		c.TGIDFrequency = split[2]
	}
	if len(split) >= 4 && split[3] != "" {
		c.Mode = split[3]
	}
	if len(split) >= 5 && split[4] != "" {
		c.ToneCode = split[4]
	}
	if len(split) >= 6 && split[5] != "" {
		parsed, parseErr := strconv.ParseInt(split[5], 10, 32)
		if parseErr != nil {
			return fmt.Errorf("error when parsing channel service type to int: %v", parseErr)
		}
		c.ServiceType = ServiceType(parsed)
	}

	conventionalOffset := 0

	if len(split) > 15 { // Conventional systems have one extra channel field, Attenuator
		conventionalOffset = 1
		if len(split) >= 7 && split[6] != "" {
			parsed, parseErr := strconv.ParseInt(split[6], 10, 32)
			if parseErr != nil {
				return fmt.Errorf("error when parsing channel attenuator to int: %v", parseErr)
			}
			c.Attenuator = int(parsed)
		}
	}

	if len(split) >= 7 && split[6] != "" {
		c.DelayValue = split[conventionalOffset+6]
	}
	if len(split) >= 8 && split[7] != "" {
		c.VolumeOffset = split[conventionalOffset+7]
	}
	if len(split) >= 9 && split[8] != "" {
		c.AlertToneType = split[conventionalOffset+8]
	}
	if len(split) >= 10 && split[9] != "" {
		c.AlertToneVolume = split[conventionalOffset+9]
	}
	if len(split) >= 11 && split[10] != "" {
		c.AlertLightColor = split[conventionalOffset+10]
	}
	if len(split) >= 12 && split[11] != "" {
		c.AlertLightType = split[conventionalOffset+11]
	}
	if len(split) >= 13 && split[12] != "" {
		c.NumberTag = split[conventionalOffset+12]
	}
	if len(split) >= 14 && split[13] != "" {
		c.Priority = split[conventionalOffset+13]
	}

	return nil
}

type Metadata struct {
	TGID      string `json:",omitempty"`
	Frequency float64
	WACN      string `json:",omitempty"`
	NAC       string `json:",omitempty"`
	UnitID    string `json:",omitempty"`

	RawTGID      string `json:",omitempty"`
	RawFrequency string `json:",omitempty"`
	RawWACN      string `json:",omitempty"`
	RawNAC       string `json:",omitempty"`
	RawUnitID    string `json:",omitempty"`

	FrequencyFmt string `json:",omitempty"`
	WACNFmt      string `json:",omitempty"`
	UnknownFmt   string `json:",omitempty"`
	NACFmt       string `json:",omitempty"`
}

func (t *Metadata) UnmarshalBinary(data []byte) error {
	fmtChunkSplit := strings.Split(string(data[0:65]), "\x00")

	if len(fmtChunkSplit) >= 1 {
		t.RawTGID = fmtChunkSplit[0]
		if len(t.RawTGID) >= 5 {
			t.TGID = t.RawTGID[5:]
		}
	}

	uidStr := string(data[99:110])

	if uidStr[0:4] == "UID:" {
		t.RawUnitID = strings.Split(uidStr, "\x00")[0]
		t.UnitID = t.RawUnitID[4:]
	}

	if len(fmtChunkSplit) >= 3 {
		t.FrequencyFmt = fmtChunkSplit[2]

		if t.FrequencyFmt != "" {
			t.RawFrequency = fmt.Sprintf(t.FrequencyFmt, data[68:70], data[70:72])

			t.RawFrequency = strings.TrimLeft(t.RawFrequency, "0")

			var parseErr error
			t.Frequency, parseErr = strconv.ParseFloat(strings.Split(t.RawFrequency, " ")[0], 64)
			if parseErr != nil {
				return fmt.Errorf("error when parsing metadata raw frequency to float64: %v", parseErr)
			}
		}
	}

	if len(fmtChunkSplit) >= 4 {
		t.WACNFmt = fmtChunkSplit[3]

		t.RawWACN = fmt.Sprintf(t.WACNFmt, data[212:216])

		t.WACN = t.RawWACN[5:]
	}

	if len(fmtChunkSplit) >= 6 {
		t.UnknownFmt = fmtChunkSplit[5]
	}

	if len(fmtChunkSplit) >= 7 {
		t.NACFmt = fmtChunkSplit[6]

		t.RawNAC = fmt.Sprintf(t.NACFmt, data[174:176])
		t.NAC = t.RawNAC[1 : len(t.RawNAC)-2]
	}

	return nil
}

type RawUnidenChunk struct {
	// Start byte 600
	Favorite   [65]byte   // 0-65 	   / 600-665
	System     [65]byte   // 65-130 	 / 665-730
	Department [65]byte   // 130-195 	 / 730-795
	Channel    [65]byte   // 195-260 	 / 795-860
	Site       [65]byte   // 260-325 	 / 860-925
	Empty      [283]byte  // 325-608 	 / 925-1208
	Metadata   [216]byte  // 608-824 	 / 1208-1424
	Remainder  [1224]byte // 824-2048  / 1424-2648
	// Total size is 2048
}

type UnidenChunk struct {
	Favorite   FavoriteInfo
	System     SystemInfo
	Department DepartmentInfo
	Channel    ChannelInfo
	Site       SiteInfo
	Metadata   Metadata
}