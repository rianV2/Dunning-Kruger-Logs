# Dunning-Kruger-Logs
a cognitive bias that causes people to overestimate their knowledge or abilities

## Data Structure
- HashTable = fast lookup, map in go, dict in python, where offer O(1)
- DoublyLinkedList = fast removal, fast insertion, use more memory need to store pointer of next and prev node
- LinkedList
- Stack = FIFO first in first out, LIFO last in first out

## Algorithm
- BFS - https://www.youtube.com/watch?v=6ZnyEApgFYg


## Big O Notation
### Time complexity
- focus on highest growth rate rather than exact execution counts
- O(NM) + O(M) = O(NM)
- if data result from O(NM), loop through data is same as O(NM), ignore constant factor, result to O(NM)
- The dominant term in Big-O notation is the largest one.
### Space complexity
- init variable of loop O(n)
- https://www.youtube.com/watch?v=rHM3zWgnPVA


## Distributed System
### Pattern
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


### Database
![CAP Theorem](cap-theorem.png)
- https://www.youtube.com/watch?v=BHqjEjzAicA&t=1s
- https://www.youtube.com/watch?v=BTKBS_GdSms
- Ex: replica master communication
- CAP Theorem (Consistency, Availability, Partition Tolerance)
    - partition means disconnected tolerance
    - CP (Consistency + Partition Tolerance)
        - Always consistent, but may be unavailable during network failures.
        - MongoDB (single-node), HBase, Zookeeper
    - AP (Availability + Partition Tolerance)
        - Always available, but may return stale data (eventual consistency).
        - DynamoDB, Cassandra, CouchDB
    - CA (Consistency + Availability)
        - Not possible in a distributed system. 
        - A system that is both fully consistent and fully available cannot tolerate network failures. 
    - sometimes not always works, example we can allow read at 1 node but cannot write
- Connection Pooling
    - when we connect to db, we open connection after using it we close the connection
    - for connection pooling, we will keep the connection open for the other operation of operation, not immediately close it
    - this ensure, we don't frequently open and close the connection
- Eventually consistent
    - We cannot have strongly consistent database data, the database need a time to make it consistent
    - Example for pub/sub event, there will be eventually consistent data between service which publish and service which consume
- Indexing
    - need to make sure a query hit appropriate index, explain analyze
    - drawback for indexing is db need more time to write data, and update the index

### Load Balancer
- round robin, alternate between pods
- worker pool, pods will take other tasks if they finish


## System Design
- its behavioral question, easiest if you ever had experience

## Test
- unit test
- integration test
- snapshot test

## Behavioral
- rust vs golang
    - our payment service got bunch of issue, where data from our db and third party not sync
    - we've been working with go, new engineer come and try to build a new payment service with rust
    - i give my opinion that we're tight with resources, introducing new learning curve, and not sure if new joiner can learn it directly
    - meanwhile we disagree, actually he built the service underground, at first the service will function as source of truth
    - but then our pm suggest to just make rust project as our payment v2
    - then new feature and issue comes, but not everyone can touch it
    - we've been trying to move it again to go

## SOLID Principle
- https://www.linkedin.com/posts/elliotone_c-solid-principles-explianed-by-elliot-one-ugcPost-7293055382741020672-Uk3C?utm_source=share&utm_medium=member_desktop&rcm=ACoAABsEHr8B2VbhRjK2xfEDzzeVePNcT92-MFM

## OOP Paradigm
- Encaptulation
- Abstraction
- Inheritance
- Polymorphism

## LLM
Fine Tuning
- train existing model for more context accurate
- Need to provide input output, and good data
- Requires training
Prompt Engineering
- adjust ai prompt aim clear guidance for ai
RAG, Retrieval-Augmented Generation
- give ai a document knowledge
LangChain, 
- python frameworks for building RAG document index
- Building vector space by chunking the documents
Temparature
- It controls the randomness in the model's word choices.
- Lower = safer and more accurate (technical), higher = riskier and more creative (poem).
How to keep chat in context
- can pass previous chats history, or same prompt.
How to split personality
- every request should send same prompt personality guidance