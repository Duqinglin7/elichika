package config

import (
	_ "modernc.org/sqlite"
	"xorm.io/xorm"
)

var (
	MasterVersion = "b66ec2295e9a00aa"
	StartUpKey    = "uWabor56EHkM30qI"
	SessionKey    = "12345678123456781234567812345678"

	// Taken from:
	// https://github.com/RayFirefist/SukuStar_Datamine/blob/master/lib/sifas_api/sifas.py#L120
	// https://github.com/RayFirefist/SukuStar_Datamine/blob/master/lib/sifas_api/sifas.py#L400
	ServerEventReceiverKey = "31f1f9dc7ac4392d1de26acf99d970e425b63335b461e720c73d6914020d6014"
	UnknownKey             = "78d53d9e645a0305602174e06b98d81f638eaf4a84db19c756866fddac360c96"

	CdnServer = "http://192.168.1.123:8080/static/ja"

	MainDb  = "assets/main.db"
	MainEng *xorm.Engine
)

func init() {
	eng, err := xorm.NewEngine("sqlite", MainDb)
	if err != nil {
		panic(err)
	}
	err = eng.Ping()
	if err != nil {
		panic(err)
	}
	MainEng = eng
	MainEng.SetMaxOpenConns(50)
	MainEng.SetMaxIdleConns(10)
}
