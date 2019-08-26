module github.com/geisonbiazus/markdown_notes/notes

go 1.12

require github.com/stretchr/testify v1.4.0

require github.com/geisonbiazus/markdown_notes/cqrs v0.0.0

replace github.com/geisonbiazus/markdown_notes/cqrs v0.0.0 => ../cqrs

require github.com/geisonbiazus/markdown_notes/validations v0.0.0

replace github.com/geisonbiazus/markdown_notes/validations v0.0.0 => ../validations
