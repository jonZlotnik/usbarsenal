module github.com/jonZlotnik/usbarsenal

go 1.18

require (
	github.com/usbarmory/crucible v0.0.0-20220412125126-837bdb65a20d
	github.com/usbarmory/tamago v0.0.0-20220714104148-d5b7c14a0fbb
)

require (
	github.com/ghodss/yaml v1.0.0 // indirect
	golang.org/x/sync v0.0.0-20220601150217-0de741cfad7f // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/usbarmory/tamago => github.com/jonZlotnik/tamago v0.0.0-20220803090757-c3c00f1b8561
