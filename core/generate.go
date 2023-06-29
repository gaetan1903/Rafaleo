package core

import (
	"fmt"
	"log"
	"os"

	"github.com/gaetan1903/rafaleo/structure"
	"gopkg.in/yaml.v3"
)

func reverseSlice(s []structure.Schema) []structure.Schema {
	length := len(s)
	reversed := make([]structure.Schema, length)

	for i, v := range s {
		reversed[length-i-1] = v
	}

	return reversed
}

func SortSchemabyForeignKey(config *structure.Configuration) {
	// Créer une map pour stocker les schémas
	schemas := make(map[string]structure.Schema)

	// Extraire les schémas
	for _, schema := range config.Schemas {
		schemas[schema.Name] = schema
	}

	// Créer une variable pour stocker les schémas triés
	var sortedSchemas []structure.Schema

	// Fonction récursive pour trier les schémas par dépendances
	var sortDependencies func(schema structure.Schema)
	sortDependencies = func(schema structure.Schema) {
		// Parcourir les colonnes du schéma
		for _, column := range schema.Columns {
			// Vérifier si la colonne a une clé étrangère
			if column.ForeignKey != "" {
				// Récupérer le schéma dépendant
				dependentSchema := schemas[column.ForeignKey]

				// Vérifier si le schéma dépendant n'est pas déjà dans la liste des schémas triés
				found := false
				for _, s := range sortedSchemas {
					if s.Name == dependentSchema.Name {
						found = true
						break
					}
				}

				// Si le schéma dépendant n'est pas déjà dans la liste des schémas triés, le trier (récursivement)
				if !found {
					sortDependencies(dependentSchema)
					sortedSchemas = append(sortedSchemas, dependentSchema)
				}
			}
		}

		// Vérifier si le schéma courant n'est pas déjà dans la liste des schémas triés
		found := false
		for _, s := range sortedSchemas {
			if s.Name == schema.Name {
				found = true
				break
			}
		}

		// Si le schéma courant n'est pas déjà dans la liste des schémas triés, l'ajouter
		if !found {
			sortedSchemas = append(sortedSchemas, schema)
		}
	}

	// Trier les schémas par dépendances
	for _, schema := range schemas {
		sortDependencies(schema)
	}

	// Assigner les schémas triés à la configuration
	config.Schemas = reverseSlice(sortedSchemas)
}

func trt_data(contents []byte, config *structure.Configuration) error {

	// Créer une variable temporaire pour stocker les données YAML
	var yamlData map[string]interface{}

	// Analyser le contenu YAML dans la variable temporaire
	err := yaml.Unmarshal(contents, &yamlData)
	if err != nil {
		log.Fatalf("Erreur lors de l'analyse du contenu YAML : %v", err)
		return err
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
		config.ProviderType = "rest"
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
		config.ProviderType = "database"
	case "sqlite":
		sqliteConfig := structure.SqLiteConfig{
			Name: providerData["name"].(string),
			Path: providerData["path"].(string),
		}
		config.Provider = sqliteConfig
		config.ProviderType = "sqlite"
	default:
		log.Fatalf("Fournisseur inconnu : %s", providerData["type"].(string))
	}

	return nil
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

	// Créer une instance de Configuration pour stocker les données YAML
	var config structure.Configuration

	trt_data(contents, &config)

	SortSchemabyForeignKey(&config)

}
