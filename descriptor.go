package securepassctl

// ApplicationDescriptor describes the basic attributes of
// a Securepass application
type ApplicationDescriptor struct {
	Label            string `json:"label"`
	Realm            string `json:"realm"`
	Group            string `json:"group"`
	Write            bool   `json:"write"`
	AllowNetworkIPv4 string `json:"allow_network_ipv4"`
	AllowNetworkIPv6 string `json:"allow_network_ipv6"`
	Privacy          bool   `json:"privacy"`
}

// UserDescriptor defines the attributes of SecurePass users
type UserDescriptor struct {
	Username string `json:"USERNAME"`
	Name     string `json:"NAME"`
	Surname  string `json:"SURNAME"`
	Email    string `json:"EMAIL"`
	Mobile   string `json:"MOBILE"`
	Nin      string `json:"NIN"`
	Rfid     string `json:"RFID"`
	Manager  string `json:"MANAGER"`
	Type     string
	Password bool
	Enabled  bool
	Token    string
}

// XattrsDescriptor defines a set of extended attributes
type XattrsDescriptor map[string]interface{}

// RadiusDescriptor defines the attributes of SecurePass RADIUS devices
type RadiusDescriptor struct {
	Radius string
	Name   string
	Secret string
	Group  string
	Realm  string
	Rfid   bool
}

// GroupDescriptor defines attributes of Group
type GroupDescriptor struct {
	Group       string
	Description string
}

// LogEntry is a SecurePass application's log entry
type LogEntry struct {
	// SecurePass response is currently broken, this
	// should be a time.Time object.
	Timestamp string
	UUID      string
	Message   string
	Level     int
	App       string
	Realm     string
}

// LogEntriesByTimestamp sorts log entries by timestamp
type LogEntriesByTimestamp []LogEntry

func (l LogEntriesByTimestamp) Less(i, j int) bool { return l[i].Timestamp < l[j].Timestamp }
func (l LogEntriesByTimestamp) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l LogEntriesByTimestamp) Len() int           { return len(l) }

// NSSConfig encapsulates the SecurePass's config section '[nss]'
type NSSConfig struct {
	Realm        string `ini:"realm"`
	DefaultGid   int    `ini:"default_gid"`
	DefaultHome  string `ini:"default_home"`
	DefaultShell string `ini:"default_shell"`
}

// SSHConfig encapsulates the SecurePass's config section '[ssh]'
type SSHConfig struct {
	Root string `ini:"root"`
}

// GlobalConfig encapsulates the SecurePass's whole configuration
type GlobalConfig struct {
	SecurePass `ini:"default"`
	NSSConfig  `ini:"nss"`
	SSHConfig  `ini:"ssh"`
}
