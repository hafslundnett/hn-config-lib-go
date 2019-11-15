package hid

// HIDclient exp
type HIDclient struct {
	Host   string
	Path   string
	Secret string
}

// New exp
func New(host, path, secret string) *HIDclient { // TODO: clean path and url of /
	return &HIDclient{
		Host:   host,
		Path:   path,
		Secret: secret,
	}
}
