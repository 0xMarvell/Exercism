
class BirdWatcher {
    private final int[] birdsPerDay;

    public BirdWatcher(int[] birdsPerDay) {
        this.birdsPerDay = birdsPerDay.clone();
    }

    public int[] getLastWeek() {
        int[] birds = {0, 2, 5, 3, 7, 8, 4};
        return birds;
    }

    public int getToday() {
        if (birdsPerDay.length == 0) {
            return 0;
        }
        return birdsPerDay[birdsPerDay.length-1];
    }

    public void incrementTodaysCount() {
        birdsPerDay[birdsPerDay.length-1]++;
    }

    public boolean hasDayWithoutBirds() {
        for (int bird : birdsPerDay){
            return bird == 0;
        }
        return false;
    }

    public int getCountForFirstDays(int numberOfDays) {
        int count = 0;

        if (numberOfDays > 7 ) {
            numberOfDays = 7;
        }

        for (int i = 0; i != numberOfDays; i++) {
            count+=birdsPerDay[i];
        }

        return count;
    }

    public int getBusyDays() {
        int busyDaysCount = 0;

        for (int i = 0; i < birdsPerDay.length; i++) {
            if (birdsPerDay[i] >= 5) {
                busyDaysCount++;
            } 
        }

     
       return busyDaysCount;
    }
}
