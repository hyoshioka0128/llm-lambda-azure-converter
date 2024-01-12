package models

type Node struct {
	Iteration      int
	Id             string
	Code           string
	Feedback       string
	SelfReflection string
	Score          float32

	ParentNode *Node
	ChildNodes []*Node
}
