public class Lasagna {
    // define the 'expectedMinutesInOven()' method
    public int expectedMinutesInOven(){
        return 40;
    }

    // define the 'remainingMinutesInOven()' method
    public int remainingMinutesInOven(int actualMinutes){
        int expectedMins = new Lasagna().expectedMinutesInOven();
        return expectedMins - actualMinutes;
    }

    // define the 'preparationTimeInMinutes()' method
    public int preparationTimeInMinutes(int numOfLayers){
        final int SINGLE_LAYER_PREP_TIME = 2;
        return numOfLayers * SINGLE_LAYER_PREP_TIME;
    }

    // define the 'totalTimeInMinutes()' method
    public int totalTimeInMinutes(int numOfLayers,int timeInOven){
        int prepTime = new Lasagna().preparationTimeInMinutes(numOfLayers);
        return prepTime + timeInOven;
    }
}
