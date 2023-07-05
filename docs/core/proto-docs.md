 <!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [nemo/auction/v1beta1/auction.proto](#nemo/auction/v1beta1/auction.proto)
    - [BaseAuction](#nemo.auction.v1beta1.BaseAuction)
    - [CollateralAuction](#nemo.auction.v1beta1.CollateralAuction)
    - [DebtAuction](#nemo.auction.v1beta1.DebtAuction)
    - [SurplusAuction](#nemo.auction.v1beta1.SurplusAuction)
    - [WeightedAddresses](#nemo.auction.v1beta1.WeightedAddresses)
  
- [nemo/auction/v1beta1/genesis.proto](#nemo/auction/v1beta1/genesis.proto)
    - [GenesisState](#nemo.auction.v1beta1.GenesisState)
    - [Params](#nemo.auction.v1beta1.Params)
  
- [nemo/auction/v1beta1/query.proto](#nemo/auction/v1beta1/query.proto)
    - [QueryAuctionRequest](#nemo.auction.v1beta1.QueryAuctionRequest)
    - [QueryAuctionResponse](#nemo.auction.v1beta1.QueryAuctionResponse)
    - [QueryAuctionsRequest](#nemo.auction.v1beta1.QueryAuctionsRequest)
    - [QueryAuctionsResponse](#nemo.auction.v1beta1.QueryAuctionsResponse)
    - [QueryNextAuctionIDRequest](#nemo.auction.v1beta1.QueryNextAuctionIDRequest)
    - [QueryNextAuctionIDResponse](#nemo.auction.v1beta1.QueryNextAuctionIDResponse)
    - [QueryParamsRequest](#nemo.auction.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#nemo.auction.v1beta1.QueryParamsResponse)
  
    - [Query](#nemo.auction.v1beta1.Query)
  
- [nemo/auction/v1beta1/tx.proto](#nemo/auction/v1beta1/tx.proto)
    - [MsgPlaceBid](#nemo.auction.v1beta1.MsgPlaceBid)
    - [MsgPlaceBidResponse](#nemo.auction.v1beta1.MsgPlaceBidResponse)
  
    - [Msg](#nemo.auction.v1beta1.Msg)
  
- [nemo/bep3/v1beta1/bep3.proto](#nemo/bep3/v1beta1/bep3.proto)
    - [AssetParam](#nemo.bep3.v1beta1.AssetParam)
    - [AssetSupply](#nemo.bep3.v1beta1.AssetSupply)
    - [AtomicSwap](#nemo.bep3.v1beta1.AtomicSwap)
    - [Params](#nemo.bep3.v1beta1.Params)
    - [SupplyLimit](#nemo.bep3.v1beta1.SupplyLimit)
  
    - [SwapDirection](#nemo.bep3.v1beta1.SwapDirection)
    - [SwapStatus](#nemo.bep3.v1beta1.SwapStatus)
  
- [nemo/bep3/v1beta1/genesis.proto](#nemo/bep3/v1beta1/genesis.proto)
    - [GenesisState](#nemo.bep3.v1beta1.GenesisState)
  
- [nemo/bep3/v1beta1/query.proto](#nemo/bep3/v1beta1/query.proto)
    - [AssetSupplyResponse](#nemo.bep3.v1beta1.AssetSupplyResponse)
    - [AtomicSwapResponse](#nemo.bep3.v1beta1.AtomicSwapResponse)
    - [QueryAssetSuppliesRequest](#nemo.bep3.v1beta1.QueryAssetSuppliesRequest)
    - [QueryAssetSuppliesResponse](#nemo.bep3.v1beta1.QueryAssetSuppliesResponse)
    - [QueryAssetSupplyRequest](#nemo.bep3.v1beta1.QueryAssetSupplyRequest)
    - [QueryAssetSupplyResponse](#nemo.bep3.v1beta1.QueryAssetSupplyResponse)
    - [QueryAtomicSwapRequest](#nemo.bep3.v1beta1.QueryAtomicSwapRequest)
    - [QueryAtomicSwapResponse](#nemo.bep3.v1beta1.QueryAtomicSwapResponse)
    - [QueryAtomicSwapsRequest](#nemo.bep3.v1beta1.QueryAtomicSwapsRequest)
    - [QueryAtomicSwapsResponse](#nemo.bep3.v1beta1.QueryAtomicSwapsResponse)
    - [QueryParamsRequest](#nemo.bep3.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#nemo.bep3.v1beta1.QueryParamsResponse)
  
    - [Query](#nemo.bep3.v1beta1.Query)
  
- [nemo/bep3/v1beta1/tx.proto](#nemo/bep3/v1beta1/tx.proto)
    - [MsgClaimAtomicSwap](#nemo.bep3.v1beta1.MsgClaimAtomicSwap)
    - [MsgClaimAtomicSwapResponse](#nemo.bep3.v1beta1.MsgClaimAtomicSwapResponse)
    - [MsgCreateAtomicSwap](#nemo.bep3.v1beta1.MsgCreateAtomicSwap)
    - [MsgCreateAtomicSwapResponse](#nemo.bep3.v1beta1.MsgCreateAtomicSwapResponse)
    - [MsgRefundAtomicSwap](#nemo.bep3.v1beta1.MsgRefundAtomicSwap)
    - [MsgRefundAtomicSwapResponse](#nemo.bep3.v1beta1.MsgRefundAtomicSwapResponse)
  
    - [Msg](#nemo.bep3.v1beta1.Msg)
  
- [nemo/cdp/v1beta1/cdp.proto](#nemo/cdp/v1beta1/cdp.proto)
    - [CDP](#nemo.cdp.v1beta1.CDP)
    - [Deposit](#nemo.cdp.v1beta1.Deposit)
    - [OwnerCDPIndex](#nemo.cdp.v1beta1.OwnerCDPIndex)
    - [TotalCollateral](#nemo.cdp.v1beta1.TotalCollateral)
    - [TotalPrincipal](#nemo.cdp.v1beta1.TotalPrincipal)
  
- [nemo/cdp/v1beta1/genesis.proto](#nemo/cdp/v1beta1/genesis.proto)
    - [CollateralParam](#nemo.cdp.v1beta1.CollateralParam)
    - [DebtParam](#nemo.cdp.v1beta1.DebtParam)
    - [GenesisAccumulationTime](#nemo.cdp.v1beta1.GenesisAccumulationTime)
    - [GenesisState](#nemo.cdp.v1beta1.GenesisState)
    - [GenesisTotalPrincipal](#nemo.cdp.v1beta1.GenesisTotalPrincipal)
    - [Params](#nemo.cdp.v1beta1.Params)
  
- [nemo/cdp/v1beta1/query.proto](#nemo/cdp/v1beta1/query.proto)
    - [CDPResponse](#nemo.cdp.v1beta1.CDPResponse)
    - [QueryAccountsRequest](#nemo.cdp.v1beta1.QueryAccountsRequest)
    - [QueryAccountsResponse](#nemo.cdp.v1beta1.QueryAccountsResponse)
    - [QueryCdpRequest](#nemo.cdp.v1beta1.QueryCdpRequest)
    - [QueryCdpResponse](#nemo.cdp.v1beta1.QueryCdpResponse)
    - [QueryCdpsRequest](#nemo.cdp.v1beta1.QueryCdpsRequest)
    - [QueryCdpsResponse](#nemo.cdp.v1beta1.QueryCdpsResponse)
    - [QueryDepositsRequest](#nemo.cdp.v1beta1.QueryDepositsRequest)
    - [QueryDepositsResponse](#nemo.cdp.v1beta1.QueryDepositsResponse)
    - [QueryParamsRequest](#nemo.cdp.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#nemo.cdp.v1beta1.QueryParamsResponse)
    - [QueryTotalCollateralRequest](#nemo.cdp.v1beta1.QueryTotalCollateralRequest)
    - [QueryTotalCollateralResponse](#nemo.cdp.v1beta1.QueryTotalCollateralResponse)
    - [QueryTotalPrincipalRequest](#nemo.cdp.v1beta1.QueryTotalPrincipalRequest)
    - [QueryTotalPrincipalResponse](#nemo.cdp.v1beta1.QueryTotalPrincipalResponse)
  
    - [Query](#nemo.cdp.v1beta1.Query)
  
- [nemo/cdp/v1beta1/tx.proto](#nemo/cdp/v1beta1/tx.proto)
    - [MsgCreateCDP](#nemo.cdp.v1beta1.MsgCreateCDP)
    - [MsgCreateCDPResponse](#nemo.cdp.v1beta1.MsgCreateCDPResponse)
    - [MsgDeposit](#nemo.cdp.v1beta1.MsgDeposit)
    - [MsgDepositResponse](#nemo.cdp.v1beta1.MsgDepositResponse)
    - [MsgDrawDebt](#nemo.cdp.v1beta1.MsgDrawDebt)
    - [MsgDrawDebtResponse](#nemo.cdp.v1beta1.MsgDrawDebtResponse)
    - [MsgLiquidate](#nemo.cdp.v1beta1.MsgLiquidate)
    - [MsgLiquidateResponse](#nemo.cdp.v1beta1.MsgLiquidateResponse)
    - [MsgRepayDebt](#nemo.cdp.v1beta1.MsgRepayDebt)
    - [MsgRepayDebtResponse](#nemo.cdp.v1beta1.MsgRepayDebtResponse)
    - [MsgWithdraw](#nemo.cdp.v1beta1.MsgWithdraw)
    - [MsgWithdrawResponse](#nemo.cdp.v1beta1.MsgWithdrawResponse)
  
    - [Msg](#nemo.cdp.v1beta1.Msg)
  
- [nemo/committee/v1beta1/committee.proto](#nemo/committee/v1beta1/committee.proto)
    - [BaseCommittee](#nemo.committee.v1beta1.BaseCommittee)
    - [MemberCommittee](#nemo.committee.v1beta1.MemberCommittee)
    - [TokenCommittee](#nemo.committee.v1beta1.TokenCommittee)
  
    - [TallyOption](#nemo.committee.v1beta1.TallyOption)
  
- [nemo/committee/v1beta1/genesis.proto](#nemo/committee/v1beta1/genesis.proto)
    - [GenesisState](#nemo.committee.v1beta1.GenesisState)
    - [Proposal](#nemo.committee.v1beta1.Proposal)
    - [Vote](#nemo.committee.v1beta1.Vote)
  
    - [VoteType](#nemo.committee.v1beta1.VoteType)
  
- [nemo/committee/v1beta1/permissions.proto](#nemo/committee/v1beta1/permissions.proto)
    - [AllowedParamsChange](#nemo.committee.v1beta1.AllowedParamsChange)
    - [CommunityCDPRepayDebtPermission](#nemo.committee.v1beta1.CommunityCDPRepayDebtPermission)
    - [CommunityCDPWithdrawCollateralPermission](#nemo.committee.v1beta1.CommunityCDPWithdrawCollateralPermission)
    - [CommunityPoolLendWithdrawPermission](#nemo.committee.v1beta1.CommunityPoolLendWithdrawPermission)
    - [GodPermission](#nemo.committee.v1beta1.GodPermission)
    - [ParamsChangePermission](#nemo.committee.v1beta1.ParamsChangePermission)
    - [SoftwareUpgradePermission](#nemo.committee.v1beta1.SoftwareUpgradePermission)
    - [SubparamRequirement](#nemo.committee.v1beta1.SubparamRequirement)
    - [TextPermission](#nemo.committee.v1beta1.TextPermission)
  
- [nemo/committee/v1beta1/proposal.proto](#nemo/committee/v1beta1/proposal.proto)
    - [CommitteeChangeProposal](#nemo.committee.v1beta1.CommitteeChangeProposal)
    - [CommitteeDeleteProposal](#nemo.committee.v1beta1.CommitteeDeleteProposal)
  
- [nemo/committee/v1beta1/query.proto](#nemo/committee/v1beta1/query.proto)
    - [QueryCommitteeRequest](#nemo.committee.v1beta1.QueryCommitteeRequest)
    - [QueryCommitteeResponse](#nemo.committee.v1beta1.QueryCommitteeResponse)
    - [QueryCommitteesRequest](#nemo.committee.v1beta1.QueryCommitteesRequest)
    - [QueryCommitteesResponse](#nemo.committee.v1beta1.QueryCommitteesResponse)
    - [QueryNextProposalIDRequest](#nemo.committee.v1beta1.QueryNextProposalIDRequest)
    - [QueryNextProposalIDResponse](#nemo.committee.v1beta1.QueryNextProposalIDResponse)
    - [QueryProposalRequest](#nemo.committee.v1beta1.QueryProposalRequest)
    - [QueryProposalResponse](#nemo.committee.v1beta1.QueryProposalResponse)
    - [QueryProposalsRequest](#nemo.committee.v1beta1.QueryProposalsRequest)
    - [QueryProposalsResponse](#nemo.committee.v1beta1.QueryProposalsResponse)
    - [QueryRawParamsRequest](#nemo.committee.v1beta1.QueryRawParamsRequest)
    - [QueryRawParamsResponse](#nemo.committee.v1beta1.QueryRawParamsResponse)
    - [QueryTallyRequest](#nemo.committee.v1beta1.QueryTallyRequest)
    - [QueryTallyResponse](#nemo.committee.v1beta1.QueryTallyResponse)
    - [QueryVoteRequest](#nemo.committee.v1beta1.QueryVoteRequest)
    - [QueryVoteResponse](#nemo.committee.v1beta1.QueryVoteResponse)
    - [QueryVotesRequest](#nemo.committee.v1beta1.QueryVotesRequest)
    - [QueryVotesResponse](#nemo.committee.v1beta1.QueryVotesResponse)
  
    - [Query](#nemo.committee.v1beta1.Query)
  
- [nemo/committee/v1beta1/tx.proto](#nemo/committee/v1beta1/tx.proto)
    - [MsgSubmitProposal](#nemo.committee.v1beta1.MsgSubmitProposal)
    - [MsgSubmitProposalResponse](#nemo.committee.v1beta1.MsgSubmitProposalResponse)
    - [MsgVote](#nemo.committee.v1beta1.MsgVote)
    - [MsgVoteResponse](#nemo.committee.v1beta1.MsgVoteResponse)
  
    - [Msg](#nemo.committee.v1beta1.Msg)
  
- [nemo/community/v1beta1/proposal.proto](#nemo/community/v1beta1/proposal.proto)
    - [CommunityCDPRepayDebtProposal](#nemo.community.v1beta1.CommunityCDPRepayDebtProposal)
    - [CommunityCDPWithdrawCollateralProposal](#nemo.community.v1beta1.CommunityCDPWithdrawCollateralProposal)
    - [CommunityPoolLendDepositProposal](#nemo.community.v1beta1.CommunityPoolLendDepositProposal)
    - [CommunityPoolLendWithdrawProposal](#nemo.community.v1beta1.CommunityPoolLendWithdrawProposal)
  
- [nemo/community/v1beta1/query.proto](#nemo/community/v1beta1/query.proto)
    - [QueryBalanceRequest](#nemo.community.v1beta1.QueryBalanceRequest)
    - [QueryBalanceResponse](#nemo.community.v1beta1.QueryBalanceResponse)
    - [QueryTotalBalanceRequest](#nemo.community.v1beta1.QueryTotalBalanceRequest)
    - [QueryTotalBalanceResponse](#nemo.community.v1beta1.QueryTotalBalanceResponse)
  
    - [Query](#nemo.community.v1beta1.Query)
  
- [nemo/community/v1beta1/tx.proto](#nemo/community/v1beta1/tx.proto)
    - [MsgFundCommunityPool](#nemo.community.v1beta1.MsgFundCommunityPool)
    - [MsgFundCommunityPoolResponse](#nemo.community.v1beta1.MsgFundCommunityPoolResponse)
  
    - [Msg](#nemo.community.v1beta1.Msg)
  
- [nemo/earn/v1beta1/strategy.proto](#nemo/earn/v1beta1/strategy.proto)
    - [StrategyType](#nemo.earn.v1beta1.StrategyType)
  
- [nemo/earn/v1beta1/vault.proto](#nemo/earn/v1beta1/vault.proto)
    - [AllowedVault](#nemo.earn.v1beta1.AllowedVault)
    - [VaultRecord](#nemo.earn.v1beta1.VaultRecord)
    - [VaultShare](#nemo.earn.v1beta1.VaultShare)
    - [VaultShareRecord](#nemo.earn.v1beta1.VaultShareRecord)
  
- [nemo/earn/v1beta1/params.proto](#nemo/earn/v1beta1/params.proto)
    - [Params](#nemo.earn.v1beta1.Params)
  
- [nemo/earn/v1beta1/genesis.proto](#nemo/earn/v1beta1/genesis.proto)
    - [GenesisState](#nemo.earn.v1beta1.GenesisState)
  
- [nemo/earn/v1beta1/proposal.proto](#nemo/earn/v1beta1/proposal.proto)
    - [CommunityPoolDepositProposal](#nemo.earn.v1beta1.CommunityPoolDepositProposal)
    - [CommunityPoolDepositProposalJSON](#nemo.earn.v1beta1.CommunityPoolDepositProposalJSON)
    - [CommunityPoolWithdrawProposal](#nemo.earn.v1beta1.CommunityPoolWithdrawProposal)
    - [CommunityPoolWithdrawProposalJSON](#nemo.earn.v1beta1.CommunityPoolWithdrawProposalJSON)
  
- [nemo/earn/v1beta1/query.proto](#nemo/earn/v1beta1/query.proto)
    - [DepositResponse](#nemo.earn.v1beta1.DepositResponse)
    - [QueryDepositsRequest](#nemo.earn.v1beta1.QueryDepositsRequest)
    - [QueryDepositsResponse](#nemo.earn.v1beta1.QueryDepositsResponse)
    - [QueryParamsRequest](#nemo.earn.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#nemo.earn.v1beta1.QueryParamsResponse)
    - [QueryTotalSupplyRequest](#nemo.earn.v1beta1.QueryTotalSupplyRequest)
    - [QueryTotalSupplyResponse](#nemo.earn.v1beta1.QueryTotalSupplyResponse)
    - [QueryVaultRequest](#nemo.earn.v1beta1.QueryVaultRequest)
    - [QueryVaultResponse](#nemo.earn.v1beta1.QueryVaultResponse)
    - [QueryVaultsRequest](#nemo.earn.v1beta1.QueryVaultsRequest)
    - [QueryVaultsResponse](#nemo.earn.v1beta1.QueryVaultsResponse)
    - [VaultResponse](#nemo.earn.v1beta1.VaultResponse)
  
    - [Query](#nemo.earn.v1beta1.Query)
  
- [nemo/earn/v1beta1/tx.proto](#nemo/earn/v1beta1/tx.proto)
    - [MsgDeposit](#nemo.earn.v1beta1.MsgDeposit)
    - [MsgDepositResponse](#nemo.earn.v1beta1.MsgDepositResponse)
    - [MsgWithdraw](#nemo.earn.v1beta1.MsgWithdraw)
    - [MsgWithdrawResponse](#nemo.earn.v1beta1.MsgWithdrawResponse)
  
    - [Msg](#nemo.earn.v1beta1.Msg)
  
- [nemo/evmutil/v1beta1/conversion_pair.proto](#nemo/evmutil/v1beta1/conversion_pair.proto)
    - [AllowedCosmosCoinERC20Token](#nemo.evmutil.v1beta1.AllowedCosmosCoinERC20Token)
    - [ConversionPair](#nemo.evmutil.v1beta1.ConversionPair)
  
- [nemo/evmutil/v1beta1/genesis.proto](#nemo/evmutil/v1beta1/genesis.proto)
    - [Account](#nemo.evmutil.v1beta1.Account)
    - [GenesisState](#nemo.evmutil.v1beta1.GenesisState)
    - [Params](#nemo.evmutil.v1beta1.Params)
  
- [nemo/evmutil/v1beta1/query.proto](#nemo/evmutil/v1beta1/query.proto)
    - [DeployedCosmosCoinContract](#nemo.evmutil.v1beta1.DeployedCosmosCoinContract)
    - [QueryDeployedCosmosCoinContractsRequest](#nemo.evmutil.v1beta1.QueryDeployedCosmosCoinContractsRequest)
    - [QueryDeployedCosmosCoinContractsResponse](#nemo.evmutil.v1beta1.QueryDeployedCosmosCoinContractsResponse)
    - [QueryParamsRequest](#nemo.evmutil.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#nemo.evmutil.v1beta1.QueryParamsResponse)
  
    - [Query](#nemo.evmutil.v1beta1.Query)
  
- [nemo/evmutil/v1beta1/tx.proto](#nemo/evmutil/v1beta1/tx.proto)
    - [MsgConvertCoinToERC20](#nemo.evmutil.v1beta1.MsgConvertCoinToERC20)
    - [MsgConvertCoinToERC20Response](#nemo.evmutil.v1beta1.MsgConvertCoinToERC20Response)
    - [MsgConvertCosmosCoinFromERC20](#nemo.evmutil.v1beta1.MsgConvertCosmosCoinFromERC20)
    - [MsgConvertCosmosCoinFromERC20Response](#nemo.evmutil.v1beta1.MsgConvertCosmosCoinFromERC20Response)
    - [MsgConvertCosmosCoinToERC20](#nemo.evmutil.v1beta1.MsgConvertCosmosCoinToERC20)
    - [MsgConvertCosmosCoinToERC20Response](#nemo.evmutil.v1beta1.MsgConvertCosmosCoinToERC20Response)
    - [MsgConvertERC20ToCoin](#nemo.evmutil.v1beta1.MsgConvertERC20ToCoin)
    - [MsgConvertERC20ToCoinResponse](#nemo.evmutil.v1beta1.MsgConvertERC20ToCoinResponse)
  
    - [Msg](#nemo.evmutil.v1beta1.Msg)
  
- [nemo/jinx/v1beta1/jinx.proto](#nemo/jinx/v1beta1/jinx.proto)
    - [Borrow](#nemo.jinx.v1beta1.Borrow)
    - [BorrowInterestFactor](#nemo.jinx.v1beta1.BorrowInterestFactor)
    - [BorrowLimit](#nemo.jinx.v1beta1.BorrowLimit)
    - [CoinsProto](#nemo.jinx.v1beta1.CoinsProto)
    - [Deposit](#nemo.jinx.v1beta1.Deposit)
    - [InterestRateModel](#nemo.jinx.v1beta1.InterestRateModel)
    - [MoneyMarket](#nemo.jinx.v1beta1.MoneyMarket)
    - [Params](#nemo.jinx.v1beta1.Params)
    - [SupplyInterestFactor](#nemo.jinx.v1beta1.SupplyInterestFactor)
  
- [nemo/jinx/v1beta1/genesis.proto](#nemo/jinx/v1beta1/genesis.proto)
    - [GenesisAccumulationTime](#nemo.jinx.v1beta1.GenesisAccumulationTime)
    - [GenesisState](#nemo.jinx.v1beta1.GenesisState)
  
- [nemo/jinx/v1beta1/query.proto](#nemo/jinx/v1beta1/query.proto)
    - [BorrowInterestFactorResponse](#nemo.jinx.v1beta1.BorrowInterestFactorResponse)
    - [BorrowResponse](#nemo.jinx.v1beta1.BorrowResponse)
    - [DepositResponse](#nemo.jinx.v1beta1.DepositResponse)
    - [InterestFactor](#nemo.jinx.v1beta1.InterestFactor)
    - [MoneyMarketInterestRate](#nemo.jinx.v1beta1.MoneyMarketInterestRate)
    - [QueryAccountsRequest](#nemo.jinx.v1beta1.QueryAccountsRequest)
    - [QueryAccountsResponse](#nemo.jinx.v1beta1.QueryAccountsResponse)
    - [QueryBorrowsRequest](#nemo.jinx.v1beta1.QueryBorrowsRequest)
    - [QueryBorrowsResponse](#nemo.jinx.v1beta1.QueryBorrowsResponse)
    - [QueryDepositsRequest](#nemo.jinx.v1beta1.QueryDepositsRequest)
    - [QueryDepositsResponse](#nemo.jinx.v1beta1.QueryDepositsResponse)
    - [QueryInterestFactorsRequest](#nemo.jinx.v1beta1.QueryInterestFactorsRequest)
    - [QueryInterestFactorsResponse](#nemo.jinx.v1beta1.QueryInterestFactorsResponse)
    - [QueryInterestRateRequest](#nemo.jinx.v1beta1.QueryInterestRateRequest)
    - [QueryInterestRateResponse](#nemo.jinx.v1beta1.QueryInterestRateResponse)
    - [QueryParamsRequest](#nemo.jinx.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#nemo.jinx.v1beta1.QueryParamsResponse)
    - [QueryReservesRequest](#nemo.jinx.v1beta1.QueryReservesRequest)
    - [QueryReservesResponse](#nemo.jinx.v1beta1.QueryReservesResponse)
    - [QueryTotalBorrowedRequest](#nemo.jinx.v1beta1.QueryTotalBorrowedRequest)
    - [QueryTotalBorrowedResponse](#nemo.jinx.v1beta1.QueryTotalBorrowedResponse)
    - [QueryTotalDepositedRequest](#nemo.jinx.v1beta1.QueryTotalDepositedRequest)
    - [QueryTotalDepositedResponse](#nemo.jinx.v1beta1.QueryTotalDepositedResponse)
    - [QueryUnsyncedBorrowsRequest](#nemo.jinx.v1beta1.QueryUnsyncedBorrowsRequest)
    - [QueryUnsyncedBorrowsResponse](#nemo.jinx.v1beta1.QueryUnsyncedBorrowsResponse)
    - [QueryUnsyncedDepositsRequest](#nemo.jinx.v1beta1.QueryUnsyncedDepositsRequest)
    - [QueryUnsyncedDepositsResponse](#nemo.jinx.v1beta1.QueryUnsyncedDepositsResponse)
    - [SupplyInterestFactorResponse](#nemo.jinx.v1beta1.SupplyInterestFactorResponse)
  
    - [Query](#nemo.jinx.v1beta1.Query)
  
- [nemo/jinx/v1beta1/tx.proto](#nemo/jinx/v1beta1/tx.proto)
    - [MsgBorrow](#nemo.jinx.v1beta1.MsgBorrow)
    - [MsgBorrowResponse](#nemo.jinx.v1beta1.MsgBorrowResponse)
    - [MsgDeposit](#nemo.jinx.v1beta1.MsgDeposit)
    - [MsgDepositResponse](#nemo.jinx.v1beta1.MsgDepositResponse)
    - [MsgLiquidate](#nemo.jinx.v1beta1.MsgLiquidate)
    - [MsgLiquidateResponse](#nemo.jinx.v1beta1.MsgLiquidateResponse)
    - [MsgRepay](#nemo.jinx.v1beta1.MsgRepay)
    - [MsgRepayResponse](#nemo.jinx.v1beta1.MsgRepayResponse)
    - [MsgWithdraw](#nemo.jinx.v1beta1.MsgWithdraw)
    - [MsgWithdrawResponse](#nemo.jinx.v1beta1.MsgWithdrawResponse)
  
    - [Msg](#nemo.jinx.v1beta1.Msg)
  
- [nemo/incentive/v1beta1/apy.proto](#nemo/incentive/v1beta1/apy.proto)
    - [Apy](#nemo.incentive.v1beta1.Apy)
  
- [nemo/incentive/v1beta1/claims.proto](#nemo/incentive/v1beta1/claims.proto)
    - [BaseClaim](#nemo.incentive.v1beta1.BaseClaim)
    - [BaseMultiClaim](#nemo.incentive.v1beta1.BaseMultiClaim)
    - [DelegatorClaim](#nemo.incentive.v1beta1.DelegatorClaim)
    - [EarnClaim](#nemo.incentive.v1beta1.EarnClaim)
    - [JinxLiquidityProviderClaim](#nemo.incentive.v1beta1.JinxLiquidityProviderClaim)
    - [MultiRewardIndex](#nemo.incentive.v1beta1.MultiRewardIndex)
    - [MultiRewardIndexesProto](#nemo.incentive.v1beta1.MultiRewardIndexesProto)
    - [RewardIndex](#nemo.incentive.v1beta1.RewardIndex)
    - [RewardIndexesProto](#nemo.incentive.v1beta1.RewardIndexesProto)
    - [SavingsClaim](#nemo.incentive.v1beta1.SavingsClaim)
    - [SwapClaim](#nemo.incentive.v1beta1.SwapClaim)
    - [MUSDMintingClaim](#nemo.incentive.v1beta1.MUSDMintingClaim)
  
- [nemo/incentive/v1beta1/params.proto](#nemo/incentive/v1beta1/params.proto)
    - [MultiRewardPeriod](#nemo.incentive.v1beta1.MultiRewardPeriod)
    - [Multiplier](#nemo.incentive.v1beta1.Multiplier)
    - [MultipliersPerDenom](#nemo.incentive.v1beta1.MultipliersPerDenom)
    - [Params](#nemo.incentive.v1beta1.Params)
    - [RewardPeriod](#nemo.incentive.v1beta1.RewardPeriod)
  
- [nemo/incentive/v1beta1/genesis.proto](#nemo/incentive/v1beta1/genesis.proto)
    - [AccumulationTime](#nemo.incentive.v1beta1.AccumulationTime)
    - [GenesisRewardState](#nemo.incentive.v1beta1.GenesisRewardState)
    - [GenesisState](#nemo.incentive.v1beta1.GenesisState)
  
- [nemo/incentive/v1beta1/query.proto](#nemo/incentive/v1beta1/query.proto)
    - [QueryApyRequest](#nemo.incentive.v1beta1.QueryApyRequest)
    - [QueryApyResponse](#nemo.incentive.v1beta1.QueryApyResponse)
    - [QueryParamsRequest](#nemo.incentive.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#nemo.incentive.v1beta1.QueryParamsResponse)
    - [QueryRewardFactorsRequest](#nemo.incentive.v1beta1.QueryRewardFactorsRequest)
    - [QueryRewardFactorsResponse](#nemo.incentive.v1beta1.QueryRewardFactorsResponse)
    - [QueryRewardsRequest](#nemo.incentive.v1beta1.QueryRewardsRequest)
    - [QueryRewardsResponse](#nemo.incentive.v1beta1.QueryRewardsResponse)
  
    - [Query](#nemo.incentive.v1beta1.Query)
  
- [nemo/incentive/v1beta1/tx.proto](#nemo/incentive/v1beta1/tx.proto)
    - [MsgClaimDelegatorReward](#nemo.incentive.v1beta1.MsgClaimDelegatorReward)
    - [MsgClaimDelegatorRewardResponse](#nemo.incentive.v1beta1.MsgClaimDelegatorRewardResponse)
    - [MsgClaimEarnReward](#nemo.incentive.v1beta1.MsgClaimEarnReward)
    - [MsgClaimEarnRewardResponse](#nemo.incentive.v1beta1.MsgClaimEarnRewardResponse)
    - [MsgClaimJinxReward](#nemo.incentive.v1beta1.MsgClaimJinxReward)
    - [MsgClaimJinxRewardResponse](#nemo.incentive.v1beta1.MsgClaimJinxRewardResponse)
    - [MsgClaimSavingsReward](#nemo.incentive.v1beta1.MsgClaimSavingsReward)
    - [MsgClaimSavingsRewardResponse](#nemo.incentive.v1beta1.MsgClaimSavingsRewardResponse)
    - [MsgClaimSwapReward](#nemo.incentive.v1beta1.MsgClaimSwapReward)
    - [MsgClaimSwapRewardResponse](#nemo.incentive.v1beta1.MsgClaimSwapRewardResponse)
    - [MsgClaimMUSDMintingReward](#nemo.incentive.v1beta1.MsgClaimMUSDMintingReward)
    - [MsgClaimMUSDMintingRewardResponse](#nemo.incentive.v1beta1.MsgClaimMUSDMintingRewardResponse)
    - [Selection](#nemo.incentive.v1beta1.Selection)
  
    - [Msg](#nemo.incentive.v1beta1.Msg)
  
- [nemo/issuance/v1beta1/genesis.proto](#nemo/issuance/v1beta1/genesis.proto)
    - [Asset](#nemo.issuance.v1beta1.Asset)
    - [AssetSupply](#nemo.issuance.v1beta1.AssetSupply)
    - [GenesisState](#nemo.issuance.v1beta1.GenesisState)
    - [Params](#nemo.issuance.v1beta1.Params)
    - [RateLimit](#nemo.issuance.v1beta1.RateLimit)
  
- [nemo/issuance/v1beta1/query.proto](#nemo/issuance/v1beta1/query.proto)
    - [QueryParamsRequest](#nemo.issuance.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#nemo.issuance.v1beta1.QueryParamsResponse)
  
    - [Query](#nemo.issuance.v1beta1.Query)
  
- [nemo/issuance/v1beta1/tx.proto](#nemo/issuance/v1beta1/tx.proto)
    - [MsgBlockAddress](#nemo.issuance.v1beta1.MsgBlockAddress)
    - [MsgBlockAddressResponse](#nemo.issuance.v1beta1.MsgBlockAddressResponse)
    - [MsgIssueTokens](#nemo.issuance.v1beta1.MsgIssueTokens)
    - [MsgIssueTokensResponse](#nemo.issuance.v1beta1.MsgIssueTokensResponse)
    - [MsgRedeemTokens](#nemo.issuance.v1beta1.MsgRedeemTokens)
    - [MsgRedeemTokensResponse](#nemo.issuance.v1beta1.MsgRedeemTokensResponse)
    - [MsgSetPauseStatus](#nemo.issuance.v1beta1.MsgSetPauseStatus)
    - [MsgSetPauseStatusResponse](#nemo.issuance.v1beta1.MsgSetPauseStatusResponse)
    - [MsgUnblockAddress](#nemo.issuance.v1beta1.MsgUnblockAddress)
    - [MsgUnblockAddressResponse](#nemo.issuance.v1beta1.MsgUnblockAddressResponse)
  
    - [Msg](#nemo.issuance.v1beta1.Msg)
  
- [nemo/liquid/v1beta1/query.proto](#nemo/liquid/v1beta1/query.proto)
    - [QueryDelegatedBalanceRequest](#nemo.liquid.v1beta1.QueryDelegatedBalanceRequest)
    - [QueryDelegatedBalanceResponse](#nemo.liquid.v1beta1.QueryDelegatedBalanceResponse)
    - [QueryTotalSupplyRequest](#nemo.liquid.v1beta1.QueryTotalSupplyRequest)
    - [QueryTotalSupplyResponse](#nemo.liquid.v1beta1.QueryTotalSupplyResponse)
  
    - [Query](#nemo.liquid.v1beta1.Query)
  
- [nemo/liquid/v1beta1/tx.proto](#nemo/liquid/v1beta1/tx.proto)
    - [MsgBurnDerivative](#nemo.liquid.v1beta1.MsgBurnDerivative)
    - [MsgBurnDerivativeResponse](#nemo.liquid.v1beta1.MsgBurnDerivativeResponse)
    - [MsgMintDerivative](#nemo.liquid.v1beta1.MsgMintDerivative)
    - [MsgMintDerivativeResponse](#nemo.liquid.v1beta1.MsgMintDerivativeResponse)
  
    - [Msg](#nemo.liquid.v1beta1.Msg)
  
- [nemo/nemodist/v1beta1/params.proto](#nemo/nemodist/v1beta1/params.proto)
    - [CoreReward](#nemo.nemodist.v1beta1.CoreReward)
    - [InfrastructureParams](#nemo.nemodist.v1beta1.InfrastructureParams)
    - [Params](#nemo.nemodist.v1beta1.Params)
    - [PartnerReward](#nemo.nemodist.v1beta1.PartnerReward)
    - [Period](#nemo.nemodist.v1beta1.Period)
  
- [nemo/nemodist/v1beta1/genesis.proto](#nemo/nemodist/v1beta1/genesis.proto)
    - [GenesisState](#nemo.nemodist.v1beta1.GenesisState)
  
- [nemo/nemodist/v1beta1/proposal.proto](#nemo/nemodist/v1beta1/proposal.proto)
    - [CommunityPoolMultiSpendProposal](#nemo.nemodist.v1beta1.CommunityPoolMultiSpendProposal)
    - [CommunityPoolMultiSpendProposalJSON](#nemo.nemodist.v1beta1.CommunityPoolMultiSpendProposalJSON)
    - [MultiSpendRecipient](#nemo.nemodist.v1beta1.MultiSpendRecipient)
  
- [nemo/nemodist/v1beta1/query.proto](#nemo/nemodist/v1beta1/query.proto)
    - [QueryBalanceRequest](#nemo.nemodist.v1beta1.QueryBalanceRequest)
    - [QueryBalanceResponse](#nemo.nemodist.v1beta1.QueryBalanceResponse)
    - [QueryParamsRequest](#nemo.nemodist.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#nemo.nemodist.v1beta1.QueryParamsResponse)
  
    - [Query](#nemo.nemodist.v1beta1.Query)
  
- [nemo/pricefeed/v1beta1/store.proto](#nemo/pricefeed/v1beta1/store.proto)
    - [CurrentPrice](#nemo.pricefeed.v1beta1.CurrentPrice)
    - [Market](#nemo.pricefeed.v1beta1.Market)
    - [Params](#nemo.pricefeed.v1beta1.Params)
    - [PostedPrice](#nemo.pricefeed.v1beta1.PostedPrice)
  
- [nemo/pricefeed/v1beta1/genesis.proto](#nemo/pricefeed/v1beta1/genesis.proto)
    - [GenesisState](#nemo.pricefeed.v1beta1.GenesisState)
  
- [nemo/pricefeed/v1beta1/query.proto](#nemo/pricefeed/v1beta1/query.proto)
    - [CurrentPriceResponse](#nemo.pricefeed.v1beta1.CurrentPriceResponse)
    - [MarketResponse](#nemo.pricefeed.v1beta1.MarketResponse)
    - [PostedPriceResponse](#nemo.pricefeed.v1beta1.PostedPriceResponse)
    - [QueryMarketsRequest](#nemo.pricefeed.v1beta1.QueryMarketsRequest)
    - [QueryMarketsResponse](#nemo.pricefeed.v1beta1.QueryMarketsResponse)
    - [QueryOraclesRequest](#nemo.pricefeed.v1beta1.QueryOraclesRequest)
    - [QueryOraclesResponse](#nemo.pricefeed.v1beta1.QueryOraclesResponse)
    - [QueryParamsRequest](#nemo.pricefeed.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#nemo.pricefeed.v1beta1.QueryParamsResponse)
    - [QueryPriceRequest](#nemo.pricefeed.v1beta1.QueryPriceRequest)
    - [QueryPriceResponse](#nemo.pricefeed.v1beta1.QueryPriceResponse)
    - [QueryPricesRequest](#nemo.pricefeed.v1beta1.QueryPricesRequest)
    - [QueryPricesResponse](#nemo.pricefeed.v1beta1.QueryPricesResponse)
    - [QueryRawPricesRequest](#nemo.pricefeed.v1beta1.QueryRawPricesRequest)
    - [QueryRawPricesResponse](#nemo.pricefeed.v1beta1.QueryRawPricesResponse)
  
    - [Query](#nemo.pricefeed.v1beta1.Query)
  
- [nemo/pricefeed/v1beta1/tx.proto](#nemo/pricefeed/v1beta1/tx.proto)
    - [MsgPostPrice](#nemo.pricefeed.v1beta1.MsgPostPrice)
    - [MsgPostPriceResponse](#nemo.pricefeed.v1beta1.MsgPostPriceResponse)
  
    - [Msg](#nemo.pricefeed.v1beta1.Msg)
  
- [nemo/router/v1beta1/tx.proto](#nemo/router/v1beta1/tx.proto)
    - [MsgDelegateMintDeposit](#nemo.router.v1beta1.MsgDelegateMintDeposit)
    - [MsgDelegateMintDepositResponse](#nemo.router.v1beta1.MsgDelegateMintDepositResponse)
    - [MsgMintDeposit](#nemo.router.v1beta1.MsgMintDeposit)
    - [MsgMintDepositResponse](#nemo.router.v1beta1.MsgMintDepositResponse)
    - [MsgWithdrawBurn](#nemo.router.v1beta1.MsgWithdrawBurn)
    - [MsgWithdrawBurnResponse](#nemo.router.v1beta1.MsgWithdrawBurnResponse)
    - [MsgWithdrawBurnUndelegate](#nemo.router.v1beta1.MsgWithdrawBurnUndelegate)
    - [MsgWithdrawBurnUndelegateResponse](#nemo.router.v1beta1.MsgWithdrawBurnUndelegateResponse)
  
    - [Msg](#nemo.router.v1beta1.Msg)
  
- [nemo/savings/v1beta1/store.proto](#nemo/savings/v1beta1/store.proto)
    - [Deposit](#nemo.savings.v1beta1.Deposit)
    - [Params](#nemo.savings.v1beta1.Params)
  
- [nemo/savings/v1beta1/genesis.proto](#nemo/savings/v1beta1/genesis.proto)
    - [GenesisState](#nemo.savings.v1beta1.GenesisState)
  
- [nemo/savings/v1beta1/query.proto](#nemo/savings/v1beta1/query.proto)
    - [QueryDepositsRequest](#nemo.savings.v1beta1.QueryDepositsRequest)
    - [QueryDepositsResponse](#nemo.savings.v1beta1.QueryDepositsResponse)
    - [QueryParamsRequest](#nemo.savings.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#nemo.savings.v1beta1.QueryParamsResponse)
    - [QueryTotalSupplyRequest](#nemo.savings.v1beta1.QueryTotalSupplyRequest)
    - [QueryTotalSupplyResponse](#nemo.savings.v1beta1.QueryTotalSupplyResponse)
  
    - [Query](#nemo.savings.v1beta1.Query)
  
- [nemo/savings/v1beta1/tx.proto](#nemo/savings/v1beta1/tx.proto)
    - [MsgDeposit](#nemo.savings.v1beta1.MsgDeposit)
    - [MsgDepositResponse](#nemo.savings.v1beta1.MsgDepositResponse)
    - [MsgWithdraw](#nemo.savings.v1beta1.MsgWithdraw)
    - [MsgWithdrawResponse](#nemo.savings.v1beta1.MsgWithdrawResponse)
  
    - [Msg](#nemo.savings.v1beta1.Msg)
  
- [nemo/swap/v1beta1/swap.proto](#nemo/swap/v1beta1/swap.proto)
    - [AllowedPool](#nemo.swap.v1beta1.AllowedPool)
    - [Params](#nemo.swap.v1beta1.Params)
    - [PoolRecord](#nemo.swap.v1beta1.PoolRecord)
    - [ShareRecord](#nemo.swap.v1beta1.ShareRecord)
  
- [nemo/swap/v1beta1/genesis.proto](#nemo/swap/v1beta1/genesis.proto)
    - [GenesisState](#nemo.swap.v1beta1.GenesisState)
  
- [nemo/swap/v1beta1/query.proto](#nemo/swap/v1beta1/query.proto)
    - [DepositResponse](#nemo.swap.v1beta1.DepositResponse)
    - [PoolResponse](#nemo.swap.v1beta1.PoolResponse)
    - [QueryDepositsRequest](#nemo.swap.v1beta1.QueryDepositsRequest)
    - [QueryDepositsResponse](#nemo.swap.v1beta1.QueryDepositsResponse)
    - [QueryParamsRequest](#nemo.swap.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#nemo.swap.v1beta1.QueryParamsResponse)
    - [QueryPoolsRequest](#nemo.swap.v1beta1.QueryPoolsRequest)
    - [QueryPoolsResponse](#nemo.swap.v1beta1.QueryPoolsResponse)
  
    - [Query](#nemo.swap.v1beta1.Query)
  
- [nemo/swap/v1beta1/tx.proto](#nemo/swap/v1beta1/tx.proto)
    - [MsgDeposit](#nemo.swap.v1beta1.MsgDeposit)
    - [MsgDepositResponse](#nemo.swap.v1beta1.MsgDepositResponse)
    - [MsgSwapExactForTokens](#nemo.swap.v1beta1.MsgSwapExactForTokens)
    - [MsgSwapExactForTokensResponse](#nemo.swap.v1beta1.MsgSwapExactForTokensResponse)
    - [MsgSwapForExactTokens](#nemo.swap.v1beta1.MsgSwapForExactTokens)
    - [MsgSwapForExactTokensResponse](#nemo.swap.v1beta1.MsgSwapForExactTokensResponse)
    - [MsgWithdraw](#nemo.swap.v1beta1.MsgWithdraw)
    - [MsgWithdrawResponse](#nemo.swap.v1beta1.MsgWithdrawResponse)
  
    - [Msg](#nemo.swap.v1beta1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="nemo/auction/v1beta1/auction.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/auction/v1beta1/auction.proto



<a name="nemo.auction.v1beta1.BaseAuction"></a>

### BaseAuction
BaseAuction defines common attributes of all auctions


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `initiator` | [string](#string) |  |  |
| `lot` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `bidder` | [bytes](#bytes) |  |  |
| `bid` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `has_received_bids` | [bool](#bool) |  |  |
| `end_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `max_end_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="nemo.auction.v1beta1.CollateralAuction"></a>

### CollateralAuction
CollateralAuction is a two phase auction.
Initially, in forward auction phase, bids can be placed up to a max bid.
Then it switches to a reverse auction phase, where the initial amount up for auction is bid down.
Unsold Lot is sent to LotReturns, being divided among the addresses by weight.
Collateral auctions are normally used to sell off collateral seized from CDPs.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_auction` | [BaseAuction](#nemo.auction.v1beta1.BaseAuction) |  |  |
| `corresponding_debt` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `max_bid` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `lot_returns` | [WeightedAddresses](#nemo.auction.v1beta1.WeightedAddresses) |  |  |






<a name="nemo.auction.v1beta1.DebtAuction"></a>

### DebtAuction
DebtAuction is a reverse auction that mints what it pays out.
It is normally used to acquire pegged asset to cover the CDP system's debts that were not covered by selling
collateral.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_auction` | [BaseAuction](#nemo.auction.v1beta1.BaseAuction) |  |  |
| `corresponding_debt` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="nemo.auction.v1beta1.SurplusAuction"></a>

### SurplusAuction
SurplusAuction is a forward auction that burns what it receives from bids.
It is normally used to sell off excess pegged asset acquired by the CDP system.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_auction` | [BaseAuction](#nemo.auction.v1beta1.BaseAuction) |  |  |






<a name="nemo.auction.v1beta1.WeightedAddresses"></a>

### WeightedAddresses
WeightedAddresses is a type for storing some addresses and associated weights.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `addresses` | [bytes](#bytes) | repeated |  |
| `weights` | [bytes](#bytes) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/auction/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/auction/v1beta1/genesis.proto



<a name="nemo.auction.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the auction module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `next_auction_id` | [uint64](#uint64) |  |  |
| `params` | [Params](#nemo.auction.v1beta1.Params) |  |  |
| `auctions` | [google.protobuf.Any](#google.protobuf.Any) | repeated | Genesis auctions |






<a name="nemo.auction.v1beta1.Params"></a>

### Params
Params defines the parameters for the issuance module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `max_auction_duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| `forward_bid_duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| `reverse_bid_duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| `increment_surplus` | [bytes](#bytes) |  |  |
| `increment_debt` | [bytes](#bytes) |  |  |
| `increment_collateral` | [bytes](#bytes) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/auction/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/auction/v1beta1/query.proto



<a name="nemo.auction.v1beta1.QueryAuctionRequest"></a>

### QueryAuctionRequest
QueryAuctionRequest is the request type for the Query/Auction RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `auction_id` | [uint64](#uint64) |  |  |






<a name="nemo.auction.v1beta1.QueryAuctionResponse"></a>

### QueryAuctionResponse
QueryAuctionResponse is the response type for the Query/Auction RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `auction` | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="nemo.auction.v1beta1.QueryAuctionsRequest"></a>

### QueryAuctionsRequest
QueryAuctionsRequest is the request type for the Query/Auctions RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `type` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `denom` | [string](#string) |  |  |
| `phase` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="nemo.auction.v1beta1.QueryAuctionsResponse"></a>

### QueryAuctionsResponse
QueryAuctionsResponse is the response type for the Query/Auctions RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `auctions` | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="nemo.auction.v1beta1.QueryNextAuctionIDRequest"></a>

### QueryNextAuctionIDRequest
QueryNextAuctionIDRequest defines the request type for querying x/auction next auction ID.






<a name="nemo.auction.v1beta1.QueryNextAuctionIDResponse"></a>

### QueryNextAuctionIDResponse
QueryNextAuctionIDResponse defines the response type for querying x/auction next auction ID.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |






<a name="nemo.auction.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest defines the request type for querying x/auction parameters.






<a name="nemo.auction.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse defines the response type for querying x/auction parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#nemo.auction.v1beta1.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.auction.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for auction module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#nemo.auction.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#nemo.auction.v1beta1.QueryParamsResponse) | Params queries all parameters of the auction module. | GET|/nemo/auction/v1beta1/params|
| `Auction` | [QueryAuctionRequest](#nemo.auction.v1beta1.QueryAuctionRequest) | [QueryAuctionResponse](#nemo.auction.v1beta1.QueryAuctionResponse) | Auction queries an individual Auction by auction ID | GET|/nemo/auction/v1beta1/auctions/{auction_id}|
| `Auctions` | [QueryAuctionsRequest](#nemo.auction.v1beta1.QueryAuctionsRequest) | [QueryAuctionsResponse](#nemo.auction.v1beta1.QueryAuctionsResponse) | Auctions queries auctions filtered by asset denom, owner address, phase, and auction type | GET|/nemo/auction/v1beta1/auctions|
| `NextAuctionID` | [QueryNextAuctionIDRequest](#nemo.auction.v1beta1.QueryNextAuctionIDRequest) | [QueryNextAuctionIDResponse](#nemo.auction.v1beta1.QueryNextAuctionIDResponse) | NextAuctionID queries the next auction ID | GET|/nemo/auction/v1beta1/next-auction-id|

 <!-- end services -->



<a name="nemo/auction/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/auction/v1beta1/tx.proto



<a name="nemo.auction.v1beta1.MsgPlaceBid"></a>

### MsgPlaceBid
MsgPlaceBid represents a message used by bidders to place bids on auctions


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `auction_id` | [uint64](#uint64) |  |  |
| `bidder` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="nemo.auction.v1beta1.MsgPlaceBidResponse"></a>

### MsgPlaceBidResponse
MsgPlaceBidResponse defines the Msg/PlaceBid response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.auction.v1beta1.Msg"></a>

### Msg
Msg defines the auction Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `PlaceBid` | [MsgPlaceBid](#nemo.auction.v1beta1.MsgPlaceBid) | [MsgPlaceBidResponse](#nemo.auction.v1beta1.MsgPlaceBidResponse) | PlaceBid message type used by bidders to place bids on auctions | |

 <!-- end services -->



<a name="nemo/bep3/v1beta1/bep3.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/bep3/v1beta1/bep3.proto



<a name="nemo.bep3.v1beta1.AssetParam"></a>

### AssetParam
AssetParam defines parameters for each bep3 asset.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  | denom represents the denominatin for this asset |
| `coin_id` | [int64](#int64) |  | coin_id represents the registered coin type to use (https://github.com/satoshilabs/slips/blob/master/slip-0044.md) |
| `supply_limit` | [SupplyLimit](#nemo.bep3.v1beta1.SupplyLimit) |  | supply_limit defines the maximum supply allowed for the asset - a total or time based rate limit |
| `active` | [bool](#bool) |  | active specifies if the asset is live or paused |
| `deputy_address` | [bytes](#bytes) |  | deputy_address the nemo address of the deputy |
| `fixed_fee` | [string](#string) |  | fixed_fee defines the fee for incoming swaps |
| `min_swap_amount` | [string](#string) |  | min_swap_amount defines the minimum amount able to be swapped in a single message |
| `max_swap_amount` | [string](#string) |  | max_swap_amount defines the maximum amount able to be swapped in a single message |
| `min_block_lock` | [uint64](#uint64) |  | min_block_lock defined the minimum blocks to lock |
| `max_block_lock` | [uint64](#uint64) |  | min_block_lock defined the maximum blocks to lock |






<a name="nemo.bep3.v1beta1.AssetSupply"></a>

### AssetSupply
AssetSupply defines information about an asset's supply.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `incoming_supply` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | incoming_supply represents the incoming supply of an asset |
| `outgoing_supply` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | outgoing_supply represents the outgoing supply of an asset |
| `current_supply` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | current_supply represents the current on-chain supply of an asset |
| `time_limited_current_supply` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | time_limited_current_supply represents the time limited current supply of an asset |
| `time_elapsed` | [google.protobuf.Duration](#google.protobuf.Duration) |  | time_elapsed represents the time elapsed |






<a name="nemo.bep3.v1beta1.AtomicSwap"></a>

### AtomicSwap
AtomicSwap defines an atomic swap between chains for the pricefeed module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | amount represents the amount being swapped |
| `random_number_hash` | [bytes](#bytes) |  | random_number_hash represents the hash of the random number |
| `expire_height` | [uint64](#uint64) |  | expire_height represents the height when the swap expires |
| `timestamp` | [int64](#int64) |  | timestamp represents the timestamp of the swap |
| `sender` | [bytes](#bytes) |  | sender is the nemo chain sender of the swap |
| `recipient` | [bytes](#bytes) |  | recipient is the nemo chain recipient of the swap |
| `sender_other_chain` | [string](#string) |  | sender_other_chain is the sender on the other chain |
| `recipient_other_chain` | [string](#string) |  | recipient_other_chain is the recipient on the other chain |
| `closed_block` | [int64](#int64) |  | closed_block is the block when the swap is closed |
| `status` | [SwapStatus](#nemo.bep3.v1beta1.SwapStatus) |  | status represents the current status of the swap |
| `cross_chain` | [bool](#bool) |  | cross_chain identifies whether the atomic swap is cross chain |
| `direction` | [SwapDirection](#nemo.bep3.v1beta1.SwapDirection) |  | direction identifies if the swap is incoming or outgoing |






<a name="nemo.bep3.v1beta1.Params"></a>

### Params
Params defines the parameters for the bep3 module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `asset_params` | [AssetParam](#nemo.bep3.v1beta1.AssetParam) | repeated | asset_params define the parameters for each bep3 asset |






<a name="nemo.bep3.v1beta1.SupplyLimit"></a>

### SupplyLimit
SupplyLimit define the absolute and time-based limits for an assets's supply.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `limit` | [string](#string) |  | limit defines the total supply allowed |
| `time_limited` | [bool](#bool) |  | time_limited enables or disables time based supply limiting |
| `time_period` | [google.protobuf.Duration](#google.protobuf.Duration) |  | time_period specifies the duration that time_based_limit is evalulated |
| `time_based_limit` | [string](#string) |  | time_based_limit defines the maximum supply that can be swapped within time_period |





 <!-- end messages -->


<a name="nemo.bep3.v1beta1.SwapDirection"></a>

### SwapDirection
SwapDirection is the direction of an AtomicSwap

| Name | Number | Description |
| ---- | ------ | ----------- |
| SWAP_DIRECTION_UNSPECIFIED | 0 | SWAP_DIRECTION_UNSPECIFIED represents unspecified or invalid swap direcation |
| SWAP_DIRECTION_INCOMING | 1 | SWAP_DIRECTION_INCOMING represents is incoming swap (to the nemo chain) |
| SWAP_DIRECTION_OUTGOING | 2 | SWAP_DIRECTION_OUTGOING represents an outgoing swap (from the nemo chain) |



<a name="nemo.bep3.v1beta1.SwapStatus"></a>

### SwapStatus
SwapStatus is the status of an AtomicSwap

| Name | Number | Description |
| ---- | ------ | ----------- |
| SWAP_STATUS_UNSPECIFIED | 0 | SWAP_STATUS_UNSPECIFIED represents an unspecified status |
| SWAP_STATUS_OPEN | 1 | SWAP_STATUS_OPEN represents an open swap |
| SWAP_STATUS_COMPLETED | 2 | SWAP_STATUS_COMPLETED represents a completed swap |
| SWAP_STATUS_EXPIRED | 3 | SWAP_STATUS_EXPIRED represents an expired swap |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/bep3/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/bep3/v1beta1/genesis.proto



<a name="nemo.bep3.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the pricefeed module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#nemo.bep3.v1beta1.Params) |  | params defines all the paramaters of the module. |
| `atomic_swaps` | [AtomicSwap](#nemo.bep3.v1beta1.AtomicSwap) | repeated | atomic_swaps represents the state of stored atomic swaps |
| `supplies` | [AssetSupply](#nemo.bep3.v1beta1.AssetSupply) | repeated | supplies represents the supply information of each atomic swap |
| `previous_block_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | previous_block_time represents the time of the previous block |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/bep3/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/bep3/v1beta1/query.proto



<a name="nemo.bep3.v1beta1.AssetSupplyResponse"></a>

### AssetSupplyResponse
AssetSupplyResponse defines information about an asset's supply.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `incoming_supply` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | incoming_supply represents the incoming supply of an asset |
| `outgoing_supply` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | outgoing_supply represents the outgoing supply of an asset |
| `current_supply` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | current_supply represents the current on-chain supply of an asset |
| `time_limited_current_supply` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | time_limited_current_supply represents the time limited current supply of an asset |
| `time_elapsed` | [google.protobuf.Duration](#google.protobuf.Duration) |  | time_elapsed represents the time elapsed |






<a name="nemo.bep3.v1beta1.AtomicSwapResponse"></a>

### AtomicSwapResponse
AtomicSwapResponse represents the returned atomic swap properties


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  | id represents the id of the atomic swap |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | amount represents the amount being swapped |
| `random_number_hash` | [string](#string) |  | random_number_hash represents the hash of the random number |
| `expire_height` | [uint64](#uint64) |  | expire_height represents the height when the swap expires |
| `timestamp` | [int64](#int64) |  | timestamp represents the timestamp of the swap |
| `sender` | [string](#string) |  | sender is the nemo chain sender of the swap |
| `recipient` | [string](#string) |  | recipient is the nemo chain recipient of the swap |
| `sender_other_chain` | [string](#string) |  | sender_other_chain is the sender on the other chain |
| `recipient_other_chain` | [string](#string) |  | recipient_other_chain is the recipient on the other chain |
| `closed_block` | [int64](#int64) |  | closed_block is the block when the swap is closed |
| `status` | [SwapStatus](#nemo.bep3.v1beta1.SwapStatus) |  | status represents the current status of the swap |
| `cross_chain` | [bool](#bool) |  | cross_chain identifies whether the atomic swap is cross chain |
| `direction` | [SwapDirection](#nemo.bep3.v1beta1.SwapDirection) |  | direction identifies if the swap is incoming or outgoing |






<a name="nemo.bep3.v1beta1.QueryAssetSuppliesRequest"></a>

### QueryAssetSuppliesRequest
QueryAssetSuppliesRequest is the request type for the Query/AssetSupplies RPC method.






<a name="nemo.bep3.v1beta1.QueryAssetSuppliesResponse"></a>

### QueryAssetSuppliesResponse
QueryAssetSuppliesResponse is the response type for the Query/AssetSupplies RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `asset_supplies` | [AssetSupplyResponse](#nemo.bep3.v1beta1.AssetSupplyResponse) | repeated | asset_supplies represents the supplies of returned assets |






<a name="nemo.bep3.v1beta1.QueryAssetSupplyRequest"></a>

### QueryAssetSupplyRequest
QueryAssetSupplyRequest is the request type for the Query/AssetSupply RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  | denom filters the asset response for the specified denom |






<a name="nemo.bep3.v1beta1.QueryAssetSupplyResponse"></a>

### QueryAssetSupplyResponse
QueryAssetSupplyResponse is the response type for the Query/AssetSupply RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `asset_supply` | [AssetSupplyResponse](#nemo.bep3.v1beta1.AssetSupplyResponse) |  | asset_supply represents the supply of the asset |






<a name="nemo.bep3.v1beta1.QueryAtomicSwapRequest"></a>

### QueryAtomicSwapRequest
QueryAtomicSwapRequest is the request type for the Query/AtomicSwap RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `swap_id` | [string](#string) |  | swap_id represents the id of the swap to query |






<a name="nemo.bep3.v1beta1.QueryAtomicSwapResponse"></a>

### QueryAtomicSwapResponse
QueryAtomicSwapResponse is the response type for the Query/AtomicSwap RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `atomic_swap` | [AtomicSwapResponse](#nemo.bep3.v1beta1.AtomicSwapResponse) |  |  |






<a name="nemo.bep3.v1beta1.QueryAtomicSwapsRequest"></a>

### QueryAtomicSwapsRequest
QueryAtomicSwapsRequest is the request type for the Query/AtomicSwaps RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `involve` | [string](#string) |  | involve filters by address |
| `expiration` | [uint64](#uint64) |  | expiration filters by expiration block height |
| `status` | [SwapStatus](#nemo.bep3.v1beta1.SwapStatus) |  | status filters by swap status |
| `direction` | [SwapDirection](#nemo.bep3.v1beta1.SwapDirection) |  | direction fitlers by swap direction |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="nemo.bep3.v1beta1.QueryAtomicSwapsResponse"></a>

### QueryAtomicSwapsResponse
QueryAtomicSwapsResponse is the response type for the Query/AtomicSwaps RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `atomic_swaps` | [AtomicSwapResponse](#nemo.bep3.v1beta1.AtomicSwapResponse) | repeated | atomic_swap represents the returned atomic swaps for the request |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="nemo.bep3.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest defines the request type for querying x/bep3 parameters.






<a name="nemo.bep3.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse defines the response type for querying x/bep3 parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#nemo.bep3.v1beta1.Params) |  | params represents the parameters of the module |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.bep3.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for bep3 module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#nemo.bep3.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#nemo.bep3.v1beta1.QueryParamsResponse) | Params queries module params | GET|/nemo/bep3/v1beta1/params|
| `AssetSupply` | [QueryAssetSupplyRequest](#nemo.bep3.v1beta1.QueryAssetSupplyRequest) | [QueryAssetSupplyResponse](#nemo.bep3.v1beta1.QueryAssetSupplyResponse) | AssetSupply queries info about an asset's supply | GET|/nemo/bep3/v1beta1/assetsupply/{denom}|
| `AssetSupplies` | [QueryAssetSuppliesRequest](#nemo.bep3.v1beta1.QueryAssetSuppliesRequest) | [QueryAssetSuppliesResponse](#nemo.bep3.v1beta1.QueryAssetSuppliesResponse) | AssetSupplies queries a list of asset supplies | GET|/nemo/bep3/v1beta1/assetsupplies|
| `AtomicSwap` | [QueryAtomicSwapRequest](#nemo.bep3.v1beta1.QueryAtomicSwapRequest) | [QueryAtomicSwapResponse](#nemo.bep3.v1beta1.QueryAtomicSwapResponse) | AtomicSwap queries info about an atomic swap | GET|/nemo/bep3/v1beta1/atomicswap/{swap_id}|
| `AtomicSwaps` | [QueryAtomicSwapsRequest](#nemo.bep3.v1beta1.QueryAtomicSwapsRequest) | [QueryAtomicSwapsResponse](#nemo.bep3.v1beta1.QueryAtomicSwapsResponse) | AtomicSwaps queries a list of atomic swaps | GET|/nemo/bep3/v1beta1/atomicswaps|

 <!-- end services -->



<a name="nemo/bep3/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/bep3/v1beta1/tx.proto



<a name="nemo.bep3.v1beta1.MsgClaimAtomicSwap"></a>

### MsgClaimAtomicSwap
MsgClaimAtomicSwap defines the Msg/ClaimAtomicSwap request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from` | [string](#string) |  |  |
| `swap_id` | [string](#string) |  |  |
| `random_number` | [string](#string) |  |  |






<a name="nemo.bep3.v1beta1.MsgClaimAtomicSwapResponse"></a>

### MsgClaimAtomicSwapResponse
MsgClaimAtomicSwapResponse defines the Msg/ClaimAtomicSwap response type.






<a name="nemo.bep3.v1beta1.MsgCreateAtomicSwap"></a>

### MsgCreateAtomicSwap
MsgCreateAtomicSwap defines the Msg/CreateAtomicSwap request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from` | [string](#string) |  |  |
| `to` | [string](#string) |  |  |
| `recipient_other_chain` | [string](#string) |  |  |
| `sender_other_chain` | [string](#string) |  |  |
| `random_number_hash` | [string](#string) |  |  |
| `timestamp` | [int64](#int64) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `height_span` | [uint64](#uint64) |  |  |






<a name="nemo.bep3.v1beta1.MsgCreateAtomicSwapResponse"></a>

### MsgCreateAtomicSwapResponse
MsgCreateAtomicSwapResponse defines the Msg/CreateAtomicSwap response type.






<a name="nemo.bep3.v1beta1.MsgRefundAtomicSwap"></a>

### MsgRefundAtomicSwap
MsgRefundAtomicSwap defines the Msg/RefundAtomicSwap request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from` | [string](#string) |  |  |
| `swap_id` | [string](#string) |  |  |






<a name="nemo.bep3.v1beta1.MsgRefundAtomicSwapResponse"></a>

### MsgRefundAtomicSwapResponse
MsgRefundAtomicSwapResponse defines the Msg/RefundAtomicSwap response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.bep3.v1beta1.Msg"></a>

### Msg
Msg defines the bep3 Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreateAtomicSwap` | [MsgCreateAtomicSwap](#nemo.bep3.v1beta1.MsgCreateAtomicSwap) | [MsgCreateAtomicSwapResponse](#nemo.bep3.v1beta1.MsgCreateAtomicSwapResponse) | CreateAtomicSwap defines a method for creating an atomic swap | |
| `ClaimAtomicSwap` | [MsgClaimAtomicSwap](#nemo.bep3.v1beta1.MsgClaimAtomicSwap) | [MsgClaimAtomicSwapResponse](#nemo.bep3.v1beta1.MsgClaimAtomicSwapResponse) | ClaimAtomicSwap defines a method for claiming an atomic swap | |
| `RefundAtomicSwap` | [MsgRefundAtomicSwap](#nemo.bep3.v1beta1.MsgRefundAtomicSwap) | [MsgRefundAtomicSwapResponse](#nemo.bep3.v1beta1.MsgRefundAtomicSwapResponse) | RefundAtomicSwap defines a method for refunding an atomic swap | |

 <!-- end services -->



<a name="nemo/cdp/v1beta1/cdp.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/cdp/v1beta1/cdp.proto



<a name="nemo.cdp.v1beta1.CDP"></a>

### CDP
CDP defines the state of a single collateralized debt position.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `owner` | [bytes](#bytes) |  |  |
| `type` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `principal` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `accumulated_fees` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `fees_updated` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `interest_factor` | [string](#string) |  |  |






<a name="nemo.cdp.v1beta1.Deposit"></a>

### Deposit
Deposit defines an amount of coins deposited by an account to a cdp


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp_id` | [uint64](#uint64) |  |  |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="nemo.cdp.v1beta1.OwnerCDPIndex"></a>

### OwnerCDPIndex
OwnerCDPIndex defines the cdp ids for a single cdp owner


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp_ids` | [uint64](#uint64) | repeated |  |






<a name="nemo.cdp.v1beta1.TotalCollateral"></a>

### TotalCollateral
TotalCollateral defines the total collateral of a given collateral type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="nemo.cdp.v1beta1.TotalPrincipal"></a>

### TotalPrincipal
TotalPrincipal defines the total principal of a given collateral type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/cdp/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/cdp/v1beta1/genesis.proto



<a name="nemo.cdp.v1beta1.CollateralParam"></a>

### CollateralParam
CollateralParam defines governance parameters for each collateral type within the cdp module


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `type` | [string](#string) |  |  |
| `liquidation_ratio` | [string](#string) |  |  |
| `debt_limit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `stability_fee` | [string](#string) |  |  |
| `auction_size` | [string](#string) |  |  |
| `liquidation_penalty` | [string](#string) |  |  |
| `spot_market_id` | [string](#string) |  |  |
| `liquidation_market_id` | [string](#string) |  |  |
| `keeper_reward_percentage` | [string](#string) |  |  |
| `check_collateralization_index_count` | [string](#string) |  |  |
| `conversion_factor` | [string](#string) |  |  |






<a name="nemo.cdp.v1beta1.DebtParam"></a>

### DebtParam
DebtParam defines governance params for debt assets


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `reference_asset` | [string](#string) |  |  |
| `conversion_factor` | [string](#string) |  |  |
| `debt_floor` | [string](#string) |  |  |






<a name="nemo.cdp.v1beta1.GenesisAccumulationTime"></a>

### GenesisAccumulationTime
GenesisAccumulationTime defines the previous distribution time and its corresponding denom


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `previous_accumulation_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `interest_factor` | [string](#string) |  |  |






<a name="nemo.cdp.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the cdp module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#nemo.cdp.v1beta1.Params) |  | params defines all the paramaters of the module. |
| `cdps` | [CDP](#nemo.cdp.v1beta1.CDP) | repeated |  |
| `deposits` | [Deposit](#nemo.cdp.v1beta1.Deposit) | repeated |  |
| `starting_cdp_id` | [uint64](#uint64) |  |  |
| `debt_denom` | [string](#string) |  |  |
| `gov_denom` | [string](#string) |  |  |
| `previous_accumulation_times` | [GenesisAccumulationTime](#nemo.cdp.v1beta1.GenesisAccumulationTime) | repeated |  |
| `total_principals` | [GenesisTotalPrincipal](#nemo.cdp.v1beta1.GenesisTotalPrincipal) | repeated |  |






<a name="nemo.cdp.v1beta1.GenesisTotalPrincipal"></a>

### GenesisTotalPrincipal
GenesisTotalPrincipal defines the total principal and its corresponding collateral type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `total_principal` | [string](#string) |  |  |






<a name="nemo.cdp.v1beta1.Params"></a>

### Params
Params defines the parameters for the cdp module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_params` | [CollateralParam](#nemo.cdp.v1beta1.CollateralParam) | repeated |  |
| `debt_param` | [DebtParam](#nemo.cdp.v1beta1.DebtParam) |  |  |
| `global_debt_limit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `surplus_auction_threshold` | [string](#string) |  |  |
| `surplus_auction_lot` | [string](#string) |  |  |
| `debt_auction_threshold` | [string](#string) |  |  |
| `debt_auction_lot` | [string](#string) |  |  |
| `circuit_breaker` | [bool](#bool) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/cdp/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/cdp/v1beta1/query.proto



<a name="nemo.cdp.v1beta1.CDPResponse"></a>

### CDPResponse
CDPResponse defines the state of a single collateralized debt position.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `owner` | [string](#string) |  |  |
| `type` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `principal` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `accumulated_fees` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `fees_updated` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `interest_factor` | [string](#string) |  |  |
| `collateral_value` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateralization_ratio` | [string](#string) |  |  |






<a name="nemo.cdp.v1beta1.QueryAccountsRequest"></a>

### QueryAccountsRequest
QueryAccountsRequest defines the request type for the Query/Accounts RPC method.






<a name="nemo.cdp.v1beta1.QueryAccountsResponse"></a>

### QueryAccountsResponse
QueryAccountsResponse defines the response type for the Query/Accounts RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `accounts` | [cosmos.auth.v1beta1.ModuleAccount](#cosmos.auth.v1beta1.ModuleAccount) | repeated |  |






<a name="nemo.cdp.v1beta1.QueryCdpRequest"></a>

### QueryCdpRequest
QueryCdpRequest defines the request type for the Query/Cdp RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |






<a name="nemo.cdp.v1beta1.QueryCdpResponse"></a>

### QueryCdpResponse
QueryCdpResponse defines the response type for the Query/Cdp RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp` | [CDPResponse](#nemo.cdp.v1beta1.CDPResponse) |  |  |






<a name="nemo.cdp.v1beta1.QueryCdpsRequest"></a>

### QueryCdpsRequest
QueryCdpsRequest is the params for a filtered CDP query, the request type for the Query/Cdps RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `id` | [uint64](#uint64) |  |  |
| `ratio` | [string](#string) |  | sdk.Dec as a string |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="nemo.cdp.v1beta1.QueryCdpsResponse"></a>

### QueryCdpsResponse
QueryCdpsResponse defines the response type for the Query/Cdps RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdps` | [CDPResponse](#nemo.cdp.v1beta1.CDPResponse) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="nemo.cdp.v1beta1.QueryDepositsRequest"></a>

### QueryDepositsRequest
QueryDepositsRequest defines the request type for the Query/Deposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |






<a name="nemo.cdp.v1beta1.QueryDepositsResponse"></a>

### QueryDepositsResponse
QueryDepositsResponse defines the response type for the Query/Deposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposits` | [Deposit](#nemo.cdp.v1beta1.Deposit) | repeated |  |






<a name="nemo.cdp.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest defines the request type for the Query/Params RPC method.






<a name="nemo.cdp.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse defines the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#nemo.cdp.v1beta1.Params) |  |  |






<a name="nemo.cdp.v1beta1.QueryTotalCollateralRequest"></a>

### QueryTotalCollateralRequest
QueryTotalCollateralRequest defines the request type for the Query/TotalCollateral RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |






<a name="nemo.cdp.v1beta1.QueryTotalCollateralResponse"></a>

### QueryTotalCollateralResponse
QueryTotalCollateralResponse defines the response type for the Query/TotalCollateral RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `total_collateral` | [TotalCollateral](#nemo.cdp.v1beta1.TotalCollateral) | repeated |  |






<a name="nemo.cdp.v1beta1.QueryTotalPrincipalRequest"></a>

### QueryTotalPrincipalRequest
QueryTotalPrincipalRequest defines the request type for the Query/TotalPrincipal RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |






<a name="nemo.cdp.v1beta1.QueryTotalPrincipalResponse"></a>

### QueryTotalPrincipalResponse
QueryTotalPrincipalResponse defines the response type for the Query/TotalPrincipal RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `total_principal` | [TotalPrincipal](#nemo.cdp.v1beta1.TotalPrincipal) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.cdp.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for cdp module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#nemo.cdp.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#nemo.cdp.v1beta1.QueryParamsResponse) | Params queries all parameters of the cdp module. | GET|/nemo/cdp/v1beta1/params|
| `Accounts` | [QueryAccountsRequest](#nemo.cdp.v1beta1.QueryAccountsRequest) | [QueryAccountsResponse](#nemo.cdp.v1beta1.QueryAccountsResponse) | Accounts queries the CDP module accounts. | GET|/nemo/cdp/v1beta1/accounts|
| `TotalPrincipal` | [QueryTotalPrincipalRequest](#nemo.cdp.v1beta1.QueryTotalPrincipalRequest) | [QueryTotalPrincipalResponse](#nemo.cdp.v1beta1.QueryTotalPrincipalResponse) | TotalPrincipal queries the total principal of a given collateral type. | GET|/nemo/cdp/v1beta1/totalPrincipal|
| `TotalCollateral` | [QueryTotalCollateralRequest](#nemo.cdp.v1beta1.QueryTotalCollateralRequest) | [QueryTotalCollateralResponse](#nemo.cdp.v1beta1.QueryTotalCollateralResponse) | TotalCollateral queries the total collateral of a given collateral type. | GET|/nemo/cdp/v1beta1/totalCollateral|
| `Cdps` | [QueryCdpsRequest](#nemo.cdp.v1beta1.QueryCdpsRequest) | [QueryCdpsResponse](#nemo.cdp.v1beta1.QueryCdpsResponse) | Cdps queries all active CDPs. | GET|/nemo/cdp/v1beta1/cdps|
| `Cdp` | [QueryCdpRequest](#nemo.cdp.v1beta1.QueryCdpRequest) | [QueryCdpResponse](#nemo.cdp.v1beta1.QueryCdpResponse) | Cdp queries a CDP with the input owner address and collateral type. | GET|/nemo/cdp/v1beta1/cdps/{owner}/{collateral_type}|
| `Deposits` | [QueryDepositsRequest](#nemo.cdp.v1beta1.QueryDepositsRequest) | [QueryDepositsResponse](#nemo.cdp.v1beta1.QueryDepositsResponse) | Deposits queries deposits associated with the CDP owned by an address for a collateral type. | GET|/nemo/cdp/v1beta1/cdps/deposits/{owner}/{collateral_type}|

 <!-- end services -->



<a name="nemo/cdp/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/cdp/v1beta1/tx.proto



<a name="nemo.cdp.v1beta1.MsgCreateCDP"></a>

### MsgCreateCDP
MsgCreateCDP defines a message to create a new CDP.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `principal` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="nemo.cdp.v1beta1.MsgCreateCDPResponse"></a>

### MsgCreateCDPResponse
MsgCreateCDPResponse defines the Msg/CreateCDP response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cdp_id` | [uint64](#uint64) |  |  |






<a name="nemo.cdp.v1beta1.MsgDeposit"></a>

### MsgDeposit
MsgDeposit defines a message to deposit to a CDP.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="nemo.cdp.v1beta1.MsgDepositResponse"></a>

### MsgDepositResponse
MsgDepositResponse defines the Msg/Deposit response type.






<a name="nemo.cdp.v1beta1.MsgDrawDebt"></a>

### MsgDrawDebt
MsgDrawDebt defines a message to draw debt from a CDP.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |
| `principal` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="nemo.cdp.v1beta1.MsgDrawDebtResponse"></a>

### MsgDrawDebtResponse
MsgDrawDebtResponse defines the Msg/DrawDebt response type.






<a name="nemo.cdp.v1beta1.MsgLiquidate"></a>

### MsgLiquidate
MsgLiquidate defines a message to attempt to liquidate a CDP whos
collateralization ratio is under its liquidation ratio.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `keeper` | [string](#string) |  |  |
| `borrower` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="nemo.cdp.v1beta1.MsgLiquidateResponse"></a>

### MsgLiquidateResponse
MsgLiquidateResponse defines the Msg/Liquidate response type.






<a name="nemo.cdp.v1beta1.MsgRepayDebt"></a>

### MsgRepayDebt
MsgRepayDebt defines a message to repay debt from a CDP.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |
| `payment` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="nemo.cdp.v1beta1.MsgRepayDebtResponse"></a>

### MsgRepayDebtResponse
MsgRepayDebtResponse defines the Msg/RepayDebt response type.






<a name="nemo.cdp.v1beta1.MsgWithdraw"></a>

### MsgWithdraw
MsgWithdraw defines a message to withdraw collateral from a CDP.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `collateral_type` | [string](#string) |  |  |






<a name="nemo.cdp.v1beta1.MsgWithdrawResponse"></a>

### MsgWithdrawResponse
MsgWithdrawResponse defines the Msg/Withdraw response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.cdp.v1beta1.Msg"></a>

### Msg
Msg defines the cdp Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreateCDP` | [MsgCreateCDP](#nemo.cdp.v1beta1.MsgCreateCDP) | [MsgCreateCDPResponse](#nemo.cdp.v1beta1.MsgCreateCDPResponse) | CreateCDP defines a method to create a new CDP. | |
| `Deposit` | [MsgDeposit](#nemo.cdp.v1beta1.MsgDeposit) | [MsgDepositResponse](#nemo.cdp.v1beta1.MsgDepositResponse) | Deposit defines a method to deposit to a CDP. | |
| `Withdraw` | [MsgWithdraw](#nemo.cdp.v1beta1.MsgWithdraw) | [MsgWithdrawResponse](#nemo.cdp.v1beta1.MsgWithdrawResponse) | Withdraw defines a method to withdraw collateral from a CDP. | |
| `DrawDebt` | [MsgDrawDebt](#nemo.cdp.v1beta1.MsgDrawDebt) | [MsgDrawDebtResponse](#nemo.cdp.v1beta1.MsgDrawDebtResponse) | DrawDebt defines a method to draw debt from a CDP. | |
| `RepayDebt` | [MsgRepayDebt](#nemo.cdp.v1beta1.MsgRepayDebt) | [MsgRepayDebtResponse](#nemo.cdp.v1beta1.MsgRepayDebtResponse) | RepayDebt defines a method to repay debt from a CDP. | |
| `Liquidate` | [MsgLiquidate](#nemo.cdp.v1beta1.MsgLiquidate) | [MsgLiquidateResponse](#nemo.cdp.v1beta1.MsgLiquidateResponse) | Liquidate defines a method to attempt to liquidate a CDP whos collateralization ratio is under its liquidation ratio. | |

 <!-- end services -->



<a name="nemo/committee/v1beta1/committee.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/committee/v1beta1/committee.proto



<a name="nemo.committee.v1beta1.BaseCommittee"></a>

### BaseCommittee
BaseCommittee is a common type shared by all Committees


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `description` | [string](#string) |  |  |
| `members` | [bytes](#bytes) | repeated |  |
| `permissions` | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |
| `vote_threshold` | [string](#string) |  | Smallest percentage that must vote for a proposal to pass |
| `proposal_duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  | The length of time a proposal remains active for. Proposals will close earlier if they get enough votes. |
| `tally_option` | [TallyOption](#nemo.committee.v1beta1.TallyOption) |  |  |






<a name="nemo.committee.v1beta1.MemberCommittee"></a>

### MemberCommittee
MemberCommittee is an alias of BaseCommittee


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_committee` | [BaseCommittee](#nemo.committee.v1beta1.BaseCommittee) |  |  |






<a name="nemo.committee.v1beta1.TokenCommittee"></a>

### TokenCommittee
TokenCommittee supports voting on proposals by token holders


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_committee` | [BaseCommittee](#nemo.committee.v1beta1.BaseCommittee) |  |  |
| `quorum` | [string](#string) |  |  |
| `tally_denom` | [string](#string) |  |  |





 <!-- end messages -->


<a name="nemo.committee.v1beta1.TallyOption"></a>

### TallyOption
TallyOption enumerates the valid types of a tally.

| Name | Number | Description |
| ---- | ------ | ----------- |
| TALLY_OPTION_UNSPECIFIED | 0 | TALLY_OPTION_UNSPECIFIED defines a null tally option. |
| TALLY_OPTION_FIRST_PAST_THE_POST | 1 | Votes are tallied each block and the proposal passes as soon as the vote threshold is reached |
| TALLY_OPTION_DEADLINE | 2 | Votes are tallied exactly once, when the deadline time is reached |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/committee/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/committee/v1beta1/genesis.proto



<a name="nemo.committee.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the committee module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `next_proposal_id` | [uint64](#uint64) |  |  |
| `committees` | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |
| `proposals` | [Proposal](#nemo.committee.v1beta1.Proposal) | repeated |  |
| `votes` | [Vote](#nemo.committee.v1beta1.Vote) | repeated |  |






<a name="nemo.committee.v1beta1.Proposal"></a>

### Proposal
Proposal is an internal record of a governance proposal submitted to a committee.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `content` | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| `id` | [uint64](#uint64) |  |  |
| `committee_id` | [uint64](#uint64) |  |  |
| `deadline` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="nemo.committee.v1beta1.Vote"></a>

### Vote
Vote is an internal record of a single governance vote.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `voter` | [bytes](#bytes) |  |  |
| `vote_type` | [VoteType](#nemo.committee.v1beta1.VoteType) |  |  |





 <!-- end messages -->


<a name="nemo.committee.v1beta1.VoteType"></a>

### VoteType
VoteType enumerates the valid types of a vote.

| Name | Number | Description |
| ---- | ------ | ----------- |
| VOTE_TYPE_UNSPECIFIED | 0 | VOTE_TYPE_UNSPECIFIED defines a no-op vote option. |
| VOTE_TYPE_YES | 1 | VOTE_TYPE_YES defines a yes vote option. |
| VOTE_TYPE_NO | 2 | VOTE_TYPE_NO defines a no vote option. |
| VOTE_TYPE_ABSTAIN | 3 | VOTE_TYPE_ABSTAIN defines an abstain vote option. |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/committee/v1beta1/permissions.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/committee/v1beta1/permissions.proto



<a name="nemo.committee.v1beta1.AllowedParamsChange"></a>

### AllowedParamsChange
AllowedParamsChange contains data on the allowed parameter changes for subspace, key, and sub params requirements.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `subspace` | [string](#string) |  |  |
| `key` | [string](#string) |  |  |
| `single_subparam_allowed_attrs` | [string](#string) | repeated | Requirements for when the subparam value is a single record. This contains list of allowed attribute keys that can be changed on the subparam record. |
| `multi_subparams_requirements` | [SubparamRequirement](#nemo.committee.v1beta1.SubparamRequirement) | repeated | Requirements for when the subparam value is a list of records. The requirements contains requirements for each record in the list. |






<a name="nemo.committee.v1beta1.CommunityCDPRepayDebtPermission"></a>

### CommunityCDPRepayDebtPermission
CommunityCDPRepayDebtPermission allows submission of CommunityCDPRepayDebtProposal






<a name="nemo.committee.v1beta1.CommunityCDPWithdrawCollateralPermission"></a>

### CommunityCDPWithdrawCollateralPermission
CommunityCDPWithdrawCollateralPermission allows submission of CommunityCDPWithdrawCollateralProposal






<a name="nemo.committee.v1beta1.CommunityPoolLendWithdrawPermission"></a>

### CommunityPoolLendWithdrawPermission
CommunityPoolLendWithdrawPermission allows submission of CommunityPoolLendWithdrawProposal






<a name="nemo.committee.v1beta1.GodPermission"></a>

### GodPermission
GodPermission allows any governance proposal. It is used mainly for testing.






<a name="nemo.committee.v1beta1.ParamsChangePermission"></a>

### ParamsChangePermission
ParamsChangePermission allows any parameter or sub parameter change proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `allowed_params_changes` | [AllowedParamsChange](#nemo.committee.v1beta1.AllowedParamsChange) | repeated |  |






<a name="nemo.committee.v1beta1.SoftwareUpgradePermission"></a>

### SoftwareUpgradePermission
SoftwareUpgradePermission permission type for software upgrade proposals






<a name="nemo.committee.v1beta1.SubparamRequirement"></a>

### SubparamRequirement
SubparamRequirement contains requirements for a single record in a subparam value list


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [string](#string) |  | The required attr key of the param record. |
| `val` | [string](#string) |  | The required param value for the param record key. The key and value is used to match to the target param record. |
| `allowed_subparam_attr_changes` | [string](#string) | repeated | The sub param attrs that are allowed to be changed. |






<a name="nemo.committee.v1beta1.TextPermission"></a>

### TextPermission
TextPermission allows any text governance proposal.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/committee/v1beta1/proposal.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/committee/v1beta1/proposal.proto



<a name="nemo.committee.v1beta1.CommitteeChangeProposal"></a>

### CommitteeChangeProposal
CommitteeChangeProposal is a gov proposal for creating a new committee or modifying an existing one.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `new_committee` | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="nemo.committee.v1beta1.CommitteeDeleteProposal"></a>

### CommitteeDeleteProposal
CommitteeDeleteProposal is a gov proposal for removing a committee.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `committee_id` | [uint64](#uint64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/committee/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/committee/v1beta1/query.proto



<a name="nemo.committee.v1beta1.QueryCommitteeRequest"></a>

### QueryCommitteeRequest
QueryCommitteeRequest defines the request type for querying x/committee committee.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `committee_id` | [uint64](#uint64) |  |  |






<a name="nemo.committee.v1beta1.QueryCommitteeResponse"></a>

### QueryCommitteeResponse
QueryCommitteeResponse defines the response type for querying x/committee committee.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `committee` | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="nemo.committee.v1beta1.QueryCommitteesRequest"></a>

### QueryCommitteesRequest
QueryCommitteesRequest defines the request type for querying x/committee committees.






<a name="nemo.committee.v1beta1.QueryCommitteesResponse"></a>

### QueryCommitteesResponse
QueryCommitteesResponse defines the response type for querying x/committee committees.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `committees` | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |






<a name="nemo.committee.v1beta1.QueryNextProposalIDRequest"></a>

### QueryNextProposalIDRequest
QueryNextProposalIDRequest defines the request type for querying x/committee NextProposalID.






<a name="nemo.committee.v1beta1.QueryNextProposalIDResponse"></a>

### QueryNextProposalIDResponse
QueryNextProposalIDRequest defines the response type for querying x/committee NextProposalID.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `next_proposal_id` | [uint64](#uint64) |  |  |






<a name="nemo.committee.v1beta1.QueryProposalRequest"></a>

### QueryProposalRequest
QueryProposalRequest defines the request type for querying x/committee proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |






<a name="nemo.committee.v1beta1.QueryProposalResponse"></a>

### QueryProposalResponse
QueryProposalResponse defines the response type for querying x/committee proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pub_proposal` | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| `id` | [uint64](#uint64) |  |  |
| `committee_id` | [uint64](#uint64) |  |  |
| `deadline` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="nemo.committee.v1beta1.QueryProposalsRequest"></a>

### QueryProposalsRequest
QueryProposalsRequest defines the request type for querying x/committee proposals.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `committee_id` | [uint64](#uint64) |  |  |






<a name="nemo.committee.v1beta1.QueryProposalsResponse"></a>

### QueryProposalsResponse
QueryProposalsResponse defines the response type for querying x/committee proposals.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposals` | [QueryProposalResponse](#nemo.committee.v1beta1.QueryProposalResponse) | repeated |  |






<a name="nemo.committee.v1beta1.QueryRawParamsRequest"></a>

### QueryRawParamsRequest
QueryRawParamsRequest defines the request type for querying x/committee raw params.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `subspace` | [string](#string) |  |  |
| `key` | [string](#string) |  |  |






<a name="nemo.committee.v1beta1.QueryRawParamsResponse"></a>

### QueryRawParamsResponse
QueryRawParamsResponse defines the response type for querying x/committee raw params.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `raw_data` | [string](#string) |  |  |






<a name="nemo.committee.v1beta1.QueryTallyRequest"></a>

### QueryTallyRequest
QueryTallyRequest defines the request type for querying x/committee tally.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |






<a name="nemo.committee.v1beta1.QueryTallyResponse"></a>

### QueryTallyResponse
QueryTallyResponse defines the response type for querying x/committee tally.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `yes_votes` | [string](#string) |  |  |
| `no_votes` | [string](#string) |  |  |
| `current_votes` | [string](#string) |  |  |
| `possible_votes` | [string](#string) |  |  |
| `vote_threshold` | [string](#string) |  |  |
| `quorum` | [string](#string) |  |  |






<a name="nemo.committee.v1beta1.QueryVoteRequest"></a>

### QueryVoteRequest
QueryVoteRequest defines the request type for querying x/committee vote.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `voter` | [string](#string) |  |  |






<a name="nemo.committee.v1beta1.QueryVoteResponse"></a>

### QueryVoteResponse
QueryVoteResponse defines the response type for querying x/committee vote.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `voter` | [string](#string) |  |  |
| `vote_type` | [VoteType](#nemo.committee.v1beta1.VoteType) |  |  |






<a name="nemo.committee.v1beta1.QueryVotesRequest"></a>

### QueryVotesRequest
QueryVotesRequest defines the request type for querying x/committee votes.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="nemo.committee.v1beta1.QueryVotesResponse"></a>

### QueryVotesResponse
QueryVotesResponse defines the response type for querying x/committee votes.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `votes` | [QueryVoteResponse](#nemo.committee.v1beta1.QueryVoteResponse) | repeated | votes defined the queried votes. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.committee.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for committee module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Committees` | [QueryCommitteesRequest](#nemo.committee.v1beta1.QueryCommitteesRequest) | [QueryCommitteesResponse](#nemo.committee.v1beta1.QueryCommitteesResponse) | Committees queries all committess of the committee module. | GET|/nemo/committee/v1beta1/committees|
| `Committee` | [QueryCommitteeRequest](#nemo.committee.v1beta1.QueryCommitteeRequest) | [QueryCommitteeResponse](#nemo.committee.v1beta1.QueryCommitteeResponse) | Committee queries a committee based on committee ID. | GET|/nemo/committee/v1beta1/committees/{committee_id}|
| `Proposals` | [QueryProposalsRequest](#nemo.committee.v1beta1.QueryProposalsRequest) | [QueryProposalsResponse](#nemo.committee.v1beta1.QueryProposalsResponse) | Proposals queries proposals based on committee ID. | GET|/nemo/committee/v1beta1/proposals|
| `Proposal` | [QueryProposalRequest](#nemo.committee.v1beta1.QueryProposalRequest) | [QueryProposalResponse](#nemo.committee.v1beta1.QueryProposalResponse) | Deposits queries a proposal based on proposal ID. | GET|/nemo/committee/v1beta1/proposals/{proposal_id}|
| `NextProposalID` | [QueryNextProposalIDRequest](#nemo.committee.v1beta1.QueryNextProposalIDRequest) | [QueryNextProposalIDResponse](#nemo.committee.v1beta1.QueryNextProposalIDResponse) | NextProposalID queries the next proposal ID of the committee module. | GET|/nemo/committee/v1beta1/next-proposal-id|
| `Votes` | [QueryVotesRequest](#nemo.committee.v1beta1.QueryVotesRequest) | [QueryVotesResponse](#nemo.committee.v1beta1.QueryVotesResponse) | Votes queries all votes for a single proposal ID. | GET|/nemo/committee/v1beta1/proposals/{proposal_id}/votes|
| `Vote` | [QueryVoteRequest](#nemo.committee.v1beta1.QueryVoteRequest) | [QueryVoteResponse](#nemo.committee.v1beta1.QueryVoteResponse) | Vote queries the vote of a single voter for a single proposal ID. | GET|/nemo/committee/v1beta1/proposals/{proposal_id}/votes/{voter}|
| `Tally` | [QueryTallyRequest](#nemo.committee.v1beta1.QueryTallyRequest) | [QueryTallyResponse](#nemo.committee.v1beta1.QueryTallyResponse) | Tally queries the tally of a single proposal ID. | GET|/nemo/committee/v1beta1/proposals/{proposal_id}/tally|
| `RawParams` | [QueryRawParamsRequest](#nemo.committee.v1beta1.QueryRawParamsRequest) | [QueryRawParamsResponse](#nemo.committee.v1beta1.QueryRawParamsResponse) | RawParams queries the raw params data of any subspace and key. | GET|/nemo/committee/v1beta1/raw-params|

 <!-- end services -->



<a name="nemo/committee/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/committee/v1beta1/tx.proto



<a name="nemo.committee.v1beta1.MsgSubmitProposal"></a>

### MsgSubmitProposal
MsgSubmitProposal is used by committee members to create a new proposal that they can vote on.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pub_proposal` | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| `proposer` | [string](#string) |  |  |
| `committee_id` | [uint64](#uint64) |  |  |






<a name="nemo.committee.v1beta1.MsgSubmitProposalResponse"></a>

### MsgSubmitProposalResponse
MsgSubmitProposalResponse defines the SubmitProposal response type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |






<a name="nemo.committee.v1beta1.MsgVote"></a>

### MsgVote
MsgVote is submitted by committee members to vote on proposals.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `voter` | [string](#string) |  |  |
| `vote_type` | [VoteType](#nemo.committee.v1beta1.VoteType) |  |  |






<a name="nemo.committee.v1beta1.MsgVoteResponse"></a>

### MsgVoteResponse
MsgVoteResponse defines the Vote response type





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.committee.v1beta1.Msg"></a>

### Msg
Msg defines the committee Msg service

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `SubmitProposal` | [MsgSubmitProposal](#nemo.committee.v1beta1.MsgSubmitProposal) | [MsgSubmitProposalResponse](#nemo.committee.v1beta1.MsgSubmitProposalResponse) | SubmitProposal defines a method for submitting a committee proposal | |
| `Vote` | [MsgVote](#nemo.committee.v1beta1.MsgVote) | [MsgVoteResponse](#nemo.committee.v1beta1.MsgVoteResponse) | Vote defines a method for voting on a proposal | |

 <!-- end services -->



<a name="nemo/community/v1beta1/proposal.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/community/v1beta1/proposal.proto



<a name="nemo.community.v1beta1.CommunityCDPRepayDebtProposal"></a>

### CommunityCDPRepayDebtProposal
CommunityCDPRepayDebtProposal repays a cdp debt position owned by the community module
This proposal exists primarily to allow committees to repay community module cdp debts.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |
| `payment` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="nemo.community.v1beta1.CommunityCDPWithdrawCollateralProposal"></a>

### CommunityCDPWithdrawCollateralProposal
CommunityCDPWithdrawCollateralProposal withdraws cdp collateral owned by the community module
This proposal exists primarily to allow committees to withdraw community module cdp collateral.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `collateral_type` | [string](#string) |  |  |
| `collateral` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="nemo.community.v1beta1.CommunityPoolLendDepositProposal"></a>

### CommunityPoolLendDepositProposal
CommunityPoolLendDepositProposal deposits from the community pool into lend


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="nemo.community.v1beta1.CommunityPoolLendWithdrawProposal"></a>

### CommunityPoolLendWithdrawProposal
CommunityPoolLendWithdrawProposal withdraws a lend position back to the community pool


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/community/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/community/v1beta1/query.proto



<a name="nemo.community.v1beta1.QueryBalanceRequest"></a>

### QueryBalanceRequest
QueryBalanceRequest defines the request type for querying x/community balance.






<a name="nemo.community.v1beta1.QueryBalanceResponse"></a>

### QueryBalanceResponse
QueryBalanceResponse defines the response type for querying x/community balance.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="nemo.community.v1beta1.QueryTotalBalanceRequest"></a>

### QueryTotalBalanceRequest
QueryTotalBalanceRequest defines the request type for querying total community pool balance.






<a name="nemo.community.v1beta1.QueryTotalBalanceResponse"></a>

### QueryTotalBalanceResponse
QueryTotalBalanceResponse defines the response type for querying total
community pool balance. This matches the x/distribution CommunityPool query response.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated | pool defines community pool's coins. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.community.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for x/community.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Balance` | [QueryBalanceRequest](#nemo.community.v1beta1.QueryBalanceRequest) | [QueryBalanceResponse](#nemo.community.v1beta1.QueryBalanceResponse) | Balance queries the balance of all coins of x/community module. | GET|/nemo/community/v1beta1/balance|
| `TotalBalance` | [QueryTotalBalanceRequest](#nemo.community.v1beta1.QueryTotalBalanceRequest) | [QueryTotalBalanceResponse](#nemo.community.v1beta1.QueryTotalBalanceResponse) | TotalBalance queries the balance of all coins, including x/distribution, x/community, and supplied balances. | GET|/nemo/community/v1beta1/total_balance|

 <!-- end services -->



<a name="nemo/community/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/community/v1beta1/tx.proto



<a name="nemo.community.v1beta1.MsgFundCommunityPool"></a>

### MsgFundCommunityPool
MsgFundCommunityPool allows an account to directly fund the community module account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `depositor` | [string](#string) |  |  |






<a name="nemo.community.v1beta1.MsgFundCommunityPoolResponse"></a>

### MsgFundCommunityPoolResponse
MsgFundCommunityPoolResponse defines the Msg/FundCommunityPool response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.community.v1beta1.Msg"></a>

### Msg
Msg defines the community Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `FundCommunityPool` | [MsgFundCommunityPool](#nemo.community.v1beta1.MsgFundCommunityPool) | [MsgFundCommunityPoolResponse](#nemo.community.v1beta1.MsgFundCommunityPoolResponse) | FundCommunityPool defines a method to allow an account to directly fund the community module account. | |

 <!-- end services -->



<a name="nemo/earn/v1beta1/strategy.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/earn/v1beta1/strategy.proto


 <!-- end messages -->


<a name="nemo.earn.v1beta1.StrategyType"></a>

### StrategyType
StrategyType is the type of strategy that a vault uses to optimize yields.

| Name | Number | Description |
| ---- | ------ | ----------- |
| STRATEGY_TYPE_UNSPECIFIED | 0 | STRATEGY_TYPE_UNSPECIFIED represents an unspecified or invalid strategy type. |
| STRATEGY_TYPE_JINX | 1 | STRATEGY_TYPE_JINX represents the strategy that deposits assets in the Jinx module. |
| STRATEGY_TYPE_SAVINGS | 2 | STRATEGY_TYPE_SAVINGS represents the strategy that deposits assets in the Savings module. |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/earn/v1beta1/vault.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/earn/v1beta1/vault.proto



<a name="nemo.earn.v1beta1.AllowedVault"></a>

### AllowedVault
AllowedVault is a vault that is allowed to be created. These can be
modified via parameter governance.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  | Denom is the only supported denomination of the vault for deposits and withdrawals. |
| `strategies` | [StrategyType](#nemo.earn.v1beta1.StrategyType) | repeated | VaultStrategy is the strategy used for this vault. |
| `is_private_vault` | [bool](#bool) |  | IsPrivateVault is true if the vault only allows depositors contained in AllowedDepositors. |
| `allowed_depositors` | [bytes](#bytes) | repeated | AllowedDepositors is a list of addresses that are allowed to deposit to this vault if IsPrivateVault is true. Addresses not contained in this list are not allowed to deposit into this vault. If IsPrivateVault is false, this should be empty and ignored. |






<a name="nemo.earn.v1beta1.VaultRecord"></a>

### VaultRecord
VaultRecord is the state of a vault.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `total_shares` | [VaultShare](#nemo.earn.v1beta1.VaultShare) |  | TotalShares is the total distributed number of shares in the vault. |






<a name="nemo.earn.v1beta1.VaultShare"></a>

### VaultShare
VaultShare defines shares of a vault owned by a depositor.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `amount` | [string](#string) |  |  |






<a name="nemo.earn.v1beta1.VaultShareRecord"></a>

### VaultShareRecord
VaultShareRecord defines the vault shares owned by a depositor.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [bytes](#bytes) |  | Depositor represents the owner of the shares |
| `shares` | [VaultShare](#nemo.earn.v1beta1.VaultShare) | repeated | Shares represent the vault shares owned by the depositor. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/earn/v1beta1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/earn/v1beta1/params.proto



<a name="nemo.earn.v1beta1.Params"></a>

### Params
Params defines the parameters of the earn module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `allowed_vaults` | [AllowedVault](#nemo.earn.v1beta1.AllowedVault) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/earn/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/earn/v1beta1/genesis.proto



<a name="nemo.earn.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the earn module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#nemo.earn.v1beta1.Params) |  | params defines all the paramaters related to earn |
| `vault_records` | [VaultRecord](#nemo.earn.v1beta1.VaultRecord) | repeated | vault_records defines the available vaults |
| `vault_share_records` | [VaultShareRecord](#nemo.earn.v1beta1.VaultShareRecord) | repeated | share_records defines the owned shares of each vault |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/earn/v1beta1/proposal.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/earn/v1beta1/proposal.proto



<a name="nemo.earn.v1beta1.CommunityPoolDepositProposal"></a>

### CommunityPoolDepositProposal
CommunityPoolDepositProposal deposits from the community pool into an earn vault


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="nemo.earn.v1beta1.CommunityPoolDepositProposalJSON"></a>

### CommunityPoolDepositProposalJSON
CommunityPoolDepositProposalJSON defines a CommunityPoolDepositProposal with a deposit


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `deposit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="nemo.earn.v1beta1.CommunityPoolWithdrawProposal"></a>

### CommunityPoolWithdrawProposal
CommunityPoolWithdrawProposal withdraws from an earn vault back to community pool


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="nemo.earn.v1beta1.CommunityPoolWithdrawProposalJSON"></a>

### CommunityPoolWithdrawProposalJSON
CommunityPoolWithdrawProposalJSON defines a CommunityPoolWithdrawProposal with a deposit


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `deposit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/earn/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/earn/v1beta1/query.proto



<a name="nemo.earn.v1beta1.DepositResponse"></a>

### DepositResponse
DepositResponse defines a deposit query response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  | depositor represents the owner of the deposit. |
| `shares` | [VaultShare](#nemo.earn.v1beta1.VaultShare) | repeated | Shares represent the issued shares from their corresponding vaults. |
| `value` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | Value represents the total accumulated value of denom coins supplied to vaults. This may be greater than or equal to amount_supplied depending on the strategy. |






<a name="nemo.earn.v1beta1.QueryDepositsRequest"></a>

### QueryDepositsRequest
QueryDepositsRequest is the request type for the Query/Deposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  | depositor optionally filters deposits by depositor |
| `denom` | [string](#string) |  | denom optionally filters deposits by vault denom |
| `value_in_staked_tokens` | [bool](#bool) |  | respond with vault value in ufury for bfury vaults |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="nemo.earn.v1beta1.QueryDepositsResponse"></a>

### QueryDepositsResponse
QueryDepositsResponse is the response type for the Query/Deposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposits` | [DepositResponse](#nemo.earn.v1beta1.DepositResponse) | repeated | deposits returns the deposits matching the requested parameters |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="nemo.earn.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest defines the request type for querying x/earn parameters.






<a name="nemo.earn.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse defines the response type for querying x/earn parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#nemo.earn.v1beta1.Params) |  | params represents the earn module parameters |






<a name="nemo.earn.v1beta1.QueryTotalSupplyRequest"></a>

### QueryTotalSupplyRequest
QueryTotalSupplyRequest defines the request type for Query/TotalSupply method.






<a name="nemo.earn.v1beta1.QueryTotalSupplyResponse"></a>

### QueryTotalSupplyResponse
TotalSupplyResponse defines the response type for the Query/TotalSupply method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `height` | [int64](#int64) |  | Height is the block height at which these totals apply |
| `result` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | Result is a list of coins supplied to earn |






<a name="nemo.earn.v1beta1.QueryVaultRequest"></a>

### QueryVaultRequest
QueryVaultRequest is the request type for the Query/Vault RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  | vault filters vault by denom |






<a name="nemo.earn.v1beta1.QueryVaultResponse"></a>

### QueryVaultResponse
QueryVaultResponse is the response type for the Query/Vault RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `vault` | [VaultResponse](#nemo.earn.v1beta1.VaultResponse) |  | vault represents the queried earn module vault |






<a name="nemo.earn.v1beta1.QueryVaultsRequest"></a>

### QueryVaultsRequest
QueryVaultsRequest is the request type for the Query/Vaults RPC method.






<a name="nemo.earn.v1beta1.QueryVaultsResponse"></a>

### QueryVaultsResponse
QueryVaultsResponse is the response type for the Query/Vaults RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `vaults` | [VaultResponse](#nemo.earn.v1beta1.VaultResponse) | repeated | vaults represents the earn module vaults |






<a name="nemo.earn.v1beta1.VaultResponse"></a>

### VaultResponse
VaultResponse is the response type for a vault.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  | denom represents the denom of the vault |
| `strategies` | [StrategyType](#nemo.earn.v1beta1.StrategyType) | repeated | VaultStrategy is the strategy used for this vault. |
| `is_private_vault` | [bool](#bool) |  | IsPrivateVault is true if the vault only allows depositors contained in AllowedDepositors. |
| `allowed_depositors` | [string](#string) | repeated | AllowedDepositors is a list of addresses that are allowed to deposit to this vault if IsPrivateVault is true. Addresses not contained in this list are not allowed to deposit into this vault. If IsPrivateVault is false, this should be empty and ignored. |
| `total_shares` | [string](#string) |  | TotalShares is the total amount of shares issued to depositors. |
| `total_value` | [string](#string) |  | TotalValue is the total value of denom coins supplied to the vault if the vault were to be liquidated. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.earn.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for earn module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#nemo.earn.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#nemo.earn.v1beta1.QueryParamsResponse) | Params queries all parameters of the earn module. | GET|/nemo/earn/v1beta1/params|
| `Vaults` | [QueryVaultsRequest](#nemo.earn.v1beta1.QueryVaultsRequest) | [QueryVaultsResponse](#nemo.earn.v1beta1.QueryVaultsResponse) | Vaults queries all vaults | GET|/nemo/earn/v1beta1/vaults|
| `Vault` | [QueryVaultRequest](#nemo.earn.v1beta1.QueryVaultRequest) | [QueryVaultResponse](#nemo.earn.v1beta1.QueryVaultResponse) | Vault queries a single vault based on the vault denom | GET|/nemo/earn/v1beta1/vaults/{denom=**}|
| `Deposits` | [QueryDepositsRequest](#nemo.earn.v1beta1.QueryDepositsRequest) | [QueryDepositsResponse](#nemo.earn.v1beta1.QueryDepositsResponse) | Deposits queries deposit details based on depositor address and vault | GET|/nemo/earn/v1beta1/deposits|
| `TotalSupply` | [QueryTotalSupplyRequest](#nemo.earn.v1beta1.QueryTotalSupplyRequest) | [QueryTotalSupplyResponse](#nemo.earn.v1beta1.QueryTotalSupplyResponse) | TotalSupply returns the total sum of all coins currently locked into the earn module. | GET|/nemo/earn/v1beta1/total_supply|

 <!-- end services -->



<a name="nemo/earn/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/earn/v1beta1/tx.proto



<a name="nemo.earn.v1beta1.MsgDeposit"></a>

### MsgDeposit
MsgDeposit represents a message for depositing assedts into a vault


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  | depositor represents the address to deposit funds from |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Amount represents the token to deposit. The vault corresponds to the denom of the amount coin. |
| `strategy` | [StrategyType](#nemo.earn.v1beta1.StrategyType) |  | Strategy is the vault strategy to use. |






<a name="nemo.earn.v1beta1.MsgDepositResponse"></a>

### MsgDepositResponse
MsgDepositResponse defines the Msg/Deposit response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `shares` | [VaultShare](#nemo.earn.v1beta1.VaultShare) |  |  |






<a name="nemo.earn.v1beta1.MsgWithdraw"></a>

### MsgWithdraw
MsgWithdraw represents a message for withdrawing liquidity from a vault


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from` | [string](#string) |  | from represents the address we are withdrawing for |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Amount represents the token to withdraw. The vault corresponds to the denom of the amount coin. |
| `strategy` | [StrategyType](#nemo.earn.v1beta1.StrategyType) |  | Strategy is the vault strategy to use. |






<a name="nemo.earn.v1beta1.MsgWithdrawResponse"></a>

### MsgWithdrawResponse
MsgWithdrawResponse defines the Msg/Withdraw response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `shares` | [VaultShare](#nemo.earn.v1beta1.VaultShare) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.earn.v1beta1.Msg"></a>

### Msg
Msg defines the earn Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Deposit` | [MsgDeposit](#nemo.earn.v1beta1.MsgDeposit) | [MsgDepositResponse](#nemo.earn.v1beta1.MsgDepositResponse) | Deposit defines a method for depositing assets into a vault | |
| `Withdraw` | [MsgWithdraw](#nemo.earn.v1beta1.MsgWithdraw) | [MsgWithdrawResponse](#nemo.earn.v1beta1.MsgWithdrawResponse) | Withdraw defines a method for withdrawing assets into a vault | |

 <!-- end services -->



<a name="nemo/evmutil/v1beta1/conversion_pair.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/evmutil/v1beta1/conversion_pair.proto



<a name="nemo.evmutil.v1beta1.AllowedCosmosCoinERC20Token"></a>

### AllowedCosmosCoinERC20Token
AllowedCosmosCoinERC20Token defines allowed cosmos-sdk denom & metadata
for evm token representations of sdk assets.
NOTE: once evm token contracts are deployed, changes to metadata for a given
cosmos_denom will not change metadata of deployed contract.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cosmos_denom` | [string](#string) |  | Denom of the sdk.Coin |
| `name` | [string](#string) |  | Name of ERC20 contract |
| `symbol` | [string](#string) |  | Symbol of ERC20 contract |
| `decimals` | [uint32](#uint32) |  | Number of decimals ERC20 contract is deployed with. |






<a name="nemo.evmutil.v1beta1.ConversionPair"></a>

### ConversionPair
ConversionPair defines a Nemo ERC20 address and corresponding denom that is
allowed to be converted between ERC20 and sdk.Coin


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nemo_erc20_address` | [bytes](#bytes) |  | ERC20 address of the token on the Nemo EVM |
| `denom` | [string](#string) |  | Denom of the corresponding sdk.Coin |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/evmutil/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/evmutil/v1beta1/genesis.proto



<a name="nemo.evmutil.v1beta1.Account"></a>

### Account
BalanceAccount defines an account in the evmutil module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [bytes](#bytes) |  |  |
| `balance` | [string](#string) |  | balance indicates the amount of afury owned by the address. |






<a name="nemo.evmutil.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the evmutil module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `accounts` | [Account](#nemo.evmutil.v1beta1.Account) | repeated |  |
| `params` | [Params](#nemo.evmutil.v1beta1.Params) |  | params defines all the parameters of the module. |






<a name="nemo.evmutil.v1beta1.Params"></a>

### Params
Params defines the evmutil module params


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `enabled_conversion_pairs` | [ConversionPair](#nemo.evmutil.v1beta1.ConversionPair) | repeated | enabled_conversion_pairs defines the list of conversion pairs allowed to be converted between Nemo ERC20 and sdk.Coin |
| `allowed_cosmos_denoms` | [AllowedCosmosCoinERC20Token](#nemo.evmutil.v1beta1.AllowedCosmosCoinERC20Token) | repeated | allowed_cosmos_denoms is a list of denom & erc20 token metadata pairs. if a denom is in the list, it is allowed to be converted to an erc20 in the evm. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/evmutil/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/evmutil/v1beta1/query.proto



<a name="nemo.evmutil.v1beta1.DeployedCosmosCoinContract"></a>

### DeployedCosmosCoinContract
DeployedCosmosCoinContract defines a deployed token contract to the evm representing a native cosmos-sdk coin


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cosmos_denom` | [string](#string) |  |  |
| `address` | [string](#string) |  |  |






<a name="nemo.evmutil.v1beta1.QueryDeployedCosmosCoinContractsRequest"></a>

### QueryDeployedCosmosCoinContractsRequest
QueryDeployedCosmosCoinContractsRequest defines the request type for Query/DeployedCosmosCoinContracts method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cosmos_denoms` | [string](#string) | repeated | optional query param to only return specific denoms in the list denoms that do not have deployed contracts will be omitted from the result must request fewer than 100 denoms at a time. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="nemo.evmutil.v1beta1.QueryDeployedCosmosCoinContractsResponse"></a>

### QueryDeployedCosmosCoinContractsResponse
QueryDeployedCosmosCoinContractsResponse defines the response type for the Query/DeployedCosmosCoinContracts method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deployed_cosmos_coin_contracts` | [DeployedCosmosCoinContract](#nemo.evmutil.v1beta1.DeployedCosmosCoinContract) | repeated | deployed_cosmos_coin_contracts is a list of cosmos-sdk coin denom and its deployed contract address |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="nemo.evmutil.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest defines the request type for querying x/evmutil parameters.






<a name="nemo.evmutil.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse defines the response type for querying x/evmutil parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#nemo.evmutil.v1beta1.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.evmutil.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for evmutil module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#nemo.evmutil.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#nemo.evmutil.v1beta1.QueryParamsResponse) | Params queries all parameters of the evmutil module. | GET|/nemo/evmutil/v1beta1/params|
| `DeployedCosmosCoinContracts` | [QueryDeployedCosmosCoinContractsRequest](#nemo.evmutil.v1beta1.QueryDeployedCosmosCoinContractsRequest) | [QueryDeployedCosmosCoinContractsResponse](#nemo.evmutil.v1beta1.QueryDeployedCosmosCoinContractsResponse) | DeployedCosmosCoinContracts queries a list cosmos coin denom and their deployed erc20 address | GET|/nemo/evmutil/v1beta1/deployed_cosmos_coin_contracts|

 <!-- end services -->



<a name="nemo/evmutil/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/evmutil/v1beta1/tx.proto



<a name="nemo.evmutil.v1beta1.MsgConvertCoinToERC20"></a>

### MsgConvertCoinToERC20
MsgConvertCoinToERC20 defines a conversion from sdk.Coin to Nemo ERC20 for EVM-native assets.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `initiator` | [string](#string) |  | Nemo bech32 address initiating the conversion. |
| `receiver` | [string](#string) |  | EVM 0x hex address that will receive the converted Nemo ERC20 tokens. |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Amount is the sdk.Coin amount to convert. |






<a name="nemo.evmutil.v1beta1.MsgConvertCoinToERC20Response"></a>

### MsgConvertCoinToERC20Response
MsgConvertCoinToERC20Response defines the response value from Msg/ConvertCoinToERC20.






<a name="nemo.evmutil.v1beta1.MsgConvertCosmosCoinFromERC20"></a>

### MsgConvertCosmosCoinFromERC20
MsgConvertCosmosCoinFromERC20 defines a conversion from ERC20 to cosmos coins for cosmos-native assets.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `initiator` | [string](#string) |  | EVM hex address initiating the conversion. |
| `receiver` | [string](#string) |  | Nemo bech32 address that will receive the cosmos coins. |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Amount is the amount to convert, expressed as a Cosmos coin. |






<a name="nemo.evmutil.v1beta1.MsgConvertCosmosCoinFromERC20Response"></a>

### MsgConvertCosmosCoinFromERC20Response
MsgConvertCosmosCoinFromERC20Response defines the response value from Msg/MsgConvertCosmosCoinFromERC20.






<a name="nemo.evmutil.v1beta1.MsgConvertCosmosCoinToERC20"></a>

### MsgConvertCosmosCoinToERC20
MsgConvertCosmosCoinToERC20 defines a conversion from cosmos sdk.Coin to ERC20 for cosmos-native assets.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `initiator` | [string](#string) |  | Nemo bech32 address initiating the conversion. |
| `receiver` | [string](#string) |  | EVM hex address that will receive the ERC20 tokens. |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Amount is the sdk.Coin amount to convert. |






<a name="nemo.evmutil.v1beta1.MsgConvertCosmosCoinToERC20Response"></a>

### MsgConvertCosmosCoinToERC20Response
MsgConvertCosmosCoinToERC20Response defines the response value from Msg/MsgConvertCosmosCoinToERC20.






<a name="nemo.evmutil.v1beta1.MsgConvertERC20ToCoin"></a>

### MsgConvertERC20ToCoin
MsgConvertERC20ToCoin defines a conversion from Nemo ERC20 to sdk.Coin for EVM-native assets.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `initiator` | [string](#string) |  | EVM 0x hex address initiating the conversion. |
| `receiver` | [string](#string) |  | Nemo bech32 address that will receive the converted sdk.Coin. |
| `nemo_erc20_address` | [string](#string) |  | EVM 0x hex address of the ERC20 contract. |
| `amount` | [string](#string) |  | ERC20 token amount to convert. |






<a name="nemo.evmutil.v1beta1.MsgConvertERC20ToCoinResponse"></a>

### MsgConvertERC20ToCoinResponse
MsgConvertERC20ToCoinResponse defines the response value from
Msg/MsgConvertERC20ToCoin.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.evmutil.v1beta1.Msg"></a>

### Msg
Msg defines the evmutil Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ConvertCoinToERC20` | [MsgConvertCoinToERC20](#nemo.evmutil.v1beta1.MsgConvertCoinToERC20) | [MsgConvertCoinToERC20Response](#nemo.evmutil.v1beta1.MsgConvertCoinToERC20Response) | ConvertCoinToERC20 defines a method for converting sdk.Coin to Nemo ERC20. | |
| `ConvertERC20ToCoin` | [MsgConvertERC20ToCoin](#nemo.evmutil.v1beta1.MsgConvertERC20ToCoin) | [MsgConvertERC20ToCoinResponse](#nemo.evmutil.v1beta1.MsgConvertERC20ToCoinResponse) | ConvertERC20ToCoin defines a method for converting Nemo ERC20 to sdk.Coin. | |
| `ConvertCosmosCoinToERC20` | [MsgConvertCosmosCoinToERC20](#nemo.evmutil.v1beta1.MsgConvertCosmosCoinToERC20) | [MsgConvertCosmosCoinToERC20Response](#nemo.evmutil.v1beta1.MsgConvertCosmosCoinToERC20Response) | ConvertCosmosCoinToERC20 defines a method for converting a cosmos sdk.Coin to an ERC20. | |
| `ConvertCosmosCoinFromERC20` | [MsgConvertCosmosCoinFromERC20](#nemo.evmutil.v1beta1.MsgConvertCosmosCoinFromERC20) | [MsgConvertCosmosCoinFromERC20Response](#nemo.evmutil.v1beta1.MsgConvertCosmosCoinFromERC20Response) | ConvertCosmosCoinFromERC20 defines a method for converting a cosmos sdk.Coin to an ERC20. | |

 <!-- end services -->



<a name="nemo/jinx/v1beta1/jinx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/jinx/v1beta1/jinx.proto



<a name="nemo.jinx.v1beta1.Borrow"></a>

### Borrow
Borrow defines an amount of coins borrowed from a jinx module account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `borrower` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `index` | [BorrowInterestFactor](#nemo.jinx.v1beta1.BorrowInterestFactor) | repeated |  |






<a name="nemo.jinx.v1beta1.BorrowInterestFactor"></a>

### BorrowInterestFactor
BorrowInterestFactor defines an individual borrow interest factor.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `value` | [string](#string) |  |  |






<a name="nemo.jinx.v1beta1.BorrowLimit"></a>

### BorrowLimit
BorrowLimit enforces restrictions on a money market.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `has_max_limit` | [bool](#bool) |  |  |
| `maximum_limit` | [string](#string) |  |  |
| `loan_to_value` | [string](#string) |  |  |






<a name="nemo.jinx.v1beta1.CoinsProto"></a>

### CoinsProto
CoinsProto defines a Protobuf wrapper around a Coins slice


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="nemo.jinx.v1beta1.Deposit"></a>

### Deposit
Deposit defines an amount of coins deposited into a jinx module account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `index` | [SupplyInterestFactor](#nemo.jinx.v1beta1.SupplyInterestFactor) | repeated |  |






<a name="nemo.jinx.v1beta1.InterestRateModel"></a>

### InterestRateModel
InterestRateModel contains information about an asset's interest rate.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_rate_apy` | [string](#string) |  |  |
| `base_multiplier` | [string](#string) |  |  |
| `kink` | [string](#string) |  |  |
| `jump_multiplier` | [string](#string) |  |  |






<a name="nemo.jinx.v1beta1.MoneyMarket"></a>

### MoneyMarket
MoneyMarket is a money market for an individual asset.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `borrow_limit` | [BorrowLimit](#nemo.jinx.v1beta1.BorrowLimit) |  |  |
| `spot_market_id` | [string](#string) |  |  |
| `conversion_factor` | [string](#string) |  |  |
| `interest_rate_model` | [InterestRateModel](#nemo.jinx.v1beta1.InterestRateModel) |  |  |
| `reserve_factor` | [string](#string) |  |  |
| `keeper_reward_percentage` | [string](#string) |  |  |






<a name="nemo.jinx.v1beta1.Params"></a>

### Params
Params defines the parameters for the jinx module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `money_markets` | [MoneyMarket](#nemo.jinx.v1beta1.MoneyMarket) | repeated |  |
| `minimum_borrow_usd_value` | [string](#string) |  |  |






<a name="nemo.jinx.v1beta1.SupplyInterestFactor"></a>

### SupplyInterestFactor
SupplyInterestFactor defines an individual borrow interest factor.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `value` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/jinx/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/jinx/v1beta1/genesis.proto



<a name="nemo.jinx.v1beta1.GenesisAccumulationTime"></a>

### GenesisAccumulationTime
GenesisAccumulationTime stores the previous distribution time and its corresponding denom.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `previous_accumulation_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `supply_interest_factor` | [string](#string) |  |  |
| `borrow_interest_factor` | [string](#string) |  |  |






<a name="nemo.jinx.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the jinx module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#nemo.jinx.v1beta1.Params) |  |  |
| `previous_accumulation_times` | [GenesisAccumulationTime](#nemo.jinx.v1beta1.GenesisAccumulationTime) | repeated |  |
| `deposits` | [Deposit](#nemo.jinx.v1beta1.Deposit) | repeated |  |
| `borrows` | [Borrow](#nemo.jinx.v1beta1.Borrow) | repeated |  |
| `total_supplied` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `total_borrowed` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `total_reserves` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/jinx/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/jinx/v1beta1/query.proto



<a name="nemo.jinx.v1beta1.BorrowInterestFactorResponse"></a>

### BorrowInterestFactorResponse
BorrowInterestFactorResponse defines an individual borrow interest factor.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `value` | [string](#string) |  | sdk.Dec as string |






<a name="nemo.jinx.v1beta1.BorrowResponse"></a>

### BorrowResponse
BorrowResponse defines an amount of coins borrowed from a jinx module account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `borrower` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `index` | [BorrowInterestFactorResponse](#nemo.jinx.v1beta1.BorrowInterestFactorResponse) | repeated |  |






<a name="nemo.jinx.v1beta1.DepositResponse"></a>

### DepositResponse
DepositResponse defines an amount of coins deposited into a jinx module account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `index` | [SupplyInterestFactorResponse](#nemo.jinx.v1beta1.SupplyInterestFactorResponse) | repeated |  |






<a name="nemo.jinx.v1beta1.InterestFactor"></a>

### InterestFactor
InterestFactor is a unique type returned by interest factor queries


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `borrow_interest_factor` | [string](#string) |  | sdk.Dec as String |
| `supply_interest_factor` | [string](#string) |  | sdk.Dec as String |






<a name="nemo.jinx.v1beta1.MoneyMarketInterestRate"></a>

### MoneyMarketInterestRate
MoneyMarketInterestRate is a unique type returned by interest rate queries


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `supply_interest_rate` | [string](#string) |  | sdk.Dec as String |
| `borrow_interest_rate` | [string](#string) |  | sdk.Dec as String |






<a name="nemo.jinx.v1beta1.QueryAccountsRequest"></a>

### QueryAccountsRequest
QueryAccountsRequest is the request type for the Query/Accounts RPC method.






<a name="nemo.jinx.v1beta1.QueryAccountsResponse"></a>

### QueryAccountsResponse
QueryAccountsResponse is the response type for the Query/Accounts RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `accounts` | [cosmos.auth.v1beta1.ModuleAccount](#cosmos.auth.v1beta1.ModuleAccount) | repeated |  |






<a name="nemo.jinx.v1beta1.QueryBorrowsRequest"></a>

### QueryBorrowsRequest
QueryBorrowsRequest is the request type for the Query/Borrows RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="nemo.jinx.v1beta1.QueryBorrowsResponse"></a>

### QueryBorrowsResponse
QueryBorrowsResponse is the response type for the Query/Borrows RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `borrows` | [BorrowResponse](#nemo.jinx.v1beta1.BorrowResponse) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="nemo.jinx.v1beta1.QueryDepositsRequest"></a>

### QueryDepositsRequest
QueryDepositsRequest is the request type for the Query/Deposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="nemo.jinx.v1beta1.QueryDepositsResponse"></a>

### QueryDepositsResponse
QueryDepositsResponse is the response type for the Query/Deposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposits` | [DepositResponse](#nemo.jinx.v1beta1.DepositResponse) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="nemo.jinx.v1beta1.QueryInterestFactorsRequest"></a>

### QueryInterestFactorsRequest
QueryInterestFactorsRequest is the request type for the Query/InterestFactors RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |






<a name="nemo.jinx.v1beta1.QueryInterestFactorsResponse"></a>

### QueryInterestFactorsResponse
QueryInterestFactorsResponse is the response type for the Query/InterestFactors RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `interest_factors` | [InterestFactor](#nemo.jinx.v1beta1.InterestFactor) | repeated |  |






<a name="nemo.jinx.v1beta1.QueryInterestRateRequest"></a>

### QueryInterestRateRequest
QueryInterestRateRequest is the request type for the Query/InterestRate RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |






<a name="nemo.jinx.v1beta1.QueryInterestRateResponse"></a>

### QueryInterestRateResponse
QueryInterestRateResponse is the response type for the Query/InterestRate RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `interest_rates` | [MoneyMarketInterestRate](#nemo.jinx.v1beta1.MoneyMarketInterestRate) | repeated |  |






<a name="nemo.jinx.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="nemo.jinx.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#nemo.jinx.v1beta1.Params) |  |  |






<a name="nemo.jinx.v1beta1.QueryReservesRequest"></a>

### QueryReservesRequest
QueryReservesRequest is the request type for the Query/Reserves RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |






<a name="nemo.jinx.v1beta1.QueryReservesResponse"></a>

### QueryReservesResponse
QueryReservesResponse is the response type for the Query/Reserves RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="nemo.jinx.v1beta1.QueryTotalBorrowedRequest"></a>

### QueryTotalBorrowedRequest
QueryTotalBorrowedRequest is the request type for the Query/TotalBorrowed RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |






<a name="nemo.jinx.v1beta1.QueryTotalBorrowedResponse"></a>

### QueryTotalBorrowedResponse
QueryTotalBorrowedResponse is the response type for the Query/TotalBorrowed RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `borrowed_coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="nemo.jinx.v1beta1.QueryTotalDepositedRequest"></a>

### QueryTotalDepositedRequest
QueryTotalDepositedRequest is the request type for the Query/TotalDeposited RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |






<a name="nemo.jinx.v1beta1.QueryTotalDepositedResponse"></a>

### QueryTotalDepositedResponse
QueryTotalDepositedResponse is the response type for the Query/TotalDeposited RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `supplied_coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="nemo.jinx.v1beta1.QueryUnsyncedBorrowsRequest"></a>

### QueryUnsyncedBorrowsRequest
QueryUnsyncedBorrowsRequest is the request type for the Query/UnsyncedBorrows RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="nemo.jinx.v1beta1.QueryUnsyncedBorrowsResponse"></a>

### QueryUnsyncedBorrowsResponse
QueryUnsyncedBorrowsResponse is the response type for the Query/UnsyncedBorrows RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `borrows` | [BorrowResponse](#nemo.jinx.v1beta1.BorrowResponse) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="nemo.jinx.v1beta1.QueryUnsyncedDepositsRequest"></a>

### QueryUnsyncedDepositsRequest
QueryUnsyncedDepositsRequest is the request type for the Query/UnsyncedDeposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="nemo.jinx.v1beta1.QueryUnsyncedDepositsResponse"></a>

### QueryUnsyncedDepositsResponse
QueryUnsyncedDepositsResponse is the response type for the Query/UnsyncedDeposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposits` | [DepositResponse](#nemo.jinx.v1beta1.DepositResponse) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="nemo.jinx.v1beta1.SupplyInterestFactorResponse"></a>

### SupplyInterestFactorResponse
SupplyInterestFactorResponse defines an individual borrow interest factor.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `value` | [string](#string) |  | sdk.Dec as string |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.jinx.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for bep3 module.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#nemo.jinx.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#nemo.jinx.v1beta1.QueryParamsResponse) | Params queries module params. | GET|/nemo/jinx/v1beta1/params|
| `Accounts` | [QueryAccountsRequest](#nemo.jinx.v1beta1.QueryAccountsRequest) | [QueryAccountsResponse](#nemo.jinx.v1beta1.QueryAccountsResponse) | Accounts queries module accounts. | GET|/nemo/jinx/v1beta1/accounts|
| `Deposits` | [QueryDepositsRequest](#nemo.jinx.v1beta1.QueryDepositsRequest) | [QueryDepositsResponse](#nemo.jinx.v1beta1.QueryDepositsResponse) | Deposits queries jinx deposits. | GET|/nemo/jinx/v1beta1/deposits|
| `UnsyncedDeposits` | [QueryUnsyncedDepositsRequest](#nemo.jinx.v1beta1.QueryUnsyncedDepositsRequest) | [QueryUnsyncedDepositsResponse](#nemo.jinx.v1beta1.QueryUnsyncedDepositsResponse) | UnsyncedDeposits queries unsynced deposits. | GET|/nemo/jinx/v1beta1/unsynced-deposits|
| `TotalDeposited` | [QueryTotalDepositedRequest](#nemo.jinx.v1beta1.QueryTotalDepositedRequest) | [QueryTotalDepositedResponse](#nemo.jinx.v1beta1.QueryTotalDepositedResponse) | TotalDeposited queries total coins deposited to jinx liquidity pools. | GET|/nemo/jinx/v1beta1/total-deposited|
| `Borrows` | [QueryBorrowsRequest](#nemo.jinx.v1beta1.QueryBorrowsRequest) | [QueryBorrowsResponse](#nemo.jinx.v1beta1.QueryBorrowsResponse) | Borrows queries jinx borrows. | GET|/nemo/jinx/v1beta1/borrows|
| `UnsyncedBorrows` | [QueryUnsyncedBorrowsRequest](#nemo.jinx.v1beta1.QueryUnsyncedBorrowsRequest) | [QueryUnsyncedBorrowsResponse](#nemo.jinx.v1beta1.QueryUnsyncedBorrowsResponse) | UnsyncedBorrows queries unsynced borrows. | GET|/nemo/jinx/v1beta1/unsynced-borrows|
| `TotalBorrowed` | [QueryTotalBorrowedRequest](#nemo.jinx.v1beta1.QueryTotalBorrowedRequest) | [QueryTotalBorrowedResponse](#nemo.jinx.v1beta1.QueryTotalBorrowedResponse) | TotalBorrowed queries total coins borrowed from jinx liquidity pools. | GET|/nemo/jinx/v1beta1/total-borrowed|
| `InterestRate` | [QueryInterestRateRequest](#nemo.jinx.v1beta1.QueryInterestRateRequest) | [QueryInterestRateResponse](#nemo.jinx.v1beta1.QueryInterestRateResponse) | InterestRate queries the jinx module interest rates. | GET|/nemo/jinx/v1beta1/interest-rate|
| `Reserves` | [QueryReservesRequest](#nemo.jinx.v1beta1.QueryReservesRequest) | [QueryReservesResponse](#nemo.jinx.v1beta1.QueryReservesResponse) | Reserves queries total jinx reserve coins. | GET|/nemo/jinx/v1beta1/reserves|
| `InterestFactors` | [QueryInterestFactorsRequest](#nemo.jinx.v1beta1.QueryInterestFactorsRequest) | [QueryInterestFactorsResponse](#nemo.jinx.v1beta1.QueryInterestFactorsResponse) | InterestFactors queries jinx module interest factors. | GET|/nemo/jinx/v1beta1/interest-factors|

 <!-- end services -->



<a name="nemo/jinx/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/jinx/v1beta1/tx.proto



<a name="nemo.jinx.v1beta1.MsgBorrow"></a>

### MsgBorrow
MsgBorrow defines the Msg/Borrow request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `borrower` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="nemo.jinx.v1beta1.MsgBorrowResponse"></a>

### MsgBorrowResponse
MsgBorrowResponse defines the Msg/Borrow response type.






<a name="nemo.jinx.v1beta1.MsgDeposit"></a>

### MsgDeposit
MsgDeposit defines the Msg/Deposit request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="nemo.jinx.v1beta1.MsgDepositResponse"></a>

### MsgDepositResponse
MsgDepositResponse defines the Msg/Deposit response type.






<a name="nemo.jinx.v1beta1.MsgLiquidate"></a>

### MsgLiquidate
MsgLiquidate defines the Msg/Liquidate request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `keeper` | [string](#string) |  |  |
| `borrower` | [string](#string) |  |  |






<a name="nemo.jinx.v1beta1.MsgLiquidateResponse"></a>

### MsgLiquidateResponse
MsgLiquidateResponse defines the Msg/Liquidate response type.






<a name="nemo.jinx.v1beta1.MsgRepay"></a>

### MsgRepay
MsgRepay defines the Msg/Repay request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="nemo.jinx.v1beta1.MsgRepayResponse"></a>

### MsgRepayResponse
MsgRepayResponse defines the Msg/Repay response type.






<a name="nemo.jinx.v1beta1.MsgWithdraw"></a>

### MsgWithdraw
MsgWithdraw defines the Msg/Withdraw request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="nemo.jinx.v1beta1.MsgWithdrawResponse"></a>

### MsgWithdrawResponse
MsgWithdrawResponse defines the Msg/Withdraw response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.jinx.v1beta1.Msg"></a>

### Msg
Msg defines the jinx Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Deposit` | [MsgDeposit](#nemo.jinx.v1beta1.MsgDeposit) | [MsgDepositResponse](#nemo.jinx.v1beta1.MsgDepositResponse) | Deposit defines a method for depositing funds to jinx liquidity pool. | |
| `Withdraw` | [MsgWithdraw](#nemo.jinx.v1beta1.MsgWithdraw) | [MsgWithdrawResponse](#nemo.jinx.v1beta1.MsgWithdrawResponse) | Withdraw defines a method for withdrawing funds from jinx liquidity pool. | |
| `Borrow` | [MsgBorrow](#nemo.jinx.v1beta1.MsgBorrow) | [MsgBorrowResponse](#nemo.jinx.v1beta1.MsgBorrowResponse) | Borrow defines a method for borrowing funds from jinx liquidity pool. | |
| `Repay` | [MsgRepay](#nemo.jinx.v1beta1.MsgRepay) | [MsgRepayResponse](#nemo.jinx.v1beta1.MsgRepayResponse) | Repay defines a method for repaying funds borrowed from jinx liquidity pool. | |
| `Liquidate` | [MsgLiquidate](#nemo.jinx.v1beta1.MsgLiquidate) | [MsgLiquidateResponse](#nemo.jinx.v1beta1.MsgLiquidateResponse) | Liquidate defines a method for attempting to liquidate a borrower that is over their loan-to-value. | |

 <!-- end services -->



<a name="nemo/incentive/v1beta1/apy.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/incentive/v1beta1/apy.proto



<a name="nemo.incentive.v1beta1.Apy"></a>

### Apy
Apy contains the calculated APY for a given collateral type at a specific
instant in time.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `apy` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/incentive/v1beta1/claims.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/incentive/v1beta1/claims.proto



<a name="nemo.incentive.v1beta1.BaseClaim"></a>

### BaseClaim
BaseClaim is a claim with a single reward coin types


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [bytes](#bytes) |  |  |
| `reward` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="nemo.incentive.v1beta1.BaseMultiClaim"></a>

### BaseMultiClaim
BaseMultiClaim is a claim with multiple reward coin types


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [bytes](#bytes) |  |  |
| `reward` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="nemo.incentive.v1beta1.DelegatorClaim"></a>

### DelegatorClaim
DelegatorClaim stores delegation rewards that can be claimed by owner


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_claim` | [BaseMultiClaim](#nemo.incentive.v1beta1.BaseMultiClaim) |  |  |
| `reward_indexes` | [MultiRewardIndex](#nemo.incentive.v1beta1.MultiRewardIndex) | repeated |  |






<a name="nemo.incentive.v1beta1.EarnClaim"></a>

### EarnClaim
EarnClaim stores the earn rewards that can be claimed by owner


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_claim` | [BaseMultiClaim](#nemo.incentive.v1beta1.BaseMultiClaim) |  |  |
| `reward_indexes` | [MultiRewardIndex](#nemo.incentive.v1beta1.MultiRewardIndex) | repeated |  |






<a name="nemo.incentive.v1beta1.JinxLiquidityProviderClaim"></a>

### JinxLiquidityProviderClaim
JinxLiquidityProviderClaim stores the jinx liquidity provider rewards that can be claimed by owner


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_claim` | [BaseMultiClaim](#nemo.incentive.v1beta1.BaseMultiClaim) |  |  |
| `supply_reward_indexes` | [MultiRewardIndex](#nemo.incentive.v1beta1.MultiRewardIndex) | repeated |  |
| `borrow_reward_indexes` | [MultiRewardIndex](#nemo.incentive.v1beta1.MultiRewardIndex) | repeated |  |






<a name="nemo.incentive.v1beta1.MultiRewardIndex"></a>

### MultiRewardIndex
MultiRewardIndex stores reward accumulation information on multiple reward types


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `reward_indexes` | [RewardIndex](#nemo.incentive.v1beta1.RewardIndex) | repeated |  |






<a name="nemo.incentive.v1beta1.MultiRewardIndexesProto"></a>

### MultiRewardIndexesProto
MultiRewardIndexesProto defines a Protobuf wrapper around a MultiRewardIndexes slice


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `multi_reward_indexes` | [MultiRewardIndex](#nemo.incentive.v1beta1.MultiRewardIndex) | repeated |  |






<a name="nemo.incentive.v1beta1.RewardIndex"></a>

### RewardIndex
RewardIndex stores reward accumulation information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `reward_factor` | [bytes](#bytes) |  |  |






<a name="nemo.incentive.v1beta1.RewardIndexesProto"></a>

### RewardIndexesProto
RewardIndexesProto defines a Protobuf wrapper around a RewardIndexes slice


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `reward_indexes` | [RewardIndex](#nemo.incentive.v1beta1.RewardIndex) | repeated |  |






<a name="nemo.incentive.v1beta1.SavingsClaim"></a>

### SavingsClaim
SavingsClaim stores the savings rewards that can be claimed by owner


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_claim` | [BaseMultiClaim](#nemo.incentive.v1beta1.BaseMultiClaim) |  |  |
| `reward_indexes` | [MultiRewardIndex](#nemo.incentive.v1beta1.MultiRewardIndex) | repeated |  |






<a name="nemo.incentive.v1beta1.SwapClaim"></a>

### SwapClaim
SwapClaim stores the swap rewards that can be claimed by owner


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_claim` | [BaseMultiClaim](#nemo.incentive.v1beta1.BaseMultiClaim) |  |  |
| `reward_indexes` | [MultiRewardIndex](#nemo.incentive.v1beta1.MultiRewardIndex) | repeated |  |






<a name="nemo.incentive.v1beta1.MUSDMintingClaim"></a>

### MUSDMintingClaim
MUSDMintingClaim is for MUSD minting rewards


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_claim` | [BaseClaim](#nemo.incentive.v1beta1.BaseClaim) |  |  |
| `reward_indexes` | [RewardIndex](#nemo.incentive.v1beta1.RewardIndex) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/incentive/v1beta1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/incentive/v1beta1/params.proto



<a name="nemo.incentive.v1beta1.MultiRewardPeriod"></a>

### MultiRewardPeriod
MultiRewardPeriod supports multiple reward types


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `active` | [bool](#bool) |  |  |
| `collateral_type` | [string](#string) |  |  |
| `start` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `end` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `rewards_per_second` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="nemo.incentive.v1beta1.Multiplier"></a>

### Multiplier
Multiplier amount the claim rewards get increased by, along with how long the claim rewards are locked


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |
| `months_lockup` | [int64](#int64) |  |  |
| `factor` | [bytes](#bytes) |  |  |






<a name="nemo.incentive.v1beta1.MultipliersPerDenom"></a>

### MultipliersPerDenom
MultipliersPerDenom is a map of denoms to a set of multipliers


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `multipliers` | [Multiplier](#nemo.incentive.v1beta1.Multiplier) | repeated |  |






<a name="nemo.incentive.v1beta1.Params"></a>

### Params
Params


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `musd_minting_reward_periods` | [RewardPeriod](#nemo.incentive.v1beta1.RewardPeriod) | repeated |  |
| `jinx_supply_reward_periods` | [MultiRewardPeriod](#nemo.incentive.v1beta1.MultiRewardPeriod) | repeated |  |
| `jinx_borrow_reward_periods` | [MultiRewardPeriod](#nemo.incentive.v1beta1.MultiRewardPeriod) | repeated |  |
| `delegator_reward_periods` | [MultiRewardPeriod](#nemo.incentive.v1beta1.MultiRewardPeriod) | repeated |  |
| `swap_reward_periods` | [MultiRewardPeriod](#nemo.incentive.v1beta1.MultiRewardPeriod) | repeated |  |
| `claim_multipliers` | [MultipliersPerDenom](#nemo.incentive.v1beta1.MultipliersPerDenom) | repeated |  |
| `claim_end` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `savings_reward_periods` | [MultiRewardPeriod](#nemo.incentive.v1beta1.MultiRewardPeriod) | repeated |  |
| `earn_reward_periods` | [MultiRewardPeriod](#nemo.incentive.v1beta1.MultiRewardPeriod) | repeated |  |






<a name="nemo.incentive.v1beta1.RewardPeriod"></a>

### RewardPeriod
RewardPeriod stores the state of an ongoing reward


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `active` | [bool](#bool) |  |  |
| `collateral_type` | [string](#string) |  |  |
| `start` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `end` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `rewards_per_second` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/incentive/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/incentive/v1beta1/genesis.proto



<a name="nemo.incentive.v1beta1.AccumulationTime"></a>

### AccumulationTime
AccumulationTime stores the previous reward distribution time and its corresponding collateral type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collateral_type` | [string](#string) |  |  |
| `previous_accumulation_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="nemo.incentive.v1beta1.GenesisRewardState"></a>

### GenesisRewardState
GenesisRewardState groups together the global state for a particular reward so it can be exported in genesis.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `accumulation_times` | [AccumulationTime](#nemo.incentive.v1beta1.AccumulationTime) | repeated |  |
| `multi_reward_indexes` | [MultiRewardIndex](#nemo.incentive.v1beta1.MultiRewardIndex) | repeated |  |






<a name="nemo.incentive.v1beta1.GenesisState"></a>

### GenesisState
GenesisState is the state that must be provided at genesis.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#nemo.incentive.v1beta1.Params) |  |  |
| `musd_reward_state` | [GenesisRewardState](#nemo.incentive.v1beta1.GenesisRewardState) |  |  |
| `jinx_supply_reward_state` | [GenesisRewardState](#nemo.incentive.v1beta1.GenesisRewardState) |  |  |
| `jinx_borrow_reward_state` | [GenesisRewardState](#nemo.incentive.v1beta1.GenesisRewardState) |  |  |
| `delegator_reward_state` | [GenesisRewardState](#nemo.incentive.v1beta1.GenesisRewardState) |  |  |
| `swap_reward_state` | [GenesisRewardState](#nemo.incentive.v1beta1.GenesisRewardState) |  |  |
| `musd_minting_claims` | [MUSDMintingClaim](#nemo.incentive.v1beta1.MUSDMintingClaim) | repeated |  |
| `jinx_liquidity_provider_claims` | [JinxLiquidityProviderClaim](#nemo.incentive.v1beta1.JinxLiquidityProviderClaim) | repeated |  |
| `delegator_claims` | [DelegatorClaim](#nemo.incentive.v1beta1.DelegatorClaim) | repeated |  |
| `swap_claims` | [SwapClaim](#nemo.incentive.v1beta1.SwapClaim) | repeated |  |
| `savings_reward_state` | [GenesisRewardState](#nemo.incentive.v1beta1.GenesisRewardState) |  |  |
| `savings_claims` | [SavingsClaim](#nemo.incentive.v1beta1.SavingsClaim) | repeated |  |
| `earn_reward_state` | [GenesisRewardState](#nemo.incentive.v1beta1.GenesisRewardState) |  |  |
| `earn_claims` | [EarnClaim](#nemo.incentive.v1beta1.EarnClaim) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/incentive/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/incentive/v1beta1/query.proto



<a name="nemo.incentive.v1beta1.QueryApyRequest"></a>

### QueryApyRequest
QueryApysRequest is the request type for the Query/Apys RPC method.






<a name="nemo.incentive.v1beta1.QueryApyResponse"></a>

### QueryApyResponse
QueryApysResponse is the response type for the Query/Apys RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `earn` | [Apy](#nemo.incentive.v1beta1.Apy) | repeated |  |






<a name="nemo.incentive.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="nemo.incentive.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#nemo.incentive.v1beta1.Params) |  |  |






<a name="nemo.incentive.v1beta1.QueryRewardFactorsRequest"></a>

### QueryRewardFactorsRequest
QueryRewardFactorsRequest is the request type for the Query/RewardFactors RPC method.






<a name="nemo.incentive.v1beta1.QueryRewardFactorsResponse"></a>

### QueryRewardFactorsResponse
QueryRewardFactorsResponse is the response type for the Query/RewardFactors RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `musd_minting_reward_factors` | [RewardIndex](#nemo.incentive.v1beta1.RewardIndex) | repeated |  |
| `jinx_supply_reward_factors` | [MultiRewardIndex](#nemo.incentive.v1beta1.MultiRewardIndex) | repeated |  |
| `jinx_borrow_reward_factors` | [MultiRewardIndex](#nemo.incentive.v1beta1.MultiRewardIndex) | repeated |  |
| `delegator_reward_factors` | [MultiRewardIndex](#nemo.incentive.v1beta1.MultiRewardIndex) | repeated |  |
| `swap_reward_factors` | [MultiRewardIndex](#nemo.incentive.v1beta1.MultiRewardIndex) | repeated |  |
| `savings_reward_factors` | [MultiRewardIndex](#nemo.incentive.v1beta1.MultiRewardIndex) | repeated |  |
| `earn_reward_factors` | [MultiRewardIndex](#nemo.incentive.v1beta1.MultiRewardIndex) | repeated |  |






<a name="nemo.incentive.v1beta1.QueryRewardsRequest"></a>

### QueryRewardsRequest
QueryRewardsRequest is the request type for the Query/Rewards RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  | owner is the address of the user to query rewards for. |
| `reward_type` | [string](#string) |  | reward_type is the type of reward to query rewards for, e.g. jinx, earn, swap. |
| `unsynchronized` | [bool](#bool) |  | unsynchronized is a flag to query rewards that are not simulated for reward synchronized for the current block. |






<a name="nemo.incentive.v1beta1.QueryRewardsResponse"></a>

### QueryRewardsResponse
QueryRewardsResponse is the response type for the Query/Rewards RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `musd_minting_claims` | [MUSDMintingClaim](#nemo.incentive.v1beta1.MUSDMintingClaim) | repeated |  |
| `jinx_liquidity_provider_claims` | [JinxLiquidityProviderClaim](#nemo.incentive.v1beta1.JinxLiquidityProviderClaim) | repeated |  |
| `delegator_claims` | [DelegatorClaim](#nemo.incentive.v1beta1.DelegatorClaim) | repeated |  |
| `swap_claims` | [SwapClaim](#nemo.incentive.v1beta1.SwapClaim) | repeated |  |
| `savings_claims` | [SavingsClaim](#nemo.incentive.v1beta1.SavingsClaim) | repeated |  |
| `earn_claims` | [EarnClaim](#nemo.incentive.v1beta1.EarnClaim) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.incentive.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for incentive module.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#nemo.incentive.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#nemo.incentive.v1beta1.QueryParamsResponse) | Params queries module params. | GET|/nemo/incentive/v1beta1/params|
| `Rewards` | [QueryRewardsRequest](#nemo.incentive.v1beta1.QueryRewardsRequest) | [QueryRewardsResponse](#nemo.incentive.v1beta1.QueryRewardsResponse) | Rewards queries reward information for a given user. | GET|/nemo/incentive/v1beta1/rewards|
| `RewardFactors` | [QueryRewardFactorsRequest](#nemo.incentive.v1beta1.QueryRewardFactorsRequest) | [QueryRewardFactorsResponse](#nemo.incentive.v1beta1.QueryRewardFactorsResponse) | Rewards queries the reward factors. | GET|/nemo/incentive/v1beta1/reward_factors|
| `Apy` | [QueryApyRequest](#nemo.incentive.v1beta1.QueryApyRequest) | [QueryApyResponse](#nemo.incentive.v1beta1.QueryApyResponse) | Apy queries incentive reward apy for a reward. | GET|/nemo/incentive/v1beta1/apy|

 <!-- end services -->



<a name="nemo/incentive/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/incentive/v1beta1/tx.proto



<a name="nemo.incentive.v1beta1.MsgClaimDelegatorReward"></a>

### MsgClaimDelegatorReward
MsgClaimDelegatorReward message type used to claim delegator rewards


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `denoms_to_claim` | [Selection](#nemo.incentive.v1beta1.Selection) | repeated |  |






<a name="nemo.incentive.v1beta1.MsgClaimDelegatorRewardResponse"></a>

### MsgClaimDelegatorRewardResponse
MsgClaimDelegatorRewardResponse defines the Msg/ClaimDelegatorReward response type.






<a name="nemo.incentive.v1beta1.MsgClaimEarnReward"></a>

### MsgClaimEarnReward
MsgClaimEarnReward message type used to claim earn rewards


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `denoms_to_claim` | [Selection](#nemo.incentive.v1beta1.Selection) | repeated |  |






<a name="nemo.incentive.v1beta1.MsgClaimEarnRewardResponse"></a>

### MsgClaimEarnRewardResponse
MsgClaimEarnRewardResponse defines the Msg/ClaimEarnReward response type.






<a name="nemo.incentive.v1beta1.MsgClaimJinxReward"></a>

### MsgClaimJinxReward
MsgClaimJinxReward message type used to claim Jinx liquidity provider rewards


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `denoms_to_claim` | [Selection](#nemo.incentive.v1beta1.Selection) | repeated |  |






<a name="nemo.incentive.v1beta1.MsgClaimJinxRewardResponse"></a>

### MsgClaimJinxRewardResponse
MsgClaimJinxRewardResponse defines the Msg/ClaimJinxReward response type.






<a name="nemo.incentive.v1beta1.MsgClaimSavingsReward"></a>

### MsgClaimSavingsReward
MsgClaimSavingsReward message type used to claim savings rewards


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `denoms_to_claim` | [Selection](#nemo.incentive.v1beta1.Selection) | repeated |  |






<a name="nemo.incentive.v1beta1.MsgClaimSavingsRewardResponse"></a>

### MsgClaimSavingsRewardResponse
MsgClaimSavingsRewardResponse defines the Msg/ClaimSavingsReward response type.






<a name="nemo.incentive.v1beta1.MsgClaimSwapReward"></a>

### MsgClaimSwapReward
MsgClaimSwapReward message type used to claim delegator rewards


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `denoms_to_claim` | [Selection](#nemo.incentive.v1beta1.Selection) | repeated |  |






<a name="nemo.incentive.v1beta1.MsgClaimSwapRewardResponse"></a>

### MsgClaimSwapRewardResponse
MsgClaimSwapRewardResponse defines the Msg/ClaimSwapReward response type.






<a name="nemo.incentive.v1beta1.MsgClaimMUSDMintingReward"></a>

### MsgClaimMUSDMintingReward
MsgClaimMUSDMintingReward message type used to claim MUSD minting rewards


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `multiplier_name` | [string](#string) |  |  |






<a name="nemo.incentive.v1beta1.MsgClaimMUSDMintingRewardResponse"></a>

### MsgClaimMUSDMintingRewardResponse
MsgClaimMUSDMintingRewardResponse defines the Msg/ClaimMUSDMintingReward response type.






<a name="nemo.incentive.v1beta1.Selection"></a>

### Selection
Selection is a pair of denom and multiplier name. It holds the choice of multiplier a user makes when they claim a
denom.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `multiplier_name` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.incentive.v1beta1.Msg"></a>

### Msg
Msg defines the incentive Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ClaimMUSDMintingReward` | [MsgClaimMUSDMintingReward](#nemo.incentive.v1beta1.MsgClaimMUSDMintingReward) | [MsgClaimMUSDMintingRewardResponse](#nemo.incentive.v1beta1.MsgClaimMUSDMintingRewardResponse) | ClaimMUSDMintingReward is a message type used to claim MUSD minting rewards | |
| `ClaimJinxReward` | [MsgClaimJinxReward](#nemo.incentive.v1beta1.MsgClaimJinxReward) | [MsgClaimJinxRewardResponse](#nemo.incentive.v1beta1.MsgClaimJinxRewardResponse) | ClaimJinxReward is a message type used to claim Jinx liquidity provider rewards | |
| `ClaimDelegatorReward` | [MsgClaimDelegatorReward](#nemo.incentive.v1beta1.MsgClaimDelegatorReward) | [MsgClaimDelegatorRewardResponse](#nemo.incentive.v1beta1.MsgClaimDelegatorRewardResponse) | ClaimDelegatorReward is a message type used to claim delegator rewards | |
| `ClaimSwapReward` | [MsgClaimSwapReward](#nemo.incentive.v1beta1.MsgClaimSwapReward) | [MsgClaimSwapRewardResponse](#nemo.incentive.v1beta1.MsgClaimSwapRewardResponse) | ClaimSwapReward is a message type used to claim swap rewards | |
| `ClaimSavingsReward` | [MsgClaimSavingsReward](#nemo.incentive.v1beta1.MsgClaimSavingsReward) | [MsgClaimSavingsRewardResponse](#nemo.incentive.v1beta1.MsgClaimSavingsRewardResponse) | ClaimSavingsReward is a message type used to claim savings rewards | |
| `ClaimEarnReward` | [MsgClaimEarnReward](#nemo.incentive.v1beta1.MsgClaimEarnReward) | [MsgClaimEarnRewardResponse](#nemo.incentive.v1beta1.MsgClaimEarnRewardResponse) | ClaimEarnReward is a message type used to claim earn rewards | |

 <!-- end services -->



<a name="nemo/issuance/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/issuance/v1beta1/genesis.proto



<a name="nemo.issuance.v1beta1.Asset"></a>

### Asset
Asset type for assets in the issuance module


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |
| `denom` | [string](#string) |  |  |
| `blocked_addresses` | [string](#string) | repeated |  |
| `paused` | [bool](#bool) |  |  |
| `blockable` | [bool](#bool) |  |  |
| `rate_limit` | [RateLimit](#nemo.issuance.v1beta1.RateLimit) |  |  |






<a name="nemo.issuance.v1beta1.AssetSupply"></a>

### AssetSupply
AssetSupply contains information about an asset's rate-limited supply (the
total supply of the asset is tracked in the top-level supply module)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `current_supply` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `time_elapsed` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |






<a name="nemo.issuance.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the issuance module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#nemo.issuance.v1beta1.Params) |  | params defines all the paramaters of the module. |
| `supplies` | [AssetSupply](#nemo.issuance.v1beta1.AssetSupply) | repeated |  |






<a name="nemo.issuance.v1beta1.Params"></a>

### Params
Params defines the parameters for the issuance module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `assets` | [Asset](#nemo.issuance.v1beta1.Asset) | repeated |  |






<a name="nemo.issuance.v1beta1.RateLimit"></a>

### RateLimit
RateLimit parameters for rate-limiting the supply of an issued asset


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `active` | [bool](#bool) |  |  |
| `limit` | [bytes](#bytes) |  |  |
| `time_period` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/issuance/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/issuance/v1beta1/query.proto



<a name="nemo.issuance.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest defines the request type for querying x/issuance parameters.






<a name="nemo.issuance.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse defines the response type for querying x/issuance parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#nemo.issuance.v1beta1.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.issuance.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for issuance module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#nemo.issuance.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#nemo.issuance.v1beta1.QueryParamsResponse) | Params queries all parameters of the issuance module. | GET|/nemo/issuance/v1beta1/params|

 <!-- end services -->



<a name="nemo/issuance/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/issuance/v1beta1/tx.proto



<a name="nemo.issuance.v1beta1.MsgBlockAddress"></a>

### MsgBlockAddress
MsgBlockAddress represents a message used by the issuer to block an address from holding or transferring tokens


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `denom` | [string](#string) |  |  |
| `blocked_address` | [string](#string) |  |  |






<a name="nemo.issuance.v1beta1.MsgBlockAddressResponse"></a>

### MsgBlockAddressResponse
MsgBlockAddressResponse defines the Msg/BlockAddress response type.






<a name="nemo.issuance.v1beta1.MsgIssueTokens"></a>

### MsgIssueTokens
MsgIssueTokens represents a message used by the issuer to issue new tokens


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `tokens` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `receiver` | [string](#string) |  |  |






<a name="nemo.issuance.v1beta1.MsgIssueTokensResponse"></a>

### MsgIssueTokensResponse
MsgIssueTokensResponse defines the Msg/IssueTokens response type.






<a name="nemo.issuance.v1beta1.MsgRedeemTokens"></a>

### MsgRedeemTokens
MsgRedeemTokens represents a message used by the issuer to redeem (burn) tokens


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `tokens` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="nemo.issuance.v1beta1.MsgRedeemTokensResponse"></a>

### MsgRedeemTokensResponse
MsgRedeemTokensResponse defines the Msg/RedeemTokens response type.






<a name="nemo.issuance.v1beta1.MsgSetPauseStatus"></a>

### MsgSetPauseStatus
MsgSetPauseStatus message type used by the issuer to pause or unpause status


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `denom` | [string](#string) |  |  |
| `status` | [bool](#bool) |  |  |






<a name="nemo.issuance.v1beta1.MsgSetPauseStatusResponse"></a>

### MsgSetPauseStatusResponse
MsgSetPauseStatusResponse defines the Msg/SetPauseStatus response type.






<a name="nemo.issuance.v1beta1.MsgUnblockAddress"></a>

### MsgUnblockAddress
MsgUnblockAddress message type used by the issuer to unblock an address from holding or transferring tokens


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `denom` | [string](#string) |  |  |
| `blocked_address` | [string](#string) |  |  |






<a name="nemo.issuance.v1beta1.MsgUnblockAddressResponse"></a>

### MsgUnblockAddressResponse
MsgUnblockAddressResponse defines the Msg/UnblockAddress response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.issuance.v1beta1.Msg"></a>

### Msg
Msg defines the issuance Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `IssueTokens` | [MsgIssueTokens](#nemo.issuance.v1beta1.MsgIssueTokens) | [MsgIssueTokensResponse](#nemo.issuance.v1beta1.MsgIssueTokensResponse) | IssueTokens message type used by the issuer to issue new tokens | |
| `RedeemTokens` | [MsgRedeemTokens](#nemo.issuance.v1beta1.MsgRedeemTokens) | [MsgRedeemTokensResponse](#nemo.issuance.v1beta1.MsgRedeemTokensResponse) | RedeemTokens message type used by the issuer to redeem (burn) tokens | |
| `BlockAddress` | [MsgBlockAddress](#nemo.issuance.v1beta1.MsgBlockAddress) | [MsgBlockAddressResponse](#nemo.issuance.v1beta1.MsgBlockAddressResponse) | BlockAddress message type used by the issuer to block an address from holding or transferring tokens | |
| `UnblockAddress` | [MsgUnblockAddress](#nemo.issuance.v1beta1.MsgUnblockAddress) | [MsgUnblockAddressResponse](#nemo.issuance.v1beta1.MsgUnblockAddressResponse) | UnblockAddress message type used by the issuer to unblock an address from holding or transferring tokens | |
| `SetPauseStatus` | [MsgSetPauseStatus](#nemo.issuance.v1beta1.MsgSetPauseStatus) | [MsgSetPauseStatusResponse](#nemo.issuance.v1beta1.MsgSetPauseStatusResponse) | SetPauseStatus message type used to pause or unpause status | |

 <!-- end services -->



<a name="nemo/liquid/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/liquid/v1beta1/query.proto



<a name="nemo.liquid.v1beta1.QueryDelegatedBalanceRequest"></a>

### QueryDelegatedBalanceRequest
QueryDelegatedBalanceRequest defines the request type for Query/DelegatedBalance method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator` | [string](#string) |  | delegator is the address of the account to query |






<a name="nemo.liquid.v1beta1.QueryDelegatedBalanceResponse"></a>

### QueryDelegatedBalanceResponse
DelegatedBalanceResponse defines the response type for the Query/DelegatedBalance method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `vested` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | vested is the amount of all delegated coins that have vested (ie not locked) |
| `vesting` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | vesting is the amount of all delegated coins that are still vesting (ie locked) |






<a name="nemo.liquid.v1beta1.QueryTotalSupplyRequest"></a>

### QueryTotalSupplyRequest
QueryTotalSupplyRequest defines the request type for Query/TotalSupply method.






<a name="nemo.liquid.v1beta1.QueryTotalSupplyResponse"></a>

### QueryTotalSupplyResponse
TotalSupplyResponse defines the response type for the Query/TotalSupply method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `height` | [int64](#int64) |  | Height is the block height at which these totals apply |
| `result` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | Result is a list of coins supplied to liquid |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.liquid.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for liquid module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `DelegatedBalance` | [QueryDelegatedBalanceRequest](#nemo.liquid.v1beta1.QueryDelegatedBalanceRequest) | [QueryDelegatedBalanceResponse](#nemo.liquid.v1beta1.QueryDelegatedBalanceResponse) | DelegatedBalance returns an account's vesting and vested coins currently delegated to validators. It ignores coins in unbonding delegations. | GET|/nemo/liquid/v1beta1/delegated_balance/{delegator}|
| `TotalSupply` | [QueryTotalSupplyRequest](#nemo.liquid.v1beta1.QueryTotalSupplyRequest) | [QueryTotalSupplyResponse](#nemo.liquid.v1beta1.QueryTotalSupplyResponse) | TotalSupply returns the total sum of all coins currently locked into the liquid module. | GET|/nemo/liquid/v1beta1/total_supply|

 <!-- end services -->



<a name="nemo/liquid/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/liquid/v1beta1/tx.proto



<a name="nemo.liquid.v1beta1.MsgBurnDerivative"></a>

### MsgBurnDerivative
MsgBurnDerivative defines the Msg/BurnDerivative request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  | sender is the owner of the derivatives to be converted |
| `validator` | [string](#string) |  | validator is the validator of the derivatives to be converted |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | amount is the quantity of derivatives to be converted |






<a name="nemo.liquid.v1beta1.MsgBurnDerivativeResponse"></a>

### MsgBurnDerivativeResponse
MsgBurnDerivativeResponse defines the Msg/BurnDerivative response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `received` | [string](#string) |  | received is the number of delegation shares sent to the sender |






<a name="nemo.liquid.v1beta1.MsgMintDerivative"></a>

### MsgMintDerivative
MsgMintDerivative defines the Msg/MintDerivative request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  | sender is the owner of the delegation to be converted |
| `validator` | [string](#string) |  | validator is the validator of the delegation to be converted |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | amount is the quantity of staked assets to be converted |






<a name="nemo.liquid.v1beta1.MsgMintDerivativeResponse"></a>

### MsgMintDerivativeResponse
MsgMintDerivativeResponse defines the Msg/MintDerivative response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `received` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | received is the amount of staking derivative minted and sent to the sender |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.liquid.v1beta1.Msg"></a>

### Msg
Msg defines the liquid Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `MintDerivative` | [MsgMintDerivative](#nemo.liquid.v1beta1.MsgMintDerivative) | [MsgMintDerivativeResponse](#nemo.liquid.v1beta1.MsgMintDerivativeResponse) | MintDerivative defines a method for converting a delegation into staking deriviatives. | |
| `BurnDerivative` | [MsgBurnDerivative](#nemo.liquid.v1beta1.MsgBurnDerivative) | [MsgBurnDerivativeResponse](#nemo.liquid.v1beta1.MsgBurnDerivativeResponse) | BurnDerivative defines a method for converting staking deriviatives into a delegation. | |

 <!-- end services -->



<a name="nemo/nemodist/v1beta1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/nemodist/v1beta1/params.proto



<a name="nemo.nemodist.v1beta1.CoreReward"></a>

### CoreReward
CoreReward defines the reward weights for core infrastructure providers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [bytes](#bytes) |  |  |
| `weight` | [string](#string) |  |  |






<a name="nemo.nemodist.v1beta1.InfrastructureParams"></a>

### InfrastructureParams
InfrastructureParams define the parameters for infrastructure rewards.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `infrastructure_periods` | [Period](#nemo.nemodist.v1beta1.Period) | repeated |  |
| `core_rewards` | [CoreReward](#nemo.nemodist.v1beta1.CoreReward) | repeated |  |
| `partner_rewards` | [PartnerReward](#nemo.nemodist.v1beta1.PartnerReward) | repeated |  |






<a name="nemo.nemodist.v1beta1.Params"></a>

### Params
Params governance parameters for nemodist module


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `active` | [bool](#bool) |  |  |
| `periods` | [Period](#nemo.nemodist.v1beta1.Period) | repeated |  |
| `infrastructure_params` | [InfrastructureParams](#nemo.nemodist.v1beta1.InfrastructureParams) |  |  |






<a name="nemo.nemodist.v1beta1.PartnerReward"></a>

### PartnerReward
PartnerRewards defines the reward schedule for partner infrastructure providers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [bytes](#bytes) |  |  |
| `rewards_per_second` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="nemo.nemodist.v1beta1.Period"></a>

### Period
Period stores the specified start and end dates, and the inflation, expressed as a decimal
representing the yearly APR of NEMO tokens that will be minted during that period


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `start` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | example "2020-03-01T15:20:00Z" |
| `end` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | example "2020-06-01T15:20:00Z" |
| `inflation` | [bytes](#bytes) |  | example "1.000000003022265980" - 10% inflation |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/nemodist/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/nemodist/v1beta1/genesis.proto



<a name="nemo.nemodist.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the nemodist module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#nemo.nemodist.v1beta1.Params) |  |  |
| `previous_block_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/nemodist/v1beta1/proposal.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/nemodist/v1beta1/proposal.proto



<a name="nemo.nemodist.v1beta1.CommunityPoolMultiSpendProposal"></a>

### CommunityPoolMultiSpendProposal
CommunityPoolMultiSpendProposal spends from the community pool by sending to one or more
addresses


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `recipient_list` | [MultiSpendRecipient](#nemo.nemodist.v1beta1.MultiSpendRecipient) | repeated |  |






<a name="nemo.nemodist.v1beta1.CommunityPoolMultiSpendProposalJSON"></a>

### CommunityPoolMultiSpendProposalJSON
CommunityPoolMultiSpendProposalJSON defines a CommunityPoolMultiSpendProposal with a deposit


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `recipient_list` | [MultiSpendRecipient](#nemo.nemodist.v1beta1.MultiSpendRecipient) | repeated |  |
| `deposit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="nemo.nemodist.v1beta1.MultiSpendRecipient"></a>

### MultiSpendRecipient
MultiSpendRecipient defines a recipient and the amount of coins they are receiving


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/nemodist/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/nemodist/v1beta1/query.proto



<a name="nemo.nemodist.v1beta1.QueryBalanceRequest"></a>

### QueryBalanceRequest
QueryBalanceRequest defines the request type for querying x/nemodist balance.






<a name="nemo.nemodist.v1beta1.QueryBalanceResponse"></a>

### QueryBalanceResponse
QueryBalanceResponse defines the response type for querying x/nemodist balance.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="nemo.nemodist.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest defines the request type for querying x/nemodist parameters.






<a name="nemo.nemodist.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse defines the response type for querying x/nemodist parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#nemo.nemodist.v1beta1.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.nemodist.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#nemo.nemodist.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#nemo.nemodist.v1beta1.QueryParamsResponse) | Params queries the parameters of x/nemodist module. | GET|/nemo/nemodist/v1beta1/parameters|
| `Balance` | [QueryBalanceRequest](#nemo.nemodist.v1beta1.QueryBalanceRequest) | [QueryBalanceResponse](#nemo.nemodist.v1beta1.QueryBalanceResponse) | Balance queries the balance of all coins of x/nemodist module. | GET|/nemo/nemodist/v1beta1/balance|

 <!-- end services -->



<a name="nemo/pricefeed/v1beta1/store.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/pricefeed/v1beta1/store.proto



<a name="nemo.pricefeed.v1beta1.CurrentPrice"></a>

### CurrentPrice
CurrentPrice defines a current price for a particular market in the pricefeed
module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |






<a name="nemo.pricefeed.v1beta1.Market"></a>

### Market
Market defines an asset in the pricefeed.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `base_asset` | [string](#string) |  |  |
| `quote_asset` | [string](#string) |  |  |
| `oracles` | [bytes](#bytes) | repeated |  |
| `active` | [bool](#bool) |  |  |






<a name="nemo.pricefeed.v1beta1.Params"></a>

### Params
Params defines the parameters for the pricefeed module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `markets` | [Market](#nemo.pricefeed.v1beta1.Market) | repeated |  |






<a name="nemo.pricefeed.v1beta1.PostedPrice"></a>

### PostedPrice
PostedPrice defines a price for market posted by a specific oracle.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `oracle_address` | [bytes](#bytes) |  |  |
| `price` | [string](#string) |  |  |
| `expiry` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/pricefeed/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/pricefeed/v1beta1/genesis.proto



<a name="nemo.pricefeed.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the pricefeed module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#nemo.pricefeed.v1beta1.Params) |  | params defines all the paramaters of the module. |
| `posted_prices` | [PostedPrice](#nemo.pricefeed.v1beta1.PostedPrice) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/pricefeed/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/pricefeed/v1beta1/query.proto



<a name="nemo.pricefeed.v1beta1.CurrentPriceResponse"></a>

### CurrentPriceResponse
CurrentPriceResponse defines a current price for a particular market in the pricefeed
module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |






<a name="nemo.pricefeed.v1beta1.MarketResponse"></a>

### MarketResponse
MarketResponse defines an asset in the pricefeed.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `base_asset` | [string](#string) |  |  |
| `quote_asset` | [string](#string) |  |  |
| `oracles` | [string](#string) | repeated |  |
| `active` | [bool](#bool) |  |  |






<a name="nemo.pricefeed.v1beta1.PostedPriceResponse"></a>

### PostedPriceResponse
PostedPriceResponse defines a price for market posted by a specific oracle.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |
| `oracle_address` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |
| `expiry` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="nemo.pricefeed.v1beta1.QueryMarketsRequest"></a>

### QueryMarketsRequest
QueryMarketsRequest is the request type for the Query/Markets RPC method.






<a name="nemo.pricefeed.v1beta1.QueryMarketsResponse"></a>

### QueryMarketsResponse
QueryMarketsResponse is the response type for the Query/Markets RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `markets` | [MarketResponse](#nemo.pricefeed.v1beta1.MarketResponse) | repeated | List of markets |






<a name="nemo.pricefeed.v1beta1.QueryOraclesRequest"></a>

### QueryOraclesRequest
QueryOraclesRequest is the request type for the Query/Oracles RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |






<a name="nemo.pricefeed.v1beta1.QueryOraclesResponse"></a>

### QueryOraclesResponse
QueryOraclesResponse is the response type for the Query/Oracles RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `oracles` | [string](#string) | repeated | List of oracle addresses |






<a name="nemo.pricefeed.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest defines the request type for querying x/pricefeed
parameters.






<a name="nemo.pricefeed.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse defines the response type for querying x/pricefeed
parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#nemo.pricefeed.v1beta1.Params) |  |  |






<a name="nemo.pricefeed.v1beta1.QueryPriceRequest"></a>

### QueryPriceRequest
QueryPriceRequest is the request type for the Query/PriceRequest RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |






<a name="nemo.pricefeed.v1beta1.QueryPriceResponse"></a>

### QueryPriceResponse
QueryPriceResponse is the response type for the Query/Prices RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `price` | [CurrentPriceResponse](#nemo.pricefeed.v1beta1.CurrentPriceResponse) |  |  |






<a name="nemo.pricefeed.v1beta1.QueryPricesRequest"></a>

### QueryPricesRequest
QueryPricesRequest is the request type for the Query/Prices RPC method.






<a name="nemo.pricefeed.v1beta1.QueryPricesResponse"></a>

### QueryPricesResponse
QueryPricesResponse is the response type for the Query/Prices RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `prices` | [CurrentPriceResponse](#nemo.pricefeed.v1beta1.CurrentPriceResponse) | repeated |  |






<a name="nemo.pricefeed.v1beta1.QueryRawPricesRequest"></a>

### QueryRawPricesRequest
QueryRawPricesRequest is the request type for the Query/RawPrices RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `market_id` | [string](#string) |  |  |






<a name="nemo.pricefeed.v1beta1.QueryRawPricesResponse"></a>

### QueryRawPricesResponse
QueryRawPricesResponse is the response type for the Query/RawPrices RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `raw_prices` | [PostedPriceResponse](#nemo.pricefeed.v1beta1.PostedPriceResponse) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.pricefeed.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for pricefeed module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#nemo.pricefeed.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#nemo.pricefeed.v1beta1.QueryParamsResponse) | Params queries all parameters of the pricefeed module. | GET|/nemo/pricefeed/v1beta1/params|
| `Price` | [QueryPriceRequest](#nemo.pricefeed.v1beta1.QueryPriceRequest) | [QueryPriceResponse](#nemo.pricefeed.v1beta1.QueryPriceResponse) | Price queries price details based on a market | GET|/nemo/pricefeed/v1beta1/prices/{market_id}|
| `Prices` | [QueryPricesRequest](#nemo.pricefeed.v1beta1.QueryPricesRequest) | [QueryPricesResponse](#nemo.pricefeed.v1beta1.QueryPricesResponse) | Prices queries all prices | GET|/nemo/pricefeed/v1beta1/prices|
| `RawPrices` | [QueryRawPricesRequest](#nemo.pricefeed.v1beta1.QueryRawPricesRequest) | [QueryRawPricesResponse](#nemo.pricefeed.v1beta1.QueryRawPricesResponse) | RawPrices queries all raw prices based on a market | GET|/nemo/pricefeed/v1beta1/rawprices/{market_id}|
| `Oracles` | [QueryOraclesRequest](#nemo.pricefeed.v1beta1.QueryOraclesRequest) | [QueryOraclesResponse](#nemo.pricefeed.v1beta1.QueryOraclesResponse) | Oracles queries all oracles based on a market | GET|/nemo/pricefeed/v1beta1/oracles/{market_id}|
| `Markets` | [QueryMarketsRequest](#nemo.pricefeed.v1beta1.QueryMarketsRequest) | [QueryMarketsResponse](#nemo.pricefeed.v1beta1.QueryMarketsResponse) | Markets queries all markets | GET|/nemo/pricefeed/v1beta1/markets|

 <!-- end services -->



<a name="nemo/pricefeed/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/pricefeed/v1beta1/tx.proto



<a name="nemo.pricefeed.v1beta1.MsgPostPrice"></a>

### MsgPostPrice
MsgPostPrice represents a method for creating a new post price


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from` | [string](#string) |  | address of client |
| `market_id` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |
| `expiry` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="nemo.pricefeed.v1beta1.MsgPostPriceResponse"></a>

### MsgPostPriceResponse
MsgPostPriceResponse defines the Msg/PostPrice response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.pricefeed.v1beta1.Msg"></a>

### Msg
Msg defines the pricefeed Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `PostPrice` | [MsgPostPrice](#nemo.pricefeed.v1beta1.MsgPostPrice) | [MsgPostPriceResponse](#nemo.pricefeed.v1beta1.MsgPostPriceResponse) | PostPrice defines a method for creating a new post price | |

 <!-- end services -->



<a name="nemo/router/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/router/v1beta1/tx.proto



<a name="nemo.router.v1beta1.MsgDelegateMintDeposit"></a>

### MsgDelegateMintDeposit
MsgDelegateMintDeposit delegates tokens to a validator, then converts them into staking derivatives,
then deposits to an earn vault.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  | depositor represents the owner of the tokens to delegate |
| `validator` | [string](#string) |  | validator is the address of the validator to delegate to |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | amount is the tokens to delegate |






<a name="nemo.router.v1beta1.MsgDelegateMintDepositResponse"></a>

### MsgDelegateMintDepositResponse
MsgDelegateMintDepositResponse defines the Msg/MsgDelegateMintDeposit response type.






<a name="nemo.router.v1beta1.MsgMintDeposit"></a>

### MsgMintDeposit
MsgMintDeposit converts a delegation into staking derivatives and deposits it all into an earn vault.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  | depositor represents the owner of the delegation to convert |
| `validator` | [string](#string) |  | validator is the validator for the depositor's delegation |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | amount is the delegation balance to convert |






<a name="nemo.router.v1beta1.MsgMintDepositResponse"></a>

### MsgMintDepositResponse
MsgMintDepositResponse defines the Msg/MsgMintDeposit response type.






<a name="nemo.router.v1beta1.MsgWithdrawBurn"></a>

### MsgWithdrawBurn
MsgWithdrawBurn removes staking derivatives from an earn vault and converts them back to a staking delegation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from` | [string](#string) |  | from is the owner of the earn vault to withdraw from |
| `validator` | [string](#string) |  | validator is the address to select the derivative denom to withdraw |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | amount is the staked token equivalent to withdraw |






<a name="nemo.router.v1beta1.MsgWithdrawBurnResponse"></a>

### MsgWithdrawBurnResponse
MsgWithdrawBurnResponse defines the Msg/MsgWithdrawBurn response type.






<a name="nemo.router.v1beta1.MsgWithdrawBurnUndelegate"></a>

### MsgWithdrawBurnUndelegate
MsgWithdrawBurnUndelegate removes staking derivatives from an earn vault, converts them to a staking delegation,
then undelegates them from their validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from` | [string](#string) |  | from is the owner of the earn vault to withdraw from |
| `validator` | [string](#string) |  | validator is the address to select the derivative denom to withdraw |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | amount is the staked token equivalent to withdraw |






<a name="nemo.router.v1beta1.MsgWithdrawBurnUndelegateResponse"></a>

### MsgWithdrawBurnUndelegateResponse
MsgWithdrawBurnUndelegateResponse defines the Msg/MsgWithdrawBurnUndelegate response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.router.v1beta1.Msg"></a>

### Msg
Msg defines the router Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `MintDeposit` | [MsgMintDeposit](#nemo.router.v1beta1.MsgMintDeposit) | [MsgMintDepositResponse](#nemo.router.v1beta1.MsgMintDepositResponse) | MintDeposit converts a delegation into staking derivatives and deposits it all into an earn vault. | |
| `DelegateMintDeposit` | [MsgDelegateMintDeposit](#nemo.router.v1beta1.MsgDelegateMintDeposit) | [MsgDelegateMintDepositResponse](#nemo.router.v1beta1.MsgDelegateMintDepositResponse) | DelegateMintDeposit delegates tokens to a validator, then converts them into staking derivatives, then deposits to an earn vault. | |
| `WithdrawBurn` | [MsgWithdrawBurn](#nemo.router.v1beta1.MsgWithdrawBurn) | [MsgWithdrawBurnResponse](#nemo.router.v1beta1.MsgWithdrawBurnResponse) | WithdrawBurn removes staking derivatives from an earn vault and converts them back to a staking delegation. | |
| `WithdrawBurnUndelegate` | [MsgWithdrawBurnUndelegate](#nemo.router.v1beta1.MsgWithdrawBurnUndelegate) | [MsgWithdrawBurnUndelegateResponse](#nemo.router.v1beta1.MsgWithdrawBurnUndelegateResponse) | WithdrawBurnUndelegate removes staking derivatives from an earn vault, converts them to a staking delegation, then undelegates them from their validator. | |

 <!-- end services -->



<a name="nemo/savings/v1beta1/store.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/savings/v1beta1/store.proto



<a name="nemo.savings.v1beta1.Deposit"></a>

### Deposit
Deposit defines an amount of coins deposited into a savings module account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="nemo.savings.v1beta1.Params"></a>

### Params
Params defines the parameters for the savings module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `supported_denoms` | [string](#string) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/savings/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/savings/v1beta1/genesis.proto



<a name="nemo.savings.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the savings module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#nemo.savings.v1beta1.Params) |  | params defines all the parameters of the module. |
| `deposits` | [Deposit](#nemo.savings.v1beta1.Deposit) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/savings/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/savings/v1beta1/query.proto



<a name="nemo.savings.v1beta1.QueryDepositsRequest"></a>

### QueryDepositsRequest
QueryDepositsRequest defines the request type for querying x/savings
deposits.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="nemo.savings.v1beta1.QueryDepositsResponse"></a>

### QueryDepositsResponse
QueryDepositsResponse defines the response type for querying x/savings
deposits.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposits` | [Deposit](#nemo.savings.v1beta1.Deposit) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="nemo.savings.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest defines the request type for querying x/savings
parameters.






<a name="nemo.savings.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse defines the response type for querying x/savings
parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#nemo.savings.v1beta1.Params) |  |  |






<a name="nemo.savings.v1beta1.QueryTotalSupplyRequest"></a>

### QueryTotalSupplyRequest
QueryTotalSupplyRequest defines the request type for Query/TotalSupply method.






<a name="nemo.savings.v1beta1.QueryTotalSupplyResponse"></a>

### QueryTotalSupplyResponse
TotalSupplyResponse defines the response type for the Query/TotalSupply method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `height` | [int64](#int64) |  | Height is the block height at which these totals apply |
| `result` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | Result is a list of coins supplied to savings |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.savings.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for savings module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#nemo.savings.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#nemo.savings.v1beta1.QueryParamsResponse) | Params queries all parameters of the savings module. | GET|/nemo/savings/v1beta1/params|
| `Deposits` | [QueryDepositsRequest](#nemo.savings.v1beta1.QueryDepositsRequest) | [QueryDepositsResponse](#nemo.savings.v1beta1.QueryDepositsResponse) | Deposits queries savings deposits. | GET|/nemo/savings/v1beta1/deposits|
| `TotalSupply` | [QueryTotalSupplyRequest](#nemo.savings.v1beta1.QueryTotalSupplyRequest) | [QueryTotalSupplyResponse](#nemo.savings.v1beta1.QueryTotalSupplyResponse) | TotalSupply returns the total sum of all coins currently locked into the savings module. | GET|/nemo/savings/v1beta1/total_supply|

 <!-- end services -->



<a name="nemo/savings/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/savings/v1beta1/tx.proto



<a name="nemo.savings.v1beta1.MsgDeposit"></a>

### MsgDeposit
MsgDeposit defines the Msg/Deposit request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="nemo.savings.v1beta1.MsgDepositResponse"></a>

### MsgDepositResponse
MsgDepositResponse defines the Msg/Deposit response type.






<a name="nemo.savings.v1beta1.MsgWithdraw"></a>

### MsgWithdraw
MsgWithdraw defines the Msg/Withdraw request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="nemo.savings.v1beta1.MsgWithdrawResponse"></a>

### MsgWithdrawResponse
MsgWithdrawResponse defines the Msg/Withdraw response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.savings.v1beta1.Msg"></a>

### Msg
Msg defines the savings Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Deposit` | [MsgDeposit](#nemo.savings.v1beta1.MsgDeposit) | [MsgDepositResponse](#nemo.savings.v1beta1.MsgDepositResponse) | Deposit defines a method for depositing funds to the savings module account | |
| `Withdraw` | [MsgWithdraw](#nemo.savings.v1beta1.MsgWithdraw) | [MsgWithdrawResponse](#nemo.savings.v1beta1.MsgWithdrawResponse) | Withdraw defines a method for withdrawing funds to the savings module account | |

 <!-- end services -->



<a name="nemo/swap/v1beta1/swap.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/swap/v1beta1/swap.proto



<a name="nemo.swap.v1beta1.AllowedPool"></a>

### AllowedPool
AllowedPool defines a pool that is allowed to be created


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token_a` | [string](#string) |  | token_a represents the a token allowed |
| `token_b` | [string](#string) |  | token_b represents the b token allowed |






<a name="nemo.swap.v1beta1.Params"></a>

### Params
Params defines the parameters for the swap module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `allowed_pools` | [AllowedPool](#nemo.swap.v1beta1.AllowedPool) | repeated | allowed_pools defines that pools that are allowed to be created |
| `swap_fee` | [string](#string) |  | swap_fee defines the swap fee for all pools |






<a name="nemo.swap.v1beta1.PoolRecord"></a>

### PoolRecord
PoolRecord represents the state of a liquidity pool
and is used to store the state of a denominated pool


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_id` | [string](#string) |  | pool_id represents the unique id of the pool |
| `reserves_a` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | reserves_a is the a token coin reserves |
| `reserves_b` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | reserves_b is the a token coin reserves |
| `total_shares` | [string](#string) |  | total_shares is the total distrubuted shares of the pool |






<a name="nemo.swap.v1beta1.ShareRecord"></a>

### ShareRecord
ShareRecord stores the shares owned for a depositor and pool


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [bytes](#bytes) |  | depositor represents the owner of the shares |
| `pool_id` | [string](#string) |  | pool_id represents the pool the shares belong to |
| `shares_owned` | [string](#string) |  | shares_owned represents the number of shares owned by depsoitor for the pool_id |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/swap/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/swap/v1beta1/genesis.proto



<a name="nemo.swap.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the swap module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#nemo.swap.v1beta1.Params) |  | params defines all the paramaters related to swap |
| `pool_records` | [PoolRecord](#nemo.swap.v1beta1.PoolRecord) | repeated | pool_records defines the available pools |
| `share_records` | [ShareRecord](#nemo.swap.v1beta1.ShareRecord) | repeated | share_records defines the owned shares of each pool |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="nemo/swap/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/swap/v1beta1/query.proto



<a name="nemo.swap.v1beta1.DepositResponse"></a>

### DepositResponse
DepositResponse defines a single deposit query response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  | depositor represents the owner of the deposit |
| `pool_id` | [string](#string) |  | pool_id represents the pool the deposit is for |
| `shares_owned` | [string](#string) |  | shares_owned presents the shares owned by the depositor for the pool |
| `shares_value` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | shares_value represents the coin value of the shares_owned |






<a name="nemo.swap.v1beta1.PoolResponse"></a>

### PoolResponse
Pool represents the state of a single pool


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  | name represents the name of the pool |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | coins represents the total reserves of the pool |
| `total_shares` | [string](#string) |  | total_shares represents the total shares of the pool |






<a name="nemo.swap.v1beta1.QueryDepositsRequest"></a>

### QueryDepositsRequest
QueryDepositsRequest is the request type for the Query/Deposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  | owner optionally filters deposits by owner |
| `pool_id` | [string](#string) |  | pool_id optionally fitlers deposits by pool id |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="nemo.swap.v1beta1.QueryDepositsResponse"></a>

### QueryDepositsResponse
QueryDepositsResponse is the response type for the Query/Deposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposits` | [DepositResponse](#nemo.swap.v1beta1.DepositResponse) | repeated | deposits returns the deposits matching the requested parameters |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="nemo.swap.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest defines the request type for querying x/swap parameters.






<a name="nemo.swap.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse defines the response type for querying x/swap parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#nemo.swap.v1beta1.Params) |  | params represents the swap module parameters |






<a name="nemo.swap.v1beta1.QueryPoolsRequest"></a>

### QueryPoolsRequest
QueryPoolsRequest is the request type for the Query/Pools RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool_id` | [string](#string) |  | pool_id filters pools by id |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="nemo.swap.v1beta1.QueryPoolsResponse"></a>

### QueryPoolsResponse
QueryPoolsResponse is the response type for the Query/Pools RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pools` | [PoolResponse](#nemo.swap.v1beta1.PoolResponse) | repeated | pools represents returned pools |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.swap.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for swap module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#nemo.swap.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#nemo.swap.v1beta1.QueryParamsResponse) | Params queries all parameters of the swap module. | GET|/nemo/swap/v1beta1/params|
| `Pools` | [QueryPoolsRequest](#nemo.swap.v1beta1.QueryPoolsRequest) | [QueryPoolsResponse](#nemo.swap.v1beta1.QueryPoolsResponse) | Pools queries pools based on pool ID | GET|/nemo/swap/v1beta1/pools|
| `Deposits` | [QueryDepositsRequest](#nemo.swap.v1beta1.QueryDepositsRequest) | [QueryDepositsResponse](#nemo.swap.v1beta1.QueryDepositsResponse) | Deposits queries deposit details based on owner address and pool | GET|/nemo/swap/v1beta1/deposits|

 <!-- end services -->



<a name="nemo/swap/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nemo/swap/v1beta1/tx.proto



<a name="nemo.swap.v1beta1.MsgDeposit"></a>

### MsgDeposit
MsgDeposit represents a message for depositing liquidity into a pool


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `depositor` | [string](#string) |  | depositor represents the address to deposit funds from |
| `token_a` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | token_a represents one token of deposit pair |
| `token_b` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | token_b represents one token of deposit pair |
| `slippage` | [string](#string) |  | slippage represents the max decimal percentage price change |
| `deadline` | [int64](#int64) |  | deadline represents the unix timestamp to complete the deposit by |






<a name="nemo.swap.v1beta1.MsgDepositResponse"></a>

### MsgDepositResponse
MsgDepositResponse defines the Msg/Deposit response type.






<a name="nemo.swap.v1beta1.MsgSwapExactForTokens"></a>

### MsgSwapExactForTokens
MsgSwapExactForTokens represents a message for trading exact coinA for coinB


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `requester` | [string](#string) |  | represents the address swaping the tokens |
| `exact_token_a` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | exact_token_a represents the exact amount to swap for token_b |
| `token_b` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | token_b represents the desired token_b to swap for |
| `slippage` | [string](#string) |  | slippage represents the maximum change in token_b allowed |
| `deadline` | [int64](#int64) |  | deadline represents the unix timestamp to complete the swap by |






<a name="nemo.swap.v1beta1.MsgSwapExactForTokensResponse"></a>

### MsgSwapExactForTokensResponse
MsgSwapExactForTokensResponse defines the Msg/SwapExactForTokens response
type.






<a name="nemo.swap.v1beta1.MsgSwapForExactTokens"></a>

### MsgSwapForExactTokens
MsgSwapForExactTokens represents a message for trading coinA for an exact
coinB


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `requester` | [string](#string) |  | represents the address swaping the tokens |
| `token_a` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | token_a represents the desired token_a to swap for |
| `exact_token_b` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | exact_token_b represents the exact token b amount to swap for token a |
| `slippage` | [string](#string) |  | slippage represents the maximum change in token_a allowed |
| `deadline` | [int64](#int64) |  | deadline represents the unix timestamp to complete the swap by |






<a name="nemo.swap.v1beta1.MsgSwapForExactTokensResponse"></a>

### MsgSwapForExactTokensResponse
MsgSwapForExactTokensResponse defines the Msg/SwapForExactTokensResponse
response type.






<a name="nemo.swap.v1beta1.MsgWithdraw"></a>

### MsgWithdraw
MsgWithdraw represents a message for withdrawing liquidity from a pool


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from` | [string](#string) |  | from represents the address we are withdrawing for |
| `shares` | [string](#string) |  | shares represents the amount of shares to withdraw |
| `min_token_a` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | min_token_a represents the minimum a token to withdraw |
| `min_token_b` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | min_token_a represents the minimum a token to withdraw |
| `deadline` | [int64](#int64) |  | deadline represents the unix timestamp to complete the withdraw by |






<a name="nemo.swap.v1beta1.MsgWithdrawResponse"></a>

### MsgWithdrawResponse
MsgWithdrawResponse defines the Msg/Withdraw response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="nemo.swap.v1beta1.Msg"></a>

### Msg
Msg defines the swap Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Deposit` | [MsgDeposit](#nemo.swap.v1beta1.MsgDeposit) | [MsgDepositResponse](#nemo.swap.v1beta1.MsgDepositResponse) | Deposit defines a method for depositing liquidity into a pool | |
| `Withdraw` | [MsgWithdraw](#nemo.swap.v1beta1.MsgWithdraw) | [MsgWithdrawResponse](#nemo.swap.v1beta1.MsgWithdrawResponse) | Withdraw defines a method for withdrawing liquidity into a pool | |
| `SwapExactForTokens` | [MsgSwapExactForTokens](#nemo.swap.v1beta1.MsgSwapExactForTokens) | [MsgSwapExactForTokensResponse](#nemo.swap.v1beta1.MsgSwapExactForTokensResponse) | SwapExactForTokens represents a message for trading exact coinA for coinB | |
| `SwapForExactTokens` | [MsgSwapForExactTokens](#nemo.swap.v1beta1.MsgSwapForExactTokens) | [MsgSwapForExactTokensResponse](#nemo.swap.v1beta1.MsgSwapForExactTokensResponse) | SwapForExactTokens represents a message for trading coinA for an exact coinB | |

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

