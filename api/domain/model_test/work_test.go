package model_test

import (
	"influ-dojo/api/domain/model"
	"influ-dojo/api/usecase/input"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Work", func() {
	Describe("CalcPoint", func() {
		Context("when any value is less than base count", func() {
			It("returns correct range value", func() {
				work := &model.Work{
					MyTweetsCount:          3,
					RepliesCount:           7,
					IncreaseFavoritesCount: 300,
				}

				work.CalcPoint(input.TweetBase, input.RepBase, input.FavBase)

				Ω(work.Point).Should(Equal(39.9))
			})
		})

		Context("when tweet count is more than base count", func() {
			It("returns correct range value", func() {
				work := &model.Work{
					MyTweetsCount:          100,
					RepliesCount:           7,
					IncreaseFavoritesCount: 300,
				}

				work.CalcPoint(input.TweetBase, input.RepBase, input.FavBase)

				Ω(work.Point).Should(Equal(65.5))
			})
		})

		Context("when reply count is more than base count", func() {
			It("returns correct range value", func() {
				work := &model.Work{
					MyTweetsCount:          3,
					RepliesCount:           200,
					IncreaseFavoritesCount: 300,
				}

				work.CalcPoint(input.TweetBase, input.RepBase, input.FavBase)

				Ω(work.Point).Should(Equal(55.2))
			})
		})

		Context("when fav count is more than base count", func() {
			It("returns correct range value", func() {
				work := &model.Work{
					MyTweetsCount:          3,
					RepliesCount:           7,
					IncreaseFavoritesCount: 800,
				}

				work.CalcPoint(input.TweetBase, input.RepBase, input.FavBase)

				Ω(work.Point).Should(Equal(59.1))
			})
		})

		Context("when every count is more than base count", func() {
			It("returns correct range value", func() {
				work := &model.Work{
					MyTweetsCount:          8,
					RepliesCount:           15,
					IncreaseFavoritesCount: 800,
				}

				work.CalcPoint(input.TweetBase, input.RepBase, input.FavBase)

				Ω(work.Point).Should(Equal(100.))
			})
		})

		Context("when any count is less than 0", func() {
			It("returns correct range value", func() {
				work := &model.Work{
					MyTweetsCount:          -4,
					RepliesCount:           -6,
					IncreaseFavoritesCount: -200,
				}

				work.CalcPoint(input.TweetBase, input.RepBase, input.FavBase)

				Ω(work.Point).Should(Equal(0.))
			})
		})
	})
})
