package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/patrikeh/go-deep/training"

	deep "github.com/patrikeh/go-deep"
)

/*
mnist classifier
mnist is a set of hand-written digits 0-9
the dataset in a sane format (as used here) can be found at:
https://pjreddie.com/projects/mnist-in-csv/
*/
func main() {
	rand.Seed(time.Now().UnixNano())

	train, err := load("./mnist_train.data")
	if err != nil {
		panic(err)
	}
	test, err := load("./mnist_test.data")
	if err != nil {
		panic(err)
	}

	for i := range train {
		for j := range train[i].Input {
			train[i].Input[j] = train[i].Input[j] / 255
		}
	}
	for i := range test {
		for j := range test[i].Input {
			test[i].Input[j] = test[i].Input[j] / 255
		}
	}
	test.Shuffle()
	train.Shuffle()

	neural := deep.NewNeural(&deep.Config{
		Inputs:     len(train[0].Input),
		Layout:     []int{50, 10},
		Activation: deep.ActivationReLU,
		Mode:       deep.ModeMultiClass,
		Weight:     deep.NewNormal(0.6, 0.1), // slight positive bias helps ReLU
		Bias:       true,
	})

	//trainer := training.NewTrainer(training.NewSGD(0.01, 0.5, 1e-6, true), 1)
	trainer := training.NewBatchTrainer(training.NewAdam(0.02, 0.9, 0.999, 1e-8), 1, 200, 8)

	fmt.Printf("training: %d, val: %d, test: %d\n", len(train), len(test), len(test))

	trainer.Train(neural, train, test, 500)
}

func load(path string) (training.Examples, error) {
	_ = "STUB: not implemented"
	return *new(training.Examples), nil
}

func toExample(in []string) training.Example {
	_ = "STUB: not implemented"
	return *new(training.Example)
}

func onehot(classes int, val float64) []float64 { _ = "STUB: not implemented"; return nil }
