# golang-todo-webapp

Go 言語の復習と Web アプリの構築練習用

# Tips

GraphQL スキーマの生成

`schema.graphsql` を編集してから以下を実行

```sh
go run github.com/99designs/gqlgen generate
```

GraphQL playground からクエリ実行

```
mutation createTodo {
  createTodo(input: { task: "todo1", description: "description" }) {
    task
    description
  }
}

query {
  todos {
    id
    task
    description
  }
}
```
