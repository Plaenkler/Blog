---
title: "Adding YAML Configuration to Go Programs"
date: 2023-01-24T00:00:00+00:00
draft: false
summary: "Easily manage configurations for your Go program using YAML files. Learn simple and readable way to manage settings in Go development."
---

# Adding YAML Configuration to Go Programs

When creating a Go program, it is useful to configure various aspects of the application, such as the database connection details. One way to manage these configurations is to use YAML files, which provide a simple and readable format for storing data. In this post, I will show you how to insert a YAML configuration into a Go program.

## Installing a suitable YAML package

The first step is to choose a library that simplifies operations with *.yaml files. The best supported and most used library is `gopkg.in/yaml.v3`.
To install the package, run the following command:

```golang
go get gopkg.in/yaml.v3
```

## Create a YAML configuration file

Next, create a new file in your project directory called "config.yaml" (or any other name you prefer) and add the following contents:

```yaml
database:
  host: localhost
  port: 5432
  user: myuser
  password: mypassword
```

This is a simple example of a YAML configuration file, but you can add as many key-value pairs as you need to configure your application.

## Parse the YAML configuration file

To parse the YAML configuration file, you will need to use the `gopkg.in/yaml.v3` package you installed earlier. Here is an example of how to parse the file and store the configuration in a struct:

```golang
package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}
```

This code starts by defining a struct called "Config" that matches the structure of the YAML configuration file. This struct has a field "Database", which is itself a struct. Each field of the struct is a string with a YAML tag specifying the key of the equivalent field in the YAML file. Basically, there is an infinite possibility to make further nestings by means of structs.

```golang
func main() {
	config := Config{}

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("could not read file: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("could not unmarshal file: %v", err)
	}

	fmt.Println(config.Database.Host)
	fmt.Println(config.API.Key)
}
```

Then the code uses the `ioutil` package to read the contents of the previously created file, which is stored in the yamlFile variable. The `yaml.v3` package is then used to parse the contents of the variable and store it in the config structure. This is done with the function `yaml.Unmarshal`.

## Use the parsed configuration

Once the YAML configuration file has been parsed and stored in a struct, you can use the configuration values in your Go program. In the example above, the fields `host` and `key` are printed to the console using `fmt.Println`. In your application, these values can now be used to establish a database connection, for example.

## Conclusion

By using the `gopkg.in/yaml.v3` package, we can easily parse a YAML configuration file and store the values in a struct, which can then be used to configure our program. This approach provides an easy way to manage configuration settings, and is a common pattern in Go development.

It is worth mentioning that this simple approach is well suited for small projects, it gives a lot of freedom but additional functions must be implemented independently. A package like [viper](https://github.com/spf13/viper) can help here, but adds more complexity and dependencies to the project.