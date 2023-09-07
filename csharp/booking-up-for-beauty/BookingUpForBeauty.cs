using System;

static class Appointment
{
    public static DateTime Schedule(string appointmentDateDescription)
    {
        return DateTime.Parse(appointmentDateDescription);
    }

    public static bool HasPassed(DateTime appointmentDate)
    {
        return appointmentDate < DateTime.Now;
    }

    public static bool IsAfternoonAppointment(DateTime appointmentDate)
    {
        DateTime afternoon = appointmentDate.Date.AddHours(12);
        DateTime evening = appointmentDate.Date.AddHours(18).AddSeconds(-1);

        return appointmentDate >= afternoon && appointmentDate <= evening;
    }

    public static string Description(DateTime appointmentDate)
    {
        var formattedDate = appointmentDate.ToString("M/d/yyyy h:mm:ss tt");
        return $"You have an appointment on {formattedDate}.";
    }

    public static DateTime AnniversaryDate()
    {
        return new DateTime(DateTime.Now.Year, 9, 15, 0, 0, 0);
    }
}
