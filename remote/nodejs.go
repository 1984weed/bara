package remote

import "github.com/go-pg/pg"

type ConfigNode struct {
	ImageName string
}

type NodeJSClient struct {
	store *pg.DB
}

func NewNodeJsClient(store *pg.DB) *NodeJSClient {
	return &NodeJSClient{store: store}
}

type Question struct {
	ID          int64
	Slug        string
	Title       string
	Description string
	ArgID       int
	AuthorID    int
	CreatedAt   int64
	UpdatedAt   int64
}

func (n *NodeJSClient) exec(slug string) string {
	question := &Question{ID: 1}
	err := n.store.Select(question)

	if err != nil {
		return ""
	}

	return "OK"
	// err = store.QueryRow("select id,content,author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)

}
