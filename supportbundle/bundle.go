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

	path += "/default"

	// unmarshal Commands
	err = sb.Commands.Unmarshal(path)
	if err != nil {
		return err
	}

	// unmarshal Docker
	err = sb.Docker.Unmarshal(path)
	if err != nil {
		return err
	}

	// unmarshal Proc
	err = sb.Proc.Unmarshal(path)
	if err != nil {
		return err
	}

	return nil
}
