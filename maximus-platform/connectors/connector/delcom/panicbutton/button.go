package panicbutton

import (
	"errors"
	"github.com/karalabe/hid"
	"repo.nefrosovet.ru/maximus-platform/connectors/connector"
	"runtime"
	"strconv"
	"time"
)

var (
	ErrNoDelcomDevices = errors.New("no delcom devices")
	ErrNoDelcomButton  = errors.New("no delcom button")
)

const (
	DelcomVendorID  uint16 = 0x0FC5
	DelcomProductID uint16 = 0xB080
	ProductType            = "USB FS IO"

	maxPacketLen        = 16
	minGetReportDelay   = 100 * time.Millisecond
	defaultEventChanLen = 10

	ButtonStateUpdate   connector.ConnectorState = "BUTTON_STATE_UPDATE"
	ConnectorTypeButton connector.ConnectorType  = "DELCOM_BTN"
	StatusACTIVE        connector.EventStatus    = "ACTIVE"
	StatusINACTIVE      connector.EventStatus    = "INACTIVE"
	StatusSTILL         connector.EventStatus    = "STILL"

	readDataFlag       = 100 // mandatory
	buttonPressedFlag  = 254
	buttonReleasedFlag = 255
)

type DelcomButton struct {
	deviceType connector.ConnectorType
	serialNo   string

	isButtonActive bool
	dev            *hid.Device
	msgChan        chan connector.ConnectorMessage
	errChan        chan error

	stopChan chan struct{}
}

func NewButton(deviceID string) (connector.Connector, error) {
	var (
		info      hid.DeviceInfo
		productID uint16 = DelcomProductID
	)

	if !hid.Supported() {
		return nil, hid.ErrUnsupportedPlatform
	}

	if deviceID != "" {
		val, err := strconv.ParseUint(deviceID, 16, 16)
		if err != nil {
			return nil, err
		}
		productID = uint16(val)
	}

	delkomInfos := hid.Enumerate(DelcomVendorID, productID)
	if len(delkomInfos) == 0 {
		return nil, ErrNoDelcomDevices
	}
	if runtime.GOOS != "linux" {
		for _, di := range delkomInfos {
			if di.Product == ProductType {
				info = di
				break
			}
		}
	} else {
		info = delkomInfos[0] // FIXME linux doesn't have 'Product' field because it's called 'iProduct'
	}

	if info == (hid.DeviceInfo{}) {
		return nil, ErrNoDelcomButton
	}

	dev, err := info.Open()
	if err != nil {
		return nil, err
	}

	return &DelcomButton{
		deviceType: ConnectorTypeButton,
		serialNo:   dev.Serial,
		dev:        dev,
		msgChan:    make(chan connector.ConnectorMessage, defaultEventChanLen),
		errChan:    make(chan error, defaultEventChanLen),

		stopChan: make(chan struct{}),
	}, nil
}

func (btn *DelcomButton) Listen() (<-chan connector.ConnectorMessage, <-chan error) {
	go func() {
		for {
			select {
			case <-btn.stopChan:
				close(btn.errChan)
				return
			case <-time.After(minGetReportDelay): // button can't be read faster
				data := make([]byte, maxPacketLen)
				data[0] = readDataFlag

				_, err := btn.dev.GetFeatureReport(data)
				if err != nil {
					btn.errChan <- err
					continue
				}

				status := btn.ackReport(data[0])
				if status == StatusACTIVE { // TODO temporary. Must be 'if status != StatusSTILL'
					btn.msgChan <- connector.BuildMessage(
						btn.ConnectorType(), status, ButtonStateUpdate,
					)
				}

				data = nil
			}
		}
	}()
	return btn.msgChan, btn.errChan
}

func (btn *DelcomButton) ConnectorType() connector.ConnectorType {
	return btn.deviceType
}

func (btn *DelcomButton) ackReport(flag byte) connector.EventStatus {
	switch flag {
	case buttonReleasedFlag:
		if !btn.isButtonActive {
			return StatusSTILL
		}
		btn.isButtonActive = false
		return StatusINACTIVE
	case buttonPressedFlag:
		if btn.isButtonActive {
			return StatusSTILL
		}
		btn.isButtonActive = true
		return StatusACTIVE
	default:
		btn.isButtonActive = false
		return StatusSTILL
	}
}

func (btn *DelcomButton) Close() {
	close(btn.stopChan)
	err := btn.dev.Close()
	if err != nil {
		panic(err)
	}
}
