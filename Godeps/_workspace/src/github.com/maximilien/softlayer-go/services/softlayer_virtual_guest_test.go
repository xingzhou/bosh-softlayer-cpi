package services_test

import (
	"errors"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	slclientfakes "github.com/maximilien/softlayer-go/client/fakes"
	common "github.com/maximilien/softlayer-go/common"
	datatypes "github.com/maximilien/softlayer-go/data_types"
	softlayer "github.com/maximilien/softlayer-go/softlayer"
)

var _ = Describe("SoftLayer_Virtual_Guest_Service", func() {
	var (
		username, apiKey string
		err              error

		fakeClient *slclientfakes.FakeSoftLayerClient

		virtualGuestService softlayer.SoftLayer_Virtual_Guest_Service

		virtualGuest         datatypes.SoftLayer_Virtual_Guest
		virtualGuestTemplate datatypes.SoftLayer_Virtual_Guest_Template
	)

	BeforeEach(func() {
		username = os.Getenv("SL_USERNAME")
		Expect(username).ToNot(Equal(""))

		apiKey = os.Getenv("SL_API_KEY")
		Expect(apiKey).ToNot(Equal(""))

		fakeClient = slclientfakes.NewFakeSoftLayerClient(username, apiKey)
		Expect(fakeClient).ToNot(BeNil())

		virtualGuestService, err = fakeClient.GetSoftLayer_Virtual_Guest_Service()
		Expect(err).ToNot(HaveOccurred())
		Expect(virtualGuestService).ToNot(BeNil())

		virtualGuest = datatypes.SoftLayer_Virtual_Guest{}
		virtualGuestTemplate = datatypes.SoftLayer_Virtual_Guest_Template{}
	})

	Context("#GetName", func() {
		It("returns the name for the service", func() {
			name := virtualGuestService.GetName()
			Expect(name).To(Equal("SoftLayer_Virtual_Guest"))
		})
	})

	Context("#CreateObject", func() {
		BeforeEach(func() {
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Service_createObject.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("creates a new SoftLayer_Virtual_Guest instance", func() {
			virtualGuestTemplate = datatypes.SoftLayer_Virtual_Guest_Template{
				Hostname:  "fake-hostname",
				Domain:    "fake.domain.com",
				StartCpus: 2,
				MaxMemory: 1024,
				Datacenter: datatypes.Datacenter{
					Name: "fake-datacenter-name",
				},
				HourlyBillingFlag:            true,
				LocalDiskFlag:                false,
				DedicatedAccountHostOnlyFlag: false,
			}
			virtualGuest, err = virtualGuestService.CreateObject(virtualGuestTemplate)
			Expect(err).ToNot(HaveOccurred())
			Expect(virtualGuest.Hostname).To(Equal("fake-hostname"))
			Expect(virtualGuest.Domain).To(Equal("fake.domain.com"))
			Expect(virtualGuest.StartCpus).To(Equal(2))
			Expect(virtualGuest.MaxMemory).To(Equal(1024))
			Expect(virtualGuest.DedicatedAccountHostOnlyFlag).To(BeFalse())
		})

		It("flags all missing required parameters for SoftLayer_Virtual_Guest/createObject.json POST call", func() {
			virtualGuestTemplate = datatypes.SoftLayer_Virtual_Guest_Template{}
			_, err := virtualGuestService.CreateObject(virtualGuestTemplate)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Hostname"))
			Expect(err.Error()).To(ContainSubstring("Domain"))
			Expect(err.Error()).To(ContainSubstring("StartCpus"))
			Expect(err.Error()).To(ContainSubstring("MaxMemory"))
			Expect(err.Error()).To(ContainSubstring("Datacenter"))
		})
	})

	Context("#GetObject", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Service_getObject.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves SoftLayer_Virtual_Guest instance", func() {
			vg, err := virtualGuestService.GetObject(virtualGuest.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(vg.Id).To(Equal(virtualGuest.Id))
			Expect(vg.AccountId).To(Equal(278444))
			Expect(vg.CreateDate).ToNot(BeNil())
			Expect(vg.DedicatedAccountHostOnlyFlag).To(BeFalse())
			Expect(vg.Domain).To(Equal("softlayer.com"))
			Expect(vg.FullyQualifiedDomainName).To(Equal("bosh-ecpi1.softlayer.com"))
			Expect(vg.Hostname).To(Equal("bosh-ecpi1"))
			Expect(vg.Id).To(Equal(1234567))
			Expect(vg.LastPowerStateId).To(Equal(0))
			Expect(vg.LastVerifiedDate).To(BeNil())
			Expect(vg.MaxCpu).To(Equal(1))
			Expect(vg.MaxCpuUnits).To(Equal("CORE"))
			Expect(vg.MaxMemory).To(Equal(1024))
			Expect(vg.MetricPollDate).To(BeNil())
			Expect(vg.ModifyDate).ToNot(BeNil())
			Expect(vg.StartCpus).To(Equal(1))
			Expect(vg.StatusId).To(Equal(1001))
			Expect(vg.Uuid).To(Equal("85d444ce-55a0-39c0-e17a-f697f223cd8a"))
			Expect(vg.GlobalIdentifier).To(Equal("52145e01-97b6-4312-9c15-dac7f24b6c2a"))
			Expect(vg.PrimaryBackendIpAddress).To(Equal("10.106.192.42"))
			Expect(vg.PrimaryIpAddress).To(Equal("23.246.234.32"))
			Expect(vg.Location.Id).To(Equal(1234567))
			Expect(len(vg.OperatingSystem.Passwords)).To(BeNumerically(">=", 1))
			Expect(vg.OperatingSystem.Passwords[0].Password).To(Equal("test_password"))
			Expect(vg.OperatingSystem.Passwords[0].Username).To(Equal("test_username"))
		})
	})

	Context("#EditObject", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Service_editObject.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("edits an existing SoftLayer_Virtual_Guest instance", func() {
			virtualGuest := datatypes.SoftLayer_Virtual_Guest{
				Notes: "fake-notes",
			}
			edited, err := virtualGuestService.EditObject(virtualGuest.Id, virtualGuest)
			Expect(err).ToNot(HaveOccurred())
			Expect(edited).To(BeTrue())
		})
	})

	Context("#DeleteObject", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
		})

		It("sucessfully deletes the SoftLayer_Virtual_Guest instance", func() {
			fakeClient.DoRawHttpRequestResponse = []byte("true")
			deleted, err := virtualGuestService.DeleteObject(virtualGuest.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(deleted).To(BeTrue())
		})

		It("fails to delete the SoftLayer_Virtual_Guest instance", func() {
			fakeClient.DoRawHttpRequestResponse = []byte("false")
			deleted, err := virtualGuestService.DeleteObject(virtualGuest.Id)
			Expect(err).To(HaveOccurred())
			Expect(deleted).To(BeFalse())
		})
	})

	Context("#AttachEphemeralDisk", func() {
		BeforeEach(func() {
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Product_Order_placeOrder.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("reports error when providing a wrong disk size", func() {
			err := virtualGuestService.AttachEphemeralDisk(123, -1)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Ephemeral disk size can not be negative: -1"))
		})

		It("can attach a local disk without error", func() {
			err := virtualGuestService.AttachEphemeralDisk(123, 25)
			Expect(err).ToNot(HaveOccurred())
		})

		It("reports error when providing a disk size that exceeds the biggest capacity disk SL can provide", func() {
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Service_getUpgradeItemPrices.json")
			err := virtualGuestService.AttachEphemeralDisk(123, 26)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("No proper local disk for size 26"))
		})

	})

	Context("#GetPowerState", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Service_getPowerState.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves SoftLayer_Virtual_Guest_State for RUNNING instance", func() {
			vgPowerState, err := virtualGuestService.GetPowerState(virtualGuest.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(vgPowerState.KeyName).To(Equal("RUNNING"))
		})
	})

	Context("#GetPrimaryIpAddress", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
			fakeClient.DoRawHttpRequestResponse = []byte("159.99.99.99")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves SoftLayer virtual guest's primary IP address instance", func() {
			vgPrimaryIpAddress, err := virtualGuestService.GetPrimaryIpAddress(virtualGuest.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(vgPrimaryIpAddress).To(Equal("159.99.99.99"))
		})
	})

	Context("#GetActiveTransaction", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Service_getActiveTransaction.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves SoftLayer_Provisioning_Version1_Transaction for virtual guest", func() {
			activeTransaction, err := virtualGuestService.GetActiveTransaction(virtualGuest.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(activeTransaction.CreateDate).ToNot(BeNil())
			Expect(activeTransaction.ElapsedSeconds).To(BeNumerically(">", 0))
			Expect(activeTransaction.GuestId).To(Equal(virtualGuest.Id))
			Expect(activeTransaction.Id).To(BeNumerically(">", 0))
		})
	})

	Context("#GetActiveTransactions", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Service_getActiveTransactions.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves an array of SoftLayer_Provisioning_Version1_Transaction for virtual guest", func() {
			activeTransactions, err := virtualGuestService.GetActiveTransactions(virtualGuest.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(activeTransactions)).To(BeNumerically(">", 0))

			for _, activeTransaction := range activeTransactions {
				Expect(activeTransaction.CreateDate).ToNot(BeNil())
				Expect(activeTransaction.ElapsedSeconds).To(BeNumerically(">", 0))
				Expect(activeTransaction.GuestId).To(Equal(virtualGuest.Id))
				Expect(activeTransaction.Id).To(BeNumerically(">", 0))
			}
		})
	})

	Context("#GetSshKeys", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Service_getSshKeys.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully retrieves an array of SoftLayer_Security_Ssh_Key for virtual guest", func() {
			sshKeys, err := virtualGuestService.GetSshKeys(virtualGuest.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(sshKeys)).To(BeNumerically(">", 0))

			for _, sshKey := range sshKeys {
				Expect(sshKey.CreateDate).ToNot(BeNil())
				Expect(sshKey.Fingerprint).To(Equal("f6:c2:9d:57:2f:74:be:a1:db:71:f2:e5:8e:0f:84:7e"))
				Expect(sshKey.Id).To(Equal(84386))
				Expect(sshKey.Key).ToNot(Equal(""))
				Expect(sshKey.Label).To(Equal("TEST:softlayer-go"))
				Expect(sshKey.ModifyDate).To(BeNil())
				Expect(sshKey.Label).To(Equal("TEST:softlayer-go"))
			}
		})
	})

	Context("#PowerCycle", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
		})

		It("sucessfully power cycle virtual guest instance", func() {
			fakeClient.DoRawHttpRequestResponse = []byte("true")

			rebooted, err := virtualGuestService.PowerCycle(virtualGuest.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(rebooted).To(BeTrue())
		})

		It("fails to power cycle virtual guest instance", func() {
			fakeClient.DoRawHttpRequestResponse = []byte("false")

			rebooted, err := virtualGuestService.PowerCycle(virtualGuest.Id)
			Expect(err).To(HaveOccurred())
			Expect(rebooted).To(BeFalse())
		})
	})

	Context("#PowerOff", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
		})

		It("sucessfully power off virtual guest instance", func() {
			fakeClient.DoRawHttpRequestResponse = []byte("true")

			rebooted, err := virtualGuestService.PowerOff(virtualGuest.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(rebooted).To(BeTrue())
		})

		It("fails to power off virtual guest instance", func() {
			fakeClient.DoRawHttpRequestResponse = []byte("false")

			rebooted, err := virtualGuestService.PowerOff(virtualGuest.Id)
			Expect(err).To(HaveOccurred())
			Expect(rebooted).To(BeFalse())
		})
	})

	Context("#PowerOffSoft", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
		})

		It("sucessfully power off soft virtual guest instance", func() {
			fakeClient.DoRawHttpRequestResponse = []byte("true")

			rebooted, err := virtualGuestService.PowerOffSoft(virtualGuest.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(rebooted).To(BeTrue())
		})

		It("fails to power off soft virtual guest instance", func() {
			fakeClient.DoRawHttpRequestResponse = []byte("false")

			rebooted, err := virtualGuestService.PowerOffSoft(virtualGuest.Id)
			Expect(err).To(HaveOccurred())
			Expect(rebooted).To(BeFalse())
		})
	})

	Context("#PowerOn", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
		})

		It("sucessfully power on virtual guest instance", func() {
			fakeClient.DoRawHttpRequestResponse = []byte("true")

			rebooted, err := virtualGuestService.PowerOn(virtualGuest.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(rebooted).To(BeTrue())
		})

		It("fails to power on virtual guest instance", func() {
			fakeClient.DoRawHttpRequestResponse = []byte("false")

			rebooted, err := virtualGuestService.PowerOn(virtualGuest.Id)
			Expect(err).To(HaveOccurred())
			Expect(rebooted).To(BeFalse())
		})
	})

	Context("#RebootDefault", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
		})

		It("sucessfully default reboots virtual guest instance", func() {
			fakeClient.DoRawHttpRequestResponse = []byte("true")

			rebooted, err := virtualGuestService.RebootDefault(virtualGuest.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(rebooted).To(BeTrue())
		})

		It("fails to default reboot virtual guest instance", func() {
			fakeClient.DoRawHttpRequestResponse = []byte("false")

			rebooted, err := virtualGuestService.RebootDefault(virtualGuest.Id)
			Expect(err).To(HaveOccurred())
			Expect(rebooted).To(BeFalse())
		})
	})

	Context("#RebootSoft", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
		})

		It("sucessfully soft reboots virtual guest instance", func() {
			fakeClient.DoRawHttpRequestResponse = []byte("true")

			rebooted, err := virtualGuestService.RebootSoft(virtualGuest.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(rebooted).To(BeTrue())
		})

		It("fails to soft reboot virtual guest instance", func() {
			fakeClient.DoRawHttpRequestResponse = []byte("false")

			rebooted, err := virtualGuestService.RebootSoft(virtualGuest.Id)
			Expect(err).To(HaveOccurred())
			Expect(rebooted).To(BeFalse())
		})
	})

	Context("#RebootHard", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
		})

		It("sucessfully hard reboot virtual guest instance", func() {
			fakeClient.DoRawHttpRequestResponse = []byte("true")

			rebooted, err := virtualGuestService.RebootHard(virtualGuest.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(rebooted).To(BeTrue())
		})

		It("fails to hard reboot virtual guest instance", func() {
			fakeClient.DoRawHttpRequestResponse = []byte("false")

			rebooted, err := virtualGuestService.RebootHard(virtualGuest.Id)
			Expect(err).To(HaveOccurred())
			Expect(rebooted).To(BeFalse())
		})
	})

	Context("#SetUserMetadata", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Service_setMetadata.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully adds metadata strings as a dile to virtual guest's metadata disk", func() {
			retBool, err := virtualGuestService.SetMetadata(virtualGuest.Id, "fake-metadata")
			Expect(err).ToNot(HaveOccurred())

			Expect(retBool).To(BeTrue())
		})
	})

	Context("#GetUserData", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Service_getUserData.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully returns user data for the virtual guest", func() {
			attributes, err := virtualGuestService.GetUserData(virtualGuest.Id)
			Expect(err).ToNot(HaveOccurred())

			Expect(len(attributes)).To(BeNumerically("==", 2))

			Expect(attributes[0].Value).To(Equal("V2hvJ3Mgc21hcnRlcj8gRG1pdHJ5aSBvciBkci5tYXguLi4gIHRoZSBkb2MsIGFueSBkYXkgOik="))
			Expect(attributes[0].Type.Name).To(Equal("User Data"))
			Expect(attributes[0].Type.Keyname).To(Equal("USER_DATA"))

			Expect(attributes[1].Value).To(Equal("ZmFrZS1iYXNlNjQtZGF0YQo="))
			Expect(attributes[1].Type.Name).To(Equal("Fake Data"))
			Expect(attributes[1].Type.Keyname).To(Equal("FAKE_DATA"))
		})
	})

	Context("#IsPingable", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
		})

		Context("when there are no API errors", func() {
			It("checks that the virtual guest instance is pigable", func() {
				fakeClient.DoRawHttpRequestResponse = []byte("true")

				pingable, err := virtualGuestService.IsPingable(virtualGuest.Id)
				Expect(err).ToNot(HaveOccurred())
				Expect(pingable).To(BeTrue())
			})

			It("checks that the virtual guest instance is NOT pigable", func() {
				fakeClient.DoRawHttpRequestResponse = []byte("false")

				pingable, err := virtualGuestService.IsPingable(virtualGuest.Id)
				Expect(err).ToNot(HaveOccurred())
				Expect(pingable).To(BeFalse())
			})
		})

		Context("when there are API errors", func() {
			It("returns false and error", func() {
				fakeClient.DoRawHttpRequestError = errors.New("fake-error")

				pingable, err := virtualGuestService.IsPingable(virtualGuest.Id)
				Expect(err).To(HaveOccurred())
				Expect(pingable).To(BeFalse())
			})
		})

		Context("when the API returns invalid or empty result", func() {
			It("returns false and error", func() {
				fakeClient.DoRawHttpRequestResponse = []byte("fake")

				pingable, err := virtualGuestService.IsPingable(virtualGuest.Id)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Failed to checking that virtual guest is pingable"))
				Expect(pingable).To(BeFalse())
			})
		})
	})

	Context("#ConfigureMetadataDisk", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Service_configureMetadataDisk.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully configures a metadata disk for a virtual guest", func() {
			transaction, err := virtualGuestService.ConfigureMetadataDisk(virtualGuest.Id)
			Expect(err).ToNot(HaveOccurred())

			Expect(transaction.CreateDate).ToNot(BeNil())
			Expect(transaction.ElapsedSeconds).To(Equal(0))
			Expect(transaction.GuestId).To(Equal(virtualGuest.Id))
			Expect(transaction.HardwareId).To(Equal(0))
			Expect(transaction.Id).To(Equal(12476326))
			Expect(transaction.ModifyDate).ToNot(BeNil())
			Expect(transaction.StatusChangeDate).ToNot(BeNil())

			Expect(transaction.TransactionGroup.AverageTimeToComplete).To(Equal("1.62"))
			Expect(transaction.TransactionGroup.Name).To(Equal("Configure Cloud Metadata Disk"))

			Expect(transaction.TransactionStatus.AverageDuration).To(Equal(".32"))
			Expect(transaction.TransactionStatus.FriendlyName).To(Equal("Configure Cloud Metadata Disk"))
			Expect(transaction.TransactionStatus.Name).To(Equal("CLOUD_CONFIGURE_METADATA_DISK"))
		})
	})

	Context("#GetUpgradeItemPrices", func() {
		BeforeEach(func() {
			virtualGuest.Id = 1234567
			fakeClient.DoRawHttpRequestResponse, err = common.ReadJsonTestFixtures("services", "SoftLayer_Virtual_Guest_Service_getUpgradeItemPrices.json")
			Expect(err).ToNot(HaveOccurred())
		})

		It("sucessfully get the upgrade item prices for a virtual guest", func() {
			itemPrices, err := virtualGuestService.GetUpgradeItemPrices(virtualGuest.Id)
			Expect(err).ToNot(HaveOccurred())

			Expect(len(itemPrices)).To(Equal(1))
			Expect(itemPrices[0].Id).To(Equal(12345))
			Expect(itemPrices[0].Categories[0].CategoryCode).To(Equal("guest_disk1"))
		})
	})

})
