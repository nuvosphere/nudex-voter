package contracts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopic(t *testing.T) {
	assert.Equal(t, AddressRegisteredTopic.String(), "0x10d3f3c3d0c7da2f6751b14c10b9dbc6e04f5ebc6b798a6e220f3857ba1cd454")
	assert.Equal(t, TaskSubmittedTopic.String(), "0x78e6ef797d565365876ecf57b6b91c9f0cb3da890e073ffcc8a3be2c28e145e0")
	assert.Equal(t, TaskCompletedTopic.String(), "0x8d728f48e16602b5feb9528aa4cbf18ab06ee67918838342dfbd65d218e2a3ee")
	assert.Equal(t, SubmitterRotationRequestedTopic.String(), "0x810bb46f7f5182d661c517393732ca0639393a548c222be3f52830dbd81b5584")
	assert.Equal(t, SubmitterChosenTopic.String(), "0x0d6caedcf9fb56222a63417673875559577b650f769290f255258825d907867d")
	assert.Equal(t, ParticipantAddedTopic.String(), "0x31d3ac54da09405b02d1de0ee0de648de637fbdc111123be0d7fc31f2a544c0b")
	assert.Equal(t, ParticipantRemovedTopic.String(), "0x1a5e355a9a34d7eac1e439a6c610ba1fa72aa45f7645724ce5187fa19c3bd3fc")
	assert.Equal(t, DepositRecordedTopic.String(), "0xda0e7c971690dbb1d8118c31cf27f8303b471719eb78a6200d35300175974100")
	assert.Equal(t, WithdrawalRecordedTopic.String(), "0x07c8f2d211076c7cba51f2504af48025acdaf410e993e6c7f62b066a51d9b068")
	assert.Equal(t, WalletCreationRequestTopic.String(), "0xfcfa9d5597fec4480a4a1fdbc38fc3fd82fdc02523a0ab0a2a12efb8dc0baf6f")
	assert.Equal(t, DepositRequestTopic.String(), "0xf26824b39d161a7a9fdacf5c40e9fa65dd2a06d5a5b12719d68e55ae2683491f")
	assert.Equal(t, WithdrawalRequestTopic.String(), "0x4cd151cd1ef25ad1d7498e51a1ba595d51db95d66bfc1611a341aa0662003d84")
	assert.Equal(t, WalletCreationResultTopic.String(), "0x440691550bb1f6d18c60b1a17fff36325a996ba3ab5917f3003445984c5302cf")
	assert.Equal(t, DepositResultTopic.String(), "0xae9e6016838d9912f513c2adb0656673485ababaddbf853b28d121bf2ce24b9e")
	assert.Equal(t, WithdrawalResultTopic.String(), "0x9a474499969867585df13ccda2ed8f3f9ad89cd1704e038cb941e1fbdc1c08fe")
}