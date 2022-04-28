package mysql

import "github.com/LearnHub/learning-go/pkg/learning"

type LearnMySQL struct {
	learning.Learning
}

func (learner LearnMySQL) Learn() {

}

func (learner LearnMySQL) Kind() string {
	return "mysql"
}
