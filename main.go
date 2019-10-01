package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Tree struct {
	Id       int
	Children []Tree
}

type CommentsTree struct {
	Id       int
	Comment  string
	Children []CommentsTree
}

func main() {
	tree := Tree{
		Id: 1,
		Children: []Tree{
			{
				Id: 2,
			},
			{
				Id: 3,
				Children: []Tree{
					{
						Id: 4,
					},
					{
						Id: 5,
					},
				},
			},
			{
				Id: 13,
				Children: []Tree{
					{
						Id: 14,
					},
					{
						Id: 15,
						Children: []Tree{
							{
								Id: 155,
							},
						},
					},
				},
			},
		},
	}
	comments := makeComments(tree)
	bytes, _ := json.Marshal(comments)

	fmt.Print(string(bytes))
}

func makeComments(tree Tree) CommentsTree {
	wg := sync.WaitGroup{}
	ct := CommentsTree{}

	wg.Add(1)
	go fillComments(&tree, &ct, &wg)
	wg.Wait()

	return ct
}

func fillComments(tree *Tree, ct *CommentsTree, wg *sync.WaitGroup) {
	childCount := len(tree.Children)
	ct.Children = make([]CommentsTree, childCount)

	wg.Add(childCount + 1)
	for i := 0; i < childCount; i++ {
		go fillComments(&tree.Children[i], &ct.Children[i], wg)
	}

	go slowLoadingComment(tree.Id, ct, wg)
	wg.Done()
}

func slowLoadingComment(id int, ct *CommentsTree, wg *sync.WaitGroup) {
	ct = &CommentsTree {
		Id: id,
		Comment: fmt.Sprintf("Comment №: %d", id),
	}
	wg.Done()
}
