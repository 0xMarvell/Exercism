public class CarsAssemble {

    public double productionRatePerHour(int speed) {
        double successRate = 1;
        final int CARS_PER_HOUR = 221;

        if (speed == 10) {
            successRate = 0.77;
            return (double)(speed * CARS_PER_HOUR * successRate);
        } else if (speed == 9) {
            successRate = 0.8;
            return (double)(speed * CARS_PER_HOUR * successRate);
        } else if ((speed >= 5) && (speed <= 8)) {
            successRate = 0.9;
            return (double)(speed * CARS_PER_HOUR * successRate);
        } else if ((speed >= 1) && (speed <= 4)){
            return (double)(speed * CARS_PER_HOUR * successRate);
        } else {
            return 0.0;
        }
    }

    public int workingItemsPerMinute(int speed) {
        return (int)(productionRatePerHour(speed) / 60);
    }
}
