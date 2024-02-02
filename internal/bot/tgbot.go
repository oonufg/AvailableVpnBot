package bot

type TgBot struct {
	apiKey string
	TgAPI
	EndPoints
}

type EndPoints interface {
	GetVpnsList(country, proto string) string
	GetVpnFile(fileName string) []byte
}

type TgAPI interface {
	HandleUpdate()
	SendFile()
}
