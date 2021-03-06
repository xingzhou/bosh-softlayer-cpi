package vm_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/maximilien/bosh-softlayer-cpi/softlayer/vm"

	common "github.com/maximilien/bosh-softlayer-cpi/common"

	boshlog "github.com/cloudfoundry/bosh-agent/logger"

	fakevm "github.com/maximilien/bosh-softlayer-cpi/softlayer/vm/fakes"
	fakeslclient "github.com/maximilien/softlayer-go/client/fakes"
)

var _ = Describe("SoftLayerFinder", func() {
	var (
		softLayerClient        *fakeslclient.FakeSoftLayerClient
		agentEnvServiceFactory *fakevm.FakeAgentEnvServiceFactory
		logger                 boshlog.Logger
		finder                 SoftLayerFinder
	)

	BeforeEach(func() {
		softLayerClient = fakeslclient.NewFakeSoftLayerClient("fake-username", "fake-api-key")
		agentEnvServiceFactory = &fakevm.FakeAgentEnvServiceFactory{}
		logger = boshlog.NewLogger(boshlog.LevelNone)

		finder = NewSoftLayerFinder(
			softLayerClient,
			agentEnvServiceFactory,
			logger,
		)

		common.SetTestFixtureForFakeSoftLayerClient(softLayerClient, "SoftLayer_Account_Service_getVirtualGuests.json")
	})

	Describe("Find", func() {
		var (
			vmID int
		)

		Context("when the VM ID is valid and existing", func() {
			BeforeEach(func() {
				vmID = 5816394
			})

			It("finds and returns a new SoftLayerVM object with correct ID", func() {
				vm, found, err := finder.Find(vmID)
				Expect(err).ToNot(HaveOccurred())
				Expect(found).To(BeTrue(), "could not find VM")
				Expect(vm.ID()).To(Equal(vmID), "found VM but ID does not match")
			})
		})

		Context("when the VM ID does not exist", func() {
			It("fails finding the VM", func() {
				_, found, _ := finder.Find(000000)
				Expect(found).To(BeFalse())
			})
		})
	})
})
