
# **KDTree**

-   **What are K-d Trees?**
    
    K-d Trees are space-partitioning data structures for organizing points in a k-dimensional space. They are particularly useful in applications involving multidimensional key searches.
    
-   **Basic Concepts**
    
    -   **K-dimensional Space**: A space with k dimensions, where each point can be represented as a k-dimensional vector.
    -   **Space Partitioning**: Dividing a space into distinct regions or partitions.
-   **K-d Trees Structure**
    
    A K-d Tree is a binary tree where each node represents an axis-aligned hyperplane that divides the space into two parts.
    
-   **Properties**
    
    1.  **Balanced Tree**: Ideally, a K-d Tree is balanced to ensure efficient operations.
    2.  **Leaf and Internal Nodes**: Internal nodes represent partitioning hyperplanes, while leaf nodes represent data points.
-   **Creating a K-d Tree**
    
    The process involves recursively dividing the set of points into halves until each partition forms a leaf node.
    
-   **Algorithm**
    
    1.  Select a dimension.
    2.  Find the median element in that dimension.
    3.  Create a node and divide the set into two subsets.
    4.  Recursively apply the steps for each subset.



-   **Applications of K-d Trees**
    
    **K-d Trees** find applications in various fields:
    
    1.  **Computer Graphics**: For efficient rendering and ray tracing.
    2.  **Geospatial Data**: To query multidimensional data like longitude and latitude.
    3.  **Machine Learning**: In nearest neighbor searches for clustering algorithms.
-   **Pros of K-d Trees**
    
    1.  **Efficient Search Operations**: K-d trees provide efficient search operations, especially for spatial data. They allow for fast nearest neighbor searches, which are crucial in many applications like computer graphics, machine learning, and geographic information systems.
    2.  **Balanced Partitioning**: They divide the space into two halves at each level, which helps in evenly distributing the points. This balanced partitioning can be very effective for range and nearest neighbor searches.
    3.  **Dimensionality Flexibility**: K-d trees can handle data with multiple dimensions effectively, making them versatile for various applications.
    4.  **Space Efficiency:** They require minimal space overhead, as they store only the data points and the splitting planes.
    5.  **Incremental Construction**: K-d trees can be built incrementally, which is beneficial when dealing with streaming data or when the dataset is too large to be loaded into memory all at once.
-   **Cons of K-d Trees**
    
    1.  **Degraded Performance in High Dimensions**: The efficiency of K-d trees diminishes as the number of dimensions increases (a phenomenon known as the "curse of dimensionality"). In very high-dimensional spaces, the performance of K-d trees may become comparable to brute force search.
    2.  **Costly Rebalancing Operations**: While K-d trees can be built incrementally, rebalancing the tree (to maintain efficient search performance) can be costly in terms of computation, especially for dynamic datasets.
    3.  **Not Ideal for Skewed Data**: If the data is heavily skewed or clustered in certain regions of the space, the performance of K-d trees can be adversely affected.
    4.  **Complexity in Deletion Operations**: Deleting points from a K-d tree while maintaining its properties can be complex and less efficient compared to other data structures like R-trees.
    5.  **Non-Uniform Data Handling Issues**: K-d trees may not perform well when the data points are not uniformly distributed across the dimensions, as the splitting criteria might lead to inefficient partitioning of the space.
    
    


# Exercise

Your task is to implement these 3 methods:-

- `func (t *KDTree[T]) SearchNearest(target KDPoint[T]) KDPoint[T]`

- `func (t *KDTree[T]) Insert(p KDPoint[T]) `

- `func buildTree[T any](points []KDPoint[T], dstFn KDistanceCalculator[T]) *Node[T]`

and pass all the tests