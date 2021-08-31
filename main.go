package main

import (
    "fmt"
)


// Graph Structure - represents an adjacency list graph
type Graph struct{
    vertices []*Vertex
}

// Vertex Structure - represents graph vertex
type Vertex struct{
    key int
    color int
    dist int
    parent *Vertex
    adjacent []*Vertex
}


// Add Vertex
func (g *Graph) addVertex(k int){

    // we only want to add new vertices
    if contains(g.vertices, k) {
        err := fmt.Errorf("Oops! You tried adding vertex %v again. This vertex already exists.", k)
        fmt.Println(err.Error())
    }else{
        // will add our new vertext to the adj list
        g.vertices = append(g.vertices, &Vertex{key: k})
    }
}
// Add Edge
func (g *Graph) addEdge(from, to int){
    // get vertex
    fromVertex := g.getVertex(from) // get address of the from vertex
    toVertex := g.getVertex(to) // get the address of the to vertex

    // check error - if edge exists already, or add edge to non-existent vertex
    if fromVertex == nil || toVertex == nil{
        err := fmt.Errorf("Oops! Invalid edge (%v ->% v)", from, to)
        fmt.Println(err.Error())
    } else if contains(fromVertex.adjacent, to){
        err :=  fmt.Errorf("Oops! Existing edge (%v -> %v)", from, to)
        fmt.Println(err.Error())
    }else{
        // add edge
        fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
    }
}
// getVertex
func (g *Graph) getVertex(k int) *Vertex {
    var vertex *Vertex
    for i, v := range g.vertices {
        if v.key == k {
            vertex = g.vertices[i]
        }
    }
    return vertex
}

// contains
func contains(s []*Vertex, k int) bool {
    for _,v := range s{
        if k == v.key{
            return true
        }
    }
    return false
}

func (g *Graph) DFS(s *Vertex) {

    for _, v := range g.vertices{
        for _, v:= range v.adjacent{
            if v != s{
                v.color = 0
                v.dist = -1
                v.parent = nil
            }
        }
    }

    s.color = 1
    s.dist = 0
    s.parent = nil

    queue := make([]*Vertex, 0) // create and init our queue
    queue = append(queue, s) // append s into the queue

    for len(queue) != 0 {
        u := queue[0]  // keep the vertex at top of queue
        queue = queue[1:] // remove value from queue

        for _, v:= range u.adjacent{
            if v.color == 0{
                v.color = 1
                v.dist = u.dist + 1
                v.parent = u
                queue = append(queue, v)
            }
        }
        fmt.Print(" ",u.key)
        u.color = 2
    }

}

// print contects of the list
func (g *Graph) Print() {
    for _, v := range g.vertices {
        fmt.Printf("\nVertex %v : ", v.key)
        for _,v := range v.adjacent {
            fmt.Printf(" %v ", v.key)
        }
    }
    fmt.Println()
}

func main(){
    test := &Graph{}

    for i := 0; i < 7; i++ {
        test.addVertex(i)
    }

    test.addEdge(1,2)
    test.addEdge(1,3)
    test.addEdge(2,3)
    test.addEdge(3,4)
    test.addEdge(2,5)
    test.addEdge(5,6)
    test.DFS(test.getVertex(1))
}
