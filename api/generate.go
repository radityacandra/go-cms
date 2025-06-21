package api

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config config.yaml ./api.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --include-tags users -config ./user/config-user.yaml ./api.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --include-tags auth -config ./auth/config-auth.yaml ./api.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --include-tags auth-private -config ./authPrivate/config-auth-private.yaml ./api.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --include-tags article-public -config ./article/config-article.yaml ./api.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --include-tags article-private -config ./articlePrivate/config-article-private.yaml ./api.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --include-tags tags -config ./tag/config-tag.yaml ./api.yaml
