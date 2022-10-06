class DifferenceOfSquaresCalculator {

    int computeSquareOfSumTo(int input) {
        int sum = 0;

        for (int i = 0; i <= input; i++) {
            sum += i;
        }

        return sum * sum;
    }

    int computeSumOfSquaresTo(int input) {
        int sumSuqared = 0;

        for (int i = 0; i < input; i++) {
            sumSuqared += i*i;
        }

        return sumSuqared;
    }

    int computeDifferenceOfSquares(int input) {
        return computeSquareOfSumTo(input) - computeSumOfSquaresTo(input);
    }

}
