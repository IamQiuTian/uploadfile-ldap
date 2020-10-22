package g


// ConfigFile holds the user supplied configuration file - it is placed here since it is a global
var ConfigFile *string

// Config is the structure of the TOML config structure
var Config TomlConfig

type TomlConfig struct {
	Listen ListenConfig `toml:"listen"`
	LDAP   LdapConfig   `toml:"ldap"`
	Upload UploadConfig `toml:"upload"`
}

type ListenConfig struct {
	SSL  bool
	Cert string
	Key  string
	Port int
}

type LdapConfig struct {
	UseLDAP      bool
	Host         string
	Port         int
	Base         string
	GroupBase    string
	BindDN       string
	BindPassword string
	GroupName    string
}

type UploadConfig struct {
	Path string
}

type Gofilepath struct {
	Path string
}

var MySigningKey = []byte("daksjds545klakjo")

