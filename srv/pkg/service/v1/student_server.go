package v1

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"srv/pkg/api/v1"
	"time"
)

func check(e error) {
	if e != nil {
		fmt.Print("Error Happened{\n")
		fmt.Print(e)
		fmt.Print("\n}")
		//panic(e)
	}
}
func errorStream(e error, subscribe *v1.Subscribe, s *studentGradeServer, t *v1.Task) bool {
	if e != nil {
		log.Printf("Lost connection to %s", subscribe.User)
		log.Printf("Recover by sending back last message to channel")
		s.students[subscribe.User] <- *t
		return true
	}
	return false
}

// studentGradeServer is implementation of v1.StudentService proto interface
type studentGradeServer struct {
	labs         map[string][]string
	students     map[string]chan v1.Task
	tasksToGrade map[string]chan v1.GradeSolution
}

func (s *studentGradeServer) TESTCreateAssignment(ctx context.Context, task *v1.Task) (*v1.Empty, error) {
	//get all labs
	allLabs := make([]string, 0, len(s.labs))
	for k := range s.labs {
		allLabs = append(allLabs, k)
	}
	rand.Seed(time.Now().Unix())
	if len(allLabs) == 0{
		log.Printf("Tried to assign a new task with no labs available")
		return &v1.Empty{}, nil
	}
	labIndex := rand.Intn(len(allLabs))

	pickedLab := s.labs[allLabs[labIndex]]
	log.Printf("Creating new assignment")
	log.Printf("Randomly picked lab %s", allLabs[labIndex])
	log.Printf("Students signed for lab %s", pickedLab)

	for _, student := range pickedLab {
		s.students[student] <- *task
	}

	log.Print("Sent message to all students")

	return &v1.Empty{}, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (s *studentGradeServer) SignForLab(subscribe *v1.Subscribe, stream v1.StudentService_SignForLabServer) error {
	log.Printf("Requested to create user %s", subscribe.User)
	if s.students[subscribe.User] == nil {
		s.students[subscribe.User] = make(chan v1.Task, 1000)
	}

	for _, lab := range subscribe.Lab {
		if s.labs[lab] == nil {
			s.labs[lab] = []string{}
		}
		if !contains(s.labs[lab], subscribe.User) {
			s.labs[lab] = append(s.labs[lab], subscribe.User)
		}
	}

	for {
		var t = <-s.students[subscribe.User]
		log.Printf("i, %s received %+v", subscribe.User, t)
		//_ = stream.Send(&v1.Task{Type: 0, Deadline: 6969, Descriptions: []string{"text", "dolny"}})
		err := stream.Send(&t)
		if errorStream(err, subscribe, s, &t) {
			break
		}
		time.Sleep(time.Second)
	}

	log.Printf("returned, probably client died")
	return nil
}

func (s *studentGradeServer) SendSolution(solution *v1.Solution, server v1.StudentService_SendSolutionServer) error {
	log.Print("received solution from ", solution.Id, solution.Comment)
	log.Print("solution saved to: .")

	for index, solutionFile := range solution.Solution {
		f, err := os.Create(solution.Id + string(index))
		check(err)
		noOfBytes, err := f.Write(solutionFile)
		check(err)
		log.Printf("created solution for file %s, size: %d", solution.Id+string(index), noOfBytes)
	}

	s.tasksToGrade[solution.Id] = make(chan v1.GradeSolution, 1000)
	grade := <-s.tasksToGrade[solution.Id]

	err := server.Send(&grade)
	check(err)
	return nil
}

func NewStudentServiceServer() v1.StudentServiceServer {
	return &studentGradeServer{
		labs:         make(map[string][]string),
		students:     make(map[string]chan v1.Task),
		tasksToGrade: make(map[string]chan v1.GradeSolution),
	}
}
