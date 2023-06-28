package core

import (
	"fmt"
	"log"
	"os"

	"github.com/gaetan1903/rafaleo/structure"
	"gopkg.in/yaml.v3"
)

func trt_data(contents []byte) {
	// Créer une instance de Configuration pour stocker les données YAML
	var config structure.Configuration

	// Créer une variable temporaire pour stocker les données YAML
	var yamlData map[string]interface{}

	// Analyser le contenu YAML dans la variable temporaire
	err := yaml.Unmarshal(contents, &yamlData)
	if err != nil {
		log.Fatalf("Erreur lors de l'analyse du contenu YAML : %v", err)
	}

	// Extraire les valeurs du YAML et les assigner à la structure de configuration
	config.Version = yamlData["version"].(string)

	// Extraire les schémas
	schemasData := yamlData["schemas"].([]interface{})
	for _, schemaData := range schemasData {
		schema := structure.Schema{
			Name: schemaData.(map[string]interface{})["name"].(string),
		}
		columnsData := schemaData.(map[string]interface{})["columns"].([]interface{})
		for _, columnData := range columnsData {
			column := structure.Column{
				Name:     columnData.(map[string]interface{})["name"].(string),
				Type:     columnData.(map[string]interface{})["type"].(string),
				Optional: columnData.(map[string]interface{})["optional"].(bool),
			}
			schema.Columns = append(schema.Columns, column)
		}
		config.Schemas = append(config.Schemas, schema)
	}

	// Extraire les données du fournisseur
	providerDataType := yamlData["provider"].(map[string]interface{})
	providerData := yamlData["provider"].(map[string]interface{})[providerDataType["type"].(string)].(map[string]interface{})
	switch providerDataType["type"] {
	case "rest":
		restConfig := structure.RestConfig{
			URL: providerData["url"].(string),
			EntityEndpoint: structure.Endpoint{
				SchemaName: providerData["entity_endpoint"].(map[string]interface{})["schema_name"].(string),
				Path:       providerData["entity_endpoint"].(map[string]interface{})["path"].(string),
			},
			Method: providerData["method"].(string),
		}
		headersData := providerData["headers"].([]interface{})
		for _, headerData := range headersData {
			header := structure.Header{
				Name:  headerData.(map[string]interface{})["name"].(string),
				Value: headerData.(map[string]interface{})["value"].(string),
			}
			restConfig.Headers = append(restConfig.Headers, header)
		}
		config.Provider = restConfig
	case "database":
		databaseConfig := structure.DatabaseConfig{
			Name:     providerData["name"].(string),
			Host:     providerData["host"].(string),
			Port:     int(providerData["port"].(int)),
			User:     providerData["user"].(string),
			Password: providerData["password"].(string),
			DbType:   providerData["db_type"].(string),
		}
		config.Provider = databaseConfig
	case "sqlite":
		sqliteConfig := structure.SqLiteConfig{
			Name: providerData["name"].(string),
			Path: providerData["path"].(string),
		}
		config.Provider = sqliteConfig
	default:
		log.Fatalf("Fournisseur inconnu : %s", providerData["type"].(string))
	}

	// Afficher les données extraites
	fmt.Printf("Version: %s\n", config.Version)
	fmt.Println("Schemas:")
	for _, schema := range config.Schemas {
		fmt.Printf("- Name: %s\n", schema.Name)
		fmt.Println("  Columns:")
		for _, column := range schema.Columns {
			fmt.Printf("  - Name: %s\n", column.Name)
			fmt.Printf("    Type: %s\n", column.Type)
			fmt.Printf("    Optional: %t\n", column.Optional)
		}
	}

	// Afficher les données spécifiques au fournisseur
	switch provider := config.Provider.(type) {
	case structure.RestConfig:
		fmt.Println("Provider: REST")
		fmt.Printf("URL: %s\n", provider.URL)
		fmt.Println("EntityEndpoint:")
		fmt.Printf("  SchemaName: %s\n", provider.EntityEndpoint.SchemaName)
		fmt.Printf("  Path: %s\n", provider.EntityEndpoint.Path)
		fmt.Printf("Method: %s\n", provider.Method)
		fmt.Println("Headers:")
		for _, header := range provider.Headers {
			fmt.Printf("  - Name: %s\n", header.Name)
			fmt.Printf("    Value: %s\n", header.Value)
		}
	case structure.DatabaseConfig:
		fmt.Println("Provider: Database")
		fmt.Printf("DbType: %s\n", provider.DbType)
		fmt.Printf("Name: %s\n", provider.Name)
		fmt.Printf("Host: %s\n", provider.Host)
		fmt.Printf("Port: %d\n", provider.Port)
		fmt.Printf("User: %s\n", provider.User)
		fmt.Printf("Password: %s\n", provider.Password)
	case structure.SqLiteConfig:
		fmt.Println("Provider: SQLite")
		fmt.Printf("Name: %s\n", provider.Name)
		fmt.Printf("Path: %s\n", provider.Path)
	default:
		fmt.Println("Provider: Unknown")
	}

}

func Generate(configFile string) {
	if _, err := os.Stat(configFile); err != nil {
		if os.IsNotExist(err) {
			fmt.Fprintln(os.Stderr, "\033[31mFile", configFile, "does not exist.\033[0m")
			return
		}
		fmt.Fprintln(os.Stderr, "\033[31mError:", err.Error(), "\033[0m")
		return
	}

	// Open the file
	file, err := os.Open(configFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "\033[31mError:", err.Error(), "\033[0m")
		return
	}
	defer file.Close()

	// Read the file contents
	contents, err := os.ReadFile(configFile)

	if err != nil {
		fmt.Fprintln(os.Stderr, "\033[31mError:", err.Error(), "\033[0m")
		return
	}

	trt_data(contents)

}
