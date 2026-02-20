package config

import (
	"flag"
	"log/slog"

	// "fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type HttpServer struct{

	Addr string
}
//struct embedding of HTTPserver inn COnfig
type Config struct {

	Env string  `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`

	StoragePath string  `yaml:"storage_path" env-required:"true"` // struct tags 


	HttpServer `yaml:"http_server"`
}


func MustLoad()*Config{
	godotenv.Load()
	var configPath string
configPath=os.Getenv("CONFIG_PATH")
slog.Info("using config path", slog.String("path", configPath))
if configPath==""{
	flags:= flag.String("config","","path to configuration file")

	flag.Parse()

	configPath=*flags 

	if configPath==""{
		log.Fatal("config path is not set")
	}

}
if _,err:= os.Stat(configPath); os.IsNotExist(err){
	log.Fatalf("config does not exist: %s",configPath)
}

var cfg Config

err:= cleanenv.ReadConfig(configPath,&cfg)

if err!= nil{
	log.Fatalf("can not read config :%s",err.Error())
}

return &cfg
}