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
	assert.Equal(t, DepositRecordedTopic.String(), "0xc81b018d055616352576702d0318bf7fc5c5b37693d9d4555113e2490d87dd80")
	assert.Equal(t, WithdrawalRecordedTopic.String(), "0x2afe20970cc53fcbca49f0fd13ca943d027d7636ad0d9ac543a995a0cd03c9ec")
	assert.Equal(t, WalletCreationRequestTopic.String(), "0x497b7d7bb002736945c16f28175435ba9f4ad4863018be370d6901eba91c9008")
	assert.Equal(t, DepositRequestTopic.String(), "0xf26824b39d161a7a9fdacf5c40e9fa65dd2a06d5a5b12719d68e55ae2683491f")
	assert.Equal(t, WithdrawalRequestTopic.String(), "0x4cd151cd1ef25ad1d7498e51a1ba595d51db95d66bfc1611a341aa0662003d84")
	assert.Equal(t, WalletCreationResultTopic.String(), "0x440691550bb1f6d18c60b1a17fff36325a996ba3ab5917f3003445984c5302cf")
	assert.Equal(t, DepositResultTopic.String(), "0xae9e6016838d9912f513c2adb0656673485ababaddbf853b28d121bf2ce24b9e")
	assert.Equal(t, WithdrawalResultTopic.String(), "0x9a474499969867585df13ccda2ed8f3f9ad89cd1704e038cb941e1fbdc1c08fe")
}
