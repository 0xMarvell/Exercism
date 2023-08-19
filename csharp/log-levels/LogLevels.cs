using System;

static class LogLine
{
    public static string Message(string logLine)
    {
        int targetIndex = logLine.IndexOf(':');
        string msg = logLine[(targetIndex + 1)..];

        return msg.Trim();
    }

    public static string LogLevel(string logLine)
    {
        string level = "";
        int targetIndex = logLine.IndexOf(':');
        string l = logLine[..targetIndex];
        char[] chars = l.ToCharArray();

        foreach (char c in chars)
        {
            if (Char.IsLetter(c))
            {
                level += c.ToString();
            }
        }

        return level.ToLower();
    }

    public static string Reformat(string logLine)
    {
        return $"{Message(logLine)} ({LogLevel(logLine)})";
    }
}
