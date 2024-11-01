package layer2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTopic(t *testing.T) {
	assert.Equal(t, AddressRegisteredTopic.String(), "0x10d3f3c3d0c7da2f6751b14c10b9dbc6e04f5ebc6b798a6e220f3857ba1cd454")
	assert.Equal(t, TaskSubmittedTopic.String(), "0x78e6ef797d565365876ecf57b6b91c9f0cb3da890e073ffcc8a3be2c28e145e0")
	assert.Equal(t, SubmitterChosenTopic.String(), "0x0d6caedcf9fb56222a63417673875559577b650f769290f255258825d907867d")
	assert.Equal(t, ParticipantAddedTopic.String(), "0x31d3ac54da09405b02d1de0ee0de648de637fbdc111123be0d7fc31f2a544c0b")
	assert.Equal(t, ParticipantRemovedTopic.String(), "0x1a5e355a9a34d7eac1e439a6c610ba1fa72aa45f7645724ce5187fa19c3bd3fc")
}
