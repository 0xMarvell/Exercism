class DifferenceOfSquaresCalculator {

    int computeSquareOfSumTo(int input) {
        int sum = 0;
        for(int i=0; i <= input; i++) {
          sum += i;
        }
        return (int)Math.pow((double)sum, 2);
      }

      int computeSumOfSquaresTo(int input) {
        int sum = 0;
        for(int i=0; i <= input; i++) {
          sum += (int)Math.pow(i, 2);
        }
        return sum;
      }
      
      int computeDifferenceOfSquares(int input) {
        int sum = 0;
        int squareSum = 0;
        for(int i=0; i <= input; i++) {
          sum += i;
          squareSum += (int)Math.pow(i, 2);
        }
        return (int)Math.pow((double)sum, 2) - squareSum;
      }

}
