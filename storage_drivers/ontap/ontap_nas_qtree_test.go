// Copyright 2020 NetApp, Inc. All Rights Reserved.

package ontap

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	tridentconfig "github.com/netapp/trident/config"
	drivers "github.com/netapp/trident/storage_drivers"
	"github.com/netapp/trident/storage_drivers/ontap/api"
	"github.com/netapp/trident/storage_drivers/ontap/api/azgo"
)

func newNASQtreeStorageDriver(showSensitive *bool) *NASQtreeStorageDriver {
	config := &drivers.OntapStorageDriverConfig{}
	sp := func(s string) *string { return &s }

	config.CommonStorageDriverConfig = &drivers.CommonStorageDriverConfig{}
	config.CommonStorageDriverConfig.DebugTraceFlags = make(map[string]bool)
	config.CommonStorageDriverConfig.DebugTraceFlags["method"] = true
	if showSensitive != nil {
		config.CommonStorageDriverConfig.DebugTraceFlags["sensitive"] = *showSensitive
	}

	config.ManagementLIF = "127.0.0.1"
	config.SVM = "SVM1"
	config.Aggregate = "aggr1"
	config.Username = "ontap-nas-qtree-user"
	config.Password = "password1!"
	config.StorageDriverName = "ontap-nas-economy"
	config.StoragePrefix = sp("test_")

	nasqtreeDriver := &NASQtreeStorageDriver{}
	nasqtreeDriver.Config = *config

	// ClientConfig holds the configuration data for Client objects
	clientConfig := api.ClientConfig{
		ManagementLIF:           config.ManagementLIF,
		SVM:                     "SVM1",
		Username:                "client_username",
		Password:                "client_password",
		DriverContext:           tridentconfig.DriverContext("driverContext"),
		ContextBasedZapiRecords: 100,
		DebugTraceFlags:         nil,
	}

	nasqtreeDriver.API = api.NewClient(clientConfig)
	nasqtreeDriver.Telemetry = &Telemetry{
		Plugin:        nasqtreeDriver.Name(),
		SVM:           nasqtreeDriver.GetConfig().SVM,
		StoragePrefix: *nasqtreeDriver.GetConfig().StoragePrefix,
		Driver:        nasqtreeDriver,
		done:          make(chan struct{}),
	}

	return nasqtreeDriver
}

func TestOntapNasQtreeStorageDriverConfigString(t *testing.T) {

	var qtreeDrivers = []NASQtreeStorageDriver{
		*newNASQtreeStorageDriver(&[]bool{true}[0]),
		*newNASQtreeStorageDriver(&[]bool{false}[0]),
		*newNASQtreeStorageDriver(nil),
	}

	sensitiveIncludeList := map[string]string{
		"username":        "ontap-nas-qtree-user",
		"password":        "password1!",
		"client username": "client_username",
		"client password": "client_password",
	}

	sensitiveExcludeList := map[string]string{
		"some information": "<REDACTED>",
	}

	externalIncludeList := map[string]string{
		"<REDACTED>":                   "<REDACTED>",
		"username":                     "Username:<REDACTED>",
		"password":                     "Password:<REDACTED>",
		"api":                          "API:<REDACTED>",
		"chap username":                "ChapUsername:<REDACTED>",
		"chap initiator secret":        "ChapInitiatorSecret:<REDACTED>",
		"chap target username":         "ChapTargetUsername:<REDACTED>",
		"chap target initiator secret": "ChapTargetInitiatorSecret:<REDACTED>",
	}

	for _, qtreeDriver := range qtreeDrivers {
		sensitive, ok := qtreeDriver.Config.DebugTraceFlags["sensitive"]

		switch {

		case !ok || (ok && !sensitive):
			for key, val := range externalIncludeList {
				assert.Contains(t, qtreeDriver.String(), val,
					"ontap-nas-economy driver does not contain %v", key)
				assert.Contains(t, qtreeDriver.GoString(), val,
					"ontap-nas-economy driver does not contain %v", key)
			}

			for key, val := range sensitiveIncludeList {
				assert.NotContains(t, qtreeDriver.String(), val,
					"ontap-nas-economy driver contains %v", key)
				assert.NotContains(t, qtreeDriver.GoString(), val,
					"ontap-nas-economy driver contains %v", key)
			}

		case ok && sensitive:
			for key, val := range sensitiveIncludeList {
				assert.Contains(t, qtreeDriver.String(), val,
					"ontap-nas-economy driver does not contain %v", key)
				assert.Contains(t, qtreeDriver.GoString(), val,
					"ontap-nas-economy driver does not contain %v", key)
			}

			for key, val := range sensitiveExcludeList {
				assert.NotContains(t, qtreeDriver.String(), val,
					"ontap-nas-economy driver redacts %v", key)
				assert.NotContains(t, qtreeDriver.GoString(), val,
					"ontap-nas-economy driver redacts %v", key)
			}
		}
	}
}

func TestOntapNasQTreeCalculateOptimalFlexVolSize(t *testing.T) {
	tests := []struct {
		name                      string
		flexvol                   string
		newQtreeSize              uint64
		totalDiskLimit            uint64
		percentageSnapshotReserve int
		sizeUsedBySnapshots       int
		expectedFlexvolSize       uint64
	}{
		{
			name:                      "3 gb snapshot add 1 gb",
			newQtreeSize:              1073741824,
			totalDiskLimit:            1073741824,
			percentageSnapshotReserve: 0,
			sizeUsedBySnapshots:       3158216704,
			expectedFlexvolSize:       5305700352,
		},
		{
			name:                      "0 gb snapshot add 4 gb",
			newQtreeSize:              1073741824,
			totalDiskLimit:            3221225472,
			percentageSnapshotReserve: 0,
			sizeUsedBySnapshots:       0,
			expectedFlexvolSize:       4294967296,
		},
		{
			name:                      "1 gb snapshot add 1 gb 20 snap reserve",
			newQtreeSize:              1073741824,
			totalDiskLimit:            1073741824,
			percentageSnapshotReserve: 20,
			sizeUsedBySnapshots:       983871488,
			expectedFlexvolSize:       3131355136,
		},
		{
			name:                      "1 gb snapshot add 7 gb 20 snap reserve",
			newQtreeSize:              7516192768,
			totalDiskLimit:            1073741824,
			percentageSnapshotReserve: 20,
			sizeUsedBySnapshots:       1022541824,
			expectedFlexvolSize:       10737418240,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			volAttrs := &azgo.VolumeAttributesType{
				VolumeSpaceAttributesPtr: &azgo.VolumeSpaceAttributesType{
					PercentageSnapshotReservePtr: &test.percentageSnapshotReserve,
					SizeUsedBySnapshotsPtr:       &test.sizeUsedBySnapshots,
				},
			}
			newFlexvolSize := calculateOptimalSizeForFlexvol(context.Background(), test.flexvol, volAttrs,
				test.newQtreeSize, test.totalDiskLimit)
			assert.Equal(t, test.expectedFlexvolSize, newFlexvolSize)
		})
	}
}
