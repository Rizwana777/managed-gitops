package metrics

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

var _ = Describe("Test for Operation DB metrics counter", func() {
	Context("Prometheus metrics responds to count of operation DB rows", func() {
		It("tests SetTotalCountOfOperationDBRows, SetCountOfOperationDBRows, SetCountOfOperationDBRowsInCompleteAndNonCompleteState function on operation DB rows", func() {

			ClearDBMetrics()

			totalNumberOfOperationDBRows := testutil.ToFloat64(OperationDBRows)
			numberOfOperationDBRowsInWaitingState := testutil.ToFloat64(OperationDBRowsInWaitingState)
			numberOfOperationDBRowsIn_InProgressState := testutil.ToFloat64(OperationDBRowsIn_InProgressState)
			numberOfOperationDBRowsInCompletedState := testutil.ToFloat64(OperationDBRowsInCompletedState)
			numberOfOperationDBRowsInFailedState := testutil.ToFloat64(OperationDBRowsInErrorState)
			totalNumberOfOperationDBRowsInCompletedState := testutil.ToFloat64(TotalOperationDBRowsInCompletedState)
			totalNumberOfOperationDBRowsInNonCompleteState := testutil.ToFloat64(TotalOperationDBRowsInNonCompleteState)

			By("verify SetTotalCountOfOperationDBRows function by passing count as 4")
			SetTotalCountOfOperationDBRows(4)
			newTotalNumberOfOperationDBRows := testutil.ToFloat64(OperationDBRows)
			newNumberOfOperationDBRowsInWaitingState := testutil.ToFloat64(OperationDBRowsInWaitingState)
			newNumberOfOperationDBRowsIn_InProgressState := testutil.ToFloat64(OperationDBRowsIn_InProgressState)
			newNumberOfOperationDBRowsInCompletedState := testutil.ToFloat64(OperationDBRowsInCompletedState)
			newNumberOfOperationDBRowsInFailedState := testutil.ToFloat64(OperationDBRowsInErrorState)
			newTotalNumberOfOperationDBRowsInCompletedState := testutil.ToFloat64(TotalOperationDBRowsInCompletedState)
			newTotalNumberOfOperationDBRowsInNonCompleteState := testutil.ToFloat64(TotalOperationDBRowsInNonCompleteState)

			Expect(newTotalNumberOfOperationDBRows).To(Equal(totalNumberOfOperationDBRows + 4))
			Expect(newNumberOfOperationDBRowsInWaitingState).To(Equal(numberOfOperationDBRowsInWaitingState))
			Expect(newNumberOfOperationDBRowsIn_InProgressState).To(Equal(numberOfOperationDBRowsIn_InProgressState))
			Expect(newNumberOfOperationDBRowsInCompletedState).To(Equal(numberOfOperationDBRowsInCompletedState))
			Expect(newNumberOfOperationDBRowsInFailedState).To(Equal(numberOfOperationDBRowsInFailedState))
			Expect(newTotalNumberOfOperationDBRowsInCompletedState).To(Equal(totalNumberOfOperationDBRowsInCompletedState))
			Expect(newTotalNumberOfOperationDBRowsInNonCompleteState).To(Equal(totalNumberOfOperationDBRowsInNonCompleteState))

			By("verify SetCountOfOperationDBRows function by passing state as 'Waiting' and count as 2")
			SetCountOfOperationDBRows("Waiting", 2)
			newTotalNumberOfOperationDBRows = testutil.ToFloat64(OperationDBRows)
			newNumberOfOperationDBRowsInWaitingState = testutil.ToFloat64(OperationDBRowsInWaitingState)
			newNumberOfOperationDBRowsIn_InProgressState = testutil.ToFloat64(OperationDBRowsIn_InProgressState)
			newNumberOfOperationDBRowsInCompletedState = testutil.ToFloat64(OperationDBRowsInCompletedState)
			newNumberOfOperationDBRowsInFailedState = testutil.ToFloat64(OperationDBRowsInErrorState)
			newTotalNumberOfOperationDBRowsInCompletedState = testutil.ToFloat64(TotalOperationDBRowsInCompletedState)
			newTotalNumberOfOperationDBRowsInNonCompleteState = testutil.ToFloat64(TotalOperationDBRowsInNonCompleteState)

			Expect(newTotalNumberOfOperationDBRows).To(Equal(totalNumberOfOperationDBRows + 4))
			Expect(newNumberOfOperationDBRowsInWaitingState).To(Equal(numberOfOperationDBRowsInWaitingState + 2))
			Expect(newNumberOfOperationDBRowsIn_InProgressState).To(Equal(numberOfOperationDBRowsIn_InProgressState))
			Expect(newNumberOfOperationDBRowsInCompletedState).To(Equal(numberOfOperationDBRowsInCompletedState))
			Expect(newNumberOfOperationDBRowsInFailedState).To(Equal(numberOfOperationDBRowsInFailedState))
			Expect(newTotalNumberOfOperationDBRowsInCompletedState).To(Equal(totalNumberOfOperationDBRowsInCompletedState))
			Expect(newTotalNumberOfOperationDBRowsInNonCompleteState).To(Equal(totalNumberOfOperationDBRowsInNonCompleteState))

			By("verify SetCountOfOperationDBRows function by passing state as 'In_Progress' and count as 2")
			SetCountOfOperationDBRows("In_Progress", 2)
			newTotalNumberOfOperationDBRows = testutil.ToFloat64(OperationDBRows)
			newNumberOfOperationDBRowsInWaitingState = testutil.ToFloat64(OperationDBRowsInWaitingState)
			newNumberOfOperationDBRowsIn_InProgressState = testutil.ToFloat64(OperationDBRowsIn_InProgressState)
			newNumberOfOperationDBRowsInCompletedState = testutil.ToFloat64(OperationDBRowsInCompletedState)
			newNumberOfOperationDBRowsInFailedState = testutil.ToFloat64(OperationDBRowsInErrorState)
			newTotalNumberOfOperationDBRowsInCompletedState = testutil.ToFloat64(TotalOperationDBRowsInCompletedState)
			newTotalNumberOfOperationDBRowsInNonCompleteState = testutil.ToFloat64(TotalOperationDBRowsInNonCompleteState)

			Expect(newTotalNumberOfOperationDBRows).To(Equal(totalNumberOfOperationDBRows + 4))
			Expect(newNumberOfOperationDBRowsInWaitingState).To(Equal(numberOfOperationDBRowsInWaitingState + 2))
			Expect(newNumberOfOperationDBRowsIn_InProgressState).To(Equal(numberOfOperationDBRowsIn_InProgressState + 2))
			Expect(newNumberOfOperationDBRowsInCompletedState).To(Equal(numberOfOperationDBRowsInCompletedState))
			Expect(newNumberOfOperationDBRowsInFailedState).To(Equal(numberOfOperationDBRowsInFailedState))
			Expect(newTotalNumberOfOperationDBRowsInCompletedState).To(Equal(totalNumberOfOperationDBRowsInCompletedState))
			Expect(newTotalNumberOfOperationDBRowsInNonCompleteState).To(Equal(totalNumberOfOperationDBRowsInNonCompleteState))

			By("verify SetCountOfOperationDBRows function by passing state as 'Completed' and count as 2")
			SetCountOfOperationDBRows("Completed", 2)
			newTotalNumberOfOperationDBRows = testutil.ToFloat64(OperationDBRows)
			newNumberOfOperationDBRowsInWaitingState = testutil.ToFloat64(OperationDBRowsInWaitingState)
			newNumberOfOperationDBRowsIn_InProgressState = testutil.ToFloat64(OperationDBRowsIn_InProgressState)
			newNumberOfOperationDBRowsInCompletedState = testutil.ToFloat64(OperationDBRowsInCompletedState)
			newNumberOfOperationDBRowsInFailedState = testutil.ToFloat64(OperationDBRowsInErrorState)
			newTotalNumberOfOperationDBRowsInCompletedState = testutil.ToFloat64(TotalOperationDBRowsInCompletedState)
			newTotalNumberOfOperationDBRowsInNonCompleteState = testutil.ToFloat64(TotalOperationDBRowsInNonCompleteState)

			Expect(newTotalNumberOfOperationDBRows).To(Equal(totalNumberOfOperationDBRows + 4))
			Expect(newNumberOfOperationDBRowsInWaitingState).To(Equal(numberOfOperationDBRowsInWaitingState + 2))
			Expect(newNumberOfOperationDBRowsIn_InProgressState).To(Equal(numberOfOperationDBRowsIn_InProgressState + 2))
			Expect(newNumberOfOperationDBRowsInCompletedState).To(Equal(numberOfOperationDBRowsInCompletedState + 2))
			Expect(newNumberOfOperationDBRowsInFailedState).To(Equal(numberOfOperationDBRowsInFailedState))
			Expect(newTotalNumberOfOperationDBRowsInCompletedState).To(Equal(totalNumberOfOperationDBRowsInCompletedState))
			Expect(newTotalNumberOfOperationDBRowsInNonCompleteState).To(Equal(totalNumberOfOperationDBRowsInNonCompleteState))

			By("verify SetCountOfOperationDBRows function by passing state as 'Failed' and count as 2")
			SetCountOfOperationDBRows("Failed", 2)
			newTotalNumberOfOperationDBRows = testutil.ToFloat64(OperationDBRows)
			newNumberOfOperationDBRowsInWaitingState = testutil.ToFloat64(OperationDBRowsInWaitingState)
			newNumberOfOperationDBRowsIn_InProgressState = testutil.ToFloat64(OperationDBRowsIn_InProgressState)
			newNumberOfOperationDBRowsInCompletedState = testutil.ToFloat64(OperationDBRowsInCompletedState)
			newNumberOfOperationDBRowsInFailedState = testutil.ToFloat64(OperationDBRowsInErrorState)
			newTotalNumberOfOperationDBRowsInCompletedState = testutil.ToFloat64(TotalOperationDBRowsInCompletedState)
			newTotalNumberOfOperationDBRowsInNonCompleteState = testutil.ToFloat64(TotalOperationDBRowsInNonCompleteState)

			Expect(newTotalNumberOfOperationDBRows).To(Equal(totalNumberOfOperationDBRows + 4))
			Expect(newNumberOfOperationDBRowsInWaitingState).To(Equal(numberOfOperationDBRowsInWaitingState + 2))
			Expect(newNumberOfOperationDBRowsIn_InProgressState).To(Equal(numberOfOperationDBRowsIn_InProgressState + 2))
			Expect(newNumberOfOperationDBRowsInCompletedState).To(Equal(numberOfOperationDBRowsInCompletedState + 2))
			Expect(newNumberOfOperationDBRowsInFailedState).To(Equal(numberOfOperationDBRowsInFailedState + 2))
			Expect(newTotalNumberOfOperationDBRowsInCompletedState).To(Equal(totalNumberOfOperationDBRowsInCompletedState))
			Expect(newTotalNumberOfOperationDBRowsInNonCompleteState).To(Equal(totalNumberOfOperationDBRowsInNonCompleteState))

			By("verify SetCountOfOperationDBRowsInCompleteAndNonCompleteState function for operation DB row with Completed state")
			SetCountOfOperationDBRowsInCompleteAndNonCompleteState(true, 2)
			newTotalNumberOfOperationDBRows = testutil.ToFloat64(OperationDBRows)
			newNumberOfOperationDBRowsInWaitingState = testutil.ToFloat64(OperationDBRowsInWaitingState)
			newNumberOfOperationDBRowsIn_InProgressState = testutil.ToFloat64(OperationDBRowsIn_InProgressState)
			newNumberOfOperationDBRowsInCompletedState = testutil.ToFloat64(OperationDBRowsInCompletedState)
			newNumberOfOperationDBRowsInFailedState = testutil.ToFloat64(OperationDBRowsInErrorState)
			newTotalNumberOfOperationDBRowsInCompletedState = testutil.ToFloat64(TotalOperationDBRowsInCompletedState)
			newTotalNumberOfOperationDBRowsInNonCompleteState = testutil.ToFloat64(TotalOperationDBRowsInNonCompleteState)

			Expect(newTotalNumberOfOperationDBRows).To(Equal(totalNumberOfOperationDBRows + 4))
			Expect(newNumberOfOperationDBRowsInWaitingState).To(Equal(numberOfOperationDBRowsInWaitingState + 2))
			Expect(newNumberOfOperationDBRowsIn_InProgressState).To(Equal(numberOfOperationDBRowsIn_InProgressState + 2))
			Expect(newNumberOfOperationDBRowsInCompletedState).To(Equal(numberOfOperationDBRowsInCompletedState + 2))
			Expect(newNumberOfOperationDBRowsInFailedState).To(Equal(numberOfOperationDBRowsInFailedState + 2))
			Expect(newTotalNumberOfOperationDBRowsInCompletedState).To(Equal(totalNumberOfOperationDBRowsInCompletedState + 2))
			Expect(newTotalNumberOfOperationDBRowsInNonCompleteState).To(Equal(totalNumberOfOperationDBRowsInNonCompleteState))

			By("verify SetCountOfOperationDBRowsInCompleteAndNonCompleteState function for operation DB row with Non-Complete state")
			SetCountOfOperationDBRowsInCompleteAndNonCompleteState(false, 2)
			newTotalNumberOfOperationDBRows = testutil.ToFloat64(OperationDBRows)
			newNumberOfOperationDBRowsInWaitingState = testutil.ToFloat64(OperationDBRowsInWaitingState)
			newNumberOfOperationDBRowsIn_InProgressState = testutil.ToFloat64(OperationDBRowsIn_InProgressState)
			newNumberOfOperationDBRowsInCompletedState = testutil.ToFloat64(OperationDBRowsInCompletedState)
			newNumberOfOperationDBRowsInFailedState = testutil.ToFloat64(OperationDBRowsInErrorState)
			newTotalNumberOfOperationDBRowsInCompletedState = testutil.ToFloat64(TotalOperationDBRowsInCompletedState)
			newTotalNumberOfOperationDBRowsInNonCompleteState = testutil.ToFloat64(TotalOperationDBRowsInNonCompleteState)

			Expect(newTotalNumberOfOperationDBRows).To(Equal(totalNumberOfOperationDBRows + 4))
			Expect(newNumberOfOperationDBRowsInWaitingState).To(Equal(numberOfOperationDBRowsInWaitingState + 2))
			Expect(newNumberOfOperationDBRowsIn_InProgressState).To(Equal(numberOfOperationDBRowsIn_InProgressState + 2))
			Expect(newNumberOfOperationDBRowsInCompletedState).To(Equal(numberOfOperationDBRowsInCompletedState + 2))
			Expect(newNumberOfOperationDBRowsInFailedState).To(Equal(numberOfOperationDBRowsInFailedState + 2))
			Expect(newTotalNumberOfOperationDBRowsInCompletedState).To(Equal(totalNumberOfOperationDBRowsInCompletedState + 2))
			Expect(newTotalNumberOfOperationDBRowsInNonCompleteState).To(Equal(totalNumberOfOperationDBRowsInNonCompleteState + 2))
		})
	})
})
