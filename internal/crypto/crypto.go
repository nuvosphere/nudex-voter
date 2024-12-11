package crypto

type PublicKey interface {
	SerializeCompressed() []byte
	SerializeUncompressed() []byte
}

//type ExtendPublicKey interface {
//	PublicKey
//	Address() string
//	GetType() int
//}

type PrivateKey interface {
	Serialize() []byte
}
