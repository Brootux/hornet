package shutdown

const (
	ShutdownPriorityFlushToDatabase = iota
	ShutdownPriorityPersisters
	ShutdownPriorityRequestsProcessor
	ShutdownPriorityMilestoneSolidifier
	ShutdownPriorityMilestoneChecker
	ShutdownPrioritySolidifierGossip
	ShutdownPriorityReceiveTxWorker
	ShutdownPriorityReplyProcessor
	ShutdownPriorityBroadcastQueue
	ShutdownPriorityPacketProcessor
	ShutdownPriorityNeighborSendQueue
	ShutdownPriorityNeighbors
	ShutdownPriorityNeighborTCPServer
	ShutdownPriorityNeighborReconnecter
	ShutdownPriorityBadgerGarbageCollection
	ShutdownPriorityMetricsUpdater
	ShutdownPrioritySPA
	ShutdownPriorityAPI
	ShutdownPriorityMetricsPublishers
	ShutdownPrioritySpammer
	ShutdownPriorityStatusReport
)
