using System;

public static class LogAnalysis
{
    // TODO: define the 'SubstringAfter()' extension method on the `string` type
    public static string SubstringAfter(this string input, string delimiter)
    {
        int index = input.IndexOf(delimiter);

        if (index >= 0)
        {
            return input[(index + delimiter.Length)..];
        }

        return input;
    }

    // TODO: define the 'SubstringBetween()' extension method on the `string` type
    public static string SubstringBetween(this string input, string a, string b)
    {
        int start = input.IndexOf(a);
        int end = input.IndexOf(b, start + a.Length);

        if (start >= 0 && end >= 0)
        {
            return input.Substring((start + a.Length), (end - start - a.Length));
        }

        return string.Empty;
    }

    // TODO: define the 'Message()' extension method on the `string` type
    public static string Message(this string input)
    {
        int targetIndex = input.IndexOf(':');
        string msg = input[(targetIndex + 1)..];

        return msg.Trim();
    }

    // TODO: define the 'LogLevel()' extension method on the `string` type
    public static string LogLevel(this string input)
    {
        string level = "";
        int targetIndex = input.IndexOf(':');
        string l = input[..targetIndex];
        char[] chars = l.ToCharArray();

        foreach (char c in chars)
        {
            if (Char.IsLetter(c))
            {
                level += c.ToString();
            }
        }

        return level;
    }
}