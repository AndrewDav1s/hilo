package e2e

import (
	"context"
	"fmt"
	"time"
)

func (s *IntegrationTestSuite) TestPhotonTokenTransfers() {
	// deploy photon ERC20 token contact
	var photonERC20Addr string
	s.Run("deploy_photon_erc20", func() {
		photonERC20Addr = s.deployERC20Token("photon")
	})

	// send 100 photon tokens from Hilo to Ethereum
	s.Run("send_photon_tokens_to_eth", func() {
		ethRecipient := s.chain.validators[1].ethereumKey.address
		s.sendFromHiloToEth(0, ethRecipient, "100photon", "10photon", "3photon")

		hiloEndpoint := fmt.Sprintf("http://%s", s.valResources[0].GetHostPort("1317/tcp"))
		fromAddr := s.chain.validators[0].keyInfo.GetAddress()

		// require the sender's (validator) balance decreased
		balance, err := queryHiloDenomBalance(hiloEndpoint, fromAddr.String(), "photon")
		s.Require().NoError(err)
		s.Require().Equal(99999999887, balance)

		expEthBalance := 100
		ethEndpoint := fmt.Sprintf("http://%s", s.ethResource.GetHostPort("8545/tcp"))

		// require the Ethereum recipient balance increased
		s.Eventually(
			func() bool {
				ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
				defer cancel()

				b, err := queryEthTokenBalance(ctx, ethEndpoint, photonERC20Addr, ethRecipient)
				if err != nil {
					return false
				}

				return b == expEthBalance
			},
			time.Minute,
			5*time.Second,
		)
	})

	// send 100 photon tokens from Ethereum back to Hilo
	s.Run("send_photon_tokens_from_eth", func() {
		s.sendFromEthToHilo(1, photonERC20Addr, s.chain.validators[0].keyInfo.GetAddress().String(), "100")

		hiloEndpoint := fmt.Sprintf("http://%s", s.valResources[0].GetHostPort("1317/tcp"))
		toAddr := s.chain.validators[0].keyInfo.GetAddress()

		// require the original sender's (validator) balance increased
		balance, err := queryHiloDenomBalance(hiloEndpoint, toAddr.String(), "photon")
		s.Require().NoError(err)
		s.Require().Equal(99999999887, balance)
	})
}

func (s *IntegrationTestSuite) TestHiloTokenTransfers() {
	// deploy hilo ERC20 token contract
	var hiloERC20Addr string
	s.Run("deploy_hilo_erc20", func() {
		hiloERC20Addr = s.deployERC20Token("uhilo")
	})

	// send 300 hilo tokens from Hilo to Ethereum
	s.Run("send_uhilo_tokens_to_eth", func() {
		ethRecipient := s.chain.validators[1].ethereumKey.address
		s.sendFromHiloToEth(0, ethRecipient, "300uhilo", "10photon", "7uhilo")

		endpoint := fmt.Sprintf("http://%s", s.valResources[0].GetHostPort("1317/tcp"))
		fromAddr := s.chain.validators[0].keyInfo.GetAddress()

		balance, err := queryHiloDenomBalance(endpoint, fromAddr.String(), "uhilo")
		s.Require().NoError(err)
		s.Require().Equal(9999999693, balance)

		expEthBalance := 300
		ethEndpoint := fmt.Sprintf("http://%s", s.ethResource.GetHostPort("8545/tcp"))

		// require the Ethereum recipient balance increased
		s.Eventually(
			func() bool {
				ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
				defer cancel()

				b, err := queryEthTokenBalance(ctx, ethEndpoint, hiloERC20Addr, ethRecipient)
				if err != nil {
					return false
				}

				return b == expEthBalance
			},
			time.Minute,
			5*time.Second,
		)
	})

	// BUG / TODO: Currently, sending tokens from Ethereum back to the cosmos zone
	// is broken in cases where the native token uses non-zero decimals.
	//
	// Ref: https://github.com/PeggyJV/gravity-bridge/issues/130

	// send 300 hilo tokens Ethereum back to Hilo
	// s.Run("send_uhilo_tokens_from_eth", func() {
	// 	s.sendFromEthToHilo(1, hiloERC20Addr, s.chain.validators[0].keyInfo.GetAddress().String(), "300")

	// 	hiloEndpoint := fmt.Sprintf("http://%s", s.valResources[0].GetHostPort("1317/tcp"))
	// 	toAddr := s.chain.validators[0].keyInfo.GetAddress()

	// 	// require the original sender's (validator) balance increased
	// 	balance, err := queryHiloDenomBalance(hiloEndpoint, toAddr.String(), "uhilo")
	// 	s.Require().NoError(err)
	// 	s.Require().Equal(99999999887, balance)
	// })
}
