### Abstract  
### 摘要  
  
Raft is a consensus algorithm for managing a replicated log.  
Raft是一种管理复制日志的共识算法。

It produces a result equivalent to (multi-)Paxos, and it is as efficient as Paxos, but its structure is different from Paxos; this makes Raft more understandable than Paxos and also provides a better foundation for building practical systems.    
他与(multi-)Paxos产生相同的结果且同样高效，但他们的结构不同; 相比于Paxos，Raft更易于理解且更于构建出实际的系统。

In order to enhance understandability, Raft separates the key elements of consensus,   
such as leader election, log replication, and safety, and it enforces a stronger degree of coherency to reduce the number of states that must be considered.
为了更易于理解，Raft将共识算法的关键要素分开罗列，例如领导者选举、日志复制、安全性、通过施加更强的一致性保障让你减少对于节点状态的考虑。

Results from a user study demonstrate that Raft is easier for students to learn than Paxos.  
Raft also includes a new mechanism for changing the cluster membership, which use overlapping majorities to guarantee safety.  
用户的研究结果表明，Raft相比Paxos更加易于学习。Raft还包括了一个新的通过多数选举以保障集群成员变更安全性的机制。