package main

import (
	"context"
	"log"
	"os"

	"github.com/sontungpytn/go-sdk/v2/thirdweb"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy [command]",
	Short: "Deploy a contract",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Please input a command to run")
	},
}

var deployNftCmd = &cobra.Command{
	Use:   "nft",
	Short: "Deploy an nft collection",
	Run: func(cmd *cobra.Command, args []string) {
		if thirdwebSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := thirdwebSDK.Deployer.DeployNFTCollection(context.Background(), &thirdweb.DeployNFTCollectionMetadata{
			Name: "Goku NFT",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

var deployEditionCmd = &cobra.Command{
	Use:   "edition",
	Short: "Deploy an edition",
	Run: func(cmd *cobra.Command, args []string) {
		if thirdwebSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := thirdwebSDK.Deployer.DeployEdition(context.Background(), &thirdweb.DeployEditionMetadata{
			Name: "Go SDK",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

var deployTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Deploy an token",
	Run: func(cmd *cobra.Command, args []string) {
		if thirdwebSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := thirdwebSDK.Deployer.DeployToken(context.Background(), &thirdweb.DeployTokenMetadata{
			Name: "Go SDK",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

var deployNFTDropCmd = &cobra.Command{
	Use:   "nftdrop",
	Short: "Deploy an nft drop",
	Run: func(cmd *cobra.Command, args []string) {
		if thirdwebSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := thirdwebSDK.Deployer.DeployNFTDrop(context.Background(), &thirdweb.DeployNFTDropMetadata{
			Name: "Go Script Drop",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

var deployEditionDropCmd = &cobra.Command{
	Use:   "editiondrop",
	Short: "Deploy an edition drop",
	Run: func(cmd *cobra.Command, args []string) {
		if thirdwebSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := thirdwebSDK.Deployer.DeployEditionDrop(context.Background(), &thirdweb.DeployEditionDropMetadata{
			Name: "Go SDK",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

var deployMultiwrapCmd = &cobra.Command{
	Use:   "multiwrap",
	Short: "Deploy a multiwrap",
	Run: func(cmd *cobra.Command, args []string) {
		if thirdwebSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := thirdwebSDK.Deployer.DeployMultiwrap(context.Background(), &thirdweb.DeployMultiwrapMetadata{
			Name: "Go SDK",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

var deployMarketplaceCmd = &cobra.Command{
	Use:   "marketplace",
	Short: "Deploy a marketplace",
	Run: func(cmd *cobra.Command, args []string) {
		if thirdwebSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := thirdwebSDK.Deployer.DeployMarketplace(context.Background(), &thirdweb.DeployMarketplaceMetadata{
			Name: "Go SDK",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

func init() {
	deployCmd.AddCommand(deployNftCmd)
	deployCmd.AddCommand(deployEditionCmd)
	deployCmd.AddCommand(deployTokenCmd)
	deployCmd.AddCommand(deployNFTDropCmd)
	deployCmd.AddCommand(deployEditionDropCmd)
	deployCmd.AddCommand(deployMultiwrapCmd)
	deployCmd.AddCommand(deployMarketplaceCmd)
}
