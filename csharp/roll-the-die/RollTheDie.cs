using System;

public class Player
{
    public int RollDie()
    {
        Random rand = new();
        return rand.Next(1,19);
    }

    public double GenerateSpellStrength()
    {
        Random rand = new();
        return (double)rand.Next(100);
    }
}
