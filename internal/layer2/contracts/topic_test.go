package contracts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopic(t *testing.T) {
	assert.Equal(t, AddressRegisteredTopic.String(), "0x0ab661710c67363885e0e51920050375aff9dcd587adf3e2e468e060ee8f0e1e")
	assert.Equal(t, TaskSubmittedTopic.String(), "0x7c6cba37f838a9f6cd45be5dbe20a2a6c0a373fcb738333fbc39ab558183576f")
	assert.Equal(t, TaskUpdatedTopic.String(), "0x30a99b2ffff1813c032a6b15bb8a15c2c3d1e9bc6dcb5f5cd80238514e86f364")
	assert.Equal(t, SubmitterRotationRequestedTopic.String(), "0x810bb46f7f5182d661c517393732ca0639393a548c222be3f52830dbd81b5584")
	assert.Equal(t, SubmitterChosenTopic.String(), "0x0d6caedcf9fb56222a63417673875559577b650f769290f255258825d907867d")
	assert.Equal(t, ParticipantAddedTopic.String(), "0x31d3ac54da09405b02d1de0ee0de648de637fbdc111123be0d7fc31f2a544c0b")
	assert.Equal(t, ParticipantRemovedTopic.String(), "0x1a5e355a9a34d7eac1e439a6c610ba1fa72aa45f7645724ce5187fa19c3bd3fc")
	assert.Equal(t, ParticipantsResetTopic.String(), "0x32e9d8d19fb1e71c8dc610e5f45fd7f1e2f81babf8ea90e267475a708e09c35e")
	assert.Equal(t, DepositRecordedTopic.String(), "0x8185b6fdeefb24e6918abb5af00007e5bba2e904f6593d5517789c7b76e5d750")
	assert.Equal(t, WithdrawalRecordedTopic.String(), "0x2afe20970cc53fcbca49f0fd13ca943d027d7636ad0d9ac543a995a0cd03c9ec")
	assert.Equal(t, NIP20TokenEventBurnbTopic.String(), "0xebe23dd93b970477278ceb9abd3082df92d977d6131fb0ef75f18c3d353b565a")
	assert.Equal(t, NIP20TokenEventMintbTopic.String(), "0x685c530c280ee1f7a4e96d082303ee9ebf21cec512259c6a943eda3854e05102")
	assert.Equal(t, AssetListedTopic.String(), "0x9a321153d23feb212bd45898d567f7472a2ecc6c752d8e6757ea0914ab2b7009")
	assert.Equal(t, AssetUpdatedTopic.String(), "0xc91c349180cb8d82f404c5b3cb776276676bba19867657b1e45e84e3103d0b36")
	assert.Equal(t, AssetDelistedTopic.String(), "0x0b1d3e62bf94aca06b9bfeae6fb1f3eda5442d64eef89d34b56bba36d69348e6")
	assert.Equal(t, WalletCreationRequestTopic.String(), "0xd9ae647e221b7680b4845a324077f471012748046751cffc1ae442d3791a330b")
	assert.Equal(t, DepositRequestTopic.String(), "0x85ef748ae6693b5b635136aefd33d3655340977a64f640fba6b5b949be4d7a5f")
	assert.Equal(t, WithdrawalRequestTopic.String(), "0x50b0ce2f0e8416e5f2c2a731451fb28dfaf54ddd1431284a24ece1316011b2bf")
	assert.Equal(t, WalletCreationResultTopic.String(), "0x440691550bb1f6d18c60b1a17fff36325a996ba3ab5917f3003445984c5302cf")
	assert.Equal(t, DepositResultTopic.String(), "0xae9e6016838d9912f513c2adb0656673485ababaddbf853b28d121bf2ce24b9e")
	assert.Equal(t, WithdrawalResultTopic.String(), "0x9a474499969867585df13ccda2ed8f3f9ad89cd1704e038cb941e1fbdc1c08fe")
}
