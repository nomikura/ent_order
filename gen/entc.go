package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
	// entgql.WithWhereInputs(true),                  // GraphQLファイルにXxxWhereInput型を生成する
	// entgql.WithConfigPath("../gqlgen.yml"),        // 参照するGraphQLの設定ファイル
	// entgql.WithSchemaGenerator(),                  // entからGraphQLの型定義を生成
	// entgql.WithSchemaPath("../graph/ent.graphql"), // entから生成するGraphQLファイルのパス
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	_ = ex

	opts := []entc.Option{
		entc.Extensions(ex),
		// entc.FeatureNames(
		// 	"sql/execquery",   // entのclientから生のSQLを実行できるようにする
		// 	"schema/snapshot", // スキーマのスナップショットを作成して、コード生成が破綻しないようにする
		// 	"sql/upsert",      // upsertを使えるようにする
		// 	"privacy",
		// ),
	}

	if err := entc.Generate("../ent/schema", &gen.Config{}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
