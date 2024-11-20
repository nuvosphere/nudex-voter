package tss

import (
	"context"
	"encoding/hex"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper/testutil"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

func createSubmitter(t *testing.T) common.Address {
	pk, err := crypto.GenerateKey()
	assert.Nil(t, err)

	return crypto.PubkeyToAddress(pk.PublicKey)
}

func createDB(t *testing.T, index int) *gorm.DB {
	basePath := filepath.Join("./", strconv.Itoa(index))
	err := os.MkdirAll(basePath, os.ModePerm)
	assert.Nil(t, err)

	path := filepath.Join(basePath, "relayer_data.db")

	relayerDb, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		Logger:         gormlogger.Default.LogMode(gormlogger.Warn),
		TranslateError: true, // https://gorm.golang.ac.cn/docs/error_handling.html
	})
	assert.Nil(t, err)
	db.SetConnParam(relayerDb)

	return relayerDb
}

type Account struct {
	PK      string
	Address common.Address
}

var accounts = []Account{
	{
		PK:      "76cbb08e5321cec5f584b2b40b4666d9bbbee59eb3022e80d804e8310b17a105",
		Address: common.HexToAddress("0x3a818294ca1F3C27d7588b123Ec43F2546fa07f4"),
	},
	{
		PK:      "ffab86884b5f4696c503e8d0cef97f818d122f44017528c24ce3ac580f12b876",
		Address: common.HexToAddress("0x04d9389Cf937b1e6F2258d842e7237E955d6ab04"),
	},
	{
		PK:      "5d0ca3f7b4e63f3308a73537001065ee1d6ff3e217115444b148018a1bcbfaf7",
		Address: common.HexToAddress("0xf6D37CE75dB465DcDb4c7097bEB9c1D46b171037"),
	},
	{
		PK:      "dd4ae923532c8b47440db5497bf0591769969d3da3ed6ac1d7c2a037033404e9",
		Address: common.HexToAddress("0x1D2cd50A3cF3c55a7982AD54F9f364C1e953Bc57"),
	},
	{
		PK:      "ccb83c6d8cf4d1400ca1d90df2f9c9fafe4b1947ba51c13617603af3bef18590",
		Address: common.HexToAddress("0x5091FC3cb4E4BB014141Aa41375d8Dd73b34AfA2"),
	},
}

func TestCreateAddress(t *testing.T) {
	for i := 0; i < testutil.TestPartyCount; i++ {
		pk, err := crypto.GenerateKey()
		assert.Nil(t, err)

		address := crypto.PubkeyToAddress(pk.PublicKey)
		t.Log("pk", hex.EncodeToString(crypto.FromECDSA(pk)), "address:", address)
	}
}

func TestScheduler(t *testing.T) {
	utils.SkipCI(t)
	log.SetLevel(log.DebugLevel)

	ss := make([]*Scheduler, 0, len(accounts))

	bus := eventbus.NewBus()
	p2pMocker := NewP2PMocker(bus)

	voterContractMocker := NewVoterContractMocker()

	var submitters types.Participants

	for i := 0; i < testutil.TestPartyCount; i++ {
		submitter := accounts[i].Address
		if i == 2 {
			voterContractMocker.SetProposer(submitter)
		}

		submitters = append(submitters, submitter)
	}

	t.Log("submitters", submitters)

	for i, submitter := range submitters {
		t.Logf("index: %d submitter: %v, submitter:%v", i, submitters[i], submitter)
	}

	voterContractMocker.SetParticipants(submitters)

	for i, submitter := range submitters {
		d := createDB(t, i)
		t.Logf("index: %d submitter: %v, submitter:%v", i, submitters[i], submitter)
		s := NewScheduler(false, p2pMocker, bus, d, voterContractMocker, submitter)
		basePath := filepath.Join("./", strconv.Itoa(i))
		err := os.MkdirAll(basePath, os.ModePerm)
		assert.NoError(t, err)

		s.partyData.basePath = basePath
		ss = append(ss, s)
	}

	for _, s := range ss {
		go s.Start()
	}

	time.Sleep(10 * time.Minute)
	lo.ForEach(ss, func(item *Scheduler, index int) { item.Stop() })
}

func TestContext(t *testing.T) {
	utils.SkipCI(t)

	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < testutil.TestPartyCount; i++ {
		go func() {
			<-ctx.Done()
			t.Logf("i := %d", i)
		}()
	}

	t.Log("---------1")
	time.Sleep(10 * time.Second)
	t.Log("---------2")
	cancel()
	t.Log("---------3")
	time.Sleep(10 * time.Second)

	ticker := time.NewTicker(1 * time.Second)

	go func() {
	L:
		for {
			select {
			case <-ctx.Done():
				t.Log("---------Done")
				break L
			case <-ticker.C:
				t.Log("---------ticker")
			}
		}
	}()
	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(10 * time.Second)
}
