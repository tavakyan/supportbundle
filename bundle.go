package supportbundle

// SupportBundle is container for contents of extracted archive
type SupportBundle struct {
	Commands Commands
	Docker   Docker
	Proc     Proc
}

// UnmarshalBundle takes in the path to the bundle and unmarshals its
// contents to a SupportBundle, and returns an error if it fails.
func (sb *SupportBundle) UnmarshalBundle(path string) (err error) {

	// Unmarshal Commands
	err = sb.Commands.Unmarshal(path)
	if err != nil {
		return err
	}

	// Unmarshal Docker
	err = sb.Docker.Unmarshal(path)
	if err != nil {
		return err
	}

	// Unmarshal Proc
	err = sb.Proc.Unmarshal(path)
	if err != nil {
		return err
	}

	return nil
}
