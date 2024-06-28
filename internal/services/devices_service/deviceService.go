package devices_service

import (
	"applicant-backend/internal/entities"
	"applicant-backend/internal/services/vulnerability_service"
	"errors"
)

type DeviceService struct {
	vulnService *vulnerability_service.VulnerabilityService
}

func NewDeviceService(vulnService *vulnerability_service.VulnerabilityService) *DeviceService {
	return &DeviceService{
		vulnService: vulnService,
	}
}

func (d *DeviceService) GetDevices() []entities.Device {
	return DUMMY_DEVICES
}

func (d *DeviceService) GetDeviceById(id string) (*entities.Device, error) {
	devices := d.GetDevices()
	for _, device := range devices {
		if device.Oid == id {
			return &device, nil
		}
	}

	return nil, errors.New("Could not find Device with ID: " + id)
}

func (d *DeviceService) GetVulnerabilitiesForDevice(deviceId string) []entities.Vulnerability {
	vulns := d.vulnService.GetVulnerabilities()
	var vulnsForAsset []entities.Vulnerability
	for _, vuln := range vulns {
		if vuln.AssetId == deviceId {
			vulnsForAsset = append(vulnsForAsset, vuln)
		}
	}
	return vulnsForAsset
}
