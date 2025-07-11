package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct{
  Db_url string `json:"db_url"`
  Current_user_name string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json" 

func Read() (Config, error) {
  path, err := getConfigFilePath()
  if err != nil{
    return Config{}, err
  }

  bytedata, err := os.ReadFile(path)
  if err != nil{
    return Config{}, err 
  }

  var con Config
  err = json.Unmarshal(bytedata, &con)
  if err != nil{
    return Config{}, err 
  }

  return con, nil
}


func getConfigFilePath() (string, error){
  home, err := os.UserHomeDir()
  if err != nil{
    return "", fmt.Errorf("error retrieving home dir")
  }
  return home+"/"+configFileName, nil

}

func (conf Config) SetUser(username string) error{
  conf.Current_user_name = username 
  err := write(conf)
  if err != nil{
    return err
  }
  return nil
}

func write(cfg Config) error{
  bytedata, err := json.Marshal(&cfg)
  if err != nil{
    return err
  }

  path, err := getConfigFilePath()
  if err != nil{
    return err
  }
  err = os.WriteFile(path, bytedata, 0666)
  if err != nil{
    return err
  }

  return nil
}
