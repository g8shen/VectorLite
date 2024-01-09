
# VectorLite

Lightweight vector database designed for handling multidimensional data points. 

## Current Features

- `Retrevial` : k-NN (K Nearest Neighbor) search for approximate matching. Direct key-based retrieval.
- `Distance Metrics` : Euclidean and Cosine similarity for either magnitude or orientation analysis. 
- `Persistence & CRUD` : Saving and loading data for session continuity and accessibility through CRUD operations.
- `Categorical Data Encoding` : Incorporates one-hot encoding to transform non-numerical, categorical data to vectors.

## Set up

Go is the only prerequisites

1. Clone the repository
2. Open a terminal at the folder directory
3. Enter  `cd dbms`
4. Enter `go run .`
5. The k-NN search and CRUD demo should run
6. Play around!

## Basic Usage

### Load

```go
db := VectorDB{}
db.Load("../datastore/example.gob")
```

### Data Model 

Each data point is encapsulated in a vectory entry which consists of:

- `Vector` : An array of floats representing the data point in a multidimensional vector space e.g it could derived from text analysis or image characteristics.
- `Key` : A unique string identifier for the vector used for retrieving, updating, or deleting specific entries within the database.
- `JSON Data` : An accompanying JSON object providing additional context or metadata for the vector. 

```go
vector := Vector{0.1, 0.3, 0.7, 0.92, 3, 0, 1, 8, 13, 0.5, 0, 0, 0, 1} 
key := "sample-vector-1"         
jsonData := `{
    "label": "Sample Label",
    "category": "Example Category"
}` 

vectorEntry := VectorEntry{
    Vector: vector,
    Key:    key,
    Data:   json.RawMessage(jsonData),
}
```

### k-NN search

```go
nearestNeighbours := db.kNN(queryVector.Vector, 5, euclideanDistance)
```

### Insert 

```go
db.Insert(VectorEntry)
```

### Update

```go
db.Update("key", VectorEntry)
```

### Delete 

```go
db.Delete("key")
```
