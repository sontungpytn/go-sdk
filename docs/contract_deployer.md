
## Contract Deployments

The contract deployer lets you deploy new contracts to the blockchain using just the thirdweb SDK. You can access the contract deployer interface as follows:

```
import (
	"github.com/sontungpytn/go-sdk/v2/thirdweb"
)

privateKey := "..."
secretKey := "..."

sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
	PrivateKey: privateKey,
	SecretKey: secretKey
})

// Now you can deploy a contract
address, err := sdk.Deployer.DeployNFTCollection(
	&thirdweb.DeployNFTCollectionMetadata{
		Name: "Go NFT",
	}
})
```

```go
type ContractDeployer struct {
    *ProviderHandler
}
```

### func \(\*ContractDeployer\) [DeployEdition](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/contract_deployer.go#L117>)

```go
func (deployer *ContractDeployer) DeployEdition(ctx context.Context, metadata *DeployEditionMetadata) (string, error)
```

Deploy a new Edition contract.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployEdition(
     context.Background(),
		&thirdweb.DeployEditionMetadata{
			Name: "Go Edition",
		}
	})
```

### func \(\*ContractDeployer\) [DeployEditionDrop](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/contract_deployer.go#L174>)

```go
func (deployer *ContractDeployer) DeployEditionDrop(ctx context.Context, metadata *DeployEditionDropMetadata) (string, error)
```

Deploy a new Edition Drop contract.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployEditionDrop(
     context.Background(),
		&thirdweb.DeployEditionDropMetadata{
			Name: "Go Edition Drop",
		}
	})
```

### func \(\*ContractDeployer\) [DeployMarketplace](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/contract_deployer.go#L212>)

```go
func (deployer *ContractDeployer) DeployMarketplace(ctx context.Context, metadata *DeployMarketplaceMetadata) (string, error)
```

Deploy a new Marketplace contract.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployMarketplace(
     context.Background()
		&thirdweb.DeployMarketplaceMetadata{
			Name: "Go Marketplace",
		}
	})
```

### func \(\*ContractDeployer\) [DeployMultiwrap](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/contract_deployer.go#L193>)

```go
func (deployer *ContractDeployer) DeployMultiwrap(ctx context.Context, metadata *DeployMultiwrapMetadata) (string, error)
```

Deploy a new Multiwrap contract.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployMultiwrap(
     context.Background()
		&thirdweb.DeployMultiwrapMetadata{
			Name: "Go Multiwrap",
		}
	})
```

### func \(\*ContractDeployer\) [DeployNFTCollection](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/contract_deployer.go#L98>)

```go
func (deployer *ContractDeployer) DeployNFTCollection(ctx context.Context, metadata *DeployNFTCollectionMetadata) (string, error)
```

Deploy a new NFT Collection contract.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployNFTCollection(
     context.Background(),
		&thirdweb.DeployNFTCollectionMetadata{
			Name: "Go NFT",
		}
	})
```

### func \(\*ContractDeployer\) [DeployNFTDrop](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/contract_deployer.go#L155>)

```go
func (deployer *ContractDeployer) DeployNFTDrop(ctx context.Context, metadata *DeployNFTDropMetadata) (string, error)
```

Deploy a new NFT Drop contract.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployNFTDrop(
     context.Background(),
		&thirdweb.DeployNFTDropMetadata{
			Name: "Go NFT Drop",
		}
	})
```

### func \(\*ContractDeployer\) [DeployToken](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/contract_deployer.go#L136>)

```go
func (deployer *ContractDeployer) DeployToken(ctx context.Context, metadata *DeployTokenMetadata) (string, error)
```

Deploy a new Token contract.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployToken(
     context.Background(),
		&thirdweb.DeployTokenMetadata{
			Name: "Go Token",
		}
	})
```
