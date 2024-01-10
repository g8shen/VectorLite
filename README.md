
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
5. The k-NN search & CRUD demo should run in the terminal
	- There is a demo preview at the end of README
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

## Demo Preview

Below is the k-NN search with Euclidean distance portion of the demo. The straight-line distance between two points in multi-dimensional space, identifying nearest neighbours with values close to the query vector. This is due to Euclidean distance considers the magnitude of each dimension in the vector, opposed to Cosine which considers the direction. 

1. Starting with a struct with a mix of pre-normalized numerical values, weighted numerical values and categorical non-numerical values.

```go
	queryMedia := Media{
		Title:           "Action Movie", // Genre + Type, not actually used in query
		Type:            "Movie",        // Type of the media, either "Movie" or "Show"
		Rating:          4.0,            // Rating out of 5 (5 being the highest rating)
		Length:          0.6,            // Normalized length (e.g., 0 for short, 1 for long)
		Year:            0.8,            // Normalized release year (0 for older, 1 for recent)
		Popularity:      0.75,           // Normalized popularity (0 for least popular, 1 for most popular)
		CriticalAcclaim: 4.2,            // Critical acclaim out of 5 (5 being highly critically acclaimed)
		AudienceTarget:  0.5,            // Normalized target audience (e.g., 0 for kids, 1 for adults)
		Genre:           "Action",       // Genre of the show/movie
	}
```

2. Transform to vector to use to query, one-hot encoding the categorical values.

```go
Query Vector: [4 0.6 0.8 0.75 4.2 0.5 1 0 0 0 0 0 0 1]
```

3. Run a search against the database populated with pre-generated vectors.

```go
nearestNeighbours := db.kNN(queryVector.Vector, 5, euclideanDistance)
```

4. Output!

```go
Key: Action Movie 657, Vector: [4.5 0.3 0.9 0.9 4.2 0.2 1 0 0 0 0 0 0 1]
Key: Action Movie 936, Vector: [4 0.1 0.5 0.7 3.9 0.3 1 0 0 0 0 0 0 1]
Key: Action Movie 169, Vector: [4.3 1 0.7 0.6 3.8 0.7 1 0 0 0 0 0 0 1]
Key: Action Movie 711, Vector: [4.4 0.9 0.7 0.7 3.7 0.6 1 0 0 0 0 0 0 1]
Key: Action Movie 437, Vector: [4.1 0.9 0.9 0.2 4.7 0.9 1 0 0 0 0 0 0 1]
```

