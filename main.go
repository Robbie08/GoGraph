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

// getVertex
// contains
func contains(s []*Vertex, k int) bool {
    for _,v := range s{
        if k == v.key{
            return true
        }
    }
    return false
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

    for i := 0; i < 5; i++ {
        test.addVertex(i)
    }
    test.addVertex(1)
    test.Print()
}
