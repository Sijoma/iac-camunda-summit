package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/camunda/zeebe/clients/go/v8/pkg/entities"
	"github.com/camunda/zeebe/clients/go/v8/pkg/pb"
	"github.com/camunda/zeebe/clients/go/v8/pkg/worker"
	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"
)

func main() {
	zbClient, err := zbc.NewClient(&zbc.ClientConfig{
		GatewayAddress: os.Getenv("ZEEBE_ADDRESS"),
	})

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	topology, err := zbClient.NewTopologyCommand().Send(ctx)
	if err != nil {
		panic(err)
	}

	for _, broker := range topology.Brokers {
		fmt.Println("Broker", broker.Host, ":", broker.Port)
		for _, partition := range broker.Partitions {
			fmt.Println("  Partition", partition.PartitionId, ":", roleToString(partition.Role))
		}
	}

	files, err := os.ReadDir("/usr/bpmn")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".bpmn") {
			process, err := zbClient.NewDeployResourceCommand().
				AddResourceFile("/usr/bpmn/" + file.Name()).
				Send(ctx)
			if err != nil {
				log.Println("failed to deploy process", err)
				return
			}
			fmt.Print("Deploy process 🆕")
			fmt.Println(process.Key, process.Deployments)
		}
	}

	jobWorker := zbClient.NewJobWorker().JobType("generate-text").Handler(handleJob).Open()
	jobWorker.AwaitClose()
}

func handleJob(client worker.JobClient, job entities.Job) {
	fmt.Println(job.GetBpmnProcessId())

	variables := map[string]interface{}{}

	request, err := client.NewCompleteJobCommand().JobKey(job.Key).VariablesFromMap(variables)
	if err != nil {
		failJob(client, job)
		return
	}

	ctx := context.Background()
	_, err = request.Send(ctx)
	if err != nil {
		panic(err)
	}

	log.Println("Successfully completed job")
}

func roleToString(role pb.Partition_PartitionBrokerRole) string {
	switch role {
	case pb.Partition_LEADER:
		return "Leader"
	case pb.Partition_FOLLOWER:
		return "Follower"
	default:
		return "Unknown"
	}
}

func failJob(client worker.JobClient, job entities.Job) {
	log.Println("Failed to complete job", job.GetKey())

	ctx := context.Background()
	_, err := client.NewFailJobCommand().JobKey(job.GetKey()).Retries(job.Retries - 1).Send(ctx)
	if err != nil {
		panic(err)
	}
}
