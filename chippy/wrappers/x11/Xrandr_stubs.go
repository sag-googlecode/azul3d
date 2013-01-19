// +build no_xrandr

package x11

type(
    RROutput struct{}
    RRCrtc int
    RRMode int
    Rotation int
    SizeID int
    XRRModeFlags int
)

const(
    RR_HSyncPositive XRRModeFlags = 0
    RR_HSyncNegative XRRModeFlags = 0
    RR_VSyncPositive XRRModeFlags = 0
    RR_VSyncNegative XRRModeFlags = 0
    RR_Interlace XRRModeFlags = 0
    RR_DoubleScan XRRModeFlags = 0
    RR_CSync XRRModeFlags = 0
    RR_CSyncPositive XRRModeFlags = 0
    RR_CSyncNegative XRRModeFlags = 0
    RR_HSkewPresent XRRModeFlags = 0
    RR_BCast XRRModeFlags = 0
    RR_PixelMultiplex XRRModeFlags = 0
    RR_DoubleClock XRRModeFlags = 0
    RR_ClockDivideBy2 XRRModeFlags = 0
)

type XRRScreenSize struct{}

func (c *XRRScreenSize) Width() int {
    return 0
}

func (c *XRRScreenSize) Height() int {
    return 0
}

func (c *XRRScreenSize) Mwidth() int {
    return 0
}

func (c *XRRScreenSize) Mheight() int {
    return 0
}

type XRRScreenConfiguration struct{}

func XRRQueryExtension(display *Display) (int, int, int) {
    return 0, 0, 0
}

func XRRQueryVersion(display *Display) (int, int, int) {
    return 0, 0, 0
}

func XRRSizes(display *Display, screen int) ([]*XRRScreenSize, func()) {
    return []*XRRScreenSize{}, nil
}

func XRRRates(display *Display, screen int, sizeID int) []int16 {
    return []int16{}
}

func XRRGetScreenInfo(display *Display, window Window) (*XRRScreenConfiguration, func()) {
    return nil, nil
}

func XRRSetScreenConfig(display *Display, config *XRRScreenConfiguration, draw Drawable, size_index int, rotation Rotation, timestamp Time) int {
    return 0
}

func XRRConfigSizes(config *XRRScreenConfiguration) []*XRRScreenSize {
    return []*XRRScreenSize{}
}

func XRRConfigCurrentConfiguration(config *XRRScreenConfiguration) (SizeID, Rotation) {
    return 0, 0
}

func XRRRootToScreen(display *Display, root Window) int {
    return 0
}

func XRRSelectInput(display *Display, window Window, mask int) {
}

func XRRSetScreenConfigAndRate(display *Display, config *XRRScreenConfiguration, draw Drawable, size_index int, rotation Rotation, rate int16, timestamp Time) int {
    return 0
}

func XRRConfigRates(config *XRRScreenConfiguration, sizeID int) []int16 {
    return []int16{}
}

func XRRConfigCurrentRate(config *XRRScreenConfiguration) int16 {
    return 0
}

type XRRModeInfo struct{}

func (c *XRRModeInfo) Id() RRMode {
    return 0
}

func (c *XRRModeInfo) Width() uint {
    return 0
}

func (c *XRRModeInfo) Height() uint {
    return 0
}

func (c *XRRModeInfo) DotClock() uint {
    return 0
}

func (c *XRRModeInfo) HTotal() uint {
    return 0
}

func (c *XRRModeInfo) VTotal() uint {
    return 0
}

func (c *XRRModeInfo) ModeFlags() XRRModeFlags {
    return 0
}

type XRRScreenResources struct{}

func (c *XRRScreenResources) Crtcs() []RRCrtc {
    return []RRCrtc{}
}

func (c *XRRScreenResources) Modes() []*XRRModeInfo {
    return []*XRRModeInfo{}
}

func XRRGetScreenResources(display *Display, window Window) (*XRRScreenResources, func()) {
    return nil, nil
}

type XRROutputInfo struct{}

func (c *XRROutputInfo) Mm_width() uint {
    return 0
}

func (c *XRROutputInfo) Mm_height() uint {
    return 0
}

func (c *XRROutputInfo) Modes() []RRMode {
    return []RRMode{}
}

func XRRGetOutputInfo(display *Display, resources *XRRScreenResources, output RROutput) (*XRROutputInfo, func()) {
    return nil, nil
}

type XRRCrtcInfo struct{}

func (c *XRRCrtcInfo) Free() {
}

func (c *XRRCrtcInfo) X() int {
    return 0
}

func (c *XRRCrtcInfo) Y() int {
    return 0
}

func (c *XRRCrtcInfo) Width() uint {
    return 0
}

func (c *XRRCrtcInfo) Height() uint {
    return 0
}

func (c *XRRCrtcInfo) Mode() RRMode {
    return 0
}

func (c *XRRCrtcInfo) Rotation() Rotation {
    return 0
}

func (c *XRRCrtcInfo) Outputs() []RROutput {
    return []RROutput{}
}

func XRRGetCrtcInfo(display *Display, resources *XRRScreenResources, crtc RRCrtc) *XRRCrtcInfo {
    return nil
}

func XRRSetCrtcConfig(display *Display, resources *XRRScreenResources, crtc RRCrtc, timestamp Time, x, y int, mode RRMode, rotation Rotation, outputs []RROutput) int {
    return 0
}

func XRRGetCrtcGammaSize(display *Display, crtc RRCrtc) int {
    return 0
}

type XRRCrtcGamma struct {
    Size int
    Red, Green, Blue []uint16
}

func XRRGetCrtcGamma(display *Display, crtc RRCrtc) *XRRCrtcGamma {
    return nil
}

func XRRSetCrtcGamma(display *Display, crtc RRCrtc, gamma *XRRCrtcGamma) {
}

func XRRGetScreenResourcesCurrent(display *Display, window Window) (*XRRScreenResources, func()) {
    return nil, nil
}

