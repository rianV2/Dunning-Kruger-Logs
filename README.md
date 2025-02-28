# Dunning-Kruger-Logs
a cognitive bias that causes people to overestimate their knowledge or abilities

Data Structure
- HashTable = fast lookup, map in go, dict in python, where offer O(1)
- DoublyLinkedList = fast removal, fast insertion, use more memory need to store pointer of next and prev node
- LinkedList

Algorithm
- BFS - https://www.youtube.com/watch?v=6ZnyEApgFYg


Big O Notation
Time complexity
- focus on highest growth rate rather than exact execution counts
- O(NM) + O(M) = O(NM)
- if data result from O(NM), loop through data is same as O(NM), ignore constant factor, result to O(NM)
- The dominant term in Big-O notation is the largest one.
Space complexity
- init variable of loop O(n)
- https://www.youtube.com/watch?v=rHM3zWgnPVA


Distributed System
- Ambassador, imagine busy ceo with an assistent, proxy to handle logging/monitoring/retries
    - kubernetes envoy
- CQRS, split between read and write of database
    - read to use replica
    - write to use master
- Circuit Breaker, imagine pipe breaker to prevent all mesh, if has a service down we stop the breaker to avoid other service got down
- Event Sourcing, event status update history need to be recorded
- Leader Election, if leader node failed new leader will be elected
    - apache zookeeper, 
        - distributed locking, if master who have lock failed, it will transfer the lock to other (leader election)
        - Strong consistency (ensures only one process holds the lock)
    - etcd
- Pub/Sub, publisher of a topic, and subscriber of a topic
    - updating user profile, need to update user profile on each service
- Sharding, split set of data to different node reducing load on only single node
    - mongodb
    - cassandra
- Strangle Fig, replace strangle fig tree grow around tree and completely replacing them
    - migrating from legacy system to modern gradually rather than replace all system