package main

import (
	"sync"
	"time"
)

func main() {

}

const (
	Follower = iota
	Candidate
	Leader
)

type Raft struct {
	mu sync.Mutex

	// 服务器id
	id int
	// 服务器当前状态
	state int
	// 服务器已知的最新任期
	currentTerm int
	// 当前任期内投票
	votedFor int
	// 心跳时间
	heartbeatTime time.Time
	// 状态机日志
	log []LogEntry
	// 当前commit的日志条目序号
	commitIndex int
}

func (rf *Raft) setState(state int) {
	rf.state = state
}

func (rf *Raft) apply() {

}

type RequestVoteArgs struct {
	// 候选者任期
	Term int
	// 候选者 id
	CandidateId int
}

type RequestVoteReply struct {
	// 处理请求节点的任期号，用于候选者更新自己的任期
	Term int
	// 候选者获得选票时为true; 否则为false
	VoteGranted bool
}

type LogEntry struct {
	// 索引，表示该日志在整个日志中的位置
	Index int
	// 任期号，日志条目首次被领导者创建时的任期
	Term int
	// 命令，应用于状态机的命令
	Command interface{}
}

type AppendEntriesArgs struct {
	Term         int
	LeaderId     int
	PrevLogIndex int
	PrevLogTerm  int
	// 需要复制的日志条目，用于发送心跳消息时Entries为空
	Entries []LogEntry
	// 领导者已提交的最大的日志索引，用于跟随者提交
	LeaderCommit int
}

type AppendEntriesReply struct {
	Term    int
	Success bool
}

func (rf *Raft) RequestVote(args *RequestVoteArgs, reply *RequestVoteReply) {
	rf.mu.Lock()
	defer rf.mu.Unlock()

	reply.Term = rf.currentTerm
	reply.VoteGranted = false
	if args.Term < rf.currentTerm {
		return
	}

	// 如果收到来自更大任期的请求，则更新自己的 currentTerm，转为跟随者
	if reply.Term > rf.currentTerm {
		rf.currentTerm = reply.Term
		rf.state = Follower
		rf.votedFor = -1
	}

	if rf.votedFor == -1 || rf.votedFor == args.CandidateId {
		rf.votedFor = args.CandidateId
		reply.VoteGranted = true
		rf.heartbeatTime = time.Now()
	}
	return
}

func (rf *Raft) AppendEntries(args *AppendEntriesArgs, reply *AppendEntriesReply) {
	rf.mu.Lock()
	defer rf.mu.Unlock()

	reply.Term = rf.currentTerm
	reply.Success = false
	if args.Term < rf.currentTerm {
		return
	}

	if args.Term > rf.currentTerm {
		reply.Term = args.Term
		rf.currentTerm = args.Term
	}

	// 主要为了重置选举超时时间
	rf.setState(Follower)

	// 日志一致性检查
	// lastLogIndex := rf.getLastIndex()
	lastLogIndex := len(rf.log) - 1
	if args.PrevLogIndex > rf.log[lastLogIndex].Index ||
		rf.log[args.PrevLogIndex].Term != args.PrevLogTerm {
		return
	}

	reply.Success = true

	// 需要处理重复的RPC请求
	// 比较日志条目的任期，以确认是否能够安全地追加日志
	// 否则会导致重复应用命令
	index := args.PrevLogIndex
	for i, entry := range args.Entries {
		index++
		if index < len(rf.log) {
			if rf.log[index].Term == entry.Term {
				continue
			}

			rf.log = rf.log[:index]
		}

		rf.log = append(rf.log, args.Entries[i:]...)
	}

	if rf.commitIndex < args.LeaderCommit {
		lastLogIndex = rf.log[len(rf.log)-1].Index
		if args.LeaderCommit > lastLogIndex {
			rf.commitIndex = lastLogIndex
		} else {
			rf.commitIndex = args.LeaderCommit
		}
		// 将命令应用到自己的状态机，不同的应用有不同的实现
		rf.apply()
	}

	// 保险起见，再重置一次选举超时时间
	rf.setState(Follower)
	return
}
