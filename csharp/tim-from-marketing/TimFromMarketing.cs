using System;

static class Badge
{
    public static string Print(int? id, string name, string? department)
    {
        if (department == null && id == null)
        {
            return $"{name} - OWNER";
        }
        
        if (id == null)
        {
            return $"{name} - {department?.ToUpper()}";
        }
        else if (department == null)
        {
            return $"[{id}] - {name} - OWNER";
        }
        else
        {
            return $"[{id}] - {name} - {department?.ToUpper()}";
        }
    }
}
