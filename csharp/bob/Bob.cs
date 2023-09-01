using System;
public static class Bob
{
    public static string Response(string statement)
    {
        if (string.IsNullOrWhiteSpace(statement))
        {
            return "Fine. Be that way!";
        }
        else if (IsYelling(statement) && statement.TrimEnd().EndsWith("?"))
        {
            return "Calm down, I know what I'm doing!";
        }
        else if (IsYelling(statement))
        {
            return "Whoa, chill out!";
        }
        else if (statement.TrimEnd().EndsWith("?"))
        {
            return "Sure.";
        }
        else
        {
            return "Whatever.";
        }
    }
    private static bool IsYelling(string statement)
    {
        bool hasLetters = false;
        foreach (char c in statement)
        {
            if (char.IsLetter(c))
            {
                hasLetters = true;
                if (!char.IsUpper(c))
                {
                    return false;
                }
            }
        }
        return hasLetters;
    }
}