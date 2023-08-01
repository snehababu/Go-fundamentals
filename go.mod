module github.com/353colutions/go-essential //defines the module

go 1.17 //minimal go version

//direct requirements
require(
    github.com/pelletier/go-toml v1.9.4
    github.com/pkg/errors v0.9.1
    github.com/stretchr/testify v1.7.0
)

//indirect requirements
require(
    github.com/devecgh/go-spew v1.1.0 //indirect
    github.com/pmezard/go-difflib v1.0.0 //indirect
    gopkg.in/yaml.v3 v3.0.0-20200323202051-9f266ea9e77c //indirect

)