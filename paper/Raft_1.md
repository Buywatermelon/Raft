### 1. Introduction
### 1. 介绍

Consensus algorithms allow a collection of machines to work as a coherent group that can survive the failures of some of its members.
共识算法允许一群机器像一个整体一样工作, 并且在集群中部门机器宕机时, 集群仍可以提供服务.
 
Because of this, they play a key role in building reliable large-scale software systems.
因如此, 他们在构建可靠的大型软件系统中扮演者至关重要的角色.
 
Paxos has dominated the discussion of consensus algorithms over the last decade: most implementations of consensus are based on Paxos or influenced by it, and Paxos has become the primary vehicle used to teach student about consensus.
Paxos在过去十年主导了关于共识算法的讨论: 大部分共识算法的实现都基于Paxos或深受其影响, 并且Paxos已经成为了教授学生共识相关的主要工具. 

--- 
 
Unfortunately, Paxos is quite difficult to understand, in spite of numerous attempts to make it more approachable.
不幸的是, Paxos非常的难以理解, 尽管有许多的尝试希望让Paxos更加的平易近人. 

Furthermore, its architecture requires complex changes to support practical system.
此外, Paxos的架构使得想要使用他构建实用的系统, 你需要考虑许多复杂的变化.

As a result, both system builders and students struggle with Paxos.
因此, 不论是系统架构师还是学生都难以理解Paxos.

---

After struggling with Paxos ourselves, we set out to find a new consensus algorithms that could provide a better foundation for system building and education. 
在与Paxos艰苦斗争后, 我们尝试去寻找一种更易于用其构建实际系统且更易于教育传播的共识算法.

Our approach was unusual in that our primary goal was understandability: could we define a consensus algorithms for practical systems and describe it in a way that is significantly easier to learn than Paxos? 
我们实现共识算法的方式不同于寻常, 因为我们的主要目标是可理解性: 我们能否为实际系统定义一个共识算法, 并以一种比Paxos更容易学习的方式描述他. 

Furthermore, we wanted that algorithm to facilitate the development of intuitions that are essential for system builders. It was important not just for the algorithm to work, but for it to be obvious why it works.
此外, 我们希望希望该算法能够促进对于系统架构师必不可少的直觉的发展, 比让算法跑起来更重要的是算法为什么是这样设计的.

---

The result of this work is a consensus algorithm called Raft.
这项工作的成果是一个名为Raft的共识算法.

In designing Raft we applied specific techniques to improve understandability, including decomposition (Raft separates leader election, log replication, and safety) and state space reduction (relative to Paxos, Raft reduces the degree of nondeterminism and the ways servers can be inconsistent with each other). 
在Raft的设计中, 我们使用了特别的方式去提高可理解性. 包括拆分 (Raft可以拆分为领导者选举, 日志复制和安全性) 还有状态空间的缩减 (相比于Paxos, Raft降低了节点状态的不确定性以及减少了节点间不一致性的场景).

A user study with 43 students at two universities shows that Raft is significantly easier to understand than Paxos: after learning both algorithms, 33 of these student were able to answer questions about Raft better than questions about Paxos.
一项对两所大学中43名学生的用户调查表明, Raft相比于Paxos显著的更易理解: 在学习了两种算法后, 相比于Paxos, 这些学生中的33位能够更好的回答出关于Raft的问题.

Raft is similar in many ways to existing consensus algorithms, but it has several novel features:
- **Strong leader**: Raft use a stronger form of leadership than other consensus algorithms. For example, log entries only flow from the leader to other servers. This simplifies the management of the replicated log and makes Raft easier to understand.
- **Leader election**: Raft user randomized timers to elect leaders. This adds only a small amount of mechanism to the heartbeats already required for any consensus algorithm, while resolving conflicts simply and rapidly.
- **Membership**: Raft's mechanism for changing the set of servers in the cluster use a new *joint consensus* approach where the majorities of two different configurations overlap during transitions. This allows the cluster to continue operating normally during configuration changes.
Raft与已知的共识算法是非常相似的, 但是他有几个新颖的功能:
- **Strong Leader**: Raft中的领导者相比于其他的共识算法更为强势. 例如, 日志条目只能从领导者同步到其他的节点. 这简化;了日志复制流程并是Raft更易于理解.
- **Leader election**: Raft使用了随机计时器去选举领导者. 这只为任何共识算法都需要的心跳添加了很少量的机制, 同时简单快速的解决了活锁问题.
- **Membership**: Raft使用了*Joint Consensus*的机制保证了在集群配置变更时, 两种不同的配置不会同时存在主节点. 这允许集群在配置变更期还能够正常的提供服务. 

---

We believe that Raft is superior to Paxos and other consensus algorithms, both for educational purposes and as a foundation for implementation.
我们相信Raft相比于Paxos和其他的共识算法, 不论在教育意义上, 还是作为实际系统实现的基础, 都是更加优越的.

It is simpler and more understandable than other algorithms; it is described completely enough to meet the needs of a practical system; it has several open-source implementations and it used by several companies; its safety properties have been formally specified and proven; and its efficiency is comparable to other algorithms.
相比于其他共识算法, Raft更简单且易于理解; 他被描述的足够清楚以满足实际系统的需要; 他有许多的开源实现并且被很多公司所采用; 他的安全性已经被严谨的定义和证明; 并且他的效率与其他共识算法相当.

---
The remainder of the paper introduces the replicated state machine problem (Section 2), discusses the strengths and weaknesses of Paxos (Section 3), describes our general approach to understandability (Section 4), presents the Raft consensus algorithm (Sections 5–8), evaluates Raft (Section 9), and discusses related work (Section 10).
论文的其余部分介绍了复制状态机的问题 (Section 2), 讨论Paxos的优缺点 (Section3), 描述了我们为了可理解性的设计 (Section 4), 提出了Raft共识算法 (Section 5-8), 对Raft算法的评估(Section 9), 和讨论相关的工作(Section 10).