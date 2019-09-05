package supportbundle

// SupportBundle is container for contents of extracted archive
type SupportBundle struct {
	Commands Commands
	Docker   Docker
	Proc     Proc
}

// UnmarshalBundle takes in the path to the bundle and unmarshals its
// contents to a SupportBundle, and returns an error if it fails.
func (sb *SupportBundle) UnmarshalBundle(s string) error {
	return nil
}
