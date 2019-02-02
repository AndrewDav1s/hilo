package e2e

import (
	"fmt"
	"time"
)

func (s *IntegrationTestSuite) TestTokenTransfers() {
	// deploy hilo ERC20 token contract
	// var hiloERC20Addr string
	s.Run("deploy_hilo_erc20", func() {
		_ = s.deployERC20Token("uhilo", "hilo", "hilo", 6)
	})

	// deploy photon ERC20 token contact
	// var photonERC20Addr string
	s.Run("deploy_photon_erc20", func() {
		_ = s.deployERC20Token("photon", "photon", "photon", 0)
	})

	// send 100 photon tokens from Hilo to Ethereum
	s.Run("send_photon_tokens_to_eth", func() {
		s.sendFromHiloToEth(0, s.chain.validators[1].ethereumKey.address, "100photon", "3photon")

		endpoint := fmt.Sprintf("http://%s", s.valResources[0].GetHostPort("1317/tcp"))
		fromAddr := s.chain.validators[0].keyInfo.GetAddress()

		// require the sender's (validator) balance decreased
		balance, err := queryHiloDenomBalance(endpoint, fromAddr.String(), "photon")
		s.Require().NoError(err)
		s.Require().Equal(balance, 99999999897)

		// TODO/XXX: Test checking Ethereum account balance. This might require
		// creating go bindings to the gravity contract. For now, we sleep enough
		// time for the orchestrator to relay the event to Ethereum.
		time.Sleep(30 * time.Second)
	})

	// TODO: Re-enable once https://github.com/PeggyJV/gravity-bridge/pull/123 is
	// merged and included in a release.
	//
	// Ref: https://github.com/cicizeo/hilo/issues/10
	//
	// send 100 photon tokens from Ethereum back to Hilo
	// s.Run("send_photon_tokens_from_eth", func() {
	// 	s.sendFromEthToHilo(1, photonERC20Addr, s.chain.validators[0].keyInfo.GetAddress().String(), "100")
	// })

	// send 300 hilo tokens from Hilo to Ethereum
	s.Run("send_uhilo_tokens_to_eth", func() {
		s.sendFromHiloToEth(0, s.chain.validators[1].ethereumKey.address, "300uhilo", "7uhilo")

		endpoint := fmt.Sprintf("http://%s", s.valResources[0].GetHostPort("1317/tcp"))
		fromAddr := s.chain.validators[0].keyInfo.GetAddress()

		balance, err := queryHiloDenomBalance(endpoint, fromAddr.String(), "uhilo")
		s.Require().NoError(err)
		s.Require().Equal(balance, 9999999693)

		// TODO/XXX: Test checking Ethereum account balance. This might require
		// creating go bindings to the gravity contract. For now, we sleep enough
		// time for the orchestrator to relay the event to Ethereum.
		time.Sleep(30 * time.Second)
	})

	// TODO: Re-enable once https://github.com/PeggyJV/gravity-bridge/pull/123 is
	// merged and included in a release.
	//
	// Ref: https://github.com/cicizeo/hilo/issues/10
	//
	// send 300 hilo tokens Ethereum back to Hilo
	// s.Run("send_uhilo_tokens_from_eth", func() {
	// 	s.sendFromEthToHilo(1, hiloERC20Addr, s.chain.validators[0].keyInfo.GetAddress().String(), "300")
	// })
}
