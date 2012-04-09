package config

var configFile = "init.ini"


type Config struct {
	Key,Value string
}

type ConfigMap struct{
	Configs []*Config
}

func SetConfigFile(fileName string){
	configFile = fileName
}

func ReadToMap(){
	
}


func GetValue(key string) string{
	var result = ""
	if key == "" {
		result = "null"
	}
	return result;
}

func WriteKeyValue(key, value string){


}

func WriteConfig(cfg *Config){

}


