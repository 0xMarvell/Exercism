using System;
using System.ComponentModel.Design;

class BirdCount
{
    private int[] birdsPerDay;
    private int length;

    public BirdCount(int[] birdsPerDay)
    {
        this.birdsPerDay = birdsPerDay;
        this.length = birdsPerDay.Length;
    }

    public static int[] LastWeek()
    {
        return new int[] {0, 2, 5, 3, 7, 8, 4};
    }

    public int Today()
    {
        return birdsPerDay[length - 1];
    }

    public void IncrementTodaysCount()
    {
        birdsPerDay[length - 1]++;
    }

    public bool HasDayWithoutBirds()
    {
        foreach (int bird in birdsPerDay)
        {
            if (bird == 0)
            {
                return true;
            }
        }

        return false;
    }

    public int CountForFirstDays(int numberOfDays)
    {
        int count = 0;

        for (int i = 0; i < numberOfDays; i++)
        {
            count += birdsPerDay[i];
        }

        return count;
    }

    public int BusyDays()
    {
        int count = 0;

        foreach (int bird in birdsPerDay)
        {
            if (bird >= 5)
            {
                count++;
            }
        }

        return count;
    }
}
