module github.com/jonZlotnik/usbarsenal

go 1.18

require (
	github.com/usbarmory/crucible v0.0.0-20220412125126-837bdb65a20d
	github.com/usbarmory/tamago v0.0.0-20220714104148-d5b7c14a0fbb
)

require (
	github.com/ghodss/yaml v1.0.0 // indirect
	go.mozilla.org/pkcs7 v0.0.0-20210826202110-33d05740a352 // indirect
	golang.org/x/sync v0.0.0-20220601150217-0de741cfad7f // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/usbarmory/tamago v0.0.0-20220714104148-d5b7c14a0fbb => ./tamago
