package asset

//AppFilter is a struct to simplify project file query
type AppFilter struct {
	Extension string
	Ignore    []string
}

var (
	AfterEffect = AppFilter{Extension: ".aep", Ignore: []string{"ame", "Auto-Save"}}
	Cinema4D    = AppFilter{Extension: ".c4d", Ignore: []string{"backup"}}
	Hiero       = AppFilter{Extension: ".hrox", Ignore: []string{".autosave"}}
	Maya        = AppFilter{Extension: ".mb", Ignore: []string{}}
	Nuke        = AppFilter{Extension: ".nk", Ignore: []string{".autosave"}}
	Photoshop   = AppFilter{Extension: ".psd", Ignore: []string{}}
	Premiere    = AppFilter{Extension: ".prproj", Ignore: []string{"ame", "Auto-Save"}}
)

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
