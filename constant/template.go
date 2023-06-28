package constant

var TemplateContent string = `
version: "1.0.0"
schemas:
  - name: ""
    columns:
      - name: ""
        type: ""
        min_length: 
        max_length: 
        faker: ""
        foreign_key: ""
        optional: false

provider:
  rest:
    name: ""
    url: ""
    entity_endpoint:
      schema_name: ""
      path: ""
    method: ""
    headers:
      - name: ""
        value: ""

  database:
    db_type: ""
    name: ""
    host: ""
    port: 
    user: ""
    password: ""
    
  sqlite:
    name: ""
    path: ""
`
