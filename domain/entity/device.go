package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Device struct {
	value_objects.BaseModel
	ID                    string      `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	VesselID              string      `json:"vessel_id" gorm:"size:191;not null;uniqueIndex:vessel_device_type;"`
	Vessel                *Vessel     `json:"vessel"`
	DeviceTypeID          string      `json:"device_type_id" gorm:"size:191;not null;uniqueIndex:vessel_device_type;"`
	DeviceType            *DeviceType `json:"device_type"`
	Port                  string      `json:"port" gorm:"size:100;not null;default:'/dev/ttyAgitators'"`
	IsConstant            bool        `json:"is_constant" gorm:"default:false;"`
	ConstantValue         float32     `json:"constant_value" gorm:"default:0.0;"`
	Factor                int         `json:"factor" gorm:"default:1;"`
	NodeAddress           int         `json:"node_address" gorm:"default:1;"`
	AdditionalNodeAddress int         `json:"additional_node_address" gorm:"default:22;"` // Can be used as GPIO Pin if isConstant is true.
	ReadStart             int         `json:"read_start" gorm:"default:0;"`
	BaudRate              int         `json:"baud_rate"  gorm:"default:9600;"`
	ByteSize              int         `json:"byte_size"  gorm:"default:16;"`
	StopBits              int         `json:"stop_bits"  gorm:"default:2;"`
	TimeOut               float32     `json:"time_out"  gorm:"default:0.5;"`
	Enabled               bool        `json:"enabled" gorm:"default:true;"`
	MessageLength         int         `json:"message_length" gorm:"default:16;"`
	ClearBuffer           bool        `json:"clear_buffers_before_each_transaction" gorm:"default:True;"`
	ClosePort             bool        `json:"close_port_after_each_call" gorm:"default:True;"`
	CommunicationMethod   string      `json:"communication_method" gorm:"size:20;not null;default:'modbus'"`
	CreatedByUsername     string      `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy             *User       `json:"created_by"`
	UpdatedByUsername     string      `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy             *User       `json:"updated_by"`
}

func (Device) Tablename() string {
	return "devices"
}

func (device *Device) BeforeCreate(db *gorm.DB) error {
	device.ID = uuid.New().String()
	return nil
}

func (device *Device) Validate() error {
	errors := map[string]interface{}{}

	if device.VesselID == "" || len(device.VesselID) == 0 {
		errors["vessel"] = "Vessel Required."
	}

	if device.CreatedByUsername == "" || len(device.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Required."
	}

	if device.DeviceTypeID == "" || len(device.DeviceTypeID) == 0 {
		errors["device_type"] = "Device Type Can Not be Empty.\n"
	}

	if device.CommunicationMethod == "" || len(device.CommunicationMethod) == 0 {
		errors["communication_emthod"] = "Communication Method Can Not be Empty.\n"
	}

	if device.UpdatedByUsername == "" || len(device.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Required."
	}

	return utilities.ConvertMapToError(errors)
}
