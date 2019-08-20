package version

//AppFilter represents posiible filters for search based on Applcation Names
type AppFilter int

const (
	AfterEffects AppFilter = iota
	Cinema4D
	Hiero
	Maya
	Nuke
	Photoshop
	Premiere
)

//AppPattern returns a string to use in Regex
func (a AppFilter) AppPattern() string {
	switch a {
	case AfterEffects:
		return "(^((?!ame|Auto-Save).)+(.aep))"
	case Cinema4D:
		return "(^((?!backup|tex).)+(.c4d))"
	case Hiero:
		return "(^((?!autosave).)+(.hrox))"
	case Maya:
		return "(^((?!ame|Auto-Save).)+(.mb))"
	case Nuke:
		return "(.nk$)"
	case Photoshop:
		return "(.psd)"
	case Premiere:
		return "(^((?!ame|Auto Save).)+(.prproj))"
	default:
		return ""
	}
}

//Filter returns a new slice containing all strings in the slice that satisfy the predicate f.
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
