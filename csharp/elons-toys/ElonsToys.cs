using System;

class RemoteControlCar
{
    public int Distance = 0;
    public int BatteryPercentage = 100;
    public static RemoteControlCar Buy()
    {
        RemoteControlCar newCar = new();
        return newCar;
    }

    public string DistanceDisplay()
    {
        return $"Driven {Distance} meters";
    }

    public string BatteryDisplay()
    {
        if (BatteryPercentage < 1)
        {
            return "Battery empty";
        }
        return $"Battery at {BatteryPercentage}%";
    }

    public void Drive()
    {
        if (BatteryPercentage > 0)
        {
            Distance += 20;
            BatteryPercentage -= 1;
        }

    }
}
