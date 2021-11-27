package termbg

type Terminal uint

const (
	RxvtCompatible Terminal = iota
	Screen
	Tmux
	XtermCompatible
	Windows
	VSCode
)

func (t Terminal) String() string {
	switch t {
	case RxvtCompatible:
		return "RxvtCompatible"
	case Screen:
		return "Screen"
	case Tmux:
		return "Tmux"
	case XtermCompatible:
		return "XtermCompatible"
	case Windows:
		return "Windows"
	case VSCode:
		return "VSCode"
	default:
		return "unknown"
	}
}
