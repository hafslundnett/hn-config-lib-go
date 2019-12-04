package cert

// AppendFromSystem returns immediatley on linux systems. Placeholder function for loading certs on Windows.
func (Pool) AppendFromSystem() error {
	return nil
}
